// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcr/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Indicate to delete the COS bucket which is auto-created with the instance or not.
     */
    public readonly deleteBucket!: pulumi.Output<boolean | undefined>;
    /**
     * Instance expiration time (prepaid).
     */
    public /*out*/ readonly expiredAt!: pulumi.Output<string>;
    /**
     * Length of time to purchase an instance (in month). Must set when registry_charge_type is prepaid.
     */
    public readonly instanceChargeTypePrepaidPeriod!: pulumi.Output<number | undefined>;
    /**
     * Auto renewal flag. 1: manual renewal, 2: automatic renewal, 3: no renewal and no notification. Must set when
     * registry_charge_type is prepaid.
     */
    public readonly instanceChargeTypePrepaidRenewFlag!: pulumi.Output<number | undefined>;
    /**
     * TCR types. Valid values are: `standard`, `basic`, `premium`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Internal address for access of the TCR instance.
     */
    public /*out*/ readonly internalEndPoint!: pulumi.Output<string>;
    /**
     * Name of the TCR instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Control public network access.
     */
    public readonly openPublicOperation!: pulumi.Output<boolean | undefined>;
    /**
     * Public address for access of the TCR instance.
     */
    public /*out*/ readonly publicDomain!: pulumi.Output<string>;
    /**
     * Status of the TCR instance public network access.
     */
    public /*out*/ readonly publicStatus!: pulumi.Output<string>;
    /**
     * Charge type of instance. 1: postpaid; 2: prepaid. Default is postpaid.
     */
    public readonly registryChargeType!: pulumi.Output<number | undefined>;
    /**
     * Specify List of instance Replications, premium only. The available [source region
     * list](https://www.tencentcloud.com/document/api/1051/41101) is here.
     */
    public readonly replications!: pulumi.Output<outputs.Tcr.InstanceReplication[] | undefined>;
    /**
     * Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
     */
    public readonly securityPolicies!: pulumi.Output<outputs.Tcr.InstanceSecurityPolicy[] | undefined>;
    /**
     * Status of the TCR instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The available tags within this TCR instance.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["deleteBucket"] = state ? state.deleteBucket : undefined;
            resourceInputs["expiredAt"] = state ? state.expiredAt : undefined;
            resourceInputs["instanceChargeTypePrepaidPeriod"] = state ? state.instanceChargeTypePrepaidPeriod : undefined;
            resourceInputs["instanceChargeTypePrepaidRenewFlag"] = state ? state.instanceChargeTypePrepaidRenewFlag : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["internalEndPoint"] = state ? state.internalEndPoint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["openPublicOperation"] = state ? state.openPublicOperation : undefined;
            resourceInputs["publicDomain"] = state ? state.publicDomain : undefined;
            resourceInputs["publicStatus"] = state ? state.publicStatus : undefined;
            resourceInputs["registryChargeType"] = state ? state.registryChargeType : undefined;
            resourceInputs["replications"] = state ? state.replications : undefined;
            resourceInputs["securityPolicies"] = state ? state.securityPolicies : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["deleteBucket"] = args ? args.deleteBucket : undefined;
            resourceInputs["instanceChargeTypePrepaidPeriod"] = args ? args.instanceChargeTypePrepaidPeriod : undefined;
            resourceInputs["instanceChargeTypePrepaidRenewFlag"] = args ? args.instanceChargeTypePrepaidRenewFlag : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openPublicOperation"] = args ? args.openPublicOperation : undefined;
            resourceInputs["registryChargeType"] = args ? args.registryChargeType : undefined;
            resourceInputs["replications"] = args ? args.replications : undefined;
            resourceInputs["securityPolicies"] = args ? args.securityPolicies : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["expiredAt"] = undefined /*out*/;
            resourceInputs["internalEndPoint"] = undefined /*out*/;
            resourceInputs["publicDomain"] = undefined /*out*/;
            resourceInputs["publicStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Indicate to delete the COS bucket which is auto-created with the instance or not.
     */
    deleteBucket?: pulumi.Input<boolean>;
    /**
     * Instance expiration time (prepaid).
     */
    expiredAt?: pulumi.Input<string>;
    /**
     * Length of time to purchase an instance (in month). Must set when registry_charge_type is prepaid.
     */
    instanceChargeTypePrepaidPeriod?: pulumi.Input<number>;
    /**
     * Auto renewal flag. 1: manual renewal, 2: automatic renewal, 3: no renewal and no notification. Must set when
     * registry_charge_type is prepaid.
     */
    instanceChargeTypePrepaidRenewFlag?: pulumi.Input<number>;
    /**
     * TCR types. Valid values are: `standard`, `basic`, `premium`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Internal address for access of the TCR instance.
     */
    internalEndPoint?: pulumi.Input<string>;
    /**
     * Name of the TCR instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Control public network access.
     */
    openPublicOperation?: pulumi.Input<boolean>;
    /**
     * Public address for access of the TCR instance.
     */
    publicDomain?: pulumi.Input<string>;
    /**
     * Status of the TCR instance public network access.
     */
    publicStatus?: pulumi.Input<string>;
    /**
     * Charge type of instance. 1: postpaid; 2: prepaid. Default is postpaid.
     */
    registryChargeType?: pulumi.Input<number>;
    /**
     * Specify List of instance Replications, premium only. The available [source region
     * list](https://www.tencentcloud.com/document/api/1051/41101) is here.
     */
    replications?: pulumi.Input<pulumi.Input<inputs.Tcr.InstanceReplication>[]>;
    /**
     * Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
     */
    securityPolicies?: pulumi.Input<pulumi.Input<inputs.Tcr.InstanceSecurityPolicy>[]>;
    /**
     * Status of the TCR instance.
     */
    status?: pulumi.Input<string>;
    /**
     * The available tags within this TCR instance.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Indicate to delete the COS bucket which is auto-created with the instance or not.
     */
    deleteBucket?: pulumi.Input<boolean>;
    /**
     * Length of time to purchase an instance (in month). Must set when registry_charge_type is prepaid.
     */
    instanceChargeTypePrepaidPeriod?: pulumi.Input<number>;
    /**
     * Auto renewal flag. 1: manual renewal, 2: automatic renewal, 3: no renewal and no notification. Must set when
     * registry_charge_type is prepaid.
     */
    instanceChargeTypePrepaidRenewFlag?: pulumi.Input<number>;
    /**
     * TCR types. Valid values are: `standard`, `basic`, `premium`.
     */
    instanceType: pulumi.Input<string>;
    /**
     * Name of the TCR instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Control public network access.
     */
    openPublicOperation?: pulumi.Input<boolean>;
    /**
     * Charge type of instance. 1: postpaid; 2: prepaid. Default is postpaid.
     */
    registryChargeType?: pulumi.Input<number>;
    /**
     * Specify List of instance Replications, premium only. The available [source region
     * list](https://www.tencentcloud.com/document/api/1051/41101) is here.
     */
    replications?: pulumi.Input<pulumi.Input<inputs.Tcr.InstanceReplication>[]>;
    /**
     * Public network access allowlist policies of the TCR instance. Only available when `open_public_operation` is `true`.
     */
    securityPolicies?: pulumi.Input<pulumi.Input<inputs.Tcr.InstanceSecurityPolicy>[]>;
    /**
     * The available tags within this TCR instance.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}