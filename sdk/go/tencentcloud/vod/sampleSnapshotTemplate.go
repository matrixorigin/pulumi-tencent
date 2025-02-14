// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vod

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SampleSnapshotTemplate struct {
	pulumi.CustomResourceState

	// Template description. Length limit: 256 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source
	// video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to
	// match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black.
	// This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black
	// color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and
	// fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio
	// of the source video for the screenshot and fills the unmatched area with Gaussian blur.Default value: black.
	FillType pulumi.StringPtrOutput `pulumi:"fillType"`
	// Image format. Valid values: jpg, png. Default value: jpg.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resolution adaption. Valid values: open: enabled. In this case, `Width` represents the long side of a video, while
	// `Height` the short side; close: disabled. In this case, `Width` represents the width of a video, while `Height` the
	// height.Default value: open.
	ResolutionAdaptive pulumi.StringPtrOutput `pulumi:"resolutionAdaptive"`
	// Sampling interval. If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.
	// If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
	SampleInterval pulumi.IntOutput `pulumi:"sampleInterval"`
	// Sampled screencapturing type. Valid values: Percent: by percent. Time: by time interval.
	SampleType pulumi.StringOutput `pulumi:"sampleType"`
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
	// service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
	// application or a newly created one), they must fill in this field with the application ID.
	SubAppId pulumi.IntOutput `pulumi:"subAppId"`
	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewSampleSnapshotTemplate registers a new resource with the given unique name, arguments, and options.
func NewSampleSnapshotTemplate(ctx *pulumi.Context,
	name string, args *SampleSnapshotTemplateArgs, opts ...pulumi.ResourceOption) (*SampleSnapshotTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SampleInterval == nil {
		return nil, errors.New("invalid value for required argument 'SampleInterval'")
	}
	if args.SampleType == nil {
		return nil, errors.New("invalid value for required argument 'SampleType'")
	}
	if args.SubAppId == nil {
		return nil, errors.New("invalid value for required argument 'SubAppId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SampleSnapshotTemplate
	err := ctx.RegisterResource("tencentcloud:Vod/sampleSnapshotTemplate:SampleSnapshotTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSampleSnapshotTemplate gets an existing SampleSnapshotTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSampleSnapshotTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SampleSnapshotTemplateState, opts ...pulumi.ResourceOption) (*SampleSnapshotTemplate, error) {
	var resource SampleSnapshotTemplate
	err := ctx.ReadResource("tencentcloud:Vod/sampleSnapshotTemplate:SampleSnapshotTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SampleSnapshotTemplate resources.
type sampleSnapshotTemplateState struct {
	// Template description. Length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source
	// video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to
	// match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black.
	// This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black
	// color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and
	// fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio
	// of the source video for the screenshot and fills the unmatched area with Gaussian blur.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format. Valid values: jpg, png. Default value: jpg.
	Format *string `pulumi:"format"`
	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Height *int `pulumi:"height"`
	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Resolution adaption. Valid values: open: enabled. In this case, `Width` represents the long side of a video, while
	// `Height` the short side; close: disabled. In this case, `Width` represents the width of a video, while `Height` the
	// height.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// Sampling interval. If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.
	// If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
	SampleInterval *int `pulumi:"sampleInterval"`
	// Sampled screencapturing type. Valid values: Percent: by percent. Time: by time interval.
	SampleType *string `pulumi:"sampleType"`
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
	// service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
	// application or a newly created one), they must fill in this field with the application ID.
	SubAppId *int `pulumi:"subAppId"`
	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Width *int `pulumi:"width"`
}

type SampleSnapshotTemplateState struct {
	// Template description. Length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source
	// video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to
	// match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black.
	// This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black
	// color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and
	// fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio
	// of the source video for the screenshot and fills the unmatched area with Gaussian blur.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format. Valid values: jpg, png. Default value: jpg.
	Format pulumi.StringPtrInput
	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Height pulumi.IntPtrInput
	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Resolution adaption. Valid values: open: enabled. In this case, `Width` represents the long side of a video, while
	// `Height` the short side; close: disabled. In this case, `Width` represents the width of a video, while `Height` the
	// height.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// Sampling interval. If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.
	// If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
	SampleInterval pulumi.IntPtrInput
	// Sampled screencapturing type. Valid values: Percent: by percent. Time: by time interval.
	SampleType pulumi.StringPtrInput
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
	// service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
	// application or a newly created one), they must fill in this field with the application ID.
	SubAppId pulumi.IntPtrInput
	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Width pulumi.IntPtrInput
}

func (SampleSnapshotTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*sampleSnapshotTemplateState)(nil)).Elem()
}

type sampleSnapshotTemplateArgs struct {
	// Template description. Length limit: 256 characters.
	Comment *string `pulumi:"comment"`
	// Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source
	// video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to
	// match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black.
	// This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black
	// color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and
	// fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio
	// of the source video for the screenshot and fills the unmatched area with Gaussian blur.Default value: black.
	FillType *string `pulumi:"fillType"`
	// Image format. Valid values: jpg, png. Default value: jpg.
	Format *string `pulumi:"format"`
	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Height *int `pulumi:"height"`
	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name *string `pulumi:"name"`
	// Resolution adaption. Valid values: open: enabled. In this case, `Width` represents the long side of a video, while
	// `Height` the short side; close: disabled. In this case, `Width` represents the width of a video, while `Height` the
	// height.Default value: open.
	ResolutionAdaptive *string `pulumi:"resolutionAdaptive"`
	// Sampling interval. If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.
	// If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
	SampleInterval int `pulumi:"sampleInterval"`
	// Sampled screencapturing type. Valid values: Percent: by percent. Time: by time interval.
	SampleType string `pulumi:"sampleType"`
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
	// service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
	// application or a newly created one), they must fill in this field with the application ID.
	SubAppId int `pulumi:"subAppId"`
	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a SampleSnapshotTemplate resource.
type SampleSnapshotTemplateArgs struct {
	// Template description. Length limit: 256 characters.
	Comment pulumi.StringPtrInput
	// Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source
	// video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to
	// match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black.
	// This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black
	// color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and
	// fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio
	// of the source video for the screenshot and fills the unmatched area with Gaussian blur.Default value: black.
	FillType pulumi.StringPtrInput
	// Image format. Valid values: jpg, png. Default value: jpg.
	Format pulumi.StringPtrInput
	// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Height pulumi.IntPtrInput
	// Name of a sampled screencapturing template. Length limit: 64 characters.
	Name pulumi.StringPtrInput
	// Resolution adaption. Valid values: open: enabled. In this case, `Width` represents the long side of a video, while
	// `Height` the short side; close: disabled. In this case, `Width` represents the width of a video, while `Height` the
	// height.Default value: open.
	ResolutionAdaptive pulumi.StringPtrInput
	// Sampling interval. If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.
	// If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
	SampleInterval pulumi.IntInput
	// Sampled screencapturing type. Valid values: Percent: by percent. Time: by time interval.
	SampleType pulumi.StringInput
	// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
	// service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
	// application or a newly created one), they must fill in this field with the application ID.
	SubAppId pulumi.IntInput
	// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
	// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
	// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
	// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
	Width pulumi.IntPtrInput
}

func (SampleSnapshotTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sampleSnapshotTemplateArgs)(nil)).Elem()
}

type SampleSnapshotTemplateInput interface {
	pulumi.Input

	ToSampleSnapshotTemplateOutput() SampleSnapshotTemplateOutput
	ToSampleSnapshotTemplateOutputWithContext(ctx context.Context) SampleSnapshotTemplateOutput
}

func (*SampleSnapshotTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**SampleSnapshotTemplate)(nil)).Elem()
}

func (i *SampleSnapshotTemplate) ToSampleSnapshotTemplateOutput() SampleSnapshotTemplateOutput {
	return i.ToSampleSnapshotTemplateOutputWithContext(context.Background())
}

func (i *SampleSnapshotTemplate) ToSampleSnapshotTemplateOutputWithContext(ctx context.Context) SampleSnapshotTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SampleSnapshotTemplateOutput)
}

