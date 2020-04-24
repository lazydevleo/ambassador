// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/v3/string.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
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

// Specifies the way to match a string.
// [#next-free-field: 6]
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	MatchPattern         isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{0}
}
func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return m.Size()
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof" json:"exact,omitempty"`
}
type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}
type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof" json:"suffix,omitempty"`
}
type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof" json:"safe_regex,omitempty"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern()     {}
func (*StringMatcher_Prefix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_Suffix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{1}
}
func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return m.Size()
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.v3.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.v3.ListStringMatcher")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3/string.proto", fileDescriptor_e33cffa01bf36e0e)
}

var fileDescriptor_e33cffa01bf36e0e = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x18, 0xc5, 0x9d, 0x68, 0xbc, 0x71, 0x44, 0xf0, 0x86, 0xfb, 0x27, 0xb8, 0x08, 0x31, 0x4a, 0x2b,
	0x14, 0x12, 0xd0, 0x9d, 0xab, 0x32, 0x2b, 0x91, 0xb6, 0x48, 0xfa, 0x00, 0x32, 0xd5, 0xd1, 0x06,
	0xda, 0x4c, 0x98, 0x8c, 0x21, 0xbe, 0x41, 0x5b, 0xba, 0xea, 0xb2, 0xef, 0x53, 0xe8, 0xb2, 0x8f,
	0x50, 0x7c, 0x8a, 0xe2, 0xaa, 0xcc, 0x4c, 0x2a, 0x15, 0xd3, 0xdd, 0x0c, 0xe7, 0x77, 0xce, 0x77,
	0x66, 0x3e, 0xe8, 0x92, 0x28, 0xa5, 0x6b, 0x9f, 0xaf, 0x63, 0xe2, 0xdf, 0x62, 0x3e, 0xbb, 0x26,
	0xcc, 0x4f, 0x07, 0x7e, 0xc2, 0x59, 0x18, 0x2d, 0xbd, 0x98, 0x51, 0x4e, 0xcd, 0xbf, 0x92, 0xf1,
	0x04, 0xe3, 0xe5, 0x8c, 0x97, 0x0e, 0x5a, 0xed, 0x62, 0x2b, 0x23, 0x4b, 0x92, 0x29, 0x67, 0xab,
	0xbd, 0x9a, 0xc7, 0xd8, 0xc7, 0x51, 0x44, 0x39, 0xe6, 0x21, 0x8d, 0x12, 0x3f, 0x25, 0x2c, 0x09,
	0x69, 0xb4, 0x0b, 0x6f, 0x75, 0x54, 0xca, 0x77, 0x66, 0x4e, 0x62, 0x46, 0x66, 0xf2, 0x92, 0x43,
	0xff, 0x53, 0x7c, 0x13, 0xce, 0x31, 0x27, 0xfe, 0xd7, 0x41, 0x09, 0xee, 0xbd, 0x06, 0x1b, 0x97,
	0xb2, 0xeb, 0xb9, 0x6a, 0x60, 0xfe, 0x83, 0x3a, 0xc9, 0xf0, 0x8c, 0x5b, 0xc0, 0x01, 0xbd, 0xda,
	0xa8, 0x14, 0xa8, 0xab, 0xd9, 0x86, 0xd5, 0x98, 0x91, 0x45, 0x98, 0x59, 0x9a, 0x10, 0xd0, 0xaf,
	0x2d, 0xaa, 0x30, 0xcd, 0x01, 0xa3, 0x52, 0x90, 0x0b, 0x02, 0x49, 0x56, 0x0b, 0x81, 0x94, 0x0f,
	0x10, 0x25, 0x98, 0x17, 0x10, 0x26, 0x78, 0x41, 0xa6, 0xf2, 0x91, 0x96, 0xee, 0x80, 0x5e, 0xbd,
	0xdf, 0xf1, 0x0a, 0xff, 0xc7, 0x0b, 0x04, 0x93, 0xd7, 0x42, 0xc6, 0x16, 0xe9, 0x0f, 0x40, 0x6b,
	0x8a, 0xb0, 0x9a, 0x88, 0x90, 0xea, 0xf0, 0xf8, 0xf9, 0xe5, 0xce, 0x76, 0xa1, 0x53, 0x90, 0xb0,
	0xf7, 0x2c, 0xf4, 0x07, 0x36, 0xa4, 0x30, 0x8d, 0x31, 0xe7, 0x84, 0x45, 0x66, 0xf9, 0x03, 0x81,
	0x71, 0xc5, 0xa8, 0x34, 0xf5, 0x40, 0x97, 0x6d, 0xdc, 0x47, 0x00, 0x7f, 0x9f, 0x85, 0x09, 0xdf,
	0xff, 0x8f, 0x31, 0x34, 0x72, 0x4b, 0x62, 0x01, 0xa7, 0xdc, 0xab, 0xf7, 0xbb, 0x3f, 0xf4, 0xdd,
	0x1f, 0x28, 0x0a, 0x3f, 0x01, 0xcd, 0x00, 0xc1, 0xce, 0x3f, 0x3c, 0x11, 0x6d, 0x8f, 0x60, 0xb7,
	0xc0, 0x7f, 0x30, 0x18, 0x9d, 0xbe, 0x6e, 0x6c, 0xf0, 0xb6, 0xb1, 0xc1, 0xfb, 0xc6, 0x06, 0xb0,
	0x13, 0x52, 0x35, 0x36, 0x66, 0x34, 0x5b, 0x17, 0x37, 0x40, 0x75, 0x95, 0x30, 0x11, 0xab, 0x9d,
	0x80, 0xab, 0xaa, 0xdc, 0xf1, 0xe0, 0x33, 0x00, 0x00, 0xff, 0xff, 0x96, 0x1d, 0x64, 0xae, 0xa4,
	0x02, 0x00, 0x00,
}

