// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tse
{
    public static class GetAccessAddress
    {
        public static Task<GetAccessAddressResult> InvokeAsync(GetAccessAddressArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessAddressResult>("tencentcloud:Tse/getAccessAddress:getAccessAddress", args ?? new GetAccessAddressArgs(), options.WithDefaults());

        public static Output<GetAccessAddressResult> Invoke(GetAccessAddressInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessAddressResult>("tencentcloud:Tse/getAccessAddress:getAccessAddress", args ?? new GetAccessAddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessAddressArgs : global::Pulumi.InvokeArgs
    {
        [Input("engineRegion")]
        public string? EngineRegion { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        [Input("workload")]
        public string? Workload { get; set; }

        public GetAccessAddressArgs()
        {
        }
        public static new GetAccessAddressArgs Empty => new GetAccessAddressArgs();
    }

    public sealed class GetAccessAddressInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("engineRegion")]
        public Input<string>? EngineRegion { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        [Input("workload")]
        public Input<string>? Workload { get; set; }

        public GetAccessAddressInvokeArgs()
        {
        }
        public static new GetAccessAddressInvokeArgs Empty => new GetAccessAddressInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessAddressResult
    {
        public readonly string ConsoleInternetAddress;
        public readonly int ConsoleInternetBandWidth;
        public readonly string ConsoleIntranetAddress;
        public readonly string? EngineRegion;
        public readonly ImmutableArray<Outputs.GetAccessAddressEnvAddressInfoResult> EnvAddressInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string InternetAddress;
        public readonly int InternetBandWidth;
        public readonly string IntranetAddress;
        public readonly ImmutableArray<Outputs.GetAccessAddressLimiterAddressInfoResult> LimiterAddressInfos;
        public readonly string? ResultOutputFile;
        public readonly string? SubnetId;
        public readonly string? VpcId;
        public readonly string? Workload;

        [OutputConstructor]
        private GetAccessAddressResult(
            string consoleInternetAddress,

            int consoleInternetBandWidth,

            string consoleIntranetAddress,

            string? engineRegion,

            ImmutableArray<Outputs.GetAccessAddressEnvAddressInfoResult> envAddressInfos,

            string id,

            string instanceId,

            string internetAddress,

            int internetBandWidth,

            string intranetAddress,

            ImmutableArray<Outputs.GetAccessAddressLimiterAddressInfoResult> limiterAddressInfos,

            string? resultOutputFile,

            string? subnetId,

            string? vpcId,

            string? workload)
        {
            ConsoleInternetAddress = consoleInternetAddress;
            ConsoleInternetBandWidth = consoleInternetBandWidth;
            ConsoleIntranetAddress = consoleIntranetAddress;
            EngineRegion = engineRegion;
            EnvAddressInfos = envAddressInfos;
            Id = id;
            InstanceId = instanceId;
            InternetAddress = internetAddress;
            InternetBandWidth = internetBandWidth;
            IntranetAddress = intranetAddress;
            LimiterAddressInfos = limiterAddressInfos;
            ResultOutputFile = resultOutputFile;
            SubnetId = subnetId;
            VpcId = vpcId;
            Workload = workload;
        }
    }
}