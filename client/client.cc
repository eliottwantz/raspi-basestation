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

#define PORT 8080
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

sockaddr_in fillServerInfo()
{
    struct sockaddr_in servaddr;
    memset(&servaddr, 0, sizeof(servaddr));

    servaddr.sin_family = AF_INET;
    servaddr.sin_port = htons(PORT);
    servaddr.sin_addr.s_addr = INADDR_ANY;

    return servaddr;
}

pb::SensorState *createNewProtobufs()
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

void sendProtobuf(int sockfd, sockaddr_in &servaddr, pb::SensorState *sensor_state)
{
    std::string serialized = sensor_state->SerializeAsString();
    sendto(sockfd, serialized.c_str(), serialized.length(),
           MSG_CONFIRM, (const struct sockaddr *)&servaddr,
           sizeof(servaddr));
}

int main()
{
    int sockfd = createSocket();
    struct sockaddr_in servaddr = fillServerInfo();

    pb::SensorState *sensor_state = createNewProtobufs();

    sendProtobuf(sockfd, servaddr, sensor_state);

    close(sockfd);
    return 0;
}