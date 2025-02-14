// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn
{
    [TencentcloudResourceType("tencentcloud:Ccn/attachment:Attachment")]
    public partial class Attachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time of attaching.
        /// </summary>
        [Output("attachedTime")]
        public Output<string> AttachedTime { get; private set; } = null!;

        /// <summary>
        /// ID of the CCN.
        /// </summary>
        [Output("ccnId")]
        public Output<string> CcnId { get; private set; } = null!;

        /// <summary>
        /// Uin of the ccn attached. If not set, which means the uin of this account. This parameter is used with case when
        /// attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        /// </summary>
        [Output("ccnUin")]
        public Output<string> CcnUin { get; private set; } = null!;

        /// <summary>
        /// A network address block of the instance that is attached.
        /// </summary>
        [Output("cidrBlocks")]
        public Output<ImmutableArray<string>> CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// Remark of attachment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of instance is attached.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The region that the instance locates at.
        /// </summary>
        [Output("instanceRegion")]
        public Output<string> InstanceRegion { get; private set; } = null!;

        /// <summary>
        /// Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note:
        /// `VPNGW` type is only for whitelist customer now.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Route id list.
        /// </summary>
        [Output("routeIds")]
        public Output<ImmutableArray<string>> RouteIds { get; private set; } = null!;

        /// <summary>
        /// Ccn instance route table ID.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// States of instance is attached. Valid values: `PENDING`, `ACTIVE`, `EXPIRED`, `REJECTED`, `DELETED`, `FAILED`,
        /// `ATTACHING`, `DETACHING` and `DETACHFAILED`. `FAILED` means asynchronous forced disassociation after 2 hours.
        /// `DETACHFAILED` means asynchronous forced disassociation after 2 hours.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Attachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Attachment(string name, AttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/attachment:Attachment", name, args ?? new AttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Attachment(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/attachment:Attachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Attachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Attachment Get(string name, Input<string> id, AttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new Attachment(name, id, state, options);
        }
    }

    public sealed class AttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the CCN.
        /// </summary>
        [Input("ccnId", required: true)]
        public Input<string> CcnId { get; set; } = null!;

        /// <summary>
        /// Uin of the ccn attached. If not set, which means the uin of this account. This parameter is used with case when
        /// attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        /// </summary>
        [Input("ccnUin")]
        public Input<string>? CcnUin { get; set; }

        /// <summary>
        /// Remark of attachment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of instance is attached.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The region that the instance locates at.
        /// </summary>
        [Input("instanceRegion", required: true)]
        public Input<string> InstanceRegion { get; set; } = null!;

        /// <summary>
        /// Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note:
        /// `VPNGW` type is only for whitelist customer now.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Ccn instance route table ID.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        public AttachmentArgs()
        {
        }
        public static new AttachmentArgs Empty => new AttachmentArgs();
    }

    public sealed class AttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time of attaching.
        /// </summary>
        [Input("attachedTime")]
        public Input<string>? AttachedTime { get; set; }

        /// <summary>
        /// ID of the CCN.
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// Uin of the ccn attached. If not set, which means the uin of this account. This parameter is used with case when
        /// attaching ccn of other account to the instance of this account. For now only support instance type `VPC`.
        /// </summary>
        [Input("ccnUin")]
        public Input<string>? CcnUin { get; set; }

        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// A network address block of the instance that is attached.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// Remark of attachment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of instance is attached.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The region that the instance locates at.
        /// </summary>
        [Input("instanceRegion")]
        public Input<string>? InstanceRegion { get; set; }

        /// <summary>
        /// Type of attached instance network, and available values include `VPC`, `DIRECTCONNECT`, `BMVPC` and `VPNGW`. Note:
        /// `VPNGW` type is only for whitelist customer now.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("routeIds")]
        private InputList<string>? _routeIds;

        /// <summary>
        /// Route id list.
        /// </summary>
        public InputList<string> RouteIds
        {
            get => _routeIds ?? (_routeIds = new InputList<string>());
            set => _routeIds = value;
        }

        /// <summary>
        /// Ccn instance route table ID.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// States of instance is attached. Valid values: `PENDING`, `ACTIVE`, `EXPIRED`, `REJECTED`, `DELETED`, `FAILED`,
        /// `ATTACHING`, `DETACHING` and `DETACHFAILED`. `FAILED` means asynchronous forced disassociation after 2 hours.
        /// `DETACHFAILED` means asynchronous forced disassociation after 2 hours.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public AttachmentState()
        {
        }
        public static new AttachmentState Empty => new AttachmentState();
    }
}
