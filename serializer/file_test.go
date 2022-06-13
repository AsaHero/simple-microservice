package serializer_test

import (
	"testing"

	"github.com/AsaHero/simple-microservice/pb"
	"github.com/AsaHero/simple-microservice/sample"
	"github.com/AsaHero/simple-microservice/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_WriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin" 
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}