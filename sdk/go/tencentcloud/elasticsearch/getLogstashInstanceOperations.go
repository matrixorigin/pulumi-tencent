// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogstashInstanceOperations(ctx *pulumi.Context, args *GetLogstashInstanceOperationsArgs, opts ...pulumi.InvokeOption) (*GetLogstashInstanceOperationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogstashInstanceOperationsResult
	err := ctx.Invoke("tencentcloud:Elasticsearch/getLogstashInstanceOperations:getLogstashInstanceOperations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogstashInstanceOperations.
type GetLogstashInstanceOperationsArgs struct {
	EndTime          string  `pulumi:"endTime"`
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        string  `pulumi:"startTime"`
}

// A collection of values returned by getLogstashInstanceOperations.
type GetLogstashInstanceOperationsResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                   `pulumi:"id"`
	InstanceId       string                                   `pulumi:"instanceId"`
	Operations       []GetLogstashInstanceOperationsOperation `pulumi:"operations"`
	ResultOutputFile *string                                  `pulumi:"resultOutputFile"`
	StartTime        string                                   `pulumi:"startTime"`
}

func GetLogstashInstanceOperationsOutput(ctx *pulumi.Context, args GetLogstashInstanceOperationsOutputArgs, opts ...pulumi.InvokeOption) GetLogstashInstanceOperationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogstashInstanceOperationsResult, error) {
			args := v.(GetLogstashInstanceOperationsArgs)
			r, err := GetLogstashInstanceOperations(ctx, &args, opts...)
			var s GetLogstashInstanceOperationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogstashInstanceOperationsResultOutput)
}

// A collection of arguments for invoking getLogstashInstanceOperations.
type GetLogstashInstanceOperationsOutputArgs struct {
	EndTime          pulumi.StringInput    `pulumi:"endTime"`
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	StartTime        pulumi.StringInput    `pulumi:"startTime"`
}

func (GetLogstashInstanceOperationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogstashInstanceOperationsArgs)(nil)).Elem()
}

// A collection of values returned by getLogstashInstanceOperations.
type GetLogstashInstanceOperationsResultOutput struct{ *pulumi.OutputState }

func (GetLogstashInstanceOperationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogstashInstanceOperationsResult)(nil)).Elem()
}

func (o GetLogstashInstanceOperationsResultOutput) ToGetLogstashInstanceOperationsResultOutput() GetLogstashInstanceOperationsResultOutput {
	return o
}

func (o GetLogstashInstanceOperationsResultOutput) ToGetLogstashInstanceOperationsResultOutputWithContext(ctx context.Context) GetLogstashInstanceOperationsResultOutput {
	return o
}

func (o GetLogstashInstanceOperationsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogstashInstanceOperationsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogstashInstanceOperationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogstashInstanceOperationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogstashInstanceOperationsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogstashInstanceOperationsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetLogstashInstanceOperationsResultOutput) Operations() GetLogstashInstanceOperationsOperationArrayOutput {
	return o.ApplyT(func(v GetLogstashInstanceOperationsResult) []GetLogstashInstanceOperationsOperation {
		return v.Operations
	}).(GetLogstashInstanceOperationsOperationArrayOutput)
}

func (o GetLogstashInstanceOperationsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogstashInstanceOperationsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetLogstashInstanceOperationsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogstashInstanceOperationsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogstashInstanceOperationsResultOutput{})
}