// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver
{
    public static class GetSlowlogs
    {
        public static Task<GetSlowlogsResult> InvokeAsync(GetSlowlogsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSlowlogsResult>("tencentcloud:Sqlserver/getSlowlogs:getSlowlogs", args ?? new GetSlowlogsArgs(), options.WithDefaults());

        public static Output<GetSlowlogsResult> Invoke(GetSlowlogsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSlowlogsResult>("tencentcloud:Sqlserver/getSlowlogs:getSlowlogs", args ?? new GetSlowlogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowlogsArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetSlowlogsArgs()
        {
        }
        public static new GetSlowlogsArgs Empty => new GetSlowlogsArgs();
    }

    public sealed class GetSlowlogsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetSlowlogsInvokeArgs()
        {
        }
        public static new GetSlowlogsInvokeArgs Empty => new GetSlowlogsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSlowlogsResult
    {
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.GetSlowlogsSlowlogResult> SqlserverSlowlogs;
        public readonly string StartTime;

        [OutputConstructor]
        private GetSlowlogsResult(
            string endTime,

            string id,

            string instanceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSlowlogsSlowlogResult> slowlogs,

            string startTime)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            SqlserverSlowlogs = slowlogs;
            StartTime = startTime;
        }
    }
}