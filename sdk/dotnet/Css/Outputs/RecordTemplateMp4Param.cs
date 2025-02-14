// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css.Outputs
{

    [OutputType]
    public sealed class RecordTemplateMp4Param
    {
        public readonly int? ClassId;
        public readonly int? Enable;
        public readonly string? Procedure;
        public readonly int? RecordInterval;
        public readonly string? StorageMode;
        public readonly int? StorageTime;
        public readonly string? VodFileName;
        public readonly int? VodSubAppId;

        [OutputConstructor]
        private RecordTemplateMp4Param(
            int? classId,

            int? enable,

            string? procedure,

            int? recordInterval,

            string? storageMode,

            int? storageTime,

            string? vodFileName,

            int? vodSubAppId)
        {
            ClassId = classId;
            Enable = enable;
            Procedure = procedure;
            RecordInterval = recordInterval;
            StorageMode = storageMode;
            StorageTime = storageTime;
            VodFileName = vodFileName;
            VodSubAppId = vodSubAppId;
        }
    }
}
