// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.12
// source: payload.proto

package data

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

type HorizontalAlign int32

const (
	HorizontalAlign_CENTER HorizontalAlign = 0
	HorizontalAlign_LEFT   HorizontalAlign = 1
	HorizontalAlign_RIGHT  HorizontalAlign = 2
)

// Enum value maps for HorizontalAlign.
var (
	HorizontalAlign_name = map[int32]string{
		0: "CENTER",
		1: "LEFT",
		2: "RIGHT",
	}
	HorizontalAlign_value = map[string]int32{
		"CENTER": 0,
		"LEFT":   1,
		"RIGHT":  2,
	}
)

func (x HorizontalAlign) Enum() *HorizontalAlign {
	p := new(HorizontalAlign)
	*p = x
	return p
}

func (x HorizontalAlign) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HorizontalAlign) Descriptor() protoreflect.EnumDescriptor {
	return file_payload_proto_enumTypes[0].Descriptor()
}

func (HorizontalAlign) Type() protoreflect.EnumType {
	return &file_payload_proto_enumTypes[0]
}

func (x HorizontalAlign) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HorizontalAlign.Descriptor instead.
func (HorizontalAlign) EnumDescriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{0}
}

type VerticalAlign int32

const (
	VerticalAlign_MIDDLE VerticalAlign = 0
	VerticalAlign_TOP    VerticalAlign = 1
	VerticalAlign_BOTTOM VerticalAlign = 2
)

// Enum value maps for VerticalAlign.
var (
	VerticalAlign_name = map[int32]string{
		0: "MIDDLE",
		1: "TOP",
		2: "BOTTOM",
	}
	VerticalAlign_value = map[string]int32{
		"MIDDLE": 0,
		"TOP":    1,
		"BOTTOM": 2,
	}
)

func (x VerticalAlign) Enum() *VerticalAlign {
	p := new(VerticalAlign)
	*p = x
	return p
}

func (x VerticalAlign) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerticalAlign) Descriptor() protoreflect.EnumDescriptor {
	return file_payload_proto_enumTypes[1].Descriptor()
}

func (VerticalAlign) Type() protoreflect.EnumType {
	return &file_payload_proto_enumTypes[1]
}

