// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetSubnetsInstanceListResult
    {
        public readonly string AvailabilityZone;
        public readonly int AvailableIpCount;
        public readonly string CdcId;
        public readonly string CidrBlock;
        public readonly string CreateTime;
        public readonly bool IsDefault;
        public readonly bool IsMulticast;
        public readonly string Name;
        public readonly string RouteTableId;
        public readonly string SubnetId;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetSubnetsInstanceListResult(
            string availabilityZone,

            int availableIpCount,

            string cdcId,

            string cidrBlock,

            string createTime,

            bool isDefault,

            bool isMulticast,

            string name,

            string routeTableId,

            string subnetId,

            ImmutableDictionary<string, object> tags,

            string vpcId)
        {
            AvailabilityZone = availabilityZone;
            AvailableIpCount = availableIpCount;
            CdcId = cdcId;
            CidrBlock = cidrBlock;
            CreateTime = createTime;
            IsDefault = isDefault;
            IsMulticast = isMulticast;
            Name = name;
            RouteTableId = routeTableId;
            SubnetId = subnetId;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}