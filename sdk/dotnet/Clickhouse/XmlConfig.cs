// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clickhouse
{
    [TencentcloudResourceType("tencentcloud:Clickhouse/xmlConfig:XmlConfig")]
    public partial class XmlConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Configuration file modification information.
        /// </summary>
        [Output("modifyConfContext")]
        public Output<Outputs.XmlConfigModifyConfContext> ModifyConfContext { get; private set; } = null!;


        /// <summary>
        /// Create a XmlConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public XmlConfig(string name, XmlConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/xmlConfig:XmlConfig", name, args ?? new XmlConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private XmlConfig(string name, Input<string> id, XmlConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/xmlConfig:XmlConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing XmlConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static XmlConfig Get(string name, Input<string> id, XmlConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new XmlConfig(name, id, state, options);
        }
    }

    public sealed class XmlConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Configuration file modification information.
        /// </summary>
        [Input("modifyConfContext", required: true)]
        public Input<Inputs.XmlConfigModifyConfContextArgs> ModifyConfContext { get; set; } = null!;

        public XmlConfigArgs()
        {
        }
        public static new XmlConfigArgs Empty => new XmlConfigArgs();
    }

    public sealed class XmlConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Configuration file modification information.
        /// </summary>
        [Input("modifyConfContext")]
        public Input<Inputs.XmlConfigModifyConfContextGetArgs>? ModifyConfContext { get; set; }

        public XmlConfigState()
        {
        }
        public static new XmlConfigState Empty => new XmlConfigState();
    }
}