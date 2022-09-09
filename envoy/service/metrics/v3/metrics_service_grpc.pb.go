// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: envoy/service/metrics/v3/metrics_service.proto

package metricsv3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) StreamMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsService_StreamMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[0], "/envoy.service.metrics.v3.MetricsService/StreamMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceStreamMetricsClient{stream}
	return x, nil
}

type MetricsService_StreamMetricsClient interface {
	Send(*StreamMetricsMessage) error
	CloseAndRecv() (*StreamMetricsResponse, error)
	grpc.ClientStream
}

type metricsServiceStreamMetricsClient struct {
	grpc.ClientStream
}

func (x *metricsServiceStreamMetricsClient) Send(m *StreamMetricsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsClient) CloseAndRecv() (*StreamMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	// Envoy will connect and send StreamMetricsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure.
	StreamMetrics(MetricsService_StreamMetricsServer) error
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) StreamMetrics(MetricsService_StreamMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_StreamMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).StreamMetrics(&metricsServiceStreamMetricsServer{stream})
}

type MetricsService_StreamMetricsServer interface {
	SendAndClose(*StreamMetricsResponse) error
	Recv() (*StreamMetricsMessage, error)
	grpc.ServerStream
}

type metricsServiceStreamMetricsServer struct {
	grpc.ServerStream
}

func (x *metricsServiceStreamMetricsServer) SendAndClose(m *StreamMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceStreamMetricsServer) Recv() (*StreamMetricsMessage, error) {
	m := new(StreamMetricsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.metrics.v3.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetrics",
			Handler:       _MetricsService_StreamMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/metrics/v3/metrics_service.proto",
}
