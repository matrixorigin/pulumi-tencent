// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetListenerStatistics(ctx *pulumi.Context, args *GetListenerStatisticsArgs, opts ...pulumi.InvokeOption) (*GetListenerStatisticsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetListenerStatisticsResult
	err := ctx.Invoke("tencentcloud:Gaap/getListenerStatistics:getListenerStatistics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getListenerStatistics.
type GetListenerStatisticsArgs struct {
	EndTime          string   `pulumi:"endTime"`
	Granularity      int      `pulumi:"granularity"`
	ListenerId       string   `pulumi:"listenerId"`
	MetricNames      []string `pulumi:"metricNames"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	StartTime        string   `pulumi:"startTime"`
}

// A collection of values returned by getListenerStatistics.
type GetListenerStatisticsResult struct {
	EndTime     string `pulumi:"endTime"`
	Granularity int    `pulumi:"granularity"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                `pulumi:"id"`
	ListenerId       string                                `pulumi:"listenerId"`
	MetricNames      []string                              `pulumi:"metricNames"`
	ResultOutputFile *string                               `pulumi:"resultOutputFile"`
	StartTime        string                                `pulumi:"startTime"`
	StatisticsDatas  []GetListenerStatisticsStatisticsData `pulumi:"statisticsDatas"`
}

func GetListenerStatisticsOutput(ctx *pulumi.Context, args GetListenerStatisticsOutputArgs, opts ...pulumi.InvokeOption) GetListenerStatisticsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetListenerStatisticsResult, error) {
			args := v.(GetListenerStatisticsArgs)
			r, err := GetListenerStatistics(ctx, &args, opts...)
			var s GetListenerStatisticsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetListenerStatisticsResultOutput)
}

// A collection of arguments for invoking getListenerStatistics.
type GetListenerStatisticsOutputArgs struct {
	EndTime          pulumi.StringInput      `pulumi:"endTime"`
	Granularity      pulumi.IntInput         `pulumi:"granularity"`
	ListenerId       pulumi.StringInput      `pulumi:"listenerId"`
	MetricNames      pulumi.StringArrayInput `pulumi:"metricNames"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	StartTime        pulumi.StringInput      `pulumi:"startTime"`
}

func (GetListenerStatisticsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListenerStatisticsArgs)(nil)).Elem()
}

// A collection of values returned by getListenerStatistics.
type GetListenerStatisticsResultOutput struct{ *pulumi.OutputState }

func (GetListenerStatisticsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetListenerStatisticsResult)(nil)).Elem()
}

func (o GetListenerStatisticsResultOutput) ToGetListenerStatisticsResultOutput() GetListenerStatisticsResultOutput {
	return o
}

func (o GetListenerStatisticsResultOutput) ToGetListenerStatisticsResultOutputWithContext(ctx context.Context) GetListenerStatisticsResultOutput {
	return o
}

func (o GetListenerStatisticsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o GetListenerStatisticsResultOutput) Granularity() pulumi.IntOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) int { return v.Granularity }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetListenerStatisticsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetListenerStatisticsResultOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) string { return v.ListenerId }).(pulumi.StringOutput)
}

func (o GetListenerStatisticsResultOutput) MetricNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) []string { return v.MetricNames }).(pulumi.StringArrayOutput)
}

func (o GetListenerStatisticsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetListenerStatisticsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o GetListenerStatisticsResultOutput) StatisticsDatas() GetListenerStatisticsStatisticsDataArrayOutput {
	return o.ApplyT(func(v GetListenerStatisticsResult) []GetListenerStatisticsStatisticsData { return v.StatisticsDatas }).(GetListenerStatisticsStatisticsDataArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetListenerStatisticsResultOutput{})
}