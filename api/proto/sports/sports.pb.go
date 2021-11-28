// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: sports/sports.proto

package sports

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sport type
type Sport_SportType int32

const (
	Sport_AFL      Sport_SportType = 0
	Sport_NRL      Sport_SportType = 1
	Sport_A_LEAGUE Sport_SportType = 2
	Sport_CRICKET  Sport_SportType = 3
)

// Enum value maps for Sport_SportType.
var (
	Sport_SportType_name = map[int32]string{
		0: "AFL",
		1: "NRL",
		2: "A_LEAGUE",
		3: "CRICKET",
	}
	Sport_SportType_value = map[string]int32{
		"AFL":      0,
		"NRL":      1,
		"A_LEAGUE": 2,
		"CRICKET":  3,
	}
)

func (x Sport_SportType) Enum() *Sport_SportType {
	p := new(Sport_SportType)
	*p = x
	return p
}

func (x Sport_SportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sport_SportType) Descriptor() protoreflect.EnumDescriptor {
	return file_sports_sports_proto_enumTypes[0].Descriptor()
}

func (Sport_SportType) Type() protoreflect.EnumType {
	return &file_sports_sports_proto_enumTypes[0]
}

func (x Sport_SportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sport_SportType.Descriptor instead.
func (Sport_SportType) EnumDescriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2, 0}
}

// Requests/Responses
type ListSportsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSportsRequest) Reset() {
	*x = ListSportsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSportsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSportsRequest) ProtoMessage() {}

func (x *ListSportsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSportsRequest.ProtoReflect.Descriptor instead.
func (*ListSportsRequest) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{0}
}

// Response to ListSports call.
type ListSportsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sports []*Sport `protobuf:"bytes,1,rep,name=sports,proto3" json:"sports,omitempty"`
}

func (x *ListSportsResponse) Reset() {
	*x = ListSportsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSportsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSportsResponse) ProtoMessage() {}

func (x *ListSportsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSportsResponse.ProtoReflect.Descriptor instead.
func (*ListSportsResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{1}
}

func (x *ListSportsResponse) GetSports() []*Sport {
	if x != nil {
		return x.Sports
	}
	return nil
}

// A sport resource.
type Sport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the sport.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the official name given to the sport.
	Name      string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SportType Sport_SportType `protobuf:"varint,4,opt,name=sportType,proto3,enum=sports.Sport_SportType" json:"sportType,omitempty"`
	// Visible represents whether or not the sport is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the sport is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
}

func (x *Sport) Reset() {
	*x = Sport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sport) ProtoMessage() {}

func (x *Sport) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sport.ProtoReflect.Descriptor instead.
func (*Sport) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2}
}

func (x *Sport) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sport) GetSportType() Sport_SportType {
	if x != nil {
		return x.SportType
	}
	return Sport_AFL
}

func (x *Sport) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Sport) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

var File_sports_sports_proto protoreflect.FileDescriptor

var file_sports_sports_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x86,
	0x02, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x09,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x4e, 0x0a,
	0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x38, 0x0a,
	0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x46,
	0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x52, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x5f, 0x4c, 0x45, 0x41, 0x47, 0x55, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52,
	0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x32, 0x69, 0x0a, 0x06, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x5f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x19, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sports_sports_proto_rawDescOnce sync.Once
	file_sports_sports_proto_rawDescData = file_sports_sports_proto_rawDesc
)

func file_sports_sports_proto_rawDescGZIP() []byte {
	file_sports_sports_proto_rawDescOnce.Do(func() {
		file_sports_sports_proto_rawDescData = protoimpl.X.CompressGZIP(file_sports_sports_proto_rawDescData)
	})
	return file_sports_sports_proto_rawDescData
}

var file_sports_sports_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sports_sports_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sports_sports_proto_goTypes = []interface{}{
	(Sport_SportType)(0),          // 0: sports.Sport.SportType
	(*ListSportsRequest)(nil),     // 1: sports.ListSportsRequest
	(*ListSportsResponse)(nil),    // 2: sports.ListSportsResponse
	(*Sport)(nil),                 // 3: sports.Sport
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_sports_sports_proto_depIdxs = []int32{
	3, // 0: sports.ListSportsResponse.sports:type_name -> sports.Sport
	0, // 1: sports.Sport.sportType:type_name -> sports.Sport.SportType
	4, // 2: sports.Sport.advertised_start_time:type_name -> google.protobuf.Timestamp
	1, // 3: sports.Sports.ListSports:input_type -> sports.ListSportsRequest
	2, // 4: sports.Sports.ListSports:output_type -> sports.ListSportsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sports_sports_proto_init() }
func file_sports_sports_proto_init() {
	if File_sports_sports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sports_sports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSportsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sports_sports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSportsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sports_sports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sports_sports_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sports_sports_proto_goTypes,
		DependencyIndexes: file_sports_sports_proto_depIdxs,
		EnumInfos:         file_sports_sports_proto_enumTypes,
		MessageInfos:      file_sports_sports_proto_msgTypes,
	}.Build()
	File_sports_sports_proto = out.File
	file_sports_sports_proto_rawDesc = nil
	file_sports_sports_proto_goTypes = nil
	file_sports_sports_proto_depIdxs = nil
}
