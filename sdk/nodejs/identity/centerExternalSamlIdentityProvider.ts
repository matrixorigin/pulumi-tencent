// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CenterExternalSamlIdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing CenterExternalSamlIdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CenterExternalSamlIdentityProviderState, opts?: pulumi.CustomResourceOptions): CenterExternalSamlIdentityProvider {
        return new CenterExternalSamlIdentityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Identity/centerExternalSamlIdentityProvider:CenterExternalSamlIdentityProvider';

    /**
     * Returns true if the given object is an instance of CenterExternalSamlIdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CenterExternalSamlIdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CenterExternalSamlIdentityProvider.__pulumiType;
    }

    /**
     * Acs url.
     */
    public /*out*/ readonly acsUrl!: pulumi.Output<string>;
    /**
     * Certificate ids.
     */
    public /*out*/ readonly certificateIds!: pulumi.Output<string[]>;
    /**
     * Create time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * IdP metadata document (Base64 encoded). Provided by an IdP that supports the SAML 2.0 protocol.
     */
    public readonly encodedMetadataDocument!: pulumi.Output<string>;
    /**
     * IdP identifier.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * IdP login URL.
     */
    public readonly loginUrl!: pulumi.Output<string>;
    /**
     * SSO enabling status. Valid values: Enabled, Disabled (default).
     */
    public readonly ssoStatus!: pulumi.Output<string>;
    /**
     * Update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * X509 certificate in PEM format. If this parameter is specified, all existing certificates will be replaced.
     */
    public readonly x509Certificate!: pulumi.Output<string>;
    /**
     * Space ID.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a CenterExternalSamlIdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CenterExternalSamlIdentityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CenterExternalSamlIdentityProviderArgs | CenterExternalSamlIdentityProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CenterExternalSamlIdentityProviderState | undefined;
            resourceInputs["acsUrl"] = state ? state.acsUrl : undefined;
            resourceInputs["certificateIds"] = state ? state.certificateIds : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["encodedMetadataDocument"] = state ? state.encodedMetadataDocument : undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["loginUrl"] = state ? state.loginUrl : undefined;
            resourceInputs["ssoStatus"] = state ? state.ssoStatus : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["x509Certificate"] = state ? state.x509Certificate : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as CenterExternalSamlIdentityProviderArgs | undefined;
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["encodedMetadataDocument"] = args ? args.encodedMetadataDocument : undefined;
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["loginUrl"] = args ? args.loginUrl : undefined;
            resourceInputs["ssoStatus"] = args ? args.ssoStatus : undefined;
            resourceInputs["x509Certificate"] = args ? args.x509Certificate : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["acsUrl"] = undefined /*out*/;
            resourceInputs["certificateIds"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CenterExternalSamlIdentityProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CenterExternalSamlIdentityProvider resources.
 */
export interface CenterExternalSamlIdentityProviderState {
    /**
     * Acs url.
     */
    acsUrl?: pulumi.Input<string>;
    /**
     * Certificate ids.
     */
    certificateIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Create time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * IdP metadata document (Base64 encoded). Provided by an IdP that supports the SAML 2.0 protocol.
     */
    encodedMetadataDocument?: pulumi.Input<string>;
    /**
     * IdP identifier.
     */
    entityId?: pulumi.Input<string>;
    /**
     * IdP login URL.
     */
    loginUrl?: pulumi.Input<string>;
    /**
     * SSO enabling status. Valid values: Enabled, Disabled (default).
     */
    ssoStatus?: pulumi.Input<string>;
    /**
     * Update time.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * X509 certificate in PEM format. If this parameter is specified, all existing certificates will be replaced.
     */
    x509Certificate?: pulumi.Input<string>;
    /**
     * Space ID.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CenterExternalSamlIdentityProvider resource.
 */
export interface CenterExternalSamlIdentityProviderArgs {
    /**
     * IdP metadata document (Base64 encoded). Provided by an IdP that supports the SAML 2.0 protocol.
     */
    encodedMetadataDocument?: pulumi.Input<string>;
    /**
     * IdP identifier.
     */
    entityId?: pulumi.Input<string>;
    /**
     * IdP login URL.
     */
    loginUrl?: pulumi.Input<string>;
    /**
     * SSO enabling status. Valid values: Enabled, Disabled (default).
     */
    ssoStatus?: pulumi.Input<string>;
    /**
     * X509 certificate in PEM format. If this parameter is specified, all existing certificates will be replaced.
     */
    x509Certificate?: pulumi.Input<string>;
    /**
     * Space ID.
     */
    zoneId: pulumi.Input<string>;
}