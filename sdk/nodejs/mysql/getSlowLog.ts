// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSlowLog(args: GetSlowLogArgs, opts?: pulumi.InvokeOptions): Promise<GetSlowLogResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Mysql/getSlowLog:getSlowLog", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getSlowLog.
 */
export interface GetSlowLogArgs {
    instanceId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getSlowLog.
 */
export interface GetSlowLogResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly items: outputs.Mysql.GetSlowLogItem[];
    readonly resultOutputFile?: string;
}
export function getSlowLogOutput(args: GetSlowLogOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSlowLogResult> {
    return pulumi.output(args).apply((a: any) => getSlowLog(a, opts))
}

/**
 * A collection of arguments for invoking getSlowLog.
 */
export interface GetSlowLogOutputArgs {
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}