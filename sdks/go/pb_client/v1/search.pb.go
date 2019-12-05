// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/search.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Search spec definition
type SearchSpec struct {
	// Search query
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Search sort
	Sort string `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	// Search group bys
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Search group bys
	Groupby string `protobuf:"bytes,4,opt,name=groupby,proto3" json:"groupby,omitempty"`
	// Search columns
	Columns              string   `protobuf:"bytes,5,opt,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchSpec) Reset()         { *m = SearchSpec{} }
func (m *SearchSpec) String() string { return proto.CompactTextString(m) }
func (*SearchSpec) ProtoMessage()    {}
func (*SearchSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445a7a836953543, []int{0}
}

func (m *SearchSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchSpec.Unmarshal(m, b)
}
func (m *SearchSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchSpec.Marshal(b, m, deterministic)
}
func (m *SearchSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchSpec.Merge(m, src)
}
func (m *SearchSpec) XXX_Size() int {
	return xxx_messageInfo_SearchSpec.Size(m)
}
func (m *SearchSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SearchSpec proto.InternalMessageInfo

func (m *SearchSpec) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchSpec) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchSpec) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchSpec) GetGroupby() string {
	if m != nil {
		return m.Groupby
	}
	return ""
}

func (m *SearchSpec) GetColumns() string {
	if m != nil {
		return m.Columns
	}
	return ""
}

// Search specification
type Search struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Optional name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Search spec
	Spec *SearchSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Search) Reset()         { *m = Search{} }
func (m *Search) String() string { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()    {}
func (*Search) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445a7a836953543, []int{1}
}

func (m *Search) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Search.Unmarshal(m, b)
}
func (m *Search) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Search.Marshal(b, m, deterministic)
}
func (m *Search) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Search.Merge(m, src)
}
func (m *Search) XXX_Size() int {
	return xxx_messageInfo_Search.Size(m)
}
func (m *Search) XXX_DiscardUnknown() {
	xxx_messageInfo_Search.DiscardUnknown(m)
}

var xxx_messageInfo_Search proto.InternalMessageInfo

func (m *Search) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Search) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Search) GetSpec() *SearchSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Search) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Search) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// Request data to create/update artifacts store
type SearchBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project under namesapce
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Search body
	Search               *Search  `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchBodyRequest) Reset()         { *m = SearchBodyRequest{} }
func (m *SearchBodyRequest) String() string { return proto.CompactTextString(m) }
func (*SearchBodyRequest) ProtoMessage()    {}
func (*SearchBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445a7a836953543, []int{2}
}

func (m *SearchBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchBodyRequest.Unmarshal(m, b)
}
func (m *SearchBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchBodyRequest.Marshal(b, m, deterministic)
}
func (m *SearchBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchBodyRequest.Merge(m, src)
}
func (m *SearchBodyRequest) XXX_Size() int {
	return xxx_messageInfo_SearchBodyRequest.Size(m)
}
func (m *SearchBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchBodyRequest proto.InternalMessageInfo

func (m *SearchBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *SearchBodyRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *SearchBodyRequest) GetSearch() *Search {
	if m != nil {
		return m.Search
	}
	return nil
}

// Contains list searches
type ListSearchesResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*Search `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSearchesResponse) Reset()         { *m = ListSearchesResponse{} }
func (m *ListSearchesResponse) String() string { return proto.CompactTextString(m) }
func (*ListSearchesResponse) ProtoMessage()    {}
func (*ListSearchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5445a7a836953543, []int{3}
}

func (m *ListSearchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSearchesResponse.Unmarshal(m, b)
}
func (m *ListSearchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSearchesResponse.Marshal(b, m, deterministic)
}
func (m *ListSearchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSearchesResponse.Merge(m, src)
}
func (m *ListSearchesResponse) XXX_Size() int {
	return xxx_messageInfo_ListSearchesResponse.Size(m)
}
func (m *ListSearchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSearchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSearchesResponse proto.InternalMessageInfo

func (m *ListSearchesResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListSearchesResponse) GetResults() []*Search {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListSearchesResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListSearchesResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchSpec)(nil), "v1.SearchSpec")
	proto.RegisterType((*Search)(nil), "v1.Search")
	proto.RegisterType((*SearchBodyRequest)(nil), "v1.SearchBodyRequest")
	proto.RegisterType((*ListSearchesResponse)(nil), "v1.ListSearchesResponse")
}

func init() { proto.RegisterFile("v1/search.proto", fileDescriptor_5445a7a836953543) }

