// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tsf
{
    [TencentcloudResourceType("tencentcloud:Tsf/lane:Lane")]
    public partial class Lane : global::Pulumi.CustomResource
    {
        /// <summary>
        /// creation time.
        /// </summary>
        [Output("createTime")]
        public Output<int> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Whether to enter the application.
        /// </summary>
        [Output("entrance")]
        public Output<bool> Entrance { get; private set; } = null!;

        /// <summary>
        /// Swimlane Deployment Group Information.
        /// </summary>
        [Output("laneGroupLists")]
        public Output<ImmutableArray<Outputs.LaneLaneGroupList>> LaneGroupLists { get; private set; } = null!;

        /// <summary>
        /// Lane id.
        /// </summary>
        [Output("laneId")]
        public Output<string> LaneId { get; private set; } = null!;

        /// <summary>
        /// Lane name.
        /// </summary>
        [Output("laneName")]
        public Output<string> LaneName { get; private set; } = null!;

        /// <summary>
        /// A list of namespaces to which the swimlane has associated deployment groups.
        /// </summary>
        [Output("namespaceIdLists")]
        public Output<ImmutableArray<string>> NamespaceIdLists { get; private set; } = null!;

        /// <summary>
        /// Program id list.
        /// </summary>
        [Output("programIdLists")]
        public Output<ImmutableArray<string>> ProgramIdLists { get; private set; } = null!;

        /// <summary>
        /// Lane Remarks.
        /// </summary>
        [Output("remark")]
        public Output<string> Remark { get; private set; } = null!;

        /// <summary>
        /// update time.
        /// </summary>
        [Output("updateTime")]
        public Output<int> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Lane resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lane(string name, LaneArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/lane:Lane", name, args ?? new LaneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lane(string name, Input<string> id, LaneState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/lane:Lane", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Lane resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lane Get(string name, Input<string> id, LaneState? state = null, CustomResourceOptions? options = null)
        {
            return new Lane(name, id, state, options);
        }
    }

    public sealed class LaneArgs : global::Pulumi.ResourceArgs
    {
        [Input("laneGroupLists", required: true)]
        private InputList<Inputs.LaneLaneGroupListArgs>? _laneGroupLists;

        /// <summary>
        /// Swimlane Deployment Group Information.
        /// </summary>
        public InputList<Inputs.LaneLaneGroupListArgs> LaneGroupLists
        {
            get => _laneGroupLists ?? (_laneGroupLists = new InputList<Inputs.LaneLaneGroupListArgs>());
            set => _laneGroupLists = value;
        }

        /// <summary>
        /// Lane name.
        /// </summary>
        [Input("laneName", required: true)]
        public Input<string> LaneName { get; set; } = null!;

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// Program id list.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        /// <summary>
        /// Lane Remarks.
        /// </summary>
        [Input("remark", required: true)]
        public Input<string> Remark { get; set; } = null!;

        public LaneArgs()
        {
        }
        public static new LaneArgs Empty => new LaneArgs();
    }

    public sealed class LaneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// creation time.
        /// </summary>
        [Input("createTime")]
        public Input<int>? CreateTime { get; set; }

        /// <summary>
        /// Whether to enter the application.
        /// </summary>
        [Input("entrance")]
        public Input<bool>? Entrance { get; set; }

        [Input("laneGroupLists")]
        private InputList<Inputs.LaneLaneGroupListGetArgs>? _laneGroupLists;

        /// <summary>
        /// Swimlane Deployment Group Information.
        /// </summary>
        public InputList<Inputs.LaneLaneGroupListGetArgs> LaneGroupLists
        {
            get => _laneGroupLists ?? (_laneGroupLists = new InputList<Inputs.LaneLaneGroupListGetArgs>());
            set => _laneGroupLists = value;
        }

        /// <summary>
        /// Lane id.
        /// </summary>
        [Input("laneId")]
        public Input<string>? LaneId { get; set; }

        /// <summary>
        /// Lane name.
        /// </summary>
        [Input("laneName")]
        public Input<string>? LaneName { get; set; }

        [Input("namespaceIdLists")]
        private InputList<string>? _namespaceIdLists;

        /// <summary>
        /// A list of namespaces to which the swimlane has associated deployment groups.
        /// </summary>
        public InputList<string> NamespaceIdLists
        {
            get => _namespaceIdLists ?? (_namespaceIdLists = new InputList<string>());
            set => _namespaceIdLists = value;
        }

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// Program id list.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        /// <summary>
        /// Lane Remarks.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// update time.
        /// </summary>
        [Input("updateTime")]
        public Input<int>? UpdateTime { get; set; }

        public LaneState()
        {
        }
        public static new LaneState Empty => new LaneState();
    }
}
