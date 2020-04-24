// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/ratelimit/v3/rls.proto

package envoy_config_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
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

// Rate limit :ref:`configuration overview <config_rate_limit_service>`.
type RateLimitServiceConfig struct {
	// Specifies the gRPC service that hosts the rate limit service. The client
	// will connect to this cluster when it needs to make rate limit service
	// requests.
	GrpcService          *v3.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RateLimitServiceConfig) Reset()         { *m = RateLimitServiceConfig{} }
func (m *RateLimitServiceConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitServiceConfig) ProtoMessage()    {}
func (*RateLimitServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca555cbfdc4e6f71, []int{0}
}
func (m *RateLimitServiceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitServiceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitServiceConfig.Merge(m, src)
}
func (m *RateLimitServiceConfig) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitServiceConfig proto.InternalMessageInfo

func (m *RateLimitServiceConfig) GetGrpcService() *v3.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitServiceConfig)(nil), "envoy.config.ratelimit.v3.RateLimitServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/ratelimit/v3/rls.proto", fileDescriptor_ca555cbfdc4e6f71)
}

var fileDescriptor_ca555cbfdc4e6f71 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2c, 0x49, 0xcd, 0xc9, 0xcc,
	0xcd, 0x2c, 0xd1, 0x2f, 0x33, 0xd6, 0x2f, 0xca, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x04, 0x2b, 0xd2, 0x83, 0x28, 0xd2, 0x83, 0x2b, 0xd2, 0x2b, 0x33, 0x96, 0x52, 0x47, 0xd1,
	0x9f, 0x9c, 0x5f, 0x94, 0x0a, 0xd2, 0x9a, 0x5e, 0x54, 0x90, 0x1c, 0x5f, 0x9c, 0x5a, 0x54, 0x96,
	0x99, 0x9c, 0x0a, 0x31, 0x43, 0x4a, 0xb1, 0x34, 0xa5, 0x20, 0x51, 0x3f, 0x31, 0x2f, 0x2f, 0xbf,
	0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xbf, 0x2c, 0xb5, 0xa8, 0x38, 0x33, 0x3f, 0x2f, 0x33,
	0x2f, 0x1d, 0xaa, 0x44, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80,
	0x48, 0x28, 0xad, 0x65, 0xe4, 0x12, 0x0b, 0x4a, 0x2c, 0x49, 0xf5, 0x01, 0xd9, 0x1a, 0x0c, 0x31,
	0xd6, 0x19, 0x6c, 0xa3, 0x90, 0x1f, 0x17, 0x0f, 0xb2, 0x65, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x8a, 0x7a, 0x28, 0x2e, 0x06, 0x39, 0x4b, 0xaf, 0xcc, 0x58, 0xcf, 0xbd, 0xa8, 0x20, 0x19,
	0xaa, 0xdd, 0x89, 0xe3, 0x97, 0x13, 0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0x77, 0x3a, 0x42,
	0xd8, 0xca, 0x7c, 0xd6, 0xd1, 0x0e, 0x39, 0x23, 0x2e, 0x03, 0x5c, 0x3e, 0x36, 0xd2, 0xc3, 0xee,
	0x10, 0x2f, 0x16, 0x0e, 0x46, 0x01, 0x26, 0x2f, 0x16, 0x0e, 0x66, 0x01, 0x16, 0x27, 0xc7, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x91, 0x4b, 0x3d, 0x33, 0x1f,
	0xe2, 0x9c, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x9c, 0x61, 0xe9, 0xc4, 0x11, 0x94, 0x53, 0x1c,
	0x00, 0xf2, 0x70, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0xe7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0x46, 0xa6, 0xb9, 0xa0, 0x01, 0x00, 0x00,
}

func (m *RateLimitServiceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitServiceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitServiceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.GrpcService != nil {
		{
			size, err := m.GrpcService.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRls(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintRls(dAtA []byte, offset int, v uint64) int {
	offset -= sovRls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RateLimitServiceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovRls(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRls(x uint64) (n int) {
	return sovRls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimitServiceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRls
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
			return fmt.Errorf("proto: RateLimitServiceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitServiceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRls
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
				return ErrInvalidLengthRls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &v3.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRls
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
func skipRls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRls
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
					return 0, ErrIntOverflowRls
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
					return 0, ErrIntOverflowRls
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
				return 0, ErrInvalidLengthRls
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRls
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRls
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRls        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRls          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRls = fmt.Errorf("proto: unexpected end of group")
)
