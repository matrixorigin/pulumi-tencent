// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mps.Inputs
{

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("headSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs>? _headSets;
        public InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs> HeadSets
        {
            get => _headSets ?? (_headSets = new InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs>());
            set => _headSets = value;
        }

        [Input("tailSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetGetArgs>? _tailSets;
        public InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetGetArgs> TailSets
        {
            get => _tailSets ?? (_tailSets = new InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetGetArgs>());
            set => _tailSets = value;
        }

        public WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs()
        {
        }
        public static new WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs Empty => new WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs();
    }
}