// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetS3InputInfoResult
    {
        public readonly string S3Bucket;
        public readonly string S3Object;
        public readonly string S3Region;
        public readonly string S3SecretId;
        public readonly string S3SecretKey;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetS3InputInfoResult(
            string s3Bucket,

            string s3Object,

            string s3Region,

            string s3SecretId,

            string s3SecretKey)
        {
            S3Bucket = s3Bucket;
            S3Object = s3Object;
            S3Region = s3Region;
            S3SecretId = s3SecretId;
            S3SecretKey = s3SecretKey;
        }
    }
}
