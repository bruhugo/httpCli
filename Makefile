TARGET = main

.PHONY: all vet fmt clear

all: vet fmt
	go build -o ${TARGET}

vet:
	staticcheck ./...

fmt: 
	go fmt ./...

clear:
	rm ${TARGET}
