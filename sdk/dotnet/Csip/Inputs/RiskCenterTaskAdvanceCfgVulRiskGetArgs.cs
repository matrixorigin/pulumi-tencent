// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Csip.Inputs
{

    public sealed class RiskCenterTaskAdvanceCfgVulRiskGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("enable", required: true)]
        public Input<int> Enable { get; set; } = null!;

        [Input("riskId", required: true)]
        public Input<string> RiskId { get; set; } = null!;

        public RiskCenterTaskAdvanceCfgVulRiskGetArgs()
        {
        }
        public static new RiskCenterTaskAdvanceCfgVulRiskGetArgs Empty => new RiskCenterTaskAdvanceCfgVulRiskGetArgs();
    }
}
