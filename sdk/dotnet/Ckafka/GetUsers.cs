// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka
{
    public static class GetUsers
    {
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("tencentcloud:Ckafka/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsersResult>("tencentcloud:Ckafka/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountName")]
        public string? AccountName { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetUsersArgs()
        {
        }
        public static new GetUsersArgs Empty => new GetUsersArgs();
    }

    public sealed class GetUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetUsersInvokeArgs()
        {
        }
        public static new GetUsersInvokeArgs Empty => new GetUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        public readonly string? AccountName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetUsersUserListResult> UserLists;

        [OutputConstructor]
        private GetUsersResult(
            string? accountName,

            string id,

            string instanceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetUsersUserListResult> userLists)
        {
            AccountName = accountName;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            UserLists = userLists;
        }
    }
}