// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lite

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HbaseInstance struct {
	pulumi.CustomResourceState

	// Instance single-node disk capacity, in GB. The single-node disk capacity must be greater than or equal to 100 and less
	// than or equal to 10000, with an adjustment step size of 20.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// Instance disk type, fill in CLOUD_HSSD to indicate performance cloud storage.
	DiskType pulumi.StringOutput `pulumi:"diskType"`
	// Instance name. Length limit is 6-36 characters. Only Chinese characters, letters, numbers, -, and _ are allowed.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Instance node type, can be filled in as 4C16G, 8C32G, 16C64G, 32C128G, case insensitive.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// Instance pay mode. Value range: 0: indicates post pay mode, that is, pay-as-you-go.
	PayMode pulumi.IntOutput `pulumi:"payMode"`
	// List of tags to bind to the instance.
	Tags HbaseInstanceTagArrayOutput `pulumi:"tags"`
	// Detailed configuration of the instance availability zone, currently supports multiple availability zones, the number of
	// availability zones can only be 1 or 3, including zone name, VPC information, and number of nodes. The total number of
	// nodes across all zones must be greater than or equal to 3 and less than or equal to 50.
	ZoneSettings HbaseInstanceZoneSettingArrayOutput `pulumi:"zoneSettings"`
}

// NewHbaseInstance registers a new resource with the given unique name, arguments, and options.
func NewHbaseInstance(ctx *pulumi.Context,
	name string, args *HbaseInstanceArgs, opts ...pulumi.ResourceOption) (*HbaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskSize == nil {
		return nil, errors.New("invalid value for required argument 'DiskSize'")
	}
	if args.DiskType == nil {
		return nil, errors.New("invalid value for required argument 'DiskType'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	if args.PayMode == nil {
		return nil, errors.New("invalid value for required argument 'PayMode'")
	}
	if args.ZoneSettings == nil {
		return nil, errors.New("invalid value for required argument 'ZoneSettings'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HbaseInstance
	err := ctx.RegisterResource("tencentcloud:Lite/hbaseInstance:HbaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHbaseInstance gets an existing HbaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHbaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HbaseInstanceState, opts ...pulumi.ResourceOption) (*HbaseInstance, error) {
	var resource HbaseInstance
	err := ctx.ReadResource("tencentcloud:Lite/hbaseInstance:HbaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HbaseInstance resources.
type hbaseInstanceState struct {
	// Instance single-node disk capacity, in GB. The single-node disk capacity must be greater than or equal to 100 and less
	// than or equal to 10000, with an adjustment step size of 20.
	DiskSize *int `pulumi:"diskSize"`
	// Instance disk type, fill in CLOUD_HSSD to indicate performance cloud storage.
	DiskType *string `pulumi:"diskType"`
	// Instance name. Length limit is 6-36 characters. Only Chinese characters, letters, numbers, -, and _ are allowed.
	InstanceName *string `pulumi:"instanceName"`
	// Instance node type, can be filled in as 4C16G, 8C32G, 16C64G, 32C128G, case insensitive.
	NodeType *string `pulumi:"nodeType"`
	// Instance pay mode. Value range: 0: indicates post pay mode, that is, pay-as-you-go.
	PayMode *int `pulumi:"payMode"`
	// List of tags to bind to the instance.
	Tags []HbaseInstanceTag `pulumi:"tags"`
	// Detailed configuration of the instance availability zone, currently supports multiple availability zones, the number of
	// availability zones can only be 1 or 3, including zone name, VPC information, and number of nodes. The total number of
	// nodes across all zones must be greater than or equal to 3 and less than or equal to 50.
	ZoneSettings []HbaseInstanceZoneSetting `pulumi:"zoneSettings"`
}

type HbaseInstanceState struct {
	// Instance single-node disk capacity, in GB. The single-node disk capacity must be greater than or equal to 100 and less
	// than or equal to 10000, with an adjustment step size of 20.
	DiskSize pulumi.IntPtrInput
	// Instance disk type, fill in CLOUD_HSSD to indicate performance cloud storage.
	DiskType pulumi.StringPtrInput
	// Instance name. Length limit is 6-36 characters. Only Chinese characters, letters, numbers, -, and _ are allowed.
	InstanceName pulumi.StringPtrInput
	// Instance node type, can be filled in as 4C16G, 8C32G, 16C64G, 32C128G, case insensitive.
	NodeType pulumi.StringPtrInput
	// Instance pay mode. Value range: 0: indicates post pay mode, that is, pay-as-you-go.
	PayMode pulumi.IntPtrInput
	// List of tags to bind to the instance.
	Tags HbaseInstanceTagArrayInput
	// Detailed configuration of the instance availability zone, currently supports multiple availability zones, the number of
	// availability zones can only be 1 or 3, including zone name, VPC information, and number of nodes. The total number of
	// nodes across all zones must be greater than or equal to 3 and less than or equal to 50.
	ZoneSettings HbaseInstanceZoneSettingArrayInput
}

func (HbaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hbaseInstanceState)(nil)).Elem()
}

type hbaseInstanceArgs struct {
	// Instance single-node disk capacity, in GB. The single-node disk capacity must be greater than or equal to 100 and less
	// than or equal to 10000, with an adjustment step size of 20.
	DiskSize int `pulumi:"diskSize"`
	// Instance disk type, fill in CLOUD_HSSD to indicate performance cloud storage.
	DiskType string `pulumi:"diskType"`
	// Instance name. Length limit is 6-36 characters. Only Chinese characters, letters, numbers, -, and _ are allowed.
	InstanceName string `pulumi:"instanceName"`
	// Instance node type, can be filled in as 4C16G, 8C32G, 16C64G, 32C128G, case insensitive.
	NodeType string `pulumi:"nodeType"`
	// Instance pay mode. Value range: 0: indicates post pay mode, that is, pay-as-you-go.
	PayMode int `pulumi:"payMode"`
	// List of tags to bind to the instance.
	Tags []HbaseInstanceTag `pulumi:"tags"`
	// Detailed configuration of the instance availability zone, currently supports multiple availability zones, the number of
	// availability zones can only be 1 or 3, including zone name, VPC information, and number of nodes. The total number of
	// nodes across all zones must be greater than or equal to 3 and less than or equal to 50.
	ZoneSettings []HbaseInstanceZoneSetting `pulumi:"zoneSettings"`
}

// The set of arguments for constructing a HbaseInstance resource.
type HbaseInstanceArgs struct {
	// Instance single-node disk capacity, in GB. The single-node disk capacity must be greater than or equal to 100 and less
	// than or equal to 10000, with an adjustment step size of 20.
	DiskSize pulumi.IntInput
	// Instance disk type, fill in CLOUD_HSSD to indicate performance cloud storage.
	DiskType pulumi.StringInput
	// Instance name. Length limit is 6-36 characters. Only Chinese characters, letters, numbers, -, and _ are allowed.
	InstanceName pulumi.StringInput
	// Instance node type, can be filled in as 4C16G, 8C32G, 16C64G, 32C128G, case insensitive.
	NodeType pulumi.StringInput
	// Instance pay mode. Value range: 0: indicates post pay mode, that is, pay-as-you-go.
	PayMode pulumi.IntInput
	// List of tags to bind to the instance.
	Tags HbaseInstanceTagArrayInput
	// Detailed configuration of the instance availability zone, currently supports multiple availability zones, the number of
	// availability zones can only be 1 or 3, including zone name, VPC information, and number of nodes. The total number of
	// nodes across all zones must be greater than or equal to 3 and less than or equal to 50.
	ZoneSettings HbaseInstanceZoneSettingArrayInput
}

func (HbaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hbaseInstanceArgs)(nil)).Elem()
}

type HbaseInstanceInput interface {
	pulumi.Input

	ToHbaseInstanceOutput() HbaseInstanceOutput
	ToHbaseInstanceOutputWithContext(ctx context.Context) HbaseInstanceOutput
}

func (*HbaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**HbaseInstance)(nil)).Elem()
}

