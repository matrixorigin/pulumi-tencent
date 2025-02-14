// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Redis
{
    [TencentcloudResourceType("tencentcloud:Redis/backupConfig:BackupConfig")]
    public partial class BackupConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`,
        /// `Friday`, `Saturday` and `Sunday`.
        /// </summary>
        [Output("backupPeriods")]
        public Output<ImmutableArray<string>> BackupPeriods { get; private set; } = null!;

        /// <summary>
        /// Specifys what time the backup action should take place. And the time interval should be one hour.
        /// </summary>
        [Output("backupTime")]
        public Output<string> BackupTime { get; private set; } = null!;

        /// <summary>
        /// ID of a redis instance to which the policy will be applied.
        /// </summary>
        [Output("redisId")]
        public Output<string> RedisId { get; private set; } = null!;


        /// <summary>
        /// Create a BackupConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupConfig(string name, BackupConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/backupConfig:BackupConfig", name, args ?? new BackupConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupConfig(string name, Input<string> id, BackupConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/backupConfig:BackupConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupConfig Get(string name, Input<string> id, BackupConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupConfig(name, id, state, options);
        }
    }

    public sealed class BackupConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`,
        /// `Friday`, `Saturday` and `Sunday`.
        /// </summary>
        [Obsolete(@"It has been deprecated from version 1.58.2. It makes no difference to online config at all")]
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Specifys what time the backup action should take place. And the time interval should be one hour.
        /// </summary>
        [Input("backupTime", required: true)]
        public Input<string> BackupTime { get; set; } = null!;

        /// <summary>
        /// ID of a redis instance to which the policy will be applied.
        /// </summary>
        [Input("redisId", required: true)]
        public Input<string> RedisId { get; set; } = null!;

        public BackupConfigArgs()
        {
        }
        public static new BackupConfigArgs Empty => new BackupConfigArgs();
    }

    public sealed class BackupConfigState : global::Pulumi.ResourceArgs
    {
        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`,
        /// `Friday`, `Saturday` and `Sunday`.
        /// </summary>
        [Obsolete(@"It has been deprecated from version 1.58.2. It makes no difference to online config at all")]
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Specifys what time the backup action should take place. And the time interval should be one hour.
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// ID of a redis instance to which the policy will be applied.
        /// </summary>
        [Input("redisId")]
        public Input<string>? RedisId { get; set; }

        public BackupConfigState()
        {
        }
        public static new BackupConfigState Empty => new BackupConfigState();
    }
}
