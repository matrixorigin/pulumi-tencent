// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    public static class GetBlackHeader
    {
        public static Task<GetBlackHeaderResult> InvokeAsync(GetBlackHeaderArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBlackHeaderResult>("tencentcloud:Gaap/getBlackHeader:getBlackHeader", args ?? new GetBlackHeaderArgs(), options.WithDefaults());

        public static Output<GetBlackHeaderResult> Invoke(GetBlackHeaderInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBlackHeaderResult>("tencentcloud:Gaap/getBlackHeader:getBlackHeader", args ?? new GetBlackHeaderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBlackHeaderArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBlackHeaderArgs()
        {
        }
        public static new GetBlackHeaderArgs Empty => new GetBlackHeaderArgs();
    }

    public sealed class GetBlackHeaderInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBlackHeaderInvokeArgs()
        {
        }
        public static new GetBlackHeaderInvokeArgs Empty => new GetBlackHeaderInvokeArgs();
    }


    [OutputType]
    public sealed class GetBlackHeaderResult
    {
        public readonly ImmutableArray<string> BlackHeaders;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBlackHeaderResult(
            ImmutableArray<string> blackHeaders,

            string id,

            string? resultOutputFile)
        {
            BlackHeaders = blackHeaders;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}