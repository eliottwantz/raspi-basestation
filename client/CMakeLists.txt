cmake_minimum_required(VERSION 3.0.0)
project(myproj VERSION 0.1.0)



# Proto file
get_filename_component(sensor_proto "../pb/sensor.proto" ABSOLUTE)
get_filename_component(sensor_proto_path "${sensor_proto}" PATH)

# Generated sources
set(sensor_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/../pb/sensor.pb.cc")
set(sensor_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/../pb/sensor.pb.h")
# set(service_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/service.grpc.pb.cc")
# set(service_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/service.grpc.pb.h")
# add_custom_command(
#       OUTPUT "${sensor_proto_srcs}" "${sensor_proto_hdrs}"
#       COMMAND ${_PROTOBUF_PROTOC}
#       ARGS --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
#         -I "${sensor_proto_path}"
#         "${sensor_proto}"
#       DEPENDS "${sensor_proto}")

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}/../pb")

# service_grpc_proto
add_library(service_grpc_proto
  ${sensor_proto_srcs}
  ${sensor_proto_hdrs})
target_link_libraries(service_grpc_proto
  # ${_REFLECTION}
  # ${_GRPC_GRPCPP}
  # ${_PROTOBUF_LIBPROTOBUF}
  )


# Targets service_client
  add_executable(client
    "client.cc")
  target_link_libraries(${_target}
    service_grpc_proto
    )