purge:
	rm -rf bin/ tmp/ logs/ client/build client/node_modules/

clean:
	go clean
	go mod tidy

build-api:
	go get
	go mod tidy
	go build -o bin/app/go-chat.out

build-client:
	cd client && npm install && rm -rf build && npm run build

run-api:
	./bin/app/go-chat.out

deploy: purge build-api build-client run-api

deploy-api: build-api run-api

kill:
	# cat ./tmp/server/pid | xargs kill -9
	npx kill-port 3000
	npx kill-port 3001

dev:
	make -j 2  client server
.PHONY: dev

server:
	go get
	go mod tidy
	air
.PHONY: server

client:
	cd client && yarn start
.PHONY: client

