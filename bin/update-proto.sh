#!/bin/bash

protoc -I. argus.proto --go_out=plugins=grpc:golang
protoc -I. argus.proto --grpc_out=c++ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. argus.proto --cpp_out=c++
protoc -I. health.proto --grpc_out=c++ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. health.proto --cpp_out=c++
