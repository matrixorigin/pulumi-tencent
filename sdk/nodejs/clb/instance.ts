// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Clb/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
     */
    public readonly addressIpVersion!: pulumi.Output<string>;
    /**
     * Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string | undefined>;
    /**
     * Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
     */
    public readonly clbName!: pulumi.Output<string>;
    /**
     * The virtual service address table of the CLB.
     */
    public /*out*/ readonly clbVips!: pulumi.Output<string[]>;
    /**
     * Whether to enable delete protection.
     */
    public readonly deleteProtect!: pulumi.Output<boolean | undefined>;
    /**
     * Domain name of the CLB instance.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * If create dynamic vip CLB instance, `true` or `false`.
     */
    public readonly dynamicVip!: pulumi.Output<boolean | undefined>;
    /**
     * Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
     */
    public readonly internetBandwidthMaxOut!: pulumi.Output<number>;
    /**
     * Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
     * `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
     * backend instance security group.
     */
    public readonly loadBalancerPassToTarget!: pulumi.Output<boolean | undefined>;
    /**
     * The id of log set.
     */
    public readonly logSetId!: pulumi.Output<string | undefined>;
    /**
     * The id of log topic.
     */
    public readonly logTopicId!: pulumi.Output<string | undefined>;
    /**
     * Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
     */
    public readonly masterZoneId!: pulumi.Output<string | undefined>;
    /**
     * Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * ID of the project within the CLB instance, `0` - Default Project.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
     */
    public readonly securityGroups!: pulumi.Output<string[] | undefined>;
    /**
     * This parameter is required to create LCU-supported instances. Values:`SLA`: Super Large 4. When you have activated Super
     * Large models, `SLA` refers to Super Large 4; `clb.c2.medium`: Standard; `clb.c3.small`: Advanced 1; `clb.c3.medium`:
     * Advanced 1; `clb.c4.small`: Super Large 1; `clb.c4.medium`: Super Large 2; `clb.c4.large`: Super Large 3;
     * `clb.c4.xlarge`: Super Large 4. For more details, see [Instance
     * Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
     */
    public readonly slaType!: pulumi.Output<string>;
    /**
     * Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
     * traffic when the master is down.
     */
    public readonly slaveZoneId!: pulumi.Output<string | undefined>;
    /**
     * Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is
     * untraceable, please import resource `tencentcloud_clb_snat_ip` to handle fixed ips.
     */
    public readonly snatIps!: pulumi.Output<outputs.Clb.InstanceSnatIp[] | undefined>;
    /**
     * Indicates whether Binding IPs of other VPCs feature switch.
     */
    public readonly snatPro!: pulumi.Output<boolean | undefined>;
    /**
     * Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * The available tags within this CLB.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    public readonly targetRegionInfoRegion!: pulumi.Output<string>;
    /**
     * Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    public readonly targetRegionInfoVpcId!: pulumi.Output<string>;
    /**
     * Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
     * Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
     * (BANDWIDTH_PACKAGE).
     */
    public readonly vipIsp!: pulumi.Output<string>;
    /**
     * VPC ID of the CLB.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Available zone id, only applicable to open CLB.
     */
    public readonly zoneId!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["clbName"] = state ? state.clbName : undefined;
            resourceInputs["clbVips"] = state ? state.clbVips : undefined;
            resourceInputs["deleteProtect"] = state ? state.deleteProtect : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["dynamicVip"] = state ? state.dynamicVip : undefined;
            resourceInputs["internetBandwidthMaxOut"] = state ? state.internetBandwidthMaxOut : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["loadBalancerPassToTarget"] = state ? state.loadBalancerPassToTarget : undefined;
            resourceInputs["logSetId"] = state ? state.logSetId : undefined;
            resourceInputs["logTopicId"] = state ? state.logTopicId : undefined;
            resourceInputs["masterZoneId"] = state ? state.masterZoneId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["slaType"] = state ? state.slaType : undefined;
            resourceInputs["slaveZoneId"] = state ? state.slaveZoneId : undefined;
            resourceInputs["snatIps"] = state ? state.snatIps : undefined;
            resourceInputs["snatPro"] = state ? state.snatPro : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["targetRegionInfoRegion"] = state ? state.targetRegionInfoRegion : undefined;
            resourceInputs["targetRegionInfoVpcId"] = state ? state.targetRegionInfoVpcId : undefined;
            resourceInputs["vipIsp"] = state ? state.vipIsp : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.clbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clbName'");
            }
            if ((!args || args.networkType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkType'");
            }
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["clbName"] = args ? args.clbName : undefined;
            resourceInputs["deleteProtect"] = args ? args.deleteProtect : undefined;
            resourceInputs["dynamicVip"] = args ? args.dynamicVip : undefined;
            resourceInputs["internetBandwidthMaxOut"] = args ? args.internetBandwidthMaxOut : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["loadBalancerPassToTarget"] = args ? args.loadBalancerPassToTarget : undefined;
            resourceInputs["logSetId"] = args ? args.logSetId : undefined;
            resourceInputs["logTopicId"] = args ? args.logTopicId : undefined;
            resourceInputs["masterZoneId"] = args ? args.masterZoneId : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["slaType"] = args ? args.slaType : undefined;
            resourceInputs["slaveZoneId"] = args ? args.slaveZoneId : undefined;
            resourceInputs["snatIps"] = args ? args.snatIps : undefined;
            resourceInputs["snatPro"] = args ? args.snatPro : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetRegionInfoRegion"] = args ? args.targetRegionInfoRegion : undefined;
            resourceInputs["targetRegionInfoVpcId"] = args ? args.targetRegionInfoVpcId : undefined;
            resourceInputs["vipIsp"] = args ? args.vipIsp : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["clbVips"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
     */
    clbName?: pulumi.Input<string>;
    /**
     * The virtual service address table of the CLB.
     */
    clbVips?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable delete protection.
     */
    deleteProtect?: pulumi.Input<boolean>;
    /**
     * Domain name of the CLB instance.
     */
    domain?: pulumi.Input<string>;
    /**
     * If create dynamic vip CLB instance, `true` or `false`.
     */
    dynamicVip?: pulumi.Input<boolean>;
    /**
     * Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
     */
    internetBandwidthMaxOut?: pulumi.Input<number>;
    /**
     * Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
     * `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
     * backend instance security group.
     */
    loadBalancerPassToTarget?: pulumi.Input<boolean>;
    /**
     * The id of log set.
     */
    logSetId?: pulumi.Input<string>;
    /**
     * The id of log topic.
     */
    logTopicId?: pulumi.Input<string>;
    /**
     * Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
     */
    masterZoneId?: pulumi.Input<string>;
    /**
     * Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * ID of the project within the CLB instance, `0` - Default Project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This parameter is required to create LCU-supported instances. Values:`SLA`: Super Large 4. When you have activated Super
     * Large models, `SLA` refers to Super Large 4; `clb.c2.medium`: Standard; `clb.c3.small`: Advanced 1; `clb.c3.medium`:
     * Advanced 1; `clb.c4.small`: Super Large 1; `clb.c4.medium`: Super Large 2; `clb.c4.large`: Super Large 3;
     * `clb.c4.xlarge`: Super Large 4. For more details, see [Instance
     * Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
     */
    slaType?: pulumi.Input<string>;
    /**
     * Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
     * traffic when the master is down.
     */
    slaveZoneId?: pulumi.Input<string>;
    /**
     * Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is
     * untraceable, please import resource `tencentcloud_clb_snat_ip` to handle fixed ips.
     */
    snatIps?: pulumi.Input<pulumi.Input<inputs.Clb.InstanceSnatIp>[]>;
    /**
     * Indicates whether Binding IPs of other VPCs feature switch.
     */
    snatPro?: pulumi.Input<boolean>;
    /**
     * Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The available tags within this CLB.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoRegion?: pulumi.Input<string>;
    /**
     * Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoVpcId?: pulumi.Input<string>;
    /**
     * Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
     * Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
     * (BANDWIDTH_PACKAGE).
     */
    vipIsp?: pulumi.Input<string>;
    /**
     * VPC ID of the CLB.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Available zone id, only applicable to open CLB.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * IP version, only applicable to open CLB. Valid values are `ipv4`, `ipv6` and `IPv6FullChain`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * Bandwidth package id. If set, the `internet_charge_type` must be `BANDWIDTH_PACKAGE`.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
     */
    clbName: pulumi.Input<string>;
    /**
     * Whether to enable delete protection.
     */
    deleteProtect?: pulumi.Input<boolean>;
    /**
     * If create dynamic vip CLB instance, `true` or `false`.
     */
    dynamicVip?: pulumi.Input<boolean>;
    /**
     * Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.
     */
    internetBandwidthMaxOut?: pulumi.Input<number>;
    /**
     * Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`,
     * `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * Whether the target allow flow come from clb. If value is true, only check security group of clb, or check both clb and
     * backend instance security group.
     */
    loadBalancerPassToTarget?: pulumi.Input<boolean>;
    /**
     * The id of log set.
     */
    logSetId?: pulumi.Input<string>;
    /**
     * The id of log topic.
     */
    logTopicId?: pulumi.Input<string>;
    /**
     * Setting master zone id of cross available zone disaster recovery, only applicable to open CLB.
     */
    masterZoneId?: pulumi.Input<string>;
    /**
     * Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
     */
    networkType: pulumi.Input<string>;
    /**
     * ID of the project within the CLB instance, `0` - Default Project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This parameter is required to create LCU-supported instances. Values:`SLA`: Super Large 4. When you have activated Super
     * Large models, `SLA` refers to Super Large 4; `clb.c2.medium`: Standard; `clb.c3.small`: Advanced 1; `clb.c3.medium`:
     * Advanced 1; `clb.c4.small`: Super Large 1; `clb.c4.medium`: Super Large 2; `clb.c4.large`: Super Large 3;
     * `clb.c4.xlarge`: Super Large 4. For more details, see [Instance
     * Specifications](https://intl.cloud.tencent.com/document/product/214/84689?from_cn_redirect=1).
     */
    slaType?: pulumi.Input<string>;
    /**
     * Setting slave zone id of cross available zone disaster recovery, only applicable to open CLB. this zone will undertake
     * traffic when the master is down.
     */
    slaveZoneId?: pulumi.Input<string>;
    /**
     * Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is
     * untraceable, please import resource `tencentcloud_clb_snat_ip` to handle fixed ips.
     */
    snatIps?: pulumi.Input<pulumi.Input<inputs.Clb.InstanceSnatIp>[]>;
    /**
     * Indicates whether Binding IPs of other VPCs feature switch.
     */
    snatPro?: pulumi.Input<boolean>;
    /**
     * Subnet ID of the CLB. Effective only for CLB within the VPC. Only supports `INTERNAL` CLBs. Default is `ipv4`.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The available tags within this CLB.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Region information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoRegion?: pulumi.Input<string>;
    /**
     * Vpc information of backend services are attached the CLB instance. Only supports `OPEN` CLBs.
     */
    targetRegionInfoVpcId?: pulumi.Input<string>;
    /**
     * Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China
     * Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing
     * (BANDWIDTH_PACKAGE).
     */
    vipIsp?: pulumi.Input<string>;
    /**
     * VPC ID of the CLB.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Available zone id, only applicable to open CLB.
     */
    zoneId?: pulumi.Input<string>;
}