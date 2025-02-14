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
    public sealed class GetDescribeCertificateResultSubmittedDataResult
    {
        public readonly string AdminEmail;
        public readonly string AdminFirstName;
        public readonly string AdminLastName;
        public readonly string AdminPhoneNum;
        public readonly string AdminPosition;
        public readonly string CertificateDomain;
        public readonly string ContactEmail;
        public readonly string ContactFirstName;
        public readonly string ContactLastName;
        public readonly string ContactNumber;
        public readonly string ContactPosition;
        public readonly string CsrContent;
        public readonly string CsrType;
        public readonly ImmutableArray<string> DomainLists;
        public readonly string KeyPassword;
        public readonly string OrganizationAddress;
        public readonly string OrganizationCity;
        public readonly string OrganizationCountry;
        public readonly string OrganizationDivision;
        public readonly string OrganizationName;
        public readonly string OrganizationRegion;
        public readonly string PhoneAreaCode;
        public readonly string PhoneNumber;
        public readonly string PostalCode;
        public readonly string VerifyType;

        [OutputConstructor]
        private GetDescribeCertificateResultSubmittedDataResult(
            string adminEmail,

            string adminFirstName,

            string adminLastName,

            string adminPhoneNum,

            string adminPosition,

            string certificateDomain,

            string contactEmail,

            string contactFirstName,

            string contactLastName,

            string contactNumber,

            string contactPosition,

            string csrContent,

            string csrType,

            ImmutableArray<string> domainLists,

            string keyPassword,

            string organizationAddress,

            string organizationCity,

            string organizationCountry,

            string organizationDivision,

            string organizationName,

            string organizationRegion,

            string phoneAreaCode,

            string phoneNumber,

            string postalCode,

            string verifyType)
        {
            AdminEmail = adminEmail;
            AdminFirstName = adminFirstName;
            AdminLastName = adminLastName;
            AdminPhoneNum = adminPhoneNum;
            AdminPosition = adminPosition;
            CertificateDomain = certificateDomain;
            ContactEmail = contactEmail;
            ContactFirstName = contactFirstName;
            ContactLastName = contactLastName;
            ContactNumber = contactNumber;
            ContactPosition = contactPosition;
            CsrContent = csrContent;
            CsrType = csrType;
            DomainLists = domainLists;
            KeyPassword = keyPassword;
            OrganizationAddress = organizationAddress;
            OrganizationCity = organizationCity;
            OrganizationCountry = organizationCountry;
            OrganizationDivision = organizationDivision;
            OrganizationName = organizationName;
            OrganizationRegion = organizationRegion;
            PhoneAreaCode = phoneAreaCode;
            PhoneNumber = phoneNumber;
            PostalCode = postalCode;
            VerifyType = verifyType;
        }
    }
}
