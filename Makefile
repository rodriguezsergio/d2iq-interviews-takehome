GIT_HASH := $(shell git rev-parse --short HEAD)
LOCAL_TAG_NAME = d2iq/app
LOCAL_TAG_NAME_FULL = $(LOCAL_TAG_NAME):$(GIT_HASH)

build:
	docker build -t "$(LOCAL_TAG_NAME_FULL)" .

run: build
	docker run --rm --name d2iq -d -p 8080:8080 "$(LOCAL_TAG_NAME_FULL)"

stop:
	docker ps | grep d2iq | awk '{print $$1}' | xargs docker kill
