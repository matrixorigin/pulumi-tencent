// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Inputs
{

    public sealed class LoadBalancerForwardLoadBalancerTargetAttributeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public LoadBalancerForwardLoadBalancerTargetAttributeGetArgs()
        {
        }
        public static new LoadBalancerForwardLoadBalancerTargetAttributeGetArgs Empty => new LoadBalancerForwardLoadBalancerTargetAttributeGetArgs();
    }
}