// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc.Inputs
{

    public sealed class DetachWorkGroupPolicyOperationPolicySetArgs : global::Pulumi.ResourceArgs
    {
        [Input("catalog", required: true)]
        public Input<string> Catalog { get; set; } = null!;

        [Input("column")]
        public Input<string>? Column { get; set; }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("dataEngine")]
        public Input<string>? DataEngine { get; set; }

        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("function")]
        public Input<string>? Function { get; set; }

        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("operation", required: true)]
        public Input<string> Operation { get; set; } = null!;

        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("reAuth")]
        public Input<bool>? ReAuth { get; set; }

        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("sourceId")]
        public Input<int>? SourceId { get; set; }

        [Input("sourceName")]
        public Input<string>? SourceName { get; set; }

        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public DetachWorkGroupPolicyOperationPolicySetArgs()
        {
        }
        public static new DetachWorkGroupPolicyOperationPolicySetArgs Empty => new DetachWorkGroupPolicyOperationPolicySetArgs();
    }
}
