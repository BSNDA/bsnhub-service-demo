#!/usr/bin/make -f

export GO111MODULE = on

install:
	@echo "installing opb contract call service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/irita-opb-provider github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/cmd
