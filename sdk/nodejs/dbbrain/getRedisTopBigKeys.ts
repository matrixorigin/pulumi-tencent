// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getRedisTopBigKeys(args: GetRedisTopBigKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetRedisTopBigKeysResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dbbrain/getRedisTopBigKeys:getRedisTopBigKeys", {
        "date": args.date,
        "instanceId": args.instanceId,
        "keyType": args.keyType,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
        "sortBy": args.sortBy,
    }, opts);
}

/**
 * A collection of arguments for invoking getRedisTopBigKeys.
 */
export interface GetRedisTopBigKeysArgs {
    date: string;
    instanceId: string;
    keyType?: string;
    product: string;
    resultOutputFile?: string;
    sortBy?: string;
}

/**
 * A collection of values returned by getRedisTopBigKeys.
 */
export interface GetRedisTopBigKeysResult {
    readonly date: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly keyType?: string;
    readonly product: string;
    readonly resultOutputFile?: string;
    readonly sortBy?: string;
    readonly topKeys: outputs.Dbbrain.GetRedisTopBigKeysTopKey[];
}
export function getRedisTopBigKeysOutput(args: GetRedisTopBigKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRedisTopBigKeysResult> {
    return pulumi.output(args).apply((a: any) => getRedisTopBigKeys(a, opts))
}

/**
 * A collection of arguments for invoking getRedisTopBigKeys.
 */
export interface GetRedisTopBigKeysOutputArgs {
    date: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    keyType?: pulumi.Input<string>;
    product: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    sortBy?: pulumi.Input<string>;
}