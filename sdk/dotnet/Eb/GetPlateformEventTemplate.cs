// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eb
{
    public static class GetPlateformEventTemplate
    {
        public static Task<GetPlateformEventTemplateResult> InvokeAsync(GetPlateformEventTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlateformEventTemplateResult>("tencentcloud:Eb/getPlateformEventTemplate:getPlateformEventTemplate", args ?? new GetPlateformEventTemplateArgs(), options.WithDefaults());

        public static Output<GetPlateformEventTemplateResult> Invoke(GetPlateformEventTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlateformEventTemplateResult>("tencentcloud:Eb/getPlateformEventTemplate:getPlateformEventTemplate", args ?? new GetPlateformEventTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlateformEventTemplateArgs : global::Pulumi.InvokeArgs
    {
        [Input("eventType", required: true)]
        public string EventType { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetPlateformEventTemplateArgs()
        {
        }
        public static new GetPlateformEventTemplateArgs Empty => new GetPlateformEventTemplateArgs();
    }

    public sealed class GetPlateformEventTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("eventType", required: true)]
        public Input<string> EventType { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetPlateformEventTemplateInvokeArgs()
        {
        }
        public static new GetPlateformEventTemplateInvokeArgs Empty => new GetPlateformEventTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlateformEventTemplateResult
    {
        public readonly string EventTemplate;
        public readonly string EventType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetPlateformEventTemplateResult(
            string eventTemplate,

            string eventType,

            string id,

            string? resultOutputFile)
        {
            EventTemplate = eventTemplate;
            EventType = eventType;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}