// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDataSourceList(args?: GetDataSourceListArgs, opts?: pulumi.InvokeOptions): Promise<GetDataSourceListResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Wedata/getDataSourceList:getDataSourceList", {
        "filters": args.filters,
        "orderFields": args.orderFields,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDataSourceList.
 */
export interface GetDataSourceListArgs {
    filters?: inputs.Wedata.GetDataSourceListFilter[];
    orderFields?: inputs.Wedata.GetDataSourceListOrderField[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDataSourceList.
 */
export interface GetDataSourceListResult {
    readonly filters?: outputs.Wedata.GetDataSourceListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly orderFields?: outputs.Wedata.GetDataSourceListOrderField[];
    readonly resultOutputFile?: string;
    readonly rows: outputs.Wedata.GetDataSourceListRow[];
}
export function getDataSourceListOutput(args?: GetDataSourceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataSourceListResult> {
    return pulumi.output(args).apply((a: any) => getDataSourceList(a, opts))
}

/**
 * A collection of arguments for invoking getDataSourceList.
 */
export interface GetDataSourceListOutputArgs {
    filters?: pulumi.Input<pulumi.Input<inputs.Wedata.GetDataSourceListFilterArgs>[]>;
    orderFields?: pulumi.Input<pulumi.Input<inputs.Wedata.GetDataSourceListOrderFieldArgs>[]>;
    resultOutputFile?: pulumi.Input<string>;
}