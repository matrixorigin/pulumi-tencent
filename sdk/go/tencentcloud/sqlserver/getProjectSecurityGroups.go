// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProjectSecurityGroups(ctx *pulumi.Context, args *GetProjectSecurityGroupsArgs, opts ...pulumi.InvokeOption) (*GetProjectSecurityGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectSecurityGroupsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getProjectSecurityGroups:getProjectSecurityGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectSecurityGroups.
type GetProjectSecurityGroupsArgs struct {
	ProjectId        int     `pulumi:"projectId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getProjectSecurityGroups.
type GetProjectSecurityGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string                                     `pulumi:"id"`
	ProjectId         int                                        `pulumi:"projectId"`
	ResultOutputFile  *string                                    `pulumi:"resultOutputFile"`
	SecurityGroupSets []GetProjectSecurityGroupsSecurityGroupSet `pulumi:"securityGroupSets"`
}

func GetProjectSecurityGroupsOutput(ctx *pulumi.Context, args GetProjectSecurityGroupsOutputArgs, opts ...pulumi.InvokeOption) GetProjectSecurityGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectSecurityGroupsResult, error) {
			args := v.(GetProjectSecurityGroupsArgs)
			r, err := GetProjectSecurityGroups(ctx, &args, opts...)
			var s GetProjectSecurityGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectSecurityGroupsResultOutput)
}

// A collection of arguments for invoking getProjectSecurityGroups.
type GetProjectSecurityGroupsOutputArgs struct {
	ProjectId        pulumi.IntInput       `pulumi:"projectId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetProjectSecurityGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectSecurityGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getProjectSecurityGroups.
type GetProjectSecurityGroupsResultOutput struct{ *pulumi.OutputState }

func (GetProjectSecurityGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectSecurityGroupsResult)(nil)).Elem()
}

func (o GetProjectSecurityGroupsResultOutput) ToGetProjectSecurityGroupsResultOutput() GetProjectSecurityGroupsResultOutput {
	return o
}

func (o GetProjectSecurityGroupsResultOutput) ToGetProjectSecurityGroupsResultOutputWithContext(ctx context.Context) GetProjectSecurityGroupsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectSecurityGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProjectSecurityGroupsResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o GetProjectSecurityGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetProjectSecurityGroupsResultOutput) SecurityGroupSets() GetProjectSecurityGroupsSecurityGroupSetArrayOutput {
	return o.ApplyT(func(v GetProjectSecurityGroupsResult) []GetProjectSecurityGroupsSecurityGroupSet {
		return v.SecurityGroupSets
	}).(GetProjectSecurityGroupsSecurityGroupSetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectSecurityGroupsResultOutput{})
}