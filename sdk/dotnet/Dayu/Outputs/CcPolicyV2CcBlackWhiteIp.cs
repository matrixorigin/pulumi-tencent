// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class CcPolicyV2CcBlackWhiteIp
    {
        public readonly string BlackWhiteIp;
        public readonly string? CreateTime;
        public readonly string Domain;
        public readonly string? ModifyTime;
        public readonly string Protocol;
        public readonly string Type;

        [OutputConstructor]
        private CcPolicyV2CcBlackWhiteIp(
            string blackWhiteIp,

            string? createTime,

            string domain,

            string? modifyTime,

            string protocol,

            string type)
        {
            BlackWhiteIp = blackWhiteIp;
            CreateTime = createTime;
            Domain = domain;
            ModifyTime = modifyTime;
            Protocol = protocol;
            Type = type;
        }
    }
}