// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tdmq.Inputs
{

    public sealed class NamespaceRetentionPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("sizeInMb")]
        public Input<int>? SizeInMb { get; set; }

        [Input("timeInMinutes")]
        public Input<int>? TimeInMinutes { get; set; }

        public NamespaceRetentionPolicyGetArgs()
        {
        }
        public static new NamespaceRetentionPolicyGetArgs Empty => new NamespaceRetentionPolicyGetArgs();
    }
}