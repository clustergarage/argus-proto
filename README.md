# fim-proto

## Build

```
protoc -I fim/ fim.proto --go_out=plugins=grpc:fim --cpp_out=fim
```
