// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class CheckCertificateDomainVerificationOperationVerificationResult
    {
        public readonly int? CaCheck;
        public readonly ImmutableArray<string> CheckValues;
        public readonly string? Domain;
        public readonly bool? Frequently;
        public readonly bool? Issued;
        public readonly int? LocalCheck;
        public readonly string? LocalCheckFailReason;
        public readonly string? VerifyType;

        [OutputConstructor]
        private CheckCertificateDomainVerificationOperationVerificationResult(
            int? caCheck,

            ImmutableArray<string> checkValues,

            string? domain,

            bool? frequently,

            bool? issued,

            int? localCheck,

            string? localCheckFailReason,

            string? verifyType)
        {
            CaCheck = caCheck;
            CheckValues = checkValues;
            Domain = domain;
            Frequently = frequently;
            Issued = issued;
            LocalCheck = localCheck;
            LocalCheckFailReason = localCheckFailReason;
            VerifyType = verifyType;
        }
    }
}