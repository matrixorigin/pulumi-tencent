// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb.Outputs
{

    [OutputType]
    public sealed class GetClusterDetailDatabasesDbInfoResult
    {
        public readonly int AppId;
        public readonly string CharacterSet;
        public readonly string ClusterId;
        public readonly string CollateRule;
        public readonly string CreateTime;
        public readonly int DbId;
        public readonly string DbName;
        public readonly string Description;
        public readonly string Status;
        public readonly string Uin;
        public readonly string UpdateTime;
        public readonly ImmutableArray<Outputs.GetClusterDetailDatabasesDbInfoUserHostPrivilegeResult> UserHostPrivileges;

        [OutputConstructor]
        private GetClusterDetailDatabasesDbInfoResult(
            int appId,

            string characterSet,

            string clusterId,

            string collateRule,

            string createTime,

            int dbId,

            string dbName,

            string description,

            string status,

            string uin,

            string updateTime,

            ImmutableArray<Outputs.GetClusterDetailDatabasesDbInfoUserHostPrivilegeResult> userHostPrivileges)
        {
            AppId = appId;
            CharacterSet = characterSet;
            ClusterId = clusterId;
            CollateRule = collateRule;
            CreateTime = createTime;
            DbId = dbId;
            DbName = dbName;
            Description = description;
            Status = status;
            Uin = uin;
            UpdateTime = updateTime;
            UserHostPrivileges = userHostPrivileges;
        }
    }
}