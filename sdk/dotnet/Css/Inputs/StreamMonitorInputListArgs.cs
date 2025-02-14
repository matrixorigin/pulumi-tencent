// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css.Inputs
{

    public sealed class StreamMonitorInputListArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputApp")]
        public Input<string>? InputApp { get; set; }

        [Input("inputDomain")]
        public Input<string>? InputDomain { get; set; }

        [Input("inputStreamName", required: true)]
        public Input<string> InputStreamName { get; set; } = null!;

        [Input("inputUrl")]
        public Input<string>? InputUrl { get; set; }

        public StreamMonitorInputListArgs()
        {
        }
        public static new StreamMonitorInputListArgs Empty => new StreamMonitorInputListArgs();
    }
}
