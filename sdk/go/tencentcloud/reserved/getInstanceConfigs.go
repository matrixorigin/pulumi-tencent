// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package reserved

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetInstanceConfigs(ctx *pulumi.Context, args *GetInstanceConfigsArgs, opts ...pulumi.InvokeOption) (*GetInstanceConfigsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceConfigsResult
	err := ctx.Invoke("tencentcloud:Reserved/getInstanceConfigs:getInstanceConfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceConfigs.
type GetInstanceConfigsArgs struct {
	AvailabilityZone   *string `pulumi:"availabilityZone"`
	Duration           *int    `pulumi:"duration"`
	InstanceType       *string `pulumi:"instanceType"`
	OfferingType       *string `pulumi:"offeringType"`
	ProductDescription *string `pulumi:"productDescription"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceConfigs.
type GetInstanceConfigsResult struct {
	AvailabilityZone *string                        `pulumi:"availabilityZone"`
	ConfigLists      []GetInstanceConfigsConfigList `pulumi:"configLists"`
	Duration         *int                           `pulumi:"duration"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	InstanceType       *string `pulumi:"instanceType"`
	OfferingType       *string `pulumi:"offeringType"`
	ProductDescription *string `pulumi:"productDescription"`
	ResultOutputFile   *string `pulumi:"resultOutputFile"`
}

func GetInstanceConfigsOutput(ctx *pulumi.Context, args GetInstanceConfigsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceConfigsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceConfigsResult, error) {
			args := v.(GetInstanceConfigsArgs)
			r, err := GetInstanceConfigs(ctx, &args, opts...)
			var s GetInstanceConfigsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceConfigsResultOutput)
}

// A collection of arguments for invoking getInstanceConfigs.
type GetInstanceConfigsOutputArgs struct {
	AvailabilityZone   pulumi.StringPtrInput `pulumi:"availabilityZone"`
	Duration           pulumi.IntPtrInput    `pulumi:"duration"`
	InstanceType       pulumi.StringPtrInput `pulumi:"instanceType"`
	OfferingType       pulumi.StringPtrInput `pulumi:"offeringType"`
	ProductDescription pulumi.StringPtrInput `pulumi:"productDescription"`
	ResultOutputFile   pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceConfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConfigsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceConfigs.
type GetInstanceConfigsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceConfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceConfigsResult)(nil)).Elem()
}

func (o GetInstanceConfigsResultOutput) ToGetInstanceConfigsResultOutput() GetInstanceConfigsResultOutput {
	return o
}

func (o GetInstanceConfigsResultOutput) ToGetInstanceConfigsResultOutputWithContext(ctx context.Context) GetInstanceConfigsResultOutput {
	return o
}

func (o GetInstanceConfigsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetInstanceConfigsResultOutput) ConfigLists() GetInstanceConfigsConfigListArrayOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) []GetInstanceConfigsConfigList { return v.ConfigLists }).(GetInstanceConfigsConfigListArrayOutput)
}

func (o GetInstanceConfigsResultOutput) Duration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *int { return v.Duration }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceConfigsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstanceConfigsResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceConfigsResultOutput) OfferingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.OfferingType }).(pulumi.StringPtrOutput)
}

func (o GetInstanceConfigsResultOutput) ProductDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.ProductDescription }).(pulumi.StringPtrOutput)
}

func (o GetInstanceConfigsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceConfigsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceConfigsResultOutput{})
}