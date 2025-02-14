// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql.Outputs
{

    [OutputType]
    public sealed class AccountPrivilegesOperationModifyPrivilegeSetDatabasePrivilegeObject
    {
        public readonly string? DatabaseName;
        public readonly string ObjectName;
        public readonly string ObjectType;
        public readonly string? SchemaName;
        public readonly string? TableName;

        [OutputConstructor]
        private AccountPrivilegesOperationModifyPrivilegeSetDatabasePrivilegeObject(
            string? databaseName,

            string objectName,

            string objectType,

            string? schemaName,

            string? tableName)
        {
            DatabaseName = databaseName;
            ObjectName = objectName;
            ObjectType = objectType;
            SchemaName = schemaName;
            TableName = tableName;
        }
    }
}
