// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRoles(ctx *pulumi.Context, args *GetRolesArgs, opts ...pulumi.InvokeOption) (*GetRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRolesResult
	err := ctx.Invoke("tencentcloud:Cam/getRoles:getRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoles.
type GetRolesArgs struct {
	Description      *string `pulumi:"description"`
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RoleId           *string `pulumi:"roleId"`
}

// A collection of values returned by getRoles.
type GetRolesResult struct {
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id               string             `pulumi:"id"`
	Name             *string            `pulumi:"name"`
	ResultOutputFile *string            `pulumi:"resultOutputFile"`
	RoleId           *string            `pulumi:"roleId"`
	RoleLists        []GetRolesRoleList `pulumi:"roleLists"`
}

func GetRolesOutput(ctx *pulumi.Context, args GetRolesOutputArgs, opts ...pulumi.InvokeOption) GetRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRolesResult, error) {
			args := v.(GetRolesArgs)
			r, err := GetRoles(ctx, &args, opts...)
			var s GetRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRolesResultOutput)
}

// A collection of arguments for invoking getRoles.
type GetRolesOutputArgs struct {
	Description      pulumi.StringPtrInput `pulumi:"description"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	RoleId           pulumi.StringPtrInput `pulumi:"roleId"`
}

func (GetRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesArgs)(nil)).Elem()
}

// A collection of values returned by getRoles.
type GetRolesResultOutput struct{ *pulumi.OutputState }

func (GetRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesResult)(nil)).Elem()
}

func (o GetRolesResultOutput) ToGetRolesResultOutput() GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) ToGetRolesResultOutputWithContext(ctx context.Context) GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRolesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetRolesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRolesResultOutput) RoleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.RoleId }).(pulumi.StringPtrOutput)
}

func (o GetRolesResultOutput) RoleLists() GetRolesRoleListArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []GetRolesRoleList { return v.RoleLists }).(GetRolesRoleListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolesResultOutput{})
}