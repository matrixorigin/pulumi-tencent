// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tse
{
    [TencentcloudResourceType("tencentcloud:Tse/cngwStrategy:CngwStrategy")]
    public partial class CngwStrategy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// configuration of metric scaling.
        /// </summary>
        [Output("config")]
        public Output<Outputs.CngwStrategyConfig?> Config { get; private set; } = null!;

        /// <summary>
        /// configuration of timed scaling.
        /// </summary>
        [Output("cronConfig")]
        public Output<Outputs.CngwStrategyCronConfig?> CronConfig { get; private set; } = null!;

        /// <summary>
        /// description information, up to 120 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// gateway ID.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// strategy ID Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        [Output("strategyId")]
        public Output<string> StrategyId { get; private set; } = null!;

        /// <summary>
        /// strategy name, up to 20 characters.
        /// </summary>
        [Output("strategyName")]
        public Output<string> StrategyName { get; private set; } = null!;


        /// <summary>
        /// Create a CngwStrategy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CngwStrategy(string name, CngwStrategyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwStrategy:CngwStrategy", name, args ?? new CngwStrategyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CngwStrategy(string name, Input<string> id, CngwStrategyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwStrategy:CngwStrategy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CngwStrategy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CngwStrategy Get(string name, Input<string> id, CngwStrategyState? state = null, CustomResourceOptions? options = null)
        {
            return new CngwStrategy(name, id, state, options);
        }
    }

    public sealed class CngwStrategyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// configuration of metric scaling.
        /// </summary>
        [Input("config")]
        public Input<Inputs.CngwStrategyConfigArgs>? Config { get; set; }

        /// <summary>
        /// configuration of timed scaling.
        /// </summary>
        [Input("cronConfig")]
        public Input<Inputs.CngwStrategyCronConfigArgs>? CronConfig { get; set; }

        /// <summary>
        /// description information, up to 120 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// strategy name, up to 20 characters.
        /// </summary>
        [Input("strategyName", required: true)]
        public Input<string> StrategyName { get; set; } = null!;

        public CngwStrategyArgs()
        {
        }
        public static new CngwStrategyArgs Empty => new CngwStrategyArgs();
    }

    public sealed class CngwStrategyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// configuration of metric scaling.
        /// </summary>
        [Input("config")]
        public Input<Inputs.CngwStrategyConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// configuration of timed scaling.
        /// </summary>
        [Input("cronConfig")]
        public Input<Inputs.CngwStrategyCronConfigGetArgs>? CronConfig { get; set; }

        /// <summary>
        /// description information, up to 120 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// gateway ID.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// strategy ID Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        [Input("strategyId")]
        public Input<string>? StrategyId { get; set; }

        /// <summary>
        /// strategy name, up to 20 characters.
        /// </summary>
        [Input("strategyName")]
        public Input<string>? StrategyName { get; set; }

        public CngwStrategyState()
        {
        }
        public static new CngwStrategyState Empty => new CngwStrategyState();
    }
}