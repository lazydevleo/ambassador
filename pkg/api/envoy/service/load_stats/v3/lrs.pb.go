// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/load_stats/v3/lrs.proto

package envoy_service_load_stats_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	v31 "github.com/datawire/ambassador/pkg/api/envoy/config/endpoint/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// A load report Envoy sends to the management server.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsRequest struct {
	// Node identifier for Envoy instance.
	Node *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of load stats to report.
	ClusterStats         []*v31.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a87f04ee5f8fca3, []int{0}
}
func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(m, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return m.Size()
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*v31.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

// The management server sends envoy a LoadStatsResponse with all clusters it
// is interested in learning load stats about.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsResponse struct {
	// Clusters to report stats for.
	Clusters []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The minimum interval of time to collect stats over. This is only a minimum for two reasons:
	// 1. There may be some delay from when the timer fires until stats sampling occurs.
	// 2. For clusters that were already feature in the previous *LoadStatsResponse*, any traffic
	//    that is observed in between the corresponding previous *LoadStatsRequest* and this
	//    *LoadStatsResponse* will also be accumulated and billed to the cluster. This avoids a period
	//    of inobservability that might otherwise exists between the messages. New clusters are not
	//    subject to this consideration.
	LoadReportingInterval *types.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	// Set to *true* if the management server supports endpoint granularity
	// report.
	ReportEndpointGranularity bool     `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a87f04ee5f8fca3, []int{1}
}
func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(m, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return m.Size()
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *types.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v3.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v3.LoadStatsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v3/lrs.proto", fileDescriptor_4a87f04ee5f8fca3)
}

var fileDescriptor_4a87f04ee5f8fca3 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x71, 0x0b, 0x28, 0x78, 0x20, 0x20, 0x02, 0x2d, 0xeb, 0xa4, 0x50, 0x2a, 0x40, 0x01,
	0x31, 0x1b, 0x35, 0x9c, 0x76, 0xe0, 0x50, 0xfe, 0x09, 0x51, 0xa1, 0x91, 0x7e, 0x80, 0xca, 0x4d,
	0xbc, 0xc8, 0x52, 0xf0, 0x1b, 0x6c, 0xc7, 0xa2, 0x57, 0x4e, 0x3b, 0x71, 0xe0, 0xc8, 0xe7, 0x01,
	0x89, 0x23, 0x1f, 0x01, 0xf5, 0x63, 0x70, 0x42, 0x89, 0x93, 0xad, 0x03, 0x6d, 0xea, 0x2d, 0xf6,
	0xfb, 0x7b, 0x1e, 0xbf, 0xef, 0xf3, 0x06, 0xdf, 0xe7, 0xd2, 0xc2, 0x92, 0x6a, 0xae, 0xac, 0x48,
	0x39, 0x2d, 0x80, 0x65, 0x73, 0x6d, 0x98, 0xd1, 0xd4, 0xc6, 0xb4, 0x50, 0x9a, 0x94, 0x0a, 0x0c,
	0xf8, 0xbb, 0x0d, 0x46, 0x5a, 0x8c, 0x9c, 0x60, 0xc4, 0xc6, 0x83, 0x3b, 0xce, 0x23, 0x05, 0x79,
	0x28, 0x72, 0x9a, 0x82, 0xe2, 0xb5, 0x78, 0xc1, 0x34, 0x77, 0xea, 0xc1, 0xa3, 0x53, 0x00, 0x97,
	0x59, 0x09, 0x42, 0x9a, 0xe6, 0x85, 0xda, 0x48, 0xf1, 0x12, 0x94, 0x69, 0xd9, 0x30, 0x07, 0xc8,
	0x0b, 0x4e, 0x9b, 0xd3, 0xa2, 0x3a, 0xa4, 0x59, 0xa5, 0x98, 0x11, 0x20, 0xdb, 0xfa, 0xdd, 0x2a,
	0x2b, 0x19, 0x65, 0x52, 0x82, 0x69, 0xae, 0x35, 0xb5, 0x5c, 0x69, 0x01, 0x52, 0xc8, 0xbc, 0x45,
	0xb6, 0x2d, 0x2b, 0x44, 0xc6, 0x0c, 0xa7, 0xdd, 0x87, 0x2b, 0x8c, 0xbe, 0x23, 0x7c, 0x63, 0x0a,
	0x2c, 0x9b, 0xd5, 0x9d, 0x27, 0xfc, 0x63, 0xc5, 0xb5, 0xf1, 0x09, 0xbe, 0x28, 0x21, 0xe3, 0x01,
	0x1a, 0xa2, 0x68, 0x6b, 0x3c, 0x20, 0x6e, 0x52, 0xd7, 0x2b, 0xa9, 0x87, 0x21, 0x36, 0x26, 0xef,
	0x20, 0xe3, 0x49, 0xc3, 0xf9, 0x6f, 0xf1, 0xb5, 0xb4, 0xa8, 0xb4, 0xe1, 0xca, 0x25, 0x10, 0xf4,
	0x86, 0xfd, 0x68, 0x6b, 0xfc, 0xe0, 0xb4, 0xb0, 0x1b, 0xb2, 0x16, 0x3f, 0x77, 0xb8, 0x7b, 0xf5,
	0x6a, 0xba, 0x76, 0xda, 0x8f, 0xbf, 0xfd, 0x38, 0x0a, 0x09, 0x7e, 0x7c, 0x76, 0xbc, 0x63, 0xf2,
	0x6f, 0xc7, 0xa3, 0xcf, 0x3d, 0x7c, 0x73, 0xed, 0x52, 0x97, 0x20, 0x35, 0xf7, 0xef, 0x61, 0xaf,
	0xb5, 0xd6, 0x01, 0x1a, 0xf6, 0xa3, 0x2b, 0x13, 0xef, 0xcf, 0xe4, 0xd2, 0x57, 0xd4, 0xf3, 0x50,
	0x72, 0x5c, 0xf1, 0xdf, 0xe3, 0xed, 0xb5, 0xcc, 0x85, 0xcc, 0xe7, 0x42, 0x1a, 0xae, 0x2c, 0x2b,
	0x82, 0x5e, 0x13, 0xc0, 0x0e, 0x71, 0x0b, 0x20, 0xdd, 0x02, 0xc8, 0x8b, 0x76, 0x01, 0xc9, 0xed,
	0x5a, 0x99, 0x74, 0xc2, 0x37, 0xad, 0xce, 0x7f, 0x86, 0x77, 0x9d, 0xdb, 0xbc, 0x1b, 0x7a, 0x9e,
	0x2b, 0x26, 0xab, 0x82, 0x29, 0x61, 0x96, 0x41, 0x7f, 0x88, 0x22, 0x2f, 0xd9, 0x71, 0xc8, 0xcb,
	0x96, 0x78, 0x7d, 0x02, 0xec, 0x3f, 0xad, 0x33, 0xa0, 0x78, 0x6f, 0xc3, 0x0c, 0xdc, 0xb8, 0xe3,
	0x2f, 0x08, 0xdf, 0x9a, 0xae, 0xf7, 0x33, 0x73, 0x42, 0xdf, 0xe2, 0xeb, 0x33, 0xa3, 0x38, 0xfb,
	0x70, 0xac, 0xf1, 0xf7, 0xc8, 0x39, 0xbf, 0xef, 0x7f, 0xf9, 0x0e, 0xc8, 0xa6, 0xb8, 0x6b, 0x65,
	0x74, 0x21, 0x42, 0x4f, 0xd0, 0xe4, 0xd5, 0xcf, 0x55, 0x88, 0x7e, 0xad, 0x42, 0xf4, 0x7b, 0x15,
	0x22, 0xfc, 0x50, 0x80, 0x73, 0x29, 0x15, 0x7c, 0x5a, 0x9e, 0x67, 0x38, 0xf1, 0xa6, 0x4a, 0x1f,
	0xd4, 0x61, 0x1f, 0xa0, 0x23, 0x84, 0x16, 0x97, 0x9b, 0xe0, 0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x82, 0xfb, 0x91, 0x7b, 0x9a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v3.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (*UnimplementedLoadReportingServiceServer) StreamLoadStats(srv LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v3.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v3/lrs.proto",
}

func (m *LoadStatsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadStatsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoadStatsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClusterStats) > 0 {
		for iNdEx := len(m.ClusterStats) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClusterStats[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLrs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Node != nil {
		{
			size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLrs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoadStatsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadStatsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoadStatsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ReportEndpointGranularity {
		i--
		if m.ReportEndpointGranularity {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.LoadReportingInterval != nil {
		{
			size, err := m.LoadReportingInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLrs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Clusters) > 0 {
		for iNdEx := len(m.Clusters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Clusters[iNdEx])
			copy(dAtA[i:], m.Clusters[iNdEx])
			i = encodeVarintLrs(dAtA, i, uint64(len(m.Clusters[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintLrs(dAtA []byte, offset int, v uint64) int {
	offset -= sovLrs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoadStatsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovLrs(uint64(l))
	}
	if len(m.ClusterStats) > 0 {
		for _, e := range m.ClusterStats {
			l = e.Size()
			n += 1 + l + sovLrs(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoadStatsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Clusters) > 0 {
		for _, s := range m.Clusters {
			l = len(s)
			n += 1 + l + sovLrs(uint64(l))
		}
	}
	if m.LoadReportingInterval != nil {
		l = m.LoadReportingInterval.Size()
		n += 1 + l + sovLrs(uint64(l))
	}
	if m.ReportEndpointGranularity {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLrs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLrs(x uint64) (n int) {
	return sovLrs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoadStatsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadStatsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadStatsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLrs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &v3.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLrs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterStats = append(m.ClusterStats, &v31.ClusterStats{})
			if err := m.ClusterStats[len(m.ClusterStats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLrs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLrs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoadStatsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLrs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadStatsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadStatsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLrs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadReportingInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLrs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLrs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LoadReportingInterval == nil {
				m.LoadReportingInterval = &types.Duration{}
			}
			if err := m.LoadReportingInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportEndpointGranularity", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ReportEndpointGranularity = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLrs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLrs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLrs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLrs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLrs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLrs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLrs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLrs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLrs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLrs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLrs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLrs = fmt.Errorf("proto: unexpected end of group")
)
