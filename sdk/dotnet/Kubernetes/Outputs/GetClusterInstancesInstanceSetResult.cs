// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class GetClusterInstancesInstanceSetResult
    {
        public readonly string AutoscalingGroupId;
        public readonly string CreatedTime;
        public readonly string DrainStatus;
        public readonly string FailedReason;
        public readonly ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingResult> InstanceAdvancedSettings;
        public readonly string InstanceId;
        public readonly string InstanceRole;
        public readonly string InstanceState;
        public readonly string LanIp;
        public readonly string NodePoolId;

        [OutputConstructor]
        private GetClusterInstancesInstanceSetResult(
            string autoscalingGroupId,

            string createdTime,

            string drainStatus,

            string failedReason,

            ImmutableArray<Outputs.GetClusterInstancesInstanceSetInstanceAdvancedSettingResult> instanceAdvancedSettings,

            string instanceId,

            string instanceRole,

            string instanceState,

            string lanIp,

            string nodePoolId)
        {
            AutoscalingGroupId = autoscalingGroupId;
            CreatedTime = createdTime;
            DrainStatus = drainStatus;
            FailedReason = failedReason;
            InstanceAdvancedSettings = instanceAdvancedSettings;
            InstanceId = instanceId;
            InstanceRole = instanceRole;
            InstanceState = instanceState;
            LanIp = lanIp;
            NodePoolId = nodePoolId;
        }
    }
}