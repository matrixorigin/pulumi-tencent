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
    public sealed class GetSyncJobsListDetailResult
    {
        public readonly int CurrentStepProgress;
        public readonly int MasterSlaveDistance;
        public readonly string Message;
        public readonly int Progress;
        public readonly int SecondsBehindMaster;
        public readonly int StepAll;
        public readonly ImmutableArray<Outputs.GetSyncJobsListDetailStepInfoResult> StepInfos;
        public readonly int StepNow;

        [OutputConstructor]
        private GetSyncJobsListDetailResult(
            int currentStepProgress,

            int masterSlaveDistance,

            string message,

            int progress,

            int secondsBehindMaster,

            int stepAll,

            ImmutableArray<Outputs.GetSyncJobsListDetailStepInfoResult> stepInfos,

            int stepNow)
        {
            CurrentStepProgress = currentStepProgress;
            MasterSlaveDistance = masterSlaveDistance;
            Message = message;
            Progress = progress;
            SecondsBehindMaster = secondsBehindMaster;
            StepAll = stepAll;
            StepInfos = stepInfos;
            StepNow = stepNow;
        }
    }
}