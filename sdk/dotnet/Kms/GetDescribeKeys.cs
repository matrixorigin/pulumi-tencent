// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kms
{
    public static class GetDescribeKeys
    {
        public static Task<GetDescribeKeysResult> InvokeAsync(GetDescribeKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeKeysResult>("tencentcloud:Kms/getDescribeKeys:getDescribeKeys", args ?? new GetDescribeKeysArgs(), options.WithDefaults());

        public static Output<GetDescribeKeysResult> Invoke(GetDescribeKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeKeysResult>("tencentcloud:Kms/getDescribeKeys:getDescribeKeys", args ?? new GetDescribeKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeKeysArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyIds", required: true)]
        private List<string>? _keyIds;
        public List<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new List<string>());
            set => _keyIds = value;
        }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeKeysArgs()
        {
        }
        public static new GetDescribeKeysArgs Empty => new GetDescribeKeysArgs();
    }

    public sealed class GetDescribeKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyIds", required: true)]
        private InputList<string>? _keyIds;
        public InputList<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new InputList<string>());
            set => _keyIds = value;
        }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeKeysInvokeArgs()
        {
        }
        public static new GetDescribeKeysInvokeArgs Empty => new GetDescribeKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeKeysResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> KeyIds;
        public readonly ImmutableArray<Outputs.GetDescribeKeysKeyListResult> KeyLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeKeysResult(
            string id,

            ImmutableArray<string> keyIds,

            ImmutableArray<Outputs.GetDescribeKeysKeyListResult> keyLists,

            string? resultOutputFile)
        {
            Id = id;
            KeyIds = keyIds;
            KeyLists = keyLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}