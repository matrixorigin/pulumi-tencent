// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class GetBusinessLogConfigsResultContentConfigSchemaResult
    {
        public readonly string SchemaContent;
        public readonly string SchemaCreateTime;
        public readonly string SchemaDateFormat;
        public readonly string SchemaMultilinePattern;
        public readonly string SchemaPatternLayout;
        public readonly int SchemaType;

        [OutputConstructor]
        private GetBusinessLogConfigsResultContentConfigSchemaResult(
            string schemaContent,

            string schemaCreateTime,

            string schemaDateFormat,

            string schemaMultilinePattern,

            string schemaPatternLayout,

            int schemaType)
        {
            SchemaContent = schemaContent;
            SchemaCreateTime = schemaCreateTime;
            SchemaDateFormat = schemaDateFormat;
            SchemaMultilinePattern = schemaMultilinePattern;
            SchemaPatternLayout = schemaPatternLayout;
            SchemaType = schemaType;
        }
    }
}