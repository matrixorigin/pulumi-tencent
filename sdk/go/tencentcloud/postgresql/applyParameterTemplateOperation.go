// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplyParameterTemplateOperation struct {
	pulumi.CustomResourceState

	// Instance ID.
	DbInstanceId pulumi.StringOutput `pulumi:"dbInstanceId"`
	// Template ID.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
}

// NewApplyParameterTemplateOperation registers a new resource with the given unique name, arguments, and options.
func NewApplyParameterTemplateOperation(ctx *pulumi.Context,
	name string, args *ApplyParameterTemplateOperationArgs, opts ...pulumi.ResourceOption) (*ApplyParameterTemplateOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceId'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplyParameterTemplateOperation
	err := ctx.RegisterResource("tencentcloud:Postgresql/applyParameterTemplateOperation:ApplyParameterTemplateOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplyParameterTemplateOperation gets an existing ApplyParameterTemplateOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplyParameterTemplateOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplyParameterTemplateOperationState, opts ...pulumi.ResourceOption) (*ApplyParameterTemplateOperation, error) {
	var resource ApplyParameterTemplateOperation
	err := ctx.ReadResource("tencentcloud:Postgresql/applyParameterTemplateOperation:ApplyParameterTemplateOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplyParameterTemplateOperation resources.
type applyParameterTemplateOperationState struct {
	// Instance ID.
	DbInstanceId *string `pulumi:"dbInstanceId"`
	// Template ID.
	TemplateId *string `pulumi:"templateId"`
}

type ApplyParameterTemplateOperationState struct {
	// Instance ID.
	DbInstanceId pulumi.StringPtrInput
	// Template ID.
	TemplateId pulumi.StringPtrInput
}

func (ApplyParameterTemplateOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applyParameterTemplateOperationState)(nil)).Elem()
}

type applyParameterTemplateOperationArgs struct {
	// Instance ID.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// Template ID.
	TemplateId string `pulumi:"templateId"`
}

// The set of arguments for constructing a ApplyParameterTemplateOperation resource.
type ApplyParameterTemplateOperationArgs struct {
	// Instance ID.
	DbInstanceId pulumi.StringInput
	// Template ID.
	TemplateId pulumi.StringInput
}

func (ApplyParameterTemplateOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applyParameterTemplateOperationArgs)(nil)).Elem()
}

type ApplyParameterTemplateOperationInput interface {
	pulumi.Input

	ToApplyParameterTemplateOperationOutput() ApplyParameterTemplateOperationOutput
	ToApplyParameterTemplateOperationOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationOutput
}

func (*ApplyParameterTemplateOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplyParameterTemplateOperation)(nil)).Elem()
}

func (i *ApplyParameterTemplateOperation) ToApplyParameterTemplateOperationOutput() ApplyParameterTemplateOperationOutput {
	return i.ToApplyParameterTemplateOperationOutputWithContext(context.Background())
}

func (i *ApplyParameterTemplateOperation) ToApplyParameterTemplateOperationOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplyParameterTemplateOperationOutput)
}

// ApplyParameterTemplateOperationArrayInput is an input type that accepts ApplyParameterTemplateOperationArray and ApplyParameterTemplateOperationArrayOutput values.
// You can construct a concrete instance of `ApplyParameterTemplateOperationArrayInput` via:
//
//	ApplyParameterTemplateOperationArray{ ApplyParameterTemplateOperationArgs{...} }
type ApplyParameterTemplateOperationArrayInput interface {
	pulumi.Input

	ToApplyParameterTemplateOperationArrayOutput() ApplyParameterTemplateOperationArrayOutput
	ToApplyParameterTemplateOperationArrayOutputWithContext(context.Context) ApplyParameterTemplateOperationArrayOutput
}

type ApplyParameterTemplateOperationArray []ApplyParameterTemplateOperationInput

func (ApplyParameterTemplateOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplyParameterTemplateOperation)(nil)).Elem()
}

func (i ApplyParameterTemplateOperationArray) ToApplyParameterTemplateOperationArrayOutput() ApplyParameterTemplateOperationArrayOutput {
	return i.ToApplyParameterTemplateOperationArrayOutputWithContext(context.Background())
}

func (i ApplyParameterTemplateOperationArray) ToApplyParameterTemplateOperationArrayOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplyParameterTemplateOperationArrayOutput)
}

