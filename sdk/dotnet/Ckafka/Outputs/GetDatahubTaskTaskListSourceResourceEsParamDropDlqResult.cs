// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetDatahubTaskTaskListSourceResourceEsParamDropDlqResult
    {
        public readonly string DlqType;
        public readonly ImmutableArray<Outputs.GetDatahubTaskTaskListSourceResourceEsParamDropDlqKafkaParamResult> KafkaParams;
        public readonly int MaxRetryAttempts;
        public readonly int RetryInterval;
        public readonly ImmutableArray<Outputs.GetDatahubTaskTaskListSourceResourceEsParamDropDlqTopicParamResult> TopicParams;
        public readonly string Type;

        [OutputConstructor]
        private GetDatahubTaskTaskListSourceResourceEsParamDropDlqResult(
            string dlqType,

            ImmutableArray<Outputs.GetDatahubTaskTaskListSourceResourceEsParamDropDlqKafkaParamResult> kafkaParams,

            int maxRetryAttempts,

            int retryInterval,

            ImmutableArray<Outputs.GetDatahubTaskTaskListSourceResourceEsParamDropDlqTopicParamResult> topicParams,

            string type)
        {
            DlqType = dlqType;
            KafkaParams = kafkaParams;
            MaxRetryAttempts = maxRetryAttempts;
            RetryInterval = retryInterval;
            TopicParams = topicParams;
            Type = type;
        }
    }
}