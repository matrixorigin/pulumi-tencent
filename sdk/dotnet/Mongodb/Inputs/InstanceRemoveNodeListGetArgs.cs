// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mongodb.Inputs
{

    public sealed class InstanceRemoveNodeListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("nodeName", required: true)]
        public Input<string> NodeName { get; set; } = null!;

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceRemoveNodeListGetArgs()
        {
        }
        public static new InstanceRemoveNodeListGetArgs Empty => new InstanceRemoveNodeListGetArgs();
    }
}