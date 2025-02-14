// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus
{
    [TencentcloudResourceType("tencentcloud:Oceanus/stopJob:StopJob")]
    public partial class StopJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description information for batch job stop.
        /// </summary>
        [Output("stopJobDescriptions")]
        public Output<ImmutableArray<Outputs.StopJobStopJobDescription>> StopJobDescriptions { get; private set; } = null!;

        /// <summary>
        /// Workspace SerialId.
        /// </summary>
        [Output("workSpaceId")]
        public Output<string?> WorkSpaceId { get; private set; } = null!;


        /// <summary>
        /// Create a StopJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StopJob(string name, StopJobArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Oceanus/stopJob:StopJob", name, args ?? new StopJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StopJob(string name, Input<string> id, StopJobState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Oceanus/stopJob:StopJob", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StopJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StopJob Get(string name, Input<string> id, StopJobState? state = null, CustomResourceOptions? options = null)
        {
            return new StopJob(name, id, state, options);
        }
    }

    public sealed class StopJobArgs : global::Pulumi.ResourceArgs
    {
        [Input("stopJobDescriptions", required: true)]
        private InputList<Inputs.StopJobStopJobDescriptionArgs>? _stopJobDescriptions;

        /// <summary>
        /// The description information for batch job stop.
        /// </summary>
        public InputList<Inputs.StopJobStopJobDescriptionArgs> StopJobDescriptions
        {
            get => _stopJobDescriptions ?? (_stopJobDescriptions = new InputList<Inputs.StopJobStopJobDescriptionArgs>());
            set => _stopJobDescriptions = value;
        }

        /// <summary>
        /// Workspace SerialId.
        /// </summary>
        [Input("workSpaceId")]
        public Input<string>? WorkSpaceId { get; set; }

        public StopJobArgs()
        {
        }
        public static new StopJobArgs Empty => new StopJobArgs();
    }

    public sealed class StopJobState : global::Pulumi.ResourceArgs
    {
        [Input("stopJobDescriptions")]
        private InputList<Inputs.StopJobStopJobDescriptionGetArgs>? _stopJobDescriptions;

        /// <summary>
        /// The description information for batch job stop.
        /// </summary>
        public InputList<Inputs.StopJobStopJobDescriptionGetArgs> StopJobDescriptions
        {
            get => _stopJobDescriptions ?? (_stopJobDescriptions = new InputList<Inputs.StopJobStopJobDescriptionGetArgs>());
            set => _stopJobDescriptions = value;
        }

        /// <summary>
        /// Workspace SerialId.
        /// </summary>
        [Input("workSpaceId")]
        public Input<string>? WorkSpaceId { get; set; }

        public StopJobState()
        {
        }
        public static new StopJobState Empty => new StopJobState();
    }
}
