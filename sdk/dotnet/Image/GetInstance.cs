// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Image
{
    public static class GetInstance
    {
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("tencentcloud:Image/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("tencentcloud:Image/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstanceFilterArgs>? _filters;
        public List<Inputs.GetInstanceFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceFilterArgs>());
            set => _filters = value;
        }

        [Input("imageNameRegex")]
        public string? ImageNameRegex { get; set; }

        [Input("osName")]
        public string? OsName { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetInstanceFilterInputArgs>? _filters;
        public InputList<Inputs.GetInstanceFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetInstanceFilterInputArgs>());
            set => _filters = value;
        }

        [Input("imageNameRegex")]
        public Input<string>? ImageNameRegex { get; set; }

        [Input("osName")]
        public Input<string>? OsName { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageId;
        public readonly string ImageName;
        public readonly string? ImageNameRegex;
        public readonly string? OsName;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<Outputs.GetInstanceFilterResult> filters,

            string id,

            string imageId,

            string imageName,

            string? imageNameRegex,

            string? osName,

            string? resultOutputFile)
        {
            Filters = filters;
            Id = id;
            ImageId = imageId;
            ImageName = imageName;
            ImageNameRegex = imageNameRegex;
            OsName = osName;
            ResultOutputFile = resultOutputFile;
        }
    }
}