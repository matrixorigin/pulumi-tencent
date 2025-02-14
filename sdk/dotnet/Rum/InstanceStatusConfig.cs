// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Rum
{
    [TencentcloudResourceType("tencentcloud:Rum/instanceStatusConfig:InstanceStatusConfig")]
    public partial class InstanceStatusConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance status (`1`=creating, `2`=running, `3`=abnormal, `4`=restarting, `5`=stopping, `6`=stopped, `7`=deleted).
        /// </summary>
        [Output("instanceStatus")]
        public Output<int> InstanceStatus { get; private set; } = null!;

        /// <summary>
        /// `resume`, `stop`.
        /// </summary>
        [Output("operate")]
        public Output<string> Operate { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceStatusConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceStatusConfig(string name, InstanceStatusConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Rum/instanceStatusConfig:InstanceStatusConfig", name, args ?? new InstanceStatusConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceStatusConfig(string name, Input<string> id, InstanceStatusConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Rum/instanceStatusConfig:InstanceStatusConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceStatusConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceStatusConfig Get(string name, Input<string> id, InstanceStatusConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceStatusConfig(name, id, state, options);
        }
    }

    public sealed class InstanceStatusConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// `resume`, `stop`.
        /// </summary>
        [Input("operate", required: true)]
        public Input<string> Operate { get; set; } = null!;

        public InstanceStatusConfigArgs()
        {
        }
        public static new InstanceStatusConfigArgs Empty => new InstanceStatusConfigArgs();
    }

    public sealed class InstanceStatusConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance status (`1`=creating, `2`=running, `3`=abnormal, `4`=restarting, `5`=stopping, `6`=stopped, `7`=deleted).
        /// </summary>
        [Input("instanceStatus")]
        public Input<int>? InstanceStatus { get; set; }

        /// <summary>
        /// `resume`, `stop`.
        /// </summary>
        [Input("operate")]
        public Input<string>? Operate { get; set; }

        public InstanceStatusConfigState()
        {
        }
        public static new InstanceStatusConfigState Empty => new InstanceStatusConfigState();
    }
}
