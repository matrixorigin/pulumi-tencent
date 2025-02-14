// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ElasticPublicIpv6 extends pulumi.CustomResource {
    /**
     * Get an existing ElasticPublicIpv6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ElasticPublicIpv6State, opts?: pulumi.CustomResourceOptions): ElasticPublicIpv6 {
        return new ElasticPublicIpv6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Classic/elasticPublicIpv6:ElasticPublicIpv6';

    /**
     * Returns true if the given object is an instance of ElasticPublicIpv6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ElasticPublicIpv6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ElasticPublicIpv6.__pulumiType;
    }

    /**
     * Bandwidth package id, move the account up, and you need to pass in the ipv6 address to apply for bandwidth package
     * charging mode.
     */
    public readonly bandwidthPackageId!: pulumi.Output<string>;
    /**
     * Network billing model. IPV6 currently supports `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. The default network
     * charging mode is `TRAFFIC_POSTPAID_BY_HOUR`.
     */
    public readonly internetChargeType!: pulumi.Output<string>;
    /**
     * Bandwidth in Mbps. Default is 1Mbps.
     */
    public readonly internetMaxBandwidthOut!: pulumi.Output<number>;
    /**
     * IPV6 addresses that require public network access.
     */
    public readonly ip6Address!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a ElasticPublicIpv6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ElasticPublicIpv6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ElasticPublicIpv6Args | ElasticPublicIpv6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ElasticPublicIpv6State | undefined;
            resourceInputs["bandwidthPackageId"] = state ? state.bandwidthPackageId : undefined;
            resourceInputs["internetChargeType"] = state ? state.internetChargeType : undefined;
            resourceInputs["internetMaxBandwidthOut"] = state ? state.internetMaxBandwidthOut : undefined;
            resourceInputs["ip6Address"] = state ? state.ip6Address : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ElasticPublicIpv6Args | undefined;
            if ((!args || args.ip6Address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip6Address'");
            }
            resourceInputs["bandwidthPackageId"] = args ? args.bandwidthPackageId : undefined;
            resourceInputs["internetChargeType"] = args ? args.internetChargeType : undefined;
            resourceInputs["internetMaxBandwidthOut"] = args ? args.internetMaxBandwidthOut : undefined;
            resourceInputs["ip6Address"] = args ? args.ip6Address : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ElasticPublicIpv6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ElasticPublicIpv6 resources.
 */
export interface ElasticPublicIpv6State {
    /**
     * Bandwidth package id, move the account up, and you need to pass in the ipv6 address to apply for bandwidth package
     * charging mode.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Network billing model. IPV6 currently supports `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. The default network
     * charging mode is `TRAFFIC_POSTPAID_BY_HOUR`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * Bandwidth in Mbps. Default is 1Mbps.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * IPV6 addresses that require public network access.
     */
    ip6Address?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ElasticPublicIpv6 resource.
 */
export interface ElasticPublicIpv6Args {
    /**
     * Bandwidth package id, move the account up, and you need to pass in the ipv6 address to apply for bandwidth package
     * charging mode.
     */
    bandwidthPackageId?: pulumi.Input<string>;
    /**
     * Network billing model. IPV6 currently supports `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. The default network
     * charging mode is `TRAFFIC_POSTPAID_BY_HOUR`.
     */
    internetChargeType?: pulumi.Input<string>;
    /**
     * Bandwidth in Mbps. Default is 1Mbps.
     */
    internetMaxBandwidthOut?: pulumi.Input<number>;
    /**
     * IPV6 addresses that require public network access.
     */
    ip6Address: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
