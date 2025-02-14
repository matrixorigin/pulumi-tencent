// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CenterRoleConfiguration struct {
	pulumi.CustomResourceState

	// Create time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Access configuration description, which contains up to 1024 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. This page must be the Tencent Cloud console page. The default is null,
	// which indicates navigating to the home page of the Tencent Cloud console.
	RelayState pulumi.StringOutput `pulumi:"relayState"`
	// Role configuration id.
	RoleConfigurationId pulumi.StringOutput `pulumi:"roleConfigurationId"`
	// Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).
	RoleConfigurationName pulumi.StringOutput `pulumi:"roleConfigurationName"`
	// Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours).
	// Default value: 3600 (1 hour).
	SessionDuration pulumi.IntOutput `pulumi:"sessionDuration"`
	// Update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Space ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCenterRoleConfiguration registers a new resource with the given unique name, arguments, and options.
func NewCenterRoleConfiguration(ctx *pulumi.Context,
	name string, args *CenterRoleConfigurationArgs, opts ...pulumi.ResourceOption) (*CenterRoleConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'RoleConfigurationName'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CenterRoleConfiguration
	err := ctx.RegisterResource("tencentcloud:Identity/centerRoleConfiguration:CenterRoleConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCenterRoleConfiguration gets an existing CenterRoleConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCenterRoleConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CenterRoleConfigurationState, opts ...pulumi.ResourceOption) (*CenterRoleConfiguration, error) {
	var resource CenterRoleConfiguration
	err := ctx.ReadResource("tencentcloud:Identity/centerRoleConfiguration:CenterRoleConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CenterRoleConfiguration resources.
type centerRoleConfigurationState struct {
	// Create time.
	CreateTime *string `pulumi:"createTime"`
	// Access configuration description, which contains up to 1024 characters.
	Description *string `pulumi:"description"`
	// Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. This page must be the Tencent Cloud console page. The default is null,
	// which indicates navigating to the home page of the Tencent Cloud console.
	RelayState *string `pulumi:"relayState"`
	// Role configuration id.
	RoleConfigurationId *string `pulumi:"roleConfigurationId"`
	// Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).
	RoleConfigurationName *string `pulumi:"roleConfigurationName"`
	// Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours).
	// Default value: 3600 (1 hour).
	SessionDuration *int `pulumi:"sessionDuration"`
	// Update time.
	UpdateTime *string `pulumi:"updateTime"`
	// Space ID.
	ZoneId *string `pulumi:"zoneId"`
}

type CenterRoleConfigurationState struct {
	// Create time.
	CreateTime pulumi.StringPtrInput
	// Access configuration description, which contains up to 1024 characters.
	Description pulumi.StringPtrInput
	// Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. This page must be the Tencent Cloud console page. The default is null,
	// which indicates navigating to the home page of the Tencent Cloud console.
	RelayState pulumi.StringPtrInput
	// Role configuration id.
	RoleConfigurationId pulumi.StringPtrInput
	// Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).
	RoleConfigurationName pulumi.StringPtrInput
	// Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours).
	// Default value: 3600 (1 hour).
	SessionDuration pulumi.IntPtrInput
	// Update time.
	UpdateTime pulumi.StringPtrInput
	// Space ID.
	ZoneId pulumi.StringPtrInput
}

func (CenterRoleConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*centerRoleConfigurationState)(nil)).Elem()
}

type centerRoleConfigurationArgs struct {
	// Access configuration description, which contains up to 1024 characters.
	Description *string `pulumi:"description"`
	// Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. This page must be the Tencent Cloud console page. The default is null,
	// which indicates navigating to the home page of the Tencent Cloud console.
	RelayState *string `pulumi:"relayState"`
	// Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).
	RoleConfigurationName string `pulumi:"roleConfigurationName"`
	// Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours).
	// Default value: 3600 (1 hour).
	SessionDuration *int `pulumi:"sessionDuration"`
	// Space ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a CenterRoleConfiguration resource.
type CenterRoleConfigurationArgs struct {
	// Access configuration description, which contains up to 1024 characters.
	Description pulumi.StringPtrInput
	// Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. This page must be the Tencent Cloud console page. The default is null,
	// which indicates navigating to the home page of the Tencent Cloud console.
	RelayState pulumi.StringPtrInput
	// Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).
	RoleConfigurationName pulumi.StringInput
	// Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the
	// target account of the Tencent Cloud Organization. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours).
	// Default value: 3600 (1 hour).
	SessionDuration pulumi.IntPtrInput
	// Space ID.
	ZoneId pulumi.StringInput
}

func (CenterRoleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*centerRoleConfigurationArgs)(nil)).Elem()
}

type CenterRoleConfigurationInput interface {
	pulumi.Input

	ToCenterRoleConfigurationOutput() CenterRoleConfigurationOutput
	ToCenterRoleConfigurationOutputWithContext(ctx context.Context) CenterRoleConfigurationOutput
}

func (*CenterRoleConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**CenterRoleConfiguration)(nil)).Elem()
}

func (i *CenterRoleConfiguration) ToCenterRoleConfigurationOutput() CenterRoleConfigurationOutput {
	return i.ToCenterRoleConfigurationOutputWithContext(context.Background())
}

func (i *CenterRoleConfiguration) ToCenterRoleConfigurationOutputWithContext(ctx context.Context) CenterRoleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CenterRoleConfigurationOutput)
}

// CenterRoleConfigurationArrayInput is an input type that accepts CenterRoleConfigurationArray and CenterRoleConfigurationArrayOutput values.
// You can construct a concrete instance of `CenterRoleConfigurationArrayInput` via:
//
//	CenterRoleConfigurationArray{ CenterRoleConfigurationArgs{...} }
type CenterRoleConfigurationArrayInput interface {
	pulumi.Input

	ToCenterRoleConfigurationArrayOutput() CenterRoleConfigurationArrayOutput
	ToCenterRoleConfigurationArrayOutputWithContext(context.Context) CenterRoleConfigurationArrayOutput
}

