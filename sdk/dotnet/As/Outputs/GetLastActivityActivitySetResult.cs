// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Outputs
{

    [OutputType]
    public sealed class GetLastActivityActivitySetResult
    {
        public readonly string ActivityId;
        public readonly ImmutableArray<Outputs.GetLastActivityActivitySetActivityRelatedInstanceSetResult> ActivityRelatedInstanceSets;
        public readonly string ActivityType;
        public readonly string AutoScalingGroupId;
        public readonly string Cause;
        public readonly string CreatedTime;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetLastActivityActivitySetDetailedStatusMessageSetResult> DetailedStatusMessageSets;
        public readonly string EndTime;
        public readonly ImmutableArray<Outputs.GetLastActivityActivitySetInvocationResultSetResult> InvocationResultSets;
        public readonly ImmutableArray<Outputs.GetLastActivityActivitySetLifecycleActionResultSetResult> LifecycleActionResultSets;
        public readonly string StartTime;
        public readonly string StatusCode;
        public readonly string StatusMessage;
        public readonly string StatusMessageSimplified;

        [OutputConstructor]
        private GetLastActivityActivitySetResult(
            string activityId,

            ImmutableArray<Outputs.GetLastActivityActivitySetActivityRelatedInstanceSetResult> activityRelatedInstanceSets,

            string activityType,

            string autoScalingGroupId,

            string cause,

            string createdTime,

            string description,

            ImmutableArray<Outputs.GetLastActivityActivitySetDetailedStatusMessageSetResult> detailedStatusMessageSets,

            string endTime,

            ImmutableArray<Outputs.GetLastActivityActivitySetInvocationResultSetResult> invocationResultSets,

            ImmutableArray<Outputs.GetLastActivityActivitySetLifecycleActionResultSetResult> lifecycleActionResultSets,

            string startTime,

            string statusCode,

            string statusMessage,

            string statusMessageSimplified)
        {
            ActivityId = activityId;
            ActivityRelatedInstanceSets = activityRelatedInstanceSets;
            ActivityType = activityType;
            AutoScalingGroupId = autoScalingGroupId;
            Cause = cause;
            CreatedTime = createdTime;
            Description = description;
            DetailedStatusMessageSets = detailedStatusMessageSets;
            EndTime = endTime;
            InvocationResultSets = invocationResultSets;
            LifecycleActionResultSets = lifecycleActionResultSets;
            StartTime = startTime;
            StatusCode = statusCode;
            StatusMessage = statusMessage;
            StatusMessageSimplified = statusMessageSimplified;
        }
    }
}