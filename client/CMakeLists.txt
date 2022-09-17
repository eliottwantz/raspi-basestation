cmake_minimum_required(VERSION 3.0.0)
project(myproj VERSION 0.1.0)


# Find Protobuf installation
  # Looks for protobuf-config.cmake file installed by Protobuf's cmake installation.
set(protobuf_MODULE_COMPATIBLE TRUE)
find_package(Protobuf CONFIG REQUIRED)
message(STATUS "Using protobuf ${Protobuf_VERSION}")
set(_PROTOBUF_LIBPROTOBUF protobuf::libprotobuf)
if(CMAKE_CROSSCOMPILING)
find_program(_PROTOBUF_PROTOC protoc)
else()
set(_PROTOBUF_PROTOC $<TARGET_FILE:protobuf::protoc>)
endif()

# # Proto file
# get_filename_component(protos "../pb/sensor.proto" ABSOLUTE)
# get_filename_component(sensor_proto_path "${protos}" PATH)

# # Generated sources
# set(sensor_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/sensor.pb.cc")
# set(sensor_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/sensor.pb.h")
# add_custom_command(
#       OUTPUT "${sensor_proto_srcs}" "${sensor_proto_hdrs}"
#       COMMAND ${_PROTOBUF_PROTOC}
#       ARGS --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
#         -I "${sensor_proto_path}"
#         "${protos}"
#       DEPENDS "${protos}")

# include_directories(${Protobuf_INCLUDE_DIRS})
include_directories(${CMAKE_CURRENT_BINARY_DIR})

# protos
file(GLOB ProtoFiles "../pb/*.proto")
protobuf_generate_cpp(PROTO_SRCS PROTO_HDRS ${ProtoFiles})
add_library(protos ${PROTO_SRCS} ${PROTO_HDRS})
target_link_libraries(protos ${_REFLECTION} ${_PROTOBUF_LIBPROTOBUF})


# Targets service_client
foreach(_target
  client)
  add_executable(${_target}
    "${_target}.cc")
  target_link_libraries(${_target}
    protos
    ${_REFLECTION}
    ${_PROTOBUF_LIBPROTOBUF})
endforeach()