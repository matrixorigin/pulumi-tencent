// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AiRecognitionTemplate struct {
	pulumi.CustomResourceState

	// Asr full text recognition control parameters.
	AsrFullTextConfigure AiRecognitionTemplateAsrFullTextConfigurePtrOutput `pulumi:"asrFullTextConfigure"`
	// Asr word recognition control parameters.
	AsrWordsConfigure AiRecognitionTemplateAsrWordsConfigurePtrOutput `pulumi:"asrWordsConfigure"`
	// Ai recognition template description information, length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Face recognition control parameters.
	FaceConfigure AiRecognitionTemplateFaceConfigurePtrOutput `pulumi:"faceConfigure"`
	// Ai recognition template name, length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Ocr full text control parameters.
	OcrFullTextConfigure AiRecognitionTemplateOcrFullTextConfigurePtrOutput `pulumi:"ocrFullTextConfigure"`
	// Ocr words recognition control parameters.
	OcrWordsConfigure AiRecognitionTemplateOcrWordsConfigurePtrOutput `pulumi:"ocrWordsConfigure"`
}

// NewAiRecognitionTemplate registers a new resource with the given unique name, arguments, and options.
func NewAiRecognitionTemplate(ctx *pulumi.Context,
	name string, args *AiRecognitionTemplateArgs, opts ...pulumi.ResourceOption) (*AiRecognitionTemplate, error) {
	if args == nil {
		args = &AiRecognitionTemplateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiRecognitionTemplate
	err := ctx.RegisterResource("tencentcloud:Mps/aiRecognitionTemplate:AiRecognitionTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiRecognitionTemplate gets an existing AiRecognitionTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiRecognitionTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiRecognitionTemplateState, opts ...pulumi.ResourceOption) (*AiRecognitionTemplate, error) {
	var resource AiRecognitionTemplate
	err := ctx.ReadResource("tencentcloud:Mps/aiRecognitionTemplate:AiRecognitionTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiRecognitionTemplate resources.
type aiRecognitionTemplateState struct {
	// Asr full text recognition control parameters.
	AsrFullTextConfigure *AiRecognitionTemplateAsrFullTextConfigure `pulumi:"asrFullTextConfigure"`
	// Asr word recognition control parameters.
	AsrWordsConfigure *AiRecognitionTemplateAsrWordsConfigure `pulumi:"asrWordsConfigure"`
	// Ai recognition template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Face recognition control parameters.
	FaceConfigure *AiRecognitionTemplateFaceConfigure `pulumi:"faceConfigure"`
	// Ai recognition template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Ocr full text control parameters.
	OcrFullTextConfigure *AiRecognitionTemplateOcrFullTextConfigure `pulumi:"ocrFullTextConfigure"`
	// Ocr words recognition control parameters.
	OcrWordsConfigure *AiRecognitionTemplateOcrWordsConfigure `pulumi:"ocrWordsConfigure"`
}

type AiRecognitionTemplateState struct {
	// Asr full text recognition control parameters.
	AsrFullTextConfigure AiRecognitionTemplateAsrFullTextConfigurePtrInput
	// Asr word recognition control parameters.
	AsrWordsConfigure AiRecognitionTemplateAsrWordsConfigurePtrInput
	// Ai recognition template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Face recognition control parameters.
	FaceConfigure AiRecognitionTemplateFaceConfigurePtrInput
	// Ai recognition template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Ocr full text control parameters.
	OcrFullTextConfigure AiRecognitionTemplateOcrFullTextConfigurePtrInput
	// Ocr words recognition control parameters.
	OcrWordsConfigure AiRecognitionTemplateOcrWordsConfigurePtrInput
}

func (AiRecognitionTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiRecognitionTemplateState)(nil)).Elem()
}

type aiRecognitionTemplateArgs struct {
	// Asr full text recognition control parameters.
	AsrFullTextConfigure *AiRecognitionTemplateAsrFullTextConfigure `pulumi:"asrFullTextConfigure"`
	// Asr word recognition control parameters.
	AsrWordsConfigure *AiRecognitionTemplateAsrWordsConfigure `pulumi:"asrWordsConfigure"`
	// Ai recognition template description information, length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Face recognition control parameters.
	FaceConfigure *AiRecognitionTemplateFaceConfigure `pulumi:"faceConfigure"`
	// Ai recognition template name, length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Ocr full text control parameters.
	OcrFullTextConfigure *AiRecognitionTemplateOcrFullTextConfigure `pulumi:"ocrFullTextConfigure"`
	// Ocr words recognition control parameters.
	OcrWordsConfigure *AiRecognitionTemplateOcrWordsConfigure `pulumi:"ocrWordsConfigure"`
}

// The set of arguments for constructing a AiRecognitionTemplate resource.
type AiRecognitionTemplateArgs struct {
	// Asr full text recognition control parameters.
	AsrFullTextConfigure AiRecognitionTemplateAsrFullTextConfigurePtrInput
	// Asr word recognition control parameters.
	AsrWordsConfigure AiRecognitionTemplateAsrWordsConfigurePtrInput
	// Ai recognition template description information, length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Face recognition control parameters.
	FaceConfigure AiRecognitionTemplateFaceConfigurePtrInput
	// Ai recognition template name, length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Ocr full text control parameters.
	OcrFullTextConfigure AiRecognitionTemplateOcrFullTextConfigurePtrInput
	// Ocr words recognition control parameters.
	OcrWordsConfigure AiRecognitionTemplateOcrWordsConfigurePtrInput
}

func (AiRecognitionTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiRecognitionTemplateArgs)(nil)).Elem()
}

