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
    public sealed class GetCertificatesCertificateResult
    {
        public readonly string BeginTime;
        public readonly string Cert;
        public readonly string CreateTime;
        public readonly string Domain;
        public readonly ImmutableArray<Outputs.GetCertificatesCertificateDvAuthResult> DvAuths;
        public readonly string EndTime;
        public readonly string Id;
        public readonly string Key;
        public readonly string Name;
        public readonly string OrderId;
        public readonly string OwnerUin;
        public readonly string ProductZhName;
        public readonly int ProjectId;
        public readonly int Status;
        public readonly ImmutableArray<string> SubjectNames;
        public readonly string Type;
        public readonly string ValidityPeriod;

        [OutputConstructor]
        private GetCertificatesCertificateResult(
            string beginTime,

            string cert,

            string createTime,

            string domain,

            ImmutableArray<Outputs.GetCertificatesCertificateDvAuthResult> dvAuths,

            string endTime,

            string id,

            string key,

            string name,

            string orderId,

            string ownerUin,

            string productZhName,

            int projectId,

            int status,

            ImmutableArray<string> subjectNames,

            string type,

            string validityPeriod)
        {
            BeginTime = beginTime;
            Cert = cert;
            CreateTime = createTime;
            Domain = domain;
            DvAuths = dvAuths;
            EndTime = endTime;
            Id = id;
            Key = key;
            Name = name;
            OrderId = orderId;
            OwnerUin = ownerUin;
            ProductZhName = productZhName;
            ProjectId = projectId;
            Status = status;
            SubjectNames = subjectNames;
            Type = type;
            ValidityPeriod = validityPeriod;
        }
    }
}
