// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Identity
{
    [TencentcloudResourceType("tencentcloud:Identity/centerUser:CenterUser")]
    public partial class CenterUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// User's description. Length: Maximum 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the user. Length: Maximum 256 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The user's email address. Must be unique within the catalog. Length: Maximum 128 characters.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// The user's last name. Length: Maximum 64 characters.
        /// </summary>
        [Output("firstName")]
        public Output<string?> FirstName { get; private set; } = null!;

        /// <summary>
        /// The user's name. Length: Maximum 64 characters.
        /// </summary>
        [Output("lastName")]
        public Output<string?> LastName { get; private set; } = null!;

        /// <summary>
        /// Update time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// User id.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// User name. It must be unique in space. Modifications are not supported. Format: Contains numbers, English letters and
        /// special symbols(`+`, `=`, `,`, `.`, `@`, `-`, `_`). Length: Maximum 64 characters.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;

        /// <summary>
        /// The status of the user. Value: Enabled (default): Enabled. Disabled: Disabled.
        /// </summary>
        [Output("userStatus")]
        public Output<string> UserStatus { get; private set; } = null!;

        /// <summary>
        /// User type.
        /// </summary>
        [Output("userType")]
        public Output<string> UserType { get; private set; } = null!;

        /// <summary>
        /// Zone id.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a CenterUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CenterUser(string name, CenterUserArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Identity/centerUser:CenterUser", name, args ?? new CenterUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CenterUser(string name, Input<string> id, CenterUserState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Identity/centerUser:CenterUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CenterUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CenterUser Get(string name, Input<string> id, CenterUserState? state = null, CustomResourceOptions? options = null)
        {
            return new CenterUser(name, id, state, options);
        }
    }

    public sealed class CenterUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User's description. Length: Maximum 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the user. Length: Maximum 256 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The user's email address. Must be unique within the catalog. Length: Maximum 128 characters.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The user's last name. Length: Maximum 64 characters.
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// The user's name. Length: Maximum 64 characters.
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// User name. It must be unique in space. Modifications are not supported. Format: Contains numbers, English letters and
        /// special symbols(`+`, `=`, `,`, `.`, `@`, `-`, `_`). Length: Maximum 64 characters.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        /// <summary>
        /// The status of the user. Value: Enabled (default): Enabled. Disabled: Disabled.
        /// </summary>
        [Input("userStatus")]
        public Input<string>? UserStatus { get; set; }

        /// <summary>
        /// Zone id.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public CenterUserArgs()
        {
        }
        public static new CenterUserArgs Empty => new CenterUserArgs();
    }

    public sealed class CenterUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// User's description. Length: Maximum 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the user. Length: Maximum 256 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The user's email address. Must be unique within the catalog. Length: Maximum 128 characters.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The user's last name. Length: Maximum 64 characters.
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// The user's name. Length: Maximum 64 characters.
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// Update time.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// User id.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// User name. It must be unique in space. Modifications are not supported. Format: Contains numbers, English letters and
        /// special symbols(`+`, `=`, `,`, `.`, `@`, `-`, `_`). Length: Maximum 64 characters.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// The status of the user. Value: Enabled (default): Enabled. Disabled: Disabled.
        /// </summary>
        [Input("userStatus")]
        public Input<string>? UserStatus { get; set; }

        /// <summary>
        /// User type.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        /// <summary>
        /// Zone id.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public CenterUserState()
        {
        }
        public static new CenterUserState Empty => new CenterUserState();
    }
}