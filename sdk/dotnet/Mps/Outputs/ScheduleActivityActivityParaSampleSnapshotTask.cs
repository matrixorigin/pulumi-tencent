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
    public sealed class ScheduleActivityActivityParaSampleSnapshotTask
    {
        public readonly int Definition;
        public readonly Outputs.ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormat? ObjectNumberFormat;
        public readonly string? OutputObjectPath;
        public readonly Outputs.ScheduleActivityActivityParaSampleSnapshotTaskOutputStorage? OutputStorage;
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSet> WatermarkSets;

        [OutputConstructor]
        private ScheduleActivityActivityParaSampleSnapshotTask(
            int definition,

            Outputs.ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormat? objectNumberFormat,

            string? outputObjectPath,

            Outputs.ScheduleActivityActivityParaSampleSnapshotTaskOutputStorage? outputStorage,

            ImmutableArray<Outputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSet> watermarkSets)
        {
            Definition = definition;
            ObjectNumberFormat = objectNumberFormat;
            OutputObjectPath = outputObjectPath;
            OutputStorage = outputStorage;
            WatermarkSets = watermarkSets;
        }
    }
}
