// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfw

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type NatInstanceNewModeItems struct {
	Eips     []string `pulumi:"eips"`
	VpcLists []string `pulumi:"vpcLists"`
}

// NatInstanceNewModeItemsInput is an input type that accepts NatInstanceNewModeItemsArgs and NatInstanceNewModeItemsOutput values.
// You can construct a concrete instance of `NatInstanceNewModeItemsInput` via:
//
//	NatInstanceNewModeItemsArgs{...}
type NatInstanceNewModeItemsInput interface {
	pulumi.Input

	ToNatInstanceNewModeItemsOutput() NatInstanceNewModeItemsOutput
	ToNatInstanceNewModeItemsOutputWithContext(context.Context) NatInstanceNewModeItemsOutput
}

type NatInstanceNewModeItemsArgs struct {
	Eips     pulumi.StringArrayInput `pulumi:"eips"`
	VpcLists pulumi.StringArrayInput `pulumi:"vpcLists"`
}

func (NatInstanceNewModeItemsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NatInstanceNewModeItems)(nil)).Elem()
}

func (i NatInstanceNewModeItemsArgs) ToNatInstanceNewModeItemsOutput() NatInstanceNewModeItemsOutput {
	return i.ToNatInstanceNewModeItemsOutputWithContext(context.Background())
}

func (i NatInstanceNewModeItemsArgs) ToNatInstanceNewModeItemsOutputWithContext(ctx context.Context) NatInstanceNewModeItemsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatInstanceNewModeItemsOutput)
}

func (i NatInstanceNewModeItemsArgs) ToNatInstanceNewModeItemsPtrOutput() NatInstanceNewModeItemsPtrOutput {
	return i.ToNatInstanceNewModeItemsPtrOutputWithContext(context.Background())
}

func (i NatInstanceNewModeItemsArgs) ToNatInstanceNewModeItemsPtrOutputWithContext(ctx context.Context) NatInstanceNewModeItemsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatInstanceNewModeItemsOutput).ToNatInstanceNewModeItemsPtrOutputWithContext(ctx)
}

// NatInstanceNewModeItemsPtrInput is an input type that accepts NatInstanceNewModeItemsArgs, NatInstanceNewModeItemsPtr and NatInstanceNewModeItemsPtrOutput values.
// You can construct a concrete instance of `NatInstanceNewModeItemsPtrInput` via:
//
//	        NatInstanceNewModeItemsArgs{...}
//
//	or:
//
//	        nil
type NatInstanceNewModeItemsPtrInput interface {
	pulumi.Input

	ToNatInstanceNewModeItemsPtrOutput() NatInstanceNewModeItemsPtrOutput
	ToNatInstanceNewModeItemsPtrOutputWithContext(context.Context) NatInstanceNewModeItemsPtrOutput
}

type natInstanceNewModeItemsPtrType NatInstanceNewModeItemsArgs

func NatInstanceNewModeItemsPtr(v *NatInstanceNewModeItemsArgs) NatInstanceNewModeItemsPtrInput {
	return (*natInstanceNewModeItemsPtrType)(v)
}

func (*natInstanceNewModeItemsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NatInstanceNewModeItems)(nil)).Elem()
}

func (i *natInstanceNewModeItemsPtrType) ToNatInstanceNewModeItemsPtrOutput() NatInstanceNewModeItemsPtrOutput {
	return i.ToNatInstanceNewModeItemsPtrOutputWithContext(context.Background())
}

func (i *natInstanceNewModeItemsPtrType) ToNatInstanceNewModeItemsPtrOutputWithContext(ctx context.Context) NatInstanceNewModeItemsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatInstanceNewModeItemsPtrOutput)
}

type NatInstanceNewModeItemsOutput struct{ *pulumi.OutputState }

func (NatInstanceNewModeItemsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatInstanceNewModeItems)(nil)).Elem()
}

func (o NatInstanceNewModeItemsOutput) ToNatInstanceNewModeItemsOutput() NatInstanceNewModeItemsOutput {
	return o
}

func (o NatInstanceNewModeItemsOutput) ToNatInstanceNewModeItemsOutputWithContext(ctx context.Context) NatInstanceNewModeItemsOutput {
	return o
}

func (o NatInstanceNewModeItemsOutput) ToNatInstanceNewModeItemsPtrOutput() NatInstanceNewModeItemsPtrOutput {
	return o.ToNatInstanceNewModeItemsPtrOutputWithContext(context.Background())
}

func (o NatInstanceNewModeItemsOutput) ToNatInstanceNewModeItemsPtrOutputWithContext(ctx context.Context) NatInstanceNewModeItemsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NatInstanceNewModeItems) *NatInstanceNewModeItems {
		return &v
	}).(NatInstanceNewModeItemsPtrOutput)
}

