// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    public static class GetTimeShiftStreamList
    {
        public static Task<GetTimeShiftStreamListResult> InvokeAsync(GetTimeShiftStreamListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTimeShiftStreamListResult>("tencentcloud:Css/getTimeShiftStreamList:getTimeShiftStreamList", args ?? new GetTimeShiftStreamListArgs(), options.WithDefaults());

        public static Output<GetTimeShiftStreamListResult> Invoke(GetTimeShiftStreamListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTimeShiftStreamListResult>("tencentcloud:Css/getTimeShiftStreamList:getTimeShiftStreamList", args ?? new GetTimeShiftStreamListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTimeShiftStreamListArgs : global::Pulumi.InvokeArgs
    {
        [Input("domain")]
        public string? Domain { get; set; }

        [Input("domainGroup")]
        public string? DomainGroup { get; set; }

        [Input("endTime", required: true)]
        public int EndTime { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public int StartTime { get; set; }

        [Input("streamName")]
        public string? StreamName { get; set; }

        public GetTimeShiftStreamListArgs()
        {
        }
        public static new GetTimeShiftStreamListArgs Empty => new GetTimeShiftStreamListArgs();
    }

    public sealed class GetTimeShiftStreamListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("domainGroup")]
        public Input<string>? DomainGroup { get; set; }

        [Input("endTime", required: true)]
        public Input<int> EndTime { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public Input<int> StartTime { get; set; } = null!;

        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        public GetTimeShiftStreamListInvokeArgs()
        {
        }
        public static new GetTimeShiftStreamListInvokeArgs Empty => new GetTimeShiftStreamListInvokeArgs();
    }


    [OutputType]
    public sealed class GetTimeShiftStreamListResult
    {
        public readonly string? Domain;
        public readonly string? DomainGroup;
        public readonly int EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int StartTime;
        public readonly ImmutableArray<Outputs.GetTimeShiftStreamListStreamListResult> StreamLists;
        public readonly string? StreamName;
        public readonly int TotalSize;

        [OutputConstructor]
        private GetTimeShiftStreamListResult(
            string? domain,

            string? domainGroup,

            int endTime,

            string id,

            string? resultOutputFile,

            int startTime,

            ImmutableArray<Outputs.GetTimeShiftStreamListStreamListResult> streamLists,

            string? streamName,

            int totalSize)
        {
            Domain = domain;
            DomainGroup = domainGroup;
            EndTime = endTime;
            Id = id;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            StreamLists = streamLists;
            StreamName = streamName;
            TotalSize = totalSize;
        }
    }
}