// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Inputs
{

    public sealed class BatchManifestGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<Inputs.BatchManifestLocationGetArgs> Location { get; set; } = null!;

        [Input("spec", required: true)]
        public Input<Inputs.BatchManifestSpecGetArgs> Spec { get; set; } = null!;

        public BatchManifestGetArgs()
        {
        }
        public static new BatchManifestGetArgs Empty => new BatchManifestGetArgs();
    }
}