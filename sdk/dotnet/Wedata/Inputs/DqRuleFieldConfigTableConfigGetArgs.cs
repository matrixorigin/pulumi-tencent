// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class DqRuleFieldConfigTableConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("fieldConfigs")]
        private InputList<Inputs.DqRuleFieldConfigTableConfigFieldConfigGetArgs>? _fieldConfigs;
        public InputList<Inputs.DqRuleFieldConfigTableConfigFieldConfigGetArgs> FieldConfigs
        {
            get => _fieldConfigs ?? (_fieldConfigs = new InputList<Inputs.DqRuleFieldConfigTableConfigFieldConfigGetArgs>());
            set => _fieldConfigs = value;
        }

        [Input("tableId")]
        public Input<string>? TableId { get; set; }

        [Input("tableKey")]
        public Input<string>? TableKey { get; set; }

        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public DqRuleFieldConfigTableConfigGetArgs()
        {
        }
        public static new DqRuleFieldConfigTableConfigGetArgs Empty => new DqRuleFieldConfigTableConfigGetArgs();
    }
}