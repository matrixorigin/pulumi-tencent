// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos.Inputs
{

    public sealed class DefaultAlarmThresholdDefaultAlarmConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("alarmThreshold")]
        public Input<int>? AlarmThreshold { get; set; }

        [Input("alarmType")]
        public Input<int>? AlarmType { get; set; }

        public DefaultAlarmThresholdDefaultAlarmConfigGetArgs()
        {
        }
        public static new DefaultAlarmThresholdDefaultAlarmConfigGetArgs Empty => new DefaultAlarmThresholdDefaultAlarmConfigGetArgs();
    }
}