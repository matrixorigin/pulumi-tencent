// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SamlProvider struct {
	pulumi.CustomResourceState

	// The create time of the CAM SAML provider.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the CAM SAML provider.
	Description pulumi.StringOutput `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringOutput `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the CAM SAML provider.
	ProviderArn pulumi.StringOutput `pulumi:"providerArn"`
	// The last update time of the CAM SAML provider.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSamlProvider registers a new resource with the given unique name, arguments, and options.
func NewSamlProvider(ctx *pulumi.Context,
	name string, args *SamlProviderArgs, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.MetaData == nil {
		return nil, errors.New("invalid value for required argument 'MetaData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SamlProvider
	err := ctx.RegisterResource("tencentcloud:Cam/samlProvider:SamlProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlProvider gets an existing SamlProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlProviderState, opts ...pulumi.ResourceOption) (*SamlProvider, error) {
	var resource SamlProvider
	err := ctx.ReadResource("tencentcloud:Cam/samlProvider:SamlProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlProvider resources.
type samlProviderState struct {
	// The create time of the CAM SAML provider.
	CreateTime *string `pulumi:"createTime"`
	// The description of the CAM SAML provider.
	Description *string `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData *string `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
	// The ARN of the CAM SAML provider.
	ProviderArn *string `pulumi:"providerArn"`
	// The last update time of the CAM SAML provider.
	UpdateTime *string `pulumi:"updateTime"`
}

type SamlProviderState struct {
	// The create time of the CAM SAML provider.
	CreateTime pulumi.StringPtrInput
	// The description of the CAM SAML provider.
	Description pulumi.StringPtrInput
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringPtrInput
	// Name of CAM SAML provider.
	Name pulumi.StringPtrInput
	// The ARN of the CAM SAML provider.
	ProviderArn pulumi.StringPtrInput
	// The last update time of the CAM SAML provider.
	UpdateTime pulumi.StringPtrInput
}

func (SamlProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderState)(nil)).Elem()
}

type samlProviderArgs struct {
	// The description of the CAM SAML provider.
	Description string `pulumi:"description"`
	// The meta data document of the CAM SAML provider.
	MetaData string `pulumi:"metaData"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SamlProvider resource.
type SamlProviderArgs struct {
	// The description of the CAM SAML provider.
	Description pulumi.StringInput
	// The meta data document of the CAM SAML provider.
	MetaData pulumi.StringInput
	// Name of CAM SAML provider.
	Name pulumi.StringPtrInput
}

func (SamlProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlProviderArgs)(nil)).Elem()
}

type SamlProviderInput interface {
	pulumi.Input

	ToSamlProviderOutput() SamlProviderOutput
	ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput
}

func (*SamlProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil)).Elem()
}

func (i *SamlProvider) ToSamlProviderOutput() SamlProviderOutput {
	return i.ToSamlProviderOutputWithContext(context.Background())
}

func (i *SamlProvider) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderOutput)
}

// SamlProviderArrayInput is an input type that accepts SamlProviderArray and SamlProviderArrayOutput values.
// You can construct a concrete instance of `SamlProviderArrayInput` via:
//
//	SamlProviderArray{ SamlProviderArgs{...} }
type SamlProviderArrayInput interface {
	pulumi.Input

	ToSamlProviderArrayOutput() SamlProviderArrayOutput
	ToSamlProviderArrayOutputWithContext(context.Context) SamlProviderArrayOutput
}

type SamlProviderArray []SamlProviderInput

func (SamlProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderArray) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return i.ToSamlProviderArrayOutputWithContext(context.Background())
}

func (i SamlProviderArray) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderArrayOutput)
}

// SamlProviderMapInput is an input type that accepts SamlProviderMap and SamlProviderMapOutput values.
// You can construct a concrete instance of `SamlProviderMapInput` via:
//
//	SamlProviderMap{ "key": SamlProviderArgs{...} }
type SamlProviderMapInput interface {
	pulumi.Input

	ToSamlProviderMapOutput() SamlProviderMapOutput
	ToSamlProviderMapOutputWithContext(context.Context) SamlProviderMapOutput
}

type SamlProviderMap map[string]SamlProviderInput

func (SamlProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (i SamlProviderMap) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return i.ToSamlProviderMapOutputWithContext(context.Background())
}

func (i SamlProviderMap) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlProviderMapOutput)
}

type SamlProviderOutput struct{ *pulumi.OutputState }

func (SamlProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlProvider)(nil)).Elem()
}

func (o SamlProviderOutput) ToSamlProviderOutput() SamlProviderOutput {
	return o
}

func (o SamlProviderOutput) ToSamlProviderOutputWithContext(ctx context.Context) SamlProviderOutput {
	return o
}

// The create time of the CAM SAML provider.
func (o SamlProviderOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the CAM SAML provider.
func (o SamlProviderOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The meta data document of the CAM SAML provider.
func (o SamlProviderOutput) MetaData() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.MetaData }).(pulumi.StringOutput)
}

// Name of CAM SAML provider.
func (o SamlProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the CAM SAML provider.
func (o SamlProviderOutput) ProviderArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.ProviderArn }).(pulumi.StringOutput)
}

// The last update time of the CAM SAML provider.
func (o SamlProviderOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlProvider) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type SamlProviderArrayOutput struct{ *pulumi.OutputState }

func (SamlProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlProvider)(nil)).Elem()
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutput() SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) ToSamlProviderArrayOutputWithContext(ctx context.Context) SamlProviderArrayOutput {
	return o
}

func (o SamlProviderArrayOutput) Index(i pulumi.IntInput) SamlProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SamlProvider {
		return vs[0].([]*SamlProvider)[vs[1].(int)]
	}).(SamlProviderOutput)
}

type SamlProviderMapOutput struct{ *pulumi.OutputState }

func (SamlProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlProvider)(nil)).Elem()
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutput() SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) ToSamlProviderMapOutputWithContext(ctx context.Context) SamlProviderMapOutput {
	return o
}

func (o SamlProviderMapOutput) MapIndex(k pulumi.StringInput) SamlProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SamlProvider {
		return vs[0].(map[string]*SamlProvider)[vs[1].(string)]
	}).(SamlProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderInput)(nil)).Elem(), &SamlProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderArrayInput)(nil)).Elem(), SamlProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlProviderMapInput)(nil)).Elem(), SamlProviderMap{})
	pulumi.RegisterOutputType(SamlProviderOutput{})
	pulumi.RegisterOutputType(SamlProviderArrayOutput{})
	pulumi.RegisterOutputType(SamlProviderMapOutput{})
}