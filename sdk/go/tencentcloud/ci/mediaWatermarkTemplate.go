// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ci

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MediaWatermarkTemplate struct {
	pulumi.CustomResourceState

	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// container format.
	Watermark MediaWatermarkTemplateWatermarkOutput `pulumi:"watermark"`
}

// NewMediaWatermarkTemplate registers a new resource with the given unique name, arguments, and options.
func NewMediaWatermarkTemplate(ctx *pulumi.Context,
	name string, args *MediaWatermarkTemplateArgs, opts ...pulumi.ResourceOption) (*MediaWatermarkTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Watermark == nil {
		return nil, errors.New("invalid value for required argument 'Watermark'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MediaWatermarkTemplate
	err := ctx.RegisterResource("tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaWatermarkTemplate gets an existing MediaWatermarkTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaWatermarkTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaWatermarkTemplateState, opts ...pulumi.ResourceOption) (*MediaWatermarkTemplate, error) {
	var resource MediaWatermarkTemplate
	err := ctx.ReadResource("tencentcloud:Ci/mediaWatermarkTemplate:MediaWatermarkTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaWatermarkTemplate resources.
type mediaWatermarkTemplateState struct {
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// container format.
	Watermark *MediaWatermarkTemplateWatermark `pulumi:"watermark"`
}

type MediaWatermarkTemplateState struct {
	// bucket name.
	Bucket pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// container format.
	Watermark MediaWatermarkTemplateWatermarkPtrInput
}

func (MediaWatermarkTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaWatermarkTemplateState)(nil)).Elem()
}

type mediaWatermarkTemplateArgs struct {
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// container format.
	Watermark MediaWatermarkTemplateWatermark `pulumi:"watermark"`
}

// The set of arguments for constructing a MediaWatermarkTemplate resource.
type MediaWatermarkTemplateArgs struct {
	// bucket name.
	Bucket pulumi.StringInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// container format.
	Watermark MediaWatermarkTemplateWatermarkInput
}

func (MediaWatermarkTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaWatermarkTemplateArgs)(nil)).Elem()
}

type MediaWatermarkTemplateInput interface {
	pulumi.Input

	ToMediaWatermarkTemplateOutput() MediaWatermarkTemplateOutput
	ToMediaWatermarkTemplateOutputWithContext(ctx context.Context) MediaWatermarkTemplateOutput
}

func (*MediaWatermarkTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaWatermarkTemplate)(nil)).Elem()
}

func (i *MediaWatermarkTemplate) ToMediaWatermarkTemplateOutput() MediaWatermarkTemplateOutput {
	return i.ToMediaWatermarkTemplateOutputWithContext(context.Background())
}

func (i *MediaWatermarkTemplate) ToMediaWatermarkTemplateOutputWithContext(ctx context.Context) MediaWatermarkTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaWatermarkTemplateOutput)
}

// MediaWatermarkTemplateArrayInput is an input type that accepts MediaWatermarkTemplateArray and MediaWatermarkTemplateArrayOutput values.
// You can construct a concrete instance of `MediaWatermarkTemplateArrayInput` via:
//
//	MediaWatermarkTemplateArray{ MediaWatermarkTemplateArgs{...} }
type MediaWatermarkTemplateArrayInput interface {
	pulumi.Input

	ToMediaWatermarkTemplateArrayOutput() MediaWatermarkTemplateArrayOutput
	ToMediaWatermarkTemplateArrayOutputWithContext(context.Context) MediaWatermarkTemplateArrayOutput
}

type MediaWatermarkTemplateArray []MediaWatermarkTemplateInput

func (MediaWatermarkTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaWatermarkTemplate)(nil)).Elem()
}

func (i MediaWatermarkTemplateArray) ToMediaWatermarkTemplateArrayOutput() MediaWatermarkTemplateArrayOutput {
	return i.ToMediaWatermarkTemplateArrayOutputWithContext(context.Background())
}

func (i MediaWatermarkTemplateArray) ToMediaWatermarkTemplateArrayOutputWithContext(ctx context.Context) MediaWatermarkTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaWatermarkTemplateArrayOutput)
}

