// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    [TencentcloudResourceType("tencentcloud:Monitor/tmpTkeGlobalNotification:TmpTkeGlobalNotification")]
    public partial class TmpTkeGlobalNotification : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Output("notification")]
        public Output<Outputs.TmpTkeGlobalNotificationNotification> Notification { get; private set; } = null!;


        /// <summary>
        /// Create a TmpTkeGlobalNotification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpTkeGlobalNotification(string name, TmpTkeGlobalNotificationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeGlobalNotification:TmpTkeGlobalNotification", name, args ?? new TmpTkeGlobalNotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpTkeGlobalNotification(string name, Input<string> id, TmpTkeGlobalNotificationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeGlobalNotification:TmpTkeGlobalNotification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpTkeGlobalNotification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpTkeGlobalNotification Get(string name, Input<string> id, TmpTkeGlobalNotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpTkeGlobalNotification(name, id, state, options);
        }
    }

    public sealed class TmpTkeGlobalNotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Input("notification", required: true)]
        public Input<Inputs.TmpTkeGlobalNotificationNotificationArgs> Notification { get; set; } = null!;

        public TmpTkeGlobalNotificationArgs()
        {
        }
        public static new TmpTkeGlobalNotificationArgs Empty => new TmpTkeGlobalNotificationArgs();
    }

    public sealed class TmpTkeGlobalNotificationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Input("notification")]
        public Input<Inputs.TmpTkeGlobalNotificationNotificationGetArgs>? Notification { get; set; }

        public TmpTkeGlobalNotificationState()
        {
        }
        public static new TmpTkeGlobalNotificationState Empty => new TmpTkeGlobalNotificationState();
    }
}
