// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: transportservice.proto

package TransportService

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

type BusScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber string `protobuf:"bytes,1,opt,name=BusNumber,proto3" json:"BusNumber,omitempty"`
}

func (x *BusScheduleRequest) Reset() {
	*x = BusScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusScheduleRequest) ProtoMessage() {}

func (x *BusScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transportservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusScheduleRequest.ProtoReflect.Descriptor instead.
func (*BusScheduleRequest) Descriptor() ([]byte, []int) {
	return file_transportservice_proto_rawDescGZIP(), []int{0}
}

func (x *BusScheduleRequest) GetBusNumber() string {
	if x != nil {
		return x.BusNumber
	}
	return ""
}

type BusScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations []*Station `protobuf:"bytes,1,rep,name=Stations,proto3" json:"Stations,omitempty"`
}

func (x *BusScheduleResponse) Reset() {
	*x = BusScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusScheduleResponse) ProtoMessage() {}

func (x *BusScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transportservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusScheduleResponse.ProtoReflect.Descriptor instead.
func (*BusScheduleResponse) Descriptor() ([]byte, []int) {
	return file_transportservice_proto_rawDescGZIP(), []int{1}
}

func (x *BusScheduleResponse) GetStations() []*Station {
	if x != nil {
		return x.Stations
	}
	return nil
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_transportservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_transportservice_proto_rawDescGZIP(), []int{2}
}

func (x *Station) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BusLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber string `protobuf:"bytes,1,opt,name=BusNumber,proto3" json:"BusNumber,omitempty"`
}

func (x *BusLocationRequest) Reset() {
	*x = BusLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusLocationRequest) ProtoMessage() {}

func (x *BusLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transportservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusLocationRequest.ProtoReflect.Descriptor instead.
func (*BusLocationRequest) Descriptor() ([]byte, []int) {
	return file_transportservice_proto_rawDescGZIP(), []int{3}
}

func (x *BusLocationRequest) GetBusNumber() string {
	if x != nil {
		return x.BusNumber
	}
	return ""
}

type TrafficJamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report string `protobuf:"bytes,1,opt,name=Report,proto3" json:"Report,omitempty"`
}

func (x *TrafficJamRequest) Reset() {
	*x = TrafficJamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficJamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficJamRequest) ProtoMessage() {}

func (x *TrafficJamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transportservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficJamRequest.ProtoReflect.Descriptor instead.
func (*TrafficJamRequest) Descriptor() ([]byte, []int) {
	return file_transportservice_proto_rawDescGZIP(), []int{4}
}

func (x *TrafficJamRequest) GetReport() string {
	if x != nil {
		return x.Report
	}
	return ""
}

type TrafficJamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsReported bool `protobuf:"varint,1,opt,name=IsReported,proto3" json:"IsReported,omitempty"`
}

func (x *TrafficJamResponse) Reset() {
	*x = TrafficJamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transportservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficJamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficJamResponse) ProtoMessage() {}

func (x *TrafficJamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transportservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficJamResponse.ProtoReflect.Descriptor instead.
func (*TrafficJamResponse) Descriptor() ([]byte, []int) {
	return file_transportservice_proto_rawDescGZIP(), []int{5}
}

func (x *TrafficJamResponse) GetIsReported() bool {
	if x != nil {
		return x.IsReported
	}
	return false
}

var File_transportservice_proto protoreflect.FileDescriptor

var file_transportservice_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x75, 0x73, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x75, 0x73,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x13, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1d, 0x0a, 0x07,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x12, 0x42,
	0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x2b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x34, 0x0a, 0x12,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x32, 0xf5, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x42,
	0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x10, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x12,
	0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transportservice_proto_rawDescOnce sync.Once
	file_transportservice_proto_rawDescData = file_transportservice_proto_rawDesc
)

func file_transportservice_proto_rawDescGZIP() []byte {
	file_transportservice_proto_rawDescOnce.Do(func() {
		file_transportservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_transportservice_proto_rawDescData)
	})
	return file_transportservice_proto_rawDescData
}

var file_transportservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transportservice_proto_goTypes = []interface{}{
	(*BusScheduleRequest)(nil),  // 0: genproto.BusScheduleRequest
	(*BusScheduleResponse)(nil), // 1: genproto.BusScheduleResponse
	(*Station)(nil),             // 2: genproto.Station
	(*BusLocationRequest)(nil),  // 3: genproto.BusLocationRequest
	(*TrafficJamRequest)(nil),   // 4: genproto.TrafficJamRequest
	(*TrafficJamResponse)(nil),  // 5: genproto.TrafficJamResponse
}
var file_transportservice_proto_depIdxs = []int32{
	2, // 0: genproto.BusScheduleResponse.Stations:type_name -> genproto.Station
	0, // 1: genproto.TransportService.GetBusSchedule:input_type -> genproto.BusScheduleRequest
	3, // 2: genproto.TransportService.TrackBusLocation:input_type -> genproto.BusLocationRequest
	4, // 3: genproto.TransportService.ReportTrafficJam:input_type -> genproto.TrafficJamRequest
	1, // 4: genproto.TransportService.GetBusSchedule:output_type -> genproto.BusScheduleResponse
	2, // 5: genproto.TransportService.TrackBusLocation:output_type -> genproto.Station
	5, // 6: genproto.TransportService.ReportTrafficJam:output_type -> genproto.TrafficJamResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transportservice_proto_init() }
func file_transportservice_proto_init() {
	if File_transportservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transportservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusScheduleRequest); i {
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
		file_transportservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusScheduleResponse); i {
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
		file_transportservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
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
		file_transportservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusLocationRequest); i {
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
		file_transportservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficJamRequest); i {
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
		file_transportservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficJamResponse); i {
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
			RawDescriptor: file_transportservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transportservice_proto_goTypes,
		DependencyIndexes: file_transportservice_proto_depIdxs,
		MessageInfos:      file_transportservice_proto_msgTypes,
	}.Build()
	File_transportservice_proto = out.File
	file_transportservice_proto_rawDesc = nil
	file_transportservice_proto_goTypes = nil
	file_transportservice_proto_depIdxs = nil
}
