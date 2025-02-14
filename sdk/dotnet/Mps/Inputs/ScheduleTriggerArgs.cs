// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleTriggerArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsS3FileUploadTrigger")]
        public Input<Inputs.ScheduleTriggerAwsS3FileUploadTriggerArgs>? AwsS3FileUploadTrigger { get; set; }

        [Input("cosFileUploadTrigger")]
        public Input<Inputs.ScheduleTriggerCosFileUploadTriggerArgs>? CosFileUploadTrigger { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ScheduleTriggerArgs()
        {
        }
        public static new ScheduleTriggerArgs Empty => new ScheduleTriggerArgs();
    }
}
