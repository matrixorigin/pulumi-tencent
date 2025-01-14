// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tat
{
    [TencentcloudResourceType("tencentcloud:Tat/invocationInvokeAttachment:InvocationInvokeAttachment")]
    public partial class InvocationInvokeAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        [Output("commandId")]
        public Output<string> CommandId { get; private set; } = null!;

        /// <summary>
        /// ID of instances about to execute commands. Supported instance types: CVM LIGHTHOUSE.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The COS bucket URL for uploading logs. The URL must start with https, such as
        /// https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        /// </summary>
        [Output("outputCosBucketUrl")]
        public Output<string?> OutputCosBucketUrl { get; private set; } = null!;

        /// <summary>
        /// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a
        /// combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a
        /// subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive
        /// slashes.
        /// </summary>
        [Output("outputCosKeyPrefix")]
        public Output<string?> OutputCosKeyPrefix { get; private set; } = null!;

        /// <summary>
        /// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the
        /// custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the
        /// DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64
        /// characters and can contain [a-z], [A-Z], [0-9] and [-_].
        /// </summary>
        [Output("parameters")]
        public Output<string?> Parameters { get; private set; } = null!;

        /// <summary>
        /// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best
        /// practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root
        /// is used to execute commands on Linux and the user System is used on Windows.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for
        /// POWERSHELL commands.
        /// </summary>
        [Output("workingDirectory")]
        public Output<string?> WorkingDirectory { get; private set; } = null!;


        /// <summary>
        /// Create a InvocationInvokeAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InvocationInvokeAttachment(string name, InvocationInvokeAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tat/invocationInvokeAttachment:InvocationInvokeAttachment", name, args ?? new InvocationInvokeAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InvocationInvokeAttachment(string name, Input<string> id, InvocationInvokeAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tat/invocationInvokeAttachment:InvocationInvokeAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InvocationInvokeAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InvocationInvokeAttachment Get(string name, Input<string> id, InvocationInvokeAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InvocationInvokeAttachment(name, id, state, options);
        }
    }

    public sealed class InvocationInvokeAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        [Input("commandId", required: true)]
        public Input<string> CommandId { get; set; } = null!;

        /// <summary>
        /// ID of instances about to execute commands. Supported instance types: CVM LIGHTHOUSE.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The COS bucket URL for uploading logs. The URL must start with https, such as
        /// https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        /// </summary>
        [Input("outputCosBucketUrl")]
        public Input<string>? OutputCosBucketUrl { get; set; }

        /// <summary>
        /// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a
        /// combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a
        /// subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive
        /// slashes.
        /// </summary>
        [Input("outputCosKeyPrefix")]
        public Input<string>? OutputCosKeyPrefix { get; set; }

        /// <summary>
        /// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the
        /// custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the
        /// DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64
        /// characters and can contain [a-z], [A-Z], [0-9] and [-_].
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        /// <summary>
        /// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best
        /// practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root
        /// is used to execute commands on Linux and the user System is used on Windows.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for
        /// POWERSHELL commands.
        /// </summary>
        [Input("workingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public InvocationInvokeAttachmentArgs()
        {
        }
        public static new InvocationInvokeAttachmentArgs Empty => new InvocationInvokeAttachmentArgs();
    }

    public sealed class InvocationInvokeAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command ID.
        /// </summary>
        [Input("commandId")]
        public Input<string>? CommandId { get; set; }

        /// <summary>
        /// ID of instances about to execute commands. Supported instance types: CVM LIGHTHOUSE.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The COS bucket URL for uploading logs. The URL must start with https, such as
        /// https://BucketName-123454321.cos.ap-beijing.myqcloud.com.
        /// </summary>
        [Input("outputCosBucketUrl")]
        public Input<string>? OutputCosBucketUrl { get; set; }

        /// <summary>
        /// The COS bucket directory where the logs are saved; Check below for the rules of the directory name: 1 It must be a
        /// combination of number, letters, and visible characters, Up to 60 characters are allowed; 2 Use a slash (/) to create a
        /// subdirectory; 3 can not be used as the folder name; It cannot start with a slash (/), and cannot contain consecutive
        /// slashes.
        /// </summary>
        [Input("outputCosKeyPrefix")]
        public Input<string>? OutputCosKeyPrefix { get; set; }

        /// <summary>
        /// Custom parameters of Command. The field type is JSON encoded string. For example, {varA: 222}.key is the name of the
        /// custom parameter and value is the default value. Both key and value are strings.If no parameter value is provided, the
        /// DefaultParameters is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64
        /// characters and can contain [a-z], [A-Z], [0-9] and [-_].
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        /// <summary>
        /// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best
        /// practice for permission management. We recommend you execute TAT commands as a general user. By default, the user root
        /// is used to execute commands on Linux and the user System is used on Windows.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Command execution path. The default value is /root for SHELL commands and C:Program Filesqcloudtat_agentworkdir for
        /// POWERSHELL commands.
        /// </summary>
        [Input("workingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public InvocationInvokeAttachmentState()
        {
        }
        public static new InvocationInvokeAttachmentState Empty => new InvocationInvokeAttachmentState();
    }
}
