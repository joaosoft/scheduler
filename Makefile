env:
	docker-compose up -d

run:
	go run ./main/main.go

build:
	docker build -t scheduler:1.0 .

push:
	docker login --username joaosoft
	docker tag scheduler:1.0 joaosoft/scheduler
	docker push joaosoft/scheduler

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*