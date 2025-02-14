// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos
{
    [TencentcloudResourceType("tencentcloud:Antiddos/ipAlarmThresholdConfig:IpAlarmThresholdConfig")]
    public partial class IpAlarmThresholdConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alarm threshold, in Mbps, with a value of&amp;gt;=0; When used as an input parameter, setting 0 will delete the alarm
        /// threshold configuration;.
        /// </summary>
        [Output("alarmThreshold")]
        public Output<int> AlarmThreshold { get; private set; } = null!;

        /// <summary>
        /// Alarm threshold type, value [1 (incoming traffic alarm threshold) 2 (attack cleaning traffic alarm threshold)].
        /// </summary>
        [Output("alarmType")]
        public Output<int> AlarmType { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance ip.
        /// </summary>
        [Output("instanceIp")]
        public Output<string> InstanceIp { get; private set; } = null!;


        /// <summary>
        /// Create a IpAlarmThresholdConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpAlarmThresholdConfig(string name, IpAlarmThresholdConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ipAlarmThresholdConfig:IpAlarmThresholdConfig", name, args ?? new IpAlarmThresholdConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpAlarmThresholdConfig(string name, Input<string> id, IpAlarmThresholdConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ipAlarmThresholdConfig:IpAlarmThresholdConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpAlarmThresholdConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpAlarmThresholdConfig Get(string name, Input<string> id, IpAlarmThresholdConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new IpAlarmThresholdConfig(name, id, state, options);
        }
    }

    public sealed class IpAlarmThresholdConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alarm threshold, in Mbps, with a value of&amp;gt;=0; When used as an input parameter, setting 0 will delete the alarm
        /// threshold configuration;.
        /// </summary>
        [Input("alarmThreshold", required: true)]
        public Input<int> AlarmThreshold { get; set; } = null!;

        /// <summary>
        /// Alarm threshold type, value [1 (incoming traffic alarm threshold) 2 (attack cleaning traffic alarm threshold)].
        /// </summary>
        [Input("alarmType", required: true)]
        public Input<int> AlarmType { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Instance ip.
        /// </summary>
        [Input("instanceIp", required: true)]
        public Input<string> InstanceIp { get; set; } = null!;

        public IpAlarmThresholdConfigArgs()
        {
        }
        public static new IpAlarmThresholdConfigArgs Empty => new IpAlarmThresholdConfigArgs();
    }

    public sealed class IpAlarmThresholdConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alarm threshold, in Mbps, with a value of&amp;gt;=0; When used as an input parameter, setting 0 will delete the alarm
        /// threshold configuration;.
        /// </summary>
        [Input("alarmThreshold")]
        public Input<int>? AlarmThreshold { get; set; }

        /// <summary>
        /// Alarm threshold type, value [1 (incoming traffic alarm threshold) 2 (attack cleaning traffic alarm threshold)].
        /// </summary>
        [Input("alarmType")]
        public Input<int>? AlarmType { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance ip.
        /// </summary>
        [Input("instanceIp")]
        public Input<string>? InstanceIp { get; set; }

        public IpAlarmThresholdConfigState()
        {
        }
        public static new IpAlarmThresholdConfigState Empty => new IpAlarmThresholdConfigState();
    }
}
