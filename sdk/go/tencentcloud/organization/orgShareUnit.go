// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrgShareUnit struct {
	pulumi.CustomResourceState

	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	Area pulumi.StringOutput `pulumi:"area"`
	// Shared unit description. Up to 128 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
	// 3-128 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	UnitId pulumi.StringOutput `pulumi:"unitId"`
}

// NewOrgShareUnit registers a new resource with the given unique name, arguments, and options.
func NewOrgShareUnit(ctx *pulumi.Context,
	name string, args *OrgShareUnitArgs, opts ...pulumi.ResourceOption) (*OrgShareUnit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Area == nil {
		return nil, errors.New("invalid value for required argument 'Area'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgShareUnit
	err := ctx.RegisterResource("tencentcloud:Organization/orgShareUnit:OrgShareUnit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgShareUnit gets an existing OrgShareUnit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgShareUnit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgShareUnitState, opts ...pulumi.ResourceOption) (*OrgShareUnit, error) {
	var resource OrgShareUnit
	err := ctx.ReadResource("tencentcloud:Organization/orgShareUnit:OrgShareUnit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgShareUnit resources.
type orgShareUnitState struct {
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	Area *string `pulumi:"area"`
	// Shared unit description. Up to 128 characters.
	Description *string `pulumi:"description"`
	// Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
	// 3-128 characters.
	Name *string `pulumi:"name"`
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	UnitId *string `pulumi:"unitId"`
}

type OrgShareUnitState struct {
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	Area pulumi.StringPtrInput
	// Shared unit description. Up to 128 characters.
	Description pulumi.StringPtrInput
	// Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
	// 3-128 characters.
	Name pulumi.StringPtrInput
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	UnitId pulumi.StringPtrInput
}

func (OrgShareUnitState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgShareUnitState)(nil)).Elem()
}

type orgShareUnitArgs struct {
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	Area string `pulumi:"area"`
	// Shared unit description. Up to 128 characters.
	Description *string `pulumi:"description"`
	// Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
	// 3-128 characters.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a OrgShareUnit resource.
type OrgShareUnitArgs struct {
	// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
	Area pulumi.StringInput
	// Shared unit description. Up to 128 characters.
	Description pulumi.StringPtrInput
	// Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
	// 3-128 characters.
	Name pulumi.StringPtrInput
}

func (OrgShareUnitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgShareUnitArgs)(nil)).Elem()
}

type OrgShareUnitInput interface {
	pulumi.Input

	ToOrgShareUnitOutput() OrgShareUnitOutput
	ToOrgShareUnitOutputWithContext(ctx context.Context) OrgShareUnitOutput
}

func (*OrgShareUnit) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgShareUnit)(nil)).Elem()
}

func (i *OrgShareUnit) ToOrgShareUnitOutput() OrgShareUnitOutput {
	return i.ToOrgShareUnitOutputWithContext(context.Background())
}

func (i *OrgShareUnit) ToOrgShareUnitOutputWithContext(ctx context.Context) OrgShareUnitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgShareUnitOutput)
}

// OrgShareUnitArrayInput is an input type that accepts OrgShareUnitArray and OrgShareUnitArrayOutput values.
// You can construct a concrete instance of `OrgShareUnitArrayInput` via:
//
//	OrgShareUnitArray{ OrgShareUnitArgs{...} }
type OrgShareUnitArrayInput interface {
	pulumi.Input

	ToOrgShareUnitArrayOutput() OrgShareUnitArrayOutput
	ToOrgShareUnitArrayOutputWithContext(context.Context) OrgShareUnitArrayOutput
}

type OrgShareUnitArray []OrgShareUnitInput

func (OrgShareUnitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgShareUnit)(nil)).Elem()
}

func (i OrgShareUnitArray) ToOrgShareUnitArrayOutput() OrgShareUnitArrayOutput {
	return i.ToOrgShareUnitArrayOutputWithContext(context.Background())
}

