// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tse.Inputs
{

    public sealed class CngwStrategyConfigBehaviorScaleUpArgs : global::Pulumi.ResourceArgs
    {
        [Input("policies")]
        private InputList<Inputs.CngwStrategyConfigBehaviorScaleUpPolicyArgs>? _policies;
        public InputList<Inputs.CngwStrategyConfigBehaviorScaleUpPolicyArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.CngwStrategyConfigBehaviorScaleUpPolicyArgs>());
            set => _policies = value;
        }

        [Input("selectPolicy")]
        public Input<string>? SelectPolicy { get; set; }

        [Input("stabilizationWindowSeconds")]
        public Input<int>? StabilizationWindowSeconds { get; set; }

        public CngwStrategyConfigBehaviorScaleUpArgs()
        {
        }
        public static new CngwStrategyConfigBehaviorScaleUpArgs Empty => new CngwStrategyConfigBehaviorScaleUpArgs();
    }
}