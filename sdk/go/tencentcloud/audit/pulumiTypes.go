// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package audit

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type TrackStorage struct {
	StorageName   string `pulumi:"storageName"`
	StoragePrefix string `pulumi:"storagePrefix"`
	StorageRegion string `pulumi:"storageRegion"`
	StorageType   string `pulumi:"storageType"`
}

// TrackStorageInput is an input type that accepts TrackStorageArgs and TrackStorageOutput values.
// You can construct a concrete instance of `TrackStorageInput` via:
//
//	TrackStorageArgs{...}
type TrackStorageInput interface {
	pulumi.Input

	ToTrackStorageOutput() TrackStorageOutput
	ToTrackStorageOutputWithContext(context.Context) TrackStorageOutput
}

type TrackStorageArgs struct {
	StorageName   pulumi.StringInput `pulumi:"storageName"`
	StoragePrefix pulumi.StringInput `pulumi:"storagePrefix"`
	StorageRegion pulumi.StringInput `pulumi:"storageRegion"`
	StorageType   pulumi.StringInput `pulumi:"storageType"`
}

func (TrackStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackStorage)(nil)).Elem()
}

func (i TrackStorageArgs) ToTrackStorageOutput() TrackStorageOutput {
	return i.ToTrackStorageOutputWithContext(context.Background())
}

func (i TrackStorageArgs) ToTrackStorageOutputWithContext(ctx context.Context) TrackStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackStorageOutput)
}

func (i TrackStorageArgs) ToTrackStoragePtrOutput() TrackStoragePtrOutput {
	return i.ToTrackStoragePtrOutputWithContext(context.Background())
}

func (i TrackStorageArgs) ToTrackStoragePtrOutputWithContext(ctx context.Context) TrackStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackStorageOutput).ToTrackStoragePtrOutputWithContext(ctx)
}

// TrackStoragePtrInput is an input type that accepts TrackStorageArgs, TrackStoragePtr and TrackStoragePtrOutput values.
// You can construct a concrete instance of `TrackStoragePtrInput` via:
//
//	        TrackStorageArgs{...}
//
//	or:
//
//	        nil
type TrackStoragePtrInput interface {
	pulumi.Input

	ToTrackStoragePtrOutput() TrackStoragePtrOutput
	ToTrackStoragePtrOutputWithContext(context.Context) TrackStoragePtrOutput
}

type trackStoragePtrType TrackStorageArgs

func TrackStoragePtr(v *TrackStorageArgs) TrackStoragePtrInput {
	return (*trackStoragePtrType)(v)
}

func (*trackStoragePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackStorage)(nil)).Elem()
}

func (i *trackStoragePtrType) ToTrackStoragePtrOutput() TrackStoragePtrOutput {
	return i.ToTrackStoragePtrOutputWithContext(context.Background())
}

func (i *trackStoragePtrType) ToTrackStoragePtrOutputWithContext(ctx context.Context) TrackStoragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackStoragePtrOutput)
}

type TrackStorageOutput struct{ *pulumi.OutputState }

func (TrackStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackStorage)(nil)).Elem()
}

func (o TrackStorageOutput) ToTrackStorageOutput() TrackStorageOutput {
	return o
}

func (o TrackStorageOutput) ToTrackStorageOutputWithContext(ctx context.Context) TrackStorageOutput {
	return o
}

func (o TrackStorageOutput) ToTrackStoragePtrOutput() TrackStoragePtrOutput {
	return o.ToTrackStoragePtrOutputWithContext(context.Background())
}

func (o TrackStorageOutput) ToTrackStoragePtrOutputWithContext(ctx context.Context) TrackStoragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrackStorage) *TrackStorage {
		return &v
	}).(TrackStoragePtrOutput)
}

func (o TrackStorageOutput) StorageName() pulumi.StringOutput {
	return o.ApplyT(func(v TrackStorage) string { return v.StorageName }).(pulumi.StringOutput)
}

