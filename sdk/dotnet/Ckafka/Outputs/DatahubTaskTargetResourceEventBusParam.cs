// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskTargetResourceEventBusParam
    {
        public readonly string? FunctionName;
        public readonly string? Namespace;
        public readonly string? Qualifier;
        public readonly string Resource;
        public readonly bool SelfBuilt;
        public readonly string Type;

        [OutputConstructor]
        private DatahubTaskTargetResourceEventBusParam(
            string? functionName,

            string? @namespace,

            string? qualifier,

            string resource,

            bool selfBuilt,

            string type)
        {
            FunctionName = functionName;
            Namespace = @namespace;
            Qualifier = qualifier;
            Resource = resource;
            SelfBuilt = selfBuilt;
            Type = type;
        }
    }
}