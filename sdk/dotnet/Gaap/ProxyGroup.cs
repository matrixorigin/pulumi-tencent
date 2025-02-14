// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    [TencentcloudResourceType("tencentcloud:Gaap/proxyGroup:ProxyGroup")]
    public partial class ProxyGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Channel group alias.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// IP version, can be taken as IPv4 or IPv6 with a default value of IPv4.
        /// </summary>
        [Output("ipAddressVersion")]
        public Output<string?> IpAddressVersion { get; private set; } = null!;

        /// <summary>
        /// Package type of channel group. Available values: Thunder and Accelerator. Default is Thunder.
        /// </summary>
        [Output("packageType")]
        public Output<string?> PackageType { get; private set; } = null!;

        /// <summary>
        /// ID of the project to which the proxy group belongs.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// real server region, refer to the interface DescribeDestRegions to return the RegionId in the parameter RegionDetail.
        /// </summary>
        [Output("realServerRegion")]
        public Output<string> RealServerRegion { get; private set; } = null!;


        /// <summary>
        /// Create a ProxyGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProxyGroup(string name, ProxyGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/proxyGroup:ProxyGroup", name, args ?? new ProxyGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProxyGroup(string name, Input<string> id, ProxyGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/proxyGroup:ProxyGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProxyGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProxyGroup Get(string name, Input<string> id, ProxyGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ProxyGroup(name, id, state, options);
        }
    }

    public sealed class ProxyGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Channel group alias.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// IP version, can be taken as IPv4 or IPv6 with a default value of IPv4.
        /// </summary>
        [Input("ipAddressVersion")]
        public Input<string>? IpAddressVersion { get; set; }

        /// <summary>
        /// Package type of channel group. Available values: Thunder and Accelerator. Default is Thunder.
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// ID of the project to which the proxy group belongs.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// real server region, refer to the interface DescribeDestRegions to return the RegionId in the parameter RegionDetail.
        /// </summary>
        [Input("realServerRegion", required: true)]
        public Input<string> RealServerRegion { get; set; } = null!;

        public ProxyGroupArgs()
        {
        }
        public static new ProxyGroupArgs Empty => new ProxyGroupArgs();
    }

    public sealed class ProxyGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Channel group alias.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// IP version, can be taken as IPv4 or IPv6 with a default value of IPv4.
        /// </summary>
        [Input("ipAddressVersion")]
        public Input<string>? IpAddressVersion { get; set; }

        /// <summary>
        /// Package type of channel group. Available values: Thunder and Accelerator. Default is Thunder.
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// ID of the project to which the proxy group belongs.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// real server region, refer to the interface DescribeDestRegions to return the RegionId in the parameter RegionDetail.
        /// </summary>
        [Input("realServerRegion")]
        public Input<string>? RealServerRegion { get; set; }

        public ProxyGroupState()
        {
        }
        public static new ProxyGroupState Empty => new ProxyGroupState();
    }
}
