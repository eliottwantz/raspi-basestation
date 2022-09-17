// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: pb/sensorstate.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BrakeManager_States int32

const (
	BrakeManager_INIT              BrakeManager_States = 0
	BrakeManager_IDLE              BrakeManager_States = 1
	BrakeManager_CALIBRATION       BrakeManager_States = 2
	BrakeManager_READY             BrakeManager_States = 3
	BrakeManager_BRAKING           BrakeManager_States = 4
	BrakeManager_EMERGENCY_BRAKING BrakeManager_States = 5
)

// Enum value maps for BrakeManager_States.
var (
	BrakeManager_States_name = map[int32]string{
		0: "INIT",
		1: "IDLE",
		2: "CALIBRATION",
		3: "READY",
		4: "BRAKING",
		5: "EMERGENCY_BRAKING",
	}
	BrakeManager_States_value = map[string]int32{
		"INIT":              0,
		"IDLE":              1,
		"CALIBRATION":       2,
		"READY":             3,
		"BRAKING":           4,
		"EMERGENCY_BRAKING": 5,
	}
)

func (x BrakeManager_States) Enum() *BrakeManager_States {
	p := new(BrakeManager_States)
	*p = x
	return p
}

func (x BrakeManager_States) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrakeManager_States) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_sensorstate_proto_enumTypes[0].Descriptor()
}

func (BrakeManager_States) Type() protoreflect.EnumType {
	return &file_pb_sensorstate_proto_enumTypes[0]
}

func (x BrakeManager_States) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrakeManager_States.Descriptor instead.
func (BrakeManager_States) EnumDescriptor() ([]byte, []int) {
	return file_pb_sensorstate_proto_rawDescGZIP(), []int{0, 0}
}

type MainComputer_States int32

const (
	MainComputer_SLEEP             MainComputer_States = 0
	MainComputer_STATIC_FAULT      MainComputer_States = 1
	MainComputer_DYNAMIC_FAULT     MainComputer_States = 2
	MainComputer_SAFE_TO_APPROACH  MainComputer_States = 3
	MainComputer_INITIALIZATION    MainComputer_States = 4
	MainComputer_SAFE_TO_LAUNCH    MainComputer_States = 5
	MainComputer_ACCELERATING      MainComputer_States = 6
	MainComputer_EMERGENCY_BRAKING MainComputer_States = 7
	MainComputer_END_BRAKES        MainComputer_States = 8
)

// Enum value maps for MainComputer_States.
var (
	MainComputer_States_name = map[int32]string{
		0: "SLEEP",
		1: "STATIC_FAULT",
		2: "DYNAMIC_FAULT",
		3: "SAFE_TO_APPROACH",
		4: "INITIALIZATION",
		5: "SAFE_TO_LAUNCH",
		6: "ACCELERATING",
		7: "EMERGENCY_BRAKING",
		8: "END_BRAKES",
	}
	MainComputer_States_value = map[string]int32{
		"SLEEP":             0,
		"STATIC_FAULT":      1,
		"DYNAMIC_FAULT":     2,
		"SAFE_TO_APPROACH":  3,
		"INITIALIZATION":    4,
		"SAFE_TO_LAUNCH":    5,
		"ACCELERATING":      6,
		"EMERGENCY_BRAKING": 7,
		"END_BRAKES":        8,
	}
)

func (x MainComputer_States) Enum() *MainComputer_States {
	p := new(MainComputer_States)
	*p = x
	return p
}

func (x MainComputer_States) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MainComputer_States) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_sensorstate_proto_enumTypes[1].Descriptor()
}

func (MainComputer_States) Type() protoreflect.EnumType {
	return &file_pb_sensorstate_proto_enumTypes[1]
}

func (x MainComputer_States) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MainComputer_States.Descriptor instead.
func (MainComputer_States) EnumDescriptor() ([]byte, []int) {
	return file_pb_sensorstate_proto_rawDescGZIP(), []int{1, 0}
}

type BrakeManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State BrakeManager_States `protobuf:"varint,1,opt,name=state,proto3,enum=pb.BrakeManager_States" json:"state,omitempty"`
}

func (x *BrakeManager) Reset() {
	*x = BrakeManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sensorstate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrakeManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrakeManager) ProtoMessage() {}

func (x *BrakeManager) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sensorstate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrakeManager.ProtoReflect.Descriptor instead.
func (*BrakeManager) Descriptor() ([]byte, []int) {
	return file_pb_sensorstate_proto_rawDescGZIP(), []int{0}
}

func (x *BrakeManager) GetState() BrakeManager_States {
	if x != nil {
		return x.State
	}
	return BrakeManager_INIT
}

type MainComputer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State MainComputer_States `protobuf:"varint,1,opt,name=state,proto3,enum=pb.MainComputer_States" json:"state,omitempty"`
}

