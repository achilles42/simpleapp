.PHONY: all
all: build fmt vet lint dev

.PHONY: ci
ci: build fmt vet lint copy-test-config test


APP_EXECUTABLE="./simpleapp"

compile:
	go build -o $(APP_EXECUTABLE)

build: build-deps compile fmt vet lint

install:
	go install ./...

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