// SampleSnapshotTemplateArrayInput is an input type that accepts SampleSnapshotTemplateArray and SampleSnapshotTemplateArrayOutput values.
// You can construct a concrete instance of `SampleSnapshotTemplateArrayInput` via:
//
//	SampleSnapshotTemplateArray{ SampleSnapshotTemplateArgs{...} }
type SampleSnapshotTemplateArrayInput interface {
	pulumi.Input

	ToSampleSnapshotTemplateArrayOutput() SampleSnapshotTemplateArrayOutput
	ToSampleSnapshotTemplateArrayOutputWithContext(context.Context) SampleSnapshotTemplateArrayOutput
}

type SampleSnapshotTemplateArray []SampleSnapshotTemplateInput

func (SampleSnapshotTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SampleSnapshotTemplate)(nil)).Elem()
}

func (i SampleSnapshotTemplateArray) ToSampleSnapshotTemplateArrayOutput() SampleSnapshotTemplateArrayOutput {
	return i.ToSampleSnapshotTemplateArrayOutputWithContext(context.Background())
}

func (i SampleSnapshotTemplateArray) ToSampleSnapshotTemplateArrayOutputWithContext(ctx context.Context) SampleSnapshotTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SampleSnapshotTemplateArrayOutput)
}

// SampleSnapshotTemplateMapInput is an input type that accepts SampleSnapshotTemplateMap and SampleSnapshotTemplateMapOutput values.
// You can construct a concrete instance of `SampleSnapshotTemplateMapInput` via:
//
//	SampleSnapshotTemplateMap{ "key": SampleSnapshotTemplateArgs{...} }
type SampleSnapshotTemplateMapInput interface {
	pulumi.Input

	ToSampleSnapshotTemplateMapOutput() SampleSnapshotTemplateMapOutput
	ToSampleSnapshotTemplateMapOutputWithContext(context.Context) SampleSnapshotTemplateMapOutput
}

