
FROM golang:latest as builder

ENV GO111MODULE=auto
#ENV SERVER_HOST=grpc-server-comp:50051

WORKDIR /go/src/github.com/maei/golang_grpc_client_streaming_comp/grpc_client

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc-client-comp /go/src/github.com/maei/golang_grpc_client_streaming_comp/grpc_client/src/main.go

FROM alpine
COPY --from=builder /go/src/github.com/maei/golang_grpc_client_streaming_comp/grpc_client /
CMD ["./grpc-client-comp"]