// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getUserDomains(args?: GetUserDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Waf/getUserDomains:getUserDomains", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserDomains.
 */
export interface GetUserDomainsArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getUserDomains.
 */
export interface GetUserDomainsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly usersInfos: outputs.Waf.GetUserDomainsUsersInfo[];
}
export function getUserDomainsOutput(args?: GetUserDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserDomainsResult> {
    return pulumi.output(args).apply((a: any) => getUserDomains(a, opts))
}

/**
 * A collection of arguments for invoking getUserDomains.
 */
export interface GetUserDomainsOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}