# Include the help snippet from above...
help: ## Show this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

all:
	go build webhook-receiver.go
	docker build -t webhook-receiver .

docker-run:
	docker run -p 8080:8080 webhook-receiver

deploy: ## deploys
	kubectl apply -f webhook-receiver-deployment.yml

service: ## create the service and exposes the service
	kubectl apply -f my-service.yml