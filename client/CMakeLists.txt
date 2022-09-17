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