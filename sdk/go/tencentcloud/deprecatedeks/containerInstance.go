// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deprecatedeks

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerInstance struct {
	pulumi.CustomResourceState

	// Indicates whether to create EIP instead of specify existing EIPs. Conflict with `existed_eip_ids`.
	AutoCreateEip pulumi.BoolPtrOutput `pulumi:"autoCreateEip"`
	// ID of EIP which create automatically.
	AutoCreateEipId pulumi.StringOutput `pulumi:"autoCreateEipId"`
	// CAM role name authorized to access.
	CamRoleName pulumi.StringPtrOutput `pulumi:"camRoleName"`
	// List of CBS volume.
	CbsVolumes ContainerInstanceCbsVolumeArrayOutput `pulumi:"cbsVolumes"`
	// List of container.
	Containers ContainerInstanceContainerArrayOutput `pulumi:"containers"`
	// The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Cpu pulumi.Float64Output `pulumi:"cpu"`
	// Type of cpu, which can set to `intel` or `amd`. It also support backup list like `amd,intel` which indicates using
	// `intel` when `amd` sold out.
	CpuType pulumi.StringPtrOutput `pulumi:"cpuType"`
	// Container instance creation time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Map of DNS config options.
	DnsConfigOptions pulumi.MapOutput `pulumi:"dnsConfigOptions"`
	// IP Addresses of DNS Servers.
	DnsNamesServers pulumi.StringArrayOutput `pulumi:"dnsNamesServers"`
	// List of DNS Search Domain.
	DnsSearches pulumi.StringArrayOutput `pulumi:"dnsSearches"`
	// EIP address.
	EipAddress pulumi.StringOutput `pulumi:"eipAddress"`
	// Indicates weather the EIP release or not after instance deleted. Conflict with `existed_eip_ids`.
	EipDeletePolicy pulumi.BoolPtrOutput `pulumi:"eipDeletePolicy"`
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with
	// `existed_eip_ids`.
	EipMaxBandwidthOut pulumi.IntPtrOutput `pulumi:"eipMaxBandwidthOut"`
	// EIP service provider. Default is `BGP`, values `CMCC`,`CTCC`,`CUCC` are available for whitelist customer. Conflict with
	// `existed_eip_ids`.
	EipServiceProvider pulumi.StringPtrOutput `pulumi:"eipServiceProvider"`
	// Existed EIP ID List which used to bind container instance. Conflict with `auto_create_eip` and auto create EIP options.
	ExistedEipIds pulumi.StringArrayOutput `pulumi:"existedEipIds"`
	// Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuCount pulumi.IntPtrOutput `pulumi:"gpuCount"`
	// Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuType pulumi.StringPtrOutput `pulumi:"gpuType"`
	// List of credentials which pull from image registry.
	ImageRegistryCredentials ContainerInstanceImageRegistryCredentialArrayOutput `pulumi:"imageRegistryCredentials"`
	// List of initialized container.
	InitContainers ContainerInstanceInitContainerArrayOutput `pulumi:"initContainers"`
	// Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Memory pulumi.Float64Output `pulumi:"memory"`
	// Name of EKS container instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of NFS volume.
	NfsVolumes ContainerInstanceNfsVolumeArrayOutput `pulumi:"nfsVolumes"`
	// Private IP address.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// Container instance restart policy. Available values: `Always`, `Never`, `OnFailure`.
	RestartPolicy pulumi.StringPtrOutput `pulumi:"restartPolicy"`
	// List of security group id.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Container instance status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Subnet ID of container instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewContainerInstance registers a new resource with the given unique name, arguments, and options.
func NewContainerInstance(ctx *pulumi.Context,
	name string, args *ContainerInstanceArgs, opts ...pulumi.ResourceOption) (*ContainerInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Containers == nil {
		return nil, errors.New("invalid value for required argument 'Containers'")
	}
	if args.Cpu == nil {
		return nil, errors.New("invalid value for required argument 'Cpu'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.SecurityGroups == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroups'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerInstance
	err := ctx.RegisterResource("tencentcloud:Deprecatedeks/containerInstance:ContainerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerInstance gets an existing ContainerInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerInstanceState, opts ...pulumi.ResourceOption) (*ContainerInstance, error) {
	var resource ContainerInstance
	err := ctx.ReadResource("tencentcloud:Deprecatedeks/containerInstance:ContainerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerInstance resources.
type containerInstanceState struct {
	// Indicates whether to create EIP instead of specify existing EIPs. Conflict with `existed_eip_ids`.
	AutoCreateEip *bool `pulumi:"autoCreateEip"`
	// ID of EIP which create automatically.
	AutoCreateEipId *string `pulumi:"autoCreateEipId"`
	// CAM role name authorized to access.
	CamRoleName *string `pulumi:"camRoleName"`
	// List of CBS volume.
	CbsVolumes []ContainerInstanceCbsVolume `pulumi:"cbsVolumes"`
	// List of container.
	Containers []ContainerInstanceContainer `pulumi:"containers"`
	// The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Cpu *float64 `pulumi:"cpu"`
	// Type of cpu, which can set to `intel` or `amd`. It also support backup list like `amd,intel` which indicates using
	// `intel` when `amd` sold out.
	CpuType *string `pulumi:"cpuType"`
	// Container instance creation time.
	CreatedTime *string `pulumi:"createdTime"`
	// Map of DNS config options.
	DnsConfigOptions map[string]interface{} `pulumi:"dnsConfigOptions"`
	// IP Addresses of DNS Servers.
	DnsNamesServers []string `pulumi:"dnsNamesServers"`
	// List of DNS Search Domain.
	DnsSearches []string `pulumi:"dnsSearches"`
	// EIP address.
	EipAddress *string `pulumi:"eipAddress"`
	// Indicates weather the EIP release or not after instance deleted. Conflict with `existed_eip_ids`.
	EipDeletePolicy *bool `pulumi:"eipDeletePolicy"`
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with
	// `existed_eip_ids`.
	EipMaxBandwidthOut *int `pulumi:"eipMaxBandwidthOut"`
	// EIP service provider. Default is `BGP`, values `CMCC`,`CTCC`,`CUCC` are available for whitelist customer. Conflict with
	// `existed_eip_ids`.
	EipServiceProvider *string `pulumi:"eipServiceProvider"`
	// Existed EIP ID List which used to bind container instance. Conflict with `auto_create_eip` and auto create EIP options.
	ExistedEipIds []string `pulumi:"existedEipIds"`
	// Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuCount *int `pulumi:"gpuCount"`
	// Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuType *string `pulumi:"gpuType"`
	// List of credentials which pull from image registry.
	ImageRegistryCredentials []ContainerInstanceImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// List of initialized container.
	InitContainers []ContainerInstanceInitContainer `pulumi:"initContainers"`
	// Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Memory *float64 `pulumi:"memory"`
	// Name of EKS container instance.
	Name *string `pulumi:"name"`
	// List of NFS volume.
	NfsVolumes []ContainerInstanceNfsVolume `pulumi:"nfsVolumes"`
	// Private IP address.
	PrivateIp *string `pulumi:"privateIp"`
	// Container instance restart policy. Available values: `Always`, `Never`, `OnFailure`.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// List of security group id.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Container instance status.
	Status *string `pulumi:"status"`
	// Subnet ID of container instance.
	SubnetId *string `pulumi:"subnetId"`
	// VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type ContainerInstanceState struct {
	// Indicates whether to create EIP instead of specify existing EIPs. Conflict with `existed_eip_ids`.
	AutoCreateEip pulumi.BoolPtrInput
	// ID of EIP which create automatically.
	AutoCreateEipId pulumi.StringPtrInput
	// CAM role name authorized to access.
	CamRoleName pulumi.StringPtrInput
	// List of CBS volume.
	CbsVolumes ContainerInstanceCbsVolumeArrayInput
	// List of container.
	Containers ContainerInstanceContainerArrayInput
	// The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Cpu pulumi.Float64PtrInput
	// Type of cpu, which can set to `intel` or `amd`. It also support backup list like `amd,intel` which indicates using
	// `intel` when `amd` sold out.
	CpuType pulumi.StringPtrInput
	// Container instance creation time.
	CreatedTime pulumi.StringPtrInput
	// Map of DNS config options.
	DnsConfigOptions pulumi.MapInput
	// IP Addresses of DNS Servers.
	DnsNamesServers pulumi.StringArrayInput
	// List of DNS Search Domain.
	DnsSearches pulumi.StringArrayInput
	// EIP address.
	EipAddress pulumi.StringPtrInput
	// Indicates weather the EIP release or not after instance deleted. Conflict with `existed_eip_ids`.
	EipDeletePolicy pulumi.BoolPtrInput
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with
	// `existed_eip_ids`.
	EipMaxBandwidthOut pulumi.IntPtrInput
	// EIP service provider. Default is `BGP`, values `CMCC`,`CTCC`,`CUCC` are available for whitelist customer. Conflict with
	// `existed_eip_ids`.
	EipServiceProvider pulumi.StringPtrInput
	// Existed EIP ID List which used to bind container instance. Conflict with `auto_create_eip` and auto create EIP options.
	ExistedEipIds pulumi.StringArrayInput
	// Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuCount pulumi.IntPtrInput
	// Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuType pulumi.StringPtrInput
	// List of credentials which pull from image registry.
	ImageRegistryCredentials ContainerInstanceImageRegistryCredentialArrayInput
	// List of initialized container.
	InitContainers ContainerInstanceInitContainerArrayInput
	// Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Memory pulumi.Float64PtrInput
	// Name of EKS container instance.
	Name pulumi.StringPtrInput
	// List of NFS volume.
	NfsVolumes ContainerInstanceNfsVolumeArrayInput
	// Private IP address.
	PrivateIp pulumi.StringPtrInput
	// Container instance restart policy. Available values: `Always`, `Never`, `OnFailure`.
	RestartPolicy pulumi.StringPtrInput
	// List of security group id.
	SecurityGroups pulumi.StringArrayInput
	// Container instance status.
	Status pulumi.StringPtrInput
	// Subnet ID of container instance.
	SubnetId pulumi.StringPtrInput
	// VPC ID.
	VpcId pulumi.StringPtrInput
}

func (ContainerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerInstanceState)(nil)).Elem()
}

type containerInstanceArgs struct {
	// Indicates whether to create EIP instead of specify existing EIPs. Conflict with `existed_eip_ids`.
	AutoCreateEip *bool `pulumi:"autoCreateEip"`
	// CAM role name authorized to access.
	CamRoleName *string `pulumi:"camRoleName"`
	// List of CBS volume.
	CbsVolumes []ContainerInstanceCbsVolume `pulumi:"cbsVolumes"`
	// List of container.
	Containers []ContainerInstanceContainer `pulumi:"containers"`
	// The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Cpu float64 `pulumi:"cpu"`
	// Type of cpu, which can set to `intel` or `amd`. It also support backup list like `amd,intel` which indicates using
	// `intel` when `amd` sold out.
	CpuType *string `pulumi:"cpuType"`
	// Map of DNS config options.
	DnsConfigOptions map[string]interface{} `pulumi:"dnsConfigOptions"`
	// IP Addresses of DNS Servers.
	DnsNamesServers []string `pulumi:"dnsNamesServers"`
	// List of DNS Search Domain.
	DnsSearches []string `pulumi:"dnsSearches"`
	// Indicates weather the EIP release or not after instance deleted. Conflict with `existed_eip_ids`.
	EipDeletePolicy *bool `pulumi:"eipDeletePolicy"`
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with
	// `existed_eip_ids`.
	EipMaxBandwidthOut *int `pulumi:"eipMaxBandwidthOut"`
	// EIP service provider. Default is `BGP`, values `CMCC`,`CTCC`,`CUCC` are available for whitelist customer. Conflict with
	// `existed_eip_ids`.
	EipServiceProvider *string `pulumi:"eipServiceProvider"`
	// Existed EIP ID List which used to bind container instance. Conflict with `auto_create_eip` and auto create EIP options.
	ExistedEipIds []string `pulumi:"existedEipIds"`
	// Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuCount *int `pulumi:"gpuCount"`
	// Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuType *string `pulumi:"gpuType"`
	// List of credentials which pull from image registry.
	ImageRegistryCredentials []ContainerInstanceImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// List of initialized container.
	InitContainers []ContainerInstanceInitContainer `pulumi:"initContainers"`
	// Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Memory float64 `pulumi:"memory"`
	// Name of EKS container instance.
	Name *string `pulumi:"name"`
	// List of NFS volume.
	NfsVolumes []ContainerInstanceNfsVolume `pulumi:"nfsVolumes"`
	// Container instance restart policy. Available values: `Always`, `Never`, `OnFailure`.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// List of security group id.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Subnet ID of container instance.
	SubnetId string `pulumi:"subnetId"`
	// VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ContainerInstance resource.
type ContainerInstanceArgs struct {
	// Indicates whether to create EIP instead of specify existing EIPs. Conflict with `existed_eip_ids`.
	AutoCreateEip pulumi.BoolPtrInput
	// CAM role name authorized to access.
	CamRoleName pulumi.StringPtrInput
	// List of CBS volume.
	CbsVolumes ContainerInstanceCbsVolumeArrayInput
	// List of container.
	Containers ContainerInstanceContainerArrayInput
	// The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Cpu pulumi.Float64Input
	// Type of cpu, which can set to `intel` or `amd`. It also support backup list like `amd,intel` which indicates using
	// `intel` when `amd` sold out.
	CpuType pulumi.StringPtrInput
	// Map of DNS config options.
	DnsConfigOptions pulumi.MapInput
	// IP Addresses of DNS Servers.
	DnsNamesServers pulumi.StringArrayInput
	// List of DNS Search Domain.
	DnsSearches pulumi.StringArrayInput
	// Indicates weather the EIP release or not after instance deleted. Conflict with `existed_eip_ids`.
	EipDeletePolicy pulumi.BoolPtrInput
	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with
	// `existed_eip_ids`.
	EipMaxBandwidthOut pulumi.IntPtrInput
	// EIP service provider. Default is `BGP`, values `CMCC`,`CTCC`,`CUCC` are available for whitelist customer. Conflict with
	// `existed_eip_ids`.
	EipServiceProvider pulumi.StringPtrInput
	// Existed EIP ID List which used to bind container instance. Conflict with `auto_create_eip` and auto create EIP options.
	ExistedEipIds pulumi.StringArrayInput
	// Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuCount pulumi.IntPtrInput
	// Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	GpuType pulumi.StringPtrInput
	// List of credentials which pull from image registry.
	ImageRegistryCredentials ContainerInstanceImageRegistryCredentialArrayInput
	// List of initialized container.
	InitContainers ContainerInstanceInitContainerArrayInput
	// Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
	Memory pulumi.Float64Input
	// Name of EKS container instance.
	Name pulumi.StringPtrInput
	// List of NFS volume.
	NfsVolumes ContainerInstanceNfsVolumeArrayInput
	// Container instance restart policy. Available values: `Always`, `Never`, `OnFailure`.
	RestartPolicy pulumi.StringPtrInput
	// List of security group id.
	SecurityGroups pulumi.StringArrayInput
	// Subnet ID of container instance.
	SubnetId pulumi.StringInput
	// VPC ID.
	VpcId pulumi.StringInput
}

func (ContainerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerInstanceArgs)(nil)).Elem()
}

type ContainerInstanceInput interface {
	pulumi.Input

	ToContainerInstanceOutput() ContainerInstanceOutput
	ToContainerInstanceOutputWithContext(ctx context.Context) ContainerInstanceOutput
}

func (*ContainerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerInstance)(nil)).Elem()
}

func (i *ContainerInstance) ToContainerInstanceOutput() ContainerInstanceOutput {
	return i.ToContainerInstanceOutputWithContext(context.Background())
}

func (i *ContainerInstance) ToContainerInstanceOutputWithContext(ctx context.Context) ContainerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerInstanceOutput)
}

