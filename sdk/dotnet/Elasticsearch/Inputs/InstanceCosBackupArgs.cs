// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Elasticsearch.Inputs
{

    public sealed class InstanceCosBackupArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupTime", required: true)]
        public Input<string> BackupTime { get; set; } = null!;

        [Input("isAutoBackup", required: true)]
        public Input<bool> IsAutoBackup { get; set; } = null!;

        public InstanceCosBackupArgs()
        {
        }
        public static new InstanceCosBackupArgs Empty => new InstanceCosBackupArgs();
    }
}
