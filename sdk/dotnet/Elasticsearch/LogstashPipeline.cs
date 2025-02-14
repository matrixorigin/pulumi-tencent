// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch
{
    [TencentcloudResourceType("tencentcloud:Elasticsearch/logstashPipeline:LogstashPipeline")]
    public partial class LogstashPipeline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Logstash instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Operation type. 1: save only; 2: save and deploy.
        /// </summary>
        [Output("opType")]
        public Output<int> OpType { get; private set; } = null!;

        /// <summary>
        /// Pipeline information.
        /// </summary>
        [Output("pipeline")]
        public Output<Outputs.LogstashPipelinePipeline> Pipeline { get; private set; } = null!;


        /// <summary>
        /// Create a LogstashPipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogstashPipeline(string name, LogstashPipelineArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/logstashPipeline:LogstashPipeline", name, args ?? new LogstashPipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogstashPipeline(string name, Input<string> id, LogstashPipelineState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/logstashPipeline:LogstashPipeline", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogstashPipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogstashPipeline Get(string name, Input<string> id, LogstashPipelineState? state = null, CustomResourceOptions? options = null)
        {
            return new LogstashPipeline(name, id, state, options);
        }
    }

    public sealed class LogstashPipelineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Logstash instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Operation type. 1: save only; 2: save and deploy.
        /// </summary>
        [Input("opType", required: true)]
        public Input<int> OpType { get; set; } = null!;

        /// <summary>
        /// Pipeline information.
        /// </summary>
        [Input("pipeline", required: true)]
        public Input<Inputs.LogstashPipelinePipelineArgs> Pipeline { get; set; } = null!;

        public LogstashPipelineArgs()
        {
        }
        public static new LogstashPipelineArgs Empty => new LogstashPipelineArgs();
    }

    public sealed class LogstashPipelineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Logstash instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Operation type. 1: save only; 2: save and deploy.
        /// </summary>
        [Input("opType")]
        public Input<int>? OpType { get; set; }

        /// <summary>
        /// Pipeline information.
        /// </summary>
        [Input("pipeline")]
        public Input<Inputs.LogstashPipelinePipelineGetArgs>? Pipeline { get; set; }

        public LogstashPipelineState()
        {
        }
        public static new LogstashPipelineState Empty => new LogstashPipelineState();
    }
}
