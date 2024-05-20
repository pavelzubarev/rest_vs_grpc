// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: polygon.proto

package polygon

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_polygon_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type PolygonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *PolygonRequest) Reset() {
	*x = PolygonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonRequest) ProtoMessage() {}

func (x *PolygonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonRequest.ProtoReflect.Descriptor instead.
func (*PolygonRequest) Descriptor() ([]byte, []int) {
	return file_polygon_proto_rawDescGZIP(), []int{1}
}

func (x *PolygonRequest) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type PolygonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Area float64 `protobuf:"fixed64,1,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *PolygonResponse) Reset() {
	*x = PolygonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonResponse) ProtoMessage() {}

func (x *PolygonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonResponse.ProtoReflect.Descriptor instead.
func (*PolygonResponse) Descriptor() ([]byte, []int) {
	return file_polygon_proto_rawDescGZIP(), []int{2}
}

func (x *PolygonResponse) GetArea() float64 {
	if x != nil {
		return x.Area
	}
	return 0
}

var File_polygon_proto protoreflect.FileDescriptor

var file_polygon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x38, 0x0a,
	0x0e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x32, 0x54,
	0x0a, 0x0e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x79,
	0x67, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x6c,
	0x79, 0x67, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2e, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x3b, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_polygon_proto_rawDescOnce sync.Once
	file_polygon_proto_rawDescData = file_polygon_proto_rawDesc
)

func file_polygon_proto_rawDescGZIP() []byte {
	file_polygon_proto_rawDescOnce.Do(func() {
		file_polygon_proto_rawDescData = protoimpl.X.CompressGZIP(file_polygon_proto_rawDescData)
	})
	return file_polygon_proto_rawDescData
}

var file_polygon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_polygon_proto_goTypes = []interface{}{
	(*Point)(nil),           // 0: polygon.Point
	(*PolygonRequest)(nil),  // 1: polygon.PolygonRequest
	(*PolygonResponse)(nil), // 2: polygon.PolygonResponse
}
var file_polygon_proto_depIdxs = []int32{
	0, // 0: polygon.PolygonRequest.points:type_name -> polygon.Point
	1, // 1: polygon.PolygonService.CalculateArea:input_type -> polygon.PolygonRequest
	2, // 2: polygon.PolygonService.CalculateArea:output_type -> polygon.PolygonResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_polygon_proto_init() }
func file_polygon_proto_init() {
	if File_polygon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_polygon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_polygon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonRequest); i {
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
		file_polygon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonResponse); i {
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
			RawDescriptor: file_polygon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_polygon_proto_goTypes,
		DependencyIndexes: file_polygon_proto_depIdxs,
		MessageInfos:      file_polygon_proto_msgTypes,
	}.Build()
	File_polygon_proto = out.File
	file_polygon_proto_rawDesc = nil
	file_polygon_proto_goTypes = nil
	file_polygon_proto_depIdxs = nil
}