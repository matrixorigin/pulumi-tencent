// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class RealtimeLogDeliveryDeliveryConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Inputs.RealtimeLogDeliveryDeliveryConditionConditionArgs>? _conditions;
        public InputList<Inputs.RealtimeLogDeliveryDeliveryConditionConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.RealtimeLogDeliveryDeliveryConditionConditionArgs>());
            set => _conditions = value;
        }

        public RealtimeLogDeliveryDeliveryConditionArgs()
        {
        }
        public static new RealtimeLogDeliveryDeliveryConditionArgs Empty => new RealtimeLogDeliveryDeliveryConditionArgs();
    }
}
