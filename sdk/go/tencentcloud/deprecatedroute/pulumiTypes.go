// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deprecatedroute

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetTableRoute struct {
	CidrBlock   string `pulumi:"cidrBlock"`
	Description string `pulumi:"description"`
	NextHub     string `pulumi:"nextHub"`
	NextType    string `pulumi:"nextType"`
}

// GetTableRouteInput is an input type that accepts GetTableRouteArgs and GetTableRouteOutput values.
// You can construct a concrete instance of `GetTableRouteInput` via:
//
//	GetTableRouteArgs{...}
type GetTableRouteInput interface {
	pulumi.Input

	ToGetTableRouteOutput() GetTableRouteOutput
	ToGetTableRouteOutputWithContext(context.Context) GetTableRouteOutput
}

type GetTableRouteArgs struct {
	CidrBlock   pulumi.StringInput `pulumi:"cidrBlock"`
	Description pulumi.StringInput `pulumi:"description"`
	NextHub     pulumi.StringInput `pulumi:"nextHub"`
	NextType    pulumi.StringInput `pulumi:"nextType"`
}

func (GetTableRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTableRoute)(nil)).Elem()
}

func (i GetTableRouteArgs) ToGetTableRouteOutput() GetTableRouteOutput {
	return i.ToGetTableRouteOutputWithContext(context.Background())
}

func (i GetTableRouteArgs) ToGetTableRouteOutputWithContext(ctx context.Context) GetTableRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTableRouteOutput)
}

// GetTableRouteArrayInput is an input type that accepts GetTableRouteArray and GetTableRouteArrayOutput values.
// You can construct a concrete instance of `GetTableRouteArrayInput` via:
//
//	GetTableRouteArray{ GetTableRouteArgs{...} }
type GetTableRouteArrayInput interface {
	pulumi.Input

	ToGetTableRouteArrayOutput() GetTableRouteArrayOutput
	ToGetTableRouteArrayOutputWithContext(context.Context) GetTableRouteArrayOutput
}

type GetTableRouteArray []GetTableRouteInput

func (GetTableRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTableRoute)(nil)).Elem()
}

func (i GetTableRouteArray) ToGetTableRouteArrayOutput() GetTableRouteArrayOutput {
	return i.ToGetTableRouteArrayOutputWithContext(context.Background())
}

func (i GetTableRouteArray) ToGetTableRouteArrayOutputWithContext(ctx context.Context) GetTableRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTableRouteArrayOutput)
}

type GetTableRouteOutput struct{ *pulumi.OutputState }

func (GetTableRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTableRoute)(nil)).Elem()
}

func (o GetTableRouteOutput) ToGetTableRouteOutput() GetTableRouteOutput {
	return o
}

func (o GetTableRouteOutput) ToGetTableRouteOutputWithContext(ctx context.Context) GetTableRouteOutput {
	return o
}

func (o GetTableRouteOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableRoute) string { return v.CidrBlock }).(pulumi.StringOutput)
}

func (o GetTableRouteOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableRoute) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetTableRouteOutput) NextHub() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableRoute) string { return v.NextHub }).(pulumi.StringOutput)
}

func (o GetTableRouteOutput) NextType() pulumi.StringOutput {
	return o.ApplyT(func(v GetTableRoute) string { return v.NextType }).(pulumi.StringOutput)
}

type GetTableRouteArrayOutput struct{ *pulumi.OutputState }

func (GetTableRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetTableRoute)(nil)).Elem()
}

func (o GetTableRouteArrayOutput) ToGetTableRouteArrayOutput() GetTableRouteArrayOutput {
	return o
}

func (o GetTableRouteArrayOutput) ToGetTableRouteArrayOutputWithContext(ctx context.Context) GetTableRouteArrayOutput {
	return o
}

func (o GetTableRouteArrayOutput) Index(i pulumi.IntInput) GetTableRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetTableRoute {
		return vs[0].([]GetTableRoute)[vs[1].(int)]
	}).(GetTableRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetTableRouteInput)(nil)).Elem(), GetTableRouteArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetTableRouteArrayInput)(nil)).Elem(), GetTableRouteArray{})
	pulumi.RegisterOutputType(GetTableRouteOutput{})
	pulumi.RegisterOutputType(GetTableRouteArrayOutput{})
}