// ContainerInstanceArrayInput is an input type that accepts ContainerInstanceArray and ContainerInstanceArrayOutput values.
// You can construct a concrete instance of `ContainerInstanceArrayInput` via:
//
//	ContainerInstanceArray{ ContainerInstanceArgs{...} }
type ContainerInstanceArrayInput interface {
	pulumi.Input

	ToContainerInstanceArrayOutput() ContainerInstanceArrayOutput
	ToContainerInstanceArrayOutputWithContext(context.Context) ContainerInstanceArrayOutput
}

type ContainerInstanceArray []ContainerInstanceInput

func (ContainerInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerInstance)(nil)).Elem()
}

func (i ContainerInstanceArray) ToContainerInstanceArrayOutput() ContainerInstanceArrayOutput {
	return i.ToContainerInstanceArrayOutputWithContext(context.Background())
}

func (i ContainerInstanceArray) ToContainerInstanceArrayOutputWithContext(ctx context.Context) ContainerInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerInstanceArrayOutput)
}

// ContainerInstanceMapInput is an input type that accepts ContainerInstanceMap and ContainerInstanceMapOutput values.
// You can construct a concrete instance of `ContainerInstanceMapInput` via:
//
//	ContainerInstanceMap{ "key": ContainerInstanceArgs{...} }
type ContainerInstanceMapInput interface {
	pulumi.Input

	ToContainerInstanceMapOutput() ContainerInstanceMapOutput
	ToContainerInstanceMapOutputWithContext(context.Context) ContainerInstanceMapOutput
}

