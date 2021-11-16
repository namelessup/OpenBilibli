// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: achievements.proto

package like

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import go_common_library_time "github.com/namelessup/bilibili/library/time"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActLikeAchievement struct {
	ID                   int64                       `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Name                 string                      `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	Icon                 string                      `protobuf:"bytes,3,opt,name=Icon,proto3" json:"icon"`
	Dic                  string                      `protobuf:"bytes,4,opt,name=Dic,proto3" json:"dic"`
	Unlock               int64                       `protobuf:"varint,5,opt,name=Unlock,proto3" json:"unlock"`
	Ctime                go_common_library_time.Time `protobuf:"varint,6,opt,name=Ctime,proto3,casttype=github.com/namelessup/bilibili/library/time.Time" json:"ctime"`
	Mtime                go_common_library_time.Time `protobuf:"varint,7,opt,name=Mtime,proto3,casttype=github.com/namelessup/bilibili/library/time.Time" json:"mtime"`
	Del                  int64                       `protobuf:"varint,8,opt,name=Del,proto3" json:"del"`
	Sid                  int64                       `protobuf:"varint,9,opt,name=Sid,proto3" json:"sid"`
	Image                string                      `protobuf:"bytes,10,opt,name=Image,proto3" json:"image"`
	Award                int64                       `protobuf:"varint,11,opt,name=Award,proto3" json:"award"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ActLikeAchievement) Reset()         { *m = ActLikeAchievement{} }
func (m *ActLikeAchievement) String() string { return proto.CompactTextString(m) }
func (*ActLikeAchievement) ProtoMessage()    {}
func (*ActLikeAchievement) Descriptor() ([]byte, []int) {
	return fileDescriptor_achievements_5f3866f0cfaf036f, []int{0}
}
func (m *ActLikeAchievement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActLikeAchievement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActLikeAchievement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ActLikeAchievement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActLikeAchievement.Merge(dst, src)
}
func (m *ActLikeAchievement) XXX_Size() int {
	return m.Size()
}
func (m *ActLikeAchievement) XXX_DiscardUnknown() {
	xxx_messageInfo_ActLikeAchievement.DiscardUnknown(m)
}

var xxx_messageInfo_ActLikeAchievement proto.InternalMessageInfo

type Achievements struct {
	Achievements         []*ActLikeAchievement `protobuf:"bytes,1,rep,name=achievements" json:"achievements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Achievements) Reset()         { *m = Achievements{} }
func (m *Achievements) String() string { return proto.CompactTextString(m) }
func (*Achievements) ProtoMessage()    {}
func (*Achievements) Descriptor() ([]byte, []int) {
	return fileDescriptor_achievements_5f3866f0cfaf036f, []int{1}
}
func (m *Achievements) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Achievements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Achievements.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Achievements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Achievements.Merge(dst, src)
}
func (m *Achievements) XXX_Size() int {
	return m.Size()
}
func (m *Achievements) XXX_DiscardUnknown() {
	xxx_messageInfo_Achievements.DiscardUnknown(m)
}

var xxx_messageInfo_Achievements proto.InternalMessageInfo

type ActLikeUserAchievement struct {
	ID                   int64                       `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Aid                  int64                       `protobuf:"varint,2,opt,name=Aid,proto3" json:"aid"`
	Ctime                go_common_library_time.Time `protobuf:"varint,3,opt,name=Ctime,proto3,casttype=github.com/namelessup/bilibili/library/time.Time" json:"ctime"`
	Mtime                go_common_library_time.Time `protobuf:"varint,4,opt,name=Mtime,proto3,casttype=github.com/namelessup/bilibili/library/time.Time" json:"mtime"`
	Del                  int64                       `protobuf:"varint,5,opt,name=Del,proto3" json:"del"`
	Mid                  int64                       `protobuf:"varint,6,opt,name=Mid,proto3" json:"mid"`
	Sid                  int64                       `protobuf:"varint,7,opt,name=Sid,proto3" json:"sid"`
	Award                int64                       `protobuf:"varint,8,opt,name=Award,proto3" json:"award"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ActLikeUserAchievement) Reset()         { *m = ActLikeUserAchievement{} }
