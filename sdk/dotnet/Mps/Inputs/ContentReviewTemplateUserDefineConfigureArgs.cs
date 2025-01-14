// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ContentReviewTemplateUserDefineConfigureArgs : global::Pulumi.ResourceArgs
    {
        [Input("asrReviewInfo")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureAsrReviewInfoArgs>? AsrReviewInfo { get; set; }

        [Input("faceReviewInfo")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureFaceReviewInfoArgs>? FaceReviewInfo { get; set; }

        [Input("ocrReviewInfo")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureOcrReviewInfoArgs>? OcrReviewInfo { get; set; }

        public ContentReviewTemplateUserDefineConfigureArgs()
        {
        }
        public static new ContentReviewTemplateUserDefineConfigureArgs Empty => new ContentReviewTemplateUserDefineConfigureArgs();
    }
}
