// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Parameters struct {
	pulumi.CustomResourceState

	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Number of days to keep, no more than 30.
	Params ParametersParamArrayOutput `pulumi:"params"`
}

// NewParameters registers a new resource with the given unique name, arguments, and options.
func NewParameters(ctx *pulumi.Context,
	name string, args *ParametersArgs, opts ...pulumi.ResourceOption) (*Parameters, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Params == nil {
		return nil, errors.New("invalid value for required argument 'Params'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Parameters
	err := ctx.RegisterResource("tencentcloud:Mariadb/parameters:Parameters", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameters gets an existing Parameters resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameters(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParametersState, opts ...pulumi.ResourceOption) (*Parameters, error) {
	var resource Parameters
	err := ctx.ReadResource("tencentcloud:Mariadb/parameters:Parameters", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parameters resources.
type parametersState struct {
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Number of days to keep, no more than 30.
	Params []ParametersParam `pulumi:"params"`
}

type ParametersState struct {
	// instance id.
	InstanceId pulumi.StringPtrInput
	// Number of days to keep, no more than 30.
	Params ParametersParamArrayInput
}

func (ParametersState) ElementType() reflect.Type {
	return reflect.TypeOf((*parametersState)(nil)).Elem()
}

type parametersArgs struct {
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Number of days to keep, no more than 30.
	Params []ParametersParam `pulumi:"params"`
}

// The set of arguments for constructing a Parameters resource.
type ParametersArgs struct {
	// instance id.
	InstanceId pulumi.StringInput
	// Number of days to keep, no more than 30.
	Params ParametersParamArrayInput
}

func (ParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parametersArgs)(nil)).Elem()
}

type ParametersInput interface {
	pulumi.Input

	ToParametersOutput() ParametersOutput
	ToParametersOutputWithContext(ctx context.Context) ParametersOutput
}

func (*Parameters) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameters)(nil)).Elem()
}

func (i *Parameters) ToParametersOutput() ParametersOutput {
	return i.ToParametersOutputWithContext(context.Background())
}

func (i *Parameters) ToParametersOutputWithContext(ctx context.Context) ParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersOutput)
}

// ParametersArrayInput is an input type that accepts ParametersArray and ParametersArrayOutput values.
// You can construct a concrete instance of `ParametersArrayInput` via:
//
//	ParametersArray{ ParametersArgs{...} }
type ParametersArrayInput interface {
	pulumi.Input

	ToParametersArrayOutput() ParametersArrayOutput
	ToParametersArrayOutputWithContext(context.Context) ParametersArrayOutput
}

type ParametersArray []ParametersInput

func (ParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameters)(nil)).Elem()
}

func (i ParametersArray) ToParametersArrayOutput() ParametersArrayOutput {
	return i.ToParametersArrayOutputWithContext(context.Background())
}

func (i ParametersArray) ToParametersArrayOutputWithContext(ctx context.Context) ParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersArrayOutput)
}

// ParametersMapInput is an input type that accepts ParametersMap and ParametersMapOutput values.
// You can construct a concrete instance of `ParametersMapInput` via:
//
//	ParametersMap{ "key": ParametersArgs{...} }
type ParametersMapInput interface {
	pulumi.Input

	ToParametersMapOutput() ParametersMapOutput
	ToParametersMapOutputWithContext(context.Context) ParametersMapOutput
}

type ParametersMap map[string]ParametersInput

func (ParametersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameters)(nil)).Elem()
}

func (i ParametersMap) ToParametersMapOutput() ParametersMapOutput {
	return i.ToParametersMapOutputWithContext(context.Background())
}

func (i ParametersMap) ToParametersMapOutputWithContext(ctx context.Context) ParametersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersMapOutput)
}

type ParametersOutput struct{ *pulumi.OutputState }

func (ParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameters)(nil)).Elem()
}

func (o ParametersOutput) ToParametersOutput() ParametersOutput {
	return o
}

func (o ParametersOutput) ToParametersOutputWithContext(ctx context.Context) ParametersOutput {
	return o
}

// instance id.
func (o ParametersOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameters) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Number of days to keep, no more than 30.
func (o ParametersOutput) Params() ParametersParamArrayOutput {
	return o.ApplyT(func(v *Parameters) ParametersParamArrayOutput { return v.Params }).(ParametersParamArrayOutput)
}

type ParametersArrayOutput struct{ *pulumi.OutputState }

func (ParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameters)(nil)).Elem()
}

func (o ParametersArrayOutput) ToParametersArrayOutput() ParametersArrayOutput {
	return o
}

func (o ParametersArrayOutput) ToParametersArrayOutputWithContext(ctx context.Context) ParametersArrayOutput {
	return o
}

func (o ParametersArrayOutput) Index(i pulumi.IntInput) ParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Parameters {
		return vs[0].([]*Parameters)[vs[1].(int)]
	}).(ParametersOutput)
}

type ParametersMapOutput struct{ *pulumi.OutputState }

func (ParametersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameters)(nil)).Elem()
}

func (o ParametersMapOutput) ToParametersMapOutput() ParametersMapOutput {
	return o
}

func (o ParametersMapOutput) ToParametersMapOutputWithContext(ctx context.Context) ParametersMapOutput {
	return o
}

func (o ParametersMapOutput) MapIndex(k pulumi.StringInput) ParametersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Parameters {
		return vs[0].(map[string]*Parameters)[vs[1].(string)]
	}).(ParametersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersInput)(nil)).Elem(), &Parameters{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersArrayInput)(nil)).Elem(), ParametersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersMapInput)(nil)).Elem(), ParametersMap{})
	pulumi.RegisterOutputType(ParametersOutput{})
	pulumi.RegisterOutputType(ParametersArrayOutput{})
	pulumi.RegisterOutputType(ParametersMapOutput{})
}