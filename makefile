.PHONY: protos

protos:
   protoc -I proto myfile.proto --go_out=plugins=grpc:proto
   
   
