// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn
{
    [TencentcloudResourceType("tencentcloud:Ccn/routes:Routes")]
    public partial class Routes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// CCN Instance ID.
        /// </summary>
        [Output("ccnId")]
        public Output<string> CcnId { get; private set; } = null!;

        /// <summary>
        /// CCN Route Id List.
        /// </summary>
        [Output("routeId")]
        public Output<string> RouteId { get; private set; } = null!;

        /// <summary>
        /// `on`: Enable, `off`: Disable.
        /// </summary>
        [Output("switch")]
        public Output<string> Switch { get; private set; } = null!;


        /// <summary>
        /// Create a Routes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Routes(string name, RoutesArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/routes:Routes", name, args ?? new RoutesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Routes(string name, Input<string> id, RoutesState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/routes:Routes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Routes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Routes Get(string name, Input<string> id, RoutesState? state = null, CustomResourceOptions? options = null)
        {
            return new Routes(name, id, state, options);
        }
    }

    public sealed class RoutesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CCN Instance ID.
        /// </summary>
        [Input("ccnId", required: true)]
        public Input<string> CcnId { get; set; } = null!;

        /// <summary>
        /// CCN Route Id List.
        /// </summary>
        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        /// <summary>
        /// `on`: Enable, `off`: Disable.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public RoutesArgs()
        {
        }
        public static new RoutesArgs Empty => new RoutesArgs();
    }

    public sealed class RoutesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CCN Instance ID.
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// CCN Route Id List.
        /// </summary>
        [Input("routeId")]
        public Input<string>? RouteId { get; set; }

        /// <summary>
        /// `on`: Enable, `off`: Disable.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public RoutesState()
        {
        }
        public static new RoutesState Empty => new RoutesState();
    }
}