// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmPolicyPolicyConditionsTempConditionResult
    {
        public readonly string ComplexExpression;
        public readonly int IsUnionRule;
        public readonly ImmutableArray<Outputs.GetAlarmPolicyPolicyConditionsTempConditionRuleResult> Rules;

        [OutputConstructor]
        private GetAlarmPolicyPolicyConditionsTempConditionResult(
            string complexExpression,

            int isUnionRule,

            ImmutableArray<Outputs.GetAlarmPolicyPolicyConditionsTempConditionRuleResult> rules)
        {
            ComplexExpression = complexExpression;
            IsUnionRule = isUnionRule;
            Rules = rules;
        }
    }
}