// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiUsagePlans(ctx *pulumi.Context, args *LookupApiUsagePlansArgs, opts ...pulumi.InvokeOption) (*LookupApiUsagePlansResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiUsagePlansResult
	err := ctx.Invoke("tencentcloud:ApiGateway/getApiUsagePlans:getApiUsagePlans", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiUsagePlans.
type LookupApiUsagePlansArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	ServiceId        string  `pulumi:"serviceId"`
}

// A collection of values returned by getApiUsagePlans.
type LookupApiUsagePlansResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                   `pulumi:"id"`
	ResultOutputFile *string                  `pulumi:"resultOutputFile"`
	Results          []GetApiUsagePlansResult `pulumi:"results"`
	ServiceId        string                   `pulumi:"serviceId"`
}

func LookupApiUsagePlansOutput(ctx *pulumi.Context, args LookupApiUsagePlansOutputArgs, opts ...pulumi.InvokeOption) LookupApiUsagePlansResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiUsagePlansResult, error) {
			args := v.(LookupApiUsagePlansArgs)
			r, err := LookupApiUsagePlans(ctx, &args, opts...)
			var s LookupApiUsagePlansResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiUsagePlansResultOutput)
}

// A collection of arguments for invoking getApiUsagePlans.
type LookupApiUsagePlansOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	ServiceId        pulumi.StringInput    `pulumi:"serviceId"`
}

func (LookupApiUsagePlansOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiUsagePlansArgs)(nil)).Elem()
}

// A collection of values returned by getApiUsagePlans.
type LookupApiUsagePlansResultOutput struct{ *pulumi.OutputState }

func (LookupApiUsagePlansResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiUsagePlansResult)(nil)).Elem()
}

func (o LookupApiUsagePlansResultOutput) ToLookupApiUsagePlansResultOutput() LookupApiUsagePlansResultOutput {
	return o
}

func (o LookupApiUsagePlansResultOutput) ToLookupApiUsagePlansResultOutputWithContext(ctx context.Context) LookupApiUsagePlansResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApiUsagePlansResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiUsagePlansResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiUsagePlansResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiUsagePlansResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupApiUsagePlansResultOutput) Results() GetApiUsagePlansResultArrayOutput {
	return o.ApplyT(func(v LookupApiUsagePlansResult) []GetApiUsagePlansResult { return v.Results }).(GetApiUsagePlansResultArrayOutput)
}

func (o LookupApiUsagePlansResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiUsagePlansResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiUsagePlansResultOutput{})
}