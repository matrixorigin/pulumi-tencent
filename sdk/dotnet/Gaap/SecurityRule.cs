// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    [TencentcloudResourceType("tencentcloud:Gaap/securityRule:SecurityRule")]
    public partial class SecurityRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// A network address block of the request source.
        /// </summary>
        [Output("cidrIp")]
        public Output<string> CidrIp { get; private set; } = null!;

        /// <summary>
        /// Name of the security policy rule. Maximum length is 30.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the security policy.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        /// </summary>
        [Output("port")]
        public Output<string?> Port { get; private set; } = null!;

        /// <summary>
        /// Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityRule(string name, SecurityRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/securityRule:SecurityRule", name, args ?? new SecurityRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityRule(string name, Input<string> id, SecurityRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/securityRule:SecurityRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityRule Get(string name, Input<string> id, SecurityRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityRule(name, id, state, options);
        }
    }

    public sealed class SecurityRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// A network address block of the request source.
        /// </summary>
        [Input("cidrIp", required: true)]
        public Input<string> CidrIp { get; set; } = null!;

        /// <summary>
        /// Name of the security policy rule. Maximum length is 30.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the security policy.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public SecurityRuleArgs()
        {
        }
        public static new SecurityRuleArgs Empty => new SecurityRuleArgs();
    }

    public sealed class SecurityRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy of the rule. Valid value: `ACCEPT` and `DROP`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A network address block of the request source.
        /// </summary>
        [Input("cidrIp")]
        public Input<string>? CidrIp { get; set; }

        /// <summary>
        /// Name of the security policy rule. Maximum length is 30.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the security policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public SecurityRuleState()
        {
        }
        public static new SecurityRuleState Empty => new SecurityRuleState();
    }
}