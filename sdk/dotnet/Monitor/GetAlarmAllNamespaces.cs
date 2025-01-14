// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    public static class GetAlarmAllNamespaces
    {
        public static Task<GetAlarmAllNamespacesResult> InvokeAsync(GetAlarmAllNamespacesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlarmAllNamespacesResult>("tencentcloud:Monitor/getAlarmAllNamespaces:getAlarmAllNamespaces", args ?? new GetAlarmAllNamespacesArgs(), options.WithDefaults());

        public static Output<GetAlarmAllNamespacesResult> Invoke(GetAlarmAllNamespacesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmAllNamespacesResult>("tencentcloud:Monitor/getAlarmAllNamespaces:getAlarmAllNamespaces", args ?? new GetAlarmAllNamespacesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlarmAllNamespacesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("module", required: true)]
        public string Module { get; set; } = null!;

        [Input("monitorTypes")]
        private List<string>? _monitorTypes;
        public List<string> MonitorTypes
        {
            get => _monitorTypes ?? (_monitorTypes = new List<string>());
            set => _monitorTypes = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("sceneType", required: true)]
        public string SceneType { get; set; } = null!;

        public GetAlarmAllNamespacesArgs()
        {
        }
        public static new GetAlarmAllNamespacesArgs Empty => new GetAlarmAllNamespacesArgs();
    }

    public sealed class GetAlarmAllNamespacesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("module", required: true)]
        public Input<string> Module { get; set; } = null!;

        [Input("monitorTypes")]
        private InputList<string>? _monitorTypes;
        public InputList<string> MonitorTypes
        {
            get => _monitorTypes ?? (_monitorTypes = new InputList<string>());
            set => _monitorTypes = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("sceneType", required: true)]
        public Input<string> SceneType { get; set; } = null!;

        public GetAlarmAllNamespacesInvokeArgs()
        {
        }
        public static new GetAlarmAllNamespacesInvokeArgs Empty => new GetAlarmAllNamespacesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlarmAllNamespacesResult
    {
        public readonly ImmutableArray<Outputs.GetAlarmAllNamespacesCommonNamespaceResult> CommonNamespaces;
        public readonly ImmutableArray<Outputs.GetAlarmAllNamespacesCustomNamespacesNewsResult> CustomNamespacesNews;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string Module;
        public readonly ImmutableArray<string> MonitorTypes;
        public readonly ImmutableArray<Outputs.GetAlarmAllNamespacesQceNamespacesNewsResult> QceNamespacesNews;
        public readonly string? ResultOutputFile;
        public readonly string SceneType;

        [OutputConstructor]
        private GetAlarmAllNamespacesResult(
            ImmutableArray<Outputs.GetAlarmAllNamespacesCommonNamespaceResult> commonNamespaces,

            ImmutableArray<Outputs.GetAlarmAllNamespacesCustomNamespacesNewsResult> customNamespacesNews,

            string id,

            ImmutableArray<string> ids,

            string module,

            ImmutableArray<string> monitorTypes,

            ImmutableArray<Outputs.GetAlarmAllNamespacesQceNamespacesNewsResult> qceNamespacesNews,

            string? resultOutputFile,

            string sceneType)
        {
            CommonNamespaces = commonNamespaces;
            CustomNamespacesNews = customNamespacesNews;
            Id = id;
            Ids = ids;
            Module = module;
            MonitorTypes = monitorTypes;
            QceNamespacesNews = qceNamespacesNews;
            ResultOutputFile = resultOutputFile;
            SceneType = sceneType;
        }
    }
}