type AiRecognitionTemplateInput interface {
	pulumi.Input

	ToAiRecognitionTemplateOutput() AiRecognitionTemplateOutput
	ToAiRecognitionTemplateOutputWithContext(ctx context.Context) AiRecognitionTemplateOutput
}

func (*AiRecognitionTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**AiRecognitionTemplate)(nil)).Elem()
}

func (i *AiRecognitionTemplate) ToAiRecognitionTemplateOutput() AiRecognitionTemplateOutput {
	return i.ToAiRecognitionTemplateOutputWithContext(context.Background())
}

func (i *AiRecognitionTemplate) ToAiRecognitionTemplateOutputWithContext(ctx context.Context) AiRecognitionTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiRecognitionTemplateOutput)
}

// AiRecognitionTemplateArrayInput is an input type that accepts AiRecognitionTemplateArray and AiRecognitionTemplateArrayOutput values.
// You can construct a concrete instance of `AiRecognitionTemplateArrayInput` via:
//
//	AiRecognitionTemplateArray{ AiRecognitionTemplateArgs{...} }
type AiRecognitionTemplateArrayInput interface {
	pulumi.Input

	ToAiRecognitionTemplateArrayOutput() AiRecognitionTemplateArrayOutput
	ToAiRecognitionTemplateArrayOutputWithContext(context.Context) AiRecognitionTemplateArrayOutput
}

type AiRecognitionTemplateArray []AiRecognitionTemplateInput

func (AiRecognitionTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiRecognitionTemplate)(nil)).Elem()
}

func (i AiRecognitionTemplateArray) ToAiRecognitionTemplateArrayOutput() AiRecognitionTemplateArrayOutput {
	return i.ToAiRecognitionTemplateArrayOutputWithContext(context.Background())
}

func (i AiRecognitionTemplateArray) ToAiRecognitionTemplateArrayOutputWithContext(ctx context.Context) AiRecognitionTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiRecognitionTemplateArrayOutput)
}

// AiRecognitionTemplateMapInput is an input type that accepts AiRecognitionTemplateMap and AiRecognitionTemplateMapOutput values.
// You can construct a concrete instance of `AiRecognitionTemplateMapInput` via:
//
//	AiRecognitionTemplateMap{ "key": AiRecognitionTemplateArgs{...} }
type AiRecognitionTemplateMapInput interface {
	pulumi.Input

	ToAiRecognitionTemplateMapOutput() AiRecognitionTemplateMapOutput
	ToAiRecognitionTemplateMapOutputWithContext(context.Context) AiRecognitionTemplateMapOutput
}

type AiRecognitionTemplateMap map[string]AiRecognitionTemplateInput

func (AiRecognitionTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiRecognitionTemplate)(nil)).Elem()
}

func (i AiRecognitionTemplateMap) ToAiRecognitionTemplateMapOutput() AiRecognitionTemplateMapOutput {
	return i.ToAiRecognitionTemplateMapOutputWithContext(context.Background())
}

func (i AiRecognitionTemplateMap) ToAiRecognitionTemplateMapOutputWithContext(ctx context.Context) AiRecognitionTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiRecognitionTemplateMapOutput)
}

type AiRecognitionTemplateOutput struct{ *pulumi.OutputState }

func (AiRecognitionTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiRecognitionTemplate)(nil)).Elem()
}

