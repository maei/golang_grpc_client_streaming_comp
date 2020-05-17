package server

import (
	"github.com/maei/golang_grpc_client_streaming_comp/grpc_server/src/domain/averagepb"
	"github.com/maei/golang_grpc_client_streaming_comp/grpc_server/src/service"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

type server struct{}

var (
	s = grpc.NewServer()
)

func (*server) ComputeAverage(stream averagepb.AverageService_ComputeAverageServer) error {
	var result []int32
	for {
		req, errRec := stream.Recv()
		if errRec == io.EOF {
			res := service.AverageService.Average(result)

			response := &averagepb.AverageResponse{
				Result: res,
			}
			return stream.SendAndClose(response)
		}
		if errRec != nil {
			logger.Error("error while fetiching grpc-client stream", errRec)
		}
		result = append(result, req.GetInput())
	}
}

func StartGRPCServer() {
	logger.Info("gRPC greet-streaming started")

	lis, err := net.Listen("tcp", os.Getenv("SERVER_PORT"))
	if err != nil {
		logger.Error("error while listening gRPC Server", err)
	}
	//greetpb.RegisterGreetServiceServer(s, &server{})
	averagepb.RegisterAverageServiceServer(s, &server{})

	errServer := s.Serve(lis)
	if errServer != nil {
		logger.Error("error while serve gRPC Server", errServer)
	}
}
