// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf.Outputs
{

    [OutputType]
    public sealed class IpAccessControlItem
    {
        public readonly int Action;
        public readonly string? Id;
        public readonly string Ip;
        public readonly string Note;
        public readonly string? Source;
        public readonly int? ValidStatus;
        public readonly int ValidTs;

        [OutputConstructor]
        private IpAccessControlItem(
            int action,

            string? id,

            string ip,

            string note,

            string? source,

            int? validStatus,

            int validTs)
        {
            Action = action;
            Id = id;
            Ip = ip;
            Note = note;
            Source = source;
            ValidStatus = validStatus;
            ValidTs = validTs;
        }
    }
}