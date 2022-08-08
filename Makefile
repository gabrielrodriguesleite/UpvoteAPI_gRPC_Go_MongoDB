# run `make` to generate protobuf and grpc
generate:
	@clear
	@echo "\n ----- Build Protobuf ----- \n\c"
	@echo $(PATH) | grep -i $(HOME)/go/bin --count > /dev/null && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/*.proto && echo "OK\n\c" || echo "Error\n\c"

build-server:
	@clear
	@echo "\n ----- Build Server ----- \n\c"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o server/build/server server/*.go && echo "\nOK\n\c" || echo "\nFail to build\n\c"
	@docker build --no-cache --build-arg port=50051 -t up_vote_server server/.

build-client:
	@clear
	@echo "\n ----- Build Client ----- \n\c"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o client/build/client client/*.go && echo "\nOK\n\c" || echo "\nFail to build\n\c"
	@docker build --no-cache --build-arg port=50051 -t up_vote_client client/.

# --new=host # solução de : https://stackoverflow.com/questions/64139243/dial-tcp-127-0-0-19091-connect-connection-refused
# @docker run -d --rm --net=host up_vote_server
run-docker-server:
	@clear
	@echo "\n ----- Running docker server ----- \n\c"
	@docker run -d --rm --net=host --publish 50051:50051 up_vote_server

run-docker-client:
	@clear
	@echo "\n ----- Running docker client ----- \n\c"
	@docker run --rm --net=host up_vote_client

all: \
	build-server \
	build-client \
	run-docker-server \
	run-docker-client \