type ContainerInstanceMap map[string]ContainerInstanceInput

func (ContainerInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerInstance)(nil)).Elem()
}

func (i ContainerInstanceMap) ToContainerInstanceMapOutput() ContainerInstanceMapOutput {
	return i.ToContainerInstanceMapOutputWithContext(context.Background())
}

func (i ContainerInstanceMap) ToContainerInstanceMapOutputWithContext(ctx context.Context) ContainerInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerInstanceMapOutput)
}

type ContainerInstanceOutput struct{ *pulumi.OutputState }

func (ContainerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerInstance)(nil)).Elem()
}

func (o ContainerInstanceOutput) ToContainerInstanceOutput() ContainerInstanceOutput {
	return o
}

func (o ContainerInstanceOutput) ToContainerInstanceOutputWithContext(ctx context.Context) ContainerInstanceOutput {
	return o
}

// Indicates whether to create EIP instead of specify existing EIPs. Conflict with `existed_eip_ids`.
func (o ContainerInstanceOutput) AutoCreateEip() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.BoolPtrOutput { return v.AutoCreateEip }).(pulumi.BoolPtrOutput)
}

// ID of EIP which create automatically.
func (o ContainerInstanceOutput) AutoCreateEipId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.AutoCreateEipId }).(pulumi.StringOutput)
}

