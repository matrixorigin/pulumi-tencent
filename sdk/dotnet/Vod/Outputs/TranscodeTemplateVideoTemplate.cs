// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class TranscodeTemplateVideoTemplate
    {
        public readonly int Bitrate;
        public readonly string Codec;
        public readonly string? CodecTag;
        public readonly string? FillType;
        public readonly int Fps;
        public readonly int? Gop;
        public readonly int? Height;
        public readonly string? PreserveHdrSwitch;
        public readonly string? ResolutionAdaptive;
        public readonly int? Vcrf;
        public readonly int? Width;

        [OutputConstructor]
        private TranscodeTemplateVideoTemplate(
            int bitrate,

            string codec,

            string? codecTag,

            string? fillType,

            int fps,

            int? gop,

            int? height,

            string? preserveHdrSwitch,

            string? resolutionAdaptive,

            int? vcrf,

            int? width)
        {
            Bitrate = bitrate;
            Codec = codec;
            CodecTag = codecTag;
            FillType = fillType;
            Fps = fps;
            Gop = gop;
            Height = height;
            PreserveHdrSwitch = preserveHdrSwitch;
            ResolutionAdaptive = resolutionAdaptive;
            Vcrf = vcrf;
            Width = width;
        }
    }
}