// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes.Inputs
{

    public sealed class NodePoolAutoScalingConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupInstanceTypes")]
        private InputList<string>? _backupInstanceTypes;
        public InputList<string> BackupInstanceTypes
        {
            get => _backupInstanceTypes ?? (_backupInstanceTypes = new InputList<string>());
            set => _backupInstanceTypes = value;
        }

        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        [Input("camRoleName")]
        public Input<string>? CamRoleName { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.NodePoolAutoScalingConfigDataDiskGetArgs>? _dataDisks;
        public InputList<Inputs.NodePoolAutoScalingConfigDataDiskGetArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.NodePoolAutoScalingConfigDataDiskGetArgs>());
            set => _dataDisks = value;
        }

        [Input("enhancedMonitorService")]
        public Input<bool>? EnhancedMonitorService { get; set; }

        [Input("enhancedSecurityService")]
        public Input<bool>? EnhancedSecurityService { get; set; }

        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        [Input("hostNameStyle")]
        public Input<string>? HostNameStyle { get; set; }

        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        [Input("instanceChargeTypePrepaidPeriod")]
        public Input<int>? InstanceChargeTypePrepaidPeriod { get; set; }

        [Input("instanceChargeTypePrepaidRenewFlag")]
        public Input<string>? InstanceChargeTypePrepaidRenewFlag { get; set; }

        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("instanceNameStyle")]
        public Input<string>? InstanceNameStyle { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        [Input("keyIds")]
        private InputList<string>? _keyIds;
        public InputList<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new InputList<string>());
            set => _keyIds = value;
        }

        [Input("orderlySecurityGroupIds")]
        private InputList<string>? _orderlySecurityGroupIds;
        public InputList<string> OrderlySecurityGroupIds
        {
            get => _orderlySecurityGroupIds ?? (_orderlySecurityGroupIds = new InputList<string>());
            set => _orderlySecurityGroupIds = value;
        }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("publicIpAssigned")]
        public Input<bool>? PublicIpAssigned { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        [Obsolete(@"The order of elements in this field cannot be guaranteed. Use `orderly_security_group_ids` instead.")]
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        [Input("spotMaxPrice")]
        public Input<string>? SpotMaxPrice { get; set; }

        [Input("systemDiskSize")]
        public Input<int>? SystemDiskSize { get; set; }

        [Input("systemDiskType")]
        public Input<string>? SystemDiskType { get; set; }

        public NodePoolAutoScalingConfigGetArgs()
        {
        }
        public static new NodePoolAutoScalingConfigGetArgs Empty => new NodePoolAutoScalingConfigGetArgs();
    }
}
