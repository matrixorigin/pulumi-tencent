// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class AiAnalysisTemplateClassificationConfigureGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public AiAnalysisTemplateClassificationConfigureGetArgs()
        {
        }
        public static new AiAnalysisTemplateClassificationConfigureGetArgs Empty => new AiAnalysisTemplateClassificationConfigureGetArgs();
    }
}