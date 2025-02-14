// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    public static class GetCountryAreaMapping
    {
        public static Task<GetCountryAreaMappingResult> InvokeAsync(GetCountryAreaMappingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCountryAreaMappingResult>("tencentcloud:Gaap/getCountryAreaMapping:getCountryAreaMapping", args ?? new GetCountryAreaMappingArgs(), options.WithDefaults());

        public static Output<GetCountryAreaMappingResult> Invoke(GetCountryAreaMappingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCountryAreaMappingResult>("tencentcloud:Gaap/getCountryAreaMapping:getCountryAreaMapping", args ?? new GetCountryAreaMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCountryAreaMappingArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetCountryAreaMappingArgs()
        {
        }
        public static new GetCountryAreaMappingArgs Empty => new GetCountryAreaMappingArgs();
    }

    public sealed class GetCountryAreaMappingInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetCountryAreaMappingInvokeArgs()
        {
        }
        public static new GetCountryAreaMappingInvokeArgs Empty => new GetCountryAreaMappingInvokeArgs();
    }


    [OutputType]
    public sealed class GetCountryAreaMappingResult
    {
        public readonly ImmutableArray<Outputs.GetCountryAreaMappingCountryAreaMappingListResult> CountryAreaMappingLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetCountryAreaMappingResult(
            ImmutableArray<Outputs.GetCountryAreaMappingCountryAreaMappingListResult> countryAreaMappingLists,

            string id,

            string? resultOutputFile)
        {
            CountryAreaMappingLists = countryAreaMappingLists;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
