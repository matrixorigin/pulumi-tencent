// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ci.Inputs
{

    public sealed class MediaSnapshotTemplateSnapshotSpriteSnapshotConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("cellHeight")]
        public Input<string>? CellHeight { get; set; }

        [Input("cellWidth")]
        public Input<string>? CellWidth { get; set; }

        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        [Input("columns", required: true)]
        public Input<string> Columns { get; set; } = null!;

        [Input("lines", required: true)]
        public Input<string> Lines { get; set; } = null!;

        [Input("margin")]
        public Input<string>? Margin { get; set; }

        [Input("padding")]
        public Input<string>? Padding { get; set; }

        public MediaSnapshotTemplateSnapshotSpriteSnapshotConfigArgs()
        {
        }
        public static new MediaSnapshotTemplateSnapshotSpriteSnapshotConfigArgs Empty => new MediaSnapshotTemplateSnapshotSpriteSnapshotConfigArgs();
    }
}