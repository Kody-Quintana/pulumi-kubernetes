// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachmentType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpec `pulumi:"spec"`
	// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
	Status *VolumeAttachmentStatus `pulumi:"status"`
}

// VolumeAttachmentTypeInput is an input type that accepts VolumeAttachmentTypeArgs and VolumeAttachmentTypeOutput values.
// You can construct a concrete instance of `VolumeAttachmentTypeInput` via:
//
//          VolumeAttachmentTypeArgs{...}
type VolumeAttachmentTypeInput interface {
	pulumi.Input

	ToVolumeAttachmentTypeOutput() VolumeAttachmentTypeOutput
	ToVolumeAttachmentTypeOutputWithContext(context.Context) VolumeAttachmentTypeOutput
}

// VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachmentTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput `pulumi:"metadata"`
	// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
	Spec VolumeAttachmentSpecInput `pulumi:"spec"`
	// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
	Status VolumeAttachmentStatusPtrInput `pulumi:"status"`
}

func (VolumeAttachmentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentType)(nil)).Elem()
}

func (i VolumeAttachmentTypeArgs) ToVolumeAttachmentTypeOutput() VolumeAttachmentTypeOutput {
	return i.ToVolumeAttachmentTypeOutputWithContext(context.Background())
}

func (i VolumeAttachmentTypeArgs) ToVolumeAttachmentTypeOutputWithContext(ctx context.Context) VolumeAttachmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentTypeOutput)
}

// VolumeAttachmentTypeArrayInput is an input type that accepts VolumeAttachmentTypeArray and VolumeAttachmentTypeArrayOutput values.
// You can construct a concrete instance of `VolumeAttachmentTypeArrayInput` via:
//
//          VolumeAttachmentTypeArray{ VolumeAttachmentTypeArgs{...} }
type VolumeAttachmentTypeArrayInput interface {
	pulumi.Input

	ToVolumeAttachmentTypeArrayOutput() VolumeAttachmentTypeArrayOutput
	ToVolumeAttachmentTypeArrayOutputWithContext(context.Context) VolumeAttachmentTypeArrayOutput
}

type VolumeAttachmentTypeArray []VolumeAttachmentTypeInput

func (VolumeAttachmentTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeAttachmentType)(nil)).Elem()
}

func (i VolumeAttachmentTypeArray) ToVolumeAttachmentTypeArrayOutput() VolumeAttachmentTypeArrayOutput {
	return i.ToVolumeAttachmentTypeArrayOutputWithContext(context.Background())
}

func (i VolumeAttachmentTypeArray) ToVolumeAttachmentTypeArrayOutputWithContext(ctx context.Context) VolumeAttachmentTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentTypeArrayOutput)
}

// VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.
//
// VolumeAttachment objects are non-namespaced.
type VolumeAttachmentTypeOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentType)(nil)).Elem()
}

func (o VolumeAttachmentTypeOutput) ToVolumeAttachmentTypeOutput() VolumeAttachmentTypeOutput {
	return o
}

func (o VolumeAttachmentTypeOutput) ToVolumeAttachmentTypeOutputWithContext(ctx context.Context) VolumeAttachmentTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VolumeAttachmentTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VolumeAttachmentTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VolumeAttachmentTypeOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentType) *metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.
func (o VolumeAttachmentTypeOutput) Spec() VolumeAttachmentSpecOutput {
	return o.ApplyT(func(v VolumeAttachmentType) VolumeAttachmentSpec { return v.Spec }).(VolumeAttachmentSpecOutput)
}

// Status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.
func (o VolumeAttachmentTypeOutput) Status() VolumeAttachmentStatusPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentType) *VolumeAttachmentStatus { return v.Status }).(VolumeAttachmentStatusPtrOutput)
}

type VolumeAttachmentTypeArrayOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeAttachmentType)(nil)).Elem()
}

func (o VolumeAttachmentTypeArrayOutput) ToVolumeAttachmentTypeArrayOutput() VolumeAttachmentTypeArrayOutput {
	return o
}

