// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogsets(ctx *pulumi.Context, args *GetLogsetsArgs, opts ...pulumi.InvokeOption) (*GetLogsetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogsetsResult
	err := ctx.Invoke("tencentcloud:Cls/getLogsets:getLogsets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogsets.
type GetLogsetsArgs struct {
	Filters          []GetLogsetsFilter `pulumi:"filters"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
}

// A collection of values returned by getLogsets.
type GetLogsetsResult struct {
	Filters []GetLogsetsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	Logsets          []GetLogsetsLogset `pulumi:"logsets"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
}

func GetLogsetsOutput(ctx *pulumi.Context, args GetLogsetsOutputArgs, opts ...pulumi.InvokeOption) GetLogsetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogsetsResult, error) {
			args := v.(GetLogsetsArgs)
			r, err := GetLogsets(ctx, &args, opts...)
			var s GetLogsetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogsetsResultOutput)
}

// A collection of arguments for invoking getLogsets.
type GetLogsetsOutputArgs struct {
	Filters          GetLogsetsFilterArrayInput `pulumi:"filters"`
	ResultOutputFile pulumi.StringPtrInput      `pulumi:"resultOutputFile"`
}

func (GetLogsetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsetsArgs)(nil)).Elem()
}

// A collection of values returned by getLogsets.
type GetLogsetsResultOutput struct{ *pulumi.OutputState }

func (GetLogsetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsetsResult)(nil)).Elem()
}

func (o GetLogsetsResultOutput) ToGetLogsetsResultOutput() GetLogsetsResultOutput {
	return o
}

func (o GetLogsetsResultOutput) ToGetLogsetsResultOutputWithContext(ctx context.Context) GetLogsetsResultOutput {
	return o
}

func (o GetLogsetsResultOutput) Filters() GetLogsetsFilterArrayOutput {
	return o.ApplyT(func(v GetLogsetsResult) []GetLogsetsFilter { return v.Filters }).(GetLogsetsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogsetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogsetsResultOutput) Logsets() GetLogsetsLogsetArrayOutput {
	return o.ApplyT(func(v GetLogsetsResult) []GetLogsetsLogset { return v.Logsets }).(GetLogsetsLogsetArrayOutput)
}

func (o GetLogsetsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogsetsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogsetsResultOutput{})
}