// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRepositories(ctx *pulumi.Context, args *GetRepositoriesArgs, opts ...pulumi.InvokeOption) (*GetRepositoriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRepositoriesResult
	err := ctx.Invoke("tencentcloud:Tcr/getRepositories:getRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositories.
type GetRepositoriesArgs struct {
	InstanceId       string  `pulumi:"instanceId"`
	NamespaceName    string  `pulumi:"namespaceName"`
	RepositoryName   *string `pulumi:"repositoryName"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRepositories.
type GetRepositoriesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                          `pulumi:"id"`
	InstanceId       string                          `pulumi:"instanceId"`
	NamespaceName    string                          `pulumi:"namespaceName"`
	RepositoryLists  []GetRepositoriesRepositoryList `pulumi:"repositoryLists"`
	RepositoryName   *string                         `pulumi:"repositoryName"`
	ResultOutputFile *string                         `pulumi:"resultOutputFile"`
}

func GetRepositoriesOutput(ctx *pulumi.Context, args GetRepositoriesOutputArgs, opts ...pulumi.InvokeOption) GetRepositoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRepositoriesResult, error) {
			args := v.(GetRepositoriesArgs)
			r, err := GetRepositories(ctx, &args, opts...)
			var s GetRepositoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRepositoriesResultOutput)
}

// A collection of arguments for invoking getRepositories.
type GetRepositoriesOutputArgs struct {
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	NamespaceName    pulumi.StringInput    `pulumi:"namespaceName"`
	RepositoryName   pulumi.StringPtrInput `pulumi:"repositoryName"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetRepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesArgs)(nil)).Elem()
}

// A collection of values returned by getRepositories.
type GetRepositoriesResultOutput struct{ *pulumi.OutputState }

func (GetRepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRepositoriesResult)(nil)).Elem()
}

func (o GetRepositoriesResultOutput) ToGetRepositoriesResultOutput() GetRepositoriesResultOutput {
	return o
}

func (o GetRepositoriesResultOutput) ToGetRepositoriesResultOutputWithContext(ctx context.Context) GetRepositoriesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRepositoriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRepositoriesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetRepositoriesResultOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRepositoriesResult) string { return v.NamespaceName }).(pulumi.StringOutput)
}

func (o GetRepositoriesResultOutput) RepositoryLists() GetRepositoriesRepositoryListArrayOutput {
	return o.ApplyT(func(v GetRepositoriesResult) []GetRepositoriesRepositoryList { return v.RepositoryLists }).(GetRepositoriesRepositoryListArrayOutput)
}

func (o GetRepositoriesResultOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o GetRepositoriesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRepositoriesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRepositoriesResultOutput{})
}