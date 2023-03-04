FROM golang:1.18-alpine AS builder

WORKDIR /polygomic-edge

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o polygomic-edge main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /polygomic-edge

COPY --from=builder /polygomic-edge/polygomic-edge ./
COPY ./docker/local/polygomic-edge.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 30342 5001

ENTRYPOINT ["./polygomic-edge.sh"]
