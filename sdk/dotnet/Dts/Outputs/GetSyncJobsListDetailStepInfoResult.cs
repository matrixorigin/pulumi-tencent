// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetSyncJobsListDetailStepInfoResult
    {
        public readonly ImmutableArray<Outputs.GetSyncJobsListDetailStepInfoErrorResult> Errors;
        public readonly int Progress;
        public readonly string StartTime;
        public readonly string Status;
        public readonly string StepId;
        public readonly string StepName;
        public readonly int StepNo;
        public readonly ImmutableArray<Outputs.GetSyncJobsListDetailStepInfoWarningResult> Warnings;

        [OutputConstructor]
        private GetSyncJobsListDetailStepInfoResult(
            ImmutableArray<Outputs.GetSyncJobsListDetailStepInfoErrorResult> errors,

            int progress,

            string startTime,

            string status,

            string stepId,

            string stepName,

            int stepNo,

            ImmutableArray<Outputs.GetSyncJobsListDetailStepInfoWarningResult> warnings)
        {
            Errors = errors;
            Progress = progress;
            StartTime = startTime;
            Status = status;
            StepId = stepId;
            StepName = stepName;
            StepNo = stepNo;
            Warnings = warnings;
        }
    }
}