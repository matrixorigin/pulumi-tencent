// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPeakPoints(args: GetPeakPointsArgs, opts?: pulumi.InvokeOptions): Promise<GetPeakPointsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Waf/getPeakPoints:getPeakPoints", {
        "domain": args.domain,
        "edition": args.edition,
        "fromTime": args.fromTime,
        "instanceId": args.instanceId,
        "metricName": args.metricName,
        "resultOutputFile": args.resultOutputFile,
        "toTime": args.toTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getPeakPoints.
 */
export interface GetPeakPointsArgs {
    domain?: string;
    edition?: string;
    fromTime: string;
    instanceId?: string;
    metricName?: string;
    resultOutputFile?: string;
    toTime: string;
}

/**
 * A collection of values returned by getPeakPoints.
 */
export interface GetPeakPointsResult {
    readonly domain?: string;
    readonly edition?: string;
    readonly fromTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly metricName?: string;
    readonly points: outputs.Waf.GetPeakPointsPoint[];
    readonly resultOutputFile?: string;
    readonly toTime: string;
}
export function getPeakPointsOutput(args: GetPeakPointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPeakPointsResult> {
    return pulumi.output(args).apply((a: any) => getPeakPoints(a, opts))
}

/**
 * A collection of arguments for invoking getPeakPoints.
 */
export interface GetPeakPointsOutputArgs {
    domain?: pulumi.Input<string>;
    edition?: pulumi.Input<string>;
    fromTime: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    metricName?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    toTime: pulumi.Input<string>;
}