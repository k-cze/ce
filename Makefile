.PHONY: local run docker-build docker-run docker-clean

local:
	@export $$(grep -v '^#' .env | xargs) && go run ./cmd/ce

run: local

docker-build:
	docker build -t ce .

docker-run:
	docker run --env-file .env -p 8080:8080 ce

docker-clean:
	docker rm -f $$(docker ps -a -q --filter ancestor=ce)
