// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class GetOverviewDdosEventListEventListResult
    {
        public readonly int AttackStatus;
        public readonly string AttackType;
        public readonly string Business;
        public readonly string EndTime;
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string InstanceName;
        public readonly int Mbps;
        public readonly int Pps;
        public readonly string StartTime;
        public readonly string Vip;

        [OutputConstructor]
        private GetOverviewDdosEventListEventListResult(
            int attackStatus,

            string attackType,

            string business,

            string endTime,

            string id,

            string instanceId,

            string instanceName,

            int mbps,

            int pps,

            string startTime,

            string vip)
        {
            AttackStatus = attackStatus;
            AttackType = attackType;
            Business = business;
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            Mbps = mbps;
            Pps = pps;
            StartTime = startTime;
            Vip = vip;
        }
    }
}