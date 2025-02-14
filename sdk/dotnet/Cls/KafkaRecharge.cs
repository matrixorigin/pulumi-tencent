// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls
{
    [TencentcloudResourceType("tencentcloud:Cls/kafkaRecharge:KafkaRecharge")]
    public partial class KafkaRecharge : global::Pulumi.CustomResource
    {
        /// <summary>
        /// user consumer group name.
        /// </summary>
        [Output("consumerGroupName")]
        public Output<string?> ConsumerGroupName { get; private set; } = null!;

        /// <summary>
        /// ServerAddr is encryption addr.
        /// </summary>
        [Output("isEncryptionAddr")]
        public Output<bool> IsEncryptionAddr { get; private set; } = null!;

        /// <summary>
        /// CKafka Instance id.
        /// </summary>
        [Output("kafkaInstance")]
        public Output<string?> KafkaInstance { get; private set; } = null!;

        /// <summary>
        /// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        /// </summary>
        [Output("kafkaType")]
        public Output<int> KafkaType { get; private set; } = null!;

        /// <summary>
        /// log recharge rule.
        /// </summary>
        [Output("logRechargeRule")]
        public Output<Outputs.KafkaRechargeLogRechargeRule> LogRechargeRule { get; private set; } = null!;

        /// <summary>
        /// kafka recharge name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The translation is: -2: Earliest (default) -1: Latest.
        /// </summary>
        [Output("offset")]
        public Output<int> Offset { get; private set; } = null!;

        /// <summary>
        /// encryption protocol.
        /// </summary>
        [Output("protocol")]
        public Output<Outputs.KafkaRechargeProtocol> Protocol { get; private set; } = null!;

        /// <summary>
        /// Server addr.
        /// </summary>
        [Output("serverAddr")]
        public Output<string?> ServerAddr { get; private set; } = null!;

        /// <summary>
        /// recharge for cls TopicId.
        /// </summary>
        [Output("topicId")]
        public Output<string> TopicId { get; private set; } = null!;

        /// <summary>
        /// user need recharge kafka topic list.
        /// </summary>
        [Output("userKafkaTopics")]
        public Output<string> UserKafkaTopics { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaRecharge resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaRecharge(string name, KafkaRechargeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/kafkaRecharge:KafkaRecharge", name, args ?? new KafkaRechargeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaRecharge(string name, Input<string> id, KafkaRechargeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/kafkaRecharge:KafkaRecharge", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KafkaRecharge resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaRecharge Get(string name, Input<string> id, KafkaRechargeState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaRecharge(name, id, state, options);
        }
    }

    public sealed class KafkaRechargeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// user consumer group name.
        /// </summary>
        [Input("consumerGroupName")]
        public Input<string>? ConsumerGroupName { get; set; }

        /// <summary>
        /// ServerAddr is encryption addr.
        /// </summary>
        [Input("isEncryptionAddr")]
        public Input<bool>? IsEncryptionAddr { get; set; }

        /// <summary>
        /// CKafka Instance id.
        /// </summary>
        [Input("kafkaInstance")]
        public Input<string>? KafkaInstance { get; set; }

        /// <summary>
        /// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        /// </summary>
        [Input("kafkaType", required: true)]
        public Input<int> KafkaType { get; set; } = null!;

        /// <summary>
        /// log recharge rule.
        /// </summary>
        [Input("logRechargeRule")]
        public Input<Inputs.KafkaRechargeLogRechargeRuleArgs>? LogRechargeRule { get; set; }

        /// <summary>
        /// kafka recharge name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The translation is: -2: Earliest (default) -1: Latest.
        /// </summary>
        [Input("offset", required: true)]
        public Input<int> Offset { get; set; } = null!;

        /// <summary>
        /// encryption protocol.
        /// </summary>
        [Input("protocol")]
        public Input<Inputs.KafkaRechargeProtocolArgs>? Protocol { get; set; }

        /// <summary>
        /// Server addr.
        /// </summary>
        [Input("serverAddr")]
        public Input<string>? ServerAddr { get; set; }

        /// <summary>
        /// recharge for cls TopicId.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        /// <summary>
        /// user need recharge kafka topic list.
        /// </summary>
        [Input("userKafkaTopics", required: true)]
        public Input<string> UserKafkaTopics { get; set; } = null!;

        public KafkaRechargeArgs()
        {
        }
        public static new KafkaRechargeArgs Empty => new KafkaRechargeArgs();
    }

    public sealed class KafkaRechargeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// user consumer group name.
        /// </summary>
        [Input("consumerGroupName")]
        public Input<string>? ConsumerGroupName { get; set; }

        /// <summary>
        /// ServerAddr is encryption addr.
        /// </summary>
        [Input("isEncryptionAddr")]
        public Input<bool>? IsEncryptionAddr { get; set; }

        /// <summary>
        /// CKafka Instance id.
        /// </summary>
        [Input("kafkaInstance")]
        public Input<string>? KafkaInstance { get; set; }

        /// <summary>
        /// kafka recharge type, 0 for CKafka, 1 fro user define Kafka.
        /// </summary>
        [Input("kafkaType")]
        public Input<int>? KafkaType { get; set; }

        /// <summary>
        /// log recharge rule.
        /// </summary>
        [Input("logRechargeRule")]
        public Input<Inputs.KafkaRechargeLogRechargeRuleGetArgs>? LogRechargeRule { get; set; }

        /// <summary>
        /// kafka recharge name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The translation is: -2: Earliest (default) -1: Latest.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        /// <summary>
        /// encryption protocol.
        /// </summary>
        [Input("protocol")]
        public Input<Inputs.KafkaRechargeProtocolGetArgs>? Protocol { get; set; }

        /// <summary>
        /// Server addr.
        /// </summary>
        [Input("serverAddr")]
        public Input<string>? ServerAddr { get; set; }

        /// <summary>
        /// recharge for cls TopicId.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        /// <summary>
        /// user need recharge kafka topic list.
        /// </summary>
        [Input("userKafkaTopics")]
        public Input<string>? UserKafkaTopics { get; set; }

        public KafkaRechargeState()
        {
        }
        public static new KafkaRechargeState Empty => new KafkaRechargeState();
    }
}