func (o NatInstanceNewModeItemsOutput) Eips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NatInstanceNewModeItems) []string { return v.Eips }).(pulumi.StringArrayOutput)
}

func (o NatInstanceNewModeItemsOutput) VpcLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NatInstanceNewModeItems) []string { return v.VpcLists }).(pulumi.StringArrayOutput)
}

type NatInstanceNewModeItemsPtrOutput struct{ *pulumi.OutputState }

func (NatInstanceNewModeItemsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatInstanceNewModeItems)(nil)).Elem()
}

func (o NatInstanceNewModeItemsPtrOutput) ToNatInstanceNewModeItemsPtrOutput() NatInstanceNewModeItemsPtrOutput {
	return o
}

func (o NatInstanceNewModeItemsPtrOutput) ToNatInstanceNewModeItemsPtrOutputWithContext(ctx context.Context) NatInstanceNewModeItemsPtrOutput {
	return o
}

func (o NatInstanceNewModeItemsPtrOutput) Elem() NatInstanceNewModeItemsOutput {
	return o.ApplyT(func(v *NatInstanceNewModeItems) NatInstanceNewModeItems {
		if v != nil {
			return *v
		}
		var ret NatInstanceNewModeItems
		return ret
	}).(NatInstanceNewModeItemsOutput)
}

func (o NatInstanceNewModeItemsPtrOutput) Eips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NatInstanceNewModeItems) []string {
		if v == nil {
			return nil
		}
		return v.Eips
	}).(pulumi.StringArrayOutput)
}

func (o NatInstanceNewModeItemsPtrOutput) VpcLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NatInstanceNewModeItems) []string {
		if v == nil {
			return nil
		}
		return v.VpcLists
	}).(pulumi.StringArrayOutput)
}

type VpcInstanceVpcFwInstance struct {
	FwDeploy VpcInstanceVpcFwInstanceFwDeploy `pulumi:"fwDeploy"`
	FwInsId  *string                          `pulumi:"fwInsId"`
	Name     string                           `pulumi:"name"`
	VpcIds   []string                         `pulumi:"vpcIds"`
}

// VpcInstanceVpcFwInstanceInput is an input type that accepts VpcInstanceVpcFwInstanceArgs and VpcInstanceVpcFwInstanceOutput values.
// You can construct a concrete instance of `VpcInstanceVpcFwInstanceInput` via:
//
//	VpcInstanceVpcFwInstanceArgs{...}
type VpcInstanceVpcFwInstanceInput interface {
	pulumi.Input

	ToVpcInstanceVpcFwInstanceOutput() VpcInstanceVpcFwInstanceOutput
	ToVpcInstanceVpcFwInstanceOutputWithContext(context.Context) VpcInstanceVpcFwInstanceOutput
}

type VpcInstanceVpcFwInstanceArgs struct {
	FwDeploy VpcInstanceVpcFwInstanceFwDeployInput `pulumi:"fwDeploy"`
	FwInsId  pulumi.StringPtrInput                 `pulumi:"fwInsId"`
	Name     pulumi.StringInput                    `pulumi:"name"`
	VpcIds   pulumi.StringArrayInput               `pulumi:"vpcIds"`
}

func (VpcInstanceVpcFwInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcInstanceVpcFwInstance)(nil)).Elem()
}

func (i VpcInstanceVpcFwInstanceArgs) ToVpcInstanceVpcFwInstanceOutput() VpcInstanceVpcFwInstanceOutput {
	return i.ToVpcInstanceVpcFwInstanceOutputWithContext(context.Background())
}

func (i VpcInstanceVpcFwInstanceArgs) ToVpcInstanceVpcFwInstanceOutputWithContext(ctx context.Context) VpcInstanceVpcFwInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcInstanceVpcFwInstanceOutput)
}

// VpcInstanceVpcFwInstanceArrayInput is an input type that accepts VpcInstanceVpcFwInstanceArray and VpcInstanceVpcFwInstanceArrayOutput values.
// You can construct a concrete instance of `VpcInstanceVpcFwInstanceArrayInput` via:
//
//	VpcInstanceVpcFwInstanceArray{ VpcInstanceVpcFwInstanceArgs{...} }
type VpcInstanceVpcFwInstanceArrayInput interface {
	pulumi.Input

	ToVpcInstanceVpcFwInstanceArrayOutput() VpcInstanceVpcFwInstanceArrayOutput
	ToVpcInstanceVpcFwInstanceArrayOutputWithContext(context.Context) VpcInstanceVpcFwInstanceArrayOutput
}

type VpcInstanceVpcFwInstanceArray []VpcInstanceVpcFwInstanceInput

func (VpcInstanceVpcFwInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcInstanceVpcFwInstance)(nil)).Elem()
}

