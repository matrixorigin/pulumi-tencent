// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eb
{
    public static class GetPlatformEventPatterns
    {
        public static Task<GetPlatformEventPatternsResult> InvokeAsync(GetPlatformEventPatternsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlatformEventPatternsResult>("tencentcloud:Eb/getPlatformEventPatterns:getPlatformEventPatterns", args ?? new GetPlatformEventPatternsArgs(), options.WithDefaults());

        public static Output<GetPlatformEventPatternsResult> Invoke(GetPlatformEventPatternsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlatformEventPatternsResult>("tencentcloud:Eb/getPlatformEventPatterns:getPlatformEventPatterns", args ?? new GetPlatformEventPatternsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlatformEventPatternsArgs : global::Pulumi.InvokeArgs
    {
        [Input("productType", required: true)]
        public string ProductType { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetPlatformEventPatternsArgs()
        {
        }
        public static new GetPlatformEventPatternsArgs Empty => new GetPlatformEventPatternsArgs();
    }

    public sealed class GetPlatformEventPatternsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("productType", required: true)]
        public Input<string> ProductType { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetPlatformEventPatternsInvokeArgs()
        {
        }
        public static new GetPlatformEventPatternsInvokeArgs Empty => new GetPlatformEventPatternsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlatformEventPatternsResult
    {
        public readonly ImmutableArray<Outputs.GetPlatformEventPatternsEventPatternResult> EventPatterns;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProductType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetPlatformEventPatternsResult(
            ImmutableArray<Outputs.GetPlatformEventPatternsEventPatternResult> eventPatterns,

            string id,

            string productType,

            string? resultOutputFile)
        {
            EventPatterns = eventPatterns;
            Id = id;
            ProductType = productType;
            ResultOutputFile = resultOutputFile;
        }
    }
}