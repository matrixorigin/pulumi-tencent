// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleActivityActivityParaTranscodeTaskObjectNumberFormatGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("increment")]
        public Input<int>? Increment { get; set; }

        [Input("initialValue")]
        public Input<int>? InitialValue { get; set; }

        [Input("minLength")]
        public Input<int>? MinLength { get; set; }

        [Input("placeHolder")]
        public Input<string>? PlaceHolder { get; set; }

        public ScheduleActivityActivityParaTranscodeTaskObjectNumberFormatGetArgs()
        {
        }
        public static new ScheduleActivityActivityParaTranscodeTaskObjectNumberFormatGetArgs Empty => new ScheduleActivityActivityParaTranscodeTaskObjectNumberFormatGetArgs();
    }
}
