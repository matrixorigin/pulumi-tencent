// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dts
{
    [TencentcloudResourceType("tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation")]
    public partial class SyncJobStartOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;


        /// <summary>
        /// Create a SyncJobStartOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncJobStartOperation(string name, SyncJobStartOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation", name, args ?? new SyncJobStartOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncJobStartOperation(string name, Input<string> id, SyncJobStartOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyncJobStartOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncJobStartOperation Get(string name, Input<string> id, SyncJobStartOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncJobStartOperation(name, id, state, options);
        }
    }

    public sealed class SyncJobStartOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        public SyncJobStartOperationArgs()
        {
        }
        public static new SyncJobStartOperationArgs Empty => new SyncJobStartOperationArgs();
    }

    public sealed class SyncJobStartOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        public SyncJobStartOperationState()
        {
        }
        public static new SyncJobStartOperationState Empty => new SyncJobStartOperationState();
    }
}