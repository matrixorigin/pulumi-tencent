// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf
{
    public static class GetDomains
    {
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("tencentcloud:Waf/getDomains:getDomains", args ?? new GetDomainsArgs(), options.WithDefaults());

        public static Output<GetDomainsResult> Invoke(GetDomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainsResult>("tencentcloud:Waf/getDomains:getDomains", args ?? new GetDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsArgs : global::Pulumi.InvokeArgs
    {
        [Input("domain")]
        public string? Domain { get; set; }

        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDomainsArgs()
        {
        }
        public static new GetDomainsArgs Empty => new GetDomainsArgs();
    }

    public sealed class GetDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDomainsInvokeArgs()
        {
        }
        public static new GetDomainsInvokeArgs Empty => new GetDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        public readonly string? Domain;
        public readonly ImmutableArray<Outputs.GetDomainsDomainResult> WafDomains;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDomainsResult(
            string? domain,

            ImmutableArray<Outputs.GetDomainsDomainResult> domains,

            string id,

            string? instanceId,

            string? resultOutputFile)
        {
            Domain = domain;
            WafDomains = domains;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
        }
    }
}