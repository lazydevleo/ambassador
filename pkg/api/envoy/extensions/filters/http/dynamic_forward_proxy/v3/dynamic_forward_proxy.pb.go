// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/http/dynamic_forward_proxy/v3/dynamic_forward_proxy.proto

package envoy_extensions_filters_http_dynamic_forward_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

// Configuration for the dynamic forward proxy HTTP filter. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.filters.http.dynamic_forward_proxy]
type FilterConfig struct {
	// The DNS cache configuration that the filter will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_api_field_extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig       *v3.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_919f718544b11cb2, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

// Per route Configuration for the dynamic forward proxy HTTP filter.
type PerRouteConfig struct {
	// Types that are valid to be assigned to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewriteLiteral
	//	*PerRouteConfig_HostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_919f718544b11cb2, []int{1}
}
func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return m.Size()
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PerRouteConfig_HostRewriteLiteral struct {
	HostRewriteLiteral string `protobuf:"bytes,1,opt,name=host_rewrite_literal,json=hostRewriteLiteral,proto3,oneof" json:"host_rewrite_literal,omitempty"`
}
type PerRouteConfig_HostRewriteHeader struct {
	HostRewriteHeader string `protobuf:"bytes,2,opt,name=host_rewrite_header,json=hostRewriteHeader,proto3,oneof" json:"host_rewrite_header,omitempty"`
}

func (*PerRouteConfig_HostRewriteLiteral) isPerRouteConfig_HostRewriteSpecifier() {}
func (*PerRouteConfig_HostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier()  {}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (m *PerRouteConfig) GetHostRewriteLiteral() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewriteLiteral); ok {
		return x.HostRewriteLiteral
	}
	return ""
}

func (m *PerRouteConfig) GetHostRewriteHeader() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewriteHeader); ok {
		return x.HostRewriteHeader
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerRouteConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerRouteConfig_HostRewriteLiteral)(nil),
		(*PerRouteConfig_HostRewriteHeader)(nil),
	}
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.extensions.filters.http.dynamic_forward_proxy.v3.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.extensions.filters.http.dynamic_forward_proxy.v3.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/dynamic_forward_proxy/v3/dynamic_forward_proxy.proto", fileDescriptor_919f718544b11cb2)
}

var fileDescriptor_919f718544b11cb2 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4a, 0xfb, 0x40,
	0x14, 0xc7, 0x7f, 0x53, 0xf8, 0x89, 0x8e, 0x52, 0x6a, 0x14, 0x0d, 0x5d, 0x04, 0xed, 0xca, 0xd5,
	0x44, 0x5a, 0x70, 0xd1, 0x55, 0x49, 0xff, 0x58, 0xa4, 0x8b, 0x90, 0x0b, 0x84, 0x31, 0x99, 0x34,
	0x03, 0xe9, 0x4c, 0x98, 0x4c, 0xd3, 0xf6, 0x06, 0xe2, 0x11, 0xbc, 0x8f, 0xe0, 0xb2, 0x27, 0x10,
	0xe9, 0x31, 0x5c, 0x49, 0x32, 0x51, 0x1b, 0x5a, 0xbb, 0xe8, 0x2e, 0xcc, 0x27, 0xef, 0xf3, 0xde,
	0x77, 0xe6, 0x41, 0x87, 0xb0, 0x94, 0x2f, 0x4c, 0x32, 0x97, 0x84, 0x25, 0x94, 0xb3, 0xc4, 0x0c,
	0x68, 0x24, 0x89, 0x48, 0xcc, 0x50, 0xca, 0xd8, 0xf4, 0x17, 0x0c, 0x4f, 0xa8, 0xe7, 0x06, 0x5c,
	0xcc, 0xb0, 0xf0, 0xdd, 0x58, 0xf0, 0xf9, 0xc2, 0x4c, 0x5b, 0xdb, 0x01, 0x8a, 0x05, 0x97, 0x5c,
	0xbb, 0xcb, 0x9d, 0xe8, 0xd7, 0x89, 0x0a, 0x27, 0xca, 0x9c, 0x68, 0x7b, 0x69, 0xda, 0xaa, 0x77,
	0x36, 0x66, 0xf1, 0xf8, 0x64, 0xc2, 0xd9, 0x8e, 0x29, 0x58, 0xe2, 0x7a, 0xd8, 0x0b, 0x89, 0xea,
	0x5c, 0xbf, 0x9e, 0xfa, 0x31, 0x36, 0x31, 0x63, 0x5c, 0x62, 0x99, 0x1b, 0x52, 0x22, 0x32, 0x15,
	0x65, 0xe3, 0xe2, 0x97, 0xcb, 0x14, 0x47, 0xd4, 0xc7, 0x92, 0x98, 0xdf, 0x1f, 0x0a, 0x34, 0x96,
	0x00, 0x9e, 0x0c, 0xf2, 0x39, 0xbb, 0x9c, 0x05, 0x74, 0xac, 0x49, 0x58, 0xfb, 0xf1, 0xbb, 0x5e,
	0x7e, 0xa6, 0x83, 0x2b, 0x70, 0x73, 0xdc, 0xec, 0xa0, 0x8d, 0x84, 0x6a, 0xd2, 0x3f, 0xb3, 0xa1,
	0x1e, 0x4b, 0xba, 0x99, 0x48, 0xb9, 0xad, 0xc3, 0x4f, 0xeb, 0xff, 0x33, 0xa8, 0xd4, 0x80, 0x53,
	0xf5, 0x4b, 0xa4, 0xfd, 0xf0, 0xf2, 0xfa, 0x64, 0xf4, 0x61, 0x57, 0x75, 0x50, 0x6d, 0x8b, 0xfb,
	0xdb, 0x79, 0x7d, 0x4d, 0x1c, 0xc5, 0x21, 0x46, 0xeb, 0x09, 0x1a, 0xef, 0x00, 0x56, 0x6d, 0x22,
	0x1c, 0x3e, 0x95, 0x85, 0x5e, 0x6b, 0xc2, 0xf3, 0x90, 0x27, 0xd2, 0x15, 0x64, 0x26, 0xa8, 0x24,
	0x6e, 0x44, 0x25, 0x11, 0x38, 0xca, 0x83, 0x1d, 0x0d, 0xff, 0x39, 0x5a, 0x46, 0x1d, 0x05, 0x47,
	0x8a, 0x69, 0xb7, 0xf0, 0xac, 0x54, 0x13, 0x12, 0xec, 0x13, 0xa1, 0x57, 0x8a, 0x92, 0xd3, 0xb5,
	0x92, 0x61, 0x8e, 0xda, 0xa3, 0x2c, 0xc4, 0x3d, 0xec, 0xef, 0x19, 0xa2, 0x3c, 0xb3, 0xa5, 0xc3,
	0x8b, 0x52, 0xff, 0x24, 0x26, 0x1e, 0x0d, 0x28, 0x11, 0x16, 0x7b, 0x5b, 0x19, 0x60, 0xb9, 0x32,
	0xc0, 0xc7, 0xca, 0x00, 0xb0, 0x47, 0xb9, 0x7a, 0x18, 0x65, 0xdb, 0x6f, 0x0b, 0x2d, 0xbd, 0xa7,
	0xc8, 0x40, 0x01, 0x3b, 0x3b, 0xb7, 0xb3, 0x0d, 0xb1, 0xc1, 0xe3, 0x41, 0xbe, 0x2a, 0xad, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xdf, 0x27, 0x72, 0x36, 0x03, 0x00, 0x00,
}

