// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    [TencentcloudResourceType("tencentcloud:Css/padRuleAttachment:PadRuleAttachment")]
    public partial class PadRuleAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Push path, must same with play path, default is live.
        /// </summary>
        [Output("appName")]
        public Output<string?> AppName { get; private set; } = null!;

        /// <summary>
        /// Push domain.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Stream name.
        /// </summary>
        [Output("streamName")]
        public Output<string?> StreamName { get; private set; } = null!;

        /// <summary>
        /// Template id.
        /// </summary>
        [Output("templateId")]
        public Output<int> TemplateId { get; private set; } = null!;


        /// <summary>
        /// Create a PadRuleAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PadRuleAttachment(string name, PadRuleAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/padRuleAttachment:PadRuleAttachment", name, args ?? new PadRuleAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PadRuleAttachment(string name, Input<string> id, PadRuleAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/padRuleAttachment:PadRuleAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PadRuleAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PadRuleAttachment Get(string name, Input<string> id, PadRuleAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new PadRuleAttachment(name, id, state, options);
        }
    }

    public sealed class PadRuleAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Push path, must same with play path, default is live.
        /// </summary>
        [Input("appName")]
        public Input<string>? AppName { get; set; }

        /// <summary>
        /// Push domain.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Stream name.
        /// </summary>
        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        /// <summary>
        /// Template id.
        /// </summary>
        [Input("templateId", required: true)]
        public Input<int> TemplateId { get; set; } = null!;

        public PadRuleAttachmentArgs()
        {
        }
        public static new PadRuleAttachmentArgs Empty => new PadRuleAttachmentArgs();
    }

    public sealed class PadRuleAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Push path, must same with play path, default is live.
        /// </summary>
        [Input("appName")]
        public Input<string>? AppName { get; set; }

        /// <summary>
        /// Push domain.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Stream name.
        /// </summary>
        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        /// <summary>
        /// Template id.
        /// </summary>
        [Input("templateId")]
        public Input<int>? TemplateId { get; set; }

        public PadRuleAttachmentState()
        {
        }
        public static new PadRuleAttachmentState Empty => new PadRuleAttachmentState();
    }
}
