#!/usr/bin/make -f

export GO111MODULE = on

install:
	@echo "installing opb contract call service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/irita-opb-provider ./cmd
