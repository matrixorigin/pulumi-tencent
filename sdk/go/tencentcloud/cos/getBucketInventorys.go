// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cos

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBucketInventorys(ctx *pulumi.Context, args *GetBucketInventorysArgs, opts ...pulumi.InvokeOption) (*GetBucketInventorysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBucketInventorysResult
	err := ctx.Invoke("tencentcloud:Cos/getBucketInventorys:getBucketInventorys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketInventorys.
type GetBucketInventorysArgs struct {
	Bucket           string  `pulumi:"bucket"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBucketInventorys.
type GetBucketInventorysResult struct {
	Bucket string `pulumi:"bucket"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	Inventorys       []GetBucketInventorysInventory `pulumi:"inventorys"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

func GetBucketInventorysOutput(ctx *pulumi.Context, args GetBucketInventorysOutputArgs, opts ...pulumi.InvokeOption) GetBucketInventorysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBucketInventorysResult, error) {
			args := v.(GetBucketInventorysArgs)
			r, err := GetBucketInventorys(ctx, &args, opts...)
			var s GetBucketInventorysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBucketInventorysResultOutput)
}

// A collection of arguments for invoking getBucketInventorys.
type GetBucketInventorysOutputArgs struct {
	Bucket           pulumi.StringInput    `pulumi:"bucket"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetBucketInventorysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketInventorysArgs)(nil)).Elem()
}

// A collection of values returned by getBucketInventorys.
type GetBucketInventorysResultOutput struct{ *pulumi.OutputState }

func (GetBucketInventorysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketInventorysResult)(nil)).Elem()
}

func (o GetBucketInventorysResultOutput) ToGetBucketInventorysResultOutput() GetBucketInventorysResultOutput {
	return o
}

func (o GetBucketInventorysResultOutput) ToGetBucketInventorysResultOutputWithContext(ctx context.Context) GetBucketInventorysResultOutput {
	return o
}

func (o GetBucketInventorysResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketInventorysResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBucketInventorysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketInventorysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBucketInventorysResultOutput) Inventorys() GetBucketInventorysInventoryArrayOutput {
	return o.ApplyT(func(v GetBucketInventorysResult) []GetBucketInventorysInventory { return v.Inventorys }).(GetBucketInventorysInventoryArrayOutput)
}

func (o GetBucketInventorysResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketInventorysResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBucketInventorysResultOutput{})
}