// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb
{
    public static class GetClusters
    {
        public static Task<GetClustersResult> InvokeAsync(GetClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClustersResult>("tencentcloud:Cynosdb/getClusters:getClusters", args ?? new GetClustersArgs(), options.WithDefaults());

        public static Output<GetClustersResult> Invoke(GetClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClustersResult>("tencentcloud:Cynosdb/getClusters:getClusters", args ?? new GetClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("clusterName")]
        public string? ClusterName { get; set; }

        [Input("dbType")]
        public string? DbType { get; set; }

        [Input("projectId")]
        public int? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetClustersArgs()
        {
        }
        public static new GetClustersArgs Empty => new GetClustersArgs();
    }

    public sealed class GetClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetClustersInvokeArgs()
        {
        }
        public static new GetClustersInvokeArgs Empty => new GetClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetClustersResult
    {
        public readonly string? ClusterId;
        public readonly ImmutableArray<Outputs.GetClustersClusterListResult> ClusterLists;
        public readonly string? ClusterName;
        public readonly string? DbType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetClustersResult(
            string? clusterId,

            ImmutableArray<Outputs.GetClustersClusterListResult> clusterLists,

            string? clusterName,

            string? dbType,

            string id,

            int? projectId,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            ClusterLists = clusterLists;
            ClusterName = clusterName;
            DbType = dbType;
            Id = id;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}