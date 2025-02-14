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
    public sealed class ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameter
    {
        public readonly string? CoordinateOrigin;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplate? ImageTemplate;
        public readonly string Type;
        public readonly string? XPos;
        public readonly string? YPos;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameter(
            string? coordinateOrigin,

            Outputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetRawParameterImageTemplate? imageTemplate,

            string type,

            string? xPos,

            string? yPos)
        {
            CoordinateOrigin = coordinateOrigin;
            ImageTemplate = imageTemplate;
            Type = type;
            XPos = xPos;
            YPos = yPos;
        }
    }
}
