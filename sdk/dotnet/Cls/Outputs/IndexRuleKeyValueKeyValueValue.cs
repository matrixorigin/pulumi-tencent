// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class IndexRuleKeyValueKeyValueValue
    {
        public readonly bool? ContainZH;
        public readonly bool? SqlFlag;
        public readonly string? Tokenizer;
        public readonly string Type;

        [OutputConstructor]
        private IndexRuleKeyValueKeyValueValue(
            bool? containZH,

            bool? sqlFlag,

            string? tokenizer,

            string type)
        {
            ContainZH = containZH;
            SqlFlag = sqlFlag;
            Tokenizer = tokenizer;
            Type = type;
        }
    }
}