// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterMasterAttachment struct {
	pulumi.CustomResourceState

	// ID of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be
	// specified.
	DesiredPodNumbers pulumi.IntArrayOutput `pulumi:"desiredPodNumbers"`
	// Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default
	// to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation
	// Assistant service.
	EnhancedAutomationService pulumi.BoolPtrOutput `pulumi:"enhancedAutomationService"`
	// To specify whether to enable cloud monitor service. Default is TRUE.
	EnhancedMonitorService pulumi.BoolPtrOutput `pulumi:"enhancedMonitorService"`
	// To specify whether to enable cloud security service. Default is TRUE.
	EnhancedSecurityService pulumi.BoolPtrOutput `pulumi:"enhancedSecurityService"`
	// Custom parameters for cluster master component.
	ExtraArgs ClusterMasterAttachmentExtraArgsPtrOutput `pulumi:"extraArgs"`
	// When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed
	// when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance
	// creation interface except that uppercase characters are not supported).
	HostName pulumi.StringPtrOutput `pulumi:"hostName"`
	// ID of the CVM instance, this cvm will reinstall the system.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
	KeyIds pulumi.StringPtrOutput `pulumi:"keyIds"`
	// Advanced Node Settings. commonly used to attach existing instances.
	MasterConfig ClusterMasterAttachmentMasterConfigPtrOutput `pulumi:"masterConfig"`
	// Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER
	// independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum
	// configuration for MASTER_ETCD is 4C8G.
	NodeRole pulumi.StringOutput `pulumi:"nodeRole"`
	// Password to access, should be set if `key_ids` not set.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return
	// value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
}

