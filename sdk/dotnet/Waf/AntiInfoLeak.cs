// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Waf
{
    [TencentcloudResourceType("tencentcloud:Waf/antiInfoLeak:AntiInfoLeak")]
    public partial class AntiInfoLeak : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Rule Action. 0: alarm; 1: replacement; 2: only displaying the first four digits; 3: only displaying the last four
        /// digits; 4: blocking.
        /// </summary>
        [Output("actionType")]
        public Output<int> ActionType { get; private set; } = null!;

        /// <summary>
        /// Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// status.
        /// </summary>
        [Output("status")]
        public Output<int?> Status { get; private set; } = null!;

        /// <summary>
        /// Strategies detail.
        /// </summary>
        [Output("strategies")]
        public Output<ImmutableArray<Outputs.AntiInfoLeakStrategy>> Strategies { get; private set; } = null!;

        /// <summary>
        /// Uri.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a AntiInfoLeak resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AntiInfoLeak(string name, AntiInfoLeakArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/antiInfoLeak:AntiInfoLeak", name, args ?? new AntiInfoLeakArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AntiInfoLeak(string name, Input<string> id, AntiInfoLeakState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/antiInfoLeak:AntiInfoLeak", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AntiInfoLeak resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AntiInfoLeak Get(string name, Input<string> id, AntiInfoLeakState? state = null, CustomResourceOptions? options = null)
        {
            return new AntiInfoLeak(name, id, state, options);
        }
    }

    public sealed class AntiInfoLeakArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule Action. 0: alarm; 1: replacement; 2: only displaying the first four digits; 3: only displaying the last four
        /// digits; 4: blocking.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<int> ActionType { get; set; } = null!;

        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("strategies", required: true)]
        private InputList<Inputs.AntiInfoLeakStrategyArgs>? _strategies;

        /// <summary>
        /// Strategies detail.
        /// </summary>
        public InputList<Inputs.AntiInfoLeakStrategyArgs> Strategies
        {
            get => _strategies ?? (_strategies = new InputList<Inputs.AntiInfoLeakStrategyArgs>());
            set => _strategies = value;
        }

        /// <summary>
        /// Uri.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public AntiInfoLeakArgs()
        {
        }
        public static new AntiInfoLeakArgs Empty => new AntiInfoLeakArgs();
    }

    public sealed class AntiInfoLeakState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule Action. 0: alarm; 1: replacement; 2: only displaying the first four digits; 3: only displaying the last four
        /// digits; 4: blocking.
        /// </summary>
        [Input("actionType")]
        public Input<int>? ActionType { get; set; }

        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("strategies")]
        private InputList<Inputs.AntiInfoLeakStrategyGetArgs>? _strategies;

        /// <summary>
        /// Strategies detail.
        /// </summary>
        public InputList<Inputs.AntiInfoLeakStrategyGetArgs> Strategies
        {
            get => _strategies ?? (_strategies = new InputList<Inputs.AntiInfoLeakStrategyGetArgs>());
            set => _strategies = value;
        }

        /// <summary>
        /// Uri.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public AntiInfoLeakState()
        {
        }
        public static new AntiInfoLeakState Empty => new AntiInfoLeakState();
    }
}