// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kms.Inputs
{

    public sealed class OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("identity", required: true)]
        public Input<string> Identity { get; set; } = null!;

        public OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs()
        {
        }
        public static new OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs Empty => new OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs();
    }
}