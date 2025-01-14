// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos
{
    [TencentcloudResourceType("tencentcloud:Antiddos/ccBlackWhiteIp:CcBlackWhiteIp")]
    public partial class CcBlackWhiteIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Black white ip.
        /// </summary>
        [Output("blackWhiteIp")]
        public Output<Outputs.CcBlackWhiteIpBlackWhiteIp> BlackWhiteIp { get; private set; } = null!;

        /// <summary>
        /// domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// ip address.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// protocol.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// IP type, value [black(blacklist IP), white(whitelist IP)].
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CcBlackWhiteIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CcBlackWhiteIp(string name, CcBlackWhiteIpArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ccBlackWhiteIp:CcBlackWhiteIp", name, args ?? new CcBlackWhiteIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CcBlackWhiteIp(string name, Input<string> id, CcBlackWhiteIpState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ccBlackWhiteIp:CcBlackWhiteIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CcBlackWhiteIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CcBlackWhiteIp Get(string name, Input<string> id, CcBlackWhiteIpState? state = null, CustomResourceOptions? options = null)
        {
            return new CcBlackWhiteIp(name, id, state, options);
        }
    }

    public sealed class CcBlackWhiteIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Black white ip.
        /// </summary>
        [Input("blackWhiteIp", required: true)]
        public Input<Inputs.CcBlackWhiteIpBlackWhiteIpArgs> BlackWhiteIp { get; set; } = null!;

        /// <summary>
        /// domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// ip address.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// protocol.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// IP type, value [black(blacklist IP), white(whitelist IP)].
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CcBlackWhiteIpArgs()
        {
        }
        public static new CcBlackWhiteIpArgs Empty => new CcBlackWhiteIpArgs();
    }

    public sealed class CcBlackWhiteIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Black white ip.
        /// </summary>
        [Input("blackWhiteIp")]
        public Input<Inputs.CcBlackWhiteIpBlackWhiteIpGetArgs>? BlackWhiteIp { get; set; }

        /// <summary>
        /// domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// ip address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// IP type, value [black(blacklist IP), white(whitelist IP)].
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CcBlackWhiteIpState()
        {
        }
        public static new CcBlackWhiteIpState Empty => new CcBlackWhiteIpState();
    }
}
