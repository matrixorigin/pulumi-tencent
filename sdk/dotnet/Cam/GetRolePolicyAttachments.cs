// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cam
{
    public static class GetRolePolicyAttachments
    {
        public static Task<GetRolePolicyAttachmentsResult> InvokeAsync(GetRolePolicyAttachmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRolePolicyAttachmentsResult>("tencentcloud:Cam/getRolePolicyAttachments:getRolePolicyAttachments", args ?? new GetRolePolicyAttachmentsArgs(), options.WithDefaults());

        public static Output<GetRolePolicyAttachmentsResult> Invoke(GetRolePolicyAttachmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRolePolicyAttachmentsResult>("tencentcloud:Cam/getRolePolicyAttachments:getRolePolicyAttachments", args ?? new GetRolePolicyAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRolePolicyAttachmentsArgs : global::Pulumi.InvokeArgs
    {
        [Input("createMode")]
        public int? CreateMode { get; set; }

        [Input("policyId")]
        public string? PolicyId { get; set; }

        [Input("policyType")]
        public string? PolicyType { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("roleId", required: true)]
        public string RoleId { get; set; } = null!;

        public GetRolePolicyAttachmentsArgs()
        {
        }
        public static new GetRolePolicyAttachmentsArgs Empty => new GetRolePolicyAttachmentsArgs();
    }

    public sealed class GetRolePolicyAttachmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("createMode")]
        public Input<int>? CreateMode { get; set; }

        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public GetRolePolicyAttachmentsInvokeArgs()
        {
        }
        public static new GetRolePolicyAttachmentsInvokeArgs Empty => new GetRolePolicyAttachmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRolePolicyAttachmentsResult
    {
        public readonly int? CreateMode;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? PolicyId;
        public readonly string? PolicyType;
        public readonly string? ResultOutputFile;
        public readonly string RoleId;
        public readonly ImmutableArray<Outputs.GetRolePolicyAttachmentsRolePolicyAttachmentListResult> RolePolicyAttachmentLists;

        [OutputConstructor]
        private GetRolePolicyAttachmentsResult(
            int? createMode,

            string id,

            string? policyId,

            string? policyType,

            string? resultOutputFile,

            string roleId,

            ImmutableArray<Outputs.GetRolePolicyAttachmentsRolePolicyAttachmentListResult> rolePolicyAttachmentLists)
        {
            CreateMode = createMode;
            Id = id;
            PolicyId = policyId;
            PolicyType = policyType;
            ResultOutputFile = resultOutputFile;
            RoleId = roleId;
            RolePolicyAttachmentLists = rolePolicyAttachmentLists;
        }
    }
}