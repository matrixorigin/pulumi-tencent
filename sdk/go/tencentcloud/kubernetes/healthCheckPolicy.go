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

type HealthCheckPolicy struct {
	pulumi.CustomResourceState

	// ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Health Check Policy Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Health check policy rule list.
	Rules HealthCheckPolicyRuleArrayOutput `pulumi:"rules"`
}

// NewHealthCheckPolicy registers a new resource with the given unique name, arguments, and options.
func NewHealthCheckPolicy(ctx *pulumi.Context,
	name string, args *HealthCheckPolicyArgs, opts ...pulumi.ResourceOption) (*HealthCheckPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HealthCheckPolicy
	err := ctx.RegisterResource("tencentcloud:Kubernetes/healthCheckPolicy:HealthCheckPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthCheckPolicy gets an existing HealthCheckPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthCheckPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthCheckPolicyState, opts ...pulumi.ResourceOption) (*HealthCheckPolicy, error) {
	var resource HealthCheckPolicy
	err := ctx.ReadResource("tencentcloud:Kubernetes/healthCheckPolicy:HealthCheckPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthCheckPolicy resources.
type healthCheckPolicyState struct {
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Health Check Policy Name.
	Name *string `pulumi:"name"`
	// Health check policy rule list.
	Rules []HealthCheckPolicyRule `pulumi:"rules"`
}

type HealthCheckPolicyState struct {
	// ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// Health Check Policy Name.
	Name pulumi.StringPtrInput
	// Health check policy rule list.
	Rules HealthCheckPolicyRuleArrayInput
}

func (HealthCheckPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckPolicyState)(nil)).Elem()
}

type healthCheckPolicyArgs struct {
	// ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// Health Check Policy Name.
	Name *string `pulumi:"name"`
	// Health check policy rule list.
	Rules []HealthCheckPolicyRule `pulumi:"rules"`
}

// The set of arguments for constructing a HealthCheckPolicy resource.
type HealthCheckPolicyArgs struct {
	// ID of the cluster.
	ClusterId pulumi.StringInput
	// Health Check Policy Name.
	Name pulumi.StringPtrInput
	// Health check policy rule list.
	Rules HealthCheckPolicyRuleArrayInput
}

func (HealthCheckPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckPolicyArgs)(nil)).Elem()
}

type HealthCheckPolicyInput interface {
	pulumi.Input

	ToHealthCheckPolicyOutput() HealthCheckPolicyOutput
	ToHealthCheckPolicyOutputWithContext(ctx context.Context) HealthCheckPolicyOutput
}

func (*HealthCheckPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheckPolicy)(nil)).Elem()
}

func (i *HealthCheckPolicy) ToHealthCheckPolicyOutput() HealthCheckPolicyOutput {
	return i.ToHealthCheckPolicyOutputWithContext(context.Background())
}

func (i *HealthCheckPolicy) ToHealthCheckPolicyOutputWithContext(ctx context.Context) HealthCheckPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckPolicyOutput)
}

// HealthCheckPolicyArrayInput is an input type that accepts HealthCheckPolicyArray and HealthCheckPolicyArrayOutput values.
// You can construct a concrete instance of `HealthCheckPolicyArrayInput` via:
//
//	HealthCheckPolicyArray{ HealthCheckPolicyArgs{...} }
type HealthCheckPolicyArrayInput interface {
	pulumi.Input

	ToHealthCheckPolicyArrayOutput() HealthCheckPolicyArrayOutput
	ToHealthCheckPolicyArrayOutputWithContext(context.Context) HealthCheckPolicyArrayOutput
}

type HealthCheckPolicyArray []HealthCheckPolicyInput

func (HealthCheckPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HealthCheckPolicy)(nil)).Elem()
}

func (i HealthCheckPolicyArray) ToHealthCheckPolicyArrayOutput() HealthCheckPolicyArrayOutput {
	return i.ToHealthCheckPolicyArrayOutputWithContext(context.Background())
}

func (i HealthCheckPolicyArray) ToHealthCheckPolicyArrayOutputWithContext(ctx context.Context) HealthCheckPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckPolicyArrayOutput)
}

