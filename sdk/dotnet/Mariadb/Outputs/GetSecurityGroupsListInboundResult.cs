// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mariadb.Outputs
{

    [OutputType]
    public sealed class GetSecurityGroupsListInboundResult
    {
        public readonly string Action;
        public readonly string CidrIp;
        public readonly string IpProtocol;
        public readonly string PortRange;

        [OutputConstructor]
        private GetSecurityGroupsListInboundResult(
            string action,

            string cidrIp,

            string ipProtocol,

            string portRange)
        {
            Action = action;
            CidrIp = cidrIp;
            IpProtocol = ipProtocol;
            PortRange = portRange;
        }
    }
}