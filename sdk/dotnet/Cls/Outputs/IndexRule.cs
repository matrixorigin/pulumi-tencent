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
    public sealed class IndexRule
    {
        public readonly Outputs.IndexRuleDynamicIndex? DynamicIndex;
        public readonly Outputs.IndexRuleFullText? FullText;
        public readonly Outputs.IndexRuleKeyValue? KeyValue;
        public readonly Outputs.IndexRuleTag? Tag;

        [OutputConstructor]
        private IndexRule(
            Outputs.IndexRuleDynamicIndex? dynamicIndex,

            Outputs.IndexRuleFullText? fullText,

            Outputs.IndexRuleKeyValue? keyValue,

            Outputs.IndexRuleTag? tag)
        {
            DynamicIndex = dynamicIndex;
            FullText = fullText;
            KeyValue = keyValue;
            Tag = tag;
        }
    }
}