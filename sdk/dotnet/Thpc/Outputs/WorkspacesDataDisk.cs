// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Thpc.Outputs
{

    [OutputType]
    public sealed class WorkspacesDataDisk
    {
        public readonly bool? BurstPerformance;
        public readonly bool? DeleteWithInstance;
        public readonly string? DiskId;
        public readonly int? DiskSize;
        public readonly string? DiskType;
        public readonly bool? Encrypt;
        public readonly string? KmsKeyId;
        public readonly string? SnapshotId;
        public readonly int? ThroughputPerformance;

        [OutputConstructor]
        private WorkspacesDataDisk(
            bool? burstPerformance,

            bool? deleteWithInstance,

            string? diskId,

            int? diskSize,

            string? diskType,

            bool? encrypt,

            string? kmsKeyId,

            string? snapshotId,

            int? throughputPerformance)
        {
            BurstPerformance = burstPerformance;
            DeleteWithInstance = deleteWithInstance;
            DiskId = diskId;
            DiskSize = diskSize;
            DiskType = diskType;
            Encrypt = encrypt;
            KmsKeyId = kmsKeyId;
            SnapshotId = snapshotId;
            ThroughputPerformance = throughputPerformance;
        }
    }
}