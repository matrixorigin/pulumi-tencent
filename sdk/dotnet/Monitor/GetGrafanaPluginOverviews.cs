// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    public static class GetGrafanaPluginOverviews
    {
        public static Task<GetGrafanaPluginOverviewsResult> InvokeAsync(GetGrafanaPluginOverviewsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGrafanaPluginOverviewsResult>("tencentcloud:Monitor/getGrafanaPluginOverviews:getGrafanaPluginOverviews", args ?? new GetGrafanaPluginOverviewsArgs(), options.WithDefaults());

        public static Output<GetGrafanaPluginOverviewsResult> Invoke(GetGrafanaPluginOverviewsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGrafanaPluginOverviewsResult>("tencentcloud:Monitor/getGrafanaPluginOverviews:getGrafanaPluginOverviews", args ?? new GetGrafanaPluginOverviewsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGrafanaPluginOverviewsArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetGrafanaPluginOverviewsArgs()
        {
        }
        public static new GetGrafanaPluginOverviewsArgs Empty => new GetGrafanaPluginOverviewsArgs();
    }

    public sealed class GetGrafanaPluginOverviewsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetGrafanaPluginOverviewsInvokeArgs()
        {
        }
        public static new GetGrafanaPluginOverviewsInvokeArgs Empty => new GetGrafanaPluginOverviewsInvokeArgs();
    }


    [OutputType]
    public sealed class GetGrafanaPluginOverviewsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetGrafanaPluginOverviewsPluginSetResult> PluginSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetGrafanaPluginOverviewsResult(
            string id,

            ImmutableArray<Outputs.GetGrafanaPluginOverviewsPluginSetResult> pluginSets,

            string? resultOutputFile)
        {
            Id = id;
            PluginSets = pluginSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}