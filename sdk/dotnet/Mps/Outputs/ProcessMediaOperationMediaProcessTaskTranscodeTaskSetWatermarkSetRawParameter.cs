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
    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameter
    {
        public readonly string? CoordinateOrigin;
        public readonly Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameterImageTemplate? ImageTemplate;
        public readonly string Type;
        public readonly string? XPos;
        public readonly string? YPos;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameter(
            string? coordinateOrigin,

            Outputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameterImageTemplate? imageTemplate,

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