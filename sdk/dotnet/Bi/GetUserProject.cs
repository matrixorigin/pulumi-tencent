// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Bi
{
    public static class GetUserProject
    {
        public static Task<GetUserProjectResult> InvokeAsync(GetUserProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserProjectResult>("tencentcloud:Bi/getUserProject:getUserProject", args ?? new GetUserProjectArgs(), options.WithDefaults());

        public static Output<GetUserProjectResult> Invoke(GetUserProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserProjectResult>("tencentcloud:Bi/getUserProject:getUserProject", args ?? new GetUserProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserProjectArgs : global::Pulumi.InvokeArgs
    {
        [Input("allPage")]
        public bool? AllPage { get; set; }

        [Input("projectId")]
        public int? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetUserProjectArgs()
        {
        }
        public static new GetUserProjectArgs Empty => new GetUserProjectArgs();
    }

    public sealed class GetUserProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("allPage")]
        public Input<bool>? AllPage { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetUserProjectInvokeArgs()
        {
        }
        public static new GetUserProjectInvokeArgs Empty => new GetUserProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserProjectResult
    {
        public readonly bool? AllPage;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetUserProjectListResult> Lists;
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetUserProjectResult(
            bool? allPage,

            string id,

            ImmutableArray<Outputs.GetUserProjectListResult> lists,

            int? projectId,

            string? resultOutputFile)
        {
            AllPage = allPage;
            Id = id;
            Lists = lists;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
