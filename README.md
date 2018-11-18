# argus-proto

## Build

```
# build golang definitions
protoc -I. argus.proto --go_out=plugins=grpc:golang

# build c++ definitions
protoc -I. argus.proto --grpc_out=c++ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. argus.proto --cpp_out=c++

# [optional] build c++ health definitions; golang has these readily available
protoc -I. health.proto --grpc_out=c++ --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. health.proto --cpp_out=c++
```

### Generates

* `argus.pb.h` - the header which declares your generated message classes
* `argus.pb.cc` - which contains the implementation of your message classes
* `argus.grpc.pb.h` - the header which declares your generated service classes
* `argus.grpc.pb.cc` - which contains the implementation of your service classes

## Generate Mocks

```
mockgen github.com/clustergarage/argus-proto/golang \
  ArgusdClient,Argusd_GetWatchStateClient,Argusd_RecordMetricsClient > \
  golang/mock/argus_mock.pb.go
```
