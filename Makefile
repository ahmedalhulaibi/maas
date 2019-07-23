tag=$(shell git describe --tags | sed "s/v//g")
gitroot=$(shell git rev-parse --show-toplevel)
PLATFORM := $(shell ./osname.sh)

default: build

.PHONY: clean
clean:
	@echo "Clean started"
	@echo $(PLATFORM)
	faas-cli remove -f maas-faas.yml
	@echo "Clean complete"

.PHONY: run
run:
	@echo "Run started"
	docker run -v /var/run/docker.sock:/var/run/docker.sock -it ahmedalhulaibi/maas:latest https://github.com/ahmedalhulaibi/maas.git install-tools build
	@echo "Run complete"

verify-tools:
	@echo "Verifying tools are installed"
	$(if $(shell PATH=$(PATH) which faas-cli),,$(error "No faas-cli in PATH. Use install-tools/faas-cli.sh"))
	@echo "Done verifying tools are installed"

.PHONY: install-tools
install-tools:
	@echo "Installing tools. Platform identified as $(PLATFORM)"
	@echo "Installing faas-cli"
	$(shell install-tools/faas-cli.sh $(PLATFORM))
    ifeq ($(PLATFORM),Alpine)
    endif
    ifeq ($(PLATFORM),Debian)
    endif 
    ifeq ($(PLATFORM),Darwin)
    endif
	@echo "Done installing tools"

dep: verify-tools
	@faas-cli template pull https://github.com/openfaas-incubator/golang-http-template

Dockerfile: clean dep
	@echo "Building image from Dockerfile"
	docker build . -t ahmedalhulaibi/maas:latest
	@echo "Building image complete"

maas-faas.yml: clean dep
	@echo "Building openfaas fn from maas-faas.yml"
	faas-cli build -f ./maas-faas.yml
	@echo  "Building openfaas fn complete"

build: Dockerfile maas-faas.yml

publish-image: build
	@echo "Publishing image to dockerhub"
	docker push ahmedalhulaibi/maas:latest
	@echo "Publishing openfaas fn"
	@echo "Publishing complete"