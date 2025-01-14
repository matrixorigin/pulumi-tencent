// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class WatermarkTemplateImageTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("height")]
        public Input<string>? Height { get; set; }

        [Input("imageContent", required: true)]
        public Input<string> ImageContent { get; set; } = null!;

        [Input("repeatType")]
        public Input<string>? RepeatType { get; set; }

        [Input("transparency")]
        public Input<int>? Transparency { get; set; }

        [Input("width")]
        public Input<string>? Width { get; set; }

        public WatermarkTemplateImageTemplateGetArgs()
        {
        }
        public static new WatermarkTemplateImageTemplateGetArgs Empty => new WatermarkTemplateImageTemplateGetArgs();
    }
}
