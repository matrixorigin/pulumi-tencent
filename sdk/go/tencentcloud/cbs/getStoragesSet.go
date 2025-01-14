// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStoragesSet(ctx *pulumi.Context, args *GetStoragesSetArgs, opts ...pulumi.InvokeOption) (*GetStoragesSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStoragesSetResult
	err := ctx.Invoke("tencentcloud:Cbs/getStoragesSet:getStoragesSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStoragesSet.
type GetStoragesSetArgs struct {
	AvailabilityZone   *string  `pulumi:"availabilityZone"`
	ChargeTypes        []string `pulumi:"chargeTypes"`
	DedicatedClusterId *string  `pulumi:"dedicatedClusterId"`
	InstanceIps        []string `pulumi:"instanceIps"`
	InstanceNames      []string `pulumi:"instanceNames"`
	Portable           *bool    `pulumi:"portable"`
	ProjectId          *int     `pulumi:"projectId"`
	ResultOutputFile   *string  `pulumi:"resultOutputFile"`
	StorageId          *string  `pulumi:"storageId"`
	StorageName        *string  `pulumi:"storageName"`
	StorageStates      []string `pulumi:"storageStates"`
	StorageType        *string  `pulumi:"storageType"`
	StorageUsage       *string  `pulumi:"storageUsage"`
	TagKeys            []string `pulumi:"tagKeys"`
	TagValues          []string `pulumi:"tagValues"`
}

// A collection of values returned by getStoragesSet.
type GetStoragesSetResult struct {
	AvailabilityZone   *string  `pulumi:"availabilityZone"`
	ChargeTypes        []string `pulumi:"chargeTypes"`
	DedicatedClusterId *string  `pulumi:"dedicatedClusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                      `pulumi:"id"`
	InstanceIps      []string                    `pulumi:"instanceIps"`
	InstanceNames    []string                    `pulumi:"instanceNames"`
	Portable         *bool                       `pulumi:"portable"`
	ProjectId        *int                        `pulumi:"projectId"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
	StorageId        *string                     `pulumi:"storageId"`
	StorageLists     []GetStoragesSetStorageList `pulumi:"storageLists"`
	StorageName      *string                     `pulumi:"storageName"`
	StorageStates    []string                    `pulumi:"storageStates"`
	StorageType      *string                     `pulumi:"storageType"`
	StorageUsage     *string                     `pulumi:"storageUsage"`
	TagKeys          []string                    `pulumi:"tagKeys"`
	TagValues        []string                    `pulumi:"tagValues"`
}

func GetStoragesSetOutput(ctx *pulumi.Context, args GetStoragesSetOutputArgs, opts ...pulumi.InvokeOption) GetStoragesSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStoragesSetResult, error) {
			args := v.(GetStoragesSetArgs)
			r, err := GetStoragesSet(ctx, &args, opts...)
			var s GetStoragesSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStoragesSetResultOutput)
}

// A collection of arguments for invoking getStoragesSet.
type GetStoragesSetOutputArgs struct {
	AvailabilityZone   pulumi.StringPtrInput   `pulumi:"availabilityZone"`
	ChargeTypes        pulumi.StringArrayInput `pulumi:"chargeTypes"`
	DedicatedClusterId pulumi.StringPtrInput   `pulumi:"dedicatedClusterId"`
	InstanceIps        pulumi.StringArrayInput `pulumi:"instanceIps"`
	InstanceNames      pulumi.StringArrayInput `pulumi:"instanceNames"`
	Portable           pulumi.BoolPtrInput     `pulumi:"portable"`
	ProjectId          pulumi.IntPtrInput      `pulumi:"projectId"`
	ResultOutputFile   pulumi.StringPtrInput   `pulumi:"resultOutputFile"`
	StorageId          pulumi.StringPtrInput   `pulumi:"storageId"`
	StorageName        pulumi.StringPtrInput   `pulumi:"storageName"`
	StorageStates      pulumi.StringArrayInput `pulumi:"storageStates"`
	StorageType        pulumi.StringPtrInput   `pulumi:"storageType"`
	StorageUsage       pulumi.StringPtrInput   `pulumi:"storageUsage"`
	TagKeys            pulumi.StringArrayInput `pulumi:"tagKeys"`
	TagValues          pulumi.StringArrayInput `pulumi:"tagValues"`
}

func (GetStoragesSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesSetArgs)(nil)).Elem()
}

// A collection of values returned by getStoragesSet.
type GetStoragesSetResultOutput struct{ *pulumi.OutputState }

func (GetStoragesSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStoragesSetResult)(nil)).Elem()
}

func (o GetStoragesSetResultOutput) ToGetStoragesSetResultOutput() GetStoragesSetResultOutput {
	return o
}

func (o GetStoragesSetResultOutput) ToGetStoragesSetResultOutputWithContext(ctx context.Context) GetStoragesSetResultOutput {
	return o
}

func (o GetStoragesSetResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) ChargeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.ChargeTypes }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) DedicatedClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.DedicatedClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStoragesSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStoragesSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStoragesSetResultOutput) InstanceIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.InstanceIps }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) InstanceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.InstanceNames }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) Portable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *bool { return v.Portable }).(pulumi.BoolPtrOutput)
}

func (o GetStoragesSetResultOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

func (o GetStoragesSetResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) StorageLists() GetStoragesSetStorageListArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []GetStoragesSetStorageList { return v.StorageLists }).(GetStoragesSetStorageListArrayOutput)
}

func (o GetStoragesSetResultOutput) StorageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageName }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) StorageStates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.StorageStates }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) StorageUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStoragesSetResult) *string { return v.StorageUsage }).(pulumi.StringPtrOutput)
}

func (o GetStoragesSetResultOutput) TagKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.TagKeys }).(pulumi.StringArrayOutput)
}

func (o GetStoragesSetResultOutput) TagValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetStoragesSetResult) []string { return v.TagValues }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStoragesSetResultOutput{})
}
