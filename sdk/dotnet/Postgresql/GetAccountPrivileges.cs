// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql
{
    public static class GetAccountPrivileges
    {
        public static Task<GetAccountPrivilegesResult> InvokeAsync(GetAccountPrivilegesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountPrivilegesResult>("tencentcloud:Postgresql/getAccountPrivileges:getAccountPrivileges", args ?? new GetAccountPrivilegesArgs(), options.WithDefaults());

        public static Output<GetAccountPrivilegesResult> Invoke(GetAccountPrivilegesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountPrivilegesResult>("tencentcloud:Postgresql/getAccountPrivileges:getAccountPrivileges", args ?? new GetAccountPrivilegesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountPrivilegesArgs : global::Pulumi.InvokeArgs
    {
        [Input("databaseObjectSets", required: true)]
        private List<Inputs.GetAccountPrivilegesDatabaseObjectSetArgs>? _databaseObjectSets;
        public List<Inputs.GetAccountPrivilegesDatabaseObjectSetArgs> DatabaseObjectSets
        {
            get => _databaseObjectSets ?? (_databaseObjectSets = new List<Inputs.GetAccountPrivilegesDatabaseObjectSetArgs>());
            set => _databaseObjectSets = value;
        }

        [Input("dbInstanceId", required: true)]
        public string DbInstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetAccountPrivilegesArgs()
        {
        }
        public static new GetAccountPrivilegesArgs Empty => new GetAccountPrivilegesArgs();
    }

    public sealed class GetAccountPrivilegesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("databaseObjectSets", required: true)]
        private InputList<Inputs.GetAccountPrivilegesDatabaseObjectSetInputArgs>? _databaseObjectSets;
        public InputList<Inputs.GetAccountPrivilegesDatabaseObjectSetInputArgs> DatabaseObjectSets
        {
            get => _databaseObjectSets ?? (_databaseObjectSets = new InputList<Inputs.GetAccountPrivilegesDatabaseObjectSetInputArgs>());
            set => _databaseObjectSets = value;
        }

        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetAccountPrivilegesInvokeArgs()
        {
        }
        public static new GetAccountPrivilegesInvokeArgs Empty => new GetAccountPrivilegesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountPrivilegesResult
    {
        public readonly ImmutableArray<Outputs.GetAccountPrivilegesDatabaseObjectSetResult> DatabaseObjectSets;
        public readonly string DbInstanceId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetAccountPrivilegesPrivilegeSetResult> PrivilegeSets;
        public readonly string? ResultOutputFile;
        public readonly string UserName;

        [OutputConstructor]
        private GetAccountPrivilegesResult(
            ImmutableArray<Outputs.GetAccountPrivilegesDatabaseObjectSetResult> databaseObjectSets,

            string dbInstanceId,

            string id,

            ImmutableArray<Outputs.GetAccountPrivilegesPrivilegeSetResult> privilegeSets,

            string? resultOutputFile,

            string userName)
        {
            DatabaseObjectSets = databaseObjectSets;
            DbInstanceId = dbInstanceId;
            Id = id;
            PrivilegeSets = privilegeSets;
            ResultOutputFile = resultOutputFile;
            UserName = userName;
        }
    }
}