func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DnsCacheConfig != nil {
		{
			size, err := m.DnsCacheConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PerRouteConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerRouteConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerRouteConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HostRewriteSpecifier != nil {
		{
			size := m.HostRewriteSpecifier.Size()
			i -= size
			if _, err := m.HostRewriteSpecifier.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PerRouteConfig_HostRewriteLiteral) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerRouteConfig_HostRewriteLiteral) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.HostRewriteLiteral)
	copy(dAtA[i:], m.HostRewriteLiteral)
	i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(len(m.HostRewriteLiteral)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *PerRouteConfig_HostRewriteHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerRouteConfig_HostRewriteHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.HostRewriteHeader)
	copy(dAtA[i:], m.HostRewriteHeader)
	i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(len(m.HostRewriteHeader)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintDynamicForwardProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicForwardProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DnsCacheConfig != nil {
		l = m.DnsCacheConfig.Size()
		n += 1 + l + sovDynamicForwardProxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PerRouteConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HostRewriteSpecifier != nil {
		n += m.HostRewriteSpecifier.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PerRouteConfig_HostRewriteLiteral) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostRewriteLiteral)
	n += 1 + l + sovDynamicForwardProxy(uint64(l))
	return n
}
func (m *PerRouteConfig_HostRewriteHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostRewriteHeader)
	n += 1 + l + sovDynamicForwardProxy(uint64(l))
	return n
}

func sovDynamicForwardProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicForwardProxy(x uint64) (n int) {
	return sovDynamicForwardProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicForwardProxy
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
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsCacheConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
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
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsCacheConfig == nil {
				m.DnsCacheConfig = &v3.DnsCacheConfig{}
			}
			if err := m.DnsCacheConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicForwardProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicForwardProxy
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
func (m *PerRouteConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicForwardProxy
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
			return fmt.Errorf("proto: PerRouteConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerRouteConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostRewriteLiteral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
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
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostRewriteSpecifier = &PerRouteConfig_HostRewriteLiteral{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostRewriteHeader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
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
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostRewriteSpecifier = &PerRouteConfig_HostRewriteHeader{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicForwardProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicForwardProxy
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
func skipDynamicForwardProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicForwardProxy
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
					return 0, ErrIntOverflowDynamicForwardProxy
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
					return 0, ErrIntOverflowDynamicForwardProxy
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
				return 0, ErrInvalidLengthDynamicForwardProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicForwardProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicForwardProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicForwardProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicForwardProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicForwardProxy = fmt.Errorf("proto: unexpected end of group")
)
