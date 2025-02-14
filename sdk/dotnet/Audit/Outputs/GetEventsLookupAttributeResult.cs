// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Audit.Outputs
{

    [OutputType]
    public sealed class GetEventsLookupAttributeResult
    {
        public readonly string AttributeKey;
        public readonly string? AttributeValue;

        [OutputConstructor]
        private GetEventsLookupAttributeResult(
            string attributeKey,

            string? attributeValue)
        {
            AttributeKey = attributeKey;
            AttributeValue = attributeValue;
        }
    }
}
