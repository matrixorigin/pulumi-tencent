// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getNetDetectStateCheck(args: GetNetDetectStateCheckArgs, opts?: pulumi.InvokeOptions): Promise<GetNetDetectStateCheckResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpc/getNetDetectStateCheck:getNetDetectStateCheck", {
        "detectDestinationIps": args.detectDestinationIps,
        "netDetectId": args.netDetectId,
        "netDetectName": args.netDetectName,
        "nextHopDestination": args.nextHopDestination,
        "nextHopType": args.nextHopType,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetDetectStateCheck.
 */
export interface GetNetDetectStateCheckArgs {
    detectDestinationIps: string[];
    netDetectId?: string;
    netDetectName?: string;
    nextHopDestination: string;
    nextHopType: string;
    resultOutputFile?: string;
    subnetId?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by getNetDetectStateCheck.
 */
export interface GetNetDetectStateCheckResult {
    readonly detectDestinationIps: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly netDetectId?: string;
    readonly netDetectIpStateSets: outputs.Vpc.GetNetDetectStateCheckNetDetectIpStateSet[];
    readonly netDetectName?: string;
    readonly nextHopDestination: string;
    readonly nextHopType: string;
    readonly resultOutputFile?: string;
    readonly subnetId?: string;
    readonly vpcId?: string;
}
export function getNetDetectStateCheckOutput(args: GetNetDetectStateCheckOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetDetectStateCheckResult> {
    return pulumi.output(args).apply((a: any) => getNetDetectStateCheck(a, opts))
}

/**
 * A collection of arguments for invoking getNetDetectStateCheck.
 */
export interface GetNetDetectStateCheckOutputArgs {
    detectDestinationIps: pulumi.Input<pulumi.Input<string>[]>;
    netDetectId?: pulumi.Input<string>;
    netDetectName?: pulumi.Input<string>;
    nextHopDestination: pulumi.Input<string>;
    nextHopType: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}