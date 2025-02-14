// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elastic
{
    public static class GetPublicIpv6s
    {
        public static Task<GetPublicIpv6sResult> InvokeAsync(GetPublicIpv6sArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublicIpv6sResult>("tencentcloud:Elastic/getPublicIpv6s:getPublicIpv6s", args ?? new GetPublicIpv6sArgs(), options.WithDefaults());

        public static Output<GetPublicIpv6sResult> Invoke(GetPublicIpv6sInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublicIpv6sResult>("tencentcloud:Elastic/getPublicIpv6s:getPublicIpv6s", args ?? new GetPublicIpv6sInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublicIpv6sArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetPublicIpv6sFilterArgs>? _filters;
        public List<Inputs.GetPublicIpv6sFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetPublicIpv6sFilterArgs>());
            set => _filters = value;
        }

        [Input("ipv6AddressIds")]
        private List<string>? _ipv6AddressIds;
        public List<string> Ipv6AddressIds
        {
            get => _ipv6AddressIds ?? (_ipv6AddressIds = new List<string>());
            set => _ipv6AddressIds = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("traditional")]
        public bool? Traditional { get; set; }

        public GetPublicIpv6sArgs()
        {
        }
        public static new GetPublicIpv6sArgs Empty => new GetPublicIpv6sArgs();
    }

    public sealed class GetPublicIpv6sInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetPublicIpv6sFilterInputArgs>? _filters;
        public InputList<Inputs.GetPublicIpv6sFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetPublicIpv6sFilterInputArgs>());
            set => _filters = value;
        }

        [Input("ipv6AddressIds")]
        private InputList<string>? _ipv6AddressIds;
        public InputList<string> Ipv6AddressIds
        {
            get => _ipv6AddressIds ?? (_ipv6AddressIds = new InputList<string>());
            set => _ipv6AddressIds = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("traditional")]
        public Input<bool>? Traditional { get; set; }

        public GetPublicIpv6sInvokeArgs()
        {
        }
        public static new GetPublicIpv6sInvokeArgs Empty => new GetPublicIpv6sInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublicIpv6sResult
    {
        public readonly ImmutableArray<Outputs.GetPublicIpv6sAddressSetResult> AddressSets;
        public readonly ImmutableArray<Outputs.GetPublicIpv6sFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ipv6AddressIds;
        public readonly string? ResultOutputFile;
        public readonly bool? Traditional;

        [OutputConstructor]
        private GetPublicIpv6sResult(
            ImmutableArray<Outputs.GetPublicIpv6sAddressSetResult> addressSets,

            ImmutableArray<Outputs.GetPublicIpv6sFilterResult> filters,

            string id,

            ImmutableArray<string> ipv6AddressIds,

            string? resultOutputFile,

            bool? traditional)
        {
            AddressSets = addressSets;
            Filters = filters;
            Id = id;
            Ipv6AddressIds = ipv6AddressIds;
            ResultOutputFile = resultOutputFile;
            Traditional = traditional;
        }
    }
}
