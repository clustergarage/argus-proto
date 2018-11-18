# Proto file
get_filename_component(_ARGUS_PROTO "argus.proto" ABSOLUTE)
get_filename_component(_ARGUS_PROTO_PATH "${_ARGUS_PROTO}" PATH)

# Generated sources
set(_ARGUS_PROTO_SRCS "${CMAKE_CURRENT_BINARY_DIR}/argus.pb.cc")
set(_ARGUS_PROTO_HDRS "${CMAKE_CURRENT_BINARY_DIR}/argus.pb.h")
set(_ARGUS_GRPC_SRCS "${CMAKE_CURRENT_BINARY_DIR}/argus.grpc.pb.cc")
set(_ARGUS_GRPC_HDRS "${CMAKE_CURRENT_BINARY_DIR}/argus.grpc.pb.h")
add_custom_command(
  OUTPUT "${_ARGUS_PROTO_SRCS}" "${_ARGUS_PROTO_HDRS}" "${_ARGUS_GRPC_SRCS}" "${_ARGUS_GRPC_HDRS}"
  COMMAND ${_PROTOBUF_PROTOC}
  ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
    --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
    -I "${_ARGUS_PROTO_PATH}"
    --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
    "${_ARGUS_PROTO}"
  DEPENDS "${_ARGUS_PROTO}"
)