// CAM role name authorized to access.
func (o ContainerInstanceOutput) CamRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringPtrOutput { return v.CamRoleName }).(pulumi.StringPtrOutput)
}

// List of CBS volume.
func (o ContainerInstanceOutput) CbsVolumes() ContainerInstanceCbsVolumeArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) ContainerInstanceCbsVolumeArrayOutput { return v.CbsVolumes }).(ContainerInstanceCbsVolumeArrayOutput)
}

// List of container.
func (o ContainerInstanceOutput) Containers() ContainerInstanceContainerArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) ContainerInstanceContainerArrayOutput { return v.Containers }).(ContainerInstanceContainerArrayOutput)
}

// The number of CPU cores. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
func (o ContainerInstanceOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v *ContainerInstance) pulumi.Float64Output { return v.Cpu }).(pulumi.Float64Output)
}

// Type of cpu, which can set to `intel` or `amd`. It also support backup list like `amd,intel` which indicates using
// `intel` when `amd` sold out.
func (o ContainerInstanceOutput) CpuType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringPtrOutput { return v.CpuType }).(pulumi.StringPtrOutput)
}

// Container instance creation time.
func (o ContainerInstanceOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Map of DNS config options.
func (o ContainerInstanceOutput) DnsConfigOptions() pulumi.MapOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.MapOutput { return v.DnsConfigOptions }).(pulumi.MapOutput)
}

