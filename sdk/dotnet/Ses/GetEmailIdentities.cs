// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ses
{
    public static class GetEmailIdentities
    {
        public static Task<GetEmailIdentitiesResult> InvokeAsync(GetEmailIdentitiesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEmailIdentitiesResult>("tencentcloud:Ses/getEmailIdentities:getEmailIdentities", args ?? new GetEmailIdentitiesArgs(), options.WithDefaults());

        public static Output<GetEmailIdentitiesResult> Invoke(GetEmailIdentitiesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEmailIdentitiesResult>("tencentcloud:Ses/getEmailIdentities:getEmailIdentities", args ?? new GetEmailIdentitiesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEmailIdentitiesArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetEmailIdentitiesArgs()
        {
        }
        public static new GetEmailIdentitiesArgs Empty => new GetEmailIdentitiesArgs();
    }

    public sealed class GetEmailIdentitiesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetEmailIdentitiesInvokeArgs()
        {
        }
        public static new GetEmailIdentitiesInvokeArgs Empty => new GetEmailIdentitiesInvokeArgs();
    }


    [OutputType]
    public sealed class GetEmailIdentitiesResult
    {
        public readonly ImmutableArray<Outputs.GetEmailIdentitiesEmailIdentityResult> SesEmailIdentities;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int MaxDailyQuota;
        public readonly int MaxReputationLevel;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetEmailIdentitiesResult(
            ImmutableArray<Outputs.GetEmailIdentitiesEmailIdentityResult> emailIdentities,

            string id,

            int maxDailyQuota,

            int maxReputationLevel,

            string? resultOutputFile)
        {
            SesEmailIdentities = emailIdentities;
            Id = id;
            MaxDailyQuota = maxDailyQuota;
            MaxReputationLevel = maxReputationLevel;
            ResultOutputFile = resultOutputFile;
        }
    }
}
