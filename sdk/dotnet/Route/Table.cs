// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Route
{
    [TencentcloudResourceType("tencentcloud:Route/table:Table")]
    public partial class Table : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time of the routing table.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Indicates whether it is the default routing table.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name of routing table.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID list of the routing entries.
        /// </summary>
        [Output("routeEntryIds")]
        public Output<ImmutableArray<string>> RouteEntryIds { get; private set; } = null!;

        /// <summary>
        /// ID list of the subnets associated with this route table.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// The tags of routing table.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// ID of VPC to which the route table should be associated.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Route/table:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, TableState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Route/table:Table", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, TableState? state = null, CustomResourceOptions? options = null)
        {
            return new Table(name, id, state, options);
        }
    }

    public sealed class TableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of routing table.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of routing table.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of VPC to which the route table should be associated.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public TableArgs()
        {
        }
        public static new TableArgs Empty => new TableArgs();
    }

    public sealed class TableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time of the routing table.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Indicates whether it is the default routing table.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of routing table.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("routeEntryIds")]
        private InputList<string>? _routeEntryIds;

        /// <summary>
        /// ID list of the routing entries.
        /// </summary>
        public InputList<string> RouteEntryIds
        {
            get => _routeEntryIds ?? (_routeEntryIds = new InputList<string>());
            set => _routeEntryIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// ID list of the subnets associated with this route table.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of routing table.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of VPC to which the route table should be associated.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public TableState()
        {
        }
        public static new TableState Empty => new TableState();
    }
}