// ApplyParameterTemplateOperationMapInput is an input type that accepts ApplyParameterTemplateOperationMap and ApplyParameterTemplateOperationMapOutput values.
// You can construct a concrete instance of `ApplyParameterTemplateOperationMapInput` via:
//
//	ApplyParameterTemplateOperationMap{ "key": ApplyParameterTemplateOperationArgs{...} }
type ApplyParameterTemplateOperationMapInput interface {
	pulumi.Input

	ToApplyParameterTemplateOperationMapOutput() ApplyParameterTemplateOperationMapOutput
	ToApplyParameterTemplateOperationMapOutputWithContext(context.Context) ApplyParameterTemplateOperationMapOutput
}

type ApplyParameterTemplateOperationMap map[string]ApplyParameterTemplateOperationInput

func (ApplyParameterTemplateOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplyParameterTemplateOperation)(nil)).Elem()
}

func (i ApplyParameterTemplateOperationMap) ToApplyParameterTemplateOperationMapOutput() ApplyParameterTemplateOperationMapOutput {
	return i.ToApplyParameterTemplateOperationMapOutputWithContext(context.Background())
}

func (i ApplyParameterTemplateOperationMap) ToApplyParameterTemplateOperationMapOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplyParameterTemplateOperationMapOutput)
}

type ApplyParameterTemplateOperationOutput struct{ *pulumi.OutputState }

func (ApplyParameterTemplateOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplyParameterTemplateOperation)(nil)).Elem()
}

func (o ApplyParameterTemplateOperationOutput) ToApplyParameterTemplateOperationOutput() ApplyParameterTemplateOperationOutput {
	return o
}

func (o ApplyParameterTemplateOperationOutput) ToApplyParameterTemplateOperationOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationOutput {
	return o
}

// Instance ID.
func (o ApplyParameterTemplateOperationOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplyParameterTemplateOperation) pulumi.StringOutput { return v.DbInstanceId }).(pulumi.StringOutput)
}

// Template ID.
func (o ApplyParameterTemplateOperationOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplyParameterTemplateOperation) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

type ApplyParameterTemplateOperationArrayOutput struct{ *pulumi.OutputState }

func (ApplyParameterTemplateOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplyParameterTemplateOperation)(nil)).Elem()
}

func (o ApplyParameterTemplateOperationArrayOutput) ToApplyParameterTemplateOperationArrayOutput() ApplyParameterTemplateOperationArrayOutput {
	return o
}

func (o ApplyParameterTemplateOperationArrayOutput) ToApplyParameterTemplateOperationArrayOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationArrayOutput {
	return o
}

func (o ApplyParameterTemplateOperationArrayOutput) Index(i pulumi.IntInput) ApplyParameterTemplateOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplyParameterTemplateOperation {
		return vs[0].([]*ApplyParameterTemplateOperation)[vs[1].(int)]
	}).(ApplyParameterTemplateOperationOutput)
}

type ApplyParameterTemplateOperationMapOutput struct{ *pulumi.OutputState }

func (ApplyParameterTemplateOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplyParameterTemplateOperation)(nil)).Elem()
}

func (o ApplyParameterTemplateOperationMapOutput) ToApplyParameterTemplateOperationMapOutput() ApplyParameterTemplateOperationMapOutput {
	return o
}

func (o ApplyParameterTemplateOperationMapOutput) ToApplyParameterTemplateOperationMapOutputWithContext(ctx context.Context) ApplyParameterTemplateOperationMapOutput {
	return o
}

func (o ApplyParameterTemplateOperationMapOutput) MapIndex(k pulumi.StringInput) ApplyParameterTemplateOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplyParameterTemplateOperation {
		return vs[0].(map[string]*ApplyParameterTemplateOperation)[vs[1].(string)]
	}).(ApplyParameterTemplateOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplyParameterTemplateOperationInput)(nil)).Elem(), &ApplyParameterTemplateOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplyParameterTemplateOperationArrayInput)(nil)).Elem(), ApplyParameterTemplateOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplyParameterTemplateOperationMapInput)(nil)).Elem(), ApplyParameterTemplateOperationMap{})
	pulumi.RegisterOutputType(ApplyParameterTemplateOperationOutput{})
	pulumi.RegisterOutputType(ApplyParameterTemplateOperationArrayOutput{})
	pulumi.RegisterOutputType(ApplyParameterTemplateOperationMapOutput{})
}