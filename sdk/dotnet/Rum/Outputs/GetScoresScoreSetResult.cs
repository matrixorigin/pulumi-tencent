// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Rum.Outputs
{

    [OutputType]
    public sealed class GetScoresScoreSetResult
    {
        public readonly string ApiDuration;
        public readonly string ApiFail;
        public readonly string ApiNum;
        public readonly string CreateTime;
        public readonly string PageDuration;
        public readonly string PageError;
        public readonly string PagePv;
        public readonly string PageUv;
        public readonly int ProjectId;
        public readonly int RecordNum;
        public readonly string Score;
        public readonly string StaticDuration;
        public readonly string StaticFail;
        public readonly string StaticNum;

        [OutputConstructor]
        private GetScoresScoreSetResult(
            string apiDuration,

            string apiFail,

            string apiNum,

            string createTime,

            string pageDuration,

            string pageError,

            string pagePv,

            string pageUv,

            int projectId,

            int recordNum,

            string score,

            string staticDuration,

            string staticFail,

            string staticNum)
        {
            ApiDuration = apiDuration;
            ApiFail = apiFail;
            ApiNum = apiNum;
            CreateTime = createTime;
            PageDuration = pageDuration;
            PageError = pageError;
            PagePv = pagePv;
            PageUv = pageUv;
            ProjectId = projectId;
            RecordNum = recordNum;
            Score = score;
            StaticDuration = staticDuration;
            StaticFail = staticFail;
            StaticNum = staticNum;
        }
    }
}