// HealthCheckPolicyMapInput is an input type that accepts HealthCheckPolicyMap and HealthCheckPolicyMapOutput values.
// You can construct a concrete instance of `HealthCheckPolicyMapInput` via:
//
//	HealthCheckPolicyMap{ "key": HealthCheckPolicyArgs{...} }
type HealthCheckPolicyMapInput interface {
	pulumi.Input

	ToHealthCheckPolicyMapOutput() HealthCheckPolicyMapOutput
	ToHealthCheckPolicyMapOutputWithContext(context.Context) HealthCheckPolicyMapOutput
}

type HealthCheckPolicyMap map[string]HealthCheckPolicyInput

func (HealthCheckPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HealthCheckPolicy)(nil)).Elem()
}

func (i HealthCheckPolicyMap) ToHealthCheckPolicyMapOutput() HealthCheckPolicyMapOutput {
	return i.ToHealthCheckPolicyMapOutputWithContext(context.Background())
}

func (i HealthCheckPolicyMap) ToHealthCheckPolicyMapOutputWithContext(ctx context.Context) HealthCheckPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckPolicyMapOutput)
}

type HealthCheckPolicyOutput struct{ *pulumi.OutputState }

func (HealthCheckPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheckPolicy)(nil)).Elem()
}

func (o HealthCheckPolicyOutput) ToHealthCheckPolicyOutput() HealthCheckPolicyOutput {
	return o
}

func (o HealthCheckPolicyOutput) ToHealthCheckPolicyOutputWithContext(ctx context.Context) HealthCheckPolicyOutput {
	return o
}

// ID of the cluster.
func (o HealthCheckPolicyOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthCheckPolicy) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Health Check Policy Name.
func (o HealthCheckPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthCheckPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Health check policy rule list.
func (o HealthCheckPolicyOutput) Rules() HealthCheckPolicyRuleArrayOutput {
	return o.ApplyT(func(v *HealthCheckPolicy) HealthCheckPolicyRuleArrayOutput { return v.Rules }).(HealthCheckPolicyRuleArrayOutput)
}

type HealthCheckPolicyArrayOutput struct{ *pulumi.OutputState }

func (HealthCheckPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HealthCheckPolicy)(nil)).Elem()
}

func (o HealthCheckPolicyArrayOutput) ToHealthCheckPolicyArrayOutput() HealthCheckPolicyArrayOutput {
	return o
}

func (o HealthCheckPolicyArrayOutput) ToHealthCheckPolicyArrayOutputWithContext(ctx context.Context) HealthCheckPolicyArrayOutput {
	return o
}

func (o HealthCheckPolicyArrayOutput) Index(i pulumi.IntInput) HealthCheckPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HealthCheckPolicy {
		return vs[0].([]*HealthCheckPolicy)[vs[1].(int)]
	}).(HealthCheckPolicyOutput)
}

type HealthCheckPolicyMapOutput struct{ *pulumi.OutputState }

func (HealthCheckPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HealthCheckPolicy)(nil)).Elem()
}

func (o HealthCheckPolicyMapOutput) ToHealthCheckPolicyMapOutput() HealthCheckPolicyMapOutput {
	return o
}

func (o HealthCheckPolicyMapOutput) ToHealthCheckPolicyMapOutputWithContext(ctx context.Context) HealthCheckPolicyMapOutput {
	return o
}

func (o HealthCheckPolicyMapOutput) MapIndex(k pulumi.StringInput) HealthCheckPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HealthCheckPolicy {
		return vs[0].(map[string]*HealthCheckPolicy)[vs[1].(string)]
	}).(HealthCheckPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckPolicyInput)(nil)).Elem(), &HealthCheckPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckPolicyArrayInput)(nil)).Elem(), HealthCheckPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckPolicyMapInput)(nil)).Elem(), HealthCheckPolicyMap{})
	pulumi.RegisterOutputType(HealthCheckPolicyOutput{})
	pulumi.RegisterOutputType(HealthCheckPolicyArrayOutput{})
	pulumi.RegisterOutputType(HealthCheckPolicyMapOutput{})
}