func (x VerticalAlign) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerticalAlign.Descriptor instead.
func (VerticalAlign) EnumDescriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{1}
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Template uint64  `protobuf:"varint,2,opt,name=template,proto3" json:"template,omitempty"`
	Text     []*Text `protobuf:"bytes,3,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Payload) GetTemplate() uint64 {
	if x != nil {
		return x.Template
	}
	return 0
}

func (x *Payload) GetText() []*Text {
	if x != nil {
		return x.Text
	}
	return nil
}

type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateText    *uint32          `protobuf:"varint,1,opt,name=template_text,json=templateText,proto3,oneof" json:"template_text,omitempty"`
	Text            *string          `protobuf:"bytes,2,opt,name=text,proto3,oneof" json:"text,omitempty"`
	X               *uint32          `protobuf:"varint,3,opt,name=x,proto3,oneof" json:"x,omitempty"`
	Y               *uint32          `protobuf:"varint,4,opt,name=y,proto3,oneof" json:"y,omitempty"`
	Width           *uint32          `protobuf:"varint,5,opt,name=width,proto3,oneof" json:"width,omitempty"`
	Height          *uint32          `protobuf:"varint,6,opt,name=height,proto3,oneof" json:"height,omitempty"`
	Font            *string          `protobuf:"bytes,7,opt,name=font,proto3,oneof" json:"font,omitempty"`
	Size            *uint32          `protobuf:"varint,8,opt,name=size,proto3,oneof" json:"size,omitempty"`
	Unfilled        *bool            `protobuf:"varint,9,opt,name=unfilled,proto3,oneof" json:"unfilled,omitempty"`
	FillColor       *string          `protobuf:"bytes,10,opt,name=fill_color,json=fillColor,proto3,oneof" json:"fill_color,omitempty"`
	StrokeColor     *string          `protobuf:"bytes,11,opt,name=stroke_color,json=strokeColor,proto3,oneof" json:"stroke_color,omitempty"`
	Stroke          *uint32          `protobuf:"varint,12,opt,name=stroke,proto3,oneof" json:"stroke,omitempty"`
	HorizontalAlign *HorizontalAlign `protobuf:"varint,13,opt,name=horizontal_align,json=horizontalAlign,proto3,enum=HorizontalAlign,oneof" json:"horizontal_align,omitempty"`
	VerticalAlign   *VerticalAlign   `protobuf:"varint,14,opt,name=vertical_align,json=verticalAlign,proto3,enum=VerticalAlign,oneof" json:"vertical_align,omitempty"`
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{1}
}

func (x *Text) GetTemplateText() uint32 {
	if x != nil && x.TemplateText != nil {
		return *x.TemplateText
	}
	return 0
}

func (x *Text) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *Text) GetX() uint32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *Text) GetY() uint32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *Text) GetWidth() uint32 {
	if x != nil && x.Width != nil {
		return *x.Width
	}
	return 0
}

func (x *Text) GetHeight() uint32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

func (x *Text) GetFont() string {
	if x != nil && x.Font != nil {
		return *x.Font
	}
	return ""
}

func (x *Text) GetSize() uint32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *Text) GetUnfilled() bool {
	if x != nil && x.Unfilled != nil {
		return *x.Unfilled
	}
	return false
}

func (x *Text) GetFillColor() string {
	if x != nil && x.FillColor != nil {
		return *x.FillColor
	}
	return ""
}

func (x *Text) GetStrokeColor() string {
	if x != nil && x.StrokeColor != nil {
		return *x.StrokeColor
	}
	return ""
}

func (x *Text) GetStroke() uint32 {
	if x != nil && x.Stroke != nil {
		return *x.Stroke
	}
	return 0
}

func (x *Text) GetHorizontalAlign() HorizontalAlign {
	if x != nil && x.HorizontalAlign != nil {
		return *x.HorizontalAlign
	}
	return HorizontalAlign_CENTER
}

func (x *Text) GetVerticalAlign() VerticalAlign {
	if x != nil && x.VerticalAlign != nil {
		return *x.VerticalAlign
	}
	return VerticalAlign_MIDDLE
}

type TemplateText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text            *string         `protobuf:"bytes,1,opt,name=text,proto3,oneof" json:"text,omitempty"`
	X               uint32          `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y               uint32          `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Width           uint32          `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Height          uint32          `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Font            string          `protobuf:"bytes,6,opt,name=font,proto3" json:"font,omitempty"`
	Size            uint32          `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Unfilled        bool            `protobuf:"varint,8,opt,name=unfilled,proto3" json:"unfilled,omitempty"`
	FillColor       string          `protobuf:"bytes,9,opt,name=fill_color,json=fillColor,proto3" json:"fill_color,omitempty"`
	StrokeColor     string          `protobuf:"bytes,10,opt,name=stroke_color,json=strokeColor,proto3" json:"stroke_color,omitempty"`
	Stroke          uint32          `protobuf:"varint,11,opt,name=stroke,proto3" json:"stroke,omitempty"`
	HorizontalAlign HorizontalAlign `protobuf:"varint,12,opt,name=horizontal_align,json=horizontalAlign,proto3,enum=HorizontalAlign" json:"horizontal_align,omitempty"`
	VerticalAlign   VerticalAlign   `protobuf:"varint,13,opt,name=vertical_align,json=verticalAlign,proto3,enum=VerticalAlign" json:"vertical_align,omitempty"`
}

func (x *TemplateText) Reset() {
	*x = TemplateText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateText) ProtoMessage() {}

func (x *TemplateText) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateText.ProtoReflect.Descriptor instead.
func (*TemplateText) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateText) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *TemplateText) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *TemplateText) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *TemplateText) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *TemplateText) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TemplateText) GetFont() string {
	if x != nil {
		return x.Font
	}
	return ""
}

