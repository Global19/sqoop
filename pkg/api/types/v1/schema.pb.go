// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import gloo_api_v1 "github.com/solo-io/gloo/pkg/api/types/v1"
import gloo_api_v11 "github.com/solo-io/gloo/pkg/api/types/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// The Schema object wraps the user's GraphQL Schema, which is stored as an inline string.
// The Schema Object contains a Status field which is used by Sqoop to validate the user's input schema.
type Schema struct {
	// Schema Names must be unique and follow the following syntax rules:
	// One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// name of the resolver map to use to resolve this schema.
	// if the user leaves this empty, Sqoop will generate the skeleton of a resolver map for the user
	ResolverMap string `protobuf:"bytes,2,opt,name=resolver_map,json=resolverMap,proto3" json:"resolver_map,omitempty"`
	// inline the entire graphql schema as a string here
	InlineSchema string `protobuf:"bytes,3,opt,name=inline_schema,json=inlineSchema,proto3" json:"inline_schema,omitempty"`
	// Status indicates the validation status of the role resource.
	// Status is read-only by clients, and set by gloo during validation
	Status *gloo_api_v1.Status `protobuf:"bytes,6,opt,name=status" json:"status,omitempty" testdiff:"ignore"`
	// Metadata contains the resource metadata for the role
	Metadata *gloo_api_v11.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Schema) Reset()                    { *m = Schema{} }
func (m *Schema) String() string            { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()               {}
func (*Schema) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{0} }

func (m *Schema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Schema) GetResolverMap() string {
	if m != nil {
		return m.ResolverMap
	}
	return ""
}

func (m *Schema) GetInlineSchema() string {
	if m != nil {
		return m.InlineSchema
	}
	return ""
}

func (m *Schema) GetStatus() *gloo_api_v1.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Schema) GetMetadata() *gloo_api_v11.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Schema)(nil), "sqoop.api.v1.Schema")
}
func (this *Schema) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Schema)
	if !ok {
		that2, ok := that.(Schema)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.ResolverMap != that1.ResolverMap {
		return false
	}
	if this.InlineSchema != that1.InlineSchema {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("schema.proto", fileDescriptorSchema) }

var fileDescriptorSchema = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x15, 0x40, 0x01, 0x9c, 0x30, 0x60, 0xa8, 0x14, 0x75, 0x80, 0x12, 0x96, 0x48, 0xa8,
	0xb6, 0x02, 0x1b, 0x63, 0xf6, 0x2e, 0xe9, 0xc6, 0x52, 0xb9, 0xad, 0xeb, 0x5a, 0xd8, 0x79, 0x26,
	0x76, 0x23, 0xf1, 0x47, 0xfc, 0x13, 0x12, 0x03, 0x9f, 0xc0, 0x17, 0xa0, 0x3a, 0x2e, 0x08, 0x09,
	0xa9, 0xdb, 0xd3, 0xbb, 0xc7, 0xd7, 0xf7, 0x5d, 0x94, 0xda, 0xc5, 0x9a, 0x6b, 0x46, 0x4c, 0x0b,
	0x0e, 0x70, 0xf2, 0xa2, 0x00, 0x08, 0x33, 0x92, 0x74, 0xe5, 0xf0, 0x52, 0x80, 0x00, 0xbf, 0xa7,
	0xdb, 0xa9, 0x47, 0x86, 0x77, 0x42, 0xba, 0xf5, 0x66, 0x4e, 0x16, 0xa0, 0xa9, 0x05, 0x05, 0x63,
	0x09, 0x54, 0x28, 0x00, 0xca, 0x8c, 0xa4, 0x5d, 0x49, 0xad, 0x63, 0x6e, 0x63, 0x03, 0x3c, 0xde,
	0x03, 0x6b, 0xee, 0xd8, 0x92, 0xb9, 0xf0, 0x7d, 0xfe, 0x1e, 0xa1, 0x78, 0xea, 0xf3, 0x60, 0x8c,
	0x8e, 0x1a, 0xa6, 0x79, 0x16, 0x8d, 0xa2, 0xe2, 0xb4, 0xf6, 0x33, 0xbe, 0x41, 0x69, 0xcb, 0x2d,
	0xa8, 0x8e, 0xb7, 0x33, 0xcd, 0x4c, 0x76, 0xe0, 0xb5, 0x64, 0xb7, 0x9b, 0x30, 0x83, 0x6f, 0xd1,
	0x99, 0x6c, 0x94, 0x6c, 0xf8, 0xac, 0xbf, 0x2b, 0x3b, 0xf4, 0x4c, 0xda, 0x2f, 0x83, 0x77, 0x85,
	0xe2, 0x3e, 0x65, 0x16, 0x8f, 0xa2, 0x22, 0xb9, 0xbf, 0x20, 0xe2, 0xf7, 0x6c, 0x32, 0xf5, 0x52,
	0x35, 0xf8, 0xfa, 0xb8, 0x3e, 0x77, 0xdc, 0xba, 0xa5, 0x5c, 0xad, 0x1e, 0x73, 0x29, 0x1a, 0x68,
	0x79, 0x5e, 0x87, 0x97, 0xb8, 0x44, 0x27, 0xbb, 0xf0, 0xd9, 0xb1, 0x77, 0x19, 0xfc, 0x71, 0x99,
	0x04, 0xb1, 0xfe, 0xc1, 0x2a, 0xf2, 0xf6, 0x79, 0x15, 0x3d, 0x15, 0xff, 0x54, 0xb2, 0x6d, 0x9d,
	0x9a, 0x67, 0xe1, 0x6b, 0x71, 0xaf, 0x86, 0x5b, 0xda, 0x95, 0xf3, 0xd8, 0x97, 0xf2, 0xf0, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xae, 0x89, 0xbe, 0xa3, 0x01, 0x00, 0x00,
}
