// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Chdfs
{
    public static class GetFileSystems
    {
        public static Task<GetFileSystemsResult> InvokeAsync(GetFileSystemsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemsResult>("tencentcloud:Chdfs/getFileSystems:getFileSystems", args ?? new GetFileSystemsArgs(), options.WithDefaults());

        public static Output<GetFileSystemsResult> Invoke(GetFileSystemsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileSystemsResult>("tencentcloud:Chdfs/getFileSystems:getFileSystems", args ?? new GetFileSystemsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileSystemsArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetFileSystemsArgs()
        {
        }
        public static new GetFileSystemsArgs Empty => new GetFileSystemsArgs();
    }

    public sealed class GetFileSystemsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetFileSystemsInvokeArgs()
        {
        }
        public static new GetFileSystemsInvokeArgs Empty => new GetFileSystemsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileSystemsResult
    {
        public readonly ImmutableArray<Outputs.GetFileSystemsFileSystemResult> ChdfsFileSystems;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetFileSystemsResult(
            ImmutableArray<Outputs.GetFileSystemsFileSystemResult> fileSystems,

            string id,

            string? resultOutputFile)
        {
            ChdfsFileSystems = fileSystems;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}