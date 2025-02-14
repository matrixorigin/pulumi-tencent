// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    public static class GetAlarmMetric
    {
        public static Task<GetAlarmMetricResult> InvokeAsync(GetAlarmMetricArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlarmMetricResult>("tencentcloud:Monitor/getAlarmMetric:getAlarmMetric", args ?? new GetAlarmMetricArgs(), options.WithDefaults());

        public static Output<GetAlarmMetricResult> Invoke(GetAlarmMetricInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmMetricResult>("tencentcloud:Monitor/getAlarmMetric:getAlarmMetric", args ?? new GetAlarmMetricInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlarmMetricArgs : global::Pulumi.InvokeArgs
    {
        [Input("module", required: true)]
        public string Module { get; set; } = null!;

        [Input("monitorType", required: true)]
        public string MonitorType { get; set; } = null!;

        [Input("namespace", required: true)]
        public string Namespace { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAlarmMetricArgs()
        {
        }
        public static new GetAlarmMetricArgs Empty => new GetAlarmMetricArgs();
    }

    public sealed class GetAlarmMetricInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("module", required: true)]
        public Input<string> Module { get; set; } = null!;

        [Input("monitorType", required: true)]
        public Input<string> MonitorType { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAlarmMetricInvokeArgs()
        {
        }
        public static new GetAlarmMetricInvokeArgs Empty => new GetAlarmMetricInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlarmMetricResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetAlarmMetricMetricResult> Metrics;
        public readonly string Module;
        public readonly string MonitorType;
        public readonly string Namespace;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetAlarmMetricResult(
            string id,

            ImmutableArray<Outputs.GetAlarmMetricMetricResult> metrics,

            string module,

            string monitorType,

            string @namespace,

            string? resultOutputFile)
        {
            Id = id;
            Metrics = metrics;
            Module = module;
            MonitorType = monitorType;
            Namespace = @namespace;
            ResultOutputFile = resultOutputFile;
        }
    }
}
