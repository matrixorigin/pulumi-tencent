// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Identity.Outputs
{

    [OutputType]
    public sealed class GetCenterRoleConfigurationsRoleConfigurationResult
    {
        public readonly string? CreateTime;
        public readonly string? Description;
        public readonly bool? IsSelected;
        public readonly string? RelayState;
        public readonly string? RoleConfigurationId;
        public readonly string? RoleConfigurationName;
        public readonly int? SessionDuration;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private GetCenterRoleConfigurationsRoleConfigurationResult(
            string? createTime,

            string? description,

            bool? isSelected,

            string? relayState,

            string? roleConfigurationId,

            string? roleConfigurationName,

            int? sessionDuration,

            string? updateTime)
        {
            CreateTime = createTime;
            Description = description;
            IsSelected = isSelected;
            RelayState = relayState;
            RoleConfigurationId = roleConfigurationId;
            RoleConfigurationName = roleConfigurationName;
            SessionDuration = sessionDuration;
            UpdateTime = updateTime;
        }
    }
}