func (i VpcInstanceVpcFwInstanceArray) ToVpcInstanceVpcFwInstanceArrayOutput() VpcInstanceVpcFwInstanceArrayOutput {
	return i.ToVpcInstanceVpcFwInstanceArrayOutputWithContext(context.Background())
}

func (i VpcInstanceVpcFwInstanceArray) ToVpcInstanceVpcFwInstanceArrayOutputWithContext(ctx context.Context) VpcInstanceVpcFwInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcInstanceVpcFwInstanceArrayOutput)
}

type VpcInstanceVpcFwInstanceOutput struct{ *pulumi.OutputState }

func (VpcInstanceVpcFwInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcInstanceVpcFwInstance)(nil)).Elem()
}

func (o VpcInstanceVpcFwInstanceOutput) ToVpcInstanceVpcFwInstanceOutput() VpcInstanceVpcFwInstanceOutput {
	return o
}

func (o VpcInstanceVpcFwInstanceOutput) ToVpcInstanceVpcFwInstanceOutputWithContext(ctx context.Context) VpcInstanceVpcFwInstanceOutput {
	return o
}

func (o VpcInstanceVpcFwInstanceOutput) FwDeploy() VpcInstanceVpcFwInstanceFwDeployOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstance) VpcInstanceVpcFwInstanceFwDeploy { return v.FwDeploy }).(VpcInstanceVpcFwInstanceFwDeployOutput)
}

func (o VpcInstanceVpcFwInstanceOutput) FwInsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstance) *string { return v.FwInsId }).(pulumi.StringPtrOutput)
}

func (o VpcInstanceVpcFwInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstance) string { return v.Name }).(pulumi.StringOutput)
}

func (o VpcInstanceVpcFwInstanceOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstance) []string { return v.VpcIds }).(pulumi.StringArrayOutput)
}

type VpcInstanceVpcFwInstanceArrayOutput struct{ *pulumi.OutputState }

func (VpcInstanceVpcFwInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcInstanceVpcFwInstance)(nil)).Elem()
}

func (o VpcInstanceVpcFwInstanceArrayOutput) ToVpcInstanceVpcFwInstanceArrayOutput() VpcInstanceVpcFwInstanceArrayOutput {
	return o
}

func (o VpcInstanceVpcFwInstanceArrayOutput) ToVpcInstanceVpcFwInstanceArrayOutputWithContext(ctx context.Context) VpcInstanceVpcFwInstanceArrayOutput {
	return o
}

func (o VpcInstanceVpcFwInstanceArrayOutput) Index(i pulumi.IntInput) VpcInstanceVpcFwInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcInstanceVpcFwInstance {
		return vs[0].([]VpcInstanceVpcFwInstance)[vs[1].(int)]
	}).(VpcInstanceVpcFwInstanceOutput)
}

type VpcInstanceVpcFwInstanceFwDeploy struct {
	CrossAZone   *int     `pulumi:"crossAZone"`
	DeployRegion string   `pulumi:"deployRegion"`
	Width        int      `pulumi:"width"`
	ZoneSets     []string `pulumi:"zoneSets"`
}

// VpcInstanceVpcFwInstanceFwDeployInput is an input type that accepts VpcInstanceVpcFwInstanceFwDeployArgs and VpcInstanceVpcFwInstanceFwDeployOutput values.
// You can construct a concrete instance of `VpcInstanceVpcFwInstanceFwDeployInput` via:
//
//	VpcInstanceVpcFwInstanceFwDeployArgs{...}
type VpcInstanceVpcFwInstanceFwDeployInput interface {
	pulumi.Input

	ToVpcInstanceVpcFwInstanceFwDeployOutput() VpcInstanceVpcFwInstanceFwDeployOutput
	ToVpcInstanceVpcFwInstanceFwDeployOutputWithContext(context.Context) VpcInstanceVpcFwInstanceFwDeployOutput
}

type VpcInstanceVpcFwInstanceFwDeployArgs struct {
	CrossAZone   pulumi.IntPtrInput      `pulumi:"crossAZone"`
	DeployRegion pulumi.StringInput      `pulumi:"deployRegion"`
	Width        pulumi.IntInput         `pulumi:"width"`
	ZoneSets     pulumi.StringArrayInput `pulumi:"zoneSets"`
}

func (VpcInstanceVpcFwInstanceFwDeployArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcInstanceVpcFwInstanceFwDeploy)(nil)).Elem()
}

func (i VpcInstanceVpcFwInstanceFwDeployArgs) ToVpcInstanceVpcFwInstanceFwDeployOutput() VpcInstanceVpcFwInstanceFwDeployOutput {
	return i.ToVpcInstanceVpcFwInstanceFwDeployOutputWithContext(context.Background())
}

