// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dlc
{
    [TencentcloudResourceType("tencentcloud:Dlc/detachUserPolicyOperation:DetachUserPolicyOperation")]
    public partial class DetachUserPolicyOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Authentication policy collection.
        /// </summary>
        [Output("policySets")]
        public Output<ImmutableArray<Outputs.DetachUserPolicyOperationPolicySet>> PolicySets { get; private set; } = null!;

        /// <summary>
        /// User id, the same as the sub-user uin.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a DetachUserPolicyOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DetachUserPolicyOperation(string name, DetachUserPolicyOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/detachUserPolicyOperation:DetachUserPolicyOperation", name, args ?? new DetachUserPolicyOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DetachUserPolicyOperation(string name, Input<string> id, DetachUserPolicyOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/detachUserPolicyOperation:DetachUserPolicyOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DetachUserPolicyOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DetachUserPolicyOperation Get(string name, Input<string> id, DetachUserPolicyOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new DetachUserPolicyOperation(name, id, state, options);
        }
    }

    public sealed class DetachUserPolicyOperationArgs : global::Pulumi.ResourceArgs
    {
        [Input("policySets")]
        private InputList<Inputs.DetachUserPolicyOperationPolicySetArgs>? _policySets;

        /// <summary>
        /// Authentication policy collection.
        /// </summary>
        public InputList<Inputs.DetachUserPolicyOperationPolicySetArgs> PolicySets
        {
            get => _policySets ?? (_policySets = new InputList<Inputs.DetachUserPolicyOperationPolicySetArgs>());
            set => _policySets = value;
        }

        /// <summary>
        /// User id, the same as the sub-user uin.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public DetachUserPolicyOperationArgs()
        {
        }
        public static new DetachUserPolicyOperationArgs Empty => new DetachUserPolicyOperationArgs();
    }

    public sealed class DetachUserPolicyOperationState : global::Pulumi.ResourceArgs
    {
        [Input("policySets")]
        private InputList<Inputs.DetachUserPolicyOperationPolicySetGetArgs>? _policySets;

        /// <summary>
        /// Authentication policy collection.
        /// </summary>
        public InputList<Inputs.DetachUserPolicyOperationPolicySetGetArgs> PolicySets
        {
            get => _policySets ?? (_policySets = new InputList<Inputs.DetachUserPolicyOperationPolicySetGetArgs>());
            set => _policySets = value;
        }

        /// <summary>
        /// User id, the same as the sub-user uin.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public DetachUserPolicyOperationState()
        {
        }
        public static new DetachUserPolicyOperationState Empty => new DetachUserPolicyOperationState();
    }
}
