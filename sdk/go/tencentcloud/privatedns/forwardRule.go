// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ForwardRule struct {
	pulumi.CustomResourceState

	// Endpoint ID.
	EndPointId pulumi.StringOutput `pulumi:"endPointId"`
	// Forwarding rule name.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
	RuleType pulumi.StringOutput `pulumi:"ruleType"`
	// Private domain ID, which can be viewed on the private domain list page.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewForwardRule registers a new resource with the given unique name, arguments, and options.
func NewForwardRule(ctx *pulumi.Context,
	name string, args *ForwardRuleArgs, opts ...pulumi.ResourceOption) (*ForwardRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndPointId == nil {
		return nil, errors.New("invalid value for required argument 'EndPointId'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	if args.RuleType == nil {
		return nil, errors.New("invalid value for required argument 'RuleType'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ForwardRule
	err := ctx.RegisterResource("tencentcloud:PrivateDns/forwardRule:ForwardRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForwardRule gets an existing ForwardRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardRuleState, opts ...pulumi.ResourceOption) (*ForwardRule, error) {
	var resource ForwardRule
	err := ctx.ReadResource("tencentcloud:PrivateDns/forwardRule:ForwardRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ForwardRule resources.
type forwardRuleState struct {
	// Endpoint ID.
	EndPointId *string `pulumi:"endPointId"`
	// Forwarding rule name.
	RuleName *string `pulumi:"ruleName"`
	// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
	RuleType *string `pulumi:"ruleType"`
	// Private domain ID, which can be viewed on the private domain list page.
	ZoneId *string `pulumi:"zoneId"`
}

type ForwardRuleState struct {
	// Endpoint ID.
	EndPointId pulumi.StringPtrInput
	// Forwarding rule name.
	RuleName pulumi.StringPtrInput
	// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
	RuleType pulumi.StringPtrInput
	// Private domain ID, which can be viewed on the private domain list page.
	ZoneId pulumi.StringPtrInput
}

func (ForwardRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardRuleState)(nil)).Elem()
}

type forwardRuleArgs struct {
	// Endpoint ID.
	EndPointId string `pulumi:"endPointId"`
	// Forwarding rule name.
	RuleName string `pulumi:"ruleName"`
	// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
	RuleType string `pulumi:"ruleType"`
	// Private domain ID, which can be viewed on the private domain list page.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ForwardRule resource.
type ForwardRuleArgs struct {
	// Endpoint ID.
	EndPointId pulumi.StringInput
	// Forwarding rule name.
	RuleName pulumi.StringInput
	// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
	RuleType pulumi.StringInput
	// Private domain ID, which can be viewed on the private domain list page.
	ZoneId pulumi.StringInput
}

func (ForwardRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardRuleArgs)(nil)).Elem()
}

type ForwardRuleInput interface {
	pulumi.Input

	ToForwardRuleOutput() ForwardRuleOutput
	ToForwardRuleOutputWithContext(ctx context.Context) ForwardRuleOutput
}

func (*ForwardRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardRule)(nil)).Elem()
}

func (i *ForwardRule) ToForwardRuleOutput() ForwardRuleOutput {
	return i.ToForwardRuleOutputWithContext(context.Background())
}

func (i *ForwardRule) ToForwardRuleOutputWithContext(ctx context.Context) ForwardRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardRuleOutput)
}

// ForwardRuleArrayInput is an input type that accepts ForwardRuleArray and ForwardRuleArrayOutput values.
// You can construct a concrete instance of `ForwardRuleArrayInput` via:
//
//	ForwardRuleArray{ ForwardRuleArgs{...} }
type ForwardRuleArrayInput interface {
	pulumi.Input

	ToForwardRuleArrayOutput() ForwardRuleArrayOutput
	ToForwardRuleArrayOutputWithContext(context.Context) ForwardRuleArrayOutput
}

type ForwardRuleArray []ForwardRuleInput

