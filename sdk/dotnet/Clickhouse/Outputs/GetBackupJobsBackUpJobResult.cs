// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clickhouse.Outputs
{

    [OutputType]
    public sealed class GetBackupJobsBackUpJobResult
    {
        public readonly int BackUpSize;
        public readonly string BackUpTime;
        public readonly string BackUpType;
        public readonly string ExpireTime;
        public readonly int JobId;
        public readonly string JobStatus;
        public readonly string Snapshot;

        [OutputConstructor]
        private GetBackupJobsBackUpJobResult(
            int backUpSize,

            string backUpTime,

            string backUpType,

            string expireTime,

            int jobId,

            string jobStatus,

            string snapshot)
        {
            BackUpSize = backUpSize;
            BackUpTime = backUpTime;
            BackUpType = backUpType;
            ExpireTime = expireTime;
            JobId = jobId;
            JobStatus = jobStatus;
            Snapshot = snapshot;
        }
    }
}