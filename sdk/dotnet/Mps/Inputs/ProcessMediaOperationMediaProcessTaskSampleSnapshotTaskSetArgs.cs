// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        [Input("objectNumberFormat")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetObjectNumberFormatArgs>? ObjectNumberFormat { get; set; }

        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        [Input("outputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetOutputStorageArgs>? OutputStorage { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetArgs>? _watermarkSets;
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetWatermarkSetArgs>());
            set => _watermarkSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetArgs()
        {
        }
        public static new ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetArgs Empty => new ProcessMediaOperationMediaProcessTaskSampleSnapshotTaskSetArgs();
    }
}
