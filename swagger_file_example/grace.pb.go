// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grace.proto

package swaggerfileserver

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// 请求
type SayHelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fa39edfc5fc594, []int{0}
}
func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (*SayHelloRequest) XXX_MessageName() string {
	return "api.swaggerfileserver.SayHelloRequest"
}

type SayHelloReply struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *SayHelloReply) Reset()         { *m = SayHelloReply{} }
func (m *SayHelloReply) String() string { return proto.CompactTextString(m) }
func (*SayHelloReply) ProtoMessage()    {}
func (*SayHelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_13fa39edfc5fc594, []int{1}
}
func (m *SayHelloReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayHelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayHelloReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayHelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloReply.Merge(m, src)
}
func (m *SayHelloReply) XXX_Size() int {
	return m.Size()
}
func (m *SayHelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloReply proto.InternalMessageInfo

func (*SayHelloReply) XXX_MessageName() string {
	return "api.swaggerfileserver.SayHelloReply"
}
func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "api.swaggerfileserver.SayHelloRequest")
	golang_proto.RegisterType((*SayHelloRequest)(nil), "api.swaggerfileserver.SayHelloRequest")
	proto.RegisterType((*SayHelloReply)(nil), "api.swaggerfileserver.SayHelloReply")
	golang_proto.RegisterType((*SayHelloReply)(nil), "api.swaggerfileserver.SayHelloReply")
}

func init() { proto.RegisterFile("grace.proto", fileDescriptor_13fa39edfc5fc594) }
func init() { golang_proto.RegisterFile("grace.proto", fileDescriptor_13fa39edfc5fc594) }

var fileDescriptor_13fa39edfc5fc594 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xb3, 0xe5, 0xdf, 0xbf, 0xba, 0x22, 0xc2, 0x8a, 0x10, 0x8a, 0x2c, 0x12, 0x54, 0x44,
	0x30, 0x01, 0xc5, 0x8b, 0xe2, 0xa5, 0x17, 0x3d, 0xd7, 0x5b, 0x6f, 0xdb, 0x74, 0xbb, 0x0d, 0xa4,
	0x3b, 0x71, 0xb3, 0x51, 0xe2, 0xd1, 0x17, 0x50, 0xf0, 0xe4, 0xdb, 0x78, 0xec, 0xb1, 0xe0, 0xc5,
	0xa3, 0x6d, 0x7c, 0x00, 0x1f, 0x41, 0x3a, 0x6d, 0xa9, 0xa4, 0x05, 0x6f, 0x33, 0xfb, 0xfd, 0x66,
	0xf8, 0xf6, 0x1b, 0xba, 0xae, 0x8c, 0x08, 0xa5, 0x9f, 0x18, 0xb0, 0xc0, 0xb6, 0x45, 0x12, 0xf9,
	0xe9, 0xbd, 0x50, 0x4a, 0x9a, 0x4e, 0x14, 0xcb, 0x54, 0x9a, 0x3b, 0x69, 0x6a, 0xc7, 0x2a, 0xb2,
	0xdd, 0xac, 0xe5, 0x87, 0xd0, 0x0b, 0x14, 0x28, 0x08, 0x90, 0x6e, 0x65, 0x1d, 0xec, 0xb0, 0xc1,
	0x6a, 0xb2, 0xa5, 0xc6, 0x15, 0x80, 0x8a, 0xe5, 0x9c, 0x6a, 0x67, 0x46, 0xd8, 0x08, 0xf4, 0x54,
	0xdf, 0x29, 0xeb, 0xa9, 0x35, 0x59, 0x68, 0x4b, 0xaa, 0x48, 0xa2, 0x40, 0x68, 0x0d, 0x16, 0x47,
	0xd3, 0x89, 0xea, 0xed, 0xd3, 0xcd, 0x1b, 0x91, 0x5f, 0xcb, 0x38, 0x86, 0x86, 0xbc, 0xcd, 0x64,
	0x6a, 0x19, 0xa3, 0xff, 0xb4, 0xe8, 0x49, 0x97, 0xec, 0x92, 0xc3, 0xb5, 0x06, 0xd6, 0xde, 0x25,
	0xdd, 0x98, 0x63, 0x49, 0x9c, 0x8f, 0xa1, 0x10, 0xda, 0x13, 0xa8, 0xda, 0xc0, 0x9a, 0xb9, 0x74,
	0x25, 0x04, 0x6d, 0xa5, 0xb6, 0x6e, 0x05, 0x67, 0x67, 0xed, 0xc9, 0x2b, 0xa1, 0xd5, 0xab, 0x71,
	0x2e, 0xec, 0x89, 0xd0, 0xd5, 0xd9, 0x26, 0x76, 0xe0, 0x2f, 0xcd, 0xc7, 0x2f, 0x39, 0xaa, 0xed,
	0xfd, 0xc9, 0x25, 0x71, 0xee, 0x9d, 0x3d, 0xbe, 0x7f, 0xbd, 0x54, 0x02, 0xb6, 0x85, 0x5f, 0xc5,
	0x33, 0x04, 0xa9, 0xc8, 0xbb, 0x63, 0xa2, 0xe9, 0x7a, 0xcb, 0x9e, 0xcf, 0xc9, 0x51, 0x3d, 0xe9,
	0x0f, 0xb9, 0x33, 0x18, 0x72, 0xe7, 0x7b, 0xc8, 0x49, 0x7f, 0xc4, 0xc9, 0x60, 0xc4, 0xc9, 0xe7,
	0x88, 0x93, 0xe7, 0x82, 0x3b, 0x6f, 0x05, 0x27, 0xfd, 0x82, 0x93, 0x41, 0xc1, 0x9d, 0x8f, 0x82,
	0x3b, 0xcd, 0xfa, 0xaf, 0xd3, 0x3d, 0x74, 0x85, 0x56, 0x4a, 0x04, 0x53, 0x63, 0xb8, 0x7f, 0xc1,
	0xe4, 0xc5, 0xc2, 0x4b, 0xeb, 0x3f, 0x46, 0x7f, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x93, 0xa1,
	0x39, 0xff, 0x2b, 0x02, 0x00, 0x00,
}

func (this *SayHelloRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&swaggerfileserver.SayHelloRequest{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SayHelloReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&swaggerfileserver.SayHelloReply{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Content: "+fmt.Sprintf("%#v", this.Content)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGrace(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SayHelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayHelloRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayHelloRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGrace(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SayHelloReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayHelloReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayHelloReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintGrace(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintGrace(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrace(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SayHelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGrace(uint64(l))
	}
	return n
}

func (m *SayHelloReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovGrace(uint64(m.Code))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGrace(uint64(l))
	}
	return n
}

func sovGrace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrace(x uint64) (n int) {
	return sovGrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SayHelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrace
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
			return fmt.Errorf("proto: SayHelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayHelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrace
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
				return ErrInvalidLengthGrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SayHelloReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrace
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
			return fmt.Errorf("proto: SayHelloReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayHelloReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrace
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
				return ErrInvalidLengthGrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGrace
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrace
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
					return 0, ErrIntOverflowGrace
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
					return 0, ErrIntOverflowGrace
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
				return 0, ErrInvalidLengthGrace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrace = fmt.Errorf("proto: unexpected end of group")
)