func (i OrgShareUnitArray) ToOrgShareUnitArrayOutputWithContext(ctx context.Context) OrgShareUnitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgShareUnitArrayOutput)
}

// OrgShareUnitMapInput is an input type that accepts OrgShareUnitMap and OrgShareUnitMapOutput values.
// You can construct a concrete instance of `OrgShareUnitMapInput` via:
//
//	OrgShareUnitMap{ "key": OrgShareUnitArgs{...} }
type OrgShareUnitMapInput interface {
	pulumi.Input

	ToOrgShareUnitMapOutput() OrgShareUnitMapOutput
	ToOrgShareUnitMapOutputWithContext(context.Context) OrgShareUnitMapOutput
}

type OrgShareUnitMap map[string]OrgShareUnitInput

func (OrgShareUnitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgShareUnit)(nil)).Elem()
}

func (i OrgShareUnitMap) ToOrgShareUnitMapOutput() OrgShareUnitMapOutput {
	return i.ToOrgShareUnitMapOutputWithContext(context.Background())
}

func (i OrgShareUnitMap) ToOrgShareUnitMapOutputWithContext(ctx context.Context) OrgShareUnitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgShareUnitMapOutput)
}

type OrgShareUnitOutput struct{ *pulumi.OutputState }

func (OrgShareUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgShareUnit)(nil)).Elem()
}

func (o OrgShareUnitOutput) ToOrgShareUnitOutput() OrgShareUnitOutput {
	return o
}

func (o OrgShareUnitOutput) ToOrgShareUnitOutputWithContext(ctx context.Context) OrgShareUnitOutput {
	return o
}

// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
func (o OrgShareUnitOutput) Area() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgShareUnit) pulumi.StringOutput { return v.Area }).(pulumi.StringOutput)
}

// Shared unit description. Up to 128 characters.
func (o OrgShareUnitOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgShareUnit) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
// 3-128 characters.
func (o OrgShareUnitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgShareUnit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
func (o OrgShareUnitOutput) UnitId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgShareUnit) pulumi.StringOutput { return v.UnitId }).(pulumi.StringOutput)
}

type OrgShareUnitArrayOutput struct{ *pulumi.OutputState }

func (OrgShareUnitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgShareUnit)(nil)).Elem()
}

func (o OrgShareUnitArrayOutput) ToOrgShareUnitArrayOutput() OrgShareUnitArrayOutput {
	return o
}

func (o OrgShareUnitArrayOutput) ToOrgShareUnitArrayOutputWithContext(ctx context.Context) OrgShareUnitArrayOutput {
	return o
}

func (o OrgShareUnitArrayOutput) Index(i pulumi.IntInput) OrgShareUnitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgShareUnit {
		return vs[0].([]*OrgShareUnit)[vs[1].(int)]
	}).(OrgShareUnitOutput)
}

type OrgShareUnitMapOutput struct{ *pulumi.OutputState }

func (OrgShareUnitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgShareUnit)(nil)).Elem()
}

func (o OrgShareUnitMapOutput) ToOrgShareUnitMapOutput() OrgShareUnitMapOutput {
	return o
}

func (o OrgShareUnitMapOutput) ToOrgShareUnitMapOutputWithContext(ctx context.Context) OrgShareUnitMapOutput {
	return o
}

func (o OrgShareUnitMapOutput) MapIndex(k pulumi.StringInput) OrgShareUnitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgShareUnit {
		return vs[0].(map[string]*OrgShareUnit)[vs[1].(string)]
	}).(OrgShareUnitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgShareUnitInput)(nil)).Elem(), &OrgShareUnit{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgShareUnitArrayInput)(nil)).Elem(), OrgShareUnitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgShareUnitMapInput)(nil)).Elem(), OrgShareUnitMap{})
	pulumi.RegisterOutputType(OrgShareUnitOutput{})
	pulumi.RegisterOutputType(OrgShareUnitArrayOutput{})
	pulumi.RegisterOutputType(OrgShareUnitMapOutput{})
}