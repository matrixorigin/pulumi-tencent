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

type MediaVideoMontageTemplate struct {
	pulumi.CustomResourceState

	// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
	Audio MediaVideoMontageTemplateAudioPtrOutput `pulumi:"audio"`
	// mixing parameters.
	AudioMixes MediaVideoMontageTemplateAudioMixArrayOutput `pulumi:"audioMixes"`
	// bucket name.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// container format.
	Container MediaVideoMontageTemplateContainerOutput `pulumi:"container"`
	// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution
	// accuracy is accurate to milliseconds.
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// video information, do not upload Video, which is equivalent to deleting video information.
	Video MediaVideoMontageTemplateVideoPtrOutput `pulumi:"video"`
}

// NewMediaVideoMontageTemplate registers a new resource with the given unique name, arguments, and options.
func NewMediaVideoMontageTemplate(ctx *pulumi.Context,
	name string, args *MediaVideoMontageTemplateArgs, opts ...pulumi.ResourceOption) (*MediaVideoMontageTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Container == nil {
		return nil, errors.New("invalid value for required argument 'Container'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MediaVideoMontageTemplate
	err := ctx.RegisterResource("tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaVideoMontageTemplate gets an existing MediaVideoMontageTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaVideoMontageTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaVideoMontageTemplateState, opts ...pulumi.ResourceOption) (*MediaVideoMontageTemplate, error) {
	var resource MediaVideoMontageTemplate
	err := ctx.ReadResource("tencentcloud:Ci/mediaVideoMontageTemplate:MediaVideoMontageTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaVideoMontageTemplate resources.
type mediaVideoMontageTemplateState struct {
	// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
	Audio *MediaVideoMontageTemplateAudio `pulumi:"audio"`
	// mixing parameters.
	AudioMixes []MediaVideoMontageTemplateAudioMix `pulumi:"audioMixes"`
	// bucket name.
	Bucket *string `pulumi:"bucket"`
	// container format.
	Container *MediaVideoMontageTemplateContainer `pulumi:"container"`
	// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution
	// accuracy is accurate to milliseconds.
	Duration *string `pulumi:"duration"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// video information, do not upload Video, which is equivalent to deleting video information.
	Video *MediaVideoMontageTemplateVideo `pulumi:"video"`
}

type MediaVideoMontageTemplateState struct {
	// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
	Audio MediaVideoMontageTemplateAudioPtrInput
	// mixing parameters.
	AudioMixes MediaVideoMontageTemplateAudioMixArrayInput
	// bucket name.
	Bucket pulumi.StringPtrInput
	// container format.
	Container MediaVideoMontageTemplateContainerPtrInput
	// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution
	// accuracy is accurate to milliseconds.
	Duration pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// video information, do not upload Video, which is equivalent to deleting video information.
	Video MediaVideoMontageTemplateVideoPtrInput
}

func (MediaVideoMontageTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaVideoMontageTemplateState)(nil)).Elem()
}

type mediaVideoMontageTemplateArgs struct {
	// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
	Audio *MediaVideoMontageTemplateAudio `pulumi:"audio"`
	// mixing parameters.
	AudioMixes []MediaVideoMontageTemplateAudioMix `pulumi:"audioMixes"`
	// bucket name.
	Bucket string `pulumi:"bucket"`
	// container format.
	Container MediaVideoMontageTemplateContainer `pulumi:"container"`
	// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution
	// accuracy is accurate to milliseconds.
	Duration *string `pulumi:"duration"`
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name *string `pulumi:"name"`
	// video information, do not upload Video, which is equivalent to deleting video information.
	Video *MediaVideoMontageTemplateVideo `pulumi:"video"`
}

// The set of arguments for constructing a MediaVideoMontageTemplate resource.
type MediaVideoMontageTemplateArgs struct {
	// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
	Audio MediaVideoMontageTemplateAudioPtrInput
	// mixing parameters.
	AudioMixes MediaVideoMontageTemplateAudioMixArrayInput
	// bucket name.
	Bucket pulumi.StringInput
	// container format.
	Container MediaVideoMontageTemplateContainerInput
	// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution
	// accuracy is accurate to milliseconds.
	Duration pulumi.StringPtrInput
	// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
	Name pulumi.StringPtrInput
	// video information, do not upload Video, which is equivalent to deleting video information.
	Video MediaVideoMontageTemplateVideoPtrInput
}

func (MediaVideoMontageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaVideoMontageTemplateArgs)(nil)).Elem()
}

type MediaVideoMontageTemplateInput interface {
	pulumi.Input

	ToMediaVideoMontageTemplateOutput() MediaVideoMontageTemplateOutput
	ToMediaVideoMontageTemplateOutputWithContext(ctx context.Context) MediaVideoMontageTemplateOutput
}

func (*MediaVideoMontageTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaVideoMontageTemplate)(nil)).Elem()
}

func (i *MediaVideoMontageTemplate) ToMediaVideoMontageTemplateOutput() MediaVideoMontageTemplateOutput {
	return i.ToMediaVideoMontageTemplateOutputWithContext(context.Background())
}

func (i *MediaVideoMontageTemplate) ToMediaVideoMontageTemplateOutputWithContext(ctx context.Context) MediaVideoMontageTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaVideoMontageTemplateOutput)
}

// MediaVideoMontageTemplateArrayInput is an input type that accepts MediaVideoMontageTemplateArray and MediaVideoMontageTemplateArrayOutput values.
// You can construct a concrete instance of `MediaVideoMontageTemplateArrayInput` via:
//
//	MediaVideoMontageTemplateArray{ MediaVideoMontageTemplateArgs{...} }
type MediaVideoMontageTemplateArrayInput interface {
	pulumi.Input

	ToMediaVideoMontageTemplateArrayOutput() MediaVideoMontageTemplateArrayOutput
	ToMediaVideoMontageTemplateArrayOutputWithContext(context.Context) MediaVideoMontageTemplateArrayOutput
}

type MediaVideoMontageTemplateArray []MediaVideoMontageTemplateInput

func (MediaVideoMontageTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaVideoMontageTemplate)(nil)).Elem()
}

func (i MediaVideoMontageTemplateArray) ToMediaVideoMontageTemplateArrayOutput() MediaVideoMontageTemplateArrayOutput {
	return i.ToMediaVideoMontageTemplateArrayOutputWithContext(context.Background())
}

func (i MediaVideoMontageTemplateArray) ToMediaVideoMontageTemplateArrayOutputWithContext(ctx context.Context) MediaVideoMontageTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaVideoMontageTemplateArrayOutput)
}

// MediaVideoMontageTemplateMapInput is an input type that accepts MediaVideoMontageTemplateMap and MediaVideoMontageTemplateMapOutput values.
// You can construct a concrete instance of `MediaVideoMontageTemplateMapInput` via:
//
//	MediaVideoMontageTemplateMap{ "key": MediaVideoMontageTemplateArgs{...} }
type MediaVideoMontageTemplateMapInput interface {
	pulumi.Input

	ToMediaVideoMontageTemplateMapOutput() MediaVideoMontageTemplateMapOutput
	ToMediaVideoMontageTemplateMapOutputWithContext(context.Context) MediaVideoMontageTemplateMapOutput
}

type MediaVideoMontageTemplateMap map[string]MediaVideoMontageTemplateInput

func (MediaVideoMontageTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaVideoMontageTemplate)(nil)).Elem()
}

func (i MediaVideoMontageTemplateMap) ToMediaVideoMontageTemplateMapOutput() MediaVideoMontageTemplateMapOutput {
	return i.ToMediaVideoMontageTemplateMapOutputWithContext(context.Background())
}

func (i MediaVideoMontageTemplateMap) ToMediaVideoMontageTemplateMapOutputWithContext(ctx context.Context) MediaVideoMontageTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaVideoMontageTemplateMapOutput)
}

type MediaVideoMontageTemplateOutput struct{ *pulumi.OutputState }

func (MediaVideoMontageTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaVideoMontageTemplate)(nil)).Elem()
}

func (o MediaVideoMontageTemplateOutput) ToMediaVideoMontageTemplateOutput() MediaVideoMontageTemplateOutput {
	return o
}

func (o MediaVideoMontageTemplateOutput) ToMediaVideoMontageTemplateOutputWithContext(ctx context.Context) MediaVideoMontageTemplateOutput {
	return o
}

// audio parameters, the target file does not require Audio information, need to set Audio.Remove to true.
func (o MediaVideoMontageTemplateOutput) Audio() MediaVideoMontageTemplateAudioPtrOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) MediaVideoMontageTemplateAudioPtrOutput { return v.Audio }).(MediaVideoMontageTemplateAudioPtrOutput)
}

// mixing parameters.
func (o MediaVideoMontageTemplateOutput) AudioMixes() MediaVideoMontageTemplateAudioMixArrayOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) MediaVideoMontageTemplateAudioMixArrayOutput { return v.AudioMixes }).(MediaVideoMontageTemplateAudioMixArrayOutput)
}

// bucket name.
func (o MediaVideoMontageTemplateOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// container format.
func (o MediaVideoMontageTemplateOutput) Container() MediaVideoMontageTemplateContainerOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) MediaVideoMontageTemplateContainerOutput { return v.Container }).(MediaVideoMontageTemplateContainerOutput)
}

// Collection duration 1: Default automatic analysis duration, 2: The unit is seconds, 3: Support float format, execution
// accuracy is accurate to milliseconds.
func (o MediaVideoMontageTemplateOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
func (o MediaVideoMontageTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// video information, do not upload Video, which is equivalent to deleting video information.
func (o MediaVideoMontageTemplateOutput) Video() MediaVideoMontageTemplateVideoPtrOutput {
	return o.ApplyT(func(v *MediaVideoMontageTemplate) MediaVideoMontageTemplateVideoPtrOutput { return v.Video }).(MediaVideoMontageTemplateVideoPtrOutput)
}

type MediaVideoMontageTemplateArrayOutput struct{ *pulumi.OutputState }

func (MediaVideoMontageTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MediaVideoMontageTemplate)(nil)).Elem()
}

func (o MediaVideoMontageTemplateArrayOutput) ToMediaVideoMontageTemplateArrayOutput() MediaVideoMontageTemplateArrayOutput {
	return o
}

func (o MediaVideoMontageTemplateArrayOutput) ToMediaVideoMontageTemplateArrayOutputWithContext(ctx context.Context) MediaVideoMontageTemplateArrayOutput {
	return o
}

func (o MediaVideoMontageTemplateArrayOutput) Index(i pulumi.IntInput) MediaVideoMontageTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MediaVideoMontageTemplate {
		return vs[0].([]*MediaVideoMontageTemplate)[vs[1].(int)]
	}).(MediaVideoMontageTemplateOutput)
}

type MediaVideoMontageTemplateMapOutput struct{ *pulumi.OutputState }

func (MediaVideoMontageTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MediaVideoMontageTemplate)(nil)).Elem()
}

func (o MediaVideoMontageTemplateMapOutput) ToMediaVideoMontageTemplateMapOutput() MediaVideoMontageTemplateMapOutput {
	return o
}

func (o MediaVideoMontageTemplateMapOutput) ToMediaVideoMontageTemplateMapOutputWithContext(ctx context.Context) MediaVideoMontageTemplateMapOutput {
	return o
}

func (o MediaVideoMontageTemplateMapOutput) MapIndex(k pulumi.StringInput) MediaVideoMontageTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MediaVideoMontageTemplate {
		return vs[0].(map[string]*MediaVideoMontageTemplate)[vs[1].(string)]
	}).(MediaVideoMontageTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MediaVideoMontageTemplateInput)(nil)).Elem(), &MediaVideoMontageTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaVideoMontageTemplateArrayInput)(nil)).Elem(), MediaVideoMontageTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MediaVideoMontageTemplateMapInput)(nil)).Elem(), MediaVideoMontageTemplateMap{})
	pulumi.RegisterOutputType(MediaVideoMontageTemplateOutput{})
	pulumi.RegisterOutputType(MediaVideoMontageTemplateArrayOutput{})
	pulumi.RegisterOutputType(MediaVideoMontageTemplateMapOutput{})
}