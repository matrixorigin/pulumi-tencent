// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Rum
{
    public static class GetWhitelist
    {
        public static Task<GetWhitelistResult> InvokeAsync(GetWhitelistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWhitelistResult>("tencentcloud:Rum/getWhitelist:getWhitelist", args ?? new GetWhitelistArgs(), options.WithDefaults());

        public static Output<GetWhitelistResult> Invoke(GetWhitelistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWhitelistResult>("tencentcloud:Rum/getWhitelist:getWhitelist", args ?? new GetWhitelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWhitelistArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetWhitelistArgs()
        {
        }
        public static new GetWhitelistArgs Empty => new GetWhitelistArgs();
    }

    public sealed class GetWhitelistInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetWhitelistInvokeArgs()
        {
        }
        public static new GetWhitelistInvokeArgs Empty => new GetWhitelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetWhitelistResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetWhitelistWhitelistSetResult> WhitelistSets;

        [OutputConstructor]
        private GetWhitelistResult(
            string id,

            string instanceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetWhitelistWhitelistSetResult> whitelistSets)
        {
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            WhitelistSets = whitelistSets;
        }
    }
}