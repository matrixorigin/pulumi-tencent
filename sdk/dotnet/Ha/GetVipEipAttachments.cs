// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ha
{
    public static class GetVipEipAttachments
    {
        public static Task<GetVipEipAttachmentsResult> InvokeAsync(GetVipEipAttachmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVipEipAttachmentsResult>("tencentcloud:Ha/getVipEipAttachments:getVipEipAttachments", args ?? new GetVipEipAttachmentsArgs(), options.WithDefaults());

        public static Output<GetVipEipAttachmentsResult> Invoke(GetVipEipAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVipEipAttachmentsResult>("tencentcloud:Ha/getVipEipAttachments:getVipEipAttachments", args ?? new GetVipEipAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVipEipAttachmentsArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressIp")]
        public string? AddressIp { get; set; }

        [Input("havipId", required: true)]
        public string HavipId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetVipEipAttachmentsArgs()
        {
        }
        public static new GetVipEipAttachmentsArgs Empty => new GetVipEipAttachmentsArgs();
    }

    public sealed class GetVipEipAttachmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressIp")]
        public Input<string>? AddressIp { get; set; }

        [Input("havipId", required: true)]
        public Input<string> HavipId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetVipEipAttachmentsInvokeArgs()
        {
        }
        public static new GetVipEipAttachmentsInvokeArgs Empty => new GetVipEipAttachmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetVipEipAttachmentsResult
    {
        public readonly string? AddressIp;
        public readonly ImmutableArray<Outputs.GetVipEipAttachmentsHaVipEipAttachmentListResult> HaVipEipAttachmentLists;
        public readonly string HavipId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetVipEipAttachmentsResult(
            string? addressIp,

            ImmutableArray<Outputs.GetVipEipAttachmentsHaVipEipAttachmentListResult> haVipEipAttachmentLists,

            string havipId,

            string id,

            string? resultOutputFile)
        {
            AddressIp = addressIp;
            HaVipEipAttachmentLists = haVipEipAttachmentLists;
            HavipId = havipId;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}