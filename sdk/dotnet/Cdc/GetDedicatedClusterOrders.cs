// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdc
{
    public static class GetDedicatedClusterOrders
    {
        public static Task<GetDedicatedClusterOrdersResult> InvokeAsync(GetDedicatedClusterOrdersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDedicatedClusterOrdersResult>("tencentcloud:Cdc/getDedicatedClusterOrders:getDedicatedClusterOrders", args ?? new GetDedicatedClusterOrdersArgs(), options.WithDefaults());

        public static Output<GetDedicatedClusterOrdersResult> Invoke(GetDedicatedClusterOrdersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDedicatedClusterOrdersResult>("tencentcloud:Cdc/getDedicatedClusterOrders:getDedicatedClusterOrders", args ?? new GetDedicatedClusterOrdersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDedicatedClusterOrdersArgs : global::Pulumi.InvokeArgs
    {
        [Input("actionType")]
        public string? ActionType { get; set; }

        [Input("dedicatedClusterIds")]
        private List<string>? _dedicatedClusterIds;
        public List<string> DedicatedClusterIds
        {
            get => _dedicatedClusterIds ?? (_dedicatedClusterIds = new List<string>());
            set => _dedicatedClusterIds = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetDedicatedClusterOrdersArgs()
        {
        }
        public static new GetDedicatedClusterOrdersArgs Empty => new GetDedicatedClusterOrdersArgs();
    }

    public sealed class GetDedicatedClusterOrdersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("actionType")]
        public Input<string>? ActionType { get; set; }

        [Input("dedicatedClusterIds")]
        private InputList<string>? _dedicatedClusterIds;
        public InputList<string> DedicatedClusterIds
        {
            get => _dedicatedClusterIds ?? (_dedicatedClusterIds = new InputList<string>());
            set => _dedicatedClusterIds = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetDedicatedClusterOrdersInvokeArgs()
        {
        }
        public static new GetDedicatedClusterOrdersInvokeArgs Empty => new GetDedicatedClusterOrdersInvokeArgs();
    }


    [OutputType]
    public sealed class GetDedicatedClusterOrdersResult
    {
        public readonly string? ActionType;
        public readonly ImmutableArray<string> DedicatedClusterIds;
        public readonly ImmutableArray<Outputs.GetDedicatedClusterOrdersDedicatedClusterOrderSetResult> DedicatedClusterOrderSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string? Status;

        [OutputConstructor]
        private GetDedicatedClusterOrdersResult(
            string? actionType,

            ImmutableArray<string> dedicatedClusterIds,

            ImmutableArray<Outputs.GetDedicatedClusterOrdersDedicatedClusterOrderSetResult> dedicatedClusterOrderSets,

            string id,

            string? resultOutputFile,

            string? status)
        {
            ActionType = actionType;
            DedicatedClusterIds = dedicatedClusterIds;
            DedicatedClusterOrderSets = dedicatedClusterOrderSets;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Status = status;
        }
    }
}