// NewClusterMasterAttachment registers a new resource with the given unique name, arguments, and options.
func NewClusterMasterAttachment(ctx *pulumi.Context,
	name string, args *ClusterMasterAttachmentArgs, opts ...pulumi.ResourceOption) (*ClusterMasterAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NodeRole == nil {
		return nil, errors.New("invalid value for required argument 'NodeRole'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterMasterAttachment
	err := ctx.RegisterResource("tencentcloud:Kubernetes/clusterMasterAttachment:ClusterMasterAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterMasterAttachment gets an existing ClusterMasterAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterMasterAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterMasterAttachmentState, opts ...pulumi.ResourceOption) (*ClusterMasterAttachment, error) {
	var resource ClusterMasterAttachment
	err := ctx.ReadResource("tencentcloud:Kubernetes/clusterMasterAttachment:ClusterMasterAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterMasterAttachment resources.
type clusterMasterAttachmentState struct {
	// ID of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be
	// specified.
	DesiredPodNumbers []int `pulumi:"desiredPodNumbers"`
	// Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default
	// to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation
	// Assistant service.
	EnhancedAutomationService *bool `pulumi:"enhancedAutomationService"`
	// To specify whether to enable cloud monitor service. Default is TRUE.
	EnhancedMonitorService *bool `pulumi:"enhancedMonitorService"`
	// To specify whether to enable cloud security service. Default is TRUE.
	EnhancedSecurityService *bool `pulumi:"enhancedSecurityService"`
	// Custom parameters for cluster master component.
	ExtraArgs *ClusterMasterAttachmentExtraArgs `pulumi:"extraArgs"`
	// When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed
	// when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance
	// creation interface except that uppercase characters are not supported).
	HostName *string `pulumi:"hostName"`
	// ID of the CVM instance, this cvm will reinstall the system.
	InstanceId *string `pulumi:"instanceId"`
	// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
	KeyIds *string `pulumi:"keyIds"`
	// Advanced Node Settings. commonly used to attach existing instances.
	MasterConfig *ClusterMasterAttachmentMasterConfig `pulumi:"masterConfig"`
	// Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER
	// independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum
	// configuration for MASTER_ETCD is 4C8G.
	NodeRole *string `pulumi:"nodeRole"`
	// Password to access, should be set if `key_ids` not set.
	Password *string `pulumi:"password"`
	// The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return
	// value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
}

type ClusterMasterAttachmentState struct {
	// ID of the cluster.
	ClusterId pulumi.StringPtrInput
	// When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be
	// specified.
	DesiredPodNumbers pulumi.IntArrayInput
	// Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default
	// to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation
	// Assistant service.
	EnhancedAutomationService pulumi.BoolPtrInput
	// To specify whether to enable cloud monitor service. Default is TRUE.
	EnhancedMonitorService pulumi.BoolPtrInput
	// To specify whether to enable cloud security service. Default is TRUE.
	EnhancedSecurityService pulumi.BoolPtrInput
	// Custom parameters for cluster master component.
	ExtraArgs ClusterMasterAttachmentExtraArgsPtrInput
	// When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed
	// when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance
	// creation interface except that uppercase characters are not supported).
	HostName pulumi.StringPtrInput
	// ID of the CVM instance, this cvm will reinstall the system.
	InstanceId pulumi.StringPtrInput
	// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
	KeyIds pulumi.StringPtrInput
	// Advanced Node Settings. commonly used to attach existing instances.
	MasterConfig ClusterMasterAttachmentMasterConfigPtrInput
	// Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER
	// independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum
	// configuration for MASTER_ETCD is 4C8G.
	NodeRole pulumi.StringPtrInput
	// Password to access, should be set if `key_ids` not set.
	Password pulumi.StringPtrInput
	// The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return
	// value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds pulumi.StringArrayInput
}

func (ClusterMasterAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterMasterAttachmentState)(nil)).Elem()
}

type clusterMasterAttachmentArgs struct {
	// ID of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be
	// specified.
	DesiredPodNumbers []int `pulumi:"desiredPodNumbers"`
	// Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default
	// to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation
	// Assistant service.
	EnhancedAutomationService *bool `pulumi:"enhancedAutomationService"`
	// To specify whether to enable cloud monitor service. Default is TRUE.
	EnhancedMonitorService *bool `pulumi:"enhancedMonitorService"`
	// To specify whether to enable cloud security service. Default is TRUE.
	EnhancedSecurityService *bool `pulumi:"enhancedSecurityService"`
	// Custom parameters for cluster master component.
	ExtraArgs *ClusterMasterAttachmentExtraArgs `pulumi:"extraArgs"`
	// When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed
	// when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance
	// creation interface except that uppercase characters are not supported).
	HostName *string `pulumi:"hostName"`
	// ID of the CVM instance, this cvm will reinstall the system.
	InstanceId string `pulumi:"instanceId"`
	// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
	KeyIds *string `pulumi:"keyIds"`
	// Advanced Node Settings. commonly used to attach existing instances.
	MasterConfig *ClusterMasterAttachmentMasterConfig `pulumi:"masterConfig"`
	// Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER
	// independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum
	// configuration for MASTER_ETCD is 4C8G.
	NodeRole string `pulumi:"nodeRole"`
	// Password to access, should be set if `key_ids` not set.
	Password *string `pulumi:"password"`
	// The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return
	// value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
}

// The set of arguments for constructing a ClusterMasterAttachment resource.
type ClusterMasterAttachmentArgs struct {
	// ID of the cluster.
	ClusterId pulumi.StringInput
	// When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be
	// specified.
	DesiredPodNumbers pulumi.IntArrayInput
	// Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default
	// to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation
	// Assistant service.
	EnhancedAutomationService pulumi.BoolPtrInput
	// To specify whether to enable cloud monitor service. Default is TRUE.
	EnhancedMonitorService pulumi.BoolPtrInput
	// To specify whether to enable cloud security service. Default is TRUE.
	EnhancedSecurityService pulumi.BoolPtrInput
	// Custom parameters for cluster master component.
	ExtraArgs ClusterMasterAttachmentExtraArgsPtrInput
	// When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed
	// when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance
	// creation interface except that uppercase characters are not supported).
	HostName pulumi.StringPtrInput
	// ID of the CVM instance, this cvm will reinstall the system.
	InstanceId pulumi.StringInput
	// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
	KeyIds pulumi.StringPtrInput
	// Advanced Node Settings. commonly used to attach existing instances.
	MasterConfig ClusterMasterAttachmentMasterConfigPtrInput
	// Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER
	// independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum
	// configuration for MASTER_ETCD is 4C8G.
	NodeRole pulumi.StringInput
	// Password to access, should be set if `key_ids` not set.
	Password pulumi.StringPtrInput
	// The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return
	// value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds pulumi.StringArrayInput
}

func (ClusterMasterAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterMasterAttachmentArgs)(nil)).Elem()
}

type ClusterMasterAttachmentInput interface {
	pulumi.Input

	ToClusterMasterAttachmentOutput() ClusterMasterAttachmentOutput
	ToClusterMasterAttachmentOutputWithContext(ctx context.Context) ClusterMasterAttachmentOutput
}

func (*ClusterMasterAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterMasterAttachment)(nil)).Elem()
}

