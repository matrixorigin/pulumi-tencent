// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Teo.Inputs
{

    public sealed class RuleEngineRuleSubRuleRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.RuleEngineRuleSubRuleRuleActionArgs>? _actions;
        public InputList<Inputs.RuleEngineRuleSubRuleRuleActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.RuleEngineRuleSubRuleRuleActionArgs>());
            set => _actions = value;
        }

        [Input("ors", required: true)]
        private InputList<Inputs.RuleEngineRuleSubRuleRuleOrArgs>? _ors;
        public InputList<Inputs.RuleEngineRuleSubRuleRuleOrArgs> Ors
        {
            get => _ors ?? (_ors = new InputList<Inputs.RuleEngineRuleSubRuleRuleOrArgs>());
            set => _ors = value;
        }

        public RuleEngineRuleSubRuleRuleArgs()
        {
        }
        public static new RuleEngineRuleSubRuleRuleArgs Empty => new RuleEngineRuleSubRuleRuleArgs();
    }
}