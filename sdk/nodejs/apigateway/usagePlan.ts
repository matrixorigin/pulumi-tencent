// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class UsagePlan extends pulumi.CustomResource {
    /**
     * Get an existing UsagePlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsagePlanState, opts?: pulumi.CustomResourceOptions): UsagePlan {
        return new UsagePlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:ApiGateway/usagePlan:UsagePlan';

    /**
     * Returns true if the given object is an instance of UsagePlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UsagePlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UsagePlan.__pulumiType;
    }

    /**
     * Attach API keys list.
     */
    public /*out*/ readonly attachApiKeys!: pulumi.Output<string[]>;
    /**
     * Attach service and API list.
     */
    public /*out*/ readonly attachLists!: pulumi.Output<outputs.ApiGateway.UsagePlanAttachList[]>;
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
     */
    public readonly maxRequestNum!: pulumi.Output<number | undefined>;
    /**
     * Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
     */
    public readonly maxRequestNumPreSec!: pulumi.Output<number | undefined>;
    /**
     * Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    /**
     * Custom usage plan description.
     */
    public readonly usagePlanDesc!: pulumi.Output<string | undefined>;
    /**
     * Custom usage plan name.
     */
    public readonly usagePlanName!: pulumi.Output<string>;

    /**
     * Create a UsagePlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UsagePlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UsagePlanArgs | UsagePlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UsagePlanState | undefined;
            resourceInputs["attachApiKeys"] = state ? state.attachApiKeys : undefined;
            resourceInputs["attachLists"] = state ? state.attachLists : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["maxRequestNum"] = state ? state.maxRequestNum : undefined;
            resourceInputs["maxRequestNumPreSec"] = state ? state.maxRequestNumPreSec : undefined;
            resourceInputs["modifyTime"] = state ? state.modifyTime : undefined;
            resourceInputs["usagePlanDesc"] = state ? state.usagePlanDesc : undefined;
            resourceInputs["usagePlanName"] = state ? state.usagePlanName : undefined;
        } else {
            const args = argsOrState as UsagePlanArgs | undefined;
            if ((!args || args.usagePlanName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'usagePlanName'");
            }
            resourceInputs["maxRequestNum"] = args ? args.maxRequestNum : undefined;
            resourceInputs["maxRequestNumPreSec"] = args ? args.maxRequestNumPreSec : undefined;
            resourceInputs["usagePlanDesc"] = args ? args.usagePlanDesc : undefined;
            resourceInputs["usagePlanName"] = args ? args.usagePlanName : undefined;
            resourceInputs["attachApiKeys"] = undefined /*out*/;
            resourceInputs["attachLists"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["modifyTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UsagePlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UsagePlan resources.
 */
export interface UsagePlanState {
    /**
     * Attach API keys list.
     */
    attachApiKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Attach service and API list.
     */
    attachLists?: pulumi.Input<pulumi.Input<inputs.ApiGateway.UsagePlanAttachList>[]>;
    /**
     * Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
     */
    maxRequestNum?: pulumi.Input<number>;
    /**
     * Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
     */
    maxRequestNumPreSec?: pulumi.Input<number>;
    /**
     * Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
     */
    modifyTime?: pulumi.Input<string>;
    /**
     * Custom usage plan description.
     */
    usagePlanDesc?: pulumi.Input<string>;
    /**
     * Custom usage plan name.
     */
    usagePlanName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UsagePlan resource.
 */
export interface UsagePlanArgs {
    /**
     * Total number of requests allowed. Valid values: -1, [1,99999999]. The default value is -1, which indicates no limit.
     */
    maxRequestNum?: pulumi.Input<number>;
    /**
     * Limit of requests per second. Valid values: -1, [1,2000]. The default value is -1, which indicates no limit.
     */
    maxRequestNumPreSec?: pulumi.Input<number>;
    /**
     * Custom usage plan description.
     */
    usagePlanDesc?: pulumi.Input<string>;
    /**
     * Custom usage plan name.
     */
    usagePlanName: pulumi.Input<string>;
}