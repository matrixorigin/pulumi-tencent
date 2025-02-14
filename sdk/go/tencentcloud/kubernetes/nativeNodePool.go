// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NativeNodePool struct {
	pulumi.CustomResourceState

	// Node Annotation List.
	Annotations NativeNodePoolAnnotationArrayOutput `pulumi:"annotations"`
	// ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation time.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Whether to enable deletion protection.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// Node Labels.
	Labels NativeNodePoolLabelArrayOutput `pulumi:"labels"`
	// Node pool status.
	LifeState pulumi.StringOutput `pulumi:"lifeState"`
	// Node pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Native node pool creation parameters.
	Native NativeNodePoolNativeOutput `pulumi:"native"`
	// Node tags.
	Tags NativeNodePoolTagArrayOutput `pulumi:"tags"`
	// Node taint.
	Taints NativeNodePoolTaintArrayOutput `pulumi:"taints"`
	// Node pool type. Optional value is `Native`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
	Unschedulable pulumi.BoolOutput `pulumi:"unschedulable"`
}

// NewNativeNodePool registers a new resource with the given unique name, arguments, and options.
func NewNativeNodePool(ctx *pulumi.Context,
	name string, args *NativeNodePoolArgs, opts ...pulumi.ResourceOption) (*NativeNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Native == nil {
		return nil, errors.New("invalid value for required argument 'Native'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NativeNodePool
	err := ctx.RegisterResource("tencentcloud:Kubernetes/nativeNodePool:NativeNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNativeNodePool gets an existing NativeNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNativeNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NativeNodePoolState, opts ...pulumi.ResourceOption) (*NativeNodePool, error) {
	var resource NativeNodePool
	err := ctx.ReadResource("tencentcloud:Kubernetes/nativeNodePool:NativeNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NativeNodePool resources.
type nativeNodePoolState struct {
	// Node Annotation List.
	Annotations []NativeNodePoolAnnotation `pulumi:"annotations"`
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Creation time.
	CreatedAt *string `pulumi:"createdAt"`
	// Whether to enable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Node Labels.
	Labels []NativeNodePoolLabel `pulumi:"labels"`
	// Node pool status.
	LifeState *string `pulumi:"lifeState"`
	// Node pool name.
	Name *string `pulumi:"name"`
	// Native node pool creation parameters.
	Native *NativeNodePoolNative `pulumi:"native"`
	// Node tags.
	Tags []NativeNodePoolTag `pulumi:"tags"`
	// Node taint.
	Taints []NativeNodePoolTaint `pulumi:"taints"`
	// Node pool type. Optional value is `Native`.
	Type *string `pulumi:"type"`
	// Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
	Unschedulable *bool `pulumi:"unschedulable"`
}

type NativeNodePoolState struct {
	// Node Annotation List.
	Annotations NativeNodePoolAnnotationArrayInput
	// ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// Creation time.
	CreatedAt pulumi.StringPtrInput
	// Whether to enable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Node Labels.
	Labels NativeNodePoolLabelArrayInput
	// Node pool status.
	LifeState pulumi.StringPtrInput
	// Node pool name.
	Name pulumi.StringPtrInput
	// Native node pool creation parameters.
	Native NativeNodePoolNativePtrInput
	// Node tags.
	Tags NativeNodePoolTagArrayInput
	// Node taint.
	Taints NativeNodePoolTaintArrayInput
	// Node pool type. Optional value is `Native`.
	Type pulumi.StringPtrInput
	// Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
	Unschedulable pulumi.BoolPtrInput
}

func (NativeNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*nativeNodePoolState)(nil)).Elem()
}

type nativeNodePoolArgs struct {
	// Node Annotation List.
	Annotations []NativeNodePoolAnnotation `pulumi:"annotations"`
	// ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// Whether to enable deletion protection.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Node Labels.
	Labels []NativeNodePoolLabel `pulumi:"labels"`
	// Node pool name.
	Name *string `pulumi:"name"`
	// Native node pool creation parameters.
	Native NativeNodePoolNative `pulumi:"native"`
	// Node tags.
	Tags []NativeNodePoolTag `pulumi:"tags"`
	// Node taint.
	Taints []NativeNodePoolTaint `pulumi:"taints"`
	// Node pool type. Optional value is `Native`.
	Type string `pulumi:"type"`
	// Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
	Unschedulable *bool `pulumi:"unschedulable"`
}

// The set of arguments for constructing a NativeNodePool resource.
type NativeNodePoolArgs struct {
	// Node Annotation List.
	Annotations NativeNodePoolAnnotationArrayInput
	// ID of the cluster.
	ClusterId pulumi.StringInput
	// Whether to enable deletion protection.
	DeletionProtection pulumi.BoolPtrInput
	// Node Labels.
	Labels NativeNodePoolLabelArrayInput
	// Node pool name.
	Name pulumi.StringPtrInput
	// Native node pool creation parameters.
	Native NativeNodePoolNativeInput
	// Node tags.
	Tags NativeNodePoolTagArrayInput
	// Node taint.
	Taints NativeNodePoolTaintArrayInput
	// Node pool type. Optional value is `Native`.
	Type pulumi.StringInput
	// Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
	Unschedulable pulumi.BoolPtrInput
}

func (NativeNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nativeNodePoolArgs)(nil)).Elem()
}

type NativeNodePoolInput interface {
	pulumi.Input

	ToNativeNodePoolOutput() NativeNodePoolOutput
	ToNativeNodePoolOutputWithContext(ctx context.Context) NativeNodePoolOutput
}

func (*NativeNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**NativeNodePool)(nil)).Elem()
}