type SampleSnapshotTemplateMap map[string]SampleSnapshotTemplateInput

func (SampleSnapshotTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SampleSnapshotTemplate)(nil)).Elem()
}

func (i SampleSnapshotTemplateMap) ToSampleSnapshotTemplateMapOutput() SampleSnapshotTemplateMapOutput {
	return i.ToSampleSnapshotTemplateMapOutputWithContext(context.Background())
}

func (i SampleSnapshotTemplateMap) ToSampleSnapshotTemplateMapOutputWithContext(ctx context.Context) SampleSnapshotTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SampleSnapshotTemplateMapOutput)
}

type SampleSnapshotTemplateOutput struct{ *pulumi.OutputState }

func (SampleSnapshotTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SampleSnapshotTemplate)(nil)).Elem()
}

func (o SampleSnapshotTemplateOutput) ToSampleSnapshotTemplateOutput() SampleSnapshotTemplateOutput {
	return o
}

func (o SampleSnapshotTemplateOutput) ToSampleSnapshotTemplateOutputWithContext(ctx context.Context) SampleSnapshotTemplateOutput {
	return o
}

// Template description. Length limit: 256 characters.
func (o SampleSnapshotTemplateOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Fill type. Fill refers to the way of processing a screenshot when its aspect ratio is different from that of the source
// video. The following fill types are supported: stretch: stretch. The screenshot will be stretched frame by frame to
// match the aspect ratio of the source video, which may make the screenshot shorter or longer; black: fill with black.
// This option retains the aspect ratio of the source video for the screenshot and fills the unmatched area with black
// color blocks. white: fill with white. This option retains the aspect ratio of the source video for the screenshot and
// fills the unmatched area with white color blocks. gauss: fill with Gaussian blur. This option retains the aspect ratio
// of the source video for the screenshot and fills the unmatched area with Gaussian blur.Default value: black.
func (o SampleSnapshotTemplateOutput) FillType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.FillType }).(pulumi.StringPtrOutput)
}

