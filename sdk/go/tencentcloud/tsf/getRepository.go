// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryResult
	err := ctx.Invoke("tencentcloud:Tsf/getRepository:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryArgs struct {
	RepositoryType   *string `pulumi:"repositoryType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SearchWord       *string `pulumi:"searchWord"`
}

// A collection of values returned by getRepository.
type LookupRepositoryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                `pulumi:"id"`
	RepositoryType   *string               `pulumi:"repositoryType"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
	Results          []GetRepositoryResult `pulumi:"results"`
	SearchWord       *string               `pulumi:"searchWord"`
}

func LookupRepositoryOutput(ctx *pulumi.Context, args LookupRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryResult, error) {
			args := v.(LookupRepositoryArgs)
			r, err := LookupRepository(ctx, &args, opts...)
			var s LookupRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRepositoryResultOutput)
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryOutputArgs struct {
	RepositoryType   pulumi.StringPtrInput `pulumi:"repositoryType"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SearchWord       pulumi.StringPtrInput `pulumi:"searchWord"`
}

func (LookupRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRepository.
type LookupRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryResult)(nil)).Elem()
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutput() LookupRepositoryResultOutput {
	return o
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutputWithContext(ctx context.Context) LookupRepositoryResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRepositoryResultOutput) RepositoryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.RepositoryType }).(pulumi.StringPtrOutput)
}

func (o LookupRepositoryResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupRepositoryResultOutput) Results() GetRepositoryResultArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []GetRepositoryResult { return v.Results }).(GetRepositoryResultArrayOutput)
}

func (o LookupRepositoryResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryResultOutput{})
}