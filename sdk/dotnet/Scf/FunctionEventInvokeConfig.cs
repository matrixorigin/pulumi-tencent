// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf
{
    [TencentcloudResourceType("tencentcloud:Scf/functionEventInvokeConfig:FunctionEventInvokeConfig")]
    public partial class FunctionEventInvokeConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Async retry configuration information.
        /// </summary>
        [Output("asyncTriggerConfig")]
        public Output<Outputs.FunctionEventInvokeConfigAsyncTriggerConfig> AsyncTriggerConfig { get; private set; } = null!;

        /// <summary>
        /// Function name.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// Function namespace. Default value: default.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionEventInvokeConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionEventInvokeConfig(string name, FunctionEventInvokeConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/functionEventInvokeConfig:FunctionEventInvokeConfig", name, args ?? new FunctionEventInvokeConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionEventInvokeConfig(string name, Input<string> id, FunctionEventInvokeConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/functionEventInvokeConfig:FunctionEventInvokeConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionEventInvokeConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionEventInvokeConfig Get(string name, Input<string> id, FunctionEventInvokeConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionEventInvokeConfig(name, id, state, options);
        }
    }

    public sealed class FunctionEventInvokeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Async retry configuration information.
        /// </summary>
        [Input("asyncTriggerConfig", required: true)]
        public Input<Inputs.FunctionEventInvokeConfigAsyncTriggerConfigArgs> AsyncTriggerConfig { get; set; } = null!;

        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Function namespace. Default value: default.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public FunctionEventInvokeConfigArgs()
        {
        }
        public static new FunctionEventInvokeConfigArgs Empty => new FunctionEventInvokeConfigArgs();
    }

    public sealed class FunctionEventInvokeConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Async retry configuration information.
        /// </summary>
        [Input("asyncTriggerConfig")]
        public Input<Inputs.FunctionEventInvokeConfigAsyncTriggerConfigGetArgs>? AsyncTriggerConfig { get; set; }

        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// Function namespace. Default value: default.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public FunctionEventInvokeConfigState()
        {
        }
        public static new FunctionEventInvokeConfigState Empty => new FunctionEventInvokeConfigState();
    }
}
