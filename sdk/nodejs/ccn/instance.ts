// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
    public static readonly __pulumiType = 'tencentcloud:Ccn/instance:Instance';

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
     * The speed limit type. Valid values: `INTER_REGION_LIMIT`, `OUTER_REGION_LIMIT`. `OUTER_REGION_LIMIT` represents the
     * regional export speed limit, `INTER_REGION_LIMIT` is the inter-regional speed limit. The default is
     * `OUTER_REGION_LIMIT`.
     */
    public readonly bandwidthLimitType!: pulumi.Output<string | undefined>;
    /**
     * Billing mode. Valid values: `PREPAID`, `POSTPAID`. `PREPAID` means prepaid, which means annual and monthly subscription,
     * `POSTPAID` means post-payment, which means billing by volume. The default is `POSTPAID`. The prepaid model only supports
     * inter-regional speed limit, and the post-paid model supports inter-regional speed limit and regional export speed limit.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Creation time of resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of CCN, and maximum length does not exceed 100 bytes.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Number of attached instances.
     */
    public /*out*/ readonly instanceCount!: pulumi.Output<number>;
    /**
     * Name of the CCN to be queried, and maximum length does not exceed 60 bytes.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * CCN service quality, 'PT': Platinum, 'AU': Gold, 'AG': Silver. The default is 'AU'.
     */
    public readonly qos!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable the equivalent routing function. `true`: enabled, `false`: disabled.
     */
    public readonly routeEcmpFlag!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable the routing overlap function. `true`: enabled, `false`: disabled.
     */
    public readonly routeOverlapFlag!: pulumi.Output<boolean | undefined>;
    /**
     * States of instance. Valid values: `ISOLATED`(arrears) and `AVAILABLE`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Instance tag.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["bandwidthLimitType"] = state ? state.bandwidthLimitType : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceCount"] = state ? state.instanceCount : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["qos"] = state ? state.qos : undefined;
            resourceInputs["routeEcmpFlag"] = state ? state.routeEcmpFlag : undefined;
            resourceInputs["routeOverlapFlag"] = state ? state.routeOverlapFlag : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            resourceInputs["bandwidthLimitType"] = args ? args.bandwidthLimitType : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["qos"] = args ? args.qos : undefined;
            resourceInputs["routeEcmpFlag"] = args ? args.routeEcmpFlag : undefined;
            resourceInputs["routeOverlapFlag"] = args ? args.routeOverlapFlag : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["instanceCount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
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
     * The speed limit type. Valid values: `INTER_REGION_LIMIT`, `OUTER_REGION_LIMIT`. `OUTER_REGION_LIMIT` represents the
     * regional export speed limit, `INTER_REGION_LIMIT` is the inter-regional speed limit. The default is
     * `OUTER_REGION_LIMIT`.
     */
    bandwidthLimitType?: pulumi.Input<string>;
    /**
     * Billing mode. Valid values: `PREPAID`, `POSTPAID`. `PREPAID` means prepaid, which means annual and monthly subscription,
     * `POSTPAID` means post-payment, which means billing by volume. The default is `POSTPAID`. The prepaid model only supports
     * inter-regional speed limit, and the post-paid model supports inter-regional speed limit and regional export speed limit.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Creation time of resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of CCN, and maximum length does not exceed 100 bytes.
     */
    description?: pulumi.Input<string>;
    /**
     * Number of attached instances.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * Name of the CCN to be queried, and maximum length does not exceed 60 bytes.
     */
    name?: pulumi.Input<string>;
    /**
     * CCN service quality, 'PT': Platinum, 'AU': Gold, 'AG': Silver. The default is 'AU'.
     */
    qos?: pulumi.Input<string>;
    /**
     * Whether to enable the equivalent routing function. `true`: enabled, `false`: disabled.
     */
    routeEcmpFlag?: pulumi.Input<boolean>;
    /**
     * Whether to enable the routing overlap function. `true`: enabled, `false`: disabled.
     */
    routeOverlapFlag?: pulumi.Input<boolean>;
    /**
     * States of instance. Valid values: `ISOLATED`(arrears) and `AVAILABLE`.
     */
    state?: pulumi.Input<string>;
    /**
     * Instance tag.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The speed limit type. Valid values: `INTER_REGION_LIMIT`, `OUTER_REGION_LIMIT`. `OUTER_REGION_LIMIT` represents the
     * regional export speed limit, `INTER_REGION_LIMIT` is the inter-regional speed limit. The default is
     * `OUTER_REGION_LIMIT`.
     */
    bandwidthLimitType?: pulumi.Input<string>;
    /**
     * Billing mode. Valid values: `PREPAID`, `POSTPAID`. `PREPAID` means prepaid, which means annual and monthly subscription,
     * `POSTPAID` means post-payment, which means billing by volume. The default is `POSTPAID`. The prepaid model only supports
     * inter-regional speed limit, and the post-paid model supports inter-regional speed limit and regional export speed limit.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Description of CCN, and maximum length does not exceed 100 bytes.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the CCN to be queried, and maximum length does not exceed 60 bytes.
     */
    name?: pulumi.Input<string>;
    /**
     * CCN service quality, 'PT': Platinum, 'AU': Gold, 'AG': Silver. The default is 'AU'.
     */
    qos?: pulumi.Input<string>;
    /**
     * Whether to enable the equivalent routing function. `true`: enabled, `false`: disabled.
     */
    routeEcmpFlag?: pulumi.Input<boolean>;
    /**
     * Whether to enable the routing overlap function. `true`: enabled, `false`: disabled.
     */
    routeOverlapFlag?: pulumi.Input<boolean>;
    /**
     * Instance tag.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
