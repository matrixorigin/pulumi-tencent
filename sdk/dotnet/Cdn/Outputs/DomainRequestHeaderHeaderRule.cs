// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainRequestHeaderHeaderRule
    {
        public readonly string HeaderMode;
        public readonly string HeaderName;
        public readonly string HeaderValue;
        public readonly ImmutableArray<string> RulePaths;
        public readonly string RuleType;

        [OutputConstructor]
        private DomainRequestHeaderHeaderRule(
            string headerMode,

            string headerName,

            string headerValue,

            ImmutableArray<string> rulePaths,

            string ruleType)
        {
            HeaderMode = headerMode;
            HeaderName = headerName;
            HeaderValue = headerValue;
            RulePaths = rulePaths;
            RuleType = ruleType;
        }
    }
}