func (i *NativeNodePool) ToNativeNodePoolOutput() NativeNodePoolOutput {
	return i.ToNativeNodePoolOutputWithContext(context.Background())
}

func (i *NativeNodePool) ToNativeNodePoolOutputWithContext(ctx context.Context) NativeNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NativeNodePoolOutput)
}

// NativeNodePoolArrayInput is an input type that accepts NativeNodePoolArray and NativeNodePoolArrayOutput values.
// You can construct a concrete instance of `NativeNodePoolArrayInput` via:
//
//	NativeNodePoolArray{ NativeNodePoolArgs{...} }
type NativeNodePoolArrayInput interface {
	pulumi.Input

	ToNativeNodePoolArrayOutput() NativeNodePoolArrayOutput
	ToNativeNodePoolArrayOutputWithContext(context.Context) NativeNodePoolArrayOutput
}

type NativeNodePoolArray []NativeNodePoolInput

func (NativeNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NativeNodePool)(nil)).Elem()
}

func (i NativeNodePoolArray) ToNativeNodePoolArrayOutput() NativeNodePoolArrayOutput {
	return i.ToNativeNodePoolArrayOutputWithContext(context.Background())
}

func (i NativeNodePoolArray) ToNativeNodePoolArrayOutputWithContext(ctx context.Context) NativeNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NativeNodePoolArrayOutput)
}

// NativeNodePoolMapInput is an input type that accepts NativeNodePoolMap and NativeNodePoolMapOutput values.
// You can construct a concrete instance of `NativeNodePoolMapInput` via:
//
//	NativeNodePoolMap{ "key": NativeNodePoolArgs{...} }
type NativeNodePoolMapInput interface {
	pulumi.Input

	ToNativeNodePoolMapOutput() NativeNodePoolMapOutput
	ToNativeNodePoolMapOutputWithContext(context.Context) NativeNodePoolMapOutput
}

type NativeNodePoolMap map[string]NativeNodePoolInput

func (NativeNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NativeNodePool)(nil)).Elem()
}

func (i NativeNodePoolMap) ToNativeNodePoolMapOutput() NativeNodePoolMapOutput {
	return i.ToNativeNodePoolMapOutputWithContext(context.Background())
}

func (i NativeNodePoolMap) ToNativeNodePoolMapOutputWithContext(ctx context.Context) NativeNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NativeNodePoolMapOutput)
}

type NativeNodePoolOutput struct{ *pulumi.OutputState }

func (NativeNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NativeNodePool)(nil)).Elem()
}

func (o NativeNodePoolOutput) ToNativeNodePoolOutput() NativeNodePoolOutput {
	return o
}

