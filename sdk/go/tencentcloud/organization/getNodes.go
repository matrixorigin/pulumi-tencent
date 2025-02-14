// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNodes(ctx *pulumi.Context, args *GetNodesArgs, opts ...pulumi.InvokeOption) (*GetNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNodesResult
	err := ctx.Invoke("tencentcloud:Organization/getNodes:getNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodes.
type GetNodesArgs struct {
	ResultOutputFile *string       `pulumi:"resultOutputFile"`
	Tags             []GetNodesTag `pulumi:"tags"`
}

// A collection of values returned by getNodes.
type GetNodesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string         `pulumi:"id"`
	Items            []GetNodesItem `pulumi:"items"`
	ResultOutputFile *string        `pulumi:"resultOutputFile"`
	Tags             []GetNodesTag  `pulumi:"tags"`
}

func GetNodesOutput(ctx *pulumi.Context, args GetNodesOutputArgs, opts ...pulumi.InvokeOption) GetNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNodesResult, error) {
			args := v.(GetNodesArgs)
			r, err := GetNodes(ctx, &args, opts...)
			var s GetNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNodesResultOutput)
}

// A collection of arguments for invoking getNodes.
type GetNodesOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	Tags             GetNodesTagArrayInput `pulumi:"tags"`
}

func (GetNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesArgs)(nil)).Elem()
}

// A collection of values returned by getNodes.
type GetNodesResultOutput struct{ *pulumi.OutputState }

func (GetNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNodesResult)(nil)).Elem()
}

func (o GetNodesResultOutput) ToGetNodesResultOutput() GetNodesResultOutput {
	return o
}

func (o GetNodesResultOutput) ToGetNodesResultOutputWithContext(ctx context.Context) GetNodesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNodesResultOutput) Items() GetNodesItemArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []GetNodesItem { return v.Items }).(GetNodesItemArrayOutput)
}

func (o GetNodesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNodesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetNodesResultOutput) Tags() GetNodesTagArrayOutput {
	return o.ApplyT(func(v GetNodesResult) []GetNodesTag { return v.Tags }).(GetNodesTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNodesResultOutput{})
}
