// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Inputs
{

    public sealed class BatchManifestArgs : global::Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<Inputs.BatchManifestLocationArgs> Location { get; set; } = null!;

        [Input("spec", required: true)]
        public Input<Inputs.BatchManifestSpecArgs> Spec { get; set; } = null!;

        public BatchManifestArgs()
        {
        }
        public static new BatchManifestArgs Empty => new BatchManifestArgs();
    }
}