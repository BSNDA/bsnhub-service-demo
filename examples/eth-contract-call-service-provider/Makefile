#!/usr/bin/make -f

export GO111MODULE = on

install:
	@echo "installing eth contract call service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/irita-eth-provider ./cmd
