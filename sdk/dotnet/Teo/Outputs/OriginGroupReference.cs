// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class OriginGroupReference
    {
        public readonly string? InstanceId;
        public readonly string? InstanceName;
        public readonly string? InstanceType;

        [OutputConstructor]
        private OriginGroupReference(
            string? instanceId,

            string? instanceName,

            string? instanceType)
        {
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceType = instanceType;
        }
    }
}
