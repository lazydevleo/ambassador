// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v3/fault.proto

package envoy_extensions_filters_http_fault_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v32 "github.com/datawire/ambassador/pkg/api/envoy/config/route/v3"
	v31 "github.com/datawire/ambassador/pkg/api/envoy/extensions/filters/common/fault/v3"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *v3.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{0}
}
func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return m.Size()
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

// [#next-free-field: 14]
type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v31.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_extensions.filters.http.fault.v3.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*v32.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//   Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//   limit. It's possible for the number of active faults to rise slightly above the configured
	//   amount due to the implementation details.
	MaxActiveFaults *types.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//  This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//  will each get an independent limit.
	ResponseRateLimit *v31.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_delay_percent
	DelayPercentRuntime string `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.abort_percent
	AbortPercentRuntime string `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_duration_ms
	DelayDurationRuntime string `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.http_status
	AbortHttpStatusRuntime string `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.max_active_faults
	MaxActiveFaultsRuntime string `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.rate_limit.response_percent
	ResponseRateLimitPercentRuntime string   `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{1}
}
func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return m.Size()
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v31.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*v32.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *types.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v31.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.extensions.filters.http.fault.v3.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.extensions.filters.http.fault.v3.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/fault/v3/fault.proto", fileDescriptor_1c8de848a52219f2)
}