func (i VpcInstanceVpcFwInstanceFwDeployArgs) ToVpcInstanceVpcFwInstanceFwDeployOutputWithContext(ctx context.Context) VpcInstanceVpcFwInstanceFwDeployOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcInstanceVpcFwInstanceFwDeployOutput)
}

type VpcInstanceVpcFwInstanceFwDeployOutput struct{ *pulumi.OutputState }

func (VpcInstanceVpcFwInstanceFwDeployOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcInstanceVpcFwInstanceFwDeploy)(nil)).Elem()
}

func (o VpcInstanceVpcFwInstanceFwDeployOutput) ToVpcInstanceVpcFwInstanceFwDeployOutput() VpcInstanceVpcFwInstanceFwDeployOutput {
	return o
}

func (o VpcInstanceVpcFwInstanceFwDeployOutput) ToVpcInstanceVpcFwInstanceFwDeployOutputWithContext(ctx context.Context) VpcInstanceVpcFwInstanceFwDeployOutput {
	return o
}

func (o VpcInstanceVpcFwInstanceFwDeployOutput) CrossAZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstanceFwDeploy) *int { return v.CrossAZone }).(pulumi.IntPtrOutput)
}

func (o VpcInstanceVpcFwInstanceFwDeployOutput) DeployRegion() pulumi.StringOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstanceFwDeploy) string { return v.DeployRegion }).(pulumi.StringOutput)
}

func (o VpcInstanceVpcFwInstanceFwDeployOutput) Width() pulumi.IntOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstanceFwDeploy) int { return v.Width }).(pulumi.IntOutput)
}

func (o VpcInstanceVpcFwInstanceFwDeployOutput) ZoneSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpcInstanceVpcFwInstanceFwDeploy) []string { return v.ZoneSets }).(pulumi.StringArrayOutput)
}

type GetEdgeFwSwitchesData struct {
	AssetType    string `pulumi:"assetType"`
	InstanceId   string `pulumi:"instanceId"`
	InstanceName string `pulumi:"instanceName"`
	IntranetIp   string `pulumi:"intranetIp"`
	PublicIp     string `pulumi:"publicIp"`
	PublicIpType int    `pulumi:"publicIpType"`
	Region       string `pulumi:"region"`
	Status       int    `pulumi:"status"`
	SwitchMode   int    `pulumi:"switchMode"`
	VpcId        string `pulumi:"vpcId"`
}

// GetEdgeFwSwitchesDataInput is an input type that accepts GetEdgeFwSwitchesDataArgs and GetEdgeFwSwitchesDataOutput values.
// You can construct a concrete instance of `GetEdgeFwSwitchesDataInput` via:
//
//	GetEdgeFwSwitchesDataArgs{...}
type GetEdgeFwSwitchesDataInput interface {
	pulumi.Input

	ToGetEdgeFwSwitchesDataOutput() GetEdgeFwSwitchesDataOutput
	ToGetEdgeFwSwitchesDataOutputWithContext(context.Context) GetEdgeFwSwitchesDataOutput
}

type GetEdgeFwSwitchesDataArgs struct {
	AssetType    pulumi.StringInput `pulumi:"assetType"`
	InstanceId   pulumi.StringInput `pulumi:"instanceId"`
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	IntranetIp   pulumi.StringInput `pulumi:"intranetIp"`
	PublicIp     pulumi.StringInput `pulumi:"publicIp"`
	PublicIpType pulumi.IntInput    `pulumi:"publicIpType"`
	Region       pulumi.StringInput `pulumi:"region"`
	Status       pulumi.IntInput    `pulumi:"status"`
	SwitchMode   pulumi.IntInput    `pulumi:"switchMode"`
	VpcId        pulumi.StringInput `pulumi:"vpcId"`
}

func (GetEdgeFwSwitchesDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeFwSwitchesData)(nil)).Elem()
}

func (i GetEdgeFwSwitchesDataArgs) ToGetEdgeFwSwitchesDataOutput() GetEdgeFwSwitchesDataOutput {
	return i.ToGetEdgeFwSwitchesDataOutputWithContext(context.Background())
}

func (i GetEdgeFwSwitchesDataArgs) ToGetEdgeFwSwitchesDataOutputWithContext(ctx context.Context) GetEdgeFwSwitchesDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEdgeFwSwitchesDataOutput)
}

