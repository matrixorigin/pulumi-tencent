// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSpec(args: GetSpecArgs, opts?: pulumi.InvokeOptions): Promise<GetSpecResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Clickhouse/getSpec:getSpec", {
        "isElastic": args.isElastic,
        "payMode": args.payMode,
        "resultOutputFile": args.resultOutputFile,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpec.
 */
export interface GetSpecArgs {
    isElastic?: boolean;
    payMode?: string;
    resultOutputFile?: string;
    zone: string;
}

/**
 * A collection of values returned by getSpec.
 */
export interface GetSpecResult {
    readonly attachCbsSpecs: outputs.Clickhouse.GetSpecAttachCbsSpec[];
    readonly commonSpecs: outputs.Clickhouse.GetSpecCommonSpec[];
    readonly dataSpecs: outputs.Clickhouse.GetSpecDataSpec[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isElastic?: boolean;
    readonly payMode?: string;
    readonly resultOutputFile?: string;
    readonly zone: string;
}
export function getSpecOutput(args: GetSpecOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpecResult> {
    return pulumi.output(args).apply((a: any) => getSpec(a, opts))
}

/**
 * A collection of arguments for invoking getSpec.
 */
export interface GetSpecOutputArgs {
    isElastic?: pulumi.Input<boolean>;
    payMode?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}