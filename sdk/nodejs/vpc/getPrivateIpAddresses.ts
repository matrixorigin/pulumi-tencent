// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getPrivateIpAddresses(args: GetPrivateIpAddressesArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateIpAddressesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Vpc/getPrivateIpAddresses:getPrivateIpAddresses", {
        "privateIpAddresses": args.privateIpAddresses,
        "resultOutputFile": args.resultOutputFile,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateIpAddresses.
 */
export interface GetPrivateIpAddressesArgs {
    privateIpAddresses: string[];
    resultOutputFile?: string;
    vpcId: string;
}

/**
 * A collection of values returned by getPrivateIpAddresses.
 */
export interface GetPrivateIpAddressesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly privateIpAddresses: string[];
    readonly resultOutputFile?: string;
    readonly vpcId: string;
    readonly vpcPrivateIpAddressSets: outputs.Vpc.GetPrivateIpAddressesVpcPrivateIpAddressSet[];
}
export function getPrivateIpAddressesOutput(args: GetPrivateIpAddressesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateIpAddressesResult> {
    return pulumi.output(args).apply((a: any) => getPrivateIpAddresses(a, opts))
}

/**
 * A collection of arguments for invoking getPrivateIpAddresses.
 */
export interface GetPrivateIpAddressesOutputArgs {
    privateIpAddresses: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    vpcId: pulumi.Input<string>;
}