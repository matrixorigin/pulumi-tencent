// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class SwitchMasterSlaveOperation extends pulumi.CustomResource {
    /**
     * Get an existing SwitchMasterSlaveOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchMasterSlaveOperationState, opts?: pulumi.CustomResourceOptions): SwitchMasterSlaveOperation {
        return new SwitchMasterSlaveOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/switchMasterSlaveOperation:SwitchMasterSlaveOperation';

    /**
     * Returns true if the given object is an instance of SwitchMasterSlaveOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchMasterSlaveOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchMasterSlaveOperation.__pulumiType;
    }

    /**
     * target instance. Possible values: `first` - first standby; `second` - second standby. The default value is `first`, and
     * only multi-AZ instances support setting it to `second`.
     */
    public readonly dstSlave!: pulumi.Output<string | undefined>;
    /**
     * Whether to force switch. Default is False. Note that if you set the mandatory switch to True, there is a risk of data
     * loss on the instance, so use it with caution.
     */
    public readonly forceSwitch!: pulumi.Output<boolean | undefined>;
    /**
     * instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether to switch within the time window. The default is False, i.e. do not switch within the time window. Note that if
     * the ForceSwitch parameter is set to True, this parameter will not take effect.
     */
    public readonly waitSwitch!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SwitchMasterSlaveOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwitchMasterSlaveOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchMasterSlaveOperationArgs | SwitchMasterSlaveOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchMasterSlaveOperationState | undefined;
            resourceInputs["dstSlave"] = state ? state.dstSlave : undefined;
            resourceInputs["forceSwitch"] = state ? state.forceSwitch : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["waitSwitch"] = state ? state.waitSwitch : undefined;
        } else {
            const args = argsOrState as SwitchMasterSlaveOperationArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["dstSlave"] = args ? args.dstSlave : undefined;
            resourceInputs["forceSwitch"] = args ? args.forceSwitch : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["waitSwitch"] = args ? args.waitSwitch : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchMasterSlaveOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchMasterSlaveOperation resources.
 */
export interface SwitchMasterSlaveOperationState {
    /**
     * target instance. Possible values: `first` - first standby; `second` - second standby. The default value is `first`, and
     * only multi-AZ instances support setting it to `second`.
     */
    dstSlave?: pulumi.Input<string>;
    /**
     * Whether to force switch. Default is False. Note that if you set the mandatory switch to True, there is a risk of data
     * loss on the instance, so use it with caution.
     */
    forceSwitch?: pulumi.Input<boolean>;
    /**
     * instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether to switch within the time window. The default is False, i.e. do not switch within the time window. Note that if
     * the ForceSwitch parameter is set to True, this parameter will not take effect.
     */
    waitSwitch?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SwitchMasterSlaveOperation resource.
 */
export interface SwitchMasterSlaveOperationArgs {
    /**
     * target instance. Possible values: `first` - first standby; `second` - second standby. The default value is `first`, and
     * only multi-AZ instances support setting it to `second`.
     */
    dstSlave?: pulumi.Input<string>;
    /**
     * Whether to force switch. Default is False. Note that if you set the mandatory switch to True, there is a risk of data
     * loss on the instance, so use it with caution.
     */
    forceSwitch?: pulumi.Input<boolean>;
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Whether to switch within the time window. The default is False, i.e. do not switch within the time window. Note that if
     * the ForceSwitch parameter is set to True, this parameter will not take effect.
     */
    waitSwitch?: pulumi.Input<boolean>;
}