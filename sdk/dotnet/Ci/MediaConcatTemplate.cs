// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ci
{
    [TencentcloudResourceType("tencentcloud:Ci/mediaConcatTemplate:MediaConcatTemplate")]
    public partial class MediaConcatTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// stitching template.
        /// </summary>
        [Output("concatTemplate")]
        public Output<Outputs.MediaConcatTemplateConcatTemplate> ConcatTemplate { get; private set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a MediaConcatTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MediaConcatTemplate(string name, MediaConcatTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaConcatTemplate:MediaConcatTemplate", name, args ?? new MediaConcatTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MediaConcatTemplate(string name, Input<string> id, MediaConcatTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ci/mediaConcatTemplate:MediaConcatTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MediaConcatTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MediaConcatTemplate Get(string name, Input<string> id, MediaConcatTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new MediaConcatTemplate(name, id, state, options);
        }
    }

    public sealed class MediaConcatTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// stitching template.
        /// </summary>
        [Input("concatTemplate", required: true)]
        public Input<Inputs.MediaConcatTemplateConcatTemplateArgs> ConcatTemplate { get; set; } = null!;

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MediaConcatTemplateArgs()
        {
        }
        public static new MediaConcatTemplateArgs Empty => new MediaConcatTemplateArgs();
    }

    public sealed class MediaConcatTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// stitching template.
        /// </summary>
        [Input("concatTemplate")]
        public Input<Inputs.MediaConcatTemplateConcatTemplateGetArgs>? ConcatTemplate { get; set; }

        /// <summary>
        /// The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MediaConcatTemplateState()
        {
        }
        public static new MediaConcatTemplateState Empty => new MediaConcatTemplateState();
    }
}
