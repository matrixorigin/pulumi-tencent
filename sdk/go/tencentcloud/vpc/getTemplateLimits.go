// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTemplateLimits(ctx *pulumi.Context, args *GetTemplateLimitsArgs, opts ...pulumi.InvokeOption) (*GetTemplateLimitsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTemplateLimitsResult
	err := ctx.Invoke("tencentcloud:Vpc/getTemplateLimits:getTemplateLimits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplateLimits.
type GetTemplateLimitsArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getTemplateLimits.
type GetTemplateLimitsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                           `pulumi:"id"`
	ResultOutputFile *string                          `pulumi:"resultOutputFile"`
	TemplateLimits   []GetTemplateLimitsTemplateLimit `pulumi:"templateLimits"`
}

func GetTemplateLimitsOutput(ctx *pulumi.Context, args GetTemplateLimitsOutputArgs, opts ...pulumi.InvokeOption) GetTemplateLimitsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplateLimitsResult, error) {
			args := v.(GetTemplateLimitsArgs)
			r, err := GetTemplateLimits(ctx, &args, opts...)
			var s GetTemplateLimitsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTemplateLimitsResultOutput)
}

// A collection of arguments for invoking getTemplateLimits.
type GetTemplateLimitsOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetTemplateLimitsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateLimitsArgs)(nil)).Elem()
}

// A collection of values returned by getTemplateLimits.
type GetTemplateLimitsResultOutput struct{ *pulumi.OutputState }

func (GetTemplateLimitsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateLimitsResult)(nil)).Elem()
}

func (o GetTemplateLimitsResultOutput) ToGetTemplateLimitsResultOutput() GetTemplateLimitsResultOutput {
	return o
}

func (o GetTemplateLimitsResultOutput) ToGetTemplateLimitsResultOutputWithContext(ctx context.Context) GetTemplateLimitsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTemplateLimitsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplateLimitsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTemplateLimitsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplateLimitsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetTemplateLimitsResultOutput) TemplateLimits() GetTemplateLimitsTemplateLimitArrayOutput {
	return o.ApplyT(func(v GetTemplateLimitsResult) []GetTemplateLimitsTemplateLimit { return v.TemplateLimits }).(GetTemplateLimitsTemplateLimitArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTemplateLimitsResultOutput{})
}