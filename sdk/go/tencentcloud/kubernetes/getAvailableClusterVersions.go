// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAvailableClusterVersions(ctx *pulumi.Context, args *GetAvailableClusterVersionsArgs, opts ...pulumi.InvokeOption) (*GetAvailableClusterVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAvailableClusterVersionsResult
	err := ctx.Invoke("tencentcloud:Kubernetes/getAvailableClusterVersions:getAvailableClusterVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailableClusterVersions.
type GetAvailableClusterVersionsArgs struct {
	ClusterId        *string  `pulumi:"clusterId"`
	ClusterIds       []string `pulumi:"clusterIds"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAvailableClusterVersions.
type GetAvailableClusterVersionsResult struct {
	ClusterId  *string                              `pulumi:"clusterId"`
	ClusterIds []string                             `pulumi:"clusterIds"`
	Clusters   []GetAvailableClusterVersionsCluster `pulumi:"clusters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	ResultOutputFile *string  `pulumi:"resultOutputFile"`
	Versions         []string `pulumi:"versions"`
}

func GetAvailableClusterVersionsOutput(ctx *pulumi.Context, args GetAvailableClusterVersionsOutputArgs, opts ...pulumi.InvokeOption) GetAvailableClusterVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAvailableClusterVersionsResult, error) {
			args := v.(GetAvailableClusterVersionsArgs)
			r, err := GetAvailableClusterVersions(ctx, &args, opts...)
			var s GetAvailableClusterVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAvailableClusterVersionsResultOutput)
}

// A collection of arguments for invoking getAvailableClusterVersions.
type GetAvailableClusterVersionsOutputArgs struct {
	ClusterId        pulumi.StringPtrInput   `pulumi:"clusterId"`
	ClusterIds       pulumi.StringArrayInput `pulumi:"clusterIds"`
	ResultOutputFile pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
}

func (GetAvailableClusterVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailableClusterVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getAvailableClusterVersions.
type GetAvailableClusterVersionsResultOutput struct{ *pulumi.OutputState }

func (GetAvailableClusterVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailableClusterVersionsResult)(nil)).Elem()
}

func (o GetAvailableClusterVersionsResultOutput) ToGetAvailableClusterVersionsResultOutput() GetAvailableClusterVersionsResultOutput {
	return o
}

func (o GetAvailableClusterVersionsResultOutput) ToGetAvailableClusterVersionsResultOutputWithContext(ctx context.Context) GetAvailableClusterVersionsResultOutput {
	return o
}

func (o GetAvailableClusterVersionsResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailableClusterVersionsResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o GetAvailableClusterVersionsResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailableClusterVersionsResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

func (o GetAvailableClusterVersionsResultOutput) Clusters() GetAvailableClusterVersionsClusterArrayOutput {
	return o.ApplyT(func(v GetAvailableClusterVersionsResult) []GetAvailableClusterVersionsCluster { return v.Clusters }).(GetAvailableClusterVersionsClusterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailableClusterVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailableClusterVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAvailableClusterVersionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailableClusterVersionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAvailableClusterVersionsResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailableClusterVersionsResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailableClusterVersionsResultOutput{})
}