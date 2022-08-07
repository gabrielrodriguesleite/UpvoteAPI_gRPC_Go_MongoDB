# FROM grpc/go

FROM alpine

RUN apk --no-cache add ca-certificates 

RUN apk add --no-cache tzdata

ENV TZ=America/Sao_Paulo

WORKDIR /app

COPY ./build/server .

ARG port

ENV PORT=${port}

EXPOSE $PORT

CMD ["./server"]

# reference https://christiangiacomi.com/posts/multi-stage-docker-image-go/