func (i *HbaseInstance) ToHbaseInstanceOutput() HbaseInstanceOutput {
	return i.ToHbaseInstanceOutputWithContext(context.Background())
}

func (i *HbaseInstance) ToHbaseInstanceOutputWithContext(ctx context.Context) HbaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HbaseInstanceOutput)
}

// HbaseInstanceArrayInput is an input type that accepts HbaseInstanceArray and HbaseInstanceArrayOutput values.
// You can construct a concrete instance of `HbaseInstanceArrayInput` via:
//
//	HbaseInstanceArray{ HbaseInstanceArgs{...} }
type HbaseInstanceArrayInput interface {
	pulumi.Input

	ToHbaseInstanceArrayOutput() HbaseInstanceArrayOutput
	ToHbaseInstanceArrayOutputWithContext(context.Context) HbaseInstanceArrayOutput
}

type HbaseInstanceArray []HbaseInstanceInput

func (HbaseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HbaseInstance)(nil)).Elem()
}

func (i HbaseInstanceArray) ToHbaseInstanceArrayOutput() HbaseInstanceArrayOutput {
	return i.ToHbaseInstanceArrayOutputWithContext(context.Background())
}

func (i HbaseInstanceArray) ToHbaseInstanceArrayOutputWithContext(ctx context.Context) HbaseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HbaseInstanceArrayOutput)
}

// HbaseInstanceMapInput is an input type that accepts HbaseInstanceMap and HbaseInstanceMapOutput values.
// You can construct a concrete instance of `HbaseInstanceMapInput` via:
//
//	HbaseInstanceMap{ "key": HbaseInstanceArgs{...} }
type HbaseInstanceMapInput interface {
	pulumi.Input

	ToHbaseInstanceMapOutput() HbaseInstanceMapOutput
	ToHbaseInstanceMapOutputWithContext(context.Context) HbaseInstanceMapOutput
}

