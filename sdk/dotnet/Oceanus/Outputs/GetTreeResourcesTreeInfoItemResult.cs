// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class GetTreeResourcesTreeInfoItemResult
    {
        public readonly string FileName;
        public readonly string FolderId;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetTreeResourcesTreeInfoItemRefJobStatusCountSetResult> RefJobStatusCountSets;
        public readonly string Remark;
        public readonly string ResourceId;
        public readonly int ResourceType;

        [OutputConstructor]
        private GetTreeResourcesTreeInfoItemResult(
            string fileName,

            string folderId,

            string name,

            ImmutableArray<Outputs.GetTreeResourcesTreeInfoItemRefJobStatusCountSetResult> refJobStatusCountSets,

            string remark,

            string resourceId,

            int resourceType)
        {
            FileName = fileName;
            FolderId = folderId;
            Name = name;
            RefJobStatusCountSets = refJobStatusCountSets;
            Remark = remark;
            ResourceId = resourceId;
            ResourceType = resourceType;
        }
    }
}