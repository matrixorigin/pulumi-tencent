// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse
{
    public static class GetInstanceBlueprint
    {
        public static Task<GetInstanceBlueprintResult> InvokeAsync(GetInstanceBlueprintArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceBlueprintResult>("tencentcloud:Lighthouse/getInstanceBlueprint:getInstanceBlueprint", args ?? new GetInstanceBlueprintArgs(), options.WithDefaults());

        public static Output<GetInstanceBlueprintResult> Invoke(GetInstanceBlueprintInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceBlueprintResult>("tencentcloud:Lighthouse/getInstanceBlueprint:getInstanceBlueprint", args ?? new GetInstanceBlueprintInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceBlueprintArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceIds", required: true)]
        private List<string>? _instanceIds;
        public List<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new List<string>());
            set => _instanceIds = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceBlueprintArgs()
        {
        }
        public static new GetInstanceBlueprintArgs Empty => new GetInstanceBlueprintArgs();
    }

    public sealed class GetInstanceBlueprintInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceBlueprintInvokeArgs()
        {
        }
        public static new GetInstanceBlueprintInvokeArgs Empty => new GetInstanceBlueprintInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceBlueprintResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceBlueprintBlueprintInstanceSetResult> BlueprintInstanceSets;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> InstanceIds;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceBlueprintResult(
            ImmutableArray<Outputs.GetInstanceBlueprintBlueprintInstanceSetResult> blueprintInstanceSets,

            string id,

            ImmutableArray<string> instanceIds,

            string? resultOutputFile)
        {
            BlueprintInstanceSets = blueprintInstanceSets;
            Id = id;
            InstanceIds = instanceIds;
            ResultOutputFile = resultOutputFile;
        }
    }
}