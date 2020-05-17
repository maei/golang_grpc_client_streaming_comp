package main

import (
	"github.com/maei/golang_grpc_client_streaming_comp/grpc_client/src/service"
	"github.com/maei/shared_utils_go/logger"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logger.Info("Starting GRPC-Client")
	time.Sleep(5 * time.Second)
	service.AverageService.Average()
}
