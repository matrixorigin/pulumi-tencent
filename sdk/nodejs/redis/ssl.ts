// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Ssl extends pulumi.CustomResource {
    /**
     * Get an existing Ssl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SslState, opts?: pulumi.CustomResourceOptions): Ssl {
        return new Ssl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Redis/ssl:Ssl';

    /**
     * Returns true if the given object is an instance of Ssl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ssl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ssl.__pulumiType;
    }

    /**
     * The ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The SSL configuration status of the instance: `enabled`,`disabled`.
     */
    public readonly sslConfig!: pulumi.Output<string>;

    /**
     * Create a Ssl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SslArgs | SslState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SslState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["sslConfig"] = state ? state.sslConfig : undefined;
        } else {
            const args = argsOrState as SslArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.sslConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslConfig'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["sslConfig"] = args ? args.sslConfig : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ssl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ssl resources.
 */
export interface SslState {
    /**
     * The ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The SSL configuration status of the instance: `enabled`,`disabled`.
     */
    sslConfig?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ssl resource.
 */
export interface SslArgs {
    /**
     * The ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The SSL configuration status of the instance: `enabled`,`disabled`.
     */
    sslConfig: pulumi.Input<string>;
}