// GetEdgeFwSwitchesDataArrayInput is an input type that accepts GetEdgeFwSwitchesDataArray and GetEdgeFwSwitchesDataArrayOutput values.
// You can construct a concrete instance of `GetEdgeFwSwitchesDataArrayInput` via:
//
//	GetEdgeFwSwitchesDataArray{ GetEdgeFwSwitchesDataArgs{...} }
type GetEdgeFwSwitchesDataArrayInput interface {
	pulumi.Input

	ToGetEdgeFwSwitchesDataArrayOutput() GetEdgeFwSwitchesDataArrayOutput
	ToGetEdgeFwSwitchesDataArrayOutputWithContext(context.Context) GetEdgeFwSwitchesDataArrayOutput
}

type GetEdgeFwSwitchesDataArray []GetEdgeFwSwitchesDataInput

func (GetEdgeFwSwitchesDataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEdgeFwSwitchesData)(nil)).Elem()
}

func (i GetEdgeFwSwitchesDataArray) ToGetEdgeFwSwitchesDataArrayOutput() GetEdgeFwSwitchesDataArrayOutput {
	return i.ToGetEdgeFwSwitchesDataArrayOutputWithContext(context.Background())
}

func (i GetEdgeFwSwitchesDataArray) ToGetEdgeFwSwitchesDataArrayOutputWithContext(ctx context.Context) GetEdgeFwSwitchesDataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEdgeFwSwitchesDataArrayOutput)
}

type GetEdgeFwSwitchesDataOutput struct{ *pulumi.OutputState }

func (GetEdgeFwSwitchesDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeFwSwitchesData)(nil)).Elem()
}

func (o GetEdgeFwSwitchesDataOutput) ToGetEdgeFwSwitchesDataOutput() GetEdgeFwSwitchesDataOutput {
	return o
}

func (o GetEdgeFwSwitchesDataOutput) ToGetEdgeFwSwitchesDataOutputWithContext(ctx context.Context) GetEdgeFwSwitchesDataOutput {
	return o
}

func (o GetEdgeFwSwitchesDataOutput) AssetType() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.AssetType }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesDataOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesDataOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesDataOutput) IntranetIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.IntranetIp }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesDataOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesDataOutput) PublicIpType() pulumi.IntOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) int { return v.PublicIpType }).(pulumi.IntOutput)
}

func (o GetEdgeFwSwitchesDataOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesDataOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) int { return v.Status }).(pulumi.IntOutput)
}

func (o GetEdgeFwSwitchesDataOutput) SwitchMode() pulumi.IntOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) int { return v.SwitchMode }).(pulumi.IntOutput)
}

func (o GetEdgeFwSwitchesDataOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesData) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetEdgeFwSwitchesDataArrayOutput struct{ *pulumi.OutputState }

func (GetEdgeFwSwitchesDataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetEdgeFwSwitchesData)(nil)).Elem()
}

func (o GetEdgeFwSwitchesDataArrayOutput) ToGetEdgeFwSwitchesDataArrayOutput() GetEdgeFwSwitchesDataArrayOutput {
	return o
}

func (o GetEdgeFwSwitchesDataArrayOutput) ToGetEdgeFwSwitchesDataArrayOutputWithContext(ctx context.Context) GetEdgeFwSwitchesDataArrayOutput {
	return o
}

func (o GetEdgeFwSwitchesDataArrayOutput) Index(i pulumi.IntInput) GetEdgeFwSwitchesDataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetEdgeFwSwitchesData {
		return vs[0].([]GetEdgeFwSwitchesData)[vs[1].(int)]
	}).(GetEdgeFwSwitchesDataOutput)
}

type GetNatFwSwitchesData struct {
	Abnormal   int    `pulumi:"abnormal"`
	CvmNum     int    `pulumi:"cvmNum"`
	Enable     int    `pulumi:"enable"`
	Id         int    `pulumi:"id"`
	NatId      string `pulumi:"natId"`
	NatInsId   string `pulumi:"natInsId"`
	NatInsName string `pulumi:"natInsName"`
	NatName    string `pulumi:"natName"`
	Region     string `pulumi:"region"`
	RouteId    string `pulumi:"routeId"`
	RouteName  string `pulumi:"routeName"`
	Status     int    `pulumi:"status"`
	SubnetCidr string `pulumi:"subnetCidr"`
	SubnetId   string `pulumi:"subnetId"`
	SubnetName string `pulumi:"subnetName"`
	VpcId      string `pulumi:"vpcId"`
	VpcName    string `pulumi:"vpcName"`
}

// GetNatFwSwitchesDataInput is an input type that accepts GetNatFwSwitchesDataArgs and GetNatFwSwitchesDataOutput values.
// You can construct a concrete instance of `GetNatFwSwitchesDataInput` via:
//
//	GetNatFwSwitchesDataArgs{...}
type GetNatFwSwitchesDataInput interface {
	pulumi.Input

	ToGetNatFwSwitchesDataOutput() GetNatFwSwitchesDataOutput
	ToGetNatFwSwitchesDataOutputWithContext(context.Context) GetNatFwSwitchesDataOutput
}