func (ForwardRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ForwardRule)(nil)).Elem()
}

func (i ForwardRuleArray) ToForwardRuleArrayOutput() ForwardRuleArrayOutput {
	return i.ToForwardRuleArrayOutputWithContext(context.Background())
}

func (i ForwardRuleArray) ToForwardRuleArrayOutputWithContext(ctx context.Context) ForwardRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardRuleArrayOutput)
}

// ForwardRuleMapInput is an input type that accepts ForwardRuleMap and ForwardRuleMapOutput values.
// You can construct a concrete instance of `ForwardRuleMapInput` via:
//
//	ForwardRuleMap{ "key": ForwardRuleArgs{...} }
type ForwardRuleMapInput interface {
	pulumi.Input

	ToForwardRuleMapOutput() ForwardRuleMapOutput
	ToForwardRuleMapOutputWithContext(context.Context) ForwardRuleMapOutput
}

type ForwardRuleMap map[string]ForwardRuleInput

func (ForwardRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ForwardRule)(nil)).Elem()
}

func (i ForwardRuleMap) ToForwardRuleMapOutput() ForwardRuleMapOutput {
	return i.ToForwardRuleMapOutputWithContext(context.Background())
}

func (i ForwardRuleMap) ToForwardRuleMapOutputWithContext(ctx context.Context) ForwardRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardRuleMapOutput)
}

type ForwardRuleOutput struct{ *pulumi.OutputState }

func (ForwardRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardRule)(nil)).Elem()
}

func (o ForwardRuleOutput) ToForwardRuleOutput() ForwardRuleOutput {
	return o
}

func (o ForwardRuleOutput) ToForwardRuleOutputWithContext(ctx context.Context) ForwardRuleOutput {
	return o
}

// Endpoint ID.
func (o ForwardRuleOutput) EndPointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.EndPointId }).(pulumi.StringOutput)
}

// Forwarding rule name.
func (o ForwardRuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// Forwarding rule type. DOWN: From cloud to off-cloud; UP: From off-cloud to cloud.
func (o ForwardRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

// Private domain ID, which can be viewed on the private domain list page.
func (o ForwardRuleOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ForwardRule) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ForwardRuleArrayOutput struct{ *pulumi.OutputState }

func (ForwardRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ForwardRule)(nil)).Elem()
}

func (o ForwardRuleArrayOutput) ToForwardRuleArrayOutput() ForwardRuleArrayOutput {
	return o
}

func (o ForwardRuleArrayOutput) ToForwardRuleArrayOutputWithContext(ctx context.Context) ForwardRuleArrayOutput {
	return o
}

func (o ForwardRuleArrayOutput) Index(i pulumi.IntInput) ForwardRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ForwardRule {
		return vs[0].([]*ForwardRule)[vs[1].(int)]
	}).(ForwardRuleOutput)
}

type ForwardRuleMapOutput struct{ *pulumi.OutputState }

func (ForwardRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ForwardRule)(nil)).Elem()
}

func (o ForwardRuleMapOutput) ToForwardRuleMapOutput() ForwardRuleMapOutput {
	return o
}

func (o ForwardRuleMapOutput) ToForwardRuleMapOutputWithContext(ctx context.Context) ForwardRuleMapOutput {
	return o
}

func (o ForwardRuleMapOutput) MapIndex(k pulumi.StringInput) ForwardRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ForwardRule {
		return vs[0].(map[string]*ForwardRule)[vs[1].(string)]
	}).(ForwardRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardRuleInput)(nil)).Elem(), &ForwardRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardRuleArrayInput)(nil)).Elem(), ForwardRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardRuleMapInput)(nil)).Elem(), ForwardRuleMap{})
	pulumi.RegisterOutputType(ForwardRuleOutput{})
	pulumi.RegisterOutputType(ForwardRuleArrayOutput{})
	pulumi.RegisterOutputType(ForwardRuleMapOutput{})
}