// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tse.Inputs
{

    public sealed class CngwStrategyConfigMetricGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        [Input("targetValue")]
        public Input<int>? TargetValue { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public CngwStrategyConfigMetricGetArgs()
        {
        }
        public static new CngwStrategyConfigMetricGetArgs Empty => new CngwStrategyConfigMetricGetArgs();
    }
}
