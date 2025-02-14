// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdwdoris
{
    [TencentcloudResourceType("tencentcloud:Cdwdoris/workloadGroup:WorkloadGroup")]
    public partial class WorkloadGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Resource group configuration.
        /// </summary>
        [Output("workloadGroup")]
        public Output<Outputs.WorkloadGroupWorkloadGroup?> CdwdorisWorkloadGroup { get; private set; } = null!;


        /// <summary>
        /// Create a WorkloadGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadGroup(string name, WorkloadGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdwdoris/workloadGroup:WorkloadGroup", name, args ?? new WorkloadGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadGroup(string name, Input<string> id, WorkloadGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdwdoris/workloadGroup:WorkloadGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkloadGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadGroup Get(string name, Input<string> id, WorkloadGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkloadGroup(name, id, state, options);
        }
    }

    public sealed class WorkloadGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Resource group configuration.
        /// </summary>
        [Input("workloadGroup")]
        public Input<Inputs.WorkloadGroupWorkloadGroupArgs>? CdwdorisWorkloadGroup { get; set; }

        public WorkloadGroupArgs()
        {
        }
        public static new WorkloadGroupArgs Empty => new WorkloadGroupArgs();
    }

    public sealed class WorkloadGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Resource group configuration.
        /// </summary>
        [Input("workloadGroup")]
        public Input<Inputs.WorkloadGroupWorkloadGroupGetArgs>? CdwdorisWorkloadGroup { get; set; }

        public WorkloadGroupState()
        {
        }
        public static new WorkloadGroupState Empty => new WorkloadGroupState();
    }
}
