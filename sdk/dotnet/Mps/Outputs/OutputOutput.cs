// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class OutputOutput
    {
        public readonly ImmutableArray<string> AllowIpLists;
        public readonly string Description;
        public readonly int? MaxConcurrent;
        public readonly string OutputName;
        public readonly string OutputRegion;
        public readonly string Protocol;
        public readonly Outputs.OutputOutputRtmpSettings? RtmpSettings;
        public readonly Outputs.OutputOutputRtpSettings? RtpSettings;
        public readonly Outputs.OutputOutputSrtSettings? SrtSettings;

        [OutputConstructor]
        private OutputOutput(
            ImmutableArray<string> allowIpLists,

            string description,

            int? maxConcurrent,

            string outputName,

            string outputRegion,

            string protocol,

            Outputs.OutputOutputRtmpSettings? rtmpSettings,

            Outputs.OutputOutputRtpSettings? rtpSettings,

            Outputs.OutputOutputSrtSettings? srtSettings)
        {
            AllowIpLists = allowIpLists;
            Description = description;
            MaxConcurrent = maxConcurrent;
            OutputName = outputName;
            OutputRegion = outputRegion;
            Protocol = protocol;
            RtmpSettings = rtmpSettings;
            RtpSettings = rtpSettings;
            SrtSettings = srtSettings;
        }
    }
}
