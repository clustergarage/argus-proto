# fim-proto

## Build

```
# build golang definitions
protoc -I. fim.proto --go_out=plugins=grpc:.

# build c++ definitions
protoc -I. fim.proto --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin`
protoc -I. fim.proto --cpp_out=.
```

### Generates

* `fim.pb.h` - the header which declares your generated message classes
* `fim.pb.cc` - which contains the implementation of your message classes
* `fim.grpc.pb.h` - the header which declares your generated service classes
* `fim.grpc.pb.cc` - which contains the implementation of your service classes
