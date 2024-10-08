// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: metav1/types.proto

package metav1

import (
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

type FieldsV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw []byte `protobuf:"bytes,1,opt,name=raw,proto3,oneof" json:"raw,omitempty"`
}

func (x *FieldsV1) Reset() {
	*x = FieldsV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metav1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldsV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldsV1) ProtoMessage() {}

func (x *FieldsV1) ProtoReflect() protoreflect.Message {
	mi := &file_metav1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldsV1.ProtoReflect.Descriptor instead.
func (*FieldsV1) Descriptor() ([]byte, []int) {
	return file_metav1_types_proto_rawDescGZIP(), []int{0}
}

func (x *FieldsV1) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

type OwnerReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion         *string `protobuf:"bytes,5,opt,name=api_version,json=apiVersion,proto3,oneof" json:"api_version,omitempty"`
	Kind               *string `protobuf:"bytes,1,opt,name=kind,proto3,oneof" json:"kind,omitempty"`
	Name               *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Uid                *string `protobuf:"bytes,4,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	Controller         *bool   `protobuf:"varint,6,opt,name=controller,proto3,oneof" json:"controller,omitempty"`
	BlockOwnerDeletion *bool   `protobuf:"varint,7,opt,name=block_owner_deletion,json=blockOwnerDeletion,proto3,oneof" json:"block_owner_deletion,omitempty"`
}

func (x *OwnerReference) Reset() {
	*x = OwnerReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metav1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerReference) ProtoMessage() {}

func (x *OwnerReference) ProtoReflect() protoreflect.Message {
	mi := &file_metav1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerReference.ProtoReflect.Descriptor instead.
func (*OwnerReference) Descriptor() ([]byte, []int) {
	return file_metav1_types_proto_rawDescGZIP(), []int{1}
}

func (x *OwnerReference) GetApiVersion() string {
	if x != nil && x.ApiVersion != nil {
		return *x.ApiVersion
	}
	return ""
}

func (x *OwnerReference) GetKind() string {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return ""
}

func (x *OwnerReference) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *OwnerReference) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *OwnerReference) GetController() bool {
	if x != nil && x.Controller != nil {
		return *x.Controller
	}
	return false
}

func (x *OwnerReference) GetBlockOwnerDeletion() bool {
	if x != nil && x.BlockOwnerDeletion != nil {
		return *x.BlockOwnerDeletion
	}
	return false
}

type ManagedFieldsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manager     *string                `protobuf:"bytes,1,opt,name=manager,proto3,oneof" json:"manager,omitempty"`
	Operation   *string                `protobuf:"bytes,2,opt,name=operation,proto3,oneof" json:"operation,omitempty"`
	ApiVersion  *string                `protobuf:"bytes,3,opt,name=api_version,json=apiVersion,proto3,oneof" json:"api_version,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3,oneof" json:"time,omitempty"`
	FieldsType  *string                `protobuf:"bytes,6,opt,name=fields_type,json=fieldsType,proto3,oneof" json:"fields_type,omitempty"`
	FieldsV1    *FieldsV1              `protobuf:"bytes,7,opt,name=fields_v1,json=fieldsV1,proto3,oneof" json:"fields_v1,omitempty"`
	Subresource *string                `protobuf:"bytes,8,opt,name=subresource,proto3,oneof" json:"subresource,omitempty"`
}

func (x *ManagedFieldsEntry) Reset() {
	*x = ManagedFieldsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metav1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagedFieldsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedFieldsEntry) ProtoMessage() {}

func (x *ManagedFieldsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_metav1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedFieldsEntry.ProtoReflect.Descriptor instead.
func (*ManagedFieldsEntry) Descriptor() ([]byte, []int) {
	return file_metav1_types_proto_rawDescGZIP(), []int{2}
}

func (x *ManagedFieldsEntry) GetManager() string {
	if x != nil && x.Manager != nil {
		return *x.Manager
	}
	return ""
}

func (x *ManagedFieldsEntry) GetOperation() string {
	if x != nil && x.Operation != nil {
		return *x.Operation
	}
	return ""
}

func (x *ManagedFieldsEntry) GetApiVersion() string {
	if x != nil && x.ApiVersion != nil {
		return *x.ApiVersion
	}
	return ""
}

func (x *ManagedFieldsEntry) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ManagedFieldsEntry) GetFieldsType() string {
	if x != nil && x.FieldsType != nil {
		return *x.FieldsType
	}
	return ""
}

func (x *ManagedFieldsEntry) GetFieldsV1() *FieldsV1 {
	if x != nil {
		return x.FieldsV1
	}
	return nil
}

func (x *ManagedFieldsEntry) GetSubresource() string {
	if x != nil && x.Subresource != nil {
		return *x.Subresource
	}
	return ""
}

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                       *string                `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	GenerateName               *string                `protobuf:"bytes,2,opt,name=generate_name,json=generateName,proto3,oneof" json:"generate_name,omitempty"`
	Namespace                  *string                `protobuf:"bytes,3,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
	SelfLink                   *string                `protobuf:"bytes,4,opt,name=self_link,json=selfLink,proto3,oneof" json:"self_link,omitempty"`
	Uid                        *string                `protobuf:"bytes,5,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	ResourceVersion            *string                `protobuf:"bytes,6,opt,name=resource_version,json=resourceVersion,proto3,oneof" json:"resource_version,omitempty"`
	Generation                 *int64                 `protobuf:"varint,7,opt,name=generation,proto3,oneof" json:"generation,omitempty"`
	CreationTimestamp          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=creation_timestamp,json=creationTimestamp,proto3,oneof" json:"creation_timestamp,omitempty"`
	DeletionTimestamp          *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deletion_timestamp,json=deletionTimestamp,proto3,oneof" json:"deletion_timestamp,omitempty"`
	DeletionGracePeriodSeconds *int64                 `protobuf:"varint,10,opt,name=deletion_grace_period_seconds,json=deletionGracePeriodSeconds,proto3,oneof" json:"deletion_grace_period_seconds,omitempty"`
	Labels                     map[string]string      `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                map[string]string      `protobuf:"bytes,12,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OwnerReferences            []*OwnerReference      `protobuf:"bytes,13,rep,name=owner_references,json=ownerReferences,proto3" json:"owner_references,omitempty"`
	Finalizers                 []string               `protobuf:"bytes,14,rep,name=finalizers,proto3" json:"finalizers,omitempty"`
	ManagedFields              []*ManagedFieldsEntry  `protobuf:"bytes,17,rep,name=managed_fields,json=managedFields,proto3" json:"managed_fields,omitempty"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metav1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_metav1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_metav1_types_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectMeta) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ObjectMeta) GetGenerateName() string {
	if x != nil && x.GenerateName != nil {
		return *x.GenerateName
	}
	return ""
}

func (x *ObjectMeta) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *ObjectMeta) GetSelfLink() string {
	if x != nil && x.SelfLink != nil {
		return *x.SelfLink
	}
	return ""
}

func (x *ObjectMeta) GetUid() string {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return ""
}

func (x *ObjectMeta) GetResourceVersion() string {
	if x != nil && x.ResourceVersion != nil {
		return *x.ResourceVersion
	}
	return ""
}

func (x *ObjectMeta) GetGeneration() int64 {
	if x != nil && x.Generation != nil {
		return *x.Generation
	}
	return 0
}

func (x *ObjectMeta) GetCreationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTimestamp
	}
	return nil
}

func (x *ObjectMeta) GetDeletionTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletionTimestamp
	}
	return nil
}

func (x *ObjectMeta) GetDeletionGracePeriodSeconds() int64 {
	if x != nil && x.DeletionGracePeriodSeconds != nil {
		return *x.DeletionGracePeriodSeconds
	}
	return 0
}

func (x *ObjectMeta) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ObjectMeta) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *ObjectMeta) GetOwnerReferences() []*OwnerReference {
	if x != nil {
		return x.OwnerReferences
	}
	return nil
}

func (x *ObjectMeta) GetFinalizers() []string {
	if x != nil {
		return x.Finalizers
	}
	return nil
}

func (x *ObjectMeta) GetManagedFields() []*ManagedFieldsEntry {
	if x != nil {
		return x.ManagedFields
	}
	return nil
}

type ListMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfLink           *string `protobuf:"bytes,1,opt,name=self_link,json=selfLink,proto3,oneof" json:"self_link,omitempty"`
	ResourceVersion    *string `protobuf:"bytes,2,opt,name=resource_version,json=resourceVersion,proto3,oneof" json:"resource_version,omitempty"`
	Continue           *string `protobuf:"bytes,3,opt,name=continue,proto3,oneof" json:"continue,omitempty"`
	RemainingItemCount *int64  `protobuf:"varint,4,opt,name=remaining_item_count,json=remainingItemCount,proto3,oneof" json:"remaining_item_count,omitempty"`
}

func (x *ListMeta) Reset() {
	*x = ListMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metav1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMeta) ProtoMessage() {}

func (x *ListMeta) ProtoReflect() protoreflect.Message {
	mi := &file_metav1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMeta.ProtoReflect.Descriptor instead.
func (*ListMeta) Descriptor() ([]byte, []int) {
	return file_metav1_types_proto_rawDescGZIP(), []int{4}
}

func (x *ListMeta) GetSelfLink() string {
	if x != nil && x.SelfLink != nil {
		return *x.SelfLink
	}
	return ""
}

func (x *ListMeta) GetResourceVersion() string {
	if x != nil && x.ResourceVersion != nil {
		return *x.ResourceVersion
	}
	return ""
}

func (x *ListMeta) GetContinue() string {
	if x != nil && x.Continue != nil {
		return *x.Continue
	}
	return ""
}

func (x *ListMeta) GetRemainingItemCount() int64 {
	if x != nil && x.RemainingItemCount != nil {
		return *x.RemainingItemCount
	}
	return 0
}

var File_metav1_types_proto protoreflect.FileDescriptor

var file_metav1_types_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x31, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x61, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x61, 0x77, 0x22, 0xad, 0x02, 0x0a, 0x0e, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x35, 0x0a, 0x14, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52,
	0x12, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69, 0x64,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x42,
	0x17, 0x0a, 0x15, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb3, 0x03, 0x0a, 0x12, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x1d, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x52, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x76, 0x31, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x67, 0x6f, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x56, 0x31, 0x48, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x56, 0x31, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x76, 0x31, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x73, 0x75, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xbb,
	0x09, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69,
	0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x06, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x4e, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x07, 0x52, 0x11, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01,
	0x01, 0x12, 0x4e, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x08, 0x52, 0x11, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01,
	0x01, 0x12, 0x46, 0x0a, 0x1d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x09, 0x52, 0x1a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x56, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61,
	0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x65, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x67, 0x6f,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x61, 0x0a, 0x10, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x12, 0x61, 0x0a, 0x0e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x15, 0x0a, 0x13,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xfd, 0x01, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65, 0x6c,
	0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x12, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x75, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xa2, 0x02, 0x0a,
	0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x6b, 0x61, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2f, 0x6b, 0x61, 0x72,
	0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x76,
	0x31, 0xa2, 0x02, 0x07, 0x47, 0x43, 0x41, 0x4b, 0x50, 0x41, 0x4d, 0xaa, 0x02, 0x26, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x2e, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e,
	0x4b, 0x61, 0x72, 0x67, 0x6f, 0x2e, 0x50, 0x6b, 0x67, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x76, 0x31, 0xca, 0x02, 0x26, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5c, 0x43, 0x6f,
	0x6d, 0x5c, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x5c, 0x4b, 0x61, 0x72, 0x67, 0x6f, 0x5c, 0x50,
	0x6b, 0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x76, 0x31, 0xe2, 0x02, 0x32,
	0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5c, 0x43, 0x6f, 0x6d, 0x5c, 0x41, 0x6b, 0x75, 0x69, 0x74,
	0x79, 0x5c, 0x4b, 0x61, 0x72, 0x67, 0x6f, 0x5c, 0x50, 0x6b, 0x67, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x4d, 0x65, 0x74, 0x61, 0x76, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x2c, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x3a, 0x3a, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4b, 0x61, 0x72, 0x67, 0x6f, 0x3a,
	0x3a, 0x50, 0x6b, 0x67, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x61, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metav1_types_proto_rawDescOnce sync.Once
	file_metav1_types_proto_rawDescData = file_metav1_types_proto_rawDesc
)

func file_metav1_types_proto_rawDescGZIP() []byte {
	file_metav1_types_proto_rawDescOnce.Do(func() {
		file_metav1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_metav1_types_proto_rawDescData)
	})
	return file_metav1_types_proto_rawDescData
}

var file_metav1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_metav1_types_proto_goTypes = []interface{}{
	(*FieldsV1)(nil),              // 0: github.com.akuity.kargo.pkg.api.metav1.FieldsV1
	(*OwnerReference)(nil),        // 1: github.com.akuity.kargo.pkg.api.metav1.OwnerReference
	(*ManagedFieldsEntry)(nil),    // 2: github.com.akuity.kargo.pkg.api.metav1.ManagedFieldsEntry
	(*ObjectMeta)(nil),            // 3: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta
	(*ListMeta)(nil),              // 4: github.com.akuity.kargo.pkg.api.metav1.ListMeta
	nil,                           // 5: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.LabelsEntry
	nil,                           // 6: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_metav1_types_proto_depIdxs = []int32{
	7, // 0: github.com.akuity.kargo.pkg.api.metav1.ManagedFieldsEntry.time:type_name -> google.protobuf.Timestamp
	0, // 1: github.com.akuity.kargo.pkg.api.metav1.ManagedFieldsEntry.fields_v1:type_name -> github.com.akuity.kargo.pkg.api.metav1.FieldsV1
	7, // 2: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.creation_timestamp:type_name -> google.protobuf.Timestamp
	7, // 3: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.deletion_timestamp:type_name -> google.protobuf.Timestamp
	5, // 4: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.labels:type_name -> github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.LabelsEntry
	6, // 5: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.annotations:type_name -> github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.AnnotationsEntry
	1, // 6: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.owner_references:type_name -> github.com.akuity.kargo.pkg.api.metav1.OwnerReference
	2, // 7: github.com.akuity.kargo.pkg.api.metav1.ObjectMeta.managed_fields:type_name -> github.com.akuity.kargo.pkg.api.metav1.ManagedFieldsEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_metav1_types_proto_init() }
func file_metav1_types_proto_init() {
	if File_metav1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metav1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldsV1); i {
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
		file_metav1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerReference); i {
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
		file_metav1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagedFieldsEntry); i {
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
		file_metav1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
		file_metav1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMeta); i {
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
	file_metav1_types_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_metav1_types_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_metav1_types_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_metav1_types_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_metav1_types_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metav1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metav1_types_proto_goTypes,
		DependencyIndexes: file_metav1_types_proto_depIdxs,
		MessageInfos:      file_metav1_types_proto_msgTypes,
	}.Build()
	File_metav1_types_proto = out.File
	file_metav1_types_proto_rawDesc = nil
	file_metav1_types_proto_goTypes = nil
	file_metav1_types_proto_depIdxs = nil
}
