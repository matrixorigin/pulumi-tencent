// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb
{
    [TencentcloudResourceType("tencentcloud:Clb/targetGroupAttachments:TargetGroupAttachments")]
    public partial class TargetGroupAttachments : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Association array, the combination cannot exceed 20.
        /// </summary>
        [Output("associations")]
        public Output<ImmutableArray<Outputs.TargetGroupAttachmentsAssociation>> Associations { get; private set; } = null!;

        /// <summary>
        /// CLB instance ID, (load_balancer_id and target_group_id require at least one).
        /// </summary>
        [Output("loadBalancerId")]
        public Output<string?> LoadBalancerId { get; private set; } = null!;

        /// <summary>
        /// Target group ID, (load_balancer_id and target_group_id require at least one).
        /// </summary>
        [Output("targetGroupId")]
        public Output<string?> TargetGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a TargetGroupAttachments resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetGroupAttachments(string name, TargetGroupAttachmentsArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/targetGroupAttachments:TargetGroupAttachments", name, args ?? new TargetGroupAttachmentsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetGroupAttachments(string name, Input<string> id, TargetGroupAttachmentsState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/targetGroupAttachments:TargetGroupAttachments", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TargetGroupAttachments resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetGroupAttachments Get(string name, Input<string> id, TargetGroupAttachmentsState? state = null, CustomResourceOptions? options = null)
        {
            return new TargetGroupAttachments(name, id, state, options);
        }
    }

    public sealed class TargetGroupAttachmentsArgs : global::Pulumi.ResourceArgs
    {
        [Input("associations", required: true)]
        private InputList<Inputs.TargetGroupAttachmentsAssociationArgs>? _associations;

        /// <summary>
        /// Association array, the combination cannot exceed 20.
        /// </summary>
        public InputList<Inputs.TargetGroupAttachmentsAssociationArgs> Associations
        {
            get => _associations ?? (_associations = new InputList<Inputs.TargetGroupAttachmentsAssociationArgs>());
            set => _associations = value;
        }

        /// <summary>
        /// CLB instance ID, (load_balancer_id and target_group_id require at least one).
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Target group ID, (load_balancer_id and target_group_id require at least one).
        /// </summary>
        [Input("targetGroupId")]
        public Input<string>? TargetGroupId { get; set; }

        public TargetGroupAttachmentsArgs()
        {
        }
        public static new TargetGroupAttachmentsArgs Empty => new TargetGroupAttachmentsArgs();
    }

    public sealed class TargetGroupAttachmentsState : global::Pulumi.ResourceArgs
    {
        [Input("associations")]
        private InputList<Inputs.TargetGroupAttachmentsAssociationGetArgs>? _associations;

        /// <summary>
        /// Association array, the combination cannot exceed 20.
        /// </summary>
        public InputList<Inputs.TargetGroupAttachmentsAssociationGetArgs> Associations
        {
            get => _associations ?? (_associations = new InputList<Inputs.TargetGroupAttachmentsAssociationGetArgs>());
            set => _associations = value;
        }

        /// <summary>
        /// CLB instance ID, (load_balancer_id and target_group_id require at least one).
        /// </summary>
        [Input("loadBalancerId")]
        public Input<string>? LoadBalancerId { get; set; }

        /// <summary>
        /// Target group ID, (load_balancer_id and target_group_id require at least one).
        /// </summary>
        [Input("targetGroupId")]
        public Input<string>? TargetGroupId { get; set; }

        public TargetGroupAttachmentsState()
        {
        }
        public static new TargetGroupAttachmentsState Empty => new TargetGroupAttachmentsState();
    }
}
