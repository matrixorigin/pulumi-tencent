// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Images/getInstance:getInstance", {
        "imageId": args.imageId,
        "imageNameRegex": args.imageNameRegex,
        "imageTypes": args.imageTypes,
        "instanceType": args.instanceType,
        "osName": args.osName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    imageId?: string;
    imageNameRegex?: string;
    imageTypes?: string[];
    instanceType?: string;
    osName?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageId?: string;
    readonly imageNameRegex?: string;
    readonly imageTypes?: string[];
    readonly images: outputs.Images.GetInstanceImage[];
    readonly instanceType?: string;
    readonly osName?: string;
    readonly resultOutputFile?: string;
}
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    imageId?: pulumi.Input<string>;
    imageNameRegex?: pulumi.Input<string>;
    imageTypes?: pulumi.Input<pulumi.Input<string>[]>;
    instanceType?: pulumi.Input<string>;
    osName?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}