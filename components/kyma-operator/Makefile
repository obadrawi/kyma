APP_NAME = kyma-operator
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)
binary=$(APP_NAME)

.PHONY: build
build:
	./before-commit.sh ci

.PHONY: pull-licenses
pull-licenses:
ifdef LICENSE_PULLER_PATH
	bash $(LICENSE_PULLER_PATH)
else
	mkdir -p licenses
endif

.PHONY: build-image
build-image: pull-licenses
	cp $(binary) deploy/kyma-operator/$(binary)
	cp -R licenses deploy/kyma-operator/
	docker build -t $(APP_NAME):latest deploy/kyma-operator

.PHONY: push-image
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)

.PHONY: ci-pr
ci-pr: build build-image push-image

.PHONY: ci-master
ci-master: build build-image push-image

.PHONY: ci-release
ci-release: build build-image push-image

.PHONY: clean
clean:
	rm -f $(binary)