func (x *TemplateText) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TemplateText) GetUnfilled() bool {
	if x != nil {
		return x.Unfilled
	}
	return false
}

func (x *TemplateText) GetFillColor() string {
	if x != nil {
		return x.FillColor
	}
	return ""
}

func (x *TemplateText) GetStrokeColor() string {
	if x != nil {
		return x.StrokeColor
	}
	return ""
}

func (x *TemplateText) GetStroke() uint32 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

func (x *TemplateText) GetHorizontalAlign() HorizontalAlign {
	if x != nil {
		return x.HorizontalAlign
	}
	return HorizontalAlign_CENTER
}

func (x *TemplateText) GetVerticalAlign() VerticalAlign {
	if x != nil {
		return x.VerticalAlign
	}
	return VerticalAlign_MIDDLE
}

var File_payload_proto protoreflect.FileDescriptor

var file_payload_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x8f, 0x05, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0c, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x02, 0x52, 0x01, 0x78, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x01, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x66, 0x6f, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x04, 0x66, 0x6f, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x6e, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x6c,
	0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0a, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x0b, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x40, 0x0a, 0x10, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x6c, 0x69, 0x67, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x48, 0x6f,
	0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x48, 0x0c, 0x52,
	0x0f, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x61, 0x6c, 0x69, 0x67, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x56, 0x65,
	0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x48, 0x0d, 0x52, 0x0d, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x78,
	0x42, 0x04, 0x0a, 0x02, 0x5f, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x66, 0x6f, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x75, 0x6e, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x74,
	0x72, 0x6f, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x72, 0x6f, 0x6b, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f,
	0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x22, 0x8c, 0x03,
	0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x01, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6f, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x6f, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x6c, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x6f,
	0x6b, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x12,
	0x3b, 0x0a, 0x10, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6c,
	0x69, 0x67, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x48, 0x6f, 0x72, 0x69,
	0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x52, 0x0f, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x12, 0x35, 0x0a, 0x0e,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41,
	0x6c, 0x69, 0x67, 0x6e, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x6c,
	0x69, 0x67, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x2a, 0x32, 0x0a, 0x0f,
	0x48, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x45, 0x46, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02,
	0x2a, 0x30, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x67,
	0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x44, 0x44, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x54, 0x54, 0x4f, 0x4d,
	0x10, 0x02, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x56, 0x69, 0x6c, 0x73, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x6d, 0x65, 0x6b, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payload_proto_rawDescOnce sync.Once
	file_payload_proto_rawDescData = file_payload_proto_rawDesc
)

func file_payload_proto_rawDescGZIP() []byte {
	file_payload_proto_rawDescOnce.Do(func() {
		file_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_payload_proto_rawDescData)
	})
	return file_payload_proto_rawDescData
}

var file_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payload_proto_goTypes = []interface{}{
	(HorizontalAlign)(0), // 0: HorizontalAlign
	(VerticalAlign)(0),   // 1: VerticalAlign
	(*Payload)(nil),      // 2: Payload
	(*Text)(nil),         // 3: Text
	(*TemplateText)(nil), // 4: TemplateText
}
var file_payload_proto_depIdxs = []int32{
	3, // 0: Payload.text:type_name -> Text
	0, // 1: Text.horizontal_align:type_name -> HorizontalAlign
	1, // 2: Text.vertical_align:type_name -> VerticalAlign
	0, // 3: TemplateText.horizontal_align:type_name -> HorizontalAlign
	1, // 4: TemplateText.vertical_align:type_name -> VerticalAlign
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_payload_proto_init() }
func file_payload_proto_init() {
	if File_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
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
		file_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateText); i {
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
	file_payload_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_payload_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payload_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payload_proto_goTypes,
		DependencyIndexes: file_payload_proto_depIdxs,
		EnumInfos:         file_payload_proto_enumTypes,
		MessageInfos:      file_payload_proto_msgTypes,
	}.Build()
	File_payload_proto = out.File
	file_payload_proto_rawDesc = nil
	file_payload_proto_goTypes = nil
	file_payload_proto_depIdxs = nil
}
