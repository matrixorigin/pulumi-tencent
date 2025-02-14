// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleActivityActivityParaArgs : global::Pulumi.ResourceArgs
    {
        [Input("adaptiveDynamicStreamingTask")]
        public Input<Inputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskArgs>? AdaptiveDynamicStreamingTask { get; set; }

        [Input("aiAnalysisTask")]
        public Input<Inputs.ScheduleActivityActivityParaAiAnalysisTaskArgs>? AiAnalysisTask { get; set; }

        [Input("aiContentReviewTask")]
        public Input<Inputs.ScheduleActivityActivityParaAiContentReviewTaskArgs>? AiContentReviewTask { get; set; }

        [Input("aiRecognitionTask")]
        public Input<Inputs.ScheduleActivityActivityParaAiRecognitionTaskArgs>? AiRecognitionTask { get; set; }

        [Input("animatedGraphicTask")]
        public Input<Inputs.ScheduleActivityActivityParaAnimatedGraphicTaskArgs>? AnimatedGraphicTask { get; set; }

        [Input("imageSpriteTask")]
        public Input<Inputs.ScheduleActivityActivityParaImageSpriteTaskArgs>? ImageSpriteTask { get; set; }

        [Input("sampleSnapshotTask")]
        public Input<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskArgs>? SampleSnapshotTask { get; set; }

        [Input("snapshotByTimeOffsetTask")]
        public Input<Inputs.ScheduleActivityActivityParaSnapshotByTimeOffsetTaskArgs>? SnapshotByTimeOffsetTask { get; set; }

        [Input("transcodeTask")]
        public Input<Inputs.ScheduleActivityActivityParaTranscodeTaskArgs>? TranscodeTask { get; set; }

        public ScheduleActivityActivityParaArgs()
        {
        }
        public static new ScheduleActivityActivityParaArgs Empty => new ScheduleActivityActivityParaArgs();
    }
}
