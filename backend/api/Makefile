NAME=api
TAG=latest
IMAGE_NAME=$(NAME):$(TAG)
OUT=$(NAME)

# Enrich the commands below with appropriate logic if necessary
# Introducing a shell script(e.g.into ./scripts/ folder) could be concidered as a good practice

.PHONY: install-deps
install-deps:
	# use your chosen dependency management system
	@go get -d ./...

.PHONY: lint
lint:
	 @gometalinter ./...

.PHONY: vet
vet:
	 go vet ./...

.PHONY: build
build:
	 @go build -o bin/$(OUT) .

.PHONY: install
install: test
	 @go install ./...

.PHONY: test
test: vet
	@go test -race ./...

.PHONY: docker-build
docker-build:
	@docker build -t $(IMAGE_NAME) .
