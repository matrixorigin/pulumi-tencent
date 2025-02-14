// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc
{
    [TencentcloudResourceType("tencentcloud:Dlc/modifyDataEngineDescriptionOperation:ModifyDataEngineDescriptionOperation")]
    public partial class ModifyDataEngineDescriptionOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the engine to modify.
        /// </summary>
        [Output("dataEngineName")]
        public Output<string> DataEngineName { get; private set; } = null!;

        /// <summary>
        /// Engine description information, the maximum length is 250.
        /// </summary>
        [Output("message")]
        public Output<string> Message { get; private set; } = null!;


        /// <summary>
        /// Create a ModifyDataEngineDescriptionOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModifyDataEngineDescriptionOperation(string name, ModifyDataEngineDescriptionOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/modifyDataEngineDescriptionOperation:ModifyDataEngineDescriptionOperation", name, args ?? new ModifyDataEngineDescriptionOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModifyDataEngineDescriptionOperation(string name, Input<string> id, ModifyDataEngineDescriptionOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/modifyDataEngineDescriptionOperation:ModifyDataEngineDescriptionOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ModifyDataEngineDescriptionOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModifyDataEngineDescriptionOperation Get(string name, Input<string> id, ModifyDataEngineDescriptionOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ModifyDataEngineDescriptionOperation(name, id, state, options);
        }
    }

    public sealed class ModifyDataEngineDescriptionOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the engine to modify.
        /// </summary>
        [Input("dataEngineName", required: true)]
        public Input<string> DataEngineName { get; set; } = null!;

        /// <summary>
        /// Engine description information, the maximum length is 250.
        /// </summary>
        [Input("message", required: true)]
        public Input<string> Message { get; set; } = null!;

        public ModifyDataEngineDescriptionOperationArgs()
        {
        }
        public static new ModifyDataEngineDescriptionOperationArgs Empty => new ModifyDataEngineDescriptionOperationArgs();
    }

    public sealed class ModifyDataEngineDescriptionOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the engine to modify.
        /// </summary>
        [Input("dataEngineName")]
        public Input<string>? DataEngineName { get; set; }

        /// <summary>
        /// Engine description information, the maximum length is 250.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public ModifyDataEngineDescriptionOperationState()
        {
        }
        public static new ModifyDataEngineDescriptionOperationState Empty => new ModifyDataEngineDescriptionOperationState();
    }
}
