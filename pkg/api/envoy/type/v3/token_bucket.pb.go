// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/v3/token_bucket.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Configures a token bucket, typically used for rate limiting.
type TokenBucket struct {
	// The maximum tokens that the bucket can hold. This is also the number of tokens that the bucket
	// initially contains.
	MaxTokens uint32 `protobuf:"varint,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// The number of tokens added to the bucket during each fill interval. If not specified, defaults
	// to a single token.
	TokensPerFill *types.UInt32Value `protobuf:"bytes,2,opt,name=tokens_per_fill,json=tokensPerFill,proto3" json:"tokens_per_fill,omitempty"`
	// The fill interval that tokens are added to the bucket. During each fill interval
	// `tokens_per_fill` are added to the bucket. The bucket will never contain more than
	// `max_tokens` tokens.
	FillInterval         *types.Duration `protobuf:"bytes,3,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TokenBucket) Reset()         { *m = TokenBucket{} }
func (m *TokenBucket) String() string { return proto.CompactTextString(m) }
func (*TokenBucket) ProtoMessage()    {}
func (*TokenBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b88f84e3f7d4627, []int{0}
}
func (m *TokenBucket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenBucket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenBucket.Merge(m, src)
}
func (m *TokenBucket) XXX_Size() int {
	return m.Size()
}
func (m *TokenBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenBucket.DiscardUnknown(m)
}

var xxx_messageInfo_TokenBucket proto.InternalMessageInfo

func (m *TokenBucket) GetMaxTokens() uint32 {
	if m != nil {
		return m.MaxTokens
	}
	return 0
}

func (m *TokenBucket) GetTokensPerFill() *types.UInt32Value {
	if m != nil {
		return m.TokensPerFill
	}
	return nil
}

func (m *TokenBucket) GetFillInterval() *types.Duration {
	if m != nil {
		return m.FillInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenBucket)(nil), "envoy.type.v3.TokenBucket")
}

func init() { proto.RegisterFile("envoy/type/v3/token_bucket.proto", fileDescriptor_5b88f84e3f7d4627) }

