// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupsResult
	err := ctx.Invoke("tencentcloud:Cam/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	GroupId          *string `pulumi:"groupId"`
	Name             *string `pulumi:"name"`
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	GroupId    *string              `pulumi:"groupId"`
	GroupLists []GetGroupsGroupList `pulumi:"groupLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Name             *string `pulumi:"name"`
	Remark           *string `pulumi:"remark"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			var s GetGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	GroupId          pulumi.StringPtrInput `pulumi:"groupId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	Remark           pulumi.StringPtrInput `pulumi:"remark"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) GroupLists() GetGroupsGroupListArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroupList { return v.GroupLists }).(GetGroupsGroupListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.Remark }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}