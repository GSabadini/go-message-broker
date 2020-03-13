#!make

fmt:
	go fmt ./...

vet:
	go vet ./...

up-all:
	docker-compose up -d

down-all:
	docker-compose down

## RabbitMQ

up-rabbitmq:
	docker-compose up -d rabbitmq

down-rabbitmq:
	docker-compose down -d rabbitmq

publish-rabbitmq:
	go run main.go rabbitmq producer

consumer-rabbitmq:
	go run main.go rabbitmq consumer

## Kafka

up-kafka:
	docker-compose up -d zookeeper kafka

down-kafka:
	docker-compose down -d zookeeper kafka

publish-kafka:
	go run main.go kafka producer

consumer-kafka:
	go run main.go kafka consumer

