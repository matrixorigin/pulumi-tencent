// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tat.Outputs
{

    [OutputType]
    public sealed class GetInvocationTaskInvocationTaskSetTaskResultResult
    {
        public readonly int Dropped;
        public readonly string ExecEndTime;
        public readonly string ExecStartTime;
        public readonly int ExitCode;
        public readonly string Output;
        public readonly string OutputUploadCosErrorInfo;
        public readonly string OutputUrl;

        [OutputConstructor]
        private GetInvocationTaskInvocationTaskSetTaskResultResult(
            int dropped,

            string execEndTime,

            string execStartTime,

            int exitCode,

            string output,

            string outputUploadCosErrorInfo,

            string outputUrl)
        {
            Dropped = dropped;
            ExecEndTime = execEndTime;
            ExecStartTime = execStartTime;
            ExitCode = exitCode;
            Output = output;
            OutputUploadCosErrorInfo = outputUploadCosErrorInfo;
            OutputUrl = outputUrl;
        }
    }
}