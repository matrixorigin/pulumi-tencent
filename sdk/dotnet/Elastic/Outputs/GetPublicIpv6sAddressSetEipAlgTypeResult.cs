// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elastic.Outputs
{

    [OutputType]
    public sealed class GetPublicIpv6sAddressSetEipAlgTypeResult
    {
        public readonly bool Ftp;
        public readonly bool Sip;

        [OutputConstructor]
        private GetPublicIpv6sAddressSetEipAlgTypeResult(
            bool ftp,

            bool sip)
        {
            Ftp = ftp;
            Sip = sip;
        }
    }
}