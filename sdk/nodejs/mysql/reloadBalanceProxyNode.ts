// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ReloadBalanceProxyNode extends pulumi.CustomResource {
    /**
     * Get an existing ReloadBalanceProxyNode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReloadBalanceProxyNodeState, opts?: pulumi.CustomResourceOptions): ReloadBalanceProxyNode {
        return new ReloadBalanceProxyNode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/reloadBalanceProxyNode:ReloadBalanceProxyNode';

    /**
     * Returns true if the given object is an instance of ReloadBalanceProxyNode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReloadBalanceProxyNode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReloadBalanceProxyNode.__pulumiType;
    }

    /**
     * Proxy address id.
     */
    public readonly proxyAddressId!: pulumi.Output<string | undefined>;
    /**
     * Proxy id.
     */
    public readonly proxyGroupId!: pulumi.Output<string>;

    /**
     * Create a ReloadBalanceProxyNode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReloadBalanceProxyNodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReloadBalanceProxyNodeArgs | ReloadBalanceProxyNodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReloadBalanceProxyNodeState | undefined;
            resourceInputs["proxyAddressId"] = state ? state.proxyAddressId : undefined;
            resourceInputs["proxyGroupId"] = state ? state.proxyGroupId : undefined;
        } else {
            const args = argsOrState as ReloadBalanceProxyNodeArgs | undefined;
            if ((!args || args.proxyGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proxyGroupId'");
            }
            resourceInputs["proxyAddressId"] = args ? args.proxyAddressId : undefined;
            resourceInputs["proxyGroupId"] = args ? args.proxyGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReloadBalanceProxyNode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReloadBalanceProxyNode resources.
 */
export interface ReloadBalanceProxyNodeState {
    /**
     * Proxy address id.
     */
    proxyAddressId?: pulumi.Input<string>;
    /**
     * Proxy id.
     */
    proxyGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReloadBalanceProxyNode resource.
 */
export interface ReloadBalanceProxyNodeArgs {
    /**
     * Proxy address id.
     */
    proxyAddressId?: pulumi.Input<string>;
    /**
     * Proxy id.
     */
    proxyGroupId: pulumi.Input<string>;
}