#check_install:
#   which_swagger||GO111MODULE=off go get -u go get -u github.com/swaggo/swag/cmd/swag
#
#swagger: check_install
#   GO111MODULE=off swagger generate spec -o ./swagger.yaml  --scan-models

.PHONY: protos

protos:
	 protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
