// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Deprecatedcontainer.Outputs
{

    [OutputType]
    public sealed class GetClustersClusterResult
    {
        public readonly string ClusterId;
        public readonly string ClusterName;
        public readonly string Description;
        public readonly string KubernetesVersion;
        public readonly int NodesNum;
        public readonly string NodesStatus;
        public readonly string SecurityCertificationAuthority;
        public readonly string SecurityClusterExternalEndpoint;
        public readonly string SecurityPassword;
        public readonly string SecurityUsername;
        public readonly int TotalCpu;
        public readonly int TotalMem;

        [OutputConstructor]
        private GetClustersClusterResult(
            string clusterId,

            string clusterName,

            string description,

            string kubernetesVersion,

            int nodesNum,

            string nodesStatus,

            string securityCertificationAuthority,

            string securityClusterExternalEndpoint,

            string securityPassword,

            string securityUsername,

            int totalCpu,

            int totalMem)
        {
            ClusterId = clusterId;
            ClusterName = clusterName;
            Description = description;
            KubernetesVersion = kubernetesVersion;
            NodesNum = nodesNum;
            NodesStatus = nodesStatus;
            SecurityCertificationAuthority = securityCertificationAuthority;
            SecurityClusterExternalEndpoint = securityClusterExternalEndpoint;
            SecurityPassword = securityPassword;
            SecurityUsername = securityUsername;
            TotalCpu = totalCpu;
            TotalMem = totalMem;
        }
    }
}