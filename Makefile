GOCMD=go
GOTEST=$(GOCMD) test

test:
	$(GOTEST) -v ./...

bench:
	$(GOTEST) -v ./... -bench .
