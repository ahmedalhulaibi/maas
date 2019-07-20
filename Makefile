tag=$(shell git describe --tags | sed "s/v//g")
gitroot=$(shell git rev-parse --show-toplevel)
PLATFORM := $(shell ./osname.sh)

default: build

clean:
	@echo "Clean started"
	sleep 1
	@echo $(PLATFORM)
	@echo "Clean complete"

test:
	@echo "Test started"
	docker run -v /var/run/docker.sock:/var/run/docker.sock -it ahmedalhulaibi/maas:latest https://github.com/ahmedalhulaibi/maas.git
	@echo "Test complete"

dep:
	@echo "Dependencies installation started"
	@echo "OS=$(PLATFORM)"
	ifeq($(PLATFORM),Alpine)
		apk add --no-cache --update jq zip
	endif
	@echo "Dependencies installation complete"

build: clean dep
	@echo "Build started"
	docker build . -t ahmedalhulaibi/maas:latest
	@echo "Build complete"

publish: build
	@echo "Publishing image to dockerhub"
	docker push ahmedalhulaibi/maas:latest
	@echo "Publishing complete"