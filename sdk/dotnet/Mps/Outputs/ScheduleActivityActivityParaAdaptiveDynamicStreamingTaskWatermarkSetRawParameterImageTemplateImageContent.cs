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
    public sealed class ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContent
    {
        public readonly Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentCosInputInfo? CosInputInfo;
        public readonly Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfo? S3InputInfo;
        public readonly string Type;
        public readonly Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfo? UrlInputInfo;

        [OutputConstructor]
        private ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContent(
            Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentCosInputInfo? cosInputInfo,

            Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfo? s3InputInfo,

            string type,

            Outputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentUrlInputInfo? urlInputInfo)
        {
            CosInputInfo = cosInputInfo;
            S3InputInfo = s3InputInfo;
            Type = type;
            UrlInputInfo = urlInputInfo;
        }
    }
}