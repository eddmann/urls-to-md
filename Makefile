IMAGE_NAME=eddmann/urls-to-md
DIST_OS=darwin linux windows
DIST_ARCH=amd64

image:
	docker build -t $(IMAGE_NAME) .

deps:
	docker run --rm -v "$(PWD)":/go/src/app $(IMAGE_NAME) dep ensure

shell:
	docker run --rm -it -v "$(PWD)":/go/src/app $(IMAGE_NAME) /bin/sh

build:
	docker run --rm -v "$(PWD)":/go/src/app $(IMAGE_NAME) /bin/sh -c ' \
		rm -fr ./dist/*; \
		for GOOS in $(DIST_OS); do \
			for GOARCH in $(DIST_ARCH); do \
				GOOS=$$GOOS GOARCH=$$GOARCH go build -o ./dist/urls-to-md-$$GOOS-$$GOARCH ./urls-to-md.go; \
			done; \
		done;'
