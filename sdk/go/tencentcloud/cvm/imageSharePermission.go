// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImageSharePermission struct {
	pulumi.CustomResourceState

	// List of account IDs with which an image is shared.
	AccountIds pulumi.StringArrayOutput `pulumi:"accountIds"`
	// Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
}

// NewImageSharePermission registers a new resource with the given unique name, arguments, and options.
func NewImageSharePermission(ctx *pulumi.Context,
	name string, args *ImageSharePermissionArgs, opts ...pulumi.ResourceOption) (*ImageSharePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountIds == nil {
		return nil, errors.New("invalid value for required argument 'AccountIds'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImageSharePermission
	err := ctx.RegisterResource("tencentcloud:Cvm/imageSharePermission:ImageSharePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageSharePermission gets an existing ImageSharePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageSharePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageSharePermissionState, opts ...pulumi.ResourceOption) (*ImageSharePermission, error) {
	var resource ImageSharePermission
	err := ctx.ReadResource("tencentcloud:Cvm/imageSharePermission:ImageSharePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageSharePermission resources.
type imageSharePermissionState struct {
	// List of account IDs with which an image is shared.
	AccountIds []string `pulumi:"accountIds"`
	// Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
	ImageId *string `pulumi:"imageId"`
}

type ImageSharePermissionState struct {
	// List of account IDs with which an image is shared.
	AccountIds pulumi.StringArrayInput
	// Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
	ImageId pulumi.StringPtrInput
}

func (ImageSharePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageSharePermissionState)(nil)).Elem()
}

type imageSharePermissionArgs struct {
	// List of account IDs with which an image is shared.
	AccountIds []string `pulumi:"accountIds"`
	// Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
	ImageId string `pulumi:"imageId"`
}

// The set of arguments for constructing a ImageSharePermission resource.
type ImageSharePermissionArgs struct {
	// List of account IDs with which an image is shared.
	AccountIds pulumi.StringArrayInput
	// Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
	ImageId pulumi.StringInput
}

func (ImageSharePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageSharePermissionArgs)(nil)).Elem()
}

type ImageSharePermissionInput interface {
	pulumi.Input

	ToImageSharePermissionOutput() ImageSharePermissionOutput
	ToImageSharePermissionOutputWithContext(ctx context.Context) ImageSharePermissionOutput
}

func (*ImageSharePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageSharePermission)(nil)).Elem()
}

func (i *ImageSharePermission) ToImageSharePermissionOutput() ImageSharePermissionOutput {
	return i.ToImageSharePermissionOutputWithContext(context.Background())
}

func (i *ImageSharePermission) ToImageSharePermissionOutputWithContext(ctx context.Context) ImageSharePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionOutput)
}

// ImageSharePermissionArrayInput is an input type that accepts ImageSharePermissionArray and ImageSharePermissionArrayOutput values.
// You can construct a concrete instance of `ImageSharePermissionArrayInput` via:
//
//	ImageSharePermissionArray{ ImageSharePermissionArgs{...} }
type ImageSharePermissionArrayInput interface {
	pulumi.Input

	ToImageSharePermissionArrayOutput() ImageSharePermissionArrayOutput
	ToImageSharePermissionArrayOutputWithContext(context.Context) ImageSharePermissionArrayOutput
}

type ImageSharePermissionArray []ImageSharePermissionInput

func (ImageSharePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageSharePermission)(nil)).Elem()
}

func (i ImageSharePermissionArray) ToImageSharePermissionArrayOutput() ImageSharePermissionArrayOutput {
	return i.ToImageSharePermissionArrayOutputWithContext(context.Background())
}

func (i ImageSharePermissionArray) ToImageSharePermissionArrayOutputWithContext(ctx context.Context) ImageSharePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionArrayOutput)
}

// ImageSharePermissionMapInput is an input type that accepts ImageSharePermissionMap and ImageSharePermissionMapOutput values.
// You can construct a concrete instance of `ImageSharePermissionMapInput` via:
//
//	ImageSharePermissionMap{ "key": ImageSharePermissionArgs{...} }
type ImageSharePermissionMapInput interface {
	pulumi.Input

	ToImageSharePermissionMapOutput() ImageSharePermissionMapOutput
	ToImageSharePermissionMapOutputWithContext(context.Context) ImageSharePermissionMapOutput
}

type ImageSharePermissionMap map[string]ImageSharePermissionInput

func (ImageSharePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageSharePermission)(nil)).Elem()
}

func (i ImageSharePermissionMap) ToImageSharePermissionMapOutput() ImageSharePermissionMapOutput {
	return i.ToImageSharePermissionMapOutputWithContext(context.Background())
}

func (i ImageSharePermissionMap) ToImageSharePermissionMapOutputWithContext(ctx context.Context) ImageSharePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageSharePermissionMapOutput)
}

type ImageSharePermissionOutput struct{ *pulumi.OutputState }

func (ImageSharePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageSharePermission)(nil)).Elem()
}

func (o ImageSharePermissionOutput) ToImageSharePermissionOutput() ImageSharePermissionOutput {
	return o
}

func (o ImageSharePermissionOutput) ToImageSharePermissionOutputWithContext(ctx context.Context) ImageSharePermissionOutput {
	return o
}

// List of account IDs with which an image is shared.
func (o ImageSharePermissionOutput) AccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImageSharePermission) pulumi.StringArrayOutput { return v.AccountIds }).(pulumi.StringArrayOutput)
}

// Image ID such as `img-gvbnzy6f`. You can only specify an image in the NORMAL state.
func (o ImageSharePermissionOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageSharePermission) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

type ImageSharePermissionArrayOutput struct{ *pulumi.OutputState }

func (ImageSharePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageSharePermission)(nil)).Elem()
}

func (o ImageSharePermissionArrayOutput) ToImageSharePermissionArrayOutput() ImageSharePermissionArrayOutput {
	return o
}

func (o ImageSharePermissionArrayOutput) ToImageSharePermissionArrayOutputWithContext(ctx context.Context) ImageSharePermissionArrayOutput {
	return o
}

func (o ImageSharePermissionArrayOutput) Index(i pulumi.IntInput) ImageSharePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageSharePermission {
		return vs[0].([]*ImageSharePermission)[vs[1].(int)]
	}).(ImageSharePermissionOutput)
}

type ImageSharePermissionMapOutput struct{ *pulumi.OutputState }

func (ImageSharePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageSharePermission)(nil)).Elem()
}

func (o ImageSharePermissionMapOutput) ToImageSharePermissionMapOutput() ImageSharePermissionMapOutput {
	return o
}

func (o ImageSharePermissionMapOutput) ToImageSharePermissionMapOutputWithContext(ctx context.Context) ImageSharePermissionMapOutput {
	return o
}

func (o ImageSharePermissionMapOutput) MapIndex(k pulumi.StringInput) ImageSharePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageSharePermission {
		return vs[0].(map[string]*ImageSharePermission)[vs[1].(string)]
	}).(ImageSharePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageSharePermissionInput)(nil)).Elem(), &ImageSharePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageSharePermissionArrayInput)(nil)).Elem(), ImageSharePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageSharePermissionMapInput)(nil)).Elem(), ImageSharePermissionMap{})
	pulumi.RegisterOutputType(ImageSharePermissionOutput{})
	pulumi.RegisterOutputType(ImageSharePermissionArrayOutput{})
	pulumi.RegisterOutputType(ImageSharePermissionMapOutput{})
}