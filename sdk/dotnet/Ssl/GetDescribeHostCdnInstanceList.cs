// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    public static class GetDescribeHostCdnInstanceList
    {
        public static Task<GetDescribeHostCdnInstanceListResult> InvokeAsync(GetDescribeHostCdnInstanceListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostCdnInstanceListResult>("tencentcloud:Ssl/getDescribeHostCdnInstanceList:getDescribeHostCdnInstanceList", args ?? new GetDescribeHostCdnInstanceListArgs(), options.WithDefaults());

        public static Output<GetDescribeHostCdnInstanceListResult> Invoke(GetDescribeHostCdnInstanceListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeHostCdnInstanceListResult>("tencentcloud:Ssl/getDescribeHostCdnInstanceList:getDescribeHostCdnInstanceList", args ?? new GetDescribeHostCdnInstanceListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostCdnInstanceListArgs : global::Pulumi.InvokeArgs
    {
        [Input("asyncCache")]
        public int? AsyncCache { get; set; }

        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetDescribeHostCdnInstanceListFilterArgs>? _filters;
        public List<Inputs.GetDescribeHostCdnInstanceListFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDescribeHostCdnInstanceListFilterArgs>());
            set => _filters = value;
        }

        [Input("isCache")]
        public int? IsCache { get; set; }

        [Input("oldCertificateId")]
        public string? OldCertificateId { get; set; }

        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeHostCdnInstanceListArgs()
        {
        }
        public static new GetDescribeHostCdnInstanceListArgs Empty => new GetDescribeHostCdnInstanceListArgs();
    }

    public sealed class GetDescribeHostCdnInstanceListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("asyncCache")]
        public Input<int>? AsyncCache { get; set; }

        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("filters")]
        private InputList<Inputs.GetDescribeHostCdnInstanceListFilterInputArgs>? _filters;
        public InputList<Inputs.GetDescribeHostCdnInstanceListFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetDescribeHostCdnInstanceListFilterInputArgs>());
            set => _filters = value;
        }

        [Input("isCache")]
        public Input<int>? IsCache { get; set; }

        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeHostCdnInstanceListInvokeArgs()
        {
        }
        public static new GetDescribeHostCdnInstanceListInvokeArgs Empty => new GetDescribeHostCdnInstanceListInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeHostCdnInstanceListResult
    {
        public readonly int? AsyncCache;
        public readonly string AsyncCacheTime;
        public readonly int AsyncOffset;
        public readonly int AsyncTotalNum;
        public readonly string CertificateId;
        public readonly ImmutableArray<Outputs.GetDescribeHostCdnInstanceListFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetDescribeHostCdnInstanceListInstanceListResult> InstanceLists;
        public readonly int? IsCache;
        public readonly string? OldCertificateId;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeHostCdnInstanceListResult(
            int? asyncCache,

            string asyncCacheTime,

            int asyncOffset,

            int asyncTotalNum,

            string certificateId,

            ImmutableArray<Outputs.GetDescribeHostCdnInstanceListFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetDescribeHostCdnInstanceListInstanceListResult> instanceLists,

            int? isCache,

            string? oldCertificateId,

            string resourceType,

            string? resultOutputFile)
        {
            AsyncCache = asyncCache;
            AsyncCacheTime = asyncCacheTime;
            AsyncOffset = asyncOffset;
            AsyncTotalNum = asyncTotalNum;
            CertificateId = certificateId;
            Filters = filters;
            Id = id;
            InstanceLists = instanceLists;
            IsCache = isCache;
            OldCertificateId = oldCertificateId;
            ResourceType = resourceType;
            ResultOutputFile = resultOutputFile;
        }
    }
}