func (o TrackStorageOutput) StoragePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v TrackStorage) string { return v.StoragePrefix }).(pulumi.StringOutput)
}

func (o TrackStorageOutput) StorageRegion() pulumi.StringOutput {
	return o.ApplyT(func(v TrackStorage) string { return v.StorageRegion }).(pulumi.StringOutput)
}

func (o TrackStorageOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v TrackStorage) string { return v.StorageType }).(pulumi.StringOutput)
}

type TrackStoragePtrOutput struct{ *pulumi.OutputState }

func (TrackStoragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrackStorage)(nil)).Elem()
}

func (o TrackStoragePtrOutput) ToTrackStoragePtrOutput() TrackStoragePtrOutput {
	return o
}

func (o TrackStoragePtrOutput) ToTrackStoragePtrOutputWithContext(ctx context.Context) TrackStoragePtrOutput {
	return o
}

func (o TrackStoragePtrOutput) Elem() TrackStorageOutput {
	return o.ApplyT(func(v *TrackStorage) TrackStorage {
		if v != nil {
			return *v
		}
		var ret TrackStorage
		return ret
	}).(TrackStorageOutput)
}

func (o TrackStoragePtrOutput) StorageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackStorage) *string {
		if v == nil {
			return nil
		}
		return &v.StorageName
	}).(pulumi.StringPtrOutput)
}

func (o TrackStoragePtrOutput) StoragePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackStorage) *string {
		if v == nil {
			return nil
		}
		return &v.StoragePrefix
	}).(pulumi.StringPtrOutput)
}

func (o TrackStoragePtrOutput) StorageRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackStorage) *string {
		if v == nil {
			return nil
		}
		return &v.StorageRegion
	}).(pulumi.StringPtrOutput)
}

func (o TrackStoragePtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrackStorage) *string {
		if v == nil {
			return nil
		}
		return &v.StorageType
	}).(pulumi.StringPtrOutput)
}

type GetCosRegionsAuditCosRegionList struct {
	CosRegion     string `pulumi:"cosRegion"`
	CosRegionName string `pulumi:"cosRegionName"`
}

// GetCosRegionsAuditCosRegionListInput is an input type that accepts GetCosRegionsAuditCosRegionListArgs and GetCosRegionsAuditCosRegionListOutput values.
// You can construct a concrete instance of `GetCosRegionsAuditCosRegionListInput` via:
//
//	GetCosRegionsAuditCosRegionListArgs{...}
type GetCosRegionsAuditCosRegionListInput interface {
	pulumi.Input

	ToGetCosRegionsAuditCosRegionListOutput() GetCosRegionsAuditCosRegionListOutput
	ToGetCosRegionsAuditCosRegionListOutputWithContext(context.Context) GetCosRegionsAuditCosRegionListOutput
}

type GetCosRegionsAuditCosRegionListArgs struct {
	CosRegion     pulumi.StringInput `pulumi:"cosRegion"`
	CosRegionName pulumi.StringInput `pulumi:"cosRegionName"`
}

func (GetCosRegionsAuditCosRegionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (i GetCosRegionsAuditCosRegionListArgs) ToGetCosRegionsAuditCosRegionListOutput() GetCosRegionsAuditCosRegionListOutput {
	return i.ToGetCosRegionsAuditCosRegionListOutputWithContext(context.Background())
}

func (i GetCosRegionsAuditCosRegionListArgs) ToGetCosRegionsAuditCosRegionListOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCosRegionsAuditCosRegionListOutput)
}

// GetCosRegionsAuditCosRegionListArrayInput is an input type that accepts GetCosRegionsAuditCosRegionListArray and GetCosRegionsAuditCosRegionListArrayOutput values.
// You can construct a concrete instance of `GetCosRegionsAuditCosRegionListArrayInput` via:
//
//	GetCosRegionsAuditCosRegionListArray{ GetCosRegionsAuditCosRegionListArgs{...} }
type GetCosRegionsAuditCosRegionListArrayInput interface {
	pulumi.Input

	ToGetCosRegionsAuditCosRegionListArrayOutput() GetCosRegionsAuditCosRegionListArrayOutput
	ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(context.Context) GetCosRegionsAuditCosRegionListArrayOutput
}