// IP Addresses of DNS Servers.
func (o ContainerInstanceOutput) DnsNamesServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringArrayOutput { return v.DnsNamesServers }).(pulumi.StringArrayOutput)
}

// List of DNS Search Domain.
func (o ContainerInstanceOutput) DnsSearches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringArrayOutput { return v.DnsSearches }).(pulumi.StringArrayOutput)
}

// EIP address.
func (o ContainerInstanceOutput) EipAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.EipAddress }).(pulumi.StringOutput)
}

// Indicates weather the EIP release or not after instance deleted. Conflict with `existed_eip_ids`.
func (o ContainerInstanceOutput) EipDeletePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.BoolPtrOutput { return v.EipDeletePolicy }).(pulumi.BoolPtrOutput)
}

// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). Conflict with
// `existed_eip_ids`.
func (o ContainerInstanceOutput) EipMaxBandwidthOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.IntPtrOutput { return v.EipMaxBandwidthOut }).(pulumi.IntPtrOutput)
}

// EIP service provider. Default is `BGP`, values `CMCC`,`CTCC`,`CUCC` are available for whitelist customer. Conflict with
// `existed_eip_ids`.
func (o ContainerInstanceOutput) EipServiceProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringPtrOutput { return v.EipServiceProvider }).(pulumi.StringPtrOutput)
}

