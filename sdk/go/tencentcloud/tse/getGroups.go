// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tse

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGroups(ctx *pulumi.Context, args *LookupGroupsArgs, opts ...pulumi.InvokeOption) (*LookupGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupsResult
	err := ctx.Invoke("tencentcloud:Tse/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type LookupGroupsArgs struct {
	Filters          []GetGroupsFilter `pulumi:"filters"`
	GatewayId        string            `pulumi:"gatewayId"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
}

// A collection of values returned by getGroups.
type LookupGroupsResult struct {
	Filters   []GetGroupsFilter `pulumi:"filters"`
	GatewayId string            `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string            `pulumi:"id"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
	Results          []GetGroupsResult `pulumi:"results"`
}

func LookupGroupsOutput(ctx *pulumi.Context, args LookupGroupsOutputArgs, opts ...pulumi.InvokeOption) LookupGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupsResult, error) {
			args := v.(LookupGroupsArgs)
			r, err := LookupGroups(ctx, &args, opts...)
			var s LookupGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type LookupGroupsOutputArgs struct {
	Filters          GetGroupsFilterArrayInput `pulumi:"filters"`
	GatewayId        pulumi.StringInput        `pulumi:"gatewayId"`
	ResultOutputFile pulumi.StringPtrInput     `pulumi:"resultOutputFile"`
}

func (LookupGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type LookupGroupsResultOutput struct{ *pulumi.OutputState }

func (LookupGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupsResult)(nil)).Elem()
}

func (o LookupGroupsResultOutput) ToLookupGroupsResultOutput() LookupGroupsResultOutput {
	return o
}

func (o LookupGroupsResultOutput) ToLookupGroupsResultOutputWithContext(ctx context.Context) LookupGroupsResultOutput {
	return o
}

func (o LookupGroupsResultOutput) Filters() GetGroupsFilterArrayOutput {
	return o.ApplyT(func(v LookupGroupsResult) []GetGroupsFilter { return v.Filters }).(GetGroupsFilterArrayOutput)
}

func (o LookupGroupsResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupsResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupGroupsResultOutput) Results() GetGroupsResultArrayOutput {
	return o.ApplyT(func(v LookupGroupsResult) []GetGroupsResult { return v.Results }).(GetGroupsResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupsResultOutput{})
}