func (o VolumeAttachmentTypeArrayOutput) ToVolumeAttachmentTypeArrayOutputWithContext(ctx context.Context) VolumeAttachmentTypeArrayOutput {
	return o
}

func (o VolumeAttachmentTypeArrayOutput) Index(i pulumi.IntInput) VolumeAttachmentTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeAttachmentType {
		return vs[0].([]VolumeAttachmentType)[vs[1].(int)]
	}).(VolumeAttachmentTypeOutput)
}

// VolumeAttachmentList is a collection of VolumeAttachment objects.
type VolumeAttachmentListType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is the list of VolumeAttachments
	Items []VolumeAttachmentType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// VolumeAttachmentListTypeInput is an input type that accepts VolumeAttachmentListTypeArgs and VolumeAttachmentListTypeOutput values.
// You can construct a concrete instance of `VolumeAttachmentListTypeInput` via:
//
//          VolumeAttachmentListTypeArgs{...}
type VolumeAttachmentListTypeInput interface {
	pulumi.Input

	ToVolumeAttachmentListTypeOutput() VolumeAttachmentListTypeOutput
	ToVolumeAttachmentListTypeOutputWithContext(context.Context) VolumeAttachmentListTypeOutput
}

// VolumeAttachmentList is a collection of VolumeAttachment objects.
type VolumeAttachmentListTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// Items is the list of VolumeAttachments
	Items VolumeAttachmentTypeArrayInput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput `pulumi:"metadata"`
}

func (VolumeAttachmentListTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentListType)(nil)).Elem()
}

func (i VolumeAttachmentListTypeArgs) ToVolumeAttachmentListTypeOutput() VolumeAttachmentListTypeOutput {
	return i.ToVolumeAttachmentListTypeOutputWithContext(context.Background())
}

func (i VolumeAttachmentListTypeArgs) ToVolumeAttachmentListTypeOutputWithContext(ctx context.Context) VolumeAttachmentListTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentListTypeOutput)
}

// VolumeAttachmentList is a collection of VolumeAttachment objects.
type VolumeAttachmentListTypeOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentListTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentListType)(nil)).Elem()
}

func (o VolumeAttachmentListTypeOutput) ToVolumeAttachmentListTypeOutput() VolumeAttachmentListTypeOutput {
	return o
}

func (o VolumeAttachmentListTypeOutput) ToVolumeAttachmentListTypeOutputWithContext(ctx context.Context) VolumeAttachmentListTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o VolumeAttachmentListTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentListType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Items is the list of VolumeAttachments
func (o VolumeAttachmentListTypeOutput) Items() VolumeAttachmentTypeArrayOutput {
	return o.ApplyT(func(v VolumeAttachmentListType) []VolumeAttachmentType { return v.Items }).(VolumeAttachmentTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o VolumeAttachmentListTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentListType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o VolumeAttachmentListTypeOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentListType) *metav1.ListMeta { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSource struct {
	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
	InlineVolumeSpec *corev1.PersistentVolumeSpec `pulumi:"inlineVolumeSpec"`
	// Name of the persistent volume to attach.
	PersistentVolumeName *string `pulumi:"persistentVolumeName"`
}

// VolumeAttachmentSourceInput is an input type that accepts VolumeAttachmentSourceArgs and VolumeAttachmentSourceOutput values.
// You can construct a concrete instance of `VolumeAttachmentSourceInput` via:
//
//          VolumeAttachmentSourceArgs{...}
type VolumeAttachmentSourceInput interface {
	pulumi.Input

	ToVolumeAttachmentSourceOutput() VolumeAttachmentSourceOutput
	ToVolumeAttachmentSourceOutputWithContext(context.Context) VolumeAttachmentSourceOutput
}

// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSourceArgs struct {
	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
	InlineVolumeSpec corev1.PersistentVolumeSpecPtrInput `pulumi:"inlineVolumeSpec"`
	// Name of the persistent volume to attach.
	PersistentVolumeName pulumi.StringPtrInput `pulumi:"persistentVolumeName"`
}

func (VolumeAttachmentSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSource)(nil)).Elem()
}