func (o AiRecognitionTemplateOutput) ToAiRecognitionTemplateOutput() AiRecognitionTemplateOutput {
	return o
}

func (o AiRecognitionTemplateOutput) ToAiRecognitionTemplateOutputWithContext(ctx context.Context) AiRecognitionTemplateOutput {
	return o
}

// Asr full text recognition control parameters.
func (o AiRecognitionTemplateOutput) AsrFullTextConfigure() AiRecognitionTemplateAsrFullTextConfigurePtrOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) AiRecognitionTemplateAsrFullTextConfigurePtrOutput {
		return v.AsrFullTextConfigure
	}).(AiRecognitionTemplateAsrFullTextConfigurePtrOutput)
}

// Asr word recognition control parameters.
func (o AiRecognitionTemplateOutput) AsrWordsConfigure() AiRecognitionTemplateAsrWordsConfigurePtrOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) AiRecognitionTemplateAsrWordsConfigurePtrOutput {
		return v.AsrWordsConfigure
	}).(AiRecognitionTemplateAsrWordsConfigurePtrOutput)
}

// Ai recognition template description information, length limit: 256 characters.
func (o AiRecognitionTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Face recognition control parameters.
func (o AiRecognitionTemplateOutput) FaceConfigure() AiRecognitionTemplateFaceConfigurePtrOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) AiRecognitionTemplateFaceConfigurePtrOutput { return v.FaceConfigure }).(AiRecognitionTemplateFaceConfigurePtrOutput)
}

// Ai recognition template name, length limit: 64 characters.
func (o AiRecognitionTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Ocr full text control parameters.
func (o AiRecognitionTemplateOutput) OcrFullTextConfigure() AiRecognitionTemplateOcrFullTextConfigurePtrOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) AiRecognitionTemplateOcrFullTextConfigurePtrOutput {
		return v.OcrFullTextConfigure
	}).(AiRecognitionTemplateOcrFullTextConfigurePtrOutput)
}

// Ocr words recognition control parameters.
func (o AiRecognitionTemplateOutput) OcrWordsConfigure() AiRecognitionTemplateOcrWordsConfigurePtrOutput {
	return o.ApplyT(func(v *AiRecognitionTemplate) AiRecognitionTemplateOcrWordsConfigurePtrOutput {
		return v.OcrWordsConfigure
	}).(AiRecognitionTemplateOcrWordsConfigurePtrOutput)
}

type AiRecognitionTemplateArrayOutput struct{ *pulumi.OutputState }

func (AiRecognitionTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiRecognitionTemplate)(nil)).Elem()
}

func (o AiRecognitionTemplateArrayOutput) ToAiRecognitionTemplateArrayOutput() AiRecognitionTemplateArrayOutput {
	return o
}

func (o AiRecognitionTemplateArrayOutput) ToAiRecognitionTemplateArrayOutputWithContext(ctx context.Context) AiRecognitionTemplateArrayOutput {
	return o
}

func (o AiRecognitionTemplateArrayOutput) Index(i pulumi.IntInput) AiRecognitionTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiRecognitionTemplate {
		return vs[0].([]*AiRecognitionTemplate)[vs[1].(int)]
	}).(AiRecognitionTemplateOutput)
}

type AiRecognitionTemplateMapOutput struct{ *pulumi.OutputState }

func (AiRecognitionTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiRecognitionTemplate)(nil)).Elem()
}

func (o AiRecognitionTemplateMapOutput) ToAiRecognitionTemplateMapOutput() AiRecognitionTemplateMapOutput {
	return o
}

func (o AiRecognitionTemplateMapOutput) ToAiRecognitionTemplateMapOutputWithContext(ctx context.Context) AiRecognitionTemplateMapOutput {
	return o
}

func (o AiRecognitionTemplateMapOutput) MapIndex(k pulumi.StringInput) AiRecognitionTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiRecognitionTemplate {
		return vs[0].(map[string]*AiRecognitionTemplate)[vs[1].(string)]
	}).(AiRecognitionTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiRecognitionTemplateInput)(nil)).Elem(), &AiRecognitionTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiRecognitionTemplateArrayInput)(nil)).Elem(), AiRecognitionTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiRecognitionTemplateMapInput)(nil)).Elem(), AiRecognitionTemplateMap{})
	pulumi.RegisterOutputType(AiRecognitionTemplateOutput{})
	pulumi.RegisterOutputType(AiRecognitionTemplateArrayOutput{})
	pulumi.RegisterOutputType(AiRecognitionTemplateMapOutput{})
}