func (o NativeNodePoolOutput) ToNativeNodePoolOutputWithContext(ctx context.Context) NativeNodePoolOutput {
	return o
}

// Node Annotation List.
func (o NativeNodePoolOutput) Annotations() NativeNodePoolAnnotationArrayOutput {
	return o.ApplyT(func(v *NativeNodePool) NativeNodePoolAnnotationArrayOutput { return v.Annotations }).(NativeNodePoolAnnotationArrayOutput)
}

// ID of the cluster.
func (o NativeNodePoolOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation time.
func (o NativeNodePoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Whether to enable deletion protection.
func (o NativeNodePoolOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.BoolOutput { return v.DeletionProtection }).(pulumi.BoolOutput)
}

// Node Labels.
func (o NativeNodePoolOutput) Labels() NativeNodePoolLabelArrayOutput {
	return o.ApplyT(func(v *NativeNodePool) NativeNodePoolLabelArrayOutput { return v.Labels }).(NativeNodePoolLabelArrayOutput)
}

// Node pool status.
func (o NativeNodePoolOutput) LifeState() pulumi.StringOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.StringOutput { return v.LifeState }).(pulumi.StringOutput)
}

// Node pool name.
func (o NativeNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Native node pool creation parameters.
func (o NativeNodePoolOutput) Native() NativeNodePoolNativeOutput {
	return o.ApplyT(func(v *NativeNodePool) NativeNodePoolNativeOutput { return v.Native }).(NativeNodePoolNativeOutput)
}

// Node tags.
func (o NativeNodePoolOutput) Tags() NativeNodePoolTagArrayOutput {
	return o.ApplyT(func(v *NativeNodePool) NativeNodePoolTagArrayOutput { return v.Tags }).(NativeNodePoolTagArrayOutput)
}

// Node taint.
func (o NativeNodePoolOutput) Taints() NativeNodePoolTaintArrayOutput {
	return o.ApplyT(func(v *NativeNodePool) NativeNodePoolTaintArrayOutput { return v.Taints }).(NativeNodePoolTaintArrayOutput)
}

// Node pool type. Optional value is `Native`.
func (o NativeNodePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
func (o NativeNodePoolOutput) Unschedulable() pulumi.BoolOutput {
	return o.ApplyT(func(v *NativeNodePool) pulumi.BoolOutput { return v.Unschedulable }).(pulumi.BoolOutput)
}

type NativeNodePoolArrayOutput struct{ *pulumi.OutputState }

func (NativeNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NativeNodePool)(nil)).Elem()
}

func (o NativeNodePoolArrayOutput) ToNativeNodePoolArrayOutput() NativeNodePoolArrayOutput {
	return o
}

func (o NativeNodePoolArrayOutput) ToNativeNodePoolArrayOutputWithContext(ctx context.Context) NativeNodePoolArrayOutput {
	return o
}

func (o NativeNodePoolArrayOutput) Index(i pulumi.IntInput) NativeNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NativeNodePool {
		return vs[0].([]*NativeNodePool)[vs[1].(int)]
	}).(NativeNodePoolOutput)
}

type NativeNodePoolMapOutput struct{ *pulumi.OutputState }

func (NativeNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NativeNodePool)(nil)).Elem()
}

func (o NativeNodePoolMapOutput) ToNativeNodePoolMapOutput() NativeNodePoolMapOutput {
	return o
}

func (o NativeNodePoolMapOutput) ToNativeNodePoolMapOutputWithContext(ctx context.Context) NativeNodePoolMapOutput {
	return o
}

func (o NativeNodePoolMapOutput) MapIndex(k pulumi.StringInput) NativeNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NativeNodePool {
		return vs[0].(map[string]*NativeNodePool)[vs[1].(string)]
	}).(NativeNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NativeNodePoolInput)(nil)).Elem(), &NativeNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*NativeNodePoolArrayInput)(nil)).Elem(), NativeNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NativeNodePoolMapInput)(nil)).Elem(), NativeNodePoolMap{})
	pulumi.RegisterOutputType(NativeNodePoolOutput{})
	pulumi.RegisterOutputType(NativeNodePoolArrayOutput{})
	pulumi.RegisterOutputType(NativeNodePoolMapOutput{})
}
