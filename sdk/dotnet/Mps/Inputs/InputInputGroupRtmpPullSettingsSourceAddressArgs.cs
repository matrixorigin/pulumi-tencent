// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class InputInputGroupRtmpPullSettingsSourceAddressArgs : global::Pulumi.ResourceArgs
    {
        [Input("streamKey", required: true)]
        public Input<string> StreamKey { get; set; } = null!;

        [Input("tcUrl", required: true)]
        public Input<string> TcUrl { get; set; } = null!;

        public InputInputGroupRtmpPullSettingsSourceAddressArgs()
        {
        }
        public static new InputInputGroupRtmpPullSettingsSourceAddressArgs Empty => new InputInputGroupRtmpPullSettingsSourceAddressArgs();
    }
}