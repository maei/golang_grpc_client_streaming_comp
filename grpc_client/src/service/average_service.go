package service

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_client_streaming_comp/grpc_client/src/client"
	"github.com/maei/golang_grpc_client_streaming_comp/grpc_client/src/domain/averagepb"
	"github.com/maei/shared_utils_go/logger"
	"log"
	"math/rand"
)

var AverageService averageServiceInterface = &averageService{}

type averageServiceInterface interface {
	Average()
}

type averageService struct{}

func generateRandom() int32 {
	min := 1
	max := 10

	n := min + rand.Intn(max-min+1)
	return int32(n)
}

func (*averageService) Average() {
	conn, err := client.GRPCClient.SetClient()
	if err != nil {
		logger.Error("Error while creating grpc-client", err)
	}
	defer conn.Close()

	c := averagepb.NewAverageServiceClient(conn)

	stream, streamErr := c.ComputeAverage(context.Background())
	if streamErr != nil {
		logger.Error("Error while streaming data to grpc-server", streamErr)
	}

	for i := 0; i < 5; i++ {
		req := &averagepb.AverageRequest{
			Input: generateRandom(),
		}
		log.Println(req)
		stream.Send(req)
	}
	response, respError := stream.CloseAndRecv()
	if respError != nil {
		logger.Error("Error while fetching response from grpc-server", respError)
	}
	fmt.Printf("Response from GRPC-Server: %v\n", response)
}
