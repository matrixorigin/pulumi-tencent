// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css.Outputs
{

    [OutputType]
    public sealed class GetMonitorReportDiagnoseResultResult
    {
        public readonly ImmutableArray<string> LowFrameRateResults;
        public readonly ImmutableArray<string> StreamBrokenResults;
        public readonly ImmutableArray<string> StreamFormatResults;

        [OutputConstructor]
        private GetMonitorReportDiagnoseResultResult(
            ImmutableArray<string> lowFrameRateResults,

            ImmutableArray<string> streamBrokenResults,

            ImmutableArray<string> streamFormatResults)
        {
            LowFrameRateResults = lowFrameRateResults;
            StreamBrokenResults = streamBrokenResults;
            StreamFormatResults = streamFormatResults;
        }
    }
}