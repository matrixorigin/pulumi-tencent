// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdwpg.Inputs
{

    public sealed class InstanceChargePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        [Input("renewFlag", required: true)]
        public Input<int> RenewFlag { get; set; } = null!;

        [Input("timeSpan", required: true)]
        public Input<int> TimeSpan { get; set; } = null!;

        [Input("timeUnit", required: true)]
        public Input<string> TimeUnit { get; set; } = null!;

        public InstanceChargePropertiesArgs()
        {
        }
        public static new InstanceChargePropertiesArgs Empty => new InstanceChargePropertiesArgs();
    }
}