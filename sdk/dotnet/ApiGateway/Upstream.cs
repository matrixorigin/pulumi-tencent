// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway
{
    [TencentcloudResourceType("tencentcloud:ApiGateway/upstream:Upstream")]
    public partial class Upstream : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Load balancing algorithm, value range: ROUND-ROBIN.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Health check configuration, currently only supports VPC channels.
        /// </summary>
        [Output("healthChecker")]
        public Output<Outputs.UpstreamHealthChecker?> HealthChecker { get; private set; } = null!;

        /// <summary>
        /// Configuration of K8S container service.
        /// </summary>
        [Output("k8sServices")]
        public Output<ImmutableArray<Outputs.UpstreamK8sService>> K8sServices { get; private set; } = null!;

        /// <summary>
        /// Backend nodes.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<Outputs.UpstreamNode>> Nodes { get; private set; } = null!;

        /// <summary>
        /// Request retry count, default to 3 times.
        /// </summary>
        [Output("retries")]
        public Output<int?> Retries { get; private set; } = null!;

        /// <summary>
        /// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
        /// </summary>
        [Output("scheme")]
        public Output<string> Scheme { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// VPC Unique ID.
        /// </summary>
        [Output("uniqVpcId")]
        public Output<string> UniqVpcId { get; private set; } = null!;

        /// <summary>
        /// Backend channel description.
        /// </summary>
        [Output("upstreamDescription")]
        public Output<string?> UpstreamDescription { get; private set; } = null!;

        /// <summary>
        /// Host request header forwarded by gateway to backend.
        /// </summary>
        [Output("upstreamHost")]
        public Output<string?> UpstreamHost { get; private set; } = null!;

        /// <summary>
        /// Backend channel name.
        /// </summary>
        [Output("upstreamName")]
        public Output<string?> UpstreamName { get; private set; } = null!;

        /// <summary>
        /// Backend access type, value range: IP_PORT, K8S.
        /// </summary>
        [Output("upstreamType")]
        public Output<string?> UpstreamType { get; private set; } = null!;


        /// <summary>
        /// Create a Upstream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Upstream(string name, UpstreamArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/upstream:Upstream", name, args ?? new UpstreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Upstream(string name, Input<string> id, UpstreamState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/upstream:Upstream", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Upstream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Upstream Get(string name, Input<string> id, UpstreamState? state = null, CustomResourceOptions? options = null)
        {
            return new Upstream(name, id, state, options);
        }
    }

    public sealed class UpstreamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load balancing algorithm, value range: ROUND-ROBIN.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// Health check configuration, currently only supports VPC channels.
        /// </summary>
        [Input("healthChecker")]
        public Input<Inputs.UpstreamHealthCheckerArgs>? HealthChecker { get; set; }

        [Input("k8sServices")]
        private InputList<Inputs.UpstreamK8sServiceArgs>? _k8sServices;

        /// <summary>
        /// Configuration of K8S container service.
        /// </summary>
        public InputList<Inputs.UpstreamK8sServiceArgs> K8sServices
        {
            get => _k8sServices ?? (_k8sServices = new InputList<Inputs.UpstreamK8sServiceArgs>());
            set => _k8sServices = value;
        }

        [Input("nodes")]
        private InputList<Inputs.UpstreamNodeArgs>? _nodes;

        /// <summary>
        /// Backend nodes.
        /// </summary>
        public InputList<Inputs.UpstreamNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.UpstreamNodeArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// Request retry count, default to 3 times.
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        /// <summary>
        /// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
        /// </summary>
        [Input("scheme", required: true)]
        public Input<string> Scheme { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// VPC Unique ID.
        /// </summary>
        [Input("uniqVpcId", required: true)]
        public Input<string> UniqVpcId { get; set; } = null!;

        /// <summary>
        /// Backend channel description.
        /// </summary>
        [Input("upstreamDescription")]
        public Input<string>? UpstreamDescription { get; set; }

        /// <summary>
        /// Host request header forwarded by gateway to backend.
        /// </summary>
        [Input("upstreamHost")]
        public Input<string>? UpstreamHost { get; set; }

        /// <summary>
        /// Backend channel name.
        /// </summary>
        [Input("upstreamName")]
        public Input<string>? UpstreamName { get; set; }

        /// <summary>
        /// Backend access type, value range: IP_PORT, K8S.
        /// </summary>
        [Input("upstreamType")]
        public Input<string>? UpstreamType { get; set; }

        public UpstreamArgs()
        {
        }
        public static new UpstreamArgs Empty => new UpstreamArgs();
    }

    public sealed class UpstreamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load balancing algorithm, value range: ROUND-ROBIN.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Health check configuration, currently only supports VPC channels.
        /// </summary>
        [Input("healthChecker")]
        public Input<Inputs.UpstreamHealthCheckerGetArgs>? HealthChecker { get; set; }

        [Input("k8sServices")]
        private InputList<Inputs.UpstreamK8sServiceGetArgs>? _k8sServices;

        /// <summary>
        /// Configuration of K8S container service.
        /// </summary>
        public InputList<Inputs.UpstreamK8sServiceGetArgs> K8sServices
        {
            get => _k8sServices ?? (_k8sServices = new InputList<Inputs.UpstreamK8sServiceGetArgs>());
            set => _k8sServices = value;
        }

        [Input("nodes")]
        private InputList<Inputs.UpstreamNodeGetArgs>? _nodes;

        /// <summary>
        /// Backend nodes.
        /// </summary>
        public InputList<Inputs.UpstreamNodeGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.UpstreamNodeGetArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// Request retry count, default to 3 times.
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        /// <summary>
        /// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// VPC Unique ID.
        /// </summary>
        [Input("uniqVpcId")]
        public Input<string>? UniqVpcId { get; set; }

        /// <summary>
        /// Backend channel description.
        /// </summary>
        [Input("upstreamDescription")]
        public Input<string>? UpstreamDescription { get; set; }

        /// <summary>
        /// Host request header forwarded by gateway to backend.
        /// </summary>
        [Input("upstreamHost")]
        public Input<string>? UpstreamHost { get; set; }

        /// <summary>
        /// Backend channel name.
        /// </summary>
        [Input("upstreamName")]
        public Input<string>? UpstreamName { get; set; }

        /// <summary>
        /// Backend access type, value range: IP_PORT, K8S.
        /// </summary>
        [Input("upstreamType")]
        public Input<string>? UpstreamType { get; set; }

        public UpstreamState()
        {
        }
        public static new UpstreamState Empty => new UpstreamState();
    }
}