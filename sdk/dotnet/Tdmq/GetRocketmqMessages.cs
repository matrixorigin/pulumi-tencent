// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tdmq
{
    public static class GetRocketmqMessages
    {
        public static Task<GetRocketmqMessagesResult> InvokeAsync(GetRocketmqMessagesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRocketmqMessagesResult>("tencentcloud:Tdmq/getRocketmqMessages:getRocketmqMessages", args ?? new GetRocketmqMessagesArgs(), options.WithDefaults());

        public static Output<GetRocketmqMessagesResult> Invoke(GetRocketmqMessagesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRocketmqMessagesResult>("tencentcloud:Tdmq/getRocketmqMessages:getRocketmqMessages", args ?? new GetRocketmqMessagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRocketmqMessagesArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("msgId", required: true)]
        public string MsgId { get; set; } = null!;

        [Input("queryDlqMsg")]
        public bool? QueryDlqMsg { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("topicName", required: true)]
        public string TopicName { get; set; } = null!;

        public GetRocketmqMessagesArgs()
        {
        }
        public static new GetRocketmqMessagesArgs Empty => new GetRocketmqMessagesArgs();
    }

    public sealed class GetRocketmqMessagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("msgId", required: true)]
        public Input<string> MsgId { get; set; } = null!;

        [Input("queryDlqMsg")]
        public Input<bool>? QueryDlqMsg { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public GetRocketmqMessagesInvokeArgs()
        {
        }
        public static new GetRocketmqMessagesInvokeArgs Empty => new GetRocketmqMessagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRocketmqMessagesResult
    {
        public readonly string Body;
        public readonly string ClusterId;
        public readonly string EnvironmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetRocketmqMessagesMessageTrackResult> MessageTracks;
        public readonly string MsgId;
        public readonly string ProduceTime;
        public readonly string ProducerAddr;
        public readonly string Properties;
        public readonly bool? QueryDlqMsg;
        public readonly string? ResultOutputFile;
        public readonly string ShowTopicName;
        public readonly string TopicName;

        [OutputConstructor]
        private GetRocketmqMessagesResult(
            string body,

            string clusterId,

            string environmentId,

            string id,

            ImmutableArray<Outputs.GetRocketmqMessagesMessageTrackResult> messageTracks,

            string msgId,

            string produceTime,

            string producerAddr,

            string properties,

            bool? queryDlqMsg,

            string? resultOutputFile,

            string showTopicName,

            string topicName)
        {
            Body = body;
            ClusterId = clusterId;
            EnvironmentId = environmentId;
            Id = id;
            MessageTracks = messageTracks;
            MsgId = msgId;
            ProduceTime = produceTime;
            ProducerAddr = producerAddr;
            Properties = properties;
            QueryDlqMsg = queryDlqMsg;
            ResultOutputFile = resultOutputFile;
            ShowTopicName = showTopicName;
            TopicName = topicName;
        }
    }
}