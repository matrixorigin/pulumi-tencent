// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Bi.Outputs
{

    [OutputType]
    public sealed class DatasourceCloudServiceType
    {
        public readonly string InstanceId;
        public readonly string Region;
        public readonly string Type;

        [OutputConstructor]
        private DatasourceCloudServiceType(
            string instanceId,

            string region,

            string type)
        {
            InstanceId = instanceId;
            Region = region;
            Type = type;
        }
    }
}