func (x *MainComputer) Reset() {
	*x = MainComputer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sensorstate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainComputer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainComputer) ProtoMessage() {}

func (x *MainComputer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sensorstate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainComputer.ProtoReflect.Descriptor instead.
func (*MainComputer) Descriptor() ([]byte, []int) {
	return file_pb_sensorstate_proto_rawDescGZIP(), []int{1}
}

func (x *MainComputer) GetState() MainComputer_States {
	if x != nil {
		return x.State
	}
	return MainComputer_SLEEP
}

type SensorState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainComputer *MainComputer `protobuf:"bytes,1,opt,name=main_computer,json=mainComputer,proto3" json:"main_computer,omitempty"`
	BrakeManager *BrakeManager `protobuf:"bytes,2,opt,name=brake_manager,json=brakeManager,proto3" json:"brake_manager,omitempty"`
}

func (x *SensorState) Reset() {
	*x = SensorState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_sensorstate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorState) ProtoMessage() {}

func (x *SensorState) ProtoReflect() protoreflect.Message {
	mi := &file_pb_sensorstate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorState.ProtoReflect.Descriptor instead.
func (*SensorState) Descriptor() ([]byte, []int) {
	return file_pb_sensorstate_proto_rawDescGZIP(), []int{2}
}

func (x *SensorState) GetMainComputer() *MainComputer {
	if x != nil {
		return x.MainComputer
	}
	return nil
}

func (x *SensorState) GetBrakeManager() *BrakeManager {
	if x != nil {
		return x.BrakeManager
	}
	return nil
}

var File_pb_sensorstate_proto protoreflect.FileDescriptor

var file_pb_sensorstate_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x9b, 0x01, 0x0a, 0x0c, 0x42,
	0x72, 0x61, 0x6b, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x42, 0x72, 0x61, 0x6b, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5c, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x4c, 0x49,
	0x42, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x52, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10,
	0x04, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x42,
	0x52, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x22, 0xef, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x69,
	0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61,
	0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4c, 0x45, 0x45, 0x50, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x5f, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x41, 0x46, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x41, 0x43, 0x48, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x41, 0x46, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x4c, 0x41, 0x55, 0x4e, 0x43, 0x48, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x45, 0x4c, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59,
	0x5f, 0x42, 0x52, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e,
	0x44, 0x5f, 0x42, 0x52, 0x41, 0x4b, 0x45, 0x53, 0x10, 0x08, 0x22, 0x7b, 0x0a, 0x0b, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x72, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x35, 0x0a, 0x0d, 0x62, 0x72, 0x61, 0x6b, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x72, 0x61,
	0x6b, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x0c, 0x62, 0x72, 0x61, 0x6b, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_sensorstate_proto_rawDescOnce sync.Once
	file_pb_sensorstate_proto_rawDescData = file_pb_sensorstate_proto_rawDesc
)

func file_pb_sensorstate_proto_rawDescGZIP() []byte {
	file_pb_sensorstate_proto_rawDescOnce.Do(func() {
		file_pb_sensorstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_sensorstate_proto_rawDescData)
	})
	return file_pb_sensorstate_proto_rawDescData
}

var file_pb_sensorstate_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_sensorstate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_sensorstate_proto_goTypes = []interface{}{
	(BrakeManager_States)(0), // 0: pb.BrakeManager.States
	(MainComputer_States)(0), // 1: pb.MainComputer.States
	(*BrakeManager)(nil),     // 2: pb.BrakeManager
	(*MainComputer)(nil),     // 3: pb.MainComputer
	(*SensorState)(nil),      // 4: pb.SensorState
}
var file_pb_sensorstate_proto_depIdxs = []int32{
	0, // 0: pb.BrakeManager.state:type_name -> pb.BrakeManager.States
	1, // 1: pb.MainComputer.state:type_name -> pb.MainComputer.States
	3, // 2: pb.SensorState.main_computer:type_name -> pb.MainComputer
	2, // 3: pb.SensorState.brake_manager:type_name -> pb.BrakeManager
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_sensorstate_proto_init() }
func file_pb_sensorstate_proto_init() {
	if File_pb_sensorstate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_sensorstate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrakeManager); i {
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
		file_pb_sensorstate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainComputer); i {
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
		file_pb_sensorstate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorState); i {
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
			RawDescriptor: file_pb_sensorstate_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_sensorstate_proto_goTypes,
		DependencyIndexes: file_pb_sensorstate_proto_depIdxs,
		EnumInfos:         file_pb_sensorstate_proto_enumTypes,
		MessageInfos:      file_pb_sensorstate_proto_msgTypes,
	}.Build()
	File_pb_sensorstate_proto = out.File
	file_pb_sensorstate_proto_rawDesc = nil
	file_pb_sensorstate_proto_goTypes = nil
	file_pb_sensorstate_proto_depIdxs = nil
}
