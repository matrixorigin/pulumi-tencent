// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMachines(ctx *pulumi.Context, args *GetMachinesArgs, opts ...pulumi.InvokeOption) (*GetMachinesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMachinesResult
	err := ctx.Invoke("tencentcloud:Cls/getMachines:getMachines", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMachines.
type GetMachinesArgs struct {
	GroupId          string  `pulumi:"groupId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getMachines.
type GetMachinesResult struct {
	GroupId string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string               `pulumi:"id"`
	Machines         []GetMachinesMachine `pulumi:"machines"`
	ResultOutputFile *string              `pulumi:"resultOutputFile"`
}

func GetMachinesOutput(ctx *pulumi.Context, args GetMachinesOutputArgs, opts ...pulumi.InvokeOption) GetMachinesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMachinesResult, error) {
			args := v.(GetMachinesArgs)
			r, err := GetMachines(ctx, &args, opts...)
			var s GetMachinesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMachinesResultOutput)
}

// A collection of arguments for invoking getMachines.
type GetMachinesOutputArgs struct {
	GroupId          pulumi.StringInput    `pulumi:"groupId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetMachinesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMachinesArgs)(nil)).Elem()
}

// A collection of values returned by getMachines.
type GetMachinesResultOutput struct{ *pulumi.OutputState }

func (GetMachinesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMachinesResult)(nil)).Elem()
}

func (o GetMachinesResultOutput) ToGetMachinesResultOutput() GetMachinesResultOutput {
	return o
}

func (o GetMachinesResultOutput) ToGetMachinesResultOutputWithContext(ctx context.Context) GetMachinesResultOutput {
	return o
}

func (o GetMachinesResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinesResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMachinesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMachinesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMachinesResultOutput) Machines() GetMachinesMachineArrayOutput {
	return o.ApplyT(func(v GetMachinesResult) []GetMachinesMachine { return v.Machines }).(GetMachinesMachineArrayOutput)
}

func (o GetMachinesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMachinesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMachinesResultOutput{})
}