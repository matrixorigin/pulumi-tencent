// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Rum
{
    public static class GetSign
    {
        public static Task<GetSignResult> InvokeAsync(GetSignArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSignResult>("tencentcloud:Rum/getSign:getSign", args ?? new GetSignArgs(), options.WithDefaults());

        public static Output<GetSignResult> Invoke(GetSignInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSignResult>("tencentcloud:Rum/getSign:getSign", args ?? new GetSignInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSignArgs : global::Pulumi.InvokeArgs
    {
        [Input("fileType")]
        public int? FileType { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("timeout")]
        public int? Timeout { get; set; }

        public GetSignArgs()
        {
        }
        public static new GetSignArgs Empty => new GetSignArgs();
    }

    public sealed class GetSignInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("fileType")]
        public Input<int>? FileType { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public GetSignInvokeArgs()
        {
        }
        public static new GetSignInvokeArgs Empty => new GetSignInvokeArgs();
    }


    [OutputType]
    public sealed class GetSignResult
    {
        public readonly int ExpiredTime;
        public readonly int? FileType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string SecretId;
        public readonly string SecretKey;
        public readonly string SessionToken;
        public readonly int StartTime;
        public readonly int? Timeout;

        [OutputConstructor]
        private GetSignResult(
            int expiredTime,

            int? fileType,

            string id,

            string? resultOutputFile,

            string secretId,

            string secretKey,

            string sessionToken,

            int startTime,

            int? timeout)
        {
            ExpiredTime = expiredTime;
            FileType = fileType;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SecretId = secretId;
            SecretKey = secretKey;
            SessionToken = sessionToken;
            StartTime = startTime;
            Timeout = timeout;
        }
    }
}