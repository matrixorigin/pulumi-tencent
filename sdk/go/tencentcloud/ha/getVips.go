// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ha

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVips(ctx *pulumi.Context, args *GetVipsArgs, opts ...pulumi.InvokeOption) (*GetVipsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVipsResult
	err := ctx.Invoke("tencentcloud:Ha/getVips:getVips", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVips.
type GetVipsArgs struct {
	AddressIp        *string `pulumi:"addressIp"`
	Id               *string `pulumi:"id"`
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SubnetId         *string `pulumi:"subnetId"`
	VpcId            *string `pulumi:"vpcId"`
}

// A collection of values returned by getVips.
type GetVipsResult struct {
	AddressIp        *string            `pulumi:"addressIp"`
	HaVipLists       []GetVipsHaVipList `pulumi:"haVipLists"`
	Id               *string            `pulumi:"id"`
	Name             *string            `pulumi:"name"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	SubnetId         *string            `pulumi:"subnetId"`
	VpcId            *string            `pulumi:"vpcId"`
}

func GetVipsOutput(ctx *pulumi.Context, args GetVipsOutputArgs, opts ...pulumi.InvokeOption) GetVipsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVipsResult, error) {
			args := v.(GetVipsArgs)
			r, err := GetVips(ctx, &args, opts...)
			var s GetVipsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVipsResultOutput)
}

// A collection of arguments for invoking getVips.
type GetVipsOutputArgs struct {
	AddressIp        pulumi.StringPtrInput `pulumi:"addressIp"`
	Id               pulumi.StringPtrInput `pulumi:"id"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
	VpcId            pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetVipsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVipsArgs)(nil)).Elem()
}

// A collection of values returned by getVips.
type GetVipsResultOutput struct{ *pulumi.OutputState }

func (GetVipsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVipsResult)(nil)).Elem()
}

func (o GetVipsResultOutput) ToGetVipsResultOutput() GetVipsResultOutput {
	return o
}

func (o GetVipsResultOutput) ToGetVipsResultOutputWithContext(ctx context.Context) GetVipsResultOutput {
	return o
}

func (o GetVipsResultOutput) AddressIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipsResult) *string { return v.AddressIp }).(pulumi.StringPtrOutput)
}

func (o GetVipsResultOutput) HaVipLists() GetVipsHaVipListArrayOutput {
	return o.ApplyT(func(v GetVipsResult) []GetVipsHaVipList { return v.HaVipLists }).(GetVipsHaVipListArrayOutput)
}

func (o GetVipsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetVipsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetVipsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetVipsResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipsResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o GetVipsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVipsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVipsResultOutput{})
}