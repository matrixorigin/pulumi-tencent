// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tdmq.Outputs
{

    [OutputType]
    public sealed class GetRocketmqClusterClusterListResult
    {
        public readonly ImmutableArray<Outputs.GetRocketmqClusterClusterListConfigResult> Configs;
        public readonly ImmutableArray<Outputs.GetRocketmqClusterClusterListInfoResult> Infos;
        public readonly int Status;

        [OutputConstructor]
        private GetRocketmqClusterClusterListResult(
            ImmutableArray<Outputs.GetRocketmqClusterClusterListConfigResult> configs,

            ImmutableArray<Outputs.GetRocketmqClusterClusterListInfoResult> infos,

            int status)
        {
            Configs = configs;
            Infos = infos;
            Status = status;
        }
    }
}