// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eb.Inputs
{

    public sealed class EventTransformTransformationExtractionTextParamsArgs : global::Pulumi.ResourceArgs
    {
        [Input("regex")]
        public Input<string>? Regex { get; set; }

        [Input("separator")]
        public Input<string>? Separator { get; set; }

        public EventTransformTransformationExtractionTextParamsArgs()
        {
        }
        public static new EventTransformTransformationExtractionTextParamsArgs Empty => new EventTransformTransformationExtractionTextParamsArgs();
    }
}