type CenterRoleConfigurationArray []CenterRoleConfigurationInput

func (CenterRoleConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CenterRoleConfiguration)(nil)).Elem()
}

func (i CenterRoleConfigurationArray) ToCenterRoleConfigurationArrayOutput() CenterRoleConfigurationArrayOutput {
	return i.ToCenterRoleConfigurationArrayOutputWithContext(context.Background())
}

func (i CenterRoleConfigurationArray) ToCenterRoleConfigurationArrayOutputWithContext(ctx context.Context) CenterRoleConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CenterRoleConfigurationArrayOutput)
}

// CenterRoleConfigurationMapInput is an input type that accepts CenterRoleConfigurationMap and CenterRoleConfigurationMapOutput values.
// You can construct a concrete instance of `CenterRoleConfigurationMapInput` via:
//
//	CenterRoleConfigurationMap{ "key": CenterRoleConfigurationArgs{...} }
type CenterRoleConfigurationMapInput interface {
	pulumi.Input

	ToCenterRoleConfigurationMapOutput() CenterRoleConfigurationMapOutput
	ToCenterRoleConfigurationMapOutputWithContext(context.Context) CenterRoleConfigurationMapOutput
}

type CenterRoleConfigurationMap map[string]CenterRoleConfigurationInput

func (CenterRoleConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CenterRoleConfiguration)(nil)).Elem()
}

func (i CenterRoleConfigurationMap) ToCenterRoleConfigurationMapOutput() CenterRoleConfigurationMapOutput {
	return i.ToCenterRoleConfigurationMapOutputWithContext(context.Background())
}

func (i CenterRoleConfigurationMap) ToCenterRoleConfigurationMapOutputWithContext(ctx context.Context) CenterRoleConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CenterRoleConfigurationMapOutput)
}

type CenterRoleConfigurationOutput struct{ *pulumi.OutputState }

func (CenterRoleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CenterRoleConfiguration)(nil)).Elem()
}

func (o CenterRoleConfigurationOutput) ToCenterRoleConfigurationOutput() CenterRoleConfigurationOutput {
	return o
}

func (o CenterRoleConfigurationOutput) ToCenterRoleConfigurationOutputWithContext(ctx context.Context) CenterRoleConfigurationOutput {
	return o
}

// Create time.
func (o CenterRoleConfigurationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Access configuration description, which contains up to 1024 characters.
func (o CenterRoleConfigurationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the
// target account of the Tencent Cloud Organization. This page must be the Tencent Cloud console page. The default is null,
// which indicates navigating to the home page of the Tencent Cloud console.
func (o CenterRoleConfigurationOutput) RelayState() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.RelayState }).(pulumi.StringOutput)
}

// Role configuration id.
func (o CenterRoleConfigurationOutput) RoleConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.RoleConfigurationId }).(pulumi.StringOutput)
}

// Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).
func (o CenterRoleConfigurationOutput) RoleConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.RoleConfigurationName }).(pulumi.StringOutput)
}

// Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the
// target account of the Tencent Cloud Organization. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours).
// Default value: 3600 (1 hour).
func (o CenterRoleConfigurationOutput) SessionDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.IntOutput { return v.SessionDuration }).(pulumi.IntOutput)
}

// Update time.
func (o CenterRoleConfigurationOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Space ID.
func (o CenterRoleConfigurationOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *CenterRoleConfiguration) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type CenterRoleConfigurationArrayOutput struct{ *pulumi.OutputState }

func (CenterRoleConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CenterRoleConfiguration)(nil)).Elem()
}

func (o CenterRoleConfigurationArrayOutput) ToCenterRoleConfigurationArrayOutput() CenterRoleConfigurationArrayOutput {
	return o
}

func (o CenterRoleConfigurationArrayOutput) ToCenterRoleConfigurationArrayOutputWithContext(ctx context.Context) CenterRoleConfigurationArrayOutput {
	return o
}

func (o CenterRoleConfigurationArrayOutput) Index(i pulumi.IntInput) CenterRoleConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CenterRoleConfiguration {
		return vs[0].([]*CenterRoleConfiguration)[vs[1].(int)]
	}).(CenterRoleConfigurationOutput)
}

type CenterRoleConfigurationMapOutput struct{ *pulumi.OutputState }

func (CenterRoleConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CenterRoleConfiguration)(nil)).Elem()
}

func (o CenterRoleConfigurationMapOutput) ToCenterRoleConfigurationMapOutput() CenterRoleConfigurationMapOutput {
	return o
}

func (o CenterRoleConfigurationMapOutput) ToCenterRoleConfigurationMapOutputWithContext(ctx context.Context) CenterRoleConfigurationMapOutput {
	return o
}

func (o CenterRoleConfigurationMapOutput) MapIndex(k pulumi.StringInput) CenterRoleConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CenterRoleConfiguration {
		return vs[0].(map[string]*CenterRoleConfiguration)[vs[1].(string)]
	}).(CenterRoleConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CenterRoleConfigurationInput)(nil)).Elem(), &CenterRoleConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*CenterRoleConfigurationArrayInput)(nil)).Elem(), CenterRoleConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CenterRoleConfigurationMapInput)(nil)).Elem(), CenterRoleConfigurationMap{})
	pulumi.RegisterOutputType(CenterRoleConfigurationOutput{})
	pulumi.RegisterOutputType(CenterRoleConfigurationArrayOutput{})
	pulumi.RegisterOutputType(CenterRoleConfigurationMapOutput{})
}
