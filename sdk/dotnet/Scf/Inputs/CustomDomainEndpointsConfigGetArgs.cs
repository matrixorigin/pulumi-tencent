// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class CustomDomainEndpointsConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("pathMatch", required: true)]
        public Input<string> PathMatch { get; set; } = null!;

        [Input("pathRewrites")]
        private InputList<Inputs.CustomDomainEndpointsConfigPathRewriteGetArgs>? _pathRewrites;
        public InputList<Inputs.CustomDomainEndpointsConfigPathRewriteGetArgs> PathRewrites
        {
            get => _pathRewrites ?? (_pathRewrites = new InputList<Inputs.CustomDomainEndpointsConfigPathRewriteGetArgs>());
            set => _pathRewrites = value;
        }

        [Input("qualifier", required: true)]
        public Input<string> Qualifier { get; set; } = null!;

        public CustomDomainEndpointsConfigGetArgs()
        {
        }
        public static new CustomDomainEndpointsConfigGetArgs Empty => new CustomDomainEndpointsConfigGetArgs();
    }
}