// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdwdoris.Outputs
{

    [OutputType]
    public sealed class GetInstancesSearchTagResult
    {
        public readonly int? AllValue;
        public readonly string? TagKey;
        public readonly string? TagValue;

        [OutputConstructor]
        private GetInstancesSearchTagResult(
            int? allValue,

            string? tagKey,

            string? tagValue)
        {
            AllValue = allValue;
            TagKey = tagKey;
            TagValue = tagValue;
        }
    }
}
