// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class DqRuleCompareRuleItemValueListArgs : global::Pulumi.ResourceArgs
    {
        [Input("value")]
        public Input<string>? Value { get; set; }

        [Input("valueType")]
        public Input<int>? ValueType { get; set; }

        public DqRuleCompareRuleItemValueListArgs()
        {
        }
        public static new DqRuleCompareRuleItemValueListArgs Empty => new DqRuleCompareRuleItemValueListArgs();
    }
}
