// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Inputs
{

    public sealed class InstanceDynamicRetentionConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bottomRetention")]
        public Input<int>? BottomRetention { get; set; }

        [Input("diskQuotaPercentage")]
        public Input<int>? DiskQuotaPercentage { get; set; }

        [Input("enable")]
        public Input<int>? Enable { get; set; }

        [Input("stepForwardPercentage")]
        public Input<int>? StepForwardPercentage { get; set; }

        public InstanceDynamicRetentionConfigGetArgs()
        {
        }
        public static new InstanceDynamicRetentionConfigGetArgs Empty => new InstanceDynamicRetentionConfigGetArgs();
    }
}