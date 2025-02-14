// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class OrgShareUnit extends pulumi.CustomResource {
    /**
     * Get an existing OrgShareUnit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgShareUnitState, opts?: pulumi.CustomResourceOptions): OrgShareUnit {
        return new OrgShareUnit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Organization/orgShareUnit:OrgShareUnit';

    /**
     * Returns true if the given object is an instance of OrgShareUnit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgShareUnit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgShareUnit.__pulumiType;
    }

    /**
     * Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
     */
    public readonly area!: pulumi.Output<string>;
    /**
     * Shared unit description. Up to 128 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
     * 3-128 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
     */
    public /*out*/ readonly unitId!: pulumi.Output<string>;

    /**
     * Create a OrgShareUnit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgShareUnitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgShareUnitArgs | OrgShareUnitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgShareUnitState | undefined;
            resourceInputs["area"] = state ? state.area : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["unitId"] = state ? state.unitId : undefined;
        } else {
            const args = argsOrState as OrgShareUnitArgs | undefined;
            if ((!args || args.area === undefined) && !opts.urn) {
                throw new Error("Missing required property 'area'");
            }
            resourceInputs["area"] = args ? args.area : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["unitId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgShareUnit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgShareUnit resources.
 */
export interface OrgShareUnitState {
    /**
     * Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
     */
    area?: pulumi.Input<string>;
    /**
     * Shared unit description. Up to 128 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
     * 3-128 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
     */
    unitId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgShareUnit resource.
 */
export interface OrgShareUnitArgs {
    /**
     * Shared unit region. The regions that support sharing can be obtained through the DescribeShareAreas interface.
     */
    area: pulumi.Input<string>;
    /**
     * Shared unit description. Up to 128 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Shared unit name. It only supports a combination of uppercase and lowercase letters, numbers, -, and _, with a length of
     * 3-128 characters.
     */
    name?: pulumi.Input<string>;
}
