// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSlowLogUserHostStats(args: GetSlowLogUserHostStatsArgs, opts?: pulumi.InvokeOptions): Promise<GetSlowLogUserHostStatsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getSlowLogUserHostStats:getSlowLogUserHostStats", {
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "md5": args.md5,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlowLogUserHostStats.
 */
export interface GetSlowLogUserHostStatsArgs {
    endTime: string;
    instanceId: string;
    md5?: string;
    product?: string;
    resultOutputFile?: string;
    startTime: string;
}

/**
 * A collection of values returned by getSlowLogUserHostStats.
 */
export interface GetSlowLogUserHostStatsResult {
    readonly endTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly items: outputs.Dbbrain.GetSlowLogUserHostStatsItem[];
    readonly md5?: string;
    readonly product?: string;
    readonly resultOutputFile?: string;
    readonly startTime: string;
}
export function getSlowLogUserHostStatsOutput(args: GetSlowLogUserHostStatsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSlowLogUserHostStatsResult> {
    return pulumi.output(args).apply((a: any) => getSlowLogUserHostStats(a, opts))
}

/**
 * A collection of arguments for invoking getSlowLogUserHostStats.
 */
export interface GetSlowLogUserHostStatsOutputArgs {
    endTime: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    md5?: pulumi.Input<string>;
    product?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    startTime: pulumi.Input<string>;
}