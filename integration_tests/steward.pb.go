// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: steward.proto

package integration_tests

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

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellarId string `protobuf:"bytes,1,opt,name=cellar_id,json=cellarId,proto3" json:"cellar_id,omitempty"`
	// Types that are assignable to ContractCallData:
	//	*SubmitRequest_Uniswapv3Rebalance
	ContractCallData isSubmitRequest_ContractCallData `protobuf_oneof:"contract_call_data"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_steward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_steward_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitRequest) GetCellarId() string {
	if x != nil {
		return x.CellarId
	}
	return ""
}

func (m *SubmitRequest) GetContractCallData() isSubmitRequest_ContractCallData {
	if m != nil {
		return m.ContractCallData
	}
	return nil
}

func (x *SubmitRequest) GetUniswapv3Rebalance() *UniswapV3RebalanceParams {
	if x, ok := x.GetContractCallData().(*SubmitRequest_Uniswapv3Rebalance); ok {
		return x.Uniswapv3Rebalance
	}
	return nil
}

type isSubmitRequest_ContractCallData interface {
	isSubmitRequest_ContractCallData()
}

type SubmitRequest_Uniswapv3Rebalance struct {
	Uniswapv3Rebalance *UniswapV3RebalanceParams `protobuf:"bytes,2,opt,name=uniswapv3_rebalance,json=uniswapv3Rebalance,proto3,oneof"`
}

func (*SubmitRequest_Uniswapv3Rebalance) isSubmitRequest_ContractCallData() {}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_steward_proto_rawDescGZIP(), []int{1}
}

type UniswapV3Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpperPrice int32 `protobuf:"varint,2,opt,name=upper_price,json=upperPrice,proto3" json:"upper_price,omitempty"`
	LowerPrice int32 `protobuf:"varint,3,opt,name=lower_price,json=lowerPrice,proto3" json:"lower_price,omitempty"`
	Weight     int32 `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *UniswapV3Position) Reset() {
	*x = UniswapV3Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3Position) ProtoMessage() {}

func (x *UniswapV3Position) ProtoReflect() protoreflect.Message {
	mi := &file_steward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3Position.ProtoReflect.Descriptor instead.
func (*UniswapV3Position) Descriptor() ([]byte, []int) {
	return file_steward_proto_rawDescGZIP(), []int{2}
}

func (x *UniswapV3Position) GetUpperPrice() int32 {
	if x != nil {
		return x.UpperPrice
	}
	return 0
}

func (x *UniswapV3Position) GetLowerPrice() int32 {
	if x != nil {
		return x.LowerPrice
	}
	return 0
}

func (x *UniswapV3Position) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type UniswapV3RebalanceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CellarTickInfo []*UniswapV3Position `protobuf:"bytes,1,rep,name=cellar_tick_info,json=cellarTickInfo,proto3" json:"cellar_tick_info,omitempty"`
	CurrentPrice   uint64               `protobuf:"varint,2,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
}

func (x *UniswapV3RebalanceParams) Reset() {
	*x = UniswapV3RebalanceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniswapV3RebalanceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniswapV3RebalanceParams) ProtoMessage() {}

func (x *UniswapV3RebalanceParams) ProtoReflect() protoreflect.Message {
	mi := &file_steward_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniswapV3RebalanceParams.ProtoReflect.Descriptor instead.
func (*UniswapV3RebalanceParams) Descriptor() ([]byte, []int) {
	return file_steward_proto_rawDescGZIP(), []int{3}
}

func (x *UniswapV3RebalanceParams) GetCellarTickInfo() []*UniswapV3Position {
	if x != nil {
		return x.CellarTickInfo
	}
	return nil
}

func (x *UniswapV3RebalanceParams) GetCurrentPrice() uint64 {
	if x != nil {
		return x.CurrentPrice
	}
	return 0
}

var File_steward_proto protoreflect.FileDescriptor

var file_steward_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x9b, 0x01, 0x0a, 0x0d,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x49, 0x64, 0x12, 0x57, 0x0a, 0x13, 0x75, 0x6e,
	0x69, 0x73, 0x77, 0x61, 0x70, 0x76, 0x33, 0x5f, 0x72, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x52, 0x65,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52,
	0x12, 0x75, 0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x76, 0x33, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x0a, 0x11, 0x55,
	0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x55,
	0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x52, 0x65, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x47, 0x0a, 0x10, 0x63, 0x65, 0x6c, 0x6c, 0x61,
	0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x6e, 0x69, 0x73, 0x77, 0x61, 0x70, 0x56, 0x33, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0x51, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12,
	0x19, 0x2e, 0x73, 0x74, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x73, 0x74, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_steward_proto_rawDescOnce sync.Once
	file_steward_proto_rawDescData = file_steward_proto_rawDesc
)

func file_steward_proto_rawDescGZIP() []byte {
	file_steward_proto_rawDescOnce.Do(func() {
		file_steward_proto_rawDescData = protoimpl.X.CompressGZIP(file_steward_proto_rawDescData)
	})
	return file_steward_proto_rawDescData
}

var file_steward_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_steward_proto_goTypes = []interface{}{
	(*SubmitRequest)(nil),            // 0: steward.v1.SubmitRequest
	(*SubmitResponse)(nil),           // 1: steward.v1.SubmitResponse
	(*UniswapV3Position)(nil),        // 2: steward.v1.UniswapV3Position
	(*UniswapV3RebalanceParams)(nil), // 3: steward.v1.UniswapV3RebalanceParams
}
var file_steward_proto_depIdxs = []int32{
	3, // 0: steward.v1.SubmitRequest.uniswapv3_rebalance:type_name -> steward.v1.UniswapV3RebalanceParams
	2, // 1: steward.v1.UniswapV3RebalanceParams.cellar_tick_info:type_name -> steward.v1.UniswapV3Position
	0, // 2: steward.v1.ContractCall.Submit:input_type -> steward.v1.SubmitRequest
	1, // 3: steward.v1.ContractCall.Submit:output_type -> steward.v1.SubmitResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_steward_proto_init() }
func file_steward_proto_init() {
	if File_steward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_steward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_steward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
		file_steward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3Position); i {
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
		file_steward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniswapV3RebalanceParams); i {
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
	file_steward_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SubmitRequest_Uniswapv3Rebalance)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_steward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steward_proto_goTypes,
		DependencyIndexes: file_steward_proto_depIdxs,
		MessageInfos:      file_steward_proto_msgTypes,
	}.Build()
	File_steward_proto = out.File
	file_steward_proto_rawDesc = nil
	file_steward_proto_goTypes = nil
	file_steward_proto_depIdxs = nil
}