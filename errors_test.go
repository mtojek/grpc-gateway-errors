package grpc_gateway_errors

import (
	"github.com/mtojek/grpc-gateway-errors/client"
	"github.com/mtojek/grpc-gateway-errors/server"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestCallApi_ReceiveGRPCError(t *testing.T) {
	// given
	go server.StartServerAndGateway()
	time.Sleep(3 * time.Second)

	// when
	rpcError := client.CallServer()

	// then
	validStatus, ok := status.FromError(rpcError)
	assert.True(t, ok, "should be rpc error")
	assert.Equal(t, "This is not implemented yet", validStatus.Message())
}