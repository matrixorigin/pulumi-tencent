// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcm.Inputs
{

    public sealed class MeshConfigSidecarResourcesRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("quantity")]
        public Input<string>? Quantity { get; set; }

        public MeshConfigSidecarResourcesRequestArgs()
        {
        }
        public static new MeshConfigSidecarResourcesRequestArgs Empty => new MeshConfigSidecarResourcesRequestArgs();
    }
}