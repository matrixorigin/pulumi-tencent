// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNamespaces(ctx *pulumi.Context, args *GetNamespacesArgs, opts ...pulumi.InvokeOption) (*GetNamespacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNamespacesResult
	err := ctx.Invoke("tencentcloud:Tcr/getNamespaces:getNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespaces.
type GetNamespacesArgs struct {
	InstanceId       string  `pulumi:"instanceId"`
	NamespaceName    *string `pulumi:"namespaceName"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getNamespaces.
type GetNamespacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                       `pulumi:"id"`
	InstanceId       string                       `pulumi:"instanceId"`
	NamespaceLists   []GetNamespacesNamespaceList `pulumi:"namespaceLists"`
	NamespaceName    *string                      `pulumi:"namespaceName"`
	ResultOutputFile *string                      `pulumi:"resultOutputFile"`
}

func GetNamespacesOutput(ctx *pulumi.Context, args GetNamespacesOutputArgs, opts ...pulumi.InvokeOption) GetNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNamespacesResult, error) {
			args := v.(GetNamespacesArgs)
			r, err := GetNamespaces(ctx, &args, opts...)
			var s GetNamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNamespacesResultOutput)
}

// A collection of arguments for invoking getNamespaces.
type GetNamespacesOutputArgs struct {
	InstanceId       pulumi.StringInput    `pulumi:"instanceId"`
	NamespaceName    pulumi.StringPtrInput `pulumi:"namespaceName"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesArgs)(nil)).Elem()
}

// A collection of values returned by getNamespaces.
type GetNamespacesResultOutput struct{ *pulumi.OutputState }

func (GetNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesResult)(nil)).Elem()
}

func (o GetNamespacesResultOutput) ToGetNamespacesResultOutput() GetNamespacesResultOutput {
	return o
}

func (o GetNamespacesResultOutput) ToGetNamespacesResultOutputWithContext(ctx context.Context) GetNamespacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNamespacesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetNamespacesResultOutput) NamespaceLists() GetNamespacesNamespaceListArrayOutput {
	return o.ApplyT(func(v GetNamespacesResult) []GetNamespacesNamespaceList { return v.NamespaceLists }).(GetNamespacesNamespaceListArrayOutput)
}

func (o GetNamespacesResultOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNamespacesResult) *string { return v.NamespaceName }).(pulumi.StringPtrOutput)
}

func (o GetNamespacesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNamespacesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNamespacesResultOutput{})
}