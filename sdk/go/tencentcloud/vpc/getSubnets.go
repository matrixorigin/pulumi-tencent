// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSubnets(ctx *pulumi.Context, args *GetSubnetsArgs, opts ...pulumi.InvokeOption) (*GetSubnetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSubnetsResult
	err := ctx.Invoke("tencentcloud:Vpc/getSubnets:getSubnets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsArgs struct {
	AvailabilityZone *string                `pulumi:"availabilityZone"`
	CidrBlock        *string                `pulumi:"cidrBlock"`
	IsDefault        *bool                  `pulumi:"isDefault"`
	IsRemoteVpcSnat  *bool                  `pulumi:"isRemoteVpcSnat"`
	Name             *string                `pulumi:"name"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
	SubnetId         *string                `pulumi:"subnetId"`
	TagKey           *string                `pulumi:"tagKey"`
	Tags             map[string]interface{} `pulumi:"tags"`
	VpcId            *string                `pulumi:"vpcId"`
}

// A collection of values returned by getSubnets.
type GetSubnetsResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	CidrBlock        *string `pulumi:"cidrBlock"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                   `pulumi:"id"`
	InstanceLists    []GetSubnetsInstanceList `pulumi:"instanceLists"`
	IsDefault        *bool                    `pulumi:"isDefault"`
	IsRemoteVpcSnat  *bool                    `pulumi:"isRemoteVpcSnat"`
	Name             *string                  `pulumi:"name"`
	ResultOutputFile *string                  `pulumi:"resultOutputFile"`
	SubnetId         *string                  `pulumi:"subnetId"`
	TagKey           *string                  `pulumi:"tagKey"`
	Tags             map[string]interface{}   `pulumi:"tags"`
	VpcId            *string                  `pulumi:"vpcId"`
}

func GetSubnetsOutput(ctx *pulumi.Context, args GetSubnetsOutputArgs, opts ...pulumi.InvokeOption) GetSubnetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubnetsResult, error) {
			args := v.(GetSubnetsArgs)
			r, err := GetSubnets(ctx, &args, opts...)
			var s GetSubnetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubnetsResultOutput)
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsOutputArgs struct {
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	CidrBlock        pulumi.StringPtrInput `pulumi:"cidrBlock"`
	IsDefault        pulumi.BoolPtrInput   `pulumi:"isDefault"`
	IsRemoteVpcSnat  pulumi.BoolPtrInput   `pulumi:"isRemoteVpcSnat"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SubnetId         pulumi.StringPtrInput `pulumi:"subnetId"`
	TagKey           pulumi.StringPtrInput `pulumi:"tagKey"`
	Tags             pulumi.MapInput       `pulumi:"tags"`
	VpcId            pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetSubnetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsArgs)(nil)).Elem()
}

// A collection of values returned by getSubnets.
type GetSubnetsResultOutput struct{ *pulumi.OutputState }

func (GetSubnetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsResult)(nil)).Elem()
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutput() GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutputWithContext(ctx context.Context) GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSubnetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSubnetsResultOutput) InstanceLists() GetSubnetsInstanceListArrayOutput {
	return o.ApplyT(func(v GetSubnetsResult) []GetSubnetsInstanceList { return v.InstanceLists }).(GetSubnetsInstanceListArrayOutput)
}

func (o GetSubnetsResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o GetSubnetsResultOutput) IsRemoteVpcSnat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *bool { return v.IsRemoteVpcSnat }).(pulumi.BoolPtrOutput)
}

func (o GetSubnetsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) TagKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.TagKey }).(pulumi.StringPtrOutput)
}

func (o GetSubnetsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetSubnetsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o GetSubnetsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSubnetsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetsResultOutput{})
}