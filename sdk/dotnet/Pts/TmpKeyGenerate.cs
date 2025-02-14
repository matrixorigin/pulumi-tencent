// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts
{
    [TencentcloudResourceType("tencentcloud:Pts/tmpKeyGenerate:TmpKeyGenerate")]
    public partial class TmpKeyGenerate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Temporary access credentials.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.TmpKeyGenerateCredential>> Credentials { get; private set; } = null!;

        /// <summary>
        /// Timestamp of temporary access credential timeout (in seconds).
        /// </summary>
        [Output("expiredTime")]
        public Output<int> ExpiredTime { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Scenario ID.
        /// </summary>
        [Output("scenarioId")]
        public Output<string?> ScenarioId { get; private set; } = null!;

        /// <summary>
        /// The timestamp of the moment when the temporary access credential was obtained (in seconds).
        /// </summary>
        [Output("startTime")]
        public Output<int> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a TmpKeyGenerate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpKeyGenerate(string name, TmpKeyGenerateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/tmpKeyGenerate:TmpKeyGenerate", name, args ?? new TmpKeyGenerateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpKeyGenerate(string name, Input<string> id, TmpKeyGenerateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/tmpKeyGenerate:TmpKeyGenerate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpKeyGenerate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpKeyGenerate Get(string name, Input<string> id, TmpKeyGenerateState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpKeyGenerate(name, id, state, options);
        }
    }

    public sealed class TmpKeyGenerateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Scenario ID.
        /// </summary>
        [Input("scenarioId")]
        public Input<string>? ScenarioId { get; set; }

        public TmpKeyGenerateArgs()
        {
        }
        public static new TmpKeyGenerateArgs Empty => new TmpKeyGenerateArgs();
    }

    public sealed class TmpKeyGenerateState : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.TmpKeyGenerateCredentialGetArgs>? _credentials;

        /// <summary>
        /// Temporary access credentials.
        /// </summary>
        public InputList<Inputs.TmpKeyGenerateCredentialGetArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.TmpKeyGenerateCredentialGetArgs>());
            set => _credentials = value;
        }

        /// <summary>
        /// Timestamp of temporary access credential timeout (in seconds).
        /// </summary>
        [Input("expiredTime")]
        public Input<int>? ExpiredTime { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Scenario ID.
        /// </summary>
        [Input("scenarioId")]
        public Input<string>? ScenarioId { get; set; }

        /// <summary>
        /// The timestamp of the moment when the temporary access credential was obtained (in seconds).
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        public TmpKeyGenerateState()
        {
        }
        public static new TmpKeyGenerateState Empty => new TmpKeyGenerateState();
    }
}