func (i VolumeAttachmentSourceArgs) ToVolumeAttachmentSourceOutput() VolumeAttachmentSourceOutput {
	return i.ToVolumeAttachmentSourceOutputWithContext(context.Background())
}

func (i VolumeAttachmentSourceArgs) ToVolumeAttachmentSourceOutputWithContext(ctx context.Context) VolumeAttachmentSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentSourceOutput)
}

// VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
type VolumeAttachmentSourceOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSource)(nil)).Elem()
}

func (o VolumeAttachmentSourceOutput) ToVolumeAttachmentSourceOutput() VolumeAttachmentSourceOutput {
	return o
}

func (o VolumeAttachmentSourceOutput) ToVolumeAttachmentSourceOutputWithContext(ctx context.Context) VolumeAttachmentSourceOutput {
	return o
}

// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
func (o VolumeAttachmentSourceOutput) InlineVolumeSpec() corev1.PersistentVolumeSpecPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentSource) *corev1.PersistentVolumeSpec { return v.InlineVolumeSpec }).(corev1.PersistentVolumeSpecPtrOutput)
}

// Name of the persistent volume to attach.
func (o VolumeAttachmentSourceOutput) PersistentVolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentSource) *string { return v.PersistentVolumeName }).(pulumi.StringPtrOutput)
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpec struct {
	// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Attacher string `pulumi:"attacher"`
	// The node that the volume should be attached to.
	NodeName string `pulumi:"nodeName"`
	// Source represents the volume that should be attached.
	Source VolumeAttachmentSource `pulumi:"source"`
}

// VolumeAttachmentSpecInput is an input type that accepts VolumeAttachmentSpecArgs and VolumeAttachmentSpecOutput values.
// You can construct a concrete instance of `VolumeAttachmentSpecInput` via:
//
//          VolumeAttachmentSpecArgs{...}
type VolumeAttachmentSpecInput interface {
	pulumi.Input

	ToVolumeAttachmentSpecOutput() VolumeAttachmentSpecOutput
	ToVolumeAttachmentSpecOutputWithContext(context.Context) VolumeAttachmentSpecOutput
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpecArgs struct {
	// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	Attacher pulumi.StringInput `pulumi:"attacher"`
	// The node that the volume should be attached to.
	NodeName pulumi.StringInput `pulumi:"nodeName"`
	// Source represents the volume that should be attached.
	Source VolumeAttachmentSourceInput `pulumi:"source"`
}

func (VolumeAttachmentSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSpec)(nil)).Elem()
}

func (i VolumeAttachmentSpecArgs) ToVolumeAttachmentSpecOutput() VolumeAttachmentSpecOutput {
	return i.ToVolumeAttachmentSpecOutputWithContext(context.Background())
}

func (i VolumeAttachmentSpecArgs) ToVolumeAttachmentSpecOutputWithContext(ctx context.Context) VolumeAttachmentSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentSpecOutput)
}

// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpecOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentSpec)(nil)).Elem()
}

func (o VolumeAttachmentSpecOutput) ToVolumeAttachmentSpecOutput() VolumeAttachmentSpecOutput {
	return o
}

func (o VolumeAttachmentSpecOutput) ToVolumeAttachmentSpecOutputWithContext(ctx context.Context) VolumeAttachmentSpecOutput {
	return o
}

// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
func (o VolumeAttachmentSpecOutput) Attacher() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeAttachmentSpec) string { return v.Attacher }).(pulumi.StringOutput)
}

// The node that the volume should be attached to.
func (o VolumeAttachmentSpecOutput) NodeName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeAttachmentSpec) string { return v.NodeName }).(pulumi.StringOutput)
}

