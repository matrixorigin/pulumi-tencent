// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstanceQpsLimit(args: GetInstanceQpsLimitArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceQpsLimitResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Waf/getInstanceQpsLimit:getInstanceQpsLimit", {
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceQpsLimit.
 */
export interface GetInstanceQpsLimitArgs {
    instanceId: string;
    resultOutputFile?: string;
    type?: string;
}

/**
 * A collection of values returned by getInstanceQpsLimit.
 */
export interface GetInstanceQpsLimitResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly qpsDatas: outputs.Waf.GetInstanceQpsLimitQpsData[];
    readonly resultOutputFile?: string;
    readonly type?: string;
}
export function getInstanceQpsLimitOutput(args: GetInstanceQpsLimitOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceQpsLimitResult> {
    return pulumi.output(args).apply((a: any) => getInstanceQpsLimit(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceQpsLimit.
 */
export interface GetInstanceQpsLimitOutputArgs {
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}