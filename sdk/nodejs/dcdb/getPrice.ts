// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPrice(args: GetPriceArgs, opts?: pulumi.InvokeOptions): Promise<GetPriceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dcdb/getPrice:getPrice", {
        "amountUnit": args.amountUnit,
        "instanceCount": args.instanceCount,
        "paymode": args.paymode,
        "period": args.period,
        "resultOutputFile": args.resultOutputFile,
        "shardCount": args.shardCount,
        "shardMemory": args.shardMemory,
        "shardNodeCount": args.shardNodeCount,
        "shardStorage": args.shardStorage,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrice.
 */
export interface GetPriceArgs {
    amountUnit?: string;
    instanceCount: number;
    paymode?: string;
    period: number;
    resultOutputFile?: string;
    shardCount: number;
    shardMemory: number;
    shardNodeCount: number;
    shardStorage: number;
    zone: string;
}

/**
 * A collection of values returned by getPrice.
 */
export interface GetPriceResult {
    readonly amountUnit?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceCount: number;
    readonly originalPrice: number;
    readonly paymode?: string;
    readonly period: number;
    readonly price: number;
    readonly resultOutputFile?: string;
    readonly shardCount: number;
    readonly shardMemory: number;
    readonly shardNodeCount: number;
    readonly shardStorage: number;
    readonly zone: string;
}
export function getPriceOutput(args: GetPriceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPriceResult> {
    return pulumi.output(args).apply((a: any) => getPrice(a, opts))
}

/**
 * A collection of arguments for invoking getPrice.
 */
export interface GetPriceOutputArgs {
    amountUnit?: pulumi.Input<string>;
    instanceCount: pulumi.Input<number>;
    paymode?: pulumi.Input<string>;
    period: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    shardCount: pulumi.Input<number>;
    shardMemory: pulumi.Input<number>;
    shardNodeCount: pulumi.Input<number>;
    shardStorage: pulumi.Input<number>;
    zone: pulumi.Input<string>;
}