# Variables
APP_NAME = yamlres
IMAGE_NAME = ghcr.io/andygodish/$(APP_NAME)
TAG = dev
DOCKER_BUILD_ARGS =
DOCKER_RUN_ARGS =

.PHONY: k3d-create docker-build docker-run docker-stop docker-push docker-clean docker-build-sha docker-build-multiarch help

# Create a k3d dev cluster
k3d-create:
	@echo "Creating k3d cluster..."
    k3d cluster create $(APP_NAME) \
		--agents 2 \
		--k3s-arg "--disable=traefik@server:0" \
		--k3s-arg "--disable=local-storage@server:0" || true

# Build the Docker image
docker-build:
	@echo "Building Docker image $(IMAGE_NAME):$(TAG)..."
	docker build $(DOCKER_BUILD_ARGS) -t $(IMAGE_NAME):$(TAG) .

# Build multi-architecture Docker image
docker-build-multiarch:
	@echo "Building multi-arch Docker image $(IMAGE_NAME):$(TAG)..."
	docker buildx create --use --name multiarch-builder || true
	docker buildx build --platform linux/amd64,linux/arm64 $(DOCKER_BUILD_ARGS) -t $(IMAGE_NAME):$(TAG) --push .

# Run the container locally
# example: 
#   make docker-run DOCKER_RUN_ARGS="-v $(pwd)/resume.yaml:/app/config/resume.yaml"
docker-run:
	@echo "Running container from $(IMAGE_NAME):$(TAG)..."
	docker run -d --name $(APP_NAME) -p 8080:8080 $(DOCKER_RUN_ARGS) $(IMAGE_NAME):$(TAG)

# Stop and remove the running container
docker-stop:
	@echo "Stopping container..."
	-docker stop $(APP_NAME)
	-docker rm $(APP_NAME)

# Push the image to GHCR
docker-push:
	@echo "Pushing $(IMAGE_NAME):$(TAG) to GHCR..."
	docker push $(IMAGE_NAME):$(TAG)

# Clean up local Docker images for this project
docker-clean:
	@echo "Removing Docker images for $(APP_NAME)..."
	-docker rmi $(IMAGE_NAME):$(TAG)

# Build and tag with git SHA (useful for CI)
docker-build-sha:
	$(eval GIT_SHA := $(shell git rev-parse --short HEAD))
	@echo "Building Docker image $(IMAGE_NAME):$(GIT_SHA)..."
	docker build $(DOCKER_BUILD_ARGS) -t $(IMAGE_NAME):$(GIT_SHA) .
	docker tag $(IMAGE_NAME):$(GIT_SHA) $(IMAGE_NAME):latest

# Help command
help:
	@echo "Available targets:"
	@echo "  docker-build       - Build the Docker image with tag 'dev'"
	@echo "  docker-build-multiarch - Build multi-arch Docker image (amd64+arm64)"
	@echo "  docker-run         - Run the container locally"
	@echo "  docker-stop        - Stop and remove the running container"
	@echo "  docker-push        - Push the image to GHCR"
	@echo "  docker-clean       - Remove Docker images for this project"
	@echo "  docker-build-sha   - Build and tag with git SHA (useful for CI)"