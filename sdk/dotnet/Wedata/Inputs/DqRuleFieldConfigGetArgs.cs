// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class DqRuleFieldConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("tableConfigs")]
        private InputList<Inputs.DqRuleFieldConfigTableConfigGetArgs>? _tableConfigs;
        public InputList<Inputs.DqRuleFieldConfigTableConfigGetArgs> TableConfigs
        {
            get => _tableConfigs ?? (_tableConfigs = new InputList<Inputs.DqRuleFieldConfigTableConfigGetArgs>());
            set => _tableConfigs = value;
        }

        [Input("whereConfigs")]
        private InputList<Inputs.DqRuleFieldConfigWhereConfigGetArgs>? _whereConfigs;
        public InputList<Inputs.DqRuleFieldConfigWhereConfigGetArgs> WhereConfigs
        {
            get => _whereConfigs ?? (_whereConfigs = new InputList<Inputs.DqRuleFieldConfigWhereConfigGetArgs>());
            set => _whereConfigs = value;
        }

        public DqRuleFieldConfigGetArgs()
        {
        }
        public static new DqRuleFieldConfigGetArgs Empty => new DqRuleFieldConfigGetArgs();
    }
}