// Source represents the volume that should be attached.
func (o VolumeAttachmentSpecOutput) Source() VolumeAttachmentSourceOutput {
	return o.ApplyT(func(v VolumeAttachmentSpec) VolumeAttachmentSource { return v.Source }).(VolumeAttachmentSourceOutput)
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatus struct {
	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachError *VolumeError `pulumi:"attachError"`
	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attached bool `pulumi:"attached"`
	// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata map[string]string `pulumi:"attachmentMetadata"`
	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
	DetachError *VolumeError `pulumi:"detachError"`
}

// VolumeAttachmentStatusInput is an input type that accepts VolumeAttachmentStatusArgs and VolumeAttachmentStatusOutput values.
// You can construct a concrete instance of `VolumeAttachmentStatusInput` via:
//
//          VolumeAttachmentStatusArgs{...}
type VolumeAttachmentStatusInput interface {
	pulumi.Input

	ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput
	ToVolumeAttachmentStatusOutputWithContext(context.Context) VolumeAttachmentStatusOutput
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatusArgs struct {
	// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachError VolumeErrorPtrInput `pulumi:"attachError"`
	// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	Attached pulumi.BoolInput `pulumi:"attached"`
	// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
	AttachmentMetadata pulumi.StringMapInput `pulumi:"attachmentMetadata"`
	// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
	DetachError VolumeErrorPtrInput `pulumi:"detachError"`
}

func (VolumeAttachmentStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentStatus)(nil)).Elem()
}

func (i VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput {
	return i.ToVolumeAttachmentStatusOutputWithContext(context.Background())
}

func (i VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusOutputWithContext(ctx context.Context) VolumeAttachmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentStatusOutput)
}

func (i VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusPtrOutput() VolumeAttachmentStatusPtrOutput {
	return i.ToVolumeAttachmentStatusPtrOutputWithContext(context.Background())
}

func (i VolumeAttachmentStatusArgs) ToVolumeAttachmentStatusPtrOutputWithContext(ctx context.Context) VolumeAttachmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentStatusOutput).ToVolumeAttachmentStatusPtrOutputWithContext(ctx)
}

// VolumeAttachmentStatusPtrInput is an input type that accepts VolumeAttachmentStatusArgs, VolumeAttachmentStatusPtr and VolumeAttachmentStatusPtrOutput values.
// You can construct a concrete instance of `VolumeAttachmentStatusPtrInput` via:
//
//          VolumeAttachmentStatusArgs{...}
//
//  or:
//
//          nil
type VolumeAttachmentStatusPtrInput interface {
	pulumi.Input

	ToVolumeAttachmentStatusPtrOutput() VolumeAttachmentStatusPtrOutput
	ToVolumeAttachmentStatusPtrOutputWithContext(context.Context) VolumeAttachmentStatusPtrOutput
}

type volumeAttachmentStatusPtrType VolumeAttachmentStatusArgs

func VolumeAttachmentStatusPtr(v *VolumeAttachmentStatusArgs) VolumeAttachmentStatusPtrInput {
	return (*volumeAttachmentStatusPtrType)(v)
}

func (*volumeAttachmentStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachmentStatus)(nil)).Elem()
}

func (i *volumeAttachmentStatusPtrType) ToVolumeAttachmentStatusPtrOutput() VolumeAttachmentStatusPtrOutput {
	return i.ToVolumeAttachmentStatusPtrOutputWithContext(context.Background())
}

func (i *volumeAttachmentStatusPtrType) ToVolumeAttachmentStatusPtrOutputWithContext(ctx context.Context) VolumeAttachmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentStatusPtrOutput)
}

// VolumeAttachmentStatus is the status of a VolumeAttachment request.
type VolumeAttachmentStatusOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentStatus)(nil)).Elem()
}

func (o VolumeAttachmentStatusOutput) ToVolumeAttachmentStatusOutput() VolumeAttachmentStatusOutput {
	return o
}

func (o VolumeAttachmentStatusOutput) ToVolumeAttachmentStatusOutputWithContext(ctx context.Context) VolumeAttachmentStatusOutput {
	return o
}

func (o VolumeAttachmentStatusOutput) ToVolumeAttachmentStatusPtrOutput() VolumeAttachmentStatusPtrOutput {
	return o.ToVolumeAttachmentStatusPtrOutputWithContext(context.Background())
}

func (o VolumeAttachmentStatusOutput) ToVolumeAttachmentStatusPtrOutputWithContext(ctx context.Context) VolumeAttachmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeAttachmentStatus) *VolumeAttachmentStatus {
		return &v
	}).(VolumeAttachmentStatusPtrOutput)
}

// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) AttachError() VolumeErrorPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentStatus) *VolumeError { return v.AttachError }).(VolumeErrorPtrOutput)
}

// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) Attached() pulumi.BoolOutput {
	return o.ApplyT(func(v VolumeAttachmentStatus) bool { return v.Attached }).(pulumi.BoolOutput)
}

// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) AttachmentMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v VolumeAttachmentStatus) map[string]string { return v.AttachmentMetadata }).(pulumi.StringMapOutput)
}

// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusOutput) DetachError() VolumeErrorPtrOutput {
	return o.ApplyT(func(v VolumeAttachmentStatus) *VolumeError { return v.DetachError }).(VolumeErrorPtrOutput)
}

type VolumeAttachmentStatusPtrOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachmentStatus)(nil)).Elem()
}

func (o VolumeAttachmentStatusPtrOutput) ToVolumeAttachmentStatusPtrOutput() VolumeAttachmentStatusPtrOutput {
	return o
}

func (o VolumeAttachmentStatusPtrOutput) ToVolumeAttachmentStatusPtrOutputWithContext(ctx context.Context) VolumeAttachmentStatusPtrOutput {
	return o
}

func (o VolumeAttachmentStatusPtrOutput) Elem() VolumeAttachmentStatusOutput {
	return o.ApplyT(func(v *VolumeAttachmentStatus) VolumeAttachmentStatus {
		if v != nil {
			return *v
		}
		var ret VolumeAttachmentStatus
		return ret
	}).(VolumeAttachmentStatusOutput)
}

// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusPtrOutput) AttachError() VolumeErrorPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentStatus) *VolumeError {
		if v == nil {
			return nil
		}
		return v.AttachError
	}).(VolumeErrorPtrOutput)
}

// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusPtrOutput) Attached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentStatus) *bool {
		if v == nil {
			return nil
		}
		return &v.Attached
	}).(pulumi.BoolPtrOutput)
}

// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusPtrOutput) AttachmentMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VolumeAttachmentStatus) map[string]string {
		if v == nil {
			return nil
		}
		return v.AttachmentMetadata
	}).(pulumi.StringMapOutput)
}

// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
func (o VolumeAttachmentStatusPtrOutput) DetachError() VolumeErrorPtrOutput {
	return o.ApplyT(func(v *VolumeAttachmentStatus) *VolumeError {
		if v == nil {
			return nil
		}
		return v.DetachError
	}).(VolumeErrorPtrOutput)
}

// VolumeError captures an error encountered during a volume operation.
type VolumeError struct {
	// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Message *string `pulumi:"message"`
	// Time the error was encountered.
	Time *string `pulumi:"time"`
}

// VolumeErrorInput is an input type that accepts VolumeErrorArgs and VolumeErrorOutput values.
// You can construct a concrete instance of `VolumeErrorInput` via:
//
//          VolumeErrorArgs{...}
type VolumeErrorInput interface {
	pulumi.Input

	ToVolumeErrorOutput() VolumeErrorOutput
	ToVolumeErrorOutputWithContext(context.Context) VolumeErrorOutput
}

// VolumeError captures an error encountered during a volume operation.
type VolumeErrorArgs struct {
	// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Message pulumi.StringPtrInput `pulumi:"message"`
	// Time the error was encountered.
	Time pulumi.StringPtrInput `pulumi:"time"`
}

func (VolumeErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeError)(nil)).Elem()
}

func (i VolumeErrorArgs) ToVolumeErrorOutput() VolumeErrorOutput {
	return i.ToVolumeErrorOutputWithContext(context.Background())
}

func (i VolumeErrorArgs) ToVolumeErrorOutputWithContext(ctx context.Context) VolumeErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeErrorOutput)
}

func (i VolumeErrorArgs) ToVolumeErrorPtrOutput() VolumeErrorPtrOutput {
	return i.ToVolumeErrorPtrOutputWithContext(context.Background())
}