type GetCosRegionsAuditCosRegionListArray []GetCosRegionsAuditCosRegionListInput

func (GetCosRegionsAuditCosRegionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (i GetCosRegionsAuditCosRegionListArray) ToGetCosRegionsAuditCosRegionListArrayOutput() GetCosRegionsAuditCosRegionListArrayOutput {
	return i.ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(context.Background())
}

func (i GetCosRegionsAuditCosRegionListArray) ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCosRegionsAuditCosRegionListArrayOutput)
}

type GetCosRegionsAuditCosRegionListOutput struct{ *pulumi.OutputState }

func (GetCosRegionsAuditCosRegionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (o GetCosRegionsAuditCosRegionListOutput) ToGetCosRegionsAuditCosRegionListOutput() GetCosRegionsAuditCosRegionListOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListOutput) ToGetCosRegionsAuditCosRegionListOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListOutput) CosRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetCosRegionsAuditCosRegionList) string { return v.CosRegion }).(pulumi.StringOutput)
}

func (o GetCosRegionsAuditCosRegionListOutput) CosRegionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCosRegionsAuditCosRegionList) string { return v.CosRegionName }).(pulumi.StringOutput)
}

type GetCosRegionsAuditCosRegionListArrayOutput struct{ *pulumi.OutputState }

func (GetCosRegionsAuditCosRegionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCosRegionsAuditCosRegionList)(nil)).Elem()
}

func (o GetCosRegionsAuditCosRegionListArrayOutput) ToGetCosRegionsAuditCosRegionListArrayOutput() GetCosRegionsAuditCosRegionListArrayOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListArrayOutput) ToGetCosRegionsAuditCosRegionListArrayOutputWithContext(ctx context.Context) GetCosRegionsAuditCosRegionListArrayOutput {
	return o
}

func (o GetCosRegionsAuditCosRegionListArrayOutput) Index(i pulumi.IntInput) GetCosRegionsAuditCosRegionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCosRegionsAuditCosRegionList {
		return vs[0].([]GetCosRegionsAuditCosRegionList)[vs[1].(int)]
	}).(GetCosRegionsAuditCosRegionListOutput)
}

type GetKeyAliasAuditKeyAliasList struct {
	KeyAlias string `pulumi:"keyAlias"`
	KeyId    string `pulumi:"keyId"`
}

// GetKeyAliasAuditKeyAliasListInput is an input type that accepts GetKeyAliasAuditKeyAliasListArgs and GetKeyAliasAuditKeyAliasListOutput values.
// You can construct a concrete instance of `GetKeyAliasAuditKeyAliasListInput` via:
//
//	GetKeyAliasAuditKeyAliasListArgs{...}
type GetKeyAliasAuditKeyAliasListInput interface {
	pulumi.Input

	ToGetKeyAliasAuditKeyAliasListOutput() GetKeyAliasAuditKeyAliasListOutput
	ToGetKeyAliasAuditKeyAliasListOutputWithContext(context.Context) GetKeyAliasAuditKeyAliasListOutput
}

type GetKeyAliasAuditKeyAliasListArgs struct {
	KeyAlias pulumi.StringInput `pulumi:"keyAlias"`
	KeyId    pulumi.StringInput `pulumi:"keyId"`
}

func (GetKeyAliasAuditKeyAliasListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (i GetKeyAliasAuditKeyAliasListArgs) ToGetKeyAliasAuditKeyAliasListOutput() GetKeyAliasAuditKeyAliasListOutput {
	return i.ToGetKeyAliasAuditKeyAliasListOutputWithContext(context.Background())
}

func (i GetKeyAliasAuditKeyAliasListArgs) ToGetKeyAliasAuditKeyAliasListOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeyAliasAuditKeyAliasListOutput)
}

