// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: habitantes.proto

package misc

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

type Habitante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PosX   int32 `protobuf:"varint,1,opt,name=posX,proto3" json:"posX,omitempty"`
	PosY   int32 `protobuf:"varint,2,opt,name=posY,proto3" json:"posY,omitempty"`
	Estado int32 `protobuf:"varint,3,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Habitante) Reset() {
	*x = Habitante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitantes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habitante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habitante) ProtoMessage() {}

func (x *Habitante) ProtoReflect() protoreflect.Message {
	mi := &file_habitantes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habitante.ProtoReflect.Descriptor instead.
func (*Habitante) Descriptor() ([]byte, []int) {
	return file_habitantes_proto_rawDescGZIP(), []int{0}
}

func (x *Habitante) GetPosX() int32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *Habitante) GetPosY() int32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *Habitante) GetEstado() int32 {
	if x != nil {
		return x.Estado
	}
	return 0
}

type InicializadorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumHabitantes int32 `protobuf:"varint,1,opt,name=num_habitantes,json=numHabitantes,proto3" json:"num_habitantes,omitempty"`
}

func (x *InicializadorRequest) Reset() {
	*x = InicializadorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitantes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InicializadorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InicializadorRequest) ProtoMessage() {}

func (x *InicializadorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_habitantes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InicializadorRequest.ProtoReflect.Descriptor instead.
func (*InicializadorRequest) Descriptor() ([]byte, []int) {
	return file_habitantes_proto_rawDescGZIP(), []int{1}
}

func (x *InicializadorRequest) GetNumHabitantes() int32 {
	if x != nil {
		return x.NumHabitantes
	}
	return 0
}

type InicializadorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitantesInicial []*Habitante `protobuf:"bytes,1,rep,name=habitantes_inicial,json=habitantesInicial,proto3" json:"habitantes_inicial,omitempty"`
}

func (x *InicializadorResponse) Reset() {
	*x = InicializadorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitantes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InicializadorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InicializadorResponse) ProtoMessage() {}

func (x *InicializadorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_habitantes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InicializadorResponse.ProtoReflect.Descriptor instead.
func (*InicializadorResponse) Descriptor() ([]byte, []int) {
	return file_habitantes_proto_rawDescGZIP(), []int{2}
}

func (x *InicializadorResponse) GetHabitantesInicial() []*Habitante {
	if x != nil {
		return x.HabitantesInicial
	}
	return nil
}

type EstadoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request int32 `protobuf:"varint,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *EstadoRequest) Reset() {
	*x = EstadoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitantes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstadoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoRequest) ProtoMessage() {}

func (x *EstadoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_habitantes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoRequest.ProtoReflect.Descriptor instead.
func (*EstadoRequest) Descriptor() ([]byte, []int) {
	return file_habitantes_proto_rawDescGZIP(), []int{3}
}

func (x *EstadoRequest) GetRequest() int32 {
	if x != nil {
		return x.Request
	}
	return 0
}

type EstadoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EstadoHabitante []*Habitante `protobuf:"bytes,1,rep,name=estado_habitante,json=estadoHabitante,proto3" json:"estado_habitante,omitempty"`
}

func (x *EstadoResponse) Reset() {
	*x = EstadoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habitantes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstadoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoResponse) ProtoMessage() {}

func (x *EstadoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_habitantes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoResponse.ProtoReflect.Descriptor instead.
func (*EstadoResponse) Descriptor() ([]byte, []int) {
	return file_habitantes_proto_rawDescGZIP(), []int{4}
}

func (x *EstadoResponse) GetEstadoHabitante() []*Habitante {
	if x != nil {
		return x.EstadoHabitante
	}
	return nil
}

var File_habitantes_proto protoreflect.FileDescriptor

var file_habitantes_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x22, 0x4b,
	0x0a, 0x09, 0x48, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x59, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x3d, 0x0a, 0x14, 0x49,
	0x6e, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x5f, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x61, 0x6e, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x15, 0x49, 0x6e,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x12, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65,
	0x73, 0x5f, 0x69, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x62,
	0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x52, 0x11, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74,
	0x65, 0x73, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x22, 0x29, 0x0a, 0x0d, 0x45, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x0e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x5f, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x52, 0x0f, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x32, 0x95, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x69, 0x6f, 0x48, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x12,
	0x62, 0x0a, 0x17, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x72,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x68, 0x61, 0x62,
	0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x19, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61,
	0x6e, 0x74, 0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69, 0x72, 0x52, 0x65,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x12, 0x19, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74,
	0x65, 0x73, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x2e, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x07, 0x5a, 0x05, 0x6d, 0x69, 0x73, 0x63, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_habitantes_proto_rawDescOnce sync.Once
	file_habitantes_proto_rawDescData = file_habitantes_proto_rawDesc
)

func file_habitantes_proto_rawDescGZIP() []byte {
	file_habitantes_proto_rawDescOnce.Do(func() {
		file_habitantes_proto_rawDescData = protoimpl.X.CompressGZIP(file_habitantes_proto_rawDescData)
	})
	return file_habitantes_proto_rawDescData
}

var file_habitantes_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_habitantes_proto_goTypes = []interface{}{
	(*Habitante)(nil),             // 0: habitantes.Habitante
	(*InicializadorRequest)(nil),  // 1: habitantes.InicializadorRequest
	(*InicializadorResponse)(nil), // 2: habitantes.InicializadorResponse
	(*EstadoRequest)(nil),         // 3: habitantes.EstadoRequest
	(*EstadoResponse)(nil),        // 4: habitantes.EstadoResponse
}
var file_habitantes_proto_depIdxs = []int32{
	0, // 0: habitantes.InicializadorResponse.habitantes_inicial:type_name -> habitantes.Habitante
	0, // 1: habitantes.EstadoResponse.estado_habitante:type_name -> habitantes.Habitante
	1, // 2: habitantes.ServicioHabitantes.InicializadorHabitantes:input_type -> habitantes.InicializadorRequest
	3, // 3: habitantes.ServicioHabitantes.ActualizarEstado:input_type -> habitantes.EstadoRequest
	3, // 4: habitantes.ServicioHabitantes.ConsumirRecurso:input_type -> habitantes.EstadoRequest
	2, // 5: habitantes.ServicioHabitantes.InicializadorHabitantes:output_type -> habitantes.InicializadorResponse
	4, // 6: habitantes.ServicioHabitantes.ActualizarEstado:output_type -> habitantes.EstadoResponse
	4, // 7: habitantes.ServicioHabitantes.ConsumirRecurso:output_type -> habitantes.EstadoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_habitantes_proto_init() }
func file_habitantes_proto_init() {
	if File_habitantes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_habitantes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Habitante); i {
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
		file_habitantes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InicializadorRequest); i {
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
		file_habitantes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InicializadorResponse); i {
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
		file_habitantes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstadoRequest); i {
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
		file_habitantes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstadoResponse); i {
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
			RawDescriptor: file_habitantes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_habitantes_proto_goTypes,
		DependencyIndexes: file_habitantes_proto_depIdxs,
		MessageInfos:      file_habitantes_proto_msgTypes,
	}.Build()
	File_habitantes_proto = out.File
	file_habitantes_proto_rawDesc = nil
	file_habitantes_proto_goTypes = nil
	file_habitantes_proto_depIdxs = nil
}