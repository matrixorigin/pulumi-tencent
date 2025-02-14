// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Chdfs
{
    [TencentcloudResourceType("tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule")]
    public partial class LifeCycleRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// file system id.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// life cycle rule.
        /// </summary>
        [Output("lifeCycleRule")]
        public Output<Outputs.LifeCycleRuleLifeCycleRule> ChdfsLifeCycleRule { get; private set; } = null!;


        /// <summary>
        /// Create a LifeCycleRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LifeCycleRule(string name, LifeCycleRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule", name, args ?? new LifeCycleRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LifeCycleRule(string name, Input<string> id, LifeCycleRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Chdfs/lifeCycleRule:LifeCycleRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LifeCycleRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LifeCycleRule Get(string name, Input<string> id, LifeCycleRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new LifeCycleRule(name, id, state, options);
        }
    }

    public sealed class LifeCycleRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// file system id.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// life cycle rule.
        /// </summary>
        [Input("lifeCycleRule", required: true)]
        public Input<Inputs.LifeCycleRuleLifeCycleRuleArgs> ChdfsLifeCycleRule { get; set; } = null!;

        public LifeCycleRuleArgs()
        {
        }
        public static new LifeCycleRuleArgs Empty => new LifeCycleRuleArgs();
    }

    public sealed class LifeCycleRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// file system id.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// life cycle rule.
        /// </summary>
        [Input("lifeCycleRule")]
        public Input<Inputs.LifeCycleRuleLifeCycleRuleGetArgs>? ChdfsLifeCycleRule { get; set; }

        public LifeCycleRuleState()
        {
        }
        public static new LifeCycleRuleState Empty => new LifeCycleRuleState();
    }
}
