CHALLENGE_COMPLEXITY ?= 4
SERVER_ADDR ?= localhost:9000

run-server:
	docker build --tag go-tcs-pow-server:dev -f cmd/server/Dockerfile .
	docker run --env CHALLENGE_COMPLEXITY=$(CHALLENGE_COMPLEXITY) -p 9000:9000 go-tcs-pow-server:dev

run-client:
	docker build --tag go-tcs-pow-client:dev -f cmd/client/Dockerfile .
	docker run --env SERVER_ADDR=$(SERVER_ADDR) --network="host" go-tcs-pow-client:dev
