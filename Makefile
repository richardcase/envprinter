VAR_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VAR_REF=$(shell git rev-parse --short HEAD)
VAR_VER=$(shell cat ./VERSION)


all: clean copycontent build 

copycontent:
	mkdir -p ./build

build: buildapp buildimage

buildapp:
	CGO_ENABLED=0 GOOS=linux go build -a -o build/envprinter github.com/richardcase/envprinter/

buildimage: buildapps
	@echo Build Date: $(VAR_DATE) 
	@echo Git Ref: $(VAR_REF)
	@echo Version: $(VAR_VER)
	docker build -t richardcase/dockercoinsgo-rng -f Dockerfile-rng --build-arg BUILD_DATE=$(VAR_DATE) --build-arg VCS_REF=$(VAR_REF) --build-arg VERSION=$(VAR_VER) .

clean:
	rm -rf build/
