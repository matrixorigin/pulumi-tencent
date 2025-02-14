// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class IntegrationRealtimeTaskTaskInfoNodeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("configs")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeConfigGetArgs>? _configs;
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeConfigGetArgs>());
            set => _configs = value;
        }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("creatorUin")]
        public Input<string>? CreatorUin { get; set; }

        [Input("dataSourceType")]
        public Input<string>? DataSourceType { get; set; }

        [Input("datasourceId")]
        public Input<string>? DatasourceId { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("extConfigs")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeExtConfigGetArgs>? _extConfigs;
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeExtConfigGetArgs> ExtConfigs
        {
            get => _extConfigs ?? (_extConfigs = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeExtConfigGetArgs>());
            set => _extConfigs = value;
        }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeMapping")]
        public Input<Inputs.IntegrationRealtimeTaskTaskInfoNodeNodeMappingGetArgs>? NodeMapping { get; set; }

        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("operatorUin")]
        public Input<string>? OperatorUin { get; set; }

        [Input("ownerUin")]
        public Input<string>? OwnerUin { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("schemas")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeSchemaGetArgs>? _schemas;
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeSchemaGetArgs> Schemas
        {
            get => _schemas ?? (_schemas = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoNodeSchemaGetArgs>());
            set => _schemas = value;
        }

        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public IntegrationRealtimeTaskTaskInfoNodeGetArgs()
        {
        }
        public static new IntegrationRealtimeTaskTaskInfoNodeGetArgs Empty => new IntegrationRealtimeTaskTaskInfoNodeGetArgs();
    }
}
