// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInstancesModification(args?: GetInstancesModificationArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesModificationResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Cvm/getInstancesModification:getInstancesModification", {
        "filters": args.filters,
        "instanceIds": args.instanceIds,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstancesModification.
 */
export interface GetInstancesModificationArgs {
    filters?: inputs.Cvm.GetInstancesModificationFilter[];
    instanceIds?: string[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstancesModification.
 */
export interface GetInstancesModificationResult {
    readonly filters?: outputs.Cvm.GetInstancesModificationFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceIds?: string[];
    readonly instanceTypeConfigStatusLists: outputs.Cvm.GetInstancesModificationInstanceTypeConfigStatusList[];
    readonly resultOutputFile?: string;
}
export function getInstancesModificationOutput(args?: GetInstancesModificationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesModificationResult> {
    return pulumi.output(args).apply((a: any) => getInstancesModification(a, opts))
}

/**
 * A collection of arguments for invoking getInstancesModification.
 */
export interface GetInstancesModificationOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Cvm.GetInstancesModificationFilterArgs>[]>;
    instanceIds?: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
}