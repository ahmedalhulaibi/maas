tag=$(shell git describe --tags | sed "s/v//g")
gitroot=$(shell git rev-parse --show-toplevel)

default: build

clean:
	@echo "Cleaning"
	sleep 1
	@echo "Clean done"

test:
	docker run -v /var/run/docker.sock:/var/run/docker.sock -it ahmedalhulaibi/maas:latest https://github.com/ahmedalhulaibi/maas.git

build: clean
	docker build . -t ahmedalhulaibi/maas:latest
	@echo "Build complete"