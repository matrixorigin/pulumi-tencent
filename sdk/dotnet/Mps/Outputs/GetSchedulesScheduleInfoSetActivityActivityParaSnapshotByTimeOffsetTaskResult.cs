// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskResult
    {
        public readonly int Definition;
        public readonly ImmutableArray<string> ExtTimeOffsetSets;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskObjectNumberFormatResult> ObjectNumberFormats;
        public readonly string OutputObjectPath;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskOutputStorageResult> OutputStorages;
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSetResult> WatermarkSets;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskResult(
            int definition,

            ImmutableArray<string> extTimeOffsetSets,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskObjectNumberFormatResult> objectNumberFormats,

            string outputObjectPath,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskOutputStorageResult> outputStorages,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSetResult> watermarkSets)
        {
            Definition = definition;
            ExtTimeOffsetSets = extTimeOffsetSets;
            ObjectNumberFormats = objectNumberFormats;
            OutputObjectPath = outputObjectPath;
            OutputStorages = outputStorages;
            WatermarkSets = watermarkSets;
        }
    }
}
