// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Sg.Outputs
{

    [OutputType]
    public sealed class RuleData
    {
        public readonly string Description;
        public readonly string DestContent;
        public readonly string DestType;
        public readonly string? OrderIndex;
        public readonly string? Port;
        public readonly string? Protocol;
        public readonly string RuleAction;
        public readonly string? ServiceTemplateId;
        public readonly string SourceContent;
        public readonly string SourceType;

        [OutputConstructor]
        private RuleData(
            string description,

            string destContent,

            string destType,

            string? orderIndex,

            string? port,

            string? protocol,

            string ruleAction,

            string? serviceTemplateId,

            string sourceContent,

            string sourceType)
        {
            Description = description;
            DestContent = destContent;
            DestType = destType;
            OrderIndex = orderIndex;
            Port = port;
            Protocol = protocol;
            RuleAction = ruleAction;
            ServiceTemplateId = serviceTemplateId;
            SourceContent = sourceContent;
            SourceType = sourceType;
        }
    }
}