var fileDescriptor_1c8de848a52219f2 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xaf, 0x9b, 0x7e, 0x4e, 0x6e, 0xd5, 0xd6, 0xbd, 0xb7, 0x98, 0x82, 0x42, 0x28, 0x08,
	0x85, 0xaf, 0xb1, 0xe4, 0x64, 0x01, 0x5d, 0x20, 0x1a, 0xaa, 0x2a, 0x54, 0x05, 0x45, 0xa6, 0xb0,
	0xb5, 0xa6, 0xc9, 0x24, 0xb1, 0x64, 0xcf, 0x58, 0x33, 0xe3, 0x34, 0x59, 0xb2, 0xe3, 0x19, 0x78,
	0x1f, 0xa4, 0x2e, 0x58, 0xb0, 0x64, 0x89, 0xfa, 0x14, 0x88, 0x15, 0x9a, 0x33, 0xb6, 0x9b, 0x36,
	0xad, 0x28, 0x3b, 0x67, 0xce, 0xf9, 0x9d, 0xaf, 0xff, 0x39, 0x41, 0x1e, 0x65, 0x43, 0x3e, 0x76,
	0xe9, 0x48, 0x51, 0x26, 0x43, 0xce, 0xa4, 0xdb, 0x0b, 0x23, 0x45, 0x85, 0x74, 0x07, 0x4a, 0x25,
	0x6e, 0x8f, 0xa4, 0x91, 0x72, 0x87, 0x75, 0xf3, 0x81, 0x13, 0xc1, 0x15, 0xb7, 0x1f, 0x00, 0x83,
	0xcf, 0x18, 0x9c, 0x31, 0x58, 0x33, 0xd8, 0xb8, 0x0e, 0xeb, 0x9b, 0x4f, 0x4c, 0xec, 0x0e, 0x67,
	0xbd, 0xb0, 0xef, 0x0a, 0x9e, 0x2a, 0xaa, 0x43, 0xc1, 0x47, 0xd0, 0xe1, 0x71, 0xc2, 0x19, 0x65,
	0x4a, 0x9a, 0xa8, 0x9b, 0x8d, 0x2b, 0x2b, 0xe9, 0xf0, 0x38, 0xe6, 0xec, 0xd2, 0x5a, 0x36, 0x6f,
	0x19, 0x4a, 0x8d, 0x13, 0x88, 0x9d, 0x50, 0xd1, 0xa1, 0x2c, 0x37, 0x56, 0xfa, 0x9c, 0xf7, 0x23,
	0xea, 0xc2, 0xaf, 0xa3, 0xb4, 0xe7, 0x1e, 0x0b, 0x92, 0x24, 0xba, 0x50, 0x63, 0xbf, 0x9b, 0x76,
	0x13, 0xe2, 0x12, 0xc6, 0xb8, 0x22, 0x0a, 0x52, 0x0e, 0xa9, 0xd0, 0xb9, 0x43, 0xd6, 0xcf, 0x5c,
	0x6e, 0x0c, 0x49, 0x14, 0x76, 0x89, 0xae, 0x3c, 0xfb, 0x30, 0x86, 0xad, 0xaf, 0x16, 0x42, 0x7b,
	0xba, 0x90, 0x9d, 0x23, 0x2e, 0x94, 0x8d, 0x51, 0x59, 0x37, 0x1f, 0x48, 0x45, 0x54, 0x2a, 0x9d,
	0x99, 0xaa, 0x55, 0x5b, 0x6e, 0x96, 0x7f, 0x35, 0x17, 0x1f, 0xcd, 0xaf, 0x7e, 0x9f, 0xad, 0x9d,
	0x58, 0xad, 0x7f, 0x7c, 0xa4, 0x3d, 0xde, 0x81, 0x83, 0xfd, 0x12, 0xa1, 0xac, 0x56, 0xd2, 0xa7,
	0x4e, 0xa9, 0x6a, 0xd5, 0xca, 0x5e, 0x15, 0x9b, 0xc1, 0xea, 0x66, 0xf0, 0xb0, 0x8e, 0xf7, 0x04,
	0xe9, 0xe8, 0xb2, 0x48, 0xd4, 0x36, 0xae, 0xfe, 0x04, 0xb3, 0x5d, 0xff, 0xfc, 0xe5, 0x53, 0x05,
	0x23, 0x33, 0x64, 0x6c, 0x86, 0x9c, 0x09, 0x71, 0x4e, 0x07, 0x0f, 0x9f, 0x95, 0xd9, 0x5c, 0x43,
	0x88, 0x0a, 0xc1, 0x45, 0xa0, 0x73, 0xd8, 0xa5, 0x9f, 0x4d, 0x6b, 0x7f, 0x76, 0xd1, 0x5a, 0x9d,
	0xd9, 0xfa, 0xb8, 0x80, 0x96, 0x5a, 0x87, 0x87, 0x6d, 0xf0, 0xb5, 0xf7, 0xd1, 0x5c, 0x97, 0x46,
	0x64, 0xec, 0x58, 0x50, 0x58, 0x03, 0x5f, 0xa9, 0xb8, 0xd1, 0xa6, 0xd0, 0xdc, 0xe4, 0xda, 0xd5,
	0xac, 0x6f, 0x42, 0xd8, 0x2d, 0x34, 0x47, 0x74, 0x6e, 0x98, 0x49, 0xd9, 0xf3, 0xf0, 0xf5, 0xb6,
	0x67, 0xa2, 0x6a, 0xdf, 0x04, 0xb0, 0x1f, 0xa2, 0xd5, 0x34, 0x91, 0x4a, 0x50, 0x12, 0x07, 0x9d,
	0x28, 0x95, 0x8a, 0x0a, 0x98, 0xdc, 0x92, 0xbf, 0x92, 0xbf, 0xbf, 0x32, 0xcf, 0xf6, 0x0b, 0xb4,
	0x30, 0xa0, 0xa4, 0x4b, 0x85, 0x74, 0x66, 0xab, 0xa5, 0x5a, 0xd9, 0xbb, 0x8f, 0xcf, 0xcd, 0x09,
	0x76, 0x50, 0x67, 0x69, 0x81, 0xd7, 0x1b, 0xa2, 0x3a, 0x03, 0x2a, 0xfc, 0x1c, 0xd2, 0xa9, 0xba,
	0xfc, 0x98, 0x65, 0xc9, 0x18, 0xef, 0x52, 0xe9, 0xcc, 0x55, 0x4b, 0x3a, 0xd5, 0xd9, 0xfb, 0x5b,
	0xfd, 0x6c, 0xb7, 0xd0, 0x5a, 0x4c, 0x46, 0x81, 0x96, 0x6a, 0x48, 0x03, 0x28, 0x5f, 0x3a, 0xf3,
	0xd0, 0xeb, 0x6d, 0x6c, 0x16, 0x10, 0xe7, 0x0b, 0x88, 0xdf, 0xbf, 0x66, 0xaa, 0xee, 0x7d, 0x20,
	0x51, 0x4a, 0xfd, 0x95, 0x98, 0x8c, 0x76, 0x80, 0x82, 0x56, 0xa5, 0x3d, 0x40, 0xeb, 0x82, 0xca,
	0x84, 0x33, 0x49, 0x03, 0x41, 0x14, 0x0d, 0xa2, 0x30, 0x0e, 0x95, 0xb3, 0x00, 0xb1, 0x9e, 0xfd,
	0xa5, 0x06, 0x3e, 0x51, 0xf4, 0x40, 0xf3, 0xfe, 0x5a, 0x1e, 0xb4, 0x78, 0xb2, 0x3d, 0xf4, 0x3f,
	0x88, 0x13, 0x64, 0xfb, 0x14, 0x88, 0x94, 0xa9, 0x30, 0xa6, 0xce, 0x22, 0x8c, 0x73, 0x1d, 0x8c,
	0xf9, 0xd2, 0x19, 0x93, 0x66, 0x40, 0x86, 0x29, 0x66, 0xc9, 0x30, 0x60, 0xbc, 0xc0, 0x34, 0xd0,
	0x86, 0xc9, 0xd3, 0x4d, 0x05, 0x5c, 0x58, 0x01, 0x21, 0x80, 0xfe, 0x03, 0xeb, 0x6e, 0x66, 0xcc,
	0xa9, 0xe7, 0xe8, 0xa6, 0xc9, 0x34, 0x71, 0x51, 0x05, 0x58, 0x06, 0x70, 0x03, 0x1c, 0x5a, 0xc5,
	0x3d, 0x4d, 0xa0, 0x53, 0x62, 0x14, 0xe8, 0xbf, 0x06, 0xbd, 0x30, 0xf6, 0x1c, 0x3d, 0x40, 0xf7,
	0x2e, 0x99, 0xfe, 0x54, 0xb7, 0xcb, 0x10, 0xe4, 0xce, 0xd4, 0x4c, 0xcf, 0x77, 0xbe, 0xed, 0xe9,
	0xeb, 0x7c, 0x8a, 0x1e, 0xff, 0xf9, 0x3a, 0x8b, 0xab, 0x6b, 0xb6, 0x4f, 0x4e, 0x2b, 0xd6, 0xb7,
	0xd3, 0x8a, 0xf5, 0xe3, 0xb4, 0x62, 0xa1, 0x46, 0xc8, 0x8d, 0xe4, 0x89, 0xe0, 0xa3, 0xf1, 0x35,
	0xaf, 0xa6, 0x69, 0xfe, 0x93, 0xda, 0x7a, 0xdf, 0xda, 0xd6, 0xd1, 0x3c, 0x2c, 0x5e, 0xfd, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x60, 0xf1, 0x24, 0xe6, 0x05, 0x00, 0x00,
}