var fileDescriptor_5b88f84e3f7d4627 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0x2b, 0x31,
	0x1c, 0x85, 0x9b, 0xde, 0x72, 0x7b, 0x6f, 0x7a, 0x87, 0x2b, 0xb3, 0xd0, 0xf1, 0xdf, 0x30, 0xba,
	0x90, 0xd2, 0x45, 0x02, 0x9d, 0x9d, 0xe0, 0x26, 0x88, 0x50, 0x41, 0x28, 0x45, 0xdd, 0x0e, 0xa9,
	0x4d, 0x4b, 0x68, 0x9a, 0x84, 0x4c, 0x66, 0x6c, 0xdf, 0xc0, 0x67, 0xf0, 0x11, 0x7c, 0x0f, 0xc1,
	0xa5, 0x8f, 0x20, 0x7d, 0x01, 0xf7, 0x5d, 0xc9, 0x64, 0xa6, 0x58, 0xe9, 0x6e, 0xe0, 0x7c, 0xe7,
	0xf0, 0xcd, 0x2f, 0x30, 0x62, 0x32, 0x57, 0x0b, 0x6c, 0x17, 0x9a, 0xe1, 0x3c, 0xc6, 0x56, 0x4d,
	0x99, 0x4c, 0x86, 0xd9, 0xc3, 0x94, 0x59, 0xa4, 0x8d, 0xb2, 0xca, 0xf7, 0x1c, 0x81, 0x0a, 0x02,
	0xe5, 0xf1, 0x41, 0x38, 0x51, 0x6a, 0x22, 0x18, 0x76, 0xe1, 0x30, 0x1b, 0xe3, 0x51, 0x66, 0xa8,
	0xe5, 0x4a, 0x96, 0xf8, 0x76, 0xfe, 0x68, 0xa8, 0xd6, 0xcc, 0xa4, 0x55, 0x7e, 0x92, 0x8d, 0x34,
	0xc5, 0x54, 0x4a, 0x65, 0x5d, 0x2d, 0xc5, 0x39, 0x33, 0x29, 0x57, 0x92, 0xcb, 0x49, 0x85, 0xec,
	0xe5, 0x54, 0xf0, 0x11, 0xb5, 0x0c, 0xaf, 0x3f, 0xca, 0xe0, 0xf4, 0x13, 0xc0, 0xd6, 0x6d, 0x61,
	0x48, 0x9c, 0xa0, 0x7f, 0x06, 0xe1, 0x8c, 0xce, 0x13, 0x27, 0x9d, 0x06, 0x20, 0x02, 0x6d, 0x8f,
	0x34, 0x57, 0xa4, 0xd1, 0xa9, 0x47, 0xb5, 0xc1, 0xdf, 0x19, 0x9d, 0x3b, 0x38, 0xf5, 0x6f, 0xe0,
	0xff, 0x92, 0x49, 0x34, 0x33, 0xc9, 0x98, 0x0b, 0x11, 0xd4, 0x23, 0xd0, 0x6e, 0x75, 0x8f, 0x50,
	0x69, 0x8b, 0xd6, 0xb6, 0xe8, 0xae, 0x27, 0x6d, 0xdc, 0xbd, 0xa7, 0x22, 0x63, 0xdf, 0x53, 0x5e,
	0xd9, 0xee, 0x33, 0x73, 0xc5, 0x85, 0xf0, 0xaf, 0xa1, 0x57, 0x6c, 0x24, 0x5c, 0x5a, 0x66, 0x72,
	0x2a, 0x82, 0x5f, 0x6e, 0x6c, 0x7f, 0x6b, 0xec, 0xb2, 0x3a, 0x0d, 0x81, 0x2b, 0xd2, 0x7c, 0x01,
	0x8d, 0x3f, 0xa0, 0x53, 0x1b, 0xfc, 0x2b, 0xba, 0xbd, 0xaa, 0x7a, 0x7e, 0xfc, 0xfc, 0xfa, 0x14,
	0x06, 0x70, 0x77, 0xe3, 0xc8, 0x1b, 0x7f, 0x48, 0x2e, 0xde, 0x96, 0x21, 0x78, 0x5f, 0x86, 0xe0,
	0x63, 0x19, 0x02, 0x78, 0xc8, 0x15, 0x72, 0xa0, 0x36, 0x6a, 0xbe, 0x40, 0x3f, 0x1e, 0x86, 0xec,
	0x6c, 0xf4, 0xfa, 0x85, 0x41, 0x1f, 0x0c, 0x7f, 0x3b, 0x95, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x06, 0xf1, 0x03, 0xd9, 0xe6, 0x01, 0x00, 0x00,
}

func (m *TokenBucket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenBucket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenBucket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.FillInterval != nil {
		{
			size, err := m.FillInterval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTokenBucket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TokensPerFill != nil {
		{
			size, err := m.TokensPerFill.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTokenBucket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxTokens != 0 {
		i = encodeVarintTokenBucket(dAtA, i, uint64(m.MaxTokens))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenBucket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenBucket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenBucket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxTokens != 0 {
		n += 1 + sovTokenBucket(uint64(m.MaxTokens))
	}
	if m.TokensPerFill != nil {
		l = m.TokensPerFill.Size()
		n += 1 + l + sovTokenBucket(uint64(l))
	}
	if m.FillInterval != nil {
		l = m.FillInterval.Size()
		n += 1 + l + sovTokenBucket(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTokenBucket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenBucket(x uint64) (n int) {
	return sovTokenBucket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenBucket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenBucket
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
			return fmt.Errorf("proto: TokenBucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenBucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTokens", wireType)
			}
			m.MaxTokens = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBucket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTokens |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensPerFill", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBucket
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
				return ErrInvalidLengthTokenBucket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TokensPerFill == nil {
				m.TokensPerFill = &types.UInt32Value{}
			}
			if err := m.TokensPerFill.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FillInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenBucket
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
				return ErrInvalidLengthTokenBucket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenBucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FillInterval == nil {
				m.FillInterval = &types.Duration{}
			}
			if err := m.FillInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenBucket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTokenBucket
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTokenBucket
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
func skipTokenBucket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenBucket
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
					return 0, ErrIntOverflowTokenBucket
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
					return 0, ErrIntOverflowTokenBucket
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
				return 0, ErrInvalidLengthTokenBucket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenBucket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenBucket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenBucket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenBucket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenBucket = fmt.Errorf("proto: unexpected end of group")
)
