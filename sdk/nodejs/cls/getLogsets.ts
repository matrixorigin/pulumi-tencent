// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLogsets(args?: GetLogsetsArgs, opts?: pulumi.InvokeOptions): Promise<GetLogsetsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cls/getLogsets:getLogsets", {
        "filters": args.filters,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getLogsets.
 */
export interface GetLogsetsArgs {
    filters?: inputs.Cls.GetLogsetsFilter[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getLogsets.
 */
export interface GetLogsetsResult {
    readonly filters?: outputs.Cls.GetLogsetsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly logsets: outputs.Cls.GetLogsetsLogset[];
    readonly resultOutputFile?: string;
}
export function getLogsetsOutput(args?: GetLogsetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogsetsResult> {
    return pulumi.output(args).apply((a: any) => getLogsets(a, opts))
}

/**
 * A collection of arguments for invoking getLogsets.
 */
export interface GetLogsetsOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Cls.GetLogsetsFilterArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
}