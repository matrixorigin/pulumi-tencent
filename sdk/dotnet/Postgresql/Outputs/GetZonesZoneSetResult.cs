// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql.Outputs
{

    [OutputType]
    public sealed class GetZonesZoneSetResult
    {
        public readonly ImmutableArray<string> StandbyZoneSets;
        public readonly string Zone;
        public readonly int ZoneId;
        public readonly string ZoneName;
        public readonly string ZoneState;
        public readonly int ZoneSupportIpv6;

        [OutputConstructor]
        private GetZonesZoneSetResult(
            ImmutableArray<string> standbyZoneSets,

            string zone,

            int zoneId,

            string zoneName,

            string zoneState,

            int zoneSupportIpv6)
        {
            StandbyZoneSets = standbyZoneSets;
            Zone = zone;
            ZoneId = zoneId;
            ZoneName = zoneName;
            ZoneState = zoneState;
            ZoneSupportIpv6 = zoneSupportIpv6;
        }
    }
}