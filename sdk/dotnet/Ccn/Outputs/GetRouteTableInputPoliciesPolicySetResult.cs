// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn.Outputs
{

    [OutputType]
    public sealed class GetRouteTableInputPoliciesPolicySetResult
    {
        public readonly string? CreateTime;
        public readonly int? PolicyVersion;
        public readonly ImmutableArray<Outputs.GetRouteTableInputPoliciesPolicySetPolicyResult> Policys;

        [OutputConstructor]
        private GetRouteTableInputPoliciesPolicySetResult(
            string? createTime,

            int? policyVersion,

            ImmutableArray<Outputs.GetRouteTableInputPoliciesPolicySetPolicyResult> policys)
        {
            CreateTime = createTime;
            PolicyVersion = policyVersion;
            Policys = policys;
        }
    }
}