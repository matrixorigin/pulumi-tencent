// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse
{
    [TencentcloudResourceType("tencentcloud:Lighthouse/disk:Disk")]
    public partial class Disk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Automatically mount and initialize data disks.
        /// </summary>
        [Output("autoMountConfiguration")]
        public Output<Outputs.DiskAutoMountConfiguration?> AutoMountConfiguration { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically use the voucher. Not used by default.
        /// </summary>
        [Output("autoVoucher")]
        public Output<bool> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Specify the disk backup quota. If not uploaded, the default is no backup quota. Currently, only one disk backup quota is
        /// supported.
        /// </summary>
        [Output("diskBackupQuota")]
        public Output<int> DiskBackupQuota { get; private set; } = null!;

        /// <summary>
        /// Disk subscription related parameter settings.
        /// </summary>
        [Output("diskChargePrepaid")]
        public Output<Outputs.DiskDiskChargePrepaid> DiskChargePrepaid { get; private set; } = null!;

        /// <summary>
        /// Disk count. Values: [1, 30]. Default: 1.
        /// </summary>
        [Output("diskCount")]
        public Output<int> DiskCount { get; private set; } = null!;

        /// <summary>
        /// Disk name. Maximum length 60.
        /// </summary>
        [Output("diskName")]
        public Output<string> DiskName { get; private set; } = null!;

        /// <summary>
        /// Disk size, unit: GB.
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// Disk type. Value:CLOUD_PREMIUM, CLOUD_SSD.
        /// </summary>
        [Output("diskType")]
        public Output<string> DiskType { get; private set; } = null!;

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Lighthouse/disk:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Lighthouse/disk:Disk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, DiskState? state = null, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, state, options);
        }
    }

    public sealed class DiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatically mount and initialize data disks.
        /// </summary>
        [Input("autoMountConfiguration")]
        public Input<Inputs.DiskAutoMountConfigurationArgs>? AutoMountConfiguration { get; set; }

        /// <summary>
        /// Whether to automatically use the voucher. Not used by default.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        /// <summary>
        /// Specify the disk backup quota. If not uploaded, the default is no backup quota. Currently, only one disk backup quota is
        /// supported.
        /// </summary>
        [Input("diskBackupQuota")]
        public Input<int>? DiskBackupQuota { get; set; }

        /// <summary>
        /// Disk subscription related parameter settings.
        /// </summary>
        [Input("diskChargePrepaid", required: true)]
        public Input<Inputs.DiskDiskChargePrepaidArgs> DiskChargePrepaid { get; set; } = null!;

        /// <summary>
        /// Disk count. Values: [1, 30]. Default: 1.
        /// </summary>
        [Input("diskCount")]
        public Input<int>? DiskCount { get; set; }

        /// <summary>
        /// Disk name. Maximum length 60.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// Disk size, unit: GB.
        /// </summary>
        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        /// <summary>
        /// Disk type. Value:CLOUD_PREMIUM, CLOUD_SSD.
        /// </summary>
        [Input("diskType", required: true)]
        public Input<string> DiskType { get; set; } = null!;

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public DiskArgs()
        {
        }
        public static new DiskArgs Empty => new DiskArgs();
    }

    public sealed class DiskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Automatically mount and initialize data disks.
        /// </summary>
        [Input("autoMountConfiguration")]
        public Input<Inputs.DiskAutoMountConfigurationGetArgs>? AutoMountConfiguration { get; set; }

        /// <summary>
        /// Whether to automatically use the voucher. Not used by default.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        /// <summary>
        /// Specify the disk backup quota. If not uploaded, the default is no backup quota. Currently, only one disk backup quota is
        /// supported.
        /// </summary>
        [Input("diskBackupQuota")]
        public Input<int>? DiskBackupQuota { get; set; }

        /// <summary>
        /// Disk subscription related parameter settings.
        /// </summary>
        [Input("diskChargePrepaid")]
        public Input<Inputs.DiskDiskChargePrepaidGetArgs>? DiskChargePrepaid { get; set; }

        /// <summary>
        /// Disk count. Values: [1, 30]. Default: 1.
        /// </summary>
        [Input("diskCount")]
        public Input<int>? DiskCount { get; set; }

        /// <summary>
        /// Disk name. Maximum length 60.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// Disk size, unit: GB.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Disk type. Value:CLOUD_PREMIUM, CLOUD_SSD.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DiskState()
        {
        }
        public static new DiskState Empty => new DiskState();
    }
}