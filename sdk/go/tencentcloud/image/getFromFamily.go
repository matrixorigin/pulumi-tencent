// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package image

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFromFamily(ctx *pulumi.Context, args *GetFromFamilyArgs, opts ...pulumi.InvokeOption) (*GetFromFamilyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFromFamilyResult
	err := ctx.Invoke("tencentcloud:Image/getFromFamily:getFromFamily", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFromFamily.
type GetFromFamilyArgs struct {
	ImageFamily      string  `pulumi:"imageFamily"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getFromFamily.
type GetFromFamilyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string               `pulumi:"id"`
	ImageFamily      string               `pulumi:"imageFamily"`
	Images           []GetFromFamilyImage `pulumi:"images"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
}

func GetFromFamilyOutput(ctx *pulumi.Context, args GetFromFamilyOutputArgs, opts ...pulumi.InvokeOption) GetFromFamilyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFromFamilyResult, error) {
			args := v.(GetFromFamilyArgs)
			r, err := GetFromFamily(ctx, &args, opts...)
			var s GetFromFamilyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFromFamilyResultOutput)
}

// A collection of arguments for invoking getFromFamily.
type GetFromFamilyOutputArgs struct {
	ImageFamily      pulumi.StringInput    `pulumi:"imageFamily"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetFromFamilyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFromFamilyArgs)(nil)).Elem()
}

// A collection of values returned by getFromFamily.
type GetFromFamilyResultOutput struct{ *pulumi.OutputState }

func (GetFromFamilyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFromFamilyResult)(nil)).Elem()
}

func (o GetFromFamilyResultOutput) ToGetFromFamilyResultOutput() GetFromFamilyResultOutput {
	return o
}

func (o GetFromFamilyResultOutput) ToGetFromFamilyResultOutputWithContext(ctx context.Context) GetFromFamilyResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetFromFamilyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFromFamilyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFromFamilyResultOutput) ImageFamily() pulumi.StringOutput {
	return o.ApplyT(func(v GetFromFamilyResult) string { return v.ImageFamily }).(pulumi.StringOutput)
}

func (o GetFromFamilyResultOutput) Images() GetFromFamilyImageArrayOutput {
	return o.ApplyT(func(v GetFromFamilyResult) []GetFromFamilyImage { return v.Images }).(GetFromFamilyImageArrayOutput)
}

func (o GetFromFamilyResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFromFamilyResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFromFamilyResultOutput{})
}
