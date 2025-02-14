// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos
{
    [TencentcloudResourceType("tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig")]
    public partial class DdosSpeedLimitConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        /// </summary>
        [Output("ddosSpeedLimitConfig")]
        public Output<Outputs.DdosSpeedLimitConfigDdosSpeedLimitConfig> AntiddosDdosSpeedLimitConfig { get; private set; } = null!;

        /// <summary>
        /// InstanceId.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a DdosSpeedLimitConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosSpeedLimitConfig(string name, DdosSpeedLimitConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig", name, args ?? new DdosSpeedLimitConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosSpeedLimitConfig(string name, Input<string> id, DdosSpeedLimitConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ddosSpeedLimitConfig:DdosSpeedLimitConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DdosSpeedLimitConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosSpeedLimitConfig Get(string name, Input<string> id, DdosSpeedLimitConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DdosSpeedLimitConfig(name, id, state, options);
        }
    }

    public sealed class DdosSpeedLimitConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        /// </summary>
        [Input("ddosSpeedLimitConfig", required: true)]
        public Input<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigArgs> AntiddosDdosSpeedLimitConfig { get; set; } = null!;

        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public DdosSpeedLimitConfigArgs()
        {
        }
        public static new DdosSpeedLimitConfigArgs Empty => new DdosSpeedLimitConfigArgs();
    }

    public sealed class DdosSpeedLimitConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accessing speed limit configuration, the configuration ID cannot be empty when filling in parameters.
        /// </summary>
        [Input("ddosSpeedLimitConfig")]
        public Input<Inputs.DdosSpeedLimitConfigDdosSpeedLimitConfigGetArgs>? AntiddosDdosSpeedLimitConfig { get; set; }

        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public DdosSpeedLimitConfigState()
        {
        }
        public static new DdosSpeedLimitConfigState Empty => new DdosSpeedLimitConfigState();
    }
}