func (i *ClusterMasterAttachment) ToClusterMasterAttachmentOutput() ClusterMasterAttachmentOutput {
	return i.ToClusterMasterAttachmentOutputWithContext(context.Background())
}

func (i *ClusterMasterAttachment) ToClusterMasterAttachmentOutputWithContext(ctx context.Context) ClusterMasterAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMasterAttachmentOutput)
}

// ClusterMasterAttachmentArrayInput is an input type that accepts ClusterMasterAttachmentArray and ClusterMasterAttachmentArrayOutput values.
// You can construct a concrete instance of `ClusterMasterAttachmentArrayInput` via:
//
//	ClusterMasterAttachmentArray{ ClusterMasterAttachmentArgs{...} }
type ClusterMasterAttachmentArrayInput interface {
	pulumi.Input

	ToClusterMasterAttachmentArrayOutput() ClusterMasterAttachmentArrayOutput
	ToClusterMasterAttachmentArrayOutputWithContext(context.Context) ClusterMasterAttachmentArrayOutput
}

type ClusterMasterAttachmentArray []ClusterMasterAttachmentInput

func (ClusterMasterAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterMasterAttachment)(nil)).Elem()
}

func (i ClusterMasterAttachmentArray) ToClusterMasterAttachmentArrayOutput() ClusterMasterAttachmentArrayOutput {
	return i.ToClusterMasterAttachmentArrayOutputWithContext(context.Background())
}

func (i ClusterMasterAttachmentArray) ToClusterMasterAttachmentArrayOutputWithContext(ctx context.Context) ClusterMasterAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMasterAttachmentArrayOutput)
}

// ClusterMasterAttachmentMapInput is an input type that accepts ClusterMasterAttachmentMap and ClusterMasterAttachmentMapOutput values.
// You can construct a concrete instance of `ClusterMasterAttachmentMapInput` via:
//
//	ClusterMasterAttachmentMap{ "key": ClusterMasterAttachmentArgs{...} }
type ClusterMasterAttachmentMapInput interface {
	pulumi.Input

	ToClusterMasterAttachmentMapOutput() ClusterMasterAttachmentMapOutput
	ToClusterMasterAttachmentMapOutputWithContext(context.Context) ClusterMasterAttachmentMapOutput
}

type ClusterMasterAttachmentMap map[string]ClusterMasterAttachmentInput

func (ClusterMasterAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterMasterAttachment)(nil)).Elem()
}

func (i ClusterMasterAttachmentMap) ToClusterMasterAttachmentMapOutput() ClusterMasterAttachmentMapOutput {
	return i.ToClusterMasterAttachmentMapOutputWithContext(context.Background())
}

func (i ClusterMasterAttachmentMap) ToClusterMasterAttachmentMapOutputWithContext(ctx context.Context) ClusterMasterAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMasterAttachmentMapOutput)
}

type ClusterMasterAttachmentOutput struct{ *pulumi.OutputState }

func (ClusterMasterAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterMasterAttachment)(nil)).Elem()
}

func (o ClusterMasterAttachmentOutput) ToClusterMasterAttachmentOutput() ClusterMasterAttachmentOutput {
	return o
}

func (o ClusterMasterAttachmentOutput) ToClusterMasterAttachmentOutputWithContext(ctx context.Context) ClusterMasterAttachmentOutput {
	return o
}

// ID of the cluster.
func (o ClusterMasterAttachmentOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// When the node belongs to the podCIDR size customization mode, the maximum number of pods running on the node can be
// specified.
func (o ClusterMasterAttachmentOutput) DesiredPodNumbers() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.IntArrayOutput { return v.DesiredPodNumbers }).(pulumi.IntArrayOutput)
}

// Activate TencentCloud Automation Tools (TAT) service. If this parameter is not specified, the public image will default
// to enabling the Cloud Automation Assistant service, while other images will default to not enabling the Cloud Automation
// Assistant service.
func (o ClusterMasterAttachmentOutput) EnhancedAutomationService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.BoolPtrOutput { return v.EnhancedAutomationService }).(pulumi.BoolPtrOutput)
}

// To specify whether to enable cloud monitor service. Default is TRUE.
func (o ClusterMasterAttachmentOutput) EnhancedMonitorService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.BoolPtrOutput { return v.EnhancedMonitorService }).(pulumi.BoolPtrOutput)
}

