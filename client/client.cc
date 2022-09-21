// Client side implementation of UDP client-server model
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <cstring>

#include "sensor.pb.h"
#include "sensorstate.pb.h"

#define SSPORT 8080
#define SDPORT 8081
#define MAXLINE 1024

int createSocket()
{
    int sockfd;
    if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) < 0)
    {
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }
    return sockfd;
}

sockaddr_in fillServerInfo(uint16_t port)
{
    struct sockaddr_in servaddr;
    memset(&servaddr, 0, sizeof(servaddr));

    servaddr.sin_family = AF_INET;
    servaddr.sin_port = htons(port);
    servaddr.sin_addr.s_addr = INADDR_ANY;

    return servaddr;
}

pb::SensorState *createNewSensorState()
{
    auto sensor_state = new pb::SensorState();
    auto main_computer = new pb::MainComputer();
    auto brake_manager = new pb::BrakeManager();

    main_computer->set_state(pb::MainComputer::END_BRAKES);
    brake_manager->set_state(pb::BrakeManager::BRAKING);
    brake_manager->set_critical_pod_deceleration_instruction_timeout(true);
    brake_manager->set_critical_emergency_brakes_without_deceleration(true);
    brake_manager->set_mesured_distance_greater_as_desired(true);

    sensor_state->set_allocated_main_computer(main_computer);
    sensor_state->set_allocated_brake_manager(brake_manager);

    return sensor_state;
}

pb::SensorData *createNewSensorData()
{
    auto sensor_data = new pb::SensorData();
    sensor_data->set_sensor_id(1);
    sensor_data->set_value(34.0);
    return sensor_data;
}

void sendSensorState(int sockfd, sockaddr_in &servaddr, pb::SensorState *sensor_state)
{
    std::string serialized = sensor_state->SerializeAsString();
    sendto(sockfd, serialized.c_str(), serialized.length(),
           MSG_CONFIRM, (const struct sockaddr *)&servaddr,
           sizeof(servaddr));
}

void sendSensorData(int sockfd, sockaddr_in &servaddr, pb::SensorData *sensor_data)
{
    std::string serialized = sensor_data->SerializeAsString();
    sendto(sockfd, serialized.c_str(), serialized.length(),
           MSG_CONFIRM, (const struct sockaddr *)&servaddr,
           sizeof(servaddr));
}

int main()
{
    int ss_socket = createSocket();
    int sd_socket = createSocket();
    struct sockaddr_in ssaddr = fillServerInfo(SSPORT);
    struct sockaddr_in sdaddr = fillServerInfo(SDPORT);

    double duration;
    double interval = 0.1;
    uint count = 0;
    std::clock_t start = std::clock();
    std::clock_t start_interval = std::clock();

    while (duration <= 2.0)
    {
        if ((std::clock() - start_interval) / (double)CLOCKS_PER_SEC >= interval)
        {

            pb::SensorState *sensor_state = createNewSensorState();
            pb::SensorData *sensor_data = createNewSensorData();
            sendSensorState(ss_socket, ssaddr, sensor_state);
            sendSensorData(sd_socket, sdaddr, sensor_data);
            start_interval = std::clock();
            count++;
            std::cout << "printf: " << duration << " count: " << count << '\n';
        }
        duration = (std::clock() - start) / (double)CLOCKS_PER_SEC;
    }

    close(ss_socket);
    close(sd_socket);
    return 0;
}