type GetNatFwSwitchesDataArgs struct {
	Abnormal   pulumi.IntInput    `pulumi:"abnormal"`
	CvmNum     pulumi.IntInput    `pulumi:"cvmNum"`
	Enable     pulumi.IntInput    `pulumi:"enable"`
	Id         pulumi.IntInput    `pulumi:"id"`
	NatId      pulumi.StringInput `pulumi:"natId"`
	NatInsId   pulumi.StringInput `pulumi:"natInsId"`
	NatInsName pulumi.StringInput `pulumi:"natInsName"`
	NatName    pulumi.StringInput `pulumi:"natName"`
	Region     pulumi.StringInput `pulumi:"region"`
	RouteId    pulumi.StringInput `pulumi:"routeId"`
	RouteName  pulumi.StringInput `pulumi:"routeName"`
	Status     pulumi.IntInput    `pulumi:"status"`
	SubnetCidr pulumi.StringInput `pulumi:"subnetCidr"`
	SubnetId   pulumi.StringInput `pulumi:"subnetId"`
	SubnetName pulumi.StringInput `pulumi:"subnetName"`
	VpcId      pulumi.StringInput `pulumi:"vpcId"`
	VpcName    pulumi.StringInput `pulumi:"vpcName"`
}

func (GetNatFwSwitchesDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNatFwSwitchesData)(nil)).Elem()
}

func (i GetNatFwSwitchesDataArgs) ToGetNatFwSwitchesDataOutput() GetNatFwSwitchesDataOutput {
	return i.ToGetNatFwSwitchesDataOutputWithContext(context.Background())
}

func (i GetNatFwSwitchesDataArgs) ToGetNatFwSwitchesDataOutputWithContext(ctx context.Context) GetNatFwSwitchesDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNatFwSwitchesDataOutput)
}

// GetNatFwSwitchesDataArrayInput is an input type that accepts GetNatFwSwitchesDataArray and GetNatFwSwitchesDataArrayOutput values.
// You can construct a concrete instance of `GetNatFwSwitchesDataArrayInput` via:
//
//	GetNatFwSwitchesDataArray{ GetNatFwSwitchesDataArgs{...} }
type GetNatFwSwitchesDataArrayInput interface {
	pulumi.Input

	ToGetNatFwSwitchesDataArrayOutput() GetNatFwSwitchesDataArrayOutput
	ToGetNatFwSwitchesDataArrayOutputWithContext(context.Context) GetNatFwSwitchesDataArrayOutput
}

type GetNatFwSwitchesDataArray []GetNatFwSwitchesDataInput

func (GetNatFwSwitchesDataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNatFwSwitchesData)(nil)).Elem()
}

func (i GetNatFwSwitchesDataArray) ToGetNatFwSwitchesDataArrayOutput() GetNatFwSwitchesDataArrayOutput {
	return i.ToGetNatFwSwitchesDataArrayOutputWithContext(context.Background())
}

func (i GetNatFwSwitchesDataArray) ToGetNatFwSwitchesDataArrayOutputWithContext(ctx context.Context) GetNatFwSwitchesDataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNatFwSwitchesDataArrayOutput)
}

type GetNatFwSwitchesDataOutput struct{ *pulumi.OutputState }

func (GetNatFwSwitchesDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNatFwSwitchesData)(nil)).Elem()
}

func (o GetNatFwSwitchesDataOutput) ToGetNatFwSwitchesDataOutput() GetNatFwSwitchesDataOutput {
	return o
}

func (o GetNatFwSwitchesDataOutput) ToGetNatFwSwitchesDataOutputWithContext(ctx context.Context) GetNatFwSwitchesDataOutput {
	return o
}

func (o GetNatFwSwitchesDataOutput) Abnormal() pulumi.IntOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) int { return v.Abnormal }).(pulumi.IntOutput)
}

func (o GetNatFwSwitchesDataOutput) CvmNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) int { return v.CvmNum }).(pulumi.IntOutput)
}

func (o GetNatFwSwitchesDataOutput) Enable() pulumi.IntOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) int { return v.Enable }).(pulumi.IntOutput)
}

func (o GetNatFwSwitchesDataOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) int { return v.Id }).(pulumi.IntOutput)
}

func (o GetNatFwSwitchesDataOutput) NatId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.NatId }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) NatInsId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.NatInsId }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) NatInsName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.NatInsName }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) NatName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.NatName }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.RouteId }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) RouteName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.RouteName }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) int { return v.Status }).(pulumi.IntOutput)
}

func (o GetNatFwSwitchesDataOutput) SubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.SubnetCidr }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) SubnetName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.SubnetName }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.VpcId }).(pulumi.StringOutput)
}