// To specify whether to enable cloud security service. Default is TRUE.
func (o ClusterMasterAttachmentOutput) EnhancedSecurityService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.BoolPtrOutput { return v.EnhancedSecurityService }).(pulumi.BoolPtrOutput)
}

// Custom parameters for cluster master component.
func (o ClusterMasterAttachmentOutput) ExtraArgs() ClusterMasterAttachmentExtraArgsPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) ClusterMasterAttachmentExtraArgsPtrOutput { return v.ExtraArgs }).(ClusterMasterAttachmentExtraArgsPtrOutput)
}

// When reinstalling the system, you can specify the HostName of the instance to be modified (this parameter must be passed
// when the cluster is in HostName mode, and the rule name should be consistent with the HostName of the CVM instance
// creation interface except that uppercase characters are not supported).
func (o ClusterMasterAttachmentOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringPtrOutput { return v.HostName }).(pulumi.StringPtrOutput)
}

// ID of the CVM instance, this cvm will reinstall the system.
func (o ClusterMasterAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
func (o ClusterMasterAttachmentOutput) KeyIds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringPtrOutput { return v.KeyIds }).(pulumi.StringPtrOutput)
}

// Advanced Node Settings. commonly used to attach existing instances.
func (o ClusterMasterAttachmentOutput) MasterConfig() ClusterMasterAttachmentMasterConfigPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) ClusterMasterAttachmentMasterConfigPtrOutput { return v.MasterConfig }).(ClusterMasterAttachmentMasterConfigPtrOutput)
}

// Node role, values: MASTER_ETCD, WORKER. MASTER_ETCD needs to be specified only when creating an INDEPENDENT_CLUSTER
// independent cluster. The number of MASTER_ETCD nodes is 3-7, and it is recommended to have an odd number. The minimum
// configuration for MASTER_ETCD is 4C8G.
func (o ClusterMasterAttachmentOutput) NodeRole() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringOutput { return v.NodeRole }).(pulumi.StringOutput)
}

// Password to access, should be set if `key_ids` not set.
func (o ClusterMasterAttachmentOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The security group to which the instance belongs. This parameter can be obtained by calling the sgId field in the return
// value of DescribeSecureGroups. If this parameter is not specified, the default security group will be bound.
func (o ClusterMasterAttachmentOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterMasterAttachment) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

type ClusterMasterAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ClusterMasterAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterMasterAttachment)(nil)).Elem()
}

func (o ClusterMasterAttachmentArrayOutput) ToClusterMasterAttachmentArrayOutput() ClusterMasterAttachmentArrayOutput {
	return o
}

func (o ClusterMasterAttachmentArrayOutput) ToClusterMasterAttachmentArrayOutputWithContext(ctx context.Context) ClusterMasterAttachmentArrayOutput {
	return o
}

func (o ClusterMasterAttachmentArrayOutput) Index(i pulumi.IntInput) ClusterMasterAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterMasterAttachment {
		return vs[0].([]*ClusterMasterAttachment)[vs[1].(int)]
	}).(ClusterMasterAttachmentOutput)
}

type ClusterMasterAttachmentMapOutput struct{ *pulumi.OutputState }

func (ClusterMasterAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterMasterAttachment)(nil)).Elem()
}

func (o ClusterMasterAttachmentMapOutput) ToClusterMasterAttachmentMapOutput() ClusterMasterAttachmentMapOutput {
	return o
}

func (o ClusterMasterAttachmentMapOutput) ToClusterMasterAttachmentMapOutputWithContext(ctx context.Context) ClusterMasterAttachmentMapOutput {
	return o
}

func (o ClusterMasterAttachmentMapOutput) MapIndex(k pulumi.StringInput) ClusterMasterAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterMasterAttachment {
		return vs[0].(map[string]*ClusterMasterAttachment)[vs[1].(string)]
	}).(ClusterMasterAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMasterAttachmentInput)(nil)).Elem(), &ClusterMasterAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMasterAttachmentArrayInput)(nil)).Elem(), ClusterMasterAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMasterAttachmentMapInput)(nil)).Elem(), ClusterMasterAttachmentMap{})
	pulumi.RegisterOutputType(ClusterMasterAttachmentOutput{})
	pulumi.RegisterOutputType(ClusterMasterAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ClusterMasterAttachmentMapOutput{})
}