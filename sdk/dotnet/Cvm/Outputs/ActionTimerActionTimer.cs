// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class ActionTimerActionTimer
    {
        public readonly string? ActionTime;
        public readonly string? TimerAction;

        [OutputConstructor]
        private ActionTimerActionTimer(
            string? actionTime,

            string? timerAction)
        {
            ActionTime = actionTime;
            TimerAction = timerAction;
        }
    }
}