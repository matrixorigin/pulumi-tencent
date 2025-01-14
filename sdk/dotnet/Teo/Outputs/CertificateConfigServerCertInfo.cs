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
    public sealed class CertificateConfigServerCertInfo
    {
        public readonly string? Alias;
        public readonly string CertId;
        public readonly string? CommonName;
        public readonly string? DeployTime;
        public readonly string? ExpireTime;
        public readonly string? SignAlgo;
        public readonly string? Type;

        [OutputConstructor]
        private CertificateConfigServerCertInfo(
            string? alias,

            string certId,

            string? commonName,

            string? deployTime,

            string? expireTime,

            string? signAlgo,

            string? type)
        {
            Alias = alias;
            CertId = certId;
            CommonName = commonName;
            DeployTime = deployTime;
            ExpireTime = expireTime;
            SignAlgo = signAlgo;
            Type = type;
        }
    }
}
