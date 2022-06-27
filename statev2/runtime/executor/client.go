package executor

import (
	"github.com/hermeznetwork/hermez-core/log"
	"github.com/hermeznetwork/hermez-core/statev2/runtime/executor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewExecutorClient(c Config) (pb.ExecutorServiceClient, *grpc.ClientConn) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	executorConn, err := grpc.Dial(c.URI, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	executorClient := pb.NewExecutorServiceClient(executorConn)
	return executorClient, executorConn
}