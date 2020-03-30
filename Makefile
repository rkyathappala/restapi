# Service Name.
ifeq ($(OS), Windows_NT)
	NAME := restapi.exe
else
	NAME := restapi
endif

VERSION=${BUILD_VERSION}

# Attempt to determine the version infromation form git
ifeq ($(VERSION),)
	# Service version and build information drived from git.
	TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
	TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
	COMMIT := $(shell git rev-parse --short HEAD)
	DATE := $(shell git log -1 --format=%cd --date=format:"%Y%m%d")
	VERSION := $(TAG:v%=%)

	# If the tag commit is not empty, check if in the latest commit and use as version.
	ifneq ($(TAG_COMMIT),)
		ifneq ($(COMMIT), $(TAG_COMMIT))
			VERSION := v$(VERSION)-$(COMMIT)-$(DATE)
		endif
	endif

	# Else, if the version is empty use the commit commit as version (prefix with dev-).
	ifeq ($(VERSION),)
		VERSION := dev-$(COMMIT)-$(DATE)
	endif

	# If local, flag as dirty.
	ifneq ($(shell git status --porcelain),)
		VERSION := $(VERSION)-dirty
	endif
endif

all: clean restapi

restapi:
	@printf "Building ${NAME} ${VERSION} ${BUILD} ... "
	@go build .
	@printf "Done.\n"

new_image: clean restapi
	@rm -f image/$(NAME)
	@mv $(NAME) image/$(NAME)
	@echo "${NAME} ${VERSION} ${BUILD} is now tracked in image/"

clean:
	@-rm -f $(NAME)
