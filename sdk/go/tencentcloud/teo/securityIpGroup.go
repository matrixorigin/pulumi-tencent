// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityIpGroup struct {
	pulumi.CustomResourceState

	// IP group information, replace all when modifying.
	IpGroup SecurityIpGroupIpGroupOutput `pulumi:"ipGroup"`
	// Site ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewSecurityIpGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityIpGroup(ctx *pulumi.Context,
	name string, args *SecurityIpGroupArgs, opts ...pulumi.ResourceOption) (*SecurityIpGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpGroup == nil {
		return nil, errors.New("invalid value for required argument 'IpGroup'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityIpGroup
	err := ctx.RegisterResource("tencentcloud:Teo/securityIpGroup:SecurityIpGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityIpGroup gets an existing SecurityIpGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityIpGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityIpGroupState, opts ...pulumi.ResourceOption) (*SecurityIpGroup, error) {
	var resource SecurityIpGroup
	err := ctx.ReadResource("tencentcloud:Teo/securityIpGroup:SecurityIpGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityIpGroup resources.
type securityIpGroupState struct {
	// IP group information, replace all when modifying.
	IpGroup *SecurityIpGroupIpGroup `pulumi:"ipGroup"`
	// Site ID.
	ZoneId *string `pulumi:"zoneId"`
}

type SecurityIpGroupState struct {
	// IP group information, replace all when modifying.
	IpGroup SecurityIpGroupIpGroupPtrInput
	// Site ID.
	ZoneId pulumi.StringPtrInput
}

func (SecurityIpGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityIpGroupState)(nil)).Elem()
}

type securityIpGroupArgs struct {
	// IP group information, replace all when modifying.
	IpGroup SecurityIpGroupIpGroup `pulumi:"ipGroup"`
	// Site ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a SecurityIpGroup resource.
type SecurityIpGroupArgs struct {
	// IP group information, replace all when modifying.
	IpGroup SecurityIpGroupIpGroupInput
	// Site ID.
	ZoneId pulumi.StringInput
}

func (SecurityIpGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityIpGroupArgs)(nil)).Elem()
}

type SecurityIpGroupInput interface {
	pulumi.Input

	ToSecurityIpGroupOutput() SecurityIpGroupOutput
	ToSecurityIpGroupOutputWithContext(ctx context.Context) SecurityIpGroupOutput
}

func (*SecurityIpGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityIpGroup)(nil)).Elem()
}

func (i *SecurityIpGroup) ToSecurityIpGroupOutput() SecurityIpGroupOutput {
	return i.ToSecurityIpGroupOutputWithContext(context.Background())
}

func (i *SecurityIpGroup) ToSecurityIpGroupOutputWithContext(ctx context.Context) SecurityIpGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityIpGroupOutput)
}

// SecurityIpGroupArrayInput is an input type that accepts SecurityIpGroupArray and SecurityIpGroupArrayOutput values.
// You can construct a concrete instance of `SecurityIpGroupArrayInput` via:
//
//	SecurityIpGroupArray{ SecurityIpGroupArgs{...} }
type SecurityIpGroupArrayInput interface {
	pulumi.Input

	ToSecurityIpGroupArrayOutput() SecurityIpGroupArrayOutput
	ToSecurityIpGroupArrayOutputWithContext(context.Context) SecurityIpGroupArrayOutput
}

type SecurityIpGroupArray []SecurityIpGroupInput

func (SecurityIpGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityIpGroup)(nil)).Elem()
}

func (i SecurityIpGroupArray) ToSecurityIpGroupArrayOutput() SecurityIpGroupArrayOutput {
	return i.ToSecurityIpGroupArrayOutputWithContext(context.Background())
}

func (i SecurityIpGroupArray) ToSecurityIpGroupArrayOutputWithContext(ctx context.Context) SecurityIpGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityIpGroupArrayOutput)
}

// SecurityIpGroupMapInput is an input type that accepts SecurityIpGroupMap and SecurityIpGroupMapOutput values.
// You can construct a concrete instance of `SecurityIpGroupMapInput` via:
//
//	SecurityIpGroupMap{ "key": SecurityIpGroupArgs{...} }
type SecurityIpGroupMapInput interface {
	pulumi.Input

	ToSecurityIpGroupMapOutput() SecurityIpGroupMapOutput
	ToSecurityIpGroupMapOutputWithContext(context.Context) SecurityIpGroupMapOutput
}

type SecurityIpGroupMap map[string]SecurityIpGroupInput

func (SecurityIpGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityIpGroup)(nil)).Elem()
}

func (i SecurityIpGroupMap) ToSecurityIpGroupMapOutput() SecurityIpGroupMapOutput {
	return i.ToSecurityIpGroupMapOutputWithContext(context.Background())
}

func (i SecurityIpGroupMap) ToSecurityIpGroupMapOutputWithContext(ctx context.Context) SecurityIpGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityIpGroupMapOutput)
}

type SecurityIpGroupOutput struct{ *pulumi.OutputState }

func (SecurityIpGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityIpGroup)(nil)).Elem()
}

func (o SecurityIpGroupOutput) ToSecurityIpGroupOutput() SecurityIpGroupOutput {
	return o
}

func (o SecurityIpGroupOutput) ToSecurityIpGroupOutputWithContext(ctx context.Context) SecurityIpGroupOutput {
	return o
}

// IP group information, replace all when modifying.
func (o SecurityIpGroupOutput) IpGroup() SecurityIpGroupIpGroupOutput {
	return o.ApplyT(func(v *SecurityIpGroup) SecurityIpGroupIpGroupOutput { return v.IpGroup }).(SecurityIpGroupIpGroupOutput)
}

// Site ID.
func (o SecurityIpGroupOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityIpGroup) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type SecurityIpGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityIpGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityIpGroup)(nil)).Elem()
}

func (o SecurityIpGroupArrayOutput) ToSecurityIpGroupArrayOutput() SecurityIpGroupArrayOutput {
	return o
}

func (o SecurityIpGroupArrayOutput) ToSecurityIpGroupArrayOutputWithContext(ctx context.Context) SecurityIpGroupArrayOutput {
	return o
}

func (o SecurityIpGroupArrayOutput) Index(i pulumi.IntInput) SecurityIpGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityIpGroup {
		return vs[0].([]*SecurityIpGroup)[vs[1].(int)]
	}).(SecurityIpGroupOutput)
}

type SecurityIpGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityIpGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityIpGroup)(nil)).Elem()
}

func (o SecurityIpGroupMapOutput) ToSecurityIpGroupMapOutput() SecurityIpGroupMapOutput {
	return o
}

func (o SecurityIpGroupMapOutput) ToSecurityIpGroupMapOutputWithContext(ctx context.Context) SecurityIpGroupMapOutput {
	return o
}

func (o SecurityIpGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityIpGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityIpGroup {
		return vs[0].(map[string]*SecurityIpGroup)[vs[1].(string)]
	}).(SecurityIpGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityIpGroupInput)(nil)).Elem(), &SecurityIpGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityIpGroupArrayInput)(nil)).Elem(), SecurityIpGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityIpGroupMapInput)(nil)).Elem(), SecurityIpGroupMap{})
	pulumi.RegisterOutputType(SecurityIpGroupOutput{})
	pulumi.RegisterOutputType(SecurityIpGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityIpGroupMapOutput{})
}