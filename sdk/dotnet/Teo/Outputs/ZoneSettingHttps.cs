// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class ZoneSettingHttps
    {
        public readonly Outputs.ZoneSettingHttpsHsts? Hsts;
        public readonly string? Http2;
        public readonly string? OcspStapling;
        public readonly ImmutableArray<string> TlsVersions;

        [OutputConstructor]
        private ZoneSettingHttps(
            Outputs.ZoneSettingHttpsHsts? hsts,

            string? http2,

            string? ocspStapling,

            ImmutableArray<string> tlsVersions)
        {
            Hsts = hsts;
            Http2 = http2;
            OcspStapling = ocspStapling;
            TlsVersions = tlsVersions;
        }
    }
}