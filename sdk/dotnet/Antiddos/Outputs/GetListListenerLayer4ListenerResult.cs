// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class GetListListenerLayer4ListenerResult
    {
        public readonly int BackendPort;
        public readonly int FrontendPort;
        public readonly ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailRuleResult> InstanceDetailRules;
        public readonly ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailResult> InstanceDetails;
        public readonly string Protocol;
        public readonly ImmutableArray<Outputs.GetListListenerLayer4ListenerRealServerResult> RealServers;

        [OutputConstructor]
        private GetListListenerLayer4ListenerResult(
            int backendPort,

            int frontendPort,

            ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailRuleResult> instanceDetailRules,

            ImmutableArray<Outputs.GetListListenerLayer4ListenerInstanceDetailResult> instanceDetails,

            string protocol,

            ImmutableArray<Outputs.GetListListenerLayer4ListenerRealServerResult> realServers)
        {
            BackendPort = backendPort;
            FrontendPort = frontendPort;
            InstanceDetailRules = instanceDetailRules;
            InstanceDetails = instanceDetails;
            Protocol = protocol;
            RealServers = realServers;
        }
    }
}