// Image format. Valid values: jpg, png. Default value: jpg.
func (o SampleSnapshotTemplateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// Maximum value of the height (or short side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
func (o SampleSnapshotTemplateOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// Name of a sampled screencapturing template. Length limit: 64 characters.
func (o SampleSnapshotTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resolution adaption. Valid values: open: enabled. In this case, `Width` represents the long side of a video, while
// `Height` the short side; close: disabled. In this case, `Width` represents the width of a video, while `Height` the
// height.Default value: open.
func (o SampleSnapshotTemplateOutput) ResolutionAdaptive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringPtrOutput { return v.ResolutionAdaptive }).(pulumi.StringPtrOutput)
}

// Sampling interval. If `SampleType` is `Percent`, sampling will be performed at an interval of the specified percentage.
// If `SampleType` is `Time`, sampling will be performed at the specified time interval in seconds.
func (o SampleSnapshotTemplateOutput) SampleInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntOutput { return v.SampleInterval }).(pulumi.IntOutput)
}

// Sampled screencapturing type. Valid values: Percent: by percent. Time: by time interval.
func (o SampleSnapshotTemplateOutput) SampleType() pulumi.StringOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.StringOutput { return v.SampleType }).(pulumi.StringOutput)
}

// The VOD [application](https://intl.cloud.tencent.com/document/product/266/14574) ID. For customers who activate VOD
// service from December 25, 2023, if they want to access resources in a VOD application (whether it's the default
// application or a newly created one), they must fill in this field with the application ID.
func (o SampleSnapshotTemplateOutput) SubAppId() pulumi.IntOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntOutput { return v.SubAppId }).(pulumi.IntOutput)
}

// Maximum value of the width (or long side) of a screenshot in px. Value range: 0 and [128, 4,096]. If both `Width` and
// `Height` are 0, the resolution will be the same as that of the source video; If `Width` is 0, but `Height` is not 0,
// `Width` will be proportionally scaled; If `Width` is not 0, but `Height` is 0, `Height` will be proportionally scaled;
// If both `Width` and `Height` are not 0, the custom resolution will be used.Default value: 0.
func (o SampleSnapshotTemplateOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SampleSnapshotTemplate) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type SampleSnapshotTemplateArrayOutput struct{ *pulumi.OutputState }

func (SampleSnapshotTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SampleSnapshotTemplate)(nil)).Elem()
}

func (o SampleSnapshotTemplateArrayOutput) ToSampleSnapshotTemplateArrayOutput() SampleSnapshotTemplateArrayOutput {
	return o
}

func (o SampleSnapshotTemplateArrayOutput) ToSampleSnapshotTemplateArrayOutputWithContext(ctx context.Context) SampleSnapshotTemplateArrayOutput {
	return o
}

func (o SampleSnapshotTemplateArrayOutput) Index(i pulumi.IntInput) SampleSnapshotTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SampleSnapshotTemplate {
		return vs[0].([]*SampleSnapshotTemplate)[vs[1].(int)]
	}).(SampleSnapshotTemplateOutput)
}

type SampleSnapshotTemplateMapOutput struct{ *pulumi.OutputState }

func (SampleSnapshotTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SampleSnapshotTemplate)(nil)).Elem()
}

func (o SampleSnapshotTemplateMapOutput) ToSampleSnapshotTemplateMapOutput() SampleSnapshotTemplateMapOutput {
	return o
}

func (o SampleSnapshotTemplateMapOutput) ToSampleSnapshotTemplateMapOutputWithContext(ctx context.Context) SampleSnapshotTemplateMapOutput {
	return o
}

func (o SampleSnapshotTemplateMapOutput) MapIndex(k pulumi.StringInput) SampleSnapshotTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SampleSnapshotTemplate {
		return vs[0].(map[string]*SampleSnapshotTemplate)[vs[1].(string)]
	}).(SampleSnapshotTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SampleSnapshotTemplateInput)(nil)).Elem(), &SampleSnapshotTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*SampleSnapshotTemplateArrayInput)(nil)).Elem(), SampleSnapshotTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SampleSnapshotTemplateMapInput)(nil)).Elem(), SampleSnapshotTemplateMap{})
	pulumi.RegisterOutputType(SampleSnapshotTemplateOutput{})
	pulumi.RegisterOutputType(SampleSnapshotTemplateArrayOutput{})
	pulumi.RegisterOutputType(SampleSnapshotTemplateMapOutput{})
}
