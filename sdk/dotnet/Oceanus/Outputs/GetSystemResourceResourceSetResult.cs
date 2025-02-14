// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class GetSystemResourceResourceSetResult
    {
        public readonly int LatestResourceConfigVersion;
        public readonly string Name;
        public readonly string Region;
        public readonly string Remark;
        public readonly string ResourceId;
        public readonly int ResourceType;

        [OutputConstructor]
        private GetSystemResourceResourceSetResult(
            int latestResourceConfigVersion,

            string name,

            string region,

            string remark,

            string resourceId,

            int resourceType)
        {
            LatestResourceConfigVersion = latestResourceConfigVersion;
            Name = name;
            Region = region;
            Remark = remark;
            ResourceId = resourceId;
            ResourceType = resourceType;
        }
    }
}
