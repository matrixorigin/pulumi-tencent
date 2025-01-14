// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdc
{
    public static class GetDedicatedClusterInstanceTypes
    {
        public static Task<GetDedicatedClusterInstanceTypesResult> InvokeAsync(GetDedicatedClusterInstanceTypesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDedicatedClusterInstanceTypesResult>("tencentcloud:Cdc/getDedicatedClusterInstanceTypes:getDedicatedClusterInstanceTypes", args ?? new GetDedicatedClusterInstanceTypesArgs(), options.WithDefaults());

        public static Output<GetDedicatedClusterInstanceTypesResult> Invoke(GetDedicatedClusterInstanceTypesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDedicatedClusterInstanceTypesResult>("tencentcloud:Cdc/getDedicatedClusterInstanceTypes:getDedicatedClusterInstanceTypes", args ?? new GetDedicatedClusterInstanceTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDedicatedClusterInstanceTypesArgs : global::Pulumi.InvokeArgs
    {
        [Input("dedicatedClusterId", required: true)]
        public string DedicatedClusterId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDedicatedClusterInstanceTypesArgs()
        {
        }
        public static new GetDedicatedClusterInstanceTypesArgs Empty => new GetDedicatedClusterInstanceTypesArgs();
    }

    public sealed class GetDedicatedClusterInstanceTypesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dedicatedClusterId", required: true)]
        public Input<string> DedicatedClusterId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDedicatedClusterInstanceTypesInvokeArgs()
        {
        }
        public static new GetDedicatedClusterInstanceTypesInvokeArgs Empty => new GetDedicatedClusterInstanceTypesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDedicatedClusterInstanceTypesResult
    {
        public readonly string DedicatedClusterId;
        public readonly ImmutableArray<Outputs.GetDedicatedClusterInstanceTypesDedicatedClusterInstanceTypeSetResult> DedicatedClusterInstanceTypeSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDedicatedClusterInstanceTypesResult(
            string dedicatedClusterId,

            ImmutableArray<Outputs.GetDedicatedClusterInstanceTypesDedicatedClusterInstanceTypeSetResult> dedicatedClusterInstanceTypeSets,

            string id,

            string? resultOutputFile)
        {
            DedicatedClusterId = dedicatedClusterId;
            DedicatedClusterInstanceTypeSets = dedicatedClusterInstanceTypeSets;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
