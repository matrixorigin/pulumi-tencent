// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RoGroupLoadOperation extends pulumi.CustomResource {
    /**
     * Get an existing RoGroupLoadOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoGroupLoadOperationState, opts?: pulumi.CustomResourceOptions): RoGroupLoadOperation {
        return new RoGroupLoadOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/roGroupLoadOperation:RoGroupLoadOperation';

    /**
     * Returns true if the given object is an instance of RoGroupLoadOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoGroupLoadOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoGroupLoadOperation.__pulumiType;
    }

    /**
     * The ID of the RO group, in the format: cdbrg-c1nl9rpv.
     */
    public readonly roGroupId!: pulumi.Output<string>;

    /**
     * Create a RoGroupLoadOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoGroupLoadOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoGroupLoadOperationArgs | RoGroupLoadOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoGroupLoadOperationState | undefined;
            resourceInputs["roGroupId"] = state ? state.roGroupId : undefined;
        } else {
            const args = argsOrState as RoGroupLoadOperationArgs | undefined;
            if ((!args || args.roGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roGroupId'");
            }
            resourceInputs["roGroupId"] = args ? args.roGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RoGroupLoadOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoGroupLoadOperation resources.
 */
export interface RoGroupLoadOperationState {
    /**
     * The ID of the RO group, in the format: cdbrg-c1nl9rpv.
     */
    roGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoGroupLoadOperation resource.
 */
export interface RoGroupLoadOperationArgs {
    /**
     * The ID of the RO group, in the format: cdbrg-c1nl9rpv.
     */
    roGroupId: pulumi.Input<string>;
}