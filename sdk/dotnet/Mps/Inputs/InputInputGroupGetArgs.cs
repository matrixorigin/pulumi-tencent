// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class InputInputGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowIpLists")]
        private InputList<string>? _allowIpLists;
        public InputList<string> AllowIpLists
        {
            get => _allowIpLists ?? (_allowIpLists = new InputList<string>());
            set => _allowIpLists = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("failOver")]
        public Input<string>? FailOver { get; set; }

        [Input("hlsPullSettings")]
        public Input<Inputs.InputInputGroupHlsPullSettingsGetArgs>? HlsPullSettings { get; set; }

        [Input("inputName", required: true)]
        public Input<string> InputName { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("resilientStream")]
        public Input<Inputs.InputInputGroupResilientStreamGetArgs>? ResilientStream { get; set; }

        [Input("rtmpPullSettings")]
        public Input<Inputs.InputInputGroupRtmpPullSettingsGetArgs>? RtmpPullSettings { get; set; }

        [Input("rtpSettings")]
        public Input<Inputs.InputInputGroupRtpSettingsGetArgs>? RtpSettings { get; set; }

        [Input("rtspPullSettings")]
        public Input<Inputs.InputInputGroupRtspPullSettingsGetArgs>? RtspPullSettings { get; set; }

        [Input("srtSettings")]
        public Input<Inputs.InputInputGroupSrtSettingsGetArgs>? SrtSettings { get; set; }

        public InputInputGroupGetArgs()
        {
        }
        public static new InputInputGroupGetArgs Empty => new InputInputGroupGetArgs();
    }
}
