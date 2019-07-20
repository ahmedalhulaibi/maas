tag=$(shell git describe --tags | sed "s/v//g")
gitroot=$(shell git rev-parse --show-toplevel)

default: build

clean:
	@echo "Cleaning"
	sleep 1
	@echo "Clean done"

build: clean
	docker build . -t ahmedalhulaibi/maas:latest
	@echo "Build complete"