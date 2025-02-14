// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Bi
{
    [TencentcloudResourceType("tencentcloud:Bi/embedTokenApply:EmbedTokenApply")]
    public partial class EmbedTokenApply : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create the generated token.
        /// </summary>
        [Output("biToken")]
        public Output<string> BiToken { get; private set; } = null!;

        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createAt")]
        public Output<string> CreateAt { get; private set; } = null!;

        /// <summary>
        /// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        /// </summary>
        [Output("expireTime")]
        public Output<string?> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Sharing page id, this is empty value 0 when embedding the board.
        /// </summary>
        [Output("pageId")]
        public Output<int?> PageId { get; private set; } = null!;

        /// <summary>
        /// Share project id.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Page means embedding the page, and panel means embedding the entire board.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        /// </summary>
        [Output("ticketNum")]
        public Output<int?> TicketNum { get; private set; } = null!;

        /// <summary>
        /// Upadte time.
        /// </summary>
        [Output("udpateAt")]
        public Output<string> UdpateAt { get; private set; } = null!;

        /// <summary>
        /// User enterprise ID (for multi-user only).
        /// </summary>
        [Output("userCorpId")]
        public Output<string?> UserCorpId { get; private set; } = null!;

        /// <summary>
        /// UserId (for multi-user only).
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a EmbedTokenApply resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EmbedTokenApply(string name, EmbedTokenApplyArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Bi/embedTokenApply:EmbedTokenApply", name, args ?? new EmbedTokenApplyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EmbedTokenApply(string name, Input<string> id, EmbedTokenApplyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Bi/embedTokenApply:EmbedTokenApply", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EmbedTokenApply resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EmbedTokenApply Get(string name, Input<string> id, EmbedTokenApplyState? state = null, CustomResourceOptions? options = null)
        {
            return new EmbedTokenApply(name, id, state, options);
        }
    }

    public sealed class EmbedTokenApplyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Sharing page id, this is empty value 0 when embedding the board.
        /// </summary>
        [Input("pageId")]
        public Input<int>? PageId { get; set; }

        /// <summary>
        /// Share project id.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Page means embedding the page, and panel means embedding the entire board.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        /// </summary>
        [Input("ticketNum")]
        public Input<int>? TicketNum { get; set; }

        /// <summary>
        /// User enterprise ID (for multi-user only).
        /// </summary>
        [Input("userCorpId")]
        public Input<string>? UserCorpId { get; set; }

        /// <summary>
        /// UserId (for multi-user only).
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public EmbedTokenApplyArgs()
        {
        }
        public static new EmbedTokenApplyArgs Empty => new EmbedTokenApplyArgs();
    }

    public sealed class EmbedTokenApplyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create the generated token.
        /// </summary>
        [Input("biToken")]
        public Input<string>? BiToken { get; set; }

        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createAt")]
        public Input<string>? CreateAt { get; set; }

        /// <summary>
        /// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Sharing page id, this is empty value 0 when embedding the board.
        /// </summary>
        [Input("pageId")]
        public Input<int>? PageId { get; set; }

        /// <summary>
        /// Share project id.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Page means embedding the page, and panel means embedding the entire board.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
        /// </summary>
        [Input("ticketNum")]
        public Input<int>? TicketNum { get; set; }

        /// <summary>
        /// Upadte time.
        /// </summary>
        [Input("udpateAt")]
        public Input<string>? UdpateAt { get; set; }

        /// <summary>
        /// User enterprise ID (for multi-user only).
        /// </summary>
        [Input("userCorpId")]
        public Input<string>? UserCorpId { get; set; }

        /// <summary>
        /// UserId (for multi-user only).
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public EmbedTokenApplyState()
        {
        }
        public static new EmbedTokenApplyState Empty => new EmbedTokenApplyState();
    }
}
