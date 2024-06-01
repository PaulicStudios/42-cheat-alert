IMAGE_NAME = paulic/42-cheat-alert
IMAGE_TAG = latest

run: build
	docker run --rm --name 42-cheat-alert $(IMAGE_NAME):$(IMAGE_TAG)

build:
	docker buildx build -t $(IMAGE_NAME):$(IMAGE_TAG) .

push: build
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