var fileDescriptor_5445a7a836953543 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x3d, 0x6f, 0xe2, 0x40,
	0x10, 0x95, 0x01, 0xc3, 0x31, 0x48, 0x77, 0xba, 0x15, 0x85, 0x45, 0x73, 0xc8, 0xba, 0x82, 0xca,
	0x08, 0xae, 0xba, 0x92, 0xd4, 0xa9, 0x96, 0xf4, 0x91, 0xb1, 0x27, 0x8e, 0x23, 0xdb, 0xbb, 0xec,
	0x87, 0x13, 0xba, 0x28, 0x3f, 0x2f, 0xbf, 0x2a, 0xda, 0x59, 0x6f, 0x90, 0xd2, 0xa4, 0x9b, 0x37,
	0xef, 0xcd, 0xbe, 0x9d, 0x37, 0xf0, 0xab, 0xdf, 0x6d, 0x35, 0xe6, 0xaa, 0x78, 0xcc, 0xa4, 0x12,
	0x46, 0xb0, 0x51, 0xbf, 0x5b, 0xfd, 0xa9, 0x84, 0xa8, 0x1a, 0xdc, 0x52, 0xe7, 0x64, 0x1f, 0xb6,
	0xa6, 0x6e, 0x51, 0x9b, 0xbc, 0x95, 0x5e, 0x94, 0xbe, 0x46, 0x00, 0x47, 0x9a, 0x3a, 0x4a, 0x2c,
	0xd8, 0x12, 0xe2, 0xb3, 0x45, 0x75, 0x49, 0xa2, 0x75, 0xb4, 0x99, 0x73, 0x0f, 0x18, 0x83, 0x89,
	0x16, 0xca, 0x24, 0x23, 0x6a, 0x52, 0xed, 0x94, 0x4d, 0xdd, 0xd6, 0x26, 0x19, 0xaf, 0xa3, 0x4d,
	0xcc, 0x3d, 0x60, 0x09, 0xcc, 0x2a, 0x25, 0xac, 0x3c, 0x5d, 0x92, 0x09, 0x89, 0x03, 0x74, 0x4c,
	0x21, 0x1a, 0xdb, 0x76, 0x3a, 0x89, 0x3d, 0x33, 0xc0, 0xf4, 0x3d, 0x82, 0xa9, 0xff, 0x82, 0x33,
	0xb2, 0xb6, 0x2e, 0x07, 0x77, 0xaa, 0x5d, 0xaf, 0xcb, 0x5b, 0x0c, 0xe6, 0xae, 0x66, 0x29, 0x4c,
	0xb4, 0xc4, 0x82, 0xbc, 0x17, 0xfb, 0x9f, 0x59, 0xbf, 0xcb, 0xae, 0x4b, 0x70, 0xe2, 0xd8, 0x7f,
	0x80, 0x42, 0x61, 0x6e, 0xb0, 0xbc, 0xcf, 0x0d, 0xfd, 0x66, 0xb1, 0x5f, 0x65, 0x3e, 0x8f, 0x2c,
	0xe4, 0x91, 0xdd, 0x85, 0x3c, 0xf8, 0x7c, 0x50, 0x1f, 0x8c, 0x1b, 0xb5, 0xb2, 0x0c, 0xa3, 0xf1,
	0xf7, 0xa3, 0x83, 0xfa, 0x60, 0xd2, 0x0a, 0x7e, 0xfb, 0x9f, 0xdc, 0x88, 0xf2, 0xc2, 0xf1, 0x6c,
	0x51, 0x53, 0x56, 0xe2, 0xb9, 0x43, 0x15, 0x52, 0x25, 0xe0, 0x12, 0x91, 0x4a, 0x3c, 0x61, 0x11,
	0x82, 0x0d, 0x90, 0xa5, 0x30, 0xf5, 0x97, 0x1c, 0x16, 0x84, 0xeb, 0x82, 0x7c, 0x60, 0xd2, 0xb7,
	0x08, 0x96, 0xb7, 0xb5, 0x36, 0xbe, 0x8d, 0x9a, 0xa3, 0x96, 0xa2, 0xd3, 0xe8, 0xcc, 0x0a, 0x61,
	0x3b, 0x43, 0x66, 0x31, 0xf7, 0x80, 0xfd, 0x85, 0x99, 0x42, 0x6d, 0x1b, 0xa3, 0x93, 0xd1, 0x7a,
	0xfc, 0xe5, 0xcd, 0x40, 0xb1, 0x15, 0xfc, 0x90, 0x0a, 0xfb, 0x5a, 0x58, 0x4d, 0xd6, 0x73, 0xfe,
	0x89, 0xe9, 0x0e, 0xf8, 0x62, 0x86, 0xbb, 0x52, 0x7d, 0x9a, 0x52, 0x18, 0xff, 0x3e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xab, 0x9b, 0x1d, 0x6b, 0x7c, 0x02, 0x00, 0x00,
}