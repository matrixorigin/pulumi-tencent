// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class ScheduleTrigger
    {
        public readonly Outputs.ScheduleTriggerAwsS3FileUploadTrigger? AwsS3FileUploadTrigger;
        public readonly Outputs.ScheduleTriggerCosFileUploadTrigger? CosFileUploadTrigger;
        public readonly string Type;

        [OutputConstructor]
        private ScheduleTrigger(
            Outputs.ScheduleTriggerAwsS3FileUploadTrigger? awsS3FileUploadTrigger,

            Outputs.ScheduleTriggerCosFileUploadTrigger? cosFileUploadTrigger,

            string type)
        {
            AwsS3FileUploadTrigger = awsS3FileUploadTrigger;
            CosFileUploadTrigger = cosFileUploadTrigger;
            Type = type;
        }
    }
}