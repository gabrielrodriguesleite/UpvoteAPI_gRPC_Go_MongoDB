# FROM grpc/go

FROM alpine

RUN apk --no-cache add ca-certificates 

RUN apk add --no-cache tzdata

ENV TZ=America/Sao_Paulo

WORKDIR /app

COPY ./build/client .

ARG api_port

ARG api_host

# ENV API_PORT=${api_port}

# ENV API_HOST=${api_host}

# EXPOSE $PORT

CMD ./client --api_host ${api_host} --api_port ${api_port}

# reference https://christiangiacomi.com/posts/multi-stage-docker-image-go/