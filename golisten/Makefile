APP_NAME=quay.io/cporter/listen
EXEC_NAME=listen
BUILD_NUMBER_FILE=buildnum.txt

.PHONY: $(BUILD_NUMBER_FILE)

$(BUILD_NUMBER_FILE):
	@if ! test -f $(BUILD_NUMBER_FILE); then echo 1000 > $(BUILD_NUMBER_FILE); fi 
	@echo $$(($$(cat $(BUILD_NUMBER_FILE)) + 1)) > $(BUILD_NUMBER_FILE)

## Build the container
image: TAG = 1.$(shell cat $(BUILD_NUMBER_FILE))

image: $(BUILD_NUMBER_FILE)
	$(info Starting image build with tag $(TAG))
	docker build -t $(APP_NAME):$(TAG) .
	docker push $(APP_NAME):$(TAG)
	## rewrite YAML
	sed -E -i .bak 's/:[0-9]\.[0-9]{4}/:$(TAG)/g' listen.yaml

## Build the container without caching
image-nc: TAG = 1.$(shell cat $(BUILD_NUMBER_FILE))

image-nc: $(BUILD_NUMBER_FILE)
	$(info Starting image build with tag $(TAG))
	docker build --no-cache -t $(APP_NAME):$(TAG) .
	docker push $(APP_NAME):$(TAG)

## Run container on port 5000
run: 
	docker run -p 80:8080/tcp --name $(EXEC_NAME) $(APP_NAME):$(TAG) 

## deploy using kubernetes
deploy: 
	oc create namespace listen || echo "namespace already created"
	oc delete -f playbooks.yaml || echo "configmap not found"
	oc create -f playbooks.yaml
	oc delete -f listen.yaml || echo "deployment not found"
	oc create -f listen.yaml

exec:
	POD=$(shell oc -n listen get pod | grep listen | awk '{ print $1 }')

