// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getProject(args?: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Bi/getProject:getProject", {
        "allPage": args.allPage,
        "keyword": args.keyword,
        "moduleCollection": args.moduleCollection,
        "pageNo": args.pageNo,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    allPage?: boolean;
    keyword?: string;
    moduleCollection?: string;
    pageNo?: number;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    readonly allPage?: boolean;
    readonly extra: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyword?: string;
    readonly lists: outputs.Bi.GetProjectList[];
    readonly moduleCollection?: string;
    readonly msg: string;
    readonly pageNo?: number;
    readonly resultOutputFile?: string;
}
export function getProjectOutput(args?: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply((a: any) => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    allPage?: pulumi.Input<boolean>;
    keyword?: pulumi.Input<string>;
    moduleCollection?: pulumi.Input<string>;
    pageNo?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
}