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
    public sealed class ScheduleTaskNotifyConfig
    {
        public readonly Outputs.ScheduleTaskNotifyConfigAwsSqs? AwsSqs;
        public readonly string? CmqModel;
        public readonly string? CmqRegion;
        public readonly string? NotifyMode;
        public readonly string? NotifyType;
        public readonly string? NotifyUrl;
        public readonly string? QueueName;
        public readonly string? TopicName;

        [OutputConstructor]
        private ScheduleTaskNotifyConfig(
            Outputs.ScheduleTaskNotifyConfigAwsSqs? awsSqs,

            string? cmqModel,

            string? cmqRegion,

            string? notifyMode,

            string? notifyType,

            string? notifyUrl,

            string? queueName,

            string? topicName)
        {
            AwsSqs = awsSqs;
            CmqModel = cmqModel;
            CmqRegion = cmqRegion;
            NotifyMode = notifyMode;
            NotifyType = notifyType;
            NotifyUrl = notifyUrl;
            QueueName = queueName;
            TopicName = topicName;
        }
    }
}
