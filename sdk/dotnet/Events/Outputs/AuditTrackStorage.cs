// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Events.Outputs
{

    [OutputType]
    public sealed class AuditTrackStorage
    {
        public readonly string? StorageAccountId;
        public readonly string? StorageAppId;
        public readonly string StorageName;
        public readonly string StoragePrefix;
        public readonly string StorageRegion;
        public readonly string StorageType;

        [OutputConstructor]
        private AuditTrackStorage(
            string? storageAccountId,

            string? storageAppId,

            string storageName,

            string storagePrefix,

            string storageRegion,

            string storageType)
        {
            StorageAccountId = storageAccountId;
            StorageAppId = storageAppId;
            StorageName = storageName;
            StoragePrefix = storagePrefix;
            StorageRegion = storageRegion;
            StorageType = storageType;
        }
    }
}