type HbaseInstanceMap map[string]HbaseInstanceInput

func (HbaseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HbaseInstance)(nil)).Elem()
}

func (i HbaseInstanceMap) ToHbaseInstanceMapOutput() HbaseInstanceMapOutput {
	return i.ToHbaseInstanceMapOutputWithContext(context.Background())
}

func (i HbaseInstanceMap) ToHbaseInstanceMapOutputWithContext(ctx context.Context) HbaseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HbaseInstanceMapOutput)
}

type HbaseInstanceOutput struct{ *pulumi.OutputState }

func (HbaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HbaseInstance)(nil)).Elem()
}

func (o HbaseInstanceOutput) ToHbaseInstanceOutput() HbaseInstanceOutput {
	return o
}

func (o HbaseInstanceOutput) ToHbaseInstanceOutputWithContext(ctx context.Context) HbaseInstanceOutput {
	return o
}

// Instance single-node disk capacity, in GB. The single-node disk capacity must be greater than or equal to 100 and less
// than or equal to 10000, with an adjustment step size of 20.
func (o HbaseInstanceOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *HbaseInstance) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// Instance disk type, fill in CLOUD_HSSD to indicate performance cloud storage.
func (o HbaseInstanceOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v *HbaseInstance) pulumi.StringOutput { return v.DiskType }).(pulumi.StringOutput)
}

// Instance name. Length limit is 6-36 characters. Only Chinese characters, letters, numbers, -, and _ are allowed.
func (o HbaseInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HbaseInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Instance node type, can be filled in as 4C16G, 8C32G, 16C64G, 32C128G, case insensitive.
func (o HbaseInstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *HbaseInstance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// Instance pay mode. Value range: 0: indicates post pay mode, that is, pay-as-you-go.
func (o HbaseInstanceOutput) PayMode() pulumi.IntOutput {
	return o.ApplyT(func(v *HbaseInstance) pulumi.IntOutput { return v.PayMode }).(pulumi.IntOutput)
}

// List of tags to bind to the instance.
func (o HbaseInstanceOutput) Tags() HbaseInstanceTagArrayOutput {
	return o.ApplyT(func(v *HbaseInstance) HbaseInstanceTagArrayOutput { return v.Tags }).(HbaseInstanceTagArrayOutput)
}

// Detailed configuration of the instance availability zone, currently supports multiple availability zones, the number of
// availability zones can only be 1 or 3, including zone name, VPC information, and number of nodes. The total number of
// nodes across all zones must be greater than or equal to 3 and less than or equal to 50.
func (o HbaseInstanceOutput) ZoneSettings() HbaseInstanceZoneSettingArrayOutput {
	return o.ApplyT(func(v *HbaseInstance) HbaseInstanceZoneSettingArrayOutput { return v.ZoneSettings }).(HbaseInstanceZoneSettingArrayOutput)
}

type HbaseInstanceArrayOutput struct{ *pulumi.OutputState }

func (HbaseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HbaseInstance)(nil)).Elem()
}

func (o HbaseInstanceArrayOutput) ToHbaseInstanceArrayOutput() HbaseInstanceArrayOutput {
	return o
}

func (o HbaseInstanceArrayOutput) ToHbaseInstanceArrayOutputWithContext(ctx context.Context) HbaseInstanceArrayOutput {
	return o
}

func (o HbaseInstanceArrayOutput) Index(i pulumi.IntInput) HbaseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HbaseInstance {
		return vs[0].([]*HbaseInstance)[vs[1].(int)]
	}).(HbaseInstanceOutput)
}

type HbaseInstanceMapOutput struct{ *pulumi.OutputState }

func (HbaseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HbaseInstance)(nil)).Elem()
}

func (o HbaseInstanceMapOutput) ToHbaseInstanceMapOutput() HbaseInstanceMapOutput {
	return o
}

func (o HbaseInstanceMapOutput) ToHbaseInstanceMapOutputWithContext(ctx context.Context) HbaseInstanceMapOutput {
	return o
}

func (o HbaseInstanceMapOutput) MapIndex(k pulumi.StringInput) HbaseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HbaseInstance {
		return vs[0].(map[string]*HbaseInstance)[vs[1].(string)]
	}).(HbaseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HbaseInstanceInput)(nil)).Elem(), &HbaseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*HbaseInstanceArrayInput)(nil)).Elem(), HbaseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HbaseInstanceMapInput)(nil)).Elem(), HbaseInstanceMap{})
	pulumi.RegisterOutputType(HbaseInstanceOutput{})
	pulumi.RegisterOutputType(HbaseInstanceArrayOutput{})
	pulumi.RegisterOutputType(HbaseInstanceMapOutput{})
}