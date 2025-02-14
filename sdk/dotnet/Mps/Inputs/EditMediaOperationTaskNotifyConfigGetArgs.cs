// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class EditMediaOperationTaskNotifyConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsSqs")]
        public Input<Inputs.EditMediaOperationTaskNotifyConfigAwsSqsGetArgs>? AwsSqs { get; set; }

        [Input("cmqModel")]
        public Input<string>? CmqModel { get; set; }

        [Input("cmqRegion")]
        public Input<string>? CmqRegion { get; set; }

        [Input("notifyMode")]
        public Input<string>? NotifyMode { get; set; }

        [Input("notifyType")]
        public Input<string>? NotifyType { get; set; }

        [Input("notifyUrl")]
        public Input<string>? NotifyUrl { get; set; }

        [Input("queueName")]
        public Input<string>? QueueName { get; set; }

        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public EditMediaOperationTaskNotifyConfigGetArgs()
        {
        }
        public static new EditMediaOperationTaskNotifyConfigGetArgs Empty => new EditMediaOperationTaskNotifyConfigGetArgs();
    }
}
