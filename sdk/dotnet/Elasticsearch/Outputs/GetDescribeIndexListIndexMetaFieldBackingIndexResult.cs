// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetDescribeIndexListIndexMetaFieldBackingIndexResult
    {
        public readonly string IndexCreateTime;
        public readonly string IndexName;
        public readonly string IndexPhrase;
        public readonly string IndexStatus;
        public readonly int IndexStorage;

        [OutputConstructor]
        private GetDescribeIndexListIndexMetaFieldBackingIndexResult(
            string indexCreateTime,

            string indexName,

            string indexPhrase,

            string indexStatus,

            int indexStorage)
        {
            IndexCreateTime = indexCreateTime;
            IndexName = indexName;
            IndexPhrase = indexPhrase;
            IndexStatus = indexStatus;
            IndexStorage = indexStorage;
        }
    }
}
