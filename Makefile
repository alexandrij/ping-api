ROOT=$(realpath $(dir $(realpath $(firstword $(MAKEFILE_LIST)))))

compile-pb:
	docker-compose -f docker-compose.yml up

run:
	go run cmd/app/app.go