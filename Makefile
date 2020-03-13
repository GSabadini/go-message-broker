#!make

up:
	docker-compose up -d

fmt:
	go fmt ./...

vet:
	go vet ./...