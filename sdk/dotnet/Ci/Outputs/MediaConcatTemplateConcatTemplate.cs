// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ci.Outputs
{

    [OutputType]
    public sealed class MediaConcatTemplateConcatTemplate
    {
        public readonly Outputs.MediaConcatTemplateConcatTemplateAudio? Audio;
        public readonly ImmutableArray<Outputs.MediaConcatTemplateConcatTemplateAudioMix> AudioMixes;
        public readonly ImmutableArray<Outputs.MediaConcatTemplateConcatTemplateConcatFragment> ConcatFragments;
        public readonly Outputs.MediaConcatTemplateConcatTemplateContainer Container;
        public readonly Outputs.MediaConcatTemplateConcatTemplateVideo? Video;

        [OutputConstructor]
        private MediaConcatTemplateConcatTemplate(
            Outputs.MediaConcatTemplateConcatTemplateAudio? audio,

            ImmutableArray<Outputs.MediaConcatTemplateConcatTemplateAudioMix> audioMixes,

            ImmutableArray<Outputs.MediaConcatTemplateConcatTemplateConcatFragment> concatFragments,

            Outputs.MediaConcatTemplateConcatTemplateContainer container,

            Outputs.MediaConcatTemplateConcatTemplateVideo? video)
        {
            Audio = audio;
            AudioMixes = audioMixes;
            ConcatFragments = concatFragments;
            Container = container;
            Video = video;
        }
    }
}