func (m *FaultAbort) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultAbort) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaultAbort) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Percentage != nil {
		{
			size, err := m.Percentage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ErrorType != nil {
		{
			size := m.ErrorType.Size()
			i -= size
			if _, err := m.ErrorType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *FaultAbort_HttpStatus) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *FaultAbort_HttpStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintFault(dAtA, i, uint64(m.HttpStatus))
	i--
	dAtA[i] = 0x10
	return len(dAtA) - i, nil
}
func (m *HTTPFault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPFault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HTTPFault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResponseRateLimitPercentRuntime) > 0 {
		i -= len(m.ResponseRateLimitPercentRuntime)
		copy(dAtA[i:], m.ResponseRateLimitPercentRuntime)
		i = encodeVarintFault(dAtA, i, uint64(len(m.ResponseRateLimitPercentRuntime)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.MaxActiveFaultsRuntime) > 0 {
		i -= len(m.MaxActiveFaultsRuntime)
		copy(dAtA[i:], m.MaxActiveFaultsRuntime)
		i = encodeVarintFault(dAtA, i, uint64(len(m.MaxActiveFaultsRuntime)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.AbortHttpStatusRuntime) > 0 {
		i -= len(m.AbortHttpStatusRuntime)
		copy(dAtA[i:], m.AbortHttpStatusRuntime)
		i = encodeVarintFault(dAtA, i, uint64(len(m.AbortHttpStatusRuntime)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.DelayDurationRuntime) > 0 {
		i -= len(m.DelayDurationRuntime)
		copy(dAtA[i:], m.DelayDurationRuntime)
		i = encodeVarintFault(dAtA, i, uint64(len(m.DelayDurationRuntime)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.AbortPercentRuntime) > 0 {
		i -= len(m.AbortPercentRuntime)
		copy(dAtA[i:], m.AbortPercentRuntime)
		i = encodeVarintFault(dAtA, i, uint64(len(m.AbortPercentRuntime)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DelayPercentRuntime) > 0 {
		i -= len(m.DelayPercentRuntime)
		copy(dAtA[i:], m.DelayPercentRuntime)
		i = encodeVarintFault(dAtA, i, uint64(len(m.DelayPercentRuntime)))
		i--
		dAtA[i] = 0x42
	}
	if m.ResponseRateLimit != nil {
		{
			size, err := m.ResponseRateLimit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.MaxActiveFaults != nil {
		{
			size, err := m.MaxActiveFaults.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.DownstreamNodes) > 0 {
		for iNdEx := len(m.DownstreamNodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DownstreamNodes[iNdEx])
			copy(dAtA[i:], m.DownstreamNodes[iNdEx])
			i = encodeVarintFault(dAtA, i, uint64(len(m.DownstreamNodes[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFault(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.UpstreamCluster) > 0 {
		i -= len(m.UpstreamCluster)
		copy(dAtA[i:], m.UpstreamCluster)
		i = encodeVarintFault(dAtA, i, uint64(len(m.UpstreamCluster)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Abort != nil {
		{
			size, err := m.Abort.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Delay != nil {
		{
			size, err := m.Delay.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	offset -= sovFault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FaultAbort) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrorType != nil {
		n += m.ErrorType.Size()
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultAbort_HttpStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFault(uint64(m.HttpStatus))
	return n
}
func (m *HTTPFault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Delay != nil {
		l = m.Delay.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.Abort != nil {
		l = m.Abort.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.UpstreamCluster)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovFault(uint64(l))
		}
	}
	if len(m.DownstreamNodes) > 0 {
		for _, s := range m.DownstreamNodes {
			l = len(s)
			n += 1 + l + sovFault(uint64(l))
		}
	}
	if m.MaxActiveFaults != nil {
		l = m.MaxActiveFaults.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.ResponseRateLimit != nil {
		l = m.ResponseRateLimit.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.DelayPercentRuntime)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.AbortPercentRuntime)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.DelayDurationRuntime)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.AbortHttpStatusRuntime)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.MaxActiveFaultsRuntime)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.ResponseRateLimitPercentRuntime)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultAbort) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
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
			return fmt.Errorf("proto: FaultAbort: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultAbort: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStatus", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ErrorType = &FaultAbort_HttpStatus{v}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Percentage == nil {
				m.Percentage = &v3.FractionalPercent{}
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
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
func (m *HTTPFault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
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
			return fmt.Errorf("proto: HTTPFault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPFault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delay == nil {
				m.Delay = &v31.FaultDelay{}
			}
			if err := m.Delay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Abort == nil {
				m.Abort = &FaultAbort{}
			}
			if err := m.Abort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpstreamCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpstreamCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &v32.HeaderMatcher{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownstreamNodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownstreamNodes = append(m.DownstreamNodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxActiveFaults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxActiveFaults == nil {
				m.MaxActiveFaults = &types.UInt32Value{}
			}
			if err := m.MaxActiveFaults.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRateLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseRateLimit == nil {
				m.ResponseRateLimit = &v31.FaultRateLimit{}
			}
			if err := m.ResponseRateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayPercentRuntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelayPercentRuntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbortPercentRuntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbortPercentRuntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayDurationRuntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelayDurationRuntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbortHttpStatusRuntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbortHttpStatusRuntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxActiveFaultsRuntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxActiveFaultsRuntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRateLimitPercentRuntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseRateLimitPercentRuntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
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
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
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
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
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
				return 0, ErrInvalidLengthFault
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFault(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFault
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)