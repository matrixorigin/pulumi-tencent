// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUserDomains(ctx *pulumi.Context, args *GetUserDomainsArgs, opts ...pulumi.InvokeOption) (*GetUserDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserDomainsResult
	err := ctx.Invoke("tencentcloud:Waf/getUserDomains:getUserDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserDomains.
type GetUserDomainsArgs struct {
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getUserDomains.
type GetUserDomainsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                    `pulumi:"id"`
	ResultOutputFile *string                   `pulumi:"resultOutputFile"`
	UsersInfos       []GetUserDomainsUsersInfo `pulumi:"usersInfos"`
}

func GetUserDomainsOutput(ctx *pulumi.Context, args GetUserDomainsOutputArgs, opts ...pulumi.InvokeOption) GetUserDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserDomainsResult, error) {
			args := v.(GetUserDomainsArgs)
			r, err := GetUserDomains(ctx, &args, opts...)
			var s GetUserDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserDomainsResultOutput)
}

// A collection of arguments for invoking getUserDomains.
type GetUserDomainsOutputArgs struct {
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetUserDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getUserDomains.
type GetUserDomainsResultOutput struct{ *pulumi.OutputState }

func (GetUserDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserDomainsResult)(nil)).Elem()
}

func (o GetUserDomainsResultOutput) ToGetUserDomainsResultOutput() GetUserDomainsResultOutput {
	return o
}

func (o GetUserDomainsResultOutput) ToGetUserDomainsResultOutputWithContext(ctx context.Context) GetUserDomainsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserDomainsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserDomainsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetUserDomainsResultOutput) UsersInfos() GetUserDomainsUsersInfoArrayOutput {
	return o.ApplyT(func(v GetUserDomainsResult) []GetUserDomainsUsersInfo { return v.UsersInfos }).(GetUserDomainsUsersInfoArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserDomainsResultOutput{})
}