func (m *ActLikeUserAchievement) String() string { return proto.CompactTextString(m) }
func (*ActLikeUserAchievement) ProtoMessage()    {}
func (*ActLikeUserAchievement) Descriptor() ([]byte, []int) {
	return fileDescriptor_achievements_5f3866f0cfaf036f, []int{2}
}
func (m *ActLikeUserAchievement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActLikeUserAchievement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActLikeUserAchievement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ActLikeUserAchievement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActLikeUserAchievement.Merge(dst, src)
}
func (m *ActLikeUserAchievement) XXX_Size() int {
	return m.Size()
}
func (m *ActLikeUserAchievement) XXX_DiscardUnknown() {
	xxx_messageInfo_ActLikeUserAchievement.DiscardUnknown(m)
}

var xxx_messageInfo_ActLikeUserAchievement proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ActLikeAchievement)(nil), "activity.service.ActLikeAchievement")
	proto.RegisterType((*Achievements)(nil), "activity.service.Achievements")
	proto.RegisterType((*ActLikeUserAchievement)(nil), "activity.service.ActLikeUserAchievement")
}
func (m *ActLikeAchievement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActLikeAchievement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.ID))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Icon) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(len(m.Icon)))
		i += copy(dAtA[i:], m.Icon)
	}
	if len(m.Dic) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(len(m.Dic)))
		i += copy(dAtA[i:], m.Dic)
	}
	if m.Unlock != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Unlock))
	}
	if m.Ctime != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Ctime))
	}
	if m.Mtime != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Mtime))
	}
	if m.Del != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Del))
	}
	if m.Sid != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Sid))
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if m.Award != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Award))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Achievements) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Achievements) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Achievements) > 0 {
		for _, msg := range m.Achievements {
			dAtA[i] = 0xa
			i++
			i = encodeVarintAchievements(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ActLikeUserAchievement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActLikeUserAchievement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.ID))
	}
	if m.Aid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Aid))
	}
	if m.Ctime != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Ctime))
	}
	if m.Mtime != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Mtime))
	}
	if m.Del != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Del))
	}
	if m.Mid != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Mid))
	}
	if m.Sid != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Sid))
	}
	if m.Award != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintAchievements(dAtA, i, uint64(m.Award))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAchievements(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ActLikeAchievement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovAchievements(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAchievements(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovAchievements(uint64(l))
	}
	l = len(m.Dic)
	if l > 0 {
		n += 1 + l + sovAchievements(uint64(l))
	}
	if m.Unlock != 0 {
		n += 1 + sovAchievements(uint64(m.Unlock))
	}
	if m.Ctime != 0 {
		n += 1 + sovAchievements(uint64(m.Ctime))
	}
	if m.Mtime != 0 {
		n += 1 + sovAchievements(uint64(m.Mtime))
	}
	if m.Del != 0 {
		n += 1 + sovAchievements(uint64(m.Del))
	}
	if m.Sid != 0 {
		n += 1 + sovAchievements(uint64(m.Sid))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovAchievements(uint64(l))
	}
	if m.Award != 0 {
		n += 1 + sovAchievements(uint64(m.Award))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Achievements) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Achievements) > 0 {
		for _, e := range m.Achievements {
			l = e.Size()
			n += 1 + l + sovAchievements(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ActLikeUserAchievement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovAchievements(uint64(m.ID))
	}
	if m.Aid != 0 {
		n += 1 + sovAchievements(uint64(m.Aid))
	}
	if m.Ctime != 0 {
		n += 1 + sovAchievements(uint64(m.Ctime))
	}
	if m.Mtime != 0 {
		n += 1 + sovAchievements(uint64(m.Mtime))
	}
	if m.Del != 0 {
		n += 1 + sovAchievements(uint64(m.Del))
	}
	if m.Mid != 0 {
		n += 1 + sovAchievements(uint64(m.Mid))
	}
	if m.Sid != 0 {
		n += 1 + sovAchievements(uint64(m.Sid))
	}
	if m.Award != 0 {
		n += 1 + sovAchievements(uint64(m.Award))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAchievements(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAchievements(x uint64) (n int) {
	return sovAchievements(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActLikeAchievement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievements
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActLikeAchievement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActLikeAchievement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievements
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievements
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievements
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unlock", wireType)
			}
			m.Unlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Unlock |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ctime", wireType)
			}
			m.Ctime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ctime |= (go_common_library_time.Time(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mtime", wireType)
			}
			m.Mtime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mtime |= (go_common_library_time.Time(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Del", wireType)
			}
			m.Del = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Del |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sid", wireType)
			}
			m.Sid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAchievements
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Award", wireType)
			}
			m.Award = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Award |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAchievements(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievements
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
func (m *Achievements) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievements
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Achievements: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Achievements: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Achievements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAchievements
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Achievements = append(m.Achievements, &ActLikeAchievement{})
			if err := m.Achievements[len(m.Achievements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAchievements(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievements
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
func (m *ActLikeUserAchievement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAchievements
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActLikeUserAchievement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActLikeUserAchievement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aid", wireType)
			}
			m.Aid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Aid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ctime", wireType)
			}
			m.Ctime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ctime |= (go_common_library_time.Time(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mtime", wireType)
			}
			m.Mtime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mtime |= (go_common_library_time.Time(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Del", wireType)
			}
			m.Del = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Del |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sid", wireType)
			}
			m.Sid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Award", wireType)
			}
			m.Award = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAchievements
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Award |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAchievements(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAchievements
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
func skipAchievements(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAchievements
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
					return 0, ErrIntOverflowAchievements
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
					return 0, ErrIntOverflowAchievements
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAchievements
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAchievements
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
				next, err := skipAchievements(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthAchievements = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAchievements   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("achievements.proto", fileDescriptor_achievements_5f3866f0cfaf036f) }

var fileDescriptor_achievements_5f3866f0cfaf036f = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xcf, 0xf1, 0x9f, 0x24, 0x7b, 0x57, 0xa0, 0x2d, 0x4e, 0x7b, 0x80, 0xec, 0x28, 0xa2,
	0x48, 0x73, 0x8e, 0x04, 0x3d, 0x92, 0x43, 0x0a, 0x22, 0x11, 0x8a, 0x85, 0x93, 0x10, 0x9d, 0xb3,
	0x5e, 0x7c, 0xa3, 0x78, 0xbd, 0xc8, 0x76, 0x82, 0xee, 0x4d, 0xe8, 0x78, 0x12, 0xfa, 0x2b, 0xef,
	0x09, 0x2c, 0x2e, 0x74, 0x7e, 0x04, 0x2a, 0xb4, 0xe3, 0x20, 0x1f, 0x90, 0xe2, 0xa4, 0xeb, 0xbc,
	0xdf, 0x37, 0xdf, 0xee, 0x8c, 0x7e, 0x63, 0x42, 0x63, 0x71, 0x09, 0x72, 0x2b, 0x95, 0xcc, 0xab,
	0x32, 0xfc, 0x5c, 0xe8, 0x4a, 0xd3, 0x47, 0xb1, 0xa8, 0x60, 0x0b, 0xd5, 0x55, 0x58, 0xca, 0x62,
	0x0b, 0x42, 0x3e, 0x3e, 0x4f, 0xa1, 0xba, 0xdc, 0xac, 0x42, 0xa1, 0xd5, 0x34, 0xd5, 0xa9, 0x9e,
	0x62, 0xe1, 0x6a, 0xf3, 0x09, 0x4f, 0x78, 0xc0, 0xaf, 0xf6, 0x82, 0xf1, 0x37, 0x9b, 0xd0, 0x48,
	0x54, 0x6f, 0x60, 0x2d, 0xa3, 0xee, 0x7a, 0x7a, 0x4a, 0x7a, 0x8b, 0x39, 0xb3, 0x46, 0xd6, 0xc4,
	0x9e, 0x79, 0x4d, 0x1d, 0xf4, 0x20, 0xe1, 0xbd, 0xc5, 0x9c, 0x3e, 0x25, 0xce, 0xdb, 0x58, 0x49,
	0xd6, 0x1b, 0x59, 0x93, 0xe1, 0x6c, 0xd0, 0xd4, 0x81, 0x93, 0xc7, 0x4a, 0x72, 0x54, 0x8d, 0xbb,
	0x10, 0x3a, 0x67, 0x76, 0xe7, 0x82, 0xd0, 0x39, 0x47, 0x95, 0x9e, 0x11, 0x7b, 0x0e, 0x82, 0x39,
	0x68, 0xf6, 0x9b, 0x3a, 0xb0, 0x13, 0x10, 0xdc, 0x68, 0x74, 0x4c, 0xbc, 0x8b, 0x3c, 0xd3, 0x62,
	0xcd, 0x5c, 0x7c, 0x92, 0x34, 0x75, 0xe0, 0x6d, 0x50, 0xe1, 0x7b, 0x87, 0xbe, 0x24, 0xee, 0xab,
	0x0a, 0x94, 0x64, 0x1e, 0x96, 0x4c, 0x9a, 0x3a, 0x70, 0x85, 0x11, 0x7e, 0xd5, 0xc1, 0x93, 0x54,
	0x9f, 0x0b, 0xad, 0x94, 0xce, 0xa7, 0x19, 0xac, 0x8a, 0xb8, 0xb8, 0x9a, 0x1a, 0x27, 0x7c, 0x0f,
	0x4a, 0xf2, 0x36, 0x66, 0xf2, 0x4b, 0xcc, 0xf7, 0xbb, 0xbc, 0xba, 0x57, 0x1e, 0x63, 0xd8, 0xbe,
	0xcc, 0xd8, 0x00, 0xd3, 0x6d, 0xfb, 0x32, 0xe3, 0x46, 0x33, 0xd6, 0x3b, 0x48, 0xd8, 0xb0, 0xb3,
	0x4a, 0x48, 0xb8, 0xd1, 0x68, 0x40, 0xdc, 0x85, 0x8a, 0x53, 0xc9, 0x08, 0x8e, 0x3d, 0x34, 0xaf,
	0x82, 0x11, 0x78, 0xab, 0x9b, 0x82, 0xe8, 0x4b, 0x5c, 0x24, 0xec, 0x18, 0xd3, 0x58, 0x10, 0x1b,
	0x81, 0xb7, 0xfa, 0xf8, 0x03, 0x39, 0xb9, 0x43, 0xa6, 0xa4, 0xaf, 0xc9, 0xc9, 0xdd, 0x45, 0x60,
	0xd6, 0xc8, 0x9e, 0x1c, 0x3f, 0x7f, 0x16, 0xfe, 0xbb, 0x09, 0xe1, 0xff, 0x58, 0xf9, 0x5f, 0xc9,
	0xf1, 0xf7, 0x1e, 0x39, 0xdd, 0x17, 0x5d, 0x94, 0xb2, 0xb8, 0x0f, 0xff, 0x33, 0x62, 0x47, 0x90,
	0x20, 0xfe, 0xfd, 0xa4, 0xb1, 0x99, 0x34, 0x82, 0xa4, 0xe3, 0x63, 0x3f, 0x90, 0x8f, 0xf3, 0x20,
	0x3e, 0xee, 0x61, 0x3e, 0x4b, 0x48, 0xf6, 0x8b, 0x83, 0x96, 0x32, 0x5d, 0x2f, 0x21, 0xf9, 0x83,
	0xae, 0x7f, 0x18, 0x5d, 0x4b, 0x66, 0x70, 0x98, 0xcc, 0xcc, 0xbf, 0xbe, 0xf5, 0x8f, 0x6e, 0x6e,
	0xfd, 0xa3, 0xeb, 0x9d, 0x6f, 0xdd, 0xec, 0x7c, 0xeb, 0xc7, 0xce, 0xb7, 0xbe, 0xfe, 0xf4, 0x8f,
	0x3e, 0x3a, 0x19, 0xac, 0xe5, 0xca, 0xc3, 0x5f, 0xec, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x49, 0x0a, 0x1d, 0x9a, 0xb9, 0x03, 0x00, 0x00,
}