func (i VolumeErrorArgs) ToVolumeErrorPtrOutputWithContext(ctx context.Context) VolumeErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeErrorOutput).ToVolumeErrorPtrOutputWithContext(ctx)
}

// VolumeErrorPtrInput is an input type that accepts VolumeErrorArgs, VolumeErrorPtr and VolumeErrorPtrOutput values.
// You can construct a concrete instance of `VolumeErrorPtrInput` via:
//
//          VolumeErrorArgs{...}
//
//  or:
//
//          nil
type VolumeErrorPtrInput interface {
	pulumi.Input

	ToVolumeErrorPtrOutput() VolumeErrorPtrOutput
	ToVolumeErrorPtrOutputWithContext(context.Context) VolumeErrorPtrOutput
}

type volumeErrorPtrType VolumeErrorArgs

func VolumeErrorPtr(v *VolumeErrorArgs) VolumeErrorPtrInput {
	return (*volumeErrorPtrType)(v)
}

func (*volumeErrorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeError)(nil)).Elem()
}

func (i *volumeErrorPtrType) ToVolumeErrorPtrOutput() VolumeErrorPtrOutput {
	return i.ToVolumeErrorPtrOutputWithContext(context.Background())
}

func (i *volumeErrorPtrType) ToVolumeErrorPtrOutputWithContext(ctx context.Context) VolumeErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeErrorPtrOutput)
}

// VolumeError captures an error encountered during a volume operation.
type VolumeErrorOutput struct{ *pulumi.OutputState }

func (VolumeErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeError)(nil)).Elem()
}

func (o VolumeErrorOutput) ToVolumeErrorOutput() VolumeErrorOutput {
	return o
}

func (o VolumeErrorOutput) ToVolumeErrorOutputWithContext(ctx context.Context) VolumeErrorOutput {
	return o
}

func (o VolumeErrorOutput) ToVolumeErrorPtrOutput() VolumeErrorPtrOutput {
	return o.ToVolumeErrorPtrOutputWithContext(context.Background())
}

func (o VolumeErrorOutput) ToVolumeErrorPtrOutputWithContext(ctx context.Context) VolumeErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeError) *VolumeError {
		return &v
	}).(VolumeErrorPtrOutput)
}

// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
func (o VolumeErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeError) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// Time the error was encountered.
func (o VolumeErrorOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeError) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type VolumeErrorPtrOutput struct{ *pulumi.OutputState }

func (VolumeErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeError)(nil)).Elem()
}

func (o VolumeErrorPtrOutput) ToVolumeErrorPtrOutput() VolumeErrorPtrOutput {
	return o
}

func (o VolumeErrorPtrOutput) ToVolumeErrorPtrOutputWithContext(ctx context.Context) VolumeErrorPtrOutput {
	return o
}

func (o VolumeErrorPtrOutput) Elem() VolumeErrorOutput {
	return o.ApplyT(func(v *VolumeError) VolumeError {
		if v != nil {
			return *v
		}
		var ret VolumeError
		return ret
	}).(VolumeErrorOutput)
}

// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
func (o VolumeErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeError) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

// Time the error was encountered.
func (o VolumeErrorPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeError) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentTypeInput)(nil)).Elem(), VolumeAttachmentTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentTypeArrayInput)(nil)).Elem(), VolumeAttachmentTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentListTypeInput)(nil)).Elem(), VolumeAttachmentListTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentSourceInput)(nil)).Elem(), VolumeAttachmentSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentSpecInput)(nil)).Elem(), VolumeAttachmentSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentStatusInput)(nil)).Elem(), VolumeAttachmentStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentStatusPtrInput)(nil)).Elem(), VolumeAttachmentStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeErrorInput)(nil)).Elem(), VolumeErrorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeErrorPtrInput)(nil)).Elem(), VolumeErrorArgs{})
	pulumi.RegisterOutputType(VolumeAttachmentTypeOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentTypeArrayOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentListTypeOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentSourceOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentSpecOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentStatusOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentStatusPtrOutput{})
	pulumi.RegisterOutputType(VolumeErrorOutput{})
	pulumi.RegisterOutputType(VolumeErrorPtrOutput{})
}
