// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class CustomDomainEndpointsConfigPathRewriteArgs : global::Pulumi.ResourceArgs
    {
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("rewrite", required: true)]
        public Input<string> Rewrite { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CustomDomainEndpointsConfigPathRewriteArgs()
        {
        }
        public static new CustomDomainEndpointsConfigPathRewriteArgs Empty => new CustomDomainEndpointsConfigPathRewriteArgs();
    }
}