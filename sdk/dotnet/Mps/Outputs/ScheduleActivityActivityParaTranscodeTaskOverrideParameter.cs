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
    public sealed class ScheduleActivityActivityParaTranscodeTaskOverrideParameter
    {
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitle> AddOnSubtitles;
        public readonly ImmutableArray<Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStream> AddonAudioStreams;
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAudioTemplate? AudioTemplate;
        public readonly string? Container;
        public readonly int? RemoveAudio;
        public readonly int? RemoveVideo;
        public readonly string? StdExtInfo;
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplate? SubtitleTemplate;
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterTehdConfig? TehdConfig;
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterVideoTemplate? VideoTemplate;

        [OutputConstructor]
        private ScheduleActivityActivityParaTranscodeTaskOverrideParameter(
            ImmutableArray<Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddOnSubtitle> addOnSubtitles,

            ImmutableArray<Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAddonAudioStream> addonAudioStreams,

            Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterAudioTemplate? audioTemplate,

            string? container,

            int? removeAudio,

            int? removeVideo,

            string? stdExtInfo,

            Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterSubtitleTemplate? subtitleTemplate,

            Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterTehdConfig? tehdConfig,

            Outputs.ScheduleActivityActivityParaTranscodeTaskOverrideParameterVideoTemplate? videoTemplate)
        {
            AddOnSubtitles = addOnSubtitles;
            AddonAudioStreams = addonAudioStreams;
            AudioTemplate = audioTemplate;
            Container = container;
            RemoveAudio = removeAudio;
            RemoveVideo = removeVideo;
            StdExtInfo = stdExtInfo;
            SubtitleTemplate = subtitleTemplate;
            TehdConfig = tehdConfig;
            VideoTemplate = videoTemplate;
        }
    }
}