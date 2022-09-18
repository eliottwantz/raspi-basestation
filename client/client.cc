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
#include <iostream>

#include "sensor.pb.h"
#include "sensorstate.pb.h"

#define PORT 8080
#define MAXLINE 1024

// Driver code
int main()
{
    int sockfd;
    struct sockaddr_in servaddr;

    // Creating socket file descriptor
    if ((sockfd = socket(AF_INET, SOCK_DGRAM, 0)) < 0)
    {
        perror("socket creation failed");
        exit(EXIT_FAILURE);
    }

    memset(&servaddr, 0, sizeof(servaddr));

    // Filling server information
    servaddr.sin_family = AF_INET;
    servaddr.sin_port = htons(PORT);
    servaddr.sin_addr.s_addr = INADDR_ANY;

    // Creating protobuf message
    auto sensor_state = new pb::SensorState();
    auto main_computer = new pb::MainComputer();
    auto brake_manager = new pb::BrakeManager();

    main_computer->set_state(pb::MainComputer_States_ACCELERATING);
    brake_manager->set_state(pb::BrakeManager_States_BRAKING);
    brake_manager->set_hydrolic_pressure_loss(false);
    brake_manager->set_critical_pod_acceleration_mesure_timeout(true);
    brake_manager->set_critical_pod_deceleration_instruction_timeout(false);
    brake_manager->set_verin_blocked(false);
    brake_manager->set_emergency_valve_open_without_hydrolic_pressor_diminution(false);
    brake_manager->set_critical_emergency_brakes_without_deceleration(false);
    brake_manager->set_mesured_distance_less_than_desired(false);
    brake_manager->set_mesured_distance_greater_as_desired(true);

    sensor_state->set_allocated_main_computer(main_computer);
    sensor_state->set_allocated_brake_manager(brake_manager);

    std::string serialized;
    sensor_state->SerializeToString(&serialized);

    sendto(sockfd, serialized.c_str(), serialized.length(),
           MSG_CONFIRM, (const struct sockaddr *)&servaddr,
           sizeof(servaddr));

    close(sockfd);
    return 0;
}
