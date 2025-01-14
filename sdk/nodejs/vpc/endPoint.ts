// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class EndPoint extends pulumi.CustomResource {
    /**
     * Get an existing EndPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndPointState, opts?: pulumi.CustomResourceOptions): EndPoint {
        return new EndPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/endPoint:EndPoint';

    /**
     * Returns true if the given object is an instance of EndPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndPoint.__pulumiType;
    }

    /**
     * CDC instance ID.
     */
    public /*out*/ readonly cdcId!: pulumi.Output<string>;
    /**
     * Create Time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Name of endpoint.
     */
    public readonly endPointName!: pulumi.Output<string>;
    /**
     * APPID.
     */
    public /*out*/ readonly endPointOwner!: pulumi.Output<string>;
    /**
     * ID of endpoint service.
     */
    public readonly endPointServiceId!: pulumi.Output<string>;
    /**
     * VIP of endpoint ip.
     */
    public readonly endPointVip!: pulumi.Output<string | undefined>;
    /**
     * Ordered security groups associated with the endpoint.
     */
    public readonly securityGroupsIds!: pulumi.Output<string[]>;
    /**
     * state of end point.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * ID of subnet instance.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * ID of vpc instance.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a EndPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndPointArgs | EndPointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndPointState | undefined;
            resourceInputs["cdcId"] = state ? state.cdcId : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["endPointName"] = state ? state.endPointName : undefined;
            resourceInputs["endPointOwner"] = state ? state.endPointOwner : undefined;
            resourceInputs["endPointServiceId"] = state ? state.endPointServiceId : undefined;
            resourceInputs["endPointVip"] = state ? state.endPointVip : undefined;
            resourceInputs["securityGroupsIds"] = state ? state.securityGroupsIds : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as EndPointArgs | undefined;
            if ((!args || args.endPointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endPointName'");
            }
            if ((!args || args.endPointServiceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endPointServiceId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["endPointName"] = args ? args.endPointName : undefined;
            resourceInputs["endPointServiceId"] = args ? args.endPointServiceId : undefined;
            resourceInputs["endPointVip"] = args ? args.endPointVip : undefined;
            resourceInputs["securityGroupsIds"] = args ? args.securityGroupsIds : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["cdcId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endPointOwner"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndPoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndPoint resources.
 */
export interface EndPointState {
    /**
     * CDC instance ID.
     */
    cdcId?: pulumi.Input<string>;
    /**
     * Create Time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Name of endpoint.
     */
    endPointName?: pulumi.Input<string>;
    /**
     * APPID.
     */
    endPointOwner?: pulumi.Input<string>;
    /**
     * ID of endpoint service.
     */
    endPointServiceId?: pulumi.Input<string>;
    /**
     * VIP of endpoint ip.
     */
    endPointVip?: pulumi.Input<string>;
    /**
     * Ordered security groups associated with the endpoint.
     */
    securityGroupsIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * state of end point.
     */
    state?: pulumi.Input<string>;
    /**
     * ID of subnet instance.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * ID of vpc instance.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndPoint resource.
 */
export interface EndPointArgs {
    /**
     * Name of endpoint.
     */
    endPointName: pulumi.Input<string>;
    /**
     * ID of endpoint service.
     */
    endPointServiceId: pulumi.Input<string>;
    /**
     * VIP of endpoint ip.
     */
    endPointVip?: pulumi.Input<string>;
    /**
     * Ordered security groups associated with the endpoint.
     */
    securityGroupsIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of subnet instance.
     */
    subnetId: pulumi.Input<string>;
    /**
     * ID of vpc instance.
     */
    vpcId: pulumi.Input<string>;
}
