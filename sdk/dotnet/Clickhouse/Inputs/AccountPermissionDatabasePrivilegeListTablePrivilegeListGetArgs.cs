// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clickhouse.Inputs
{

    public sealed class AccountPermissionDatabasePrivilegeListTablePrivilegeListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        [Input("tablePrivileges", required: true)]
        private InputList<string>? _tablePrivileges;
        public InputList<string> TablePrivileges
        {
            get => _tablePrivileges ?? (_tablePrivileges = new InputList<string>());
            set => _tablePrivileges = value;
        }

        public AccountPermissionDatabasePrivilegeListTablePrivilegeListGetArgs()
        {
        }
        public static new AccountPermissionDatabasePrivilegeListTablePrivilegeListGetArgs Empty => new AccountPermissionDatabasePrivilegeListTablePrivilegeListGetArgs();
    }
}