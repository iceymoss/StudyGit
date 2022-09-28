cd /Users/feng/src/grpc_gateway/server/trip/api

protoc -I . trip.proto --go_out=plugins=grpc:.

protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:. trip.proto