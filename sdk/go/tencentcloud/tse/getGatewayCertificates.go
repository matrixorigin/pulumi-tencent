// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tse

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGatewayCertificates(ctx *pulumi.Context, args *LookupGatewayCertificatesArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayCertificatesResult
	err := ctx.Invoke("tencentcloud:Tse/getGatewayCertificates:getGatewayCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayCertificates.
type LookupGatewayCertificatesArgs struct {
	Filters          []GetGatewayCertificatesFilter `pulumi:"filters"`
	GatewayId        string                         `pulumi:"gatewayId"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

// A collection of values returned by getGatewayCertificates.
type LookupGatewayCertificatesResult struct {
	Filters   []GetGatewayCertificatesFilter `pulumi:"filters"`
	GatewayId string                         `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
	Results          []GetGatewayCertificatesResult `pulumi:"results"`
}

func LookupGatewayCertificatesOutput(ctx *pulumi.Context, args LookupGatewayCertificatesOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCertificatesResult, error) {
			args := v.(LookupGatewayCertificatesArgs)
			r, err := LookupGatewayCertificates(ctx, &args, opts...)
			var s LookupGatewayCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCertificatesResultOutput)
}

// A collection of arguments for invoking getGatewayCertificates.
type LookupGatewayCertificatesOutputArgs struct {
	Filters          GetGatewayCertificatesFilterArrayInput `pulumi:"filters"`
	GatewayId        pulumi.StringInput                     `pulumi:"gatewayId"`
	ResultOutputFile pulumi.StringPtrInput                  `pulumi:"resultOutputFile"`
}

func (LookupGatewayCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayCertificates.
type LookupGatewayCertificatesResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCertificatesResult)(nil)).Elem()
}

func (o LookupGatewayCertificatesResultOutput) ToLookupGatewayCertificatesResultOutput() LookupGatewayCertificatesResultOutput {
	return o
}

func (o LookupGatewayCertificatesResultOutput) ToLookupGatewayCertificatesResultOutputWithContext(ctx context.Context) LookupGatewayCertificatesResultOutput {
	return o
}

func (o LookupGatewayCertificatesResultOutput) Filters() GetGatewayCertificatesFilterArrayOutput {
	return o.ApplyT(func(v LookupGatewayCertificatesResult) []GetGatewayCertificatesFilter { return v.Filters }).(GetGatewayCertificatesFilterArrayOutput)
}

func (o LookupGatewayCertificatesResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificatesResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGatewayCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayCertificatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayCertificatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayCertificatesResultOutput) Results() GetGatewayCertificatesResultArrayOutput {
	return o.ApplyT(func(v LookupGatewayCertificatesResult) []GetGatewayCertificatesResult { return v.Results }).(GetGatewayCertificatesResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCertificatesResultOutput{})
}