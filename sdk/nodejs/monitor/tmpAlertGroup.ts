// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class TmpAlertGroup extends pulumi.CustomResource {
    /**
     * Get an existing TmpAlertGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TmpAlertGroupState, opts?: pulumi.CustomResourceOptions): TmpAlertGroup {
        return new TmpAlertGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/tmpAlertGroup:TmpAlertGroup';

    /**
     * Returns true if the given object is an instance of TmpAlertGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TmpAlertGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TmpAlertGroup.__pulumiType;
    }

    /**
     * Tencent cloud notification template id list.
     */
    public readonly ampReceivers!: pulumi.Output<string[] | undefined>;
    /**
     * User custom notification template, such as webhook, alertmanager.
     */
    public readonly customReceiver!: pulumi.Output<outputs.Monitor.TmpAlertGroupCustomReceiver | undefined>;
    /**
     * Alarm group id.
     */
    public /*out*/ readonly groupId!: pulumi.Output<string>;
    /**
     * Unique alert group name.
     */
    public readonly groupName!: pulumi.Output<string | undefined>;
    /**
     * Instance id.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * Alert message send interval, default 1 hour.
     */
    public readonly repeatInterval!: pulumi.Output<string | undefined>;
    /**
     * A list of alert rules.
     */
    public readonly rules!: pulumi.Output<outputs.Monitor.TmpAlertGroupRule[] | undefined>;

    /**
     * Create a TmpAlertGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TmpAlertGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TmpAlertGroupArgs | TmpAlertGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TmpAlertGroupState | undefined;
            resourceInputs["ampReceivers"] = state ? state.ampReceivers : undefined;
            resourceInputs["customReceiver"] = state ? state.customReceiver : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["repeatInterval"] = state ? state.repeatInterval : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as TmpAlertGroupArgs | undefined;
            resourceInputs["ampReceivers"] = args ? args.ampReceivers : undefined;
            resourceInputs["customReceiver"] = args ? args.customReceiver : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["repeatInterval"] = args ? args.repeatInterval : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["groupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TmpAlertGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TmpAlertGroup resources.
 */
export interface TmpAlertGroupState {
    /**
     * Tencent cloud notification template id list.
     */
    ampReceivers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User custom notification template, such as webhook, alertmanager.
     */
    customReceiver?: pulumi.Input<inputs.Monitor.TmpAlertGroupCustomReceiver>;
    /**
     * Alarm group id.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Unique alert group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Alert message send interval, default 1 hour.
     */
    repeatInterval?: pulumi.Input<string>;
    /**
     * A list of alert rules.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpAlertGroupRule>[]>;
}

/**
 * The set of arguments for constructing a TmpAlertGroup resource.
 */
export interface TmpAlertGroupArgs {
    /**
     * Tencent cloud notification template id list.
     */
    ampReceivers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User custom notification template, such as webhook, alertmanager.
     */
    customReceiver?: pulumi.Input<inputs.Monitor.TmpAlertGroupCustomReceiver>;
    /**
     * Unique alert group name.
     */
    groupName?: pulumi.Input<string>;
    /**
     * Instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Alert message send interval, default 1 hour.
     */
    repeatInterval?: pulumi.Input<string>;
    /**
     * A list of alert rules.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpAlertGroupRule>[]>;
}