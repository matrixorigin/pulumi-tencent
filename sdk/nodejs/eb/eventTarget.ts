// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class EventTarget extends pulumi.CustomResource {
    /**
     * Get an existing EventTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTargetState, opts?: pulumi.CustomResourceOptions): EventTarget {
        return new EventTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Eb/eventTarget:EventTarget';

    /**
     * Returns true if the given object is an instance of EventTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTarget.__pulumiType;
    }

    /**
     * event bus id.
     */
    public readonly eventBusId!: pulumi.Output<string>;
    /**
     * event rule id.
     */
    public readonly ruleId!: pulumi.Output<string>;
    /**
     * target description.
     */
    public readonly targetDescription!: pulumi.Output<outputs.Eb.EventTargetTargetDescription>;
    /**
     * target type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a EventTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTargetArgs | EventTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventTargetState | undefined;
            resourceInputs["eventBusId"] = state ? state.eventBusId : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["targetDescription"] = state ? state.targetDescription : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as EventTargetArgs | undefined;
            if ((!args || args.eventBusId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventBusId'");
            }
            if ((!args || args.ruleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleId'");
            }
            if ((!args || args.targetDescription === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetDescription'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["eventBusId"] = args ? args.eventBusId : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["targetDescription"] = args ? args.targetDescription : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTarget resources.
 */
export interface EventTargetState {
    /**
     * event bus id.
     */
    eventBusId?: pulumi.Input<string>;
    /**
     * event rule id.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * target description.
     */
    targetDescription?: pulumi.Input<inputs.Eb.EventTargetTargetDescription>;
    /**
     * target type.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTarget resource.
 */
export interface EventTargetArgs {
    /**
     * event bus id.
     */
    eventBusId: pulumi.Input<string>;
    /**
     * event rule id.
     */
    ruleId: pulumi.Input<string>;
    /**
     * target description.
     */
    targetDescription: pulumi.Input<inputs.Eb.EventTargetTargetDescription>;
    /**
     * target type.
     */
    type: pulumi.Input<string>;
}