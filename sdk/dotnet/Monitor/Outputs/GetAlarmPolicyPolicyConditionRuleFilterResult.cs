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
    public sealed class GetAlarmPolicyPolicyConditionRuleFilterResult
    {
        public readonly string Dimensions;
        public readonly string Type;

        [OutputConstructor]
        private GetAlarmPolicyPolicyConditionRuleFilterResult(
            string dimensions,

            string type)
        {
            Dimensions = dimensions;
            Type = type;
        }
    }
}
