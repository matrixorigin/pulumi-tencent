// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTargetResourceKafkaParamTableMappingArgs : global::Pulumi.ResourceArgs
    {
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public DatahubTaskTargetResourceKafkaParamTableMappingArgs()
        {
        }
        public static new DatahubTaskTargetResourceKafkaParamTableMappingArgs Empty => new DatahubTaskTargetResourceKafkaParamTableMappingArgs();
    }
}