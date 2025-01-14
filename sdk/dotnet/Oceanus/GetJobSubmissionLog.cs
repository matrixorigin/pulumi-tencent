// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus
{
    public static class GetJobSubmissionLog
    {
        public static Task<GetJobSubmissionLogResult> InvokeAsync(GetJobSubmissionLogArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJobSubmissionLogResult>("tencentcloud:Oceanus/getJobSubmissionLog:getJobSubmissionLog", args ?? new GetJobSubmissionLogArgs(), options.WithDefaults());

        public static Output<GetJobSubmissionLogResult> Invoke(GetJobSubmissionLogInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobSubmissionLogResult>("tencentcloud:Oceanus/getJobSubmissionLog:getJobSubmissionLog", args ?? new GetJobSubmissionLogInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobSubmissionLogArgs : global::Pulumi.InvokeArgs
    {
        [Input("cursor")]
        public string? Cursor { get; set; }

        [Input("endTime", required: true)]
        public int EndTime { get; set; }

        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        [Input("keyword")]
        public string? Keyword { get; set; }

        [Input("orderType")]
        public string? OrderType { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("runningOrderId")]
        public int? RunningOrderId { get; set; }

        [Input("startTime", required: true)]
        public int StartTime { get; set; }

        public GetJobSubmissionLogArgs()
        {
        }
        public static new GetJobSubmissionLogArgs Empty => new GetJobSubmissionLogArgs();
    }

    public sealed class GetJobSubmissionLogInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cursor")]
        public Input<string>? Cursor { get; set; }

        [Input("endTime", required: true)]
        public Input<int> EndTime { get; set; } = null!;

        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        [Input("keyword")]
        public Input<string>? Keyword { get; set; }

        [Input("orderType")]
        public Input<string>? OrderType { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("runningOrderId")]
        public Input<int>? RunningOrderId { get; set; }

        [Input("startTime", required: true)]
        public Input<int> StartTime { get; set; } = null!;

        public GetJobSubmissionLogInvokeArgs()
        {
        }
        public static new GetJobSubmissionLogInvokeArgs Empty => new GetJobSubmissionLogInvokeArgs();
    }


    [OutputType]
    public sealed class GetJobSubmissionLogResult
    {
        public readonly string Cursor;
        public readonly int EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string JobId;
        public readonly ImmutableArray<Outputs.GetJobSubmissionLogJobInstanceListResult> JobInstanceLists;
        public readonly string JobRequestId;
        public readonly string? Keyword;
        public readonly bool ListOver;
        public readonly ImmutableArray<Outputs.GetJobSubmissionLogLogContentListResult> LogContentLists;
        public readonly ImmutableArray<string> LogLists;
        public readonly string? OrderType;
        public readonly string? ResultOutputFile;
        public readonly int? RunningOrderId;
        public readonly int StartTime;

        [OutputConstructor]
        private GetJobSubmissionLogResult(
            string cursor,

            int endTime,

            string id,

            string jobId,

            ImmutableArray<Outputs.GetJobSubmissionLogJobInstanceListResult> jobInstanceLists,

            string jobRequestId,

            string? keyword,

            bool listOver,

            ImmutableArray<Outputs.GetJobSubmissionLogLogContentListResult> logContentLists,

            ImmutableArray<string> logLists,

            string? orderType,

            string? resultOutputFile,

            int? runningOrderId,

            int startTime)
        {
            Cursor = cursor;
            EndTime = endTime;
            Id = id;
            JobId = jobId;
            JobInstanceLists = jobInstanceLists;
            JobRequestId = jobRequestId;
            Keyword = keyword;
            ListOver = listOver;
            LogContentLists = logContentLists;
            LogLists = logLists;
            OrderType = orderType;
            ResultOutputFile = resultOutputFile;
            RunningOrderId = runningOrderId;
            StartTime = startTime;
        }
    }
}
