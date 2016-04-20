#!/bin/bash

protoc --proto_path=. --go_out=mathservice mathservice.proto
protoc --go_out=plugins=grpc:mathservice mathservice.proto
