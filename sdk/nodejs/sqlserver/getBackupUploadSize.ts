// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getBackupUploadSize(args: GetBackupUploadSizeArgs, opts?: pulumi.InvokeOptions): Promise<GetBackupUploadSizeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Sqlserver/getBackupUploadSize:getBackupUploadSize", {
        "backupMigrationId": args.backupMigrationId,
        "incrementalMigrationId": args.incrementalMigrationId,
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackupUploadSize.
 */
export interface GetBackupUploadSizeArgs {
    backupMigrationId: string;
    incrementalMigrationId?: string;
    instanceId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getBackupUploadSize.
 */
export interface GetBackupUploadSizeResult {
    readonly backupMigrationId: string;
    readonly cosUploadBackupFileSets: outputs.Sqlserver.GetBackupUploadSizeCosUploadBackupFileSet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly incrementalMigrationId?: string;
    readonly instanceId: string;
    readonly resultOutputFile?: string;
}
export function getBackupUploadSizeOutput(args: GetBackupUploadSizeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackupUploadSizeResult> {
    return pulumi.output(args).apply((a: any) => getBackupUploadSize(a, opts))
}

/**
 * A collection of arguments for invoking getBackupUploadSize.
 */
export interface GetBackupUploadSizeOutputArgs {
    backupMigrationId: pulumi.Input<string>;
    incrementalMigrationId?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}