all:
	go build webhook-receiver.go
	docker build -t webhook-receiver .

kubernetes-start:
	colima start --memory 8 --kubernetes

kubernetes-stop:
	colima stop

docker-run:
	docker run -p 8080:8080 webhook-receiver

deploy:
	kubectl apply -f webhook-receiver-deployment.yml

service:
	kubectl apply -f my-service.yml

check:
	kubectl version 
