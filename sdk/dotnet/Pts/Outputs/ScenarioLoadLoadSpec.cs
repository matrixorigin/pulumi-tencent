// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts.Outputs
{

    [OutputType]
    public sealed class ScenarioLoadLoadSpec
    {
        public readonly Outputs.ScenarioLoadLoadSpecConcurrency? Concurrency;
        public readonly Outputs.ScenarioLoadLoadSpecRequestsPerSecond? RequestsPerSecond;
        public readonly Outputs.ScenarioLoadLoadSpecScriptOrigin? ScriptOrigin;

        [OutputConstructor]
        private ScenarioLoadLoadSpec(
            Outputs.ScenarioLoadLoadSpecConcurrency? concurrency,

            Outputs.ScenarioLoadLoadSpecRequestsPerSecond? requestsPerSecond,

            Outputs.ScenarioLoadLoadSpecScriptOrigin? scriptOrigin)
        {
            Concurrency = concurrency;
            RequestsPerSecond = requestsPerSecond;
            ScriptOrigin = scriptOrigin;
        }
    }
}