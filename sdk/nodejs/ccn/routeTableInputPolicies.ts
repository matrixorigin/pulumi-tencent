// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class RouteTableInputPolicies extends pulumi.CustomResource {
    /**
     * Get an existing RouteTableInputPolicies resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableInputPoliciesState, opts?: pulumi.CustomResourceOptions): RouteTableInputPolicies {
        return new RouteTableInputPolicies(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ccn/routeTableInputPolicies:RouteTableInputPolicies';

    /**
     * Returns true if the given object is an instance of RouteTableInputPolicies.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTableInputPolicies {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTableInputPolicies.__pulumiType;
    }

    /**
     * CCN Instance ID.
     */
    public readonly ccnId!: pulumi.Output<string>;
    /**
     * Routing reception strategy.
     */
    public readonly policies!: pulumi.Output<outputs.Ccn.RouteTableInputPoliciesPolicy[] | undefined>;
    /**
     * CCN Route table ID.
     */
    public readonly routeTableId!: pulumi.Output<string>;

    /**
     * Create a RouteTableInputPolicies resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableInputPoliciesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableInputPoliciesArgs | RouteTableInputPoliciesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTableInputPoliciesState | undefined;
            resourceInputs["ccnId"] = state ? state.ccnId : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
        } else {
            const args = argsOrState as RouteTableInputPoliciesArgs | undefined;
            if ((!args || args.ccnId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ccnId'");
            }
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            resourceInputs["ccnId"] = args ? args.ccnId : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteTableInputPolicies.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTableInputPolicies resources.
 */
export interface RouteTableInputPoliciesState {
    /**
     * CCN Instance ID.
     */
    ccnId?: pulumi.Input<string>;
    /**
     * Routing reception strategy.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.Ccn.RouteTableInputPoliciesPolicy>[]>;
    /**
     * CCN Route table ID.
     */
    routeTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTableInputPolicies resource.
 */
export interface RouteTableInputPoliciesArgs {
    /**
     * CCN Instance ID.
     */
    ccnId: pulumi.Input<string>;
    /**
     * Routing reception strategy.
     */
    policies?: pulumi.Input<pulumi.Input<inputs.Ccn.RouteTableInputPoliciesPolicy>[]>;
    /**
     * CCN Route table ID.
     */
    routeTableId: pulumi.Input<string>;
}