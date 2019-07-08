// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/sender_alias.proto

package entity

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SenderAlias_Status int32

const (
	SenderAlias_UNKNOWN        SenderAlias_Status = 0
	SenderAlias_SENT           SenderAlias_Status = 1
	SenderAlias_SENT_AND_ACKED SenderAlias_Status = 2
	SenderAlias_RECEIVED       SenderAlias_Status = 3
)

var SenderAlias_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SENT",
	2: "SENT_AND_ACKED",
	3: "RECEIVED",
}

var SenderAlias_Status_value = map[string]int32{
	"UNKNOWN":        0,
	"SENT":           1,
	"SENT_AND_ACKED": 2,
	"RECEIVED":       3,
}

func (x SenderAlias_Status) String() string {
	return proto.EnumName(SenderAlias_Status_name, int32(x))
}

func (SenderAlias_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5a4b242f459f89c7, []int{0, 0}
}

type SenderAlias struct {
	ID                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt            time.Time          `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	UpdatedAt            time.Time          `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at"`
	Status               SenderAlias_Status `protobuf:"varint,5,opt,name=status,proto3,enum=berty.entity.SenderAlias_Status" json:"status,omitempty"`
	OriginDeviceID       string             `protobuf:"bytes,6,opt,name=origin_device_id,json=originDeviceId,proto3" json:"origin_device_id,omitempty"`
	ContactID            string             `protobuf:"bytes,7,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	ConversationID       string             `protobuf:"bytes,8,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	AliasIdentifier      string             `protobuf:"bytes,9,opt,name=alias_identifier,json=aliasIdentifier,proto3" json:"alias_identifier,omitempty"`
	Used                 bool               `protobuf:"varint,10,opt,name=used,proto3" json:"used,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SenderAlias) Reset()         { *m = SenderAlias{} }
func (m *SenderAlias) String() string { return proto.CompactTextString(m) }
func (*SenderAlias) ProtoMessage()    {}
func (*SenderAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a4b242f459f89c7, []int{0}
}
func (m *SenderAlias) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SenderAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SenderAlias.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SenderAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SenderAlias.Merge(m, src)
}
func (m *SenderAlias) XXX_Size() int {
	return m.Size()
}
func (m *SenderAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_SenderAlias.DiscardUnknown(m)
}

var xxx_messageInfo_SenderAlias proto.InternalMessageInfo

func (m *SenderAlias) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *SenderAlias) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *SenderAlias) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *SenderAlias) GetStatus() SenderAlias_Status {
	if m != nil {
		return m.Status
	}
	return SenderAlias_UNKNOWN
}

func (m *SenderAlias) GetOriginDeviceID() string {
	if m != nil {
		return m.OriginDeviceID
	}
	return ""
}

func (m *SenderAlias) GetContactID() string {
	if m != nil {
		return m.ContactID
	}
	return ""
}

func (m *SenderAlias) GetConversationID() string {
	if m != nil {
		return m.ConversationID
	}
	return ""
}

func (m *SenderAlias) GetAliasIdentifier() string {
	if m != nil {
		return m.AliasIdentifier
	}
	return ""
}

func (m *SenderAlias) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

func init() {
	proto.RegisterEnum("berty.entity.SenderAlias_Status", SenderAlias_Status_name, SenderAlias_Status_value)
	proto.RegisterType((*SenderAlias)(nil), "berty.entity.SenderAlias")
}

func init() { proto.RegisterFile("entity/sender_alias.proto", fileDescriptor_5a4b242f459f89c7) }

var fileDescriptor_5a4b242f459f89c7 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0xef, 0x74, 0x6b, 0x37, 0x99, 0xae, 0xdd, 0x30, 0x88, 0xc4, 0x22, 0x4d, 0xe8, 0x29, 0xc2,
	0x92, 0xc0, 0x7a, 0x11, 0xf5, 0x92, 0x26, 0x39, 0x84, 0x85, 0x2c, 0xa4, 0xab, 0x82, 0x97, 0x90,
	0x66, 0x66, 0xe3, 0xe0, 0x36, 0x53, 0x26, 0xd3, 0x85, 0xbe, 0x85, 0x47, 0x9f, 0xc1, 0x27, 0xd9,
	0xa3, 0x4f, 0x10, 0x25, 0xbe, 0x81, 0x4f, 0x20, 0x33, 0x69, 0x35, 0x57, 0x6f, 0xdf, 0x7c, 0xbf,
	0x3f, 0xf3, 0xfd, 0xbe, 0x0f, 0x3e, 0x23, 0x95, 0xa0, 0x62, 0xef, 0xd5, 0xa4, 0xc2, 0x84, 0x67,
	0xf9, 0x1d, 0xcd, 0x6b, 0x77, 0xcb, 0x99, 0x60, 0xe8, 0x6c, 0x4d, 0xb8, 0xd8, 0xbb, 0x1d, 0x61,
	0x66, 0x95, 0x8c, 0x95, 0x77, 0xc4, 0x53, 0xd8, 0x7a, 0x77, 0xeb, 0x09, 0xba, 0x21, 0xb5, 0xc8,
	0x37, 0xdb, 0x8e, 0x3e, 0x7b, 0x52, 0xb2, 0x92, 0xa9, 0xd2, 0x93, 0x55, 0xd7, 0x5d, 0x7c, 0x1b,
	0xc1, 0xc9, 0x4a, 0x79, 0xfb, 0xd2, 0x1a, 0x5d, 0xc0, 0x21, 0xc5, 0x26, 0xb0, 0x81, 0xa3, 0x2f,
	0x9f, 0xb7, 0x8d, 0x35, 0x8c, 0xc3, 0xdf, 0x8d, 0x85, 0x4a, 0xc6, 0x37, 0xaf, 0x17, 0x5b, 0x4e,
	0x37, 0x39, 0xdf, 0x67, 0x9f, 0xc9, 0x7e, 0x91, 0x0e, 0x29, 0x46, 0x01, 0x84, 0x05, 0x27, 0xb9,
	0x20, 0x38, 0xcb, 0x85, 0x39, 0xb4, 0x81, 0x33, 0xb9, 0x9c, 0xb9, 0xdd, 0x24, 0xee, 0x71, 0x12,
	0xf7, 0xe6, 0x38, 0xc9, 0x52, 0x7b, 0x68, 0xac, 0xc1, 0x97, 0x1f, 0x16, 0x48, 0xf5, 0x83, 0xce,
	0x17, 0xd2, 0x64, 0xb7, 0xc5, 0x47, 0x93, 0x93, 0xff, 0x31, 0x39, 0xe8, 0x7c, 0x81, 0x5e, 0xc1,
	0x71, 0x2d, 0x72, 0xb1, 0xab, 0xcd, 0x47, 0x36, 0x70, 0xa6, 0x97, 0xb6, 0xdb, 0xdf, 0x8e, 0xdb,
	0x8b, 0xe8, 0xae, 0x14, 0x2f, 0x3d, 0xf0, 0xd1, 0x5b, 0x68, 0x30, 0x4e, 0x4b, 0x5a, 0x65, 0x98,
	0xdc, 0xd3, 0x82, 0x64, 0x14, 0x9b, 0x63, 0x95, 0x1f, 0xb5, 0x8d, 0x35, 0xbd, 0x56, 0x58, 0xa8,
	0xa0, 0x38, 0x4c, 0xa7, 0xac, 0xff, 0xc6, 0xe8, 0x02, 0xc2, 0x82, 0x55, 0x22, 0x2f, 0x84, 0xd4,
	0x9d, 0x2a, 0xdd, 0xe3, 0xb6, 0xb1, 0xf4, 0xa0, 0xeb, 0xc6, 0x61, 0xaa, 0x1f, 0x08, 0x31, 0x46,
	0x6f, 0xe0, 0x79, 0xc1, 0xaa, 0x7b, 0xc2, 0xeb, 0x5c, 0x50, 0x56, 0x49, 0x89, 0xf6, 0xef, 0xab,
	0xa0, 0x07, 0xc9, 0xaf, 0xfa, 0xd4, 0x18, 0xa3, 0x17, 0xd0, 0x50, 0xe7, 0xcf, 0x28, 0x96, 0xa9,
	0x6e, 0x29, 0xe1, 0xa6, 0x2e, 0xd5, 0xe9, 0xb9, 0xea, 0xc7, 0x7f, 0xdb, 0x08, 0xc1, 0xd1, 0xae,
	0x26, 0xd8, 0x84, 0x36, 0x70, 0xb4, 0x54, 0xd5, 0x0b, 0x1f, 0x8e, 0xbb, 0xe4, 0x68, 0x02, 0x4f,
	0xdf, 0x25, 0x57, 0xc9, 0xf5, 0x87, 0xc4, 0x18, 0x20, 0x0d, 0x8e, 0x56, 0x51, 0x72, 0x63, 0x00,
	0x84, 0xe0, 0x54, 0x56, 0x99, 0x9f, 0x84, 0x99, 0x1f, 0x5c, 0x45, 0xa1, 0x31, 0x44, 0x67, 0x50,
	0x4b, 0xa3, 0x20, 0x8a, 0xdf, 0x47, 0xa1, 0x71, 0xb2, 0x74, 0x1e, 0xda, 0x39, 0xf8, 0xde, 0xce,
	0xc1, 0xcf, 0x76, 0x0e, 0xbe, 0xfe, 0x9a, 0x0f, 0x3e, 0x3e, 0xed, 0xb6, 0x2c, 0x48, 0xf1, 0xc9,
	0x2b, 0x18, 0x27, 0x5e, 0xb7, 0xef, 0xf5, 0x58, 0xdd, 0xed, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x77, 0xec, 0xd3, 0xda, 0xbf, 0x02, 0x00, 0x00,
}

func (m *SenderAlias) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SenderAlias) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSenderAlias(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintSenderAlias(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err1 != nil {
		return 0, err1
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSenderAlias(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err2 != nil {
		return 0, err2
	}
	i += n2
	if m.Status != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSenderAlias(dAtA, i, uint64(m.Status))
	}
	if len(m.OriginDeviceID) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSenderAlias(dAtA, i, uint64(len(m.OriginDeviceID)))
		i += copy(dAtA[i:], m.OriginDeviceID)
	}
	if len(m.ContactID) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSenderAlias(dAtA, i, uint64(len(m.ContactID)))
		i += copy(dAtA[i:], m.ContactID)
	}
	if len(m.ConversationID) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintSenderAlias(dAtA, i, uint64(len(m.ConversationID)))
		i += copy(dAtA[i:], m.ConversationID)
	}
	if len(m.AliasIdentifier) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintSenderAlias(dAtA, i, uint64(len(m.AliasIdentifier)))
		i += copy(dAtA[i:], m.AliasIdentifier)
	}
	if m.Used {
		dAtA[i] = 0x50
		i++
		if m.Used {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSenderAlias(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SenderAlias) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovSenderAlias(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovSenderAlias(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovSenderAlias(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSenderAlias(uint64(m.Status))
	}
	l = len(m.OriginDeviceID)
	if l > 0 {
		n += 1 + l + sovSenderAlias(uint64(l))
	}
	l = len(m.ContactID)
	if l > 0 {
		n += 1 + l + sovSenderAlias(uint64(l))
	}
	l = len(m.ConversationID)
	if l > 0 {
		n += 1 + l + sovSenderAlias(uint64(l))
	}
	l = len(m.AliasIdentifier)
	if l > 0 {
		n += 1 + l + sovSenderAlias(uint64(l))
	}
	if m.Used {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSenderAlias(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSenderAlias(x uint64) (n int) {
	return sovSenderAlias(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SenderAlias) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSenderAlias
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
			return fmt.Errorf("proto: SenderAlias: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SenderAlias: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SenderAlias_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginDeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginDeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContactID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConversationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConversationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AliasIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
				return ErrInvalidLengthSenderAlias
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AliasIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Used", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSenderAlias
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
			m.Used = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSenderAlias(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSenderAlias
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSenderAlias
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
func skipSenderAlias(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSenderAlias
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
					return 0, ErrIntOverflowSenderAlias
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
					return 0, ErrIntOverflowSenderAlias
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
				return 0, ErrInvalidLengthSenderAlias
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSenderAlias
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSenderAlias
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
				next, err := skipSenderAlias(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSenderAlias
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
	ErrInvalidLengthSenderAlias = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSenderAlias   = fmt.Errorf("proto: integer overflow")
)
