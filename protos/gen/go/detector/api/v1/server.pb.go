// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: detector/api/v1/server.proto

package apiv1

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

type DetectFrameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame []byte `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
}

func (x *DetectFrameRequest) Reset() {
	*x = DetectFrameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_api_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectFrameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectFrameRequest) ProtoMessage() {}

func (x *DetectFrameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_detector_api_v1_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectFrameRequest.ProtoReflect.Descriptor instead.
func (*DetectFrameRequest) Descriptor() ([]byte, []int) {
	return file_detector_api_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *DetectFrameRequest) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

type Detection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label      string  `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Confidence float32 `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	X1         float32 `protobuf:"fixed32,3,opt,name=x1,proto3" json:"x1,omitempty"`
	Y1         float32 `protobuf:"fixed32,4,opt,name=y1,proto3" json:"y1,omitempty"`
	X2         float32 `protobuf:"fixed32,5,opt,name=x2,proto3" json:"x2,omitempty"`
	Y2         float32 `protobuf:"fixed32,6,opt,name=y2,proto3" json:"y2,omitempty"`
}

func (x *Detection) Reset() {
	*x = Detection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_api_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detection) ProtoMessage() {}

func (x *Detection) ProtoReflect() protoreflect.Message {
	mi := &file_detector_api_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detection.ProtoReflect.Descriptor instead.
func (*Detection) Descriptor() ([]byte, []int) {
	return file_detector_api_v1_server_proto_rawDescGZIP(), []int{1}
}

func (x *Detection) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Detection) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *Detection) GetX1() float32 {
	if x != nil {
		return x.X1
	}
	return 0
}

func (x *Detection) GetY1() float32 {
	if x != nil {
		return x.Y1
	}
	return 0
}

func (x *Detection) GetX2() float32 {
	if x != nil {
		return x.X2
	}
	return 0
}

func (x *Detection) GetY2() float32 {
	if x != nil {
		return x.Y2
	}
	return 0
}

type DetectFrameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detections []*Detection `protobuf:"bytes,1,rep,name=detections,proto3" json:"detections,omitempty"`
}

func (x *DetectFrameResponse) Reset() {
	*x = DetectFrameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detector_api_v1_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectFrameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectFrameResponse) ProtoMessage() {}

func (x *DetectFrameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_detector_api_v1_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectFrameResponse.ProtoReflect.Descriptor instead.
func (*DetectFrameResponse) Descriptor() ([]byte, []int) {
	return file_detector_api_v1_server_proto_rawDescGZIP(), []int{2}
}

func (x *DetectFrameResponse) GetDetections() []*Detection {
	if x != nil {
		return x.Detections
	}
	return nil
}

var File_detector_api_v1_server_proto protoreflect.FileDescriptor

var file_detector_api_v1_server_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22,
	0x2a, 0x0a, 0x12, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x09,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x78, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x78, 0x31, 0x12,
	0x0e, 0x0a, 0x02, 0x79, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x79, 0x31, 0x12,
	0x0e, 0x0a, 0x02, 0x78, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x78, 0x32, 0x12,
	0x0e, 0x0a, 0x02, 0x79, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x79, 0x32, 0x22,
	0x51, 0x0a, 0x13, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0x69, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x50,
	0x49, 0x12, 0x5a, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcc, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x6e, 0x64, 0x42, 0x6f, 0x62, 0x73, 0x59, 0x6f, 0x75, 0x72, 0x55, 0x6e, 0x63, 0x6c,
	0x65, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x44, 0x41, 0x58, 0xaa, 0x02, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_detector_api_v1_server_proto_rawDescOnce sync.Once
	file_detector_api_v1_server_proto_rawDescData = file_detector_api_v1_server_proto_rawDesc
)

func file_detector_api_v1_server_proto_rawDescGZIP() []byte {
	file_detector_api_v1_server_proto_rawDescOnce.Do(func() {
		file_detector_api_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_detector_api_v1_server_proto_rawDescData)
	})
	return file_detector_api_v1_server_proto_rawDescData
}

var file_detector_api_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_detector_api_v1_server_proto_goTypes = []interface{}{
	(*DetectFrameRequest)(nil),  // 0: detector.api.v1.DetectFrameRequest
	(*Detection)(nil),           // 1: detector.api.v1.Detection
	(*DetectFrameResponse)(nil), // 2: detector.api.v1.DetectFrameResponse
}
var file_detector_api_v1_server_proto_depIdxs = []int32{
	1, // 0: detector.api.v1.DetectFrameResponse.detections:type_name -> detector.api.v1.Detection
	0, // 1: detector.api.v1.DetectorAPI.DetectFrame:input_type -> detector.api.v1.DetectFrameRequest
	2, // 2: detector.api.v1.DetectorAPI.DetectFrame:output_type -> detector.api.v1.DetectFrameResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_detector_api_v1_server_proto_init() }
func file_detector_api_v1_server_proto_init() {
	if File_detector_api_v1_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detector_api_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectFrameRequest); i {
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
		file_detector_api_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detection); i {
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
		file_detector_api_v1_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectFrameResponse); i {
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
			RawDescriptor: file_detector_api_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_detector_api_v1_server_proto_goTypes,
		DependencyIndexes: file_detector_api_v1_server_proto_depIdxs,
		MessageInfos:      file_detector_api_v1_server_proto_msgTypes,
	}.Build()
	File_detector_api_v1_server_proto = out.File
	file_detector_api_v1_server_proto_rawDesc = nil
	file_detector_api_v1_server_proto_goTypes = nil
	file_detector_api_v1_server_proto_depIdxs = nil
}
