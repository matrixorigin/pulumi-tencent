// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deprecatedvpc

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type Ipv6EniAddressIpv6Address struct {
	Address        string  `pulumi:"address"`
	AddressId      *string `pulumi:"addressId"`
	Description    *string `pulumi:"description"`
	IsWanIpBlocked *bool   `pulumi:"isWanIpBlocked"`
	Primary        *bool   `pulumi:"primary"`
	State          *string `pulumi:"state"`
}

// Ipv6EniAddressIpv6AddressInput is an input type that accepts Ipv6EniAddressIpv6AddressArgs and Ipv6EniAddressIpv6AddressOutput values.
// You can construct a concrete instance of `Ipv6EniAddressIpv6AddressInput` via:
//
//	Ipv6EniAddressIpv6AddressArgs{...}
type Ipv6EniAddressIpv6AddressInput interface {
	pulumi.Input

	ToIpv6EniAddressIpv6AddressOutput() Ipv6EniAddressIpv6AddressOutput
	ToIpv6EniAddressIpv6AddressOutputWithContext(context.Context) Ipv6EniAddressIpv6AddressOutput
}

type Ipv6EniAddressIpv6AddressArgs struct {
	Address        pulumi.StringInput    `pulumi:"address"`
	AddressId      pulumi.StringPtrInput `pulumi:"addressId"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	IsWanIpBlocked pulumi.BoolPtrInput   `pulumi:"isWanIpBlocked"`
	Primary        pulumi.BoolPtrInput   `pulumi:"primary"`
	State          pulumi.StringPtrInput `pulumi:"state"`
}

func (Ipv6EniAddressIpv6AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6EniAddressIpv6Address)(nil)).Elem()
}

func (i Ipv6EniAddressIpv6AddressArgs) ToIpv6EniAddressIpv6AddressOutput() Ipv6EniAddressIpv6AddressOutput {
	return i.ToIpv6EniAddressIpv6AddressOutputWithContext(context.Background())
}

func (i Ipv6EniAddressIpv6AddressArgs) ToIpv6EniAddressIpv6AddressOutputWithContext(ctx context.Context) Ipv6EniAddressIpv6AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6EniAddressIpv6AddressOutput)
}

// Ipv6EniAddressIpv6AddressArrayInput is an input type that accepts Ipv6EniAddressIpv6AddressArray and Ipv6EniAddressIpv6AddressArrayOutput values.
// You can construct a concrete instance of `Ipv6EniAddressIpv6AddressArrayInput` via:
//
//	Ipv6EniAddressIpv6AddressArray{ Ipv6EniAddressIpv6AddressArgs{...} }
type Ipv6EniAddressIpv6AddressArrayInput interface {
	pulumi.Input

	ToIpv6EniAddressIpv6AddressArrayOutput() Ipv6EniAddressIpv6AddressArrayOutput
	ToIpv6EniAddressIpv6AddressArrayOutputWithContext(context.Context) Ipv6EniAddressIpv6AddressArrayOutput
}

type Ipv6EniAddressIpv6AddressArray []Ipv6EniAddressIpv6AddressInput

func (Ipv6EniAddressIpv6AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ipv6EniAddressIpv6Address)(nil)).Elem()
}

func (i Ipv6EniAddressIpv6AddressArray) ToIpv6EniAddressIpv6AddressArrayOutput() Ipv6EniAddressIpv6AddressArrayOutput {
	return i.ToIpv6EniAddressIpv6AddressArrayOutputWithContext(context.Background())
}

func (i Ipv6EniAddressIpv6AddressArray) ToIpv6EniAddressIpv6AddressArrayOutputWithContext(ctx context.Context) Ipv6EniAddressIpv6AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6EniAddressIpv6AddressArrayOutput)
}

type Ipv6EniAddressIpv6AddressOutput struct{ *pulumi.OutputState }

func (Ipv6EniAddressIpv6AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6EniAddressIpv6Address)(nil)).Elem()
}

func (o Ipv6EniAddressIpv6AddressOutput) ToIpv6EniAddressIpv6AddressOutput() Ipv6EniAddressIpv6AddressOutput {
	return o
}

func (o Ipv6EniAddressIpv6AddressOutput) ToIpv6EniAddressIpv6AddressOutputWithContext(ctx context.Context) Ipv6EniAddressIpv6AddressOutput {
	return o
}

func (o Ipv6EniAddressIpv6AddressOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v Ipv6EniAddressIpv6Address) string { return v.Address }).(pulumi.StringOutput)
}

func (o Ipv6EniAddressIpv6AddressOutput) AddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6EniAddressIpv6Address) *string { return v.AddressId }).(pulumi.StringPtrOutput)
}

func (o Ipv6EniAddressIpv6AddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6EniAddressIpv6Address) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o Ipv6EniAddressIpv6AddressOutput) IsWanIpBlocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ipv6EniAddressIpv6Address) *bool { return v.IsWanIpBlocked }).(pulumi.BoolPtrOutput)
}

func (o Ipv6EniAddressIpv6AddressOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ipv6EniAddressIpv6Address) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o Ipv6EniAddressIpv6AddressOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6EniAddressIpv6Address) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv6EniAddressIpv6AddressArrayOutput struct{ *pulumi.OutputState }

func (Ipv6EniAddressIpv6AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ipv6EniAddressIpv6Address)(nil)).Elem()
}

func (o Ipv6EniAddressIpv6AddressArrayOutput) ToIpv6EniAddressIpv6AddressArrayOutput() Ipv6EniAddressIpv6AddressArrayOutput {
	return o
}

func (o Ipv6EniAddressIpv6AddressArrayOutput) ToIpv6EniAddressIpv6AddressArrayOutputWithContext(ctx context.Context) Ipv6EniAddressIpv6AddressArrayOutput {
	return o
}

func (o Ipv6EniAddressIpv6AddressArrayOutput) Index(i pulumi.IntInput) Ipv6EniAddressIpv6AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ipv6EniAddressIpv6Address {
		return vs[0].([]Ipv6EniAddressIpv6Address)[vs[1].(int)]
	}).(Ipv6EniAddressIpv6AddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6EniAddressIpv6AddressInput)(nil)).Elem(), Ipv6EniAddressIpv6AddressArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv6EniAddressIpv6AddressArrayInput)(nil)).Elem(), Ipv6EniAddressIpv6AddressArray{})
	pulumi.RegisterOutputType(Ipv6EniAddressIpv6AddressOutput{})
	pulumi.RegisterOutputType(Ipv6EniAddressIpv6AddressArrayOutput{})
}