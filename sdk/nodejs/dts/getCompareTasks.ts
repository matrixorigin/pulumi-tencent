// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getCompareTasks(args: GetCompareTasksArgs, opts?: pulumi.InvokeOptions): Promise<GetCompareTasksResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dts/getCompareTasks:getCompareTasks", {
        "jobId": args.jobId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getCompareTasks.
 */
export interface GetCompareTasksArgs {
    jobId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getCompareTasks.
 */
export interface GetCompareTasksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jobId: string;
    readonly lists: outputs.Dts.GetCompareTasksList[];
    readonly resultOutputFile?: string;
}
export function getCompareTasksOutput(args: GetCompareTasksOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCompareTasksResult> {
    return pulumi.output(args).apply((a: any) => getCompareTasks(a, opts))
}

/**
 * A collection of arguments for invoking getCompareTasks.
 */
export interface GetCompareTasksOutputArgs {
    jobId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}