// GetKeyAliasAuditKeyAliasListArrayInput is an input type that accepts GetKeyAliasAuditKeyAliasListArray and GetKeyAliasAuditKeyAliasListArrayOutput values.
// You can construct a concrete instance of `GetKeyAliasAuditKeyAliasListArrayInput` via:
//
//	GetKeyAliasAuditKeyAliasListArray{ GetKeyAliasAuditKeyAliasListArgs{...} }
type GetKeyAliasAuditKeyAliasListArrayInput interface {
	pulumi.Input

	ToGetKeyAliasAuditKeyAliasListArrayOutput() GetKeyAliasAuditKeyAliasListArrayOutput
	ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(context.Context) GetKeyAliasAuditKeyAliasListArrayOutput
}

type GetKeyAliasAuditKeyAliasListArray []GetKeyAliasAuditKeyAliasListInput

func (GetKeyAliasAuditKeyAliasListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (i GetKeyAliasAuditKeyAliasListArray) ToGetKeyAliasAuditKeyAliasListArrayOutput() GetKeyAliasAuditKeyAliasListArrayOutput {
	return i.ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(context.Background())
}

func (i GetKeyAliasAuditKeyAliasListArray) ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetKeyAliasAuditKeyAliasListArrayOutput)
}

type GetKeyAliasAuditKeyAliasListOutput struct{ *pulumi.OutputState }

func (GetKeyAliasAuditKeyAliasListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (o GetKeyAliasAuditKeyAliasListOutput) ToGetKeyAliasAuditKeyAliasListOutput() GetKeyAliasAuditKeyAliasListOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListOutput) ToGetKeyAliasAuditKeyAliasListOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListOutput) KeyAlias() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyAliasAuditKeyAliasList) string { return v.KeyAlias }).(pulumi.StringOutput)
}

func (o GetKeyAliasAuditKeyAliasListOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKeyAliasAuditKeyAliasList) string { return v.KeyId }).(pulumi.StringOutput)
}

type GetKeyAliasAuditKeyAliasListArrayOutput struct{ *pulumi.OutputState }

func (GetKeyAliasAuditKeyAliasListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetKeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (o GetKeyAliasAuditKeyAliasListArrayOutput) ToGetKeyAliasAuditKeyAliasListArrayOutput() GetKeyAliasAuditKeyAliasListArrayOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListArrayOutput) ToGetKeyAliasAuditKeyAliasListArrayOutputWithContext(ctx context.Context) GetKeyAliasAuditKeyAliasListArrayOutput {
	return o
}

func (o GetKeyAliasAuditKeyAliasListArrayOutput) Index(i pulumi.IntInput) GetKeyAliasAuditKeyAliasListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetKeyAliasAuditKeyAliasList {
		return vs[0].([]GetKeyAliasAuditKeyAliasList)[vs[1].(int)]
	}).(GetKeyAliasAuditKeyAliasListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrackStorageInput)(nil)).Elem(), TrackStorageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrackStoragePtrInput)(nil)).Elem(), TrackStorageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCosRegionsAuditCosRegionListInput)(nil)).Elem(), GetCosRegionsAuditCosRegionListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCosRegionsAuditCosRegionListArrayInput)(nil)).Elem(), GetCosRegionsAuditCosRegionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetKeyAliasAuditKeyAliasListInput)(nil)).Elem(), GetKeyAliasAuditKeyAliasListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetKeyAliasAuditKeyAliasListArrayInput)(nil)).Elem(), GetKeyAliasAuditKeyAliasListArray{})
	pulumi.RegisterOutputType(TrackStorageOutput{})
	pulumi.RegisterOutputType(TrackStoragePtrOutput{})
	pulumi.RegisterOutputType(GetCosRegionsAuditCosRegionListOutput{})
	pulumi.RegisterOutputType(GetCosRegionsAuditCosRegionListArrayOutput{})
	pulumi.RegisterOutputType(GetKeyAliasAuditKeyAliasListOutput{})
	pulumi.RegisterOutputType(GetKeyAliasAuditKeyAliasListArrayOutput{})
}