func (m *StringMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MatchPattern != nil {
		{
			size := m.MatchPattern.Size()
			i -= size
			if _, err := m.MatchPattern.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *StringMatcher_Exact) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_Exact) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Exact)
	copy(dAtA[i:], m.Exact)
	i = encodeVarintString(dAtA, i, uint64(len(m.Exact)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Prefix) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_Prefix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Prefix)
	copy(dAtA[i:], m.Prefix)
	i = encodeVarintString(dAtA, i, uint64(len(m.Prefix)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *StringMatcher_Suffix) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_Suffix) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Suffix)
	copy(dAtA[i:], m.Suffix)
	i = encodeVarintString(dAtA, i, uint64(len(m.Suffix)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *StringMatcher_SafeRegex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StringMatcher_SafeRegex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SafeRegex != nil {
		{
			size, err := m.SafeRegex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintString(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *ListStringMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListStringMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListStringMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Patterns) > 0 {
		for iNdEx := len(m.Patterns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Patterns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintString(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintString(dAtA []byte, offset int, v uint64) int {
	offset -= sovString(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StringMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchPattern != nil {
		n += m.MatchPattern.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StringMatcher_Exact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Exact)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Prefix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Prefix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_Suffix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Suffix)
	n += 1 + l + sovString(uint64(l))
	return n
}
func (m *StringMatcher_SafeRegex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SafeRegex != nil {
		l = m.SafeRegex.Size()
		n += 1 + l + sovString(uint64(l))
	}
	return n
}
func (m *ListStringMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Patterns) > 0 {
		for _, e := range m.Patterns {
			l = e.Size()
			n += 1 + l + sovString(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovString(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozString(x uint64) (n int) {
	return sovString(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StringMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
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
			return fmt.Errorf("proto: StringMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Exact{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Prefix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchPattern = &StringMatcher_Suffix{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafeRegex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RegexMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &StringMatcher_SafeRegex{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthString
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
func (m *ListStringMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowString
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
			return fmt.Errorf("proto: ListStringMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListStringMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patterns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowString
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
				return ErrInvalidLengthString
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthString
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Patterns = append(m.Patterns, &StringMatcher{})
			if err := m.Patterns[len(m.Patterns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipString(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthString
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthString
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
func skipString(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
					return 0, ErrIntOverflowString
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
				return 0, ErrInvalidLengthString
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupString
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthString
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthString        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowString          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupString = fmt.Errorf("proto: unexpected end of group")
)
