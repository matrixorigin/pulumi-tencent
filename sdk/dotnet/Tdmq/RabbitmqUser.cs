// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tdmq
{
    [TencentcloudResourceType("tencentcloud:Tdmq/rabbitmqUser:RabbitmqUser")]
    public partial class RabbitmqUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Describe.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Cluster instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The maximum number of channels for this user, if not filled in, there is no limit.
        /// </summary>
        [Output("maxChannels")]
        public Output<int?> MaxChannels { get; private set; } = null!;

        /// <summary>
        /// The maximum number of connections for this user, if not filled in, there is no limit.
        /// </summary>
        [Output("maxConnections")]
        public Output<int?> MaxConnections { get; private set; } = null!;

        /// <summary>
        /// Password, used when logging in.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular
        /// console user, monitoring: management console user, other values: non console user.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Username, used when logging in.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a RabbitmqUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RabbitmqUser(string name, RabbitmqUserArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tdmq/rabbitmqUser:RabbitmqUser", name, args ?? new RabbitmqUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RabbitmqUser(string name, Input<string> id, RabbitmqUserState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tdmq/rabbitmqUser:RabbitmqUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RabbitmqUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RabbitmqUser Get(string name, Input<string> id, RabbitmqUserState? state = null, CustomResourceOptions? options = null)
        {
            return new RabbitmqUser(name, id, state, options);
        }
    }

    public sealed class RabbitmqUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describe.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Cluster instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The maximum number of channels for this user, if not filled in, there is no limit.
        /// </summary>
        [Input("maxChannels")]
        public Input<int>? MaxChannels { get; set; }

        /// <summary>
        /// The maximum number of connections for this user, if not filled in, there is no limit.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password, used when logging in.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular
        /// console user, monitoring: management console user, other values: non console user.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Username, used when logging in.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public RabbitmqUserArgs()
        {
        }
        public static new RabbitmqUserArgs Empty => new RabbitmqUserArgs();
    }

    public sealed class RabbitmqUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describe.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Cluster instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The maximum number of channels for this user, if not filled in, there is no limit.
        /// </summary>
        [Input("maxChannels")]
        public Input<int>? MaxChannels { get; set; }

        /// <summary>
        /// The maximum number of connections for this user, if not filled in, there is no limit.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password, used when logging in.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// User tag, used to determine the permission range for changing user access to RabbitMQ Management. Management: regular
        /// console user, monitoring: management console user, other values: non console user.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Username, used when logging in.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public RabbitmqUserState()
        {
        }
        public static new RabbitmqUserState Empty => new RabbitmqUserState();
    }
}