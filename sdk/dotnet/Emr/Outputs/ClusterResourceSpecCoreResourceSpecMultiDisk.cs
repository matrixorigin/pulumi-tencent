// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterResourceSpecCoreResourceSpecMultiDisk
    {
        public readonly int? Count;
        public readonly string? DiskType;
        public readonly int? Volume;

        [OutputConstructor]
        private ClusterResourceSpecCoreResourceSpecMultiDisk(
            int? count,

            string? diskType,

            int? volume)
        {
            Count = count;
            DiskType = diskType;
            Volume = volume;
        }
    }
}
