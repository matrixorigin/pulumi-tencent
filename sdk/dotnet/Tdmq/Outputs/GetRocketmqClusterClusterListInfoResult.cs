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
    public sealed class GetRocketmqClusterClusterListInfoResult
    {
        public readonly string ClusterId;
        public readonly string ClusterName;
        public readonly int CreateTime;
        public readonly bool IsVip;
        public readonly string PublicEndPoint;
        public readonly string Region;
        public readonly string Remark;
        public readonly bool RocketmqFlag;
        public readonly bool SupportNamespaceEndpoint;
        public readonly string VpcEndPoint;
        public readonly ImmutableArray<Outputs.GetRocketmqClusterClusterListInfoVpcResult> Vpcs;

        [OutputConstructor]
        private GetRocketmqClusterClusterListInfoResult(
            string clusterId,

            string clusterName,

            int createTime,

            bool isVip,

            string publicEndPoint,

            string region,

            string remark,

            bool rocketmqFlag,

            bool supportNamespaceEndpoint,

            string vpcEndPoint,

            ImmutableArray<Outputs.GetRocketmqClusterClusterListInfoVpcResult> vpcs)
        {
            ClusterId = clusterId;
            ClusterName = clusterName;
            CreateTime = createTime;
            IsVip = isVip;
            PublicEndPoint = publicEndPoint;
            Region = region;
            Remark = remark;
            RocketmqFlag = rocketmqFlag;
            SupportNamespaceEndpoint = supportNamespaceEndpoint;
            VpcEndPoint = vpcEndPoint;
            Vpcs = vpcs;
        }
    }
}