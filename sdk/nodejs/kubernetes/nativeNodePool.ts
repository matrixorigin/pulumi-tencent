// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class NativeNodePool extends pulumi.CustomResource {
    /**
     * Get an existing NativeNodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NativeNodePoolState, opts?: pulumi.CustomResourceOptions): NativeNodePool {
        return new NativeNodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Kubernetes/nativeNodePool:NativeNodePool';

    /**
     * Returns true if the given object is an instance of NativeNodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NativeNodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NativeNodePool.__pulumiType;
    }

    /**
     * Node Annotation List.
     */
    public readonly annotations!: pulumi.Output<outputs.Kubernetes.NativeNodePoolAnnotation[]>;
    /**
     * ID of the cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Creation time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Whether to enable deletion protection.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * Node Labels.
     */
    public readonly labels!: pulumi.Output<outputs.Kubernetes.NativeNodePoolLabel[] | undefined>;
    /**
     * Node pool status.
     */
    public /*out*/ readonly lifeState!: pulumi.Output<string>;
    /**
     * Node pool name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Native node pool creation parameters.
     */
    public readonly native!: pulumi.Output<outputs.Kubernetes.NativeNodePoolNative>;
    /**
     * Node tags.
     */
    public readonly tags!: pulumi.Output<outputs.Kubernetes.NativeNodePoolTag[] | undefined>;
    /**
     * Node taint.
     */
    public readonly taints!: pulumi.Output<outputs.Kubernetes.NativeNodePoolTaint[] | undefined>;
    /**
     * Node pool type. Optional value is `Native`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
     */
    public readonly unschedulable!: pulumi.Output<boolean>;

    /**
     * Create a NativeNodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NativeNodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NativeNodePoolArgs | NativeNodePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NativeNodePoolState | undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["lifeState"] = state ? state.lifeState : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["native"] = state ? state.native : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["taints"] = state ? state.taints : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["unschedulable"] = state ? state.unschedulable : undefined;
        } else {
            const args = argsOrState as NativeNodePoolArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.native === undefined) && !opts.urn) {
                throw new Error("Missing required property 'native'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["native"] = args ? args.native : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["taints"] = args ? args.taints : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["unschedulable"] = args ? args.unschedulable : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lifeState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NativeNodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NativeNodePool resources.
 */
export interface NativeNodePoolState {
    /**
     * Node Annotation List.
     */
    annotations?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolAnnotation>[]>;
    /**
     * ID of the cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Creation time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Whether to enable deletion protection.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Node Labels.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolLabel>[]>;
    /**
     * Node pool status.
     */
    lifeState?: pulumi.Input<string>;
    /**
     * Node pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Native node pool creation parameters.
     */
    native?: pulumi.Input<inputs.Kubernetes.NativeNodePoolNative>;
    /**
     * Node tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolTag>[]>;
    /**
     * Node taint.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolTaint>[]>;
    /**
     * Node pool type. Optional value is `Native`.
     */
    type?: pulumi.Input<string>;
    /**
     * Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
     */
    unschedulable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a NativeNodePool resource.
 */
export interface NativeNodePoolArgs {
    /**
     * Node Annotation List.
     */
    annotations?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolAnnotation>[]>;
    /**
     * ID of the cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Whether to enable deletion protection.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Node Labels.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolLabel>[]>;
    /**
     * Node pool name.
     */
    name?: pulumi.Input<string>;
    /**
     * Native node pool creation parameters.
     */
    native: pulumi.Input<inputs.Kubernetes.NativeNodePoolNative>;
    /**
     * Node tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolTag>[]>;
    /**
     * Node taint.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NativeNodePoolTaint>[]>;
    /**
     * Node pool type. Optional value is `Native`.
     */
    type: pulumi.Input<string>;
    /**
     * Whether the node is not schedulable by default. The native node is not aware of it and passes false by default.
     */
    unschedulable?: pulumi.Input<boolean>;
}