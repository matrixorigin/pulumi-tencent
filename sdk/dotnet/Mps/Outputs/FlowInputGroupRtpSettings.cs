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
    public sealed class FlowInputGroupRtpSettings
    {
        public readonly string? Fec;
        public readonly int? IdleTimeout;

        [OutputConstructor]
        private FlowInputGroupRtpSettings(
            string? fec,

            int? idleTimeout)
        {
            Fec = fec;
            IdleTimeout = idleTimeout;
        }
    }
}