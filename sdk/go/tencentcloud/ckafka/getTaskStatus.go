// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTaskStatus(ctx *pulumi.Context, args *LookupTaskStatusArgs, opts ...pulumi.InvokeOption) (*LookupTaskStatusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTaskStatusResult
	err := ctx.Invoke("tencentcloud:Ckafka/getTaskStatus:getTaskStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTaskStatus.
type LookupTaskStatusArgs struct {
	FlowId           int     `pulumi:"flowId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getTaskStatus.
type LookupTaskStatusResult struct {
	FlowId int `pulumi:"flowId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                `pulumi:"id"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
	Results          []GetTaskStatusResult `pulumi:"results"`
}

func LookupTaskStatusOutput(ctx *pulumi.Context, args LookupTaskStatusOutputArgs, opts ...pulumi.InvokeOption) LookupTaskStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTaskStatusResult, error) {
			args := v.(LookupTaskStatusArgs)
			r, err := LookupTaskStatus(ctx, &args, opts...)
			var s LookupTaskStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTaskStatusResultOutput)
}

// A collection of arguments for invoking getTaskStatus.
type LookupTaskStatusOutputArgs struct {
	FlowId           pulumi.IntInput       `pulumi:"flowId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (LookupTaskStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskStatusArgs)(nil)).Elem()
}

// A collection of values returned by getTaskStatus.
type LookupTaskStatusResultOutput struct{ *pulumi.OutputState }

func (LookupTaskStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTaskStatusResult)(nil)).Elem()
}

func (o LookupTaskStatusResultOutput) ToLookupTaskStatusResultOutput() LookupTaskStatusResultOutput {
	return o
}

func (o LookupTaskStatusResultOutput) ToLookupTaskStatusResultOutputWithContext(ctx context.Context) LookupTaskStatusResultOutput {
	return o
}

func (o LookupTaskStatusResultOutput) FlowId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTaskStatusResult) int { return v.FlowId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTaskStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTaskStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTaskStatusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTaskStatusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupTaskStatusResultOutput) Results() GetTaskStatusResultArrayOutput {
	return o.ApplyT(func(v LookupTaskStatusResult) []GetTaskStatusResult { return v.Results }).(GetTaskStatusResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTaskStatusResultOutput{})
}