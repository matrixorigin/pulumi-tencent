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
    public sealed class DomainIpFilter
    {
        public readonly ImmutableArray<Outputs.DomainIpFilterFilterRule> FilterRules;
        public readonly string? FilterType;
        public readonly ImmutableArray<string> Filters;
        public readonly int? ReturnCode;
        public readonly string Switch;

        [OutputConstructor]
        private DomainIpFilter(
            ImmutableArray<Outputs.DomainIpFilterFilterRule> filterRules,

            string? filterType,

            ImmutableArray<string> filters,

            int? returnCode,

            string @switch)
        {
            FilterRules = filterRules;
            FilterType = filterType;
            Filters = filters;
            ReturnCode = returnCode;
            Switch = @switch;
        }
    }
}