// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos
{
    public static class GetOverviewAttackTrend
    {
        public static Task<GetOverviewAttackTrendResult> InvokeAsync(GetOverviewAttackTrendArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOverviewAttackTrendResult>("tencentcloud:Antiddos/getOverviewAttackTrend:getOverviewAttackTrend", args ?? new GetOverviewAttackTrendArgs(), options.WithDefaults());

        public static Output<GetOverviewAttackTrendResult> Invoke(GetOverviewAttackTrendInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOverviewAttackTrendResult>("tencentcloud:Antiddos/getOverviewAttackTrend:getOverviewAttackTrend", args ?? new GetOverviewAttackTrendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOverviewAttackTrendArgs : global::Pulumi.InvokeArgs
    {
        [Input("dimension", required: true)]
        public string Dimension { get; set; } = null!;

        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        [Input("period", required: true)]
        public int Period { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetOverviewAttackTrendArgs()
        {
        }
        public static new GetOverviewAttackTrendArgs Empty => new GetOverviewAttackTrendArgs();
    }

    public sealed class GetOverviewAttackTrendInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dimension", required: true)]
        public Input<string> Dimension { get; set; } = null!;

        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetOverviewAttackTrendInvokeArgs()
        {
        }
        public static new GetOverviewAttackTrendInvokeArgs Empty => new GetOverviewAttackTrendInvokeArgs();
    }


    [OutputType]
    public sealed class GetOverviewAttackTrendResult
    {
        public readonly ImmutableArray<int> Datas;
        public readonly string Dimension;
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int Period;
        public readonly int PeriodPointCount;
        public readonly string? ResultOutputFile;
        public readonly string StartTime;
        public readonly string Type;

        [OutputConstructor]
        private GetOverviewAttackTrendResult(
            ImmutableArray<int> datas,

            string dimension,

            string endTime,

            string id,

            int period,

            int periodPointCount,

            string? resultOutputFile,

            string startTime,

            string type)
        {
            Datas = datas;
            Dimension = dimension;
            EndTime = endTime;
            Id = id;
            Period = period;
            PeriodPointCount = periodPointCount;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            Type = type;
        }
    }
}
