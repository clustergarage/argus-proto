# fim-proto

## Build

```
# build golang definitions
protoc -I. fim.proto --go_out=plugins=grpc:golang

# build c++ definitions
protoc -I. fim.proto --grpc_out=c++ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. fim.proto --cpp_out=c++

# [optional] build c++ health definitions; golang has these readily available
protoc -I. health.proto --grpc_out=c++ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. health.proto --cpp_out=c++
```

### Generates

* `fim.pb.h` - the header which declares your generated message classes
* `fim.pb.cc` - which contains the implementation of your message classes
* `fim.grpc.pb.h` - the header which declares your generated service classes
* `fim.grpc.pb.cc` - which contains the implementation of your service classes

## Generate Mocks

```
mockgen github.com/clustergarage/fim-proto/golang \
  FimdClient,Fimd_GetWatchStateClient,Fimd_RecordMetricsClient > \
  golang/mock/fim_mock.pb.go
```
