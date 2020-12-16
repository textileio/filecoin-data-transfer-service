.DEFAULT_GOAL=install

include .bingo/Variables.mk

FTS_BUILD_FLAGS?=CGO_ENABLED=0
FTS_VERSION?="git"
GOVVV_FLAGS=$(shell $(GOVVV) -flags -version $(FTS_VERSION) -pkg $(shell go list ./buildinfo))

build: $(GOVVV)
	$(FTS_BUILD_FLAGS) go build -ldflags="${GOVVV_FLAGS}" ./...
.PHONY: build

install: $(GOVVV)
	$(FTS_BUILD_FLAGS) go install -ldflags="${GOVVV_FLAGS}" ./...
.PHONY: install

define gen_release_files
	$(GOX) -osarch=$(3) -output="build/$(2)/$(2)_${FTS_VERSION}_{{.OS}}-{{.Arch}}/$(2)" -ldflags="${GOVVV_FLAGS}" $(1)
	mkdir -p build/dist; \
	cd build/$(2); \
	for release in *; do \
		cp ../../LICENSE ../../README.md $${release}/; \
		if [ $${release} != *"windows"* ]; then \
  		POW_FILE=$(2) $(GOMPLATE) -f ../../dist/install.tmpl -o "$${release}/install"; \
			tar -czvf ../dist/$${release}.tar.gz $${release}; \
		else \
			zip -r ../dist/$${release}.zip $${release}; \
		fi; \
	done
endef

build-fts-release: $(GOX) $(GOVVV) $(GOMPLATE)
	$(call gen_release_files,./fts,fts,"linux/amd64 linux/386 linux/arm darwin/amd64 windows/amd64")
.PHONY: build-fts-release

build-releases: build-fts-release
.PHONY: build-releases

test:
	go test -short -p 2 -parallel 2 -race -timeout 45m ./... 
.PHONY: test