// MediaWatermarkTemplateMapInput is an input type that accepts MediaWatermarkTemplateMap and MediaWatermarkTemplateMapOutput values.
// You can construct a concrete instance of `MediaWatermarkTemplateMapInput` via:
//
//	MediaWatermarkTemplateMap{ "key": MediaWatermarkTemplateArgs{...} }
type MediaWatermarkTemplateMapInput interface {
	pulumi.Input

	ToMediaWatermarkTemplateMapOutput() MediaWatermarkTemplateMapOutput
	ToMediaWatermarkTemplateMapOutputWithContext(context.Context) MediaWatermarkTemplateMapOutput
}

type MediaWatermarkTemplateMap map[string]MediaWatermarkTemplateInput

func (MediaWatermarkTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaWatermarkTemplate)(nil)).Elem()
}

func (i MediaWatermarkTemplateMap) ToMediaWatermarkTemplateMapOutput() MediaWatermarkTemplateMapOutput {
	return i.ToMediaWatermarkTemplateMapOutputWithContext(context.Background())
}

func (i MediaWatermarkTemplateMap) ToMediaWatermarkTemplateMapOutputWithContext(ctx context.Context) MediaWatermarkTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaWatermarkTemplateMapOutput)
}

type MediaWatermarkTemplateOutput struct{ *pulumi.OutputState }

func (MediaWatermarkTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaWatermarkTemplate)(nil)).Elem()
}

func (o MediaWatermarkTemplateOutput) ToMediaWatermarkTemplateOutput() MediaWatermarkTemplateOutput {
	return o
}

func (o MediaWatermarkTemplateOutput) ToMediaWatermarkTemplateOutputWithContext(ctx context.Context) MediaWatermarkTemplateOutput {
	return o
}

// bucket name.
func (o MediaWatermarkTemplateOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaWatermarkTemplate) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
func (o MediaWatermarkTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaWatermarkTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// container format.
func (o MediaWatermarkTemplateOutput) Watermark() MediaWatermarkTemplateWatermarkOutput {
	return o.ApplyT(func(v *MediaWatermarkTemplate) MediaWatermarkTemplateWatermarkOutput { return v.Watermark }).(MediaWatermarkTemplateWatermarkOutput)
}

type MediaWatermarkTemplateArrayOutput struct{ *pulumi.OutputState }

func (MediaWatermarkTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaWatermarkTemplate)(nil)).Elem()
}

func (o MediaWatermarkTemplateArrayOutput) ToMediaWatermarkTemplateArrayOutput() MediaWatermarkTemplateArrayOutput {
	return o
}

func (o MediaWatermarkTemplateArrayOutput) ToMediaWatermarkTemplateArrayOutputWithContext(ctx context.Context) MediaWatermarkTemplateArrayOutput {
	return o
}

func (o MediaWatermarkTemplateArrayOutput) Index(i pulumi.IntInput) MediaWatermarkTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MediaWatermarkTemplate {
		return vs[0].([]*MediaWatermarkTemplate)[vs[1].(int)]
	}).(MediaWatermarkTemplateOutput)
}

type MediaWatermarkTemplateMapOutput struct{ *pulumi.OutputState }

func (MediaWatermarkTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaWatermarkTemplate)(nil)).Elem()
}

func (o MediaWatermarkTemplateMapOutput) ToMediaWatermarkTemplateMapOutput() MediaWatermarkTemplateMapOutput {
	return o
}

func (o MediaWatermarkTemplateMapOutput) ToMediaWatermarkTemplateMapOutputWithContext(ctx context.Context) MediaWatermarkTemplateMapOutput {
	return o
}

func (o MediaWatermarkTemplateMapOutput) MapIndex(k pulumi.StringInput) MediaWatermarkTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MediaWatermarkTemplate {
		return vs[0].(map[string]*MediaWatermarkTemplate)[vs[1].(string)]
	}).(MediaWatermarkTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaWatermarkTemplateInput)(nil)).Elem(), &MediaWatermarkTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaWatermarkTemplateArrayInput)(nil)).Elem(), MediaWatermarkTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaWatermarkTemplateMapInput)(nil)).Elem(), MediaWatermarkTemplateMap{})
	pulumi.RegisterOutputType(MediaWatermarkTemplateOutput{})
	pulumi.RegisterOutputType(MediaWatermarkTemplateArrayOutput{})
	pulumi.RegisterOutputType(MediaWatermarkTemplateMapOutput{})
}