// Existed EIP ID List which used to bind container instance. Conflict with `auto_create_eip` and auto create EIP options.
func (o ContainerInstanceOutput) ExistedEipIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringArrayOutput { return v.ExistedEipIds }).(pulumi.StringArrayOutput)
}

// Count of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
func (o ContainerInstanceOutput) GpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.IntPtrOutput { return v.GpuCount }).(pulumi.IntPtrOutput)
}

// Type of GPU. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
func (o ContainerInstanceOutput) GpuType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringPtrOutput { return v.GpuType }).(pulumi.StringPtrOutput)
}

// List of credentials which pull from image registry.
func (o ContainerInstanceOutput) ImageRegistryCredentials() ContainerInstanceImageRegistryCredentialArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) ContainerInstanceImageRegistryCredentialArrayOutput {
		return v.ImageRegistryCredentials
	}).(ContainerInstanceImageRegistryCredentialArrayOutput)
}

// List of initialized container.
func (o ContainerInstanceOutput) InitContainers() ContainerInstanceInitContainerArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) ContainerInstanceInitContainerArrayOutput { return v.InitContainers }).(ContainerInstanceInitContainerArrayOutput)
}

// Memory size. Check https://intl.cloud.tencent.com/document/product/457/34057 for specification references.
func (o ContainerInstanceOutput) Memory() pulumi.Float64Output {
	return o.ApplyT(func(v *ContainerInstance) pulumi.Float64Output { return v.Memory }).(pulumi.Float64Output)
}

// Name of EKS container instance.
func (o ContainerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of NFS volume.
func (o ContainerInstanceOutput) NfsVolumes() ContainerInstanceNfsVolumeArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) ContainerInstanceNfsVolumeArrayOutput { return v.NfsVolumes }).(ContainerInstanceNfsVolumeArrayOutput)
}

// Private IP address.
func (o ContainerInstanceOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// Container instance restart policy. Available values: `Always`, `Never`, `OnFailure`.
func (o ContainerInstanceOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringPtrOutput { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

// List of security group id.
func (o ContainerInstanceOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Container instance status.
func (o ContainerInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Subnet ID of container instance.
func (o ContainerInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// VPC ID.
func (o ContainerInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ContainerInstanceArrayOutput struct{ *pulumi.OutputState }

func (ContainerInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerInstance)(nil)).Elem()
}

func (o ContainerInstanceArrayOutput) ToContainerInstanceArrayOutput() ContainerInstanceArrayOutput {
	return o
}

func (o ContainerInstanceArrayOutput) ToContainerInstanceArrayOutputWithContext(ctx context.Context) ContainerInstanceArrayOutput {
	return o
}

func (o ContainerInstanceArrayOutput) Index(i pulumi.IntInput) ContainerInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerInstance {
		return vs[0].([]*ContainerInstance)[vs[1].(int)]
	}).(ContainerInstanceOutput)
}

type ContainerInstanceMapOutput struct{ *pulumi.OutputState }

func (ContainerInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerInstance)(nil)).Elem()
}

func (o ContainerInstanceMapOutput) ToContainerInstanceMapOutput() ContainerInstanceMapOutput {
	return o
}

func (o ContainerInstanceMapOutput) ToContainerInstanceMapOutputWithContext(ctx context.Context) ContainerInstanceMapOutput {
	return o
}

func (o ContainerInstanceMapOutput) MapIndex(k pulumi.StringInput) ContainerInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerInstance {
		return vs[0].(map[string]*ContainerInstance)[vs[1].(string)]
	}).(ContainerInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInstanceInput)(nil)).Elem(), &ContainerInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInstanceArrayInput)(nil)).Elem(), ContainerInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInstanceMapInput)(nil)).Elem(), ContainerInstanceMap{})
	pulumi.RegisterOutputType(ContainerInstanceOutput{})
	pulumi.RegisterOutputType(ContainerInstanceArrayOutput{})
	pulumi.RegisterOutputType(ContainerInstanceMapOutput{})
}