func (o GetNatFwSwitchesDataOutput) VpcName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNatFwSwitchesData) string { return v.VpcName }).(pulumi.StringOutput)
}

type GetNatFwSwitchesDataArrayOutput struct{ *pulumi.OutputState }

func (GetNatFwSwitchesDataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNatFwSwitchesData)(nil)).Elem()
}

func (o GetNatFwSwitchesDataArrayOutput) ToGetNatFwSwitchesDataArrayOutput() GetNatFwSwitchesDataArrayOutput {
	return o
}

func (o GetNatFwSwitchesDataArrayOutput) ToGetNatFwSwitchesDataArrayOutputWithContext(ctx context.Context) GetNatFwSwitchesDataArrayOutput {
	return o
}

func (o GetNatFwSwitchesDataArrayOutput) Index(i pulumi.IntInput) GetNatFwSwitchesDataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNatFwSwitchesData {
		return vs[0].([]GetNatFwSwitchesData)[vs[1].(int)]
	}).(GetNatFwSwitchesDataOutput)
}

type GetVpcFwSwitchesSwitchList struct {
	Enable     int    `pulumi:"enable"`
	Status     int    `pulumi:"status"`
	SwitchId   string `pulumi:"switchId"`
	SwitchMode int    `pulumi:"switchMode"`
	SwitchName string `pulumi:"switchName"`
}

// GetVpcFwSwitchesSwitchListInput is an input type that accepts GetVpcFwSwitchesSwitchListArgs and GetVpcFwSwitchesSwitchListOutput values.
// You can construct a concrete instance of `GetVpcFwSwitchesSwitchListInput` via:
//
//	GetVpcFwSwitchesSwitchListArgs{...}
type GetVpcFwSwitchesSwitchListInput interface {
	pulumi.Input

	ToGetVpcFwSwitchesSwitchListOutput() GetVpcFwSwitchesSwitchListOutput
	ToGetVpcFwSwitchesSwitchListOutputWithContext(context.Context) GetVpcFwSwitchesSwitchListOutput
}

type GetVpcFwSwitchesSwitchListArgs struct {
	Enable     pulumi.IntInput    `pulumi:"enable"`
	Status     pulumi.IntInput    `pulumi:"status"`
	SwitchId   pulumi.StringInput `pulumi:"switchId"`
	SwitchMode pulumi.IntInput    `pulumi:"switchMode"`
	SwitchName pulumi.StringInput `pulumi:"switchName"`
}

func (GetVpcFwSwitchesSwitchListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcFwSwitchesSwitchList)(nil)).Elem()
}

func (i GetVpcFwSwitchesSwitchListArgs) ToGetVpcFwSwitchesSwitchListOutput() GetVpcFwSwitchesSwitchListOutput {
	return i.ToGetVpcFwSwitchesSwitchListOutputWithContext(context.Background())
}

func (i GetVpcFwSwitchesSwitchListArgs) ToGetVpcFwSwitchesSwitchListOutputWithContext(ctx context.Context) GetVpcFwSwitchesSwitchListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcFwSwitchesSwitchListOutput)
}

// GetVpcFwSwitchesSwitchListArrayInput is an input type that accepts GetVpcFwSwitchesSwitchListArray and GetVpcFwSwitchesSwitchListArrayOutput values.
// You can construct a concrete instance of `GetVpcFwSwitchesSwitchListArrayInput` via:
//
//	GetVpcFwSwitchesSwitchListArray{ GetVpcFwSwitchesSwitchListArgs{...} }
type GetVpcFwSwitchesSwitchListArrayInput interface {
	pulumi.Input

	ToGetVpcFwSwitchesSwitchListArrayOutput() GetVpcFwSwitchesSwitchListArrayOutput
	ToGetVpcFwSwitchesSwitchListArrayOutputWithContext(context.Context) GetVpcFwSwitchesSwitchListArrayOutput
}

type GetVpcFwSwitchesSwitchListArray []GetVpcFwSwitchesSwitchListInput

func (GetVpcFwSwitchesSwitchListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcFwSwitchesSwitchList)(nil)).Elem()
}

func (i GetVpcFwSwitchesSwitchListArray) ToGetVpcFwSwitchesSwitchListArrayOutput() GetVpcFwSwitchesSwitchListArrayOutput {
	return i.ToGetVpcFwSwitchesSwitchListArrayOutputWithContext(context.Background())
}

func (i GetVpcFwSwitchesSwitchListArray) ToGetVpcFwSwitchesSwitchListArrayOutputWithContext(ctx context.Context) GetVpcFwSwitchesSwitchListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetVpcFwSwitchesSwitchListArrayOutput)
}

type GetVpcFwSwitchesSwitchListOutput struct{ *pulumi.OutputState }

func (GetVpcFwSwitchesSwitchListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcFwSwitchesSwitchList)(nil)).Elem()
}

func (o GetVpcFwSwitchesSwitchListOutput) ToGetVpcFwSwitchesSwitchListOutput() GetVpcFwSwitchesSwitchListOutput {
	return o
}

func (o GetVpcFwSwitchesSwitchListOutput) ToGetVpcFwSwitchesSwitchListOutputWithContext(ctx context.Context) GetVpcFwSwitchesSwitchListOutput {
	return o
}

func (o GetVpcFwSwitchesSwitchListOutput) Enable() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpcFwSwitchesSwitchList) int { return v.Enable }).(pulumi.IntOutput)
}

func (o GetVpcFwSwitchesSwitchListOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpcFwSwitchesSwitchList) int { return v.Status }).(pulumi.IntOutput)
}

func (o GetVpcFwSwitchesSwitchListOutput) SwitchId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcFwSwitchesSwitchList) string { return v.SwitchId }).(pulumi.StringOutput)
}

func (o GetVpcFwSwitchesSwitchListOutput) SwitchMode() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpcFwSwitchesSwitchList) int { return v.SwitchMode }).(pulumi.IntOutput)
}

func (o GetVpcFwSwitchesSwitchListOutput) SwitchName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcFwSwitchesSwitchList) string { return v.SwitchName }).(pulumi.StringOutput)
}

type GetVpcFwSwitchesSwitchListArrayOutput struct{ *pulumi.OutputState }

func (GetVpcFwSwitchesSwitchListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetVpcFwSwitchesSwitchList)(nil)).Elem()
}

func (o GetVpcFwSwitchesSwitchListArrayOutput) ToGetVpcFwSwitchesSwitchListArrayOutput() GetVpcFwSwitchesSwitchListArrayOutput {
	return o
}

func (o GetVpcFwSwitchesSwitchListArrayOutput) ToGetVpcFwSwitchesSwitchListArrayOutputWithContext(ctx context.Context) GetVpcFwSwitchesSwitchListArrayOutput {
	return o
}

func (o GetVpcFwSwitchesSwitchListArrayOutput) Index(i pulumi.IntInput) GetVpcFwSwitchesSwitchListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetVpcFwSwitchesSwitchList {
		return vs[0].([]GetVpcFwSwitchesSwitchList)[vs[1].(int)]
	}).(GetVpcFwSwitchesSwitchListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatInstanceNewModeItemsInput)(nil)).Elem(), NatInstanceNewModeItemsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatInstanceNewModeItemsPtrInput)(nil)).Elem(), NatInstanceNewModeItemsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInstanceVpcFwInstanceInput)(nil)).Elem(), VpcInstanceVpcFwInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInstanceVpcFwInstanceArrayInput)(nil)).Elem(), VpcInstanceVpcFwInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInstanceVpcFwInstanceFwDeployInput)(nil)).Elem(), VpcInstanceVpcFwInstanceFwDeployArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetEdgeFwSwitchesDataInput)(nil)).Elem(), GetEdgeFwSwitchesDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetEdgeFwSwitchesDataArrayInput)(nil)).Elem(), GetEdgeFwSwitchesDataArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNatFwSwitchesDataInput)(nil)).Elem(), GetNatFwSwitchesDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNatFwSwitchesDataArrayInput)(nil)).Elem(), GetNatFwSwitchesDataArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetVpcFwSwitchesSwitchListInput)(nil)).Elem(), GetVpcFwSwitchesSwitchListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetVpcFwSwitchesSwitchListArrayInput)(nil)).Elem(), GetVpcFwSwitchesSwitchListArray{})
	pulumi.RegisterOutputType(NatInstanceNewModeItemsOutput{})
	pulumi.RegisterOutputType(NatInstanceNewModeItemsPtrOutput{})
	pulumi.RegisterOutputType(VpcInstanceVpcFwInstanceOutput{})
	pulumi.RegisterOutputType(VpcInstanceVpcFwInstanceArrayOutput{})
	pulumi.RegisterOutputType(VpcInstanceVpcFwInstanceFwDeployOutput{})
	pulumi.RegisterOutputType(GetEdgeFwSwitchesDataOutput{})
	pulumi.RegisterOutputType(GetEdgeFwSwitchesDataArrayOutput{})
	pulumi.RegisterOutputType(GetNatFwSwitchesDataOutput{})
	pulumi.RegisterOutputType(GetNatFwSwitchesDataArrayOutput{})
	pulumi.RegisterOutputType(GetVpcFwSwitchesSwitchListOutput{})
	pulumi.RegisterOutputType(GetVpcFwSwitchesSwitchListArrayOutput{})
}