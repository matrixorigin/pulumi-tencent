// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class GetJobEventsEventResult
    {
        public readonly string Description;
        public readonly string Message;
        public readonly int RunningOrderId;
        public readonly string SolutionLink;
        public readonly int Timestamp;
        public readonly string Type;

        [OutputConstructor]
        private GetJobEventsEventResult(
            string description,

            string message,

            int runningOrderId,

            string solutionLink,

            int timestamp,

            string type)
        {
            Description = description;
            Message = message;
            RunningOrderId = runningOrderId;
            SolutionLink = solutionLink;
            Timestamp = timestamp;
            Type = type;
        }
    }
}
