// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateVersionActionTimerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionTime")]
        public Input<string>? ActionTime { get; set; }

        [Input("externals")]
        public Input<Inputs.LaunchTemplateVersionActionTimerExternalsGetArgs>? Externals { get; set; }

        [Input("timerAction")]
        public Input<string>? TimerAction { get; set; }

        public LaunchTemplateVersionActionTimerGetArgs()
        {
        }
        public static new LaunchTemplateVersionActionTimerGetArgs Empty => new LaunchTemplateVersionActionTimerGetArgs();
    }
}