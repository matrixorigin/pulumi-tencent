// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sqlserver.Outputs
{

    [OutputType]
    public sealed class GetBackupsListResult
    {
        public readonly ImmutableArray<string> DbLists;
        public readonly string EndTime;
        public readonly string FileName;
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string InternetUrl;
        public readonly string IntranetUrl;
        public readonly int Size;
        public readonly string StartTime;
        public readonly int Status;
        public readonly int Strategy;
        public readonly int TriggerModel;

        [OutputConstructor]
        private GetBackupsListResult(
            ImmutableArray<string> dbLists,

            string endTime,

            string fileName,

            string id,

            string instanceId,

            string internetUrl,

            string intranetUrl,

            int size,

            string startTime,

            int status,

            int strategy,

            int triggerModel)
        {
            DbLists = dbLists;
            EndTime = endTime;
            FileName = fileName;
            Id = id;
            InstanceId = instanceId;
            InternetUrl = internetUrl;
            IntranetUrl = intranetUrl;
            Size = size;
            StartTime = startTime;
            Status = status;
            Strategy = strategy;
            TriggerModel = triggerModel;
        }
    }
}