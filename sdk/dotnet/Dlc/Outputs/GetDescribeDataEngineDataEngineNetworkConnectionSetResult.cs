// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc.Outputs
{

    [OutputType]
    public sealed class GetDescribeDataEngineDataEngineNetworkConnectionSetResult
    {
        public readonly int Appid;
        public readonly string AssociateId;
        public readonly int CreateTime;
        public readonly string DatasourceConnectionCidrBlock;
        public readonly string DatasourceConnectionId;
        public readonly string DatasourceConnectionName;
        public readonly string DatasourceConnectionSubnetCidrBlock;
        public readonly string DatasourceConnectionSubnetId;
        public readonly string DatasourceConnectionVpcId;
        public readonly string HouseId;
        public readonly string HouseName;
        public readonly int Id;
        public readonly string NetworkConnectionDesc;
        public readonly int NetworkConnectionType;
        public readonly int State;
        public readonly string SubAccountUin;
        public readonly string Uin;
        public readonly int UpdateTime;

        [OutputConstructor]
        private GetDescribeDataEngineDataEngineNetworkConnectionSetResult(
            int appid,

            string associateId,

            int createTime,

            string datasourceConnectionCidrBlock,

            string datasourceConnectionId,

            string datasourceConnectionName,

            string datasourceConnectionSubnetCidrBlock,

            string datasourceConnectionSubnetId,

            string datasourceConnectionVpcId,

            string houseId,

            string houseName,

            int id,

            string networkConnectionDesc,

            int networkConnectionType,

            int state,

            string subAccountUin,

            string uin,

            int updateTime)
        {
            Appid = appid;
            AssociateId = associateId;
            CreateTime = createTime;
            DatasourceConnectionCidrBlock = datasourceConnectionCidrBlock;
            DatasourceConnectionId = datasourceConnectionId;
            DatasourceConnectionName = datasourceConnectionName;
            DatasourceConnectionSubnetCidrBlock = datasourceConnectionSubnetCidrBlock;
            DatasourceConnectionSubnetId = datasourceConnectionSubnetId;
            DatasourceConnectionVpcId = datasourceConnectionVpcId;
            HouseId = houseId;
            HouseName = houseName;
            Id = id;
            NetworkConnectionDesc = networkConnectionDesc;
            NetworkConnectionType = networkConnectionType;
            State = state;
            SubAccountUin = subAccountUin;
            Uin = uin;
            UpdateTime = updateTime;
        }
    }
}