// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class RuleEngine extends pulumi.CustomResource {
    /**
     * Get an existing RuleEngine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleEngineState, opts?: pulumi.CustomResourceOptions): RuleEngine {
        return new RuleEngine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Teo/ruleEngine:RuleEngine';

    /**
     * Returns true if the given object is an instance of RuleEngine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleEngine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleEngine.__pulumiType;
    }

    /**
     * Rule ID.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * The rule name (1 to 255 characters).
     */
    public readonly ruleName!: pulumi.Output<string>;
    /**
     * Rule items list.
     */
    public readonly rules!: pulumi.Output<outputs.Teo.RuleEngineRule[]>;
    /**
     * Rule status. Values: `enable`: Enabled; `disable`: Disabled.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * rule tag list.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * ID of the site.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a RuleEngine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleEngineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleEngineArgs | RuleEngineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleEngineState | undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as RuleEngineArgs | undefined;
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleEngine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleEngine resources.
 */
export interface RuleEngineState {
    /**
     * Rule ID.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * The rule name (1 to 255 characters).
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Rule items list.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.Teo.RuleEngineRule>[]>;
    /**
     * Rule status. Values: `enable`: Enabled; `disable`: Disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * rule tag list.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the site.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RuleEngine resource.
 */
export interface RuleEngineArgs {
    /**
     * The rule name (1 to 255 characters).
     */
    ruleName: pulumi.Input<string>;
    /**
     * Rule items list.
     */
    rules: pulumi.Input<pulumi.Input<inputs.Teo.RuleEngineRule>[]>;
    /**
     * Rule status. Values: `enable`: Enabled; `disable`: Disabled.
     */
    status: pulumi.Input<string>;
    /**
     * rule tag list.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the site.
     */
    zoneId: pulumi.Input<string>;
}