// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClusterNodePools(ctx *pulumi.Context, args *GetClusterNodePoolsArgs, opts ...pulumi.InvokeOption) (*GetClusterNodePoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetClusterNodePoolsResult
	err := ctx.Invoke("tencentcloud:Kubernetes/getClusterNodePools:getClusterNodePools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterNodePools.
type GetClusterNodePoolsArgs struct {
	ClusterId        string                      `pulumi:"clusterId"`
	Filters          []GetClusterNodePoolsFilter `pulumi:"filters"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
}

// A collection of values returned by getClusterNodePools.
type GetClusterNodePoolsResult struct {
	ClusterId string                      `pulumi:"clusterId"`
	Filters   []GetClusterNodePoolsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                           `pulumi:"id"`
	NodePoolSets     []GetClusterNodePoolsNodePoolSet `pulumi:"nodePoolSets"`
	ResultOutputFile *string                          `pulumi:"resultOutputFile"`
}

func GetClusterNodePoolsOutput(ctx *pulumi.Context, args GetClusterNodePoolsOutputArgs, opts ...pulumi.InvokeOption) GetClusterNodePoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClusterNodePoolsResult, error) {
			args := v.(GetClusterNodePoolsArgs)
			r, err := GetClusterNodePools(ctx, &args, opts...)
			var s GetClusterNodePoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClusterNodePoolsResultOutput)
}

// A collection of arguments for invoking getClusterNodePools.
type GetClusterNodePoolsOutputArgs struct {
	ClusterId        pulumi.StringInput                  `pulumi:"clusterId"`
	Filters          GetClusterNodePoolsFilterArrayInput `pulumi:"filters"`
	ResultOutputFile pulumi.StringPtrInput               `pulumi:"resultOutputFile"`
}

func (GetClusterNodePoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterNodePoolsArgs)(nil)).Elem()
}

// A collection of values returned by getClusterNodePools.
type GetClusterNodePoolsResultOutput struct{ *pulumi.OutputState }

func (GetClusterNodePoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterNodePoolsResult)(nil)).Elem()
}

func (o GetClusterNodePoolsResultOutput) ToGetClusterNodePoolsResultOutput() GetClusterNodePoolsResultOutput {
	return o
}

func (o GetClusterNodePoolsResultOutput) ToGetClusterNodePoolsResultOutputWithContext(ctx context.Context) GetClusterNodePoolsResultOutput {
	return o
}

func (o GetClusterNodePoolsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterNodePoolsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetClusterNodePoolsResultOutput) Filters() GetClusterNodePoolsFilterArrayOutput {
	return o.ApplyT(func(v GetClusterNodePoolsResult) []GetClusterNodePoolsFilter { return v.Filters }).(GetClusterNodePoolsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetClusterNodePoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterNodePoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClusterNodePoolsResultOutput) NodePoolSets() GetClusterNodePoolsNodePoolSetArrayOutput {
	return o.ApplyT(func(v GetClusterNodePoolsResult) []GetClusterNodePoolsNodePoolSet { return v.NodePoolSets }).(GetClusterNodePoolsNodePoolSetArrayOutput)
}

func (o GetClusterNodePoolsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetClusterNodePoolsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClusterNodePoolsResultOutput{})
}