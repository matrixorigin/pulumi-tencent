// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getNodes(args?: GetNodesArgs, opts?: pulumi.InvokeOptions): Promise<GetNodesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Organization/getNodes:getNodes", {
        "resultOutputFile": args.resultOutputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getNodes.
 */
export interface GetNodesArgs {
    resultOutputFile?: string;
    tags?: inputs.Organization.GetNodesTag[];
}

/**
 * A collection of values returned by getNodes.
 */
export interface GetNodesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly items: outputs.Organization.GetNodesItem[];
    readonly resultOutputFile?: string;
    readonly tags?: outputs.Organization.GetNodesTag[];
}
export function getNodesOutput(args?: GetNodesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNodesResult> {
    return pulumi.output(args).apply((a: any) => getNodes(a, opts))
}

/**
 * A collection of arguments for invoking getNodes.
 */
export interface GetNodesOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.Organization.GetNodesTagArgs>[]>;
}
