// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiKeyAttachment struct {
	pulumi.CustomResourceState

	// ID of API key.
	ApiKeyId pulumi.StringOutput `pulumi:"apiKeyId"`
	// ID of the usage plan.
	UsagePlanId pulumi.StringOutput `pulumi:"usagePlanId"`
}

// NewApiKeyAttachment registers a new resource with the given unique name, arguments, and options.
func NewApiKeyAttachment(ctx *pulumi.Context,
	name string, args *ApiKeyAttachmentArgs, opts ...pulumi.ResourceOption) (*ApiKeyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKeyId == nil {
		return nil, errors.New("invalid value for required argument 'ApiKeyId'")
	}
	if args.UsagePlanId == nil {
		return nil, errors.New("invalid value for required argument 'UsagePlanId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiKeyAttachment
	err := ctx.RegisterResource("tencentcloud:ApiGateway/apiKeyAttachment:ApiKeyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKeyAttachment gets an existing ApiKeyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKeyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyAttachmentState, opts ...pulumi.ResourceOption) (*ApiKeyAttachment, error) {
	var resource ApiKeyAttachment
	err := ctx.ReadResource("tencentcloud:ApiGateway/apiKeyAttachment:ApiKeyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKeyAttachment resources.
type apiKeyAttachmentState struct {
	// ID of API key.
	ApiKeyId *string `pulumi:"apiKeyId"`
	// ID of the usage plan.
	UsagePlanId *string `pulumi:"usagePlanId"`
}

type ApiKeyAttachmentState struct {
	// ID of API key.
	ApiKeyId pulumi.StringPtrInput
	// ID of the usage plan.
	UsagePlanId pulumi.StringPtrInput
}

func (ApiKeyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyAttachmentState)(nil)).Elem()
}

type apiKeyAttachmentArgs struct {
	// ID of API key.
	ApiKeyId string `pulumi:"apiKeyId"`
	// ID of the usage plan.
	UsagePlanId string `pulumi:"usagePlanId"`
}

// The set of arguments for constructing a ApiKeyAttachment resource.
type ApiKeyAttachmentArgs struct {
	// ID of API key.
	ApiKeyId pulumi.StringInput
	// ID of the usage plan.
	UsagePlanId pulumi.StringInput
}

func (ApiKeyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyAttachmentArgs)(nil)).Elem()
}

type ApiKeyAttachmentInput interface {
	pulumi.Input

	ToApiKeyAttachmentOutput() ApiKeyAttachmentOutput
	ToApiKeyAttachmentOutputWithContext(ctx context.Context) ApiKeyAttachmentOutput
}

func (*ApiKeyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKeyAttachment)(nil)).Elem()
}

func (i *ApiKeyAttachment) ToApiKeyAttachmentOutput() ApiKeyAttachmentOutput {
	return i.ToApiKeyAttachmentOutputWithContext(context.Background())
}

func (i *ApiKeyAttachment) ToApiKeyAttachmentOutputWithContext(ctx context.Context) ApiKeyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyAttachmentOutput)
}

// ApiKeyAttachmentArrayInput is an input type that accepts ApiKeyAttachmentArray and ApiKeyAttachmentArrayOutput values.
// You can construct a concrete instance of `ApiKeyAttachmentArrayInput` via:
//
//	ApiKeyAttachmentArray{ ApiKeyAttachmentArgs{...} }
type ApiKeyAttachmentArrayInput interface {
	pulumi.Input

	ToApiKeyAttachmentArrayOutput() ApiKeyAttachmentArrayOutput
	ToApiKeyAttachmentArrayOutputWithContext(context.Context) ApiKeyAttachmentArrayOutput
}

type ApiKeyAttachmentArray []ApiKeyAttachmentInput

func (ApiKeyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiKeyAttachment)(nil)).Elem()
}

func (i ApiKeyAttachmentArray) ToApiKeyAttachmentArrayOutput() ApiKeyAttachmentArrayOutput {
	return i.ToApiKeyAttachmentArrayOutputWithContext(context.Background())
}

func (i ApiKeyAttachmentArray) ToApiKeyAttachmentArrayOutputWithContext(ctx context.Context) ApiKeyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyAttachmentArrayOutput)
}

// ApiKeyAttachmentMapInput is an input type that accepts ApiKeyAttachmentMap and ApiKeyAttachmentMapOutput values.
// You can construct a concrete instance of `ApiKeyAttachmentMapInput` via:
//
//	ApiKeyAttachmentMap{ "key": ApiKeyAttachmentArgs{...} }
type ApiKeyAttachmentMapInput interface {
	pulumi.Input

	ToApiKeyAttachmentMapOutput() ApiKeyAttachmentMapOutput
	ToApiKeyAttachmentMapOutputWithContext(context.Context) ApiKeyAttachmentMapOutput
}

type ApiKeyAttachmentMap map[string]ApiKeyAttachmentInput

func (ApiKeyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiKeyAttachment)(nil)).Elem()
}

func (i ApiKeyAttachmentMap) ToApiKeyAttachmentMapOutput() ApiKeyAttachmentMapOutput {
	return i.ToApiKeyAttachmentMapOutputWithContext(context.Background())
}

func (i ApiKeyAttachmentMap) ToApiKeyAttachmentMapOutputWithContext(ctx context.Context) ApiKeyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyAttachmentMapOutput)
}

type ApiKeyAttachmentOutput struct{ *pulumi.OutputState }

func (ApiKeyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKeyAttachment)(nil)).Elem()
}

func (o ApiKeyAttachmentOutput) ToApiKeyAttachmentOutput() ApiKeyAttachmentOutput {
	return o
}

func (o ApiKeyAttachmentOutput) ToApiKeyAttachmentOutputWithContext(ctx context.Context) ApiKeyAttachmentOutput {
	return o
}

// ID of API key.
func (o ApiKeyAttachmentOutput) ApiKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKeyAttachment) pulumi.StringOutput { return v.ApiKeyId }).(pulumi.StringOutput)
}

// ID of the usage plan.
func (o ApiKeyAttachmentOutput) UsagePlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKeyAttachment) pulumi.StringOutput { return v.UsagePlanId }).(pulumi.StringOutput)
}

type ApiKeyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiKeyAttachment)(nil)).Elem()
}

func (o ApiKeyAttachmentArrayOutput) ToApiKeyAttachmentArrayOutput() ApiKeyAttachmentArrayOutput {
	return o
}

func (o ApiKeyAttachmentArrayOutput) ToApiKeyAttachmentArrayOutputWithContext(ctx context.Context) ApiKeyAttachmentArrayOutput {
	return o
}

func (o ApiKeyAttachmentArrayOutput) Index(i pulumi.IntInput) ApiKeyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiKeyAttachment {
		return vs[0].([]*ApiKeyAttachment)[vs[1].(int)]
	}).(ApiKeyAttachmentOutput)
}

type ApiKeyAttachmentMapOutput struct{ *pulumi.OutputState }

func (ApiKeyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiKeyAttachment)(nil)).Elem()
}

func (o ApiKeyAttachmentMapOutput) ToApiKeyAttachmentMapOutput() ApiKeyAttachmentMapOutput {
	return o
}

func (o ApiKeyAttachmentMapOutput) ToApiKeyAttachmentMapOutputWithContext(ctx context.Context) ApiKeyAttachmentMapOutput {
	return o
}

func (o ApiKeyAttachmentMapOutput) MapIndex(k pulumi.StringInput) ApiKeyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiKeyAttachment {
		return vs[0].(map[string]*ApiKeyAttachment)[vs[1].(string)]
	}).(ApiKeyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyAttachmentInput)(nil)).Elem(), &ApiKeyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyAttachmentArrayInput)(nil)).Elem(), ApiKeyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyAttachmentMapInput)(nil)).Elem(), ApiKeyAttachmentMap{})
	pulumi.RegisterOutputType(ApiKeyAttachmentOutput{})
	pulumi.RegisterOutputType(ApiKeyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ApiKeyAttachmentMapOutput{})
}