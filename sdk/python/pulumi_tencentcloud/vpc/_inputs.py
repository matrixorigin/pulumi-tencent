# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'FlowLogFlowLogStorageArgs',
    'Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksArgs',
    'NetworkAclQuintupleNetworkAclQuintupleSetArgs',
    'NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs',
    'NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs',
    'SnapshotPolicyAttachmentInstanceArgs',
    'SnapshotPolicyBackupPolicyArgs',
    'GetClassicLinkInstancesFilterArgs',
    'GetCvmInstancesFilterArgs',
    'GetNetDetectStatesFilterArgs',
]

@pulumi.input_type
class FlowLogFlowLogStorageArgs:
    def __init__(__self__, *,
                 storage_id: Optional[pulumi.Input[str]] = None,
                 storage_topic: Optional[pulumi.Input[str]] = None):
        if storage_id is not None:
            pulumi.set(__self__, "storage_id", storage_id)
        if storage_topic is not None:
            pulumi.set(__self__, "storage_topic", storage_topic)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_id", value)

    @property
    @pulumi.getter(name="storageTopic")
    def storage_topic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "storage_topic")

    @storage_topic.setter
    def storage_topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_topic", value)


@pulumi.input_type
class Ipv6SubnetCidrBlockIpv6SubnetCidrBlocksArgs:
    def __init__(__self__, *,
                 ipv6_cidr_block: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ipv6_cidr_block")

    @ipv6_cidr_block.setter
    def ipv6_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class NetworkAclQuintupleNetworkAclQuintupleSetArgs:
    def __init__(__self__, *,
                 egresses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs']]]] = None,
                 ingresses: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs']]]] = None):
        if egresses is not None:
            pulumi.set(__self__, "egresses", egresses)
        if ingresses is not None:
            pulumi.set(__self__, "ingresses", ingresses)

    @property
    @pulumi.getter
    def egresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs']]]]:
        return pulumi.get(self, "egresses")

    @egresses.setter
    def egresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs']]]]):
        pulumi.set(self, "egresses", value)

    @property
    @pulumi.getter
    def ingresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs']]]]:
        return pulumi.get(self, "ingresses")

    @ingresses.setter
    def ingresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs']]]]):
        pulumi.set(self, "ingresses", value)


@pulumi.input_type
class NetworkAclQuintupleNetworkAclQuintupleSetEgressArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[str]] = None,
                 network_acl_direction: Optional[pulumi.Input[str]] = None,
                 network_acl_quintuple_entry_id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 source_cidr: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[str]] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_cidr is not None:
            pulumi.set(__self__, "destination_cidr", destination_cidr)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if network_acl_direction is not None:
            pulumi.set(__self__, "network_acl_direction", network_acl_direction)
        if network_acl_quintuple_entry_id is not None:
            pulumi.set(__self__, "network_acl_quintuple_entry_id", network_acl_quintuple_entry_id)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if source_cidr is not None:
            pulumi.set(__self__, "source_cidr", source_cidr)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_cidr")

    @destination_cidr.setter
    def destination_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter(name="networkAclDirection")
    def network_acl_direction(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_acl_direction")

    @network_acl_direction.setter
    def network_acl_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_direction", value)

    @property
    @pulumi.getter(name="networkAclQuintupleEntryId")
    def network_acl_quintuple_entry_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_acl_quintuple_entry_id")

    @network_acl_quintuple_entry_id.setter
    def network_acl_quintuple_entry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_quintuple_entry_id", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="sourceCidr")
    def source_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_cidr")

    @source_cidr.setter
    def source_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_cidr", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_port", value)


@pulumi.input_type
class NetworkAclQuintupleNetworkAclQuintupleSetIngressArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[str]] = None,
                 network_acl_direction: Optional[pulumi.Input[str]] = None,
                 network_acl_quintuple_entry_id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 source_cidr: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[str]] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_cidr is not None:
            pulumi.set(__self__, "destination_cidr", destination_cidr)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if network_acl_direction is not None:
            pulumi.set(__self__, "network_acl_direction", network_acl_direction)
        if network_acl_quintuple_entry_id is not None:
            pulumi.set(__self__, "network_acl_quintuple_entry_id", network_acl_quintuple_entry_id)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if source_cidr is not None:
            pulumi.set(__self__, "source_cidr", source_cidr)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_cidr")

    @destination_cidr.setter
    def destination_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter(name="networkAclDirection")
    def network_acl_direction(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_acl_direction")

    @network_acl_direction.setter
    def network_acl_direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_direction", value)

    @property
    @pulumi.getter(name="networkAclQuintupleEntryId")
    def network_acl_quintuple_entry_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_acl_quintuple_entry_id")

    @network_acl_quintuple_entry_id.setter
    def network_acl_quintuple_entry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_quintuple_entry_id", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="sourceCidr")
    def source_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_cidr")

    @source_cidr.setter
    def source_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_cidr", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_port", value)


@pulumi.input_type
class SnapshotPolicyAttachmentInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 instance_region: pulumi.Input[str],
                 instance_type: pulumi.Input[str],
                 instance_name: Optional[pulumi.Input[str]] = None,
                 snapshot_policy_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        pulumi.set(__self__, "instance_type", instance_type)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if snapshot_policy_id is not None:
            pulumi.set(__self__, "snapshot_policy_id", snapshot_policy_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_region")

    @instance_region.setter
    def instance_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_region", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "snapshot_policy_id")

    @snapshot_policy_id.setter
    def snapshot_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_policy_id", value)


@pulumi.input_type
class SnapshotPolicyBackupPolicyArgs:
    def __init__(__self__, *,
                 backup_day: pulumi.Input[str],
                 backup_time: pulumi.Input[str]):
        pulumi.set(__self__, "backup_day", backup_day)
        pulumi.set(__self__, "backup_time", backup_time)

    @property
    @pulumi.getter(name="backupDay")
    def backup_day(self) -> pulumi.Input[str]:
        return pulumi.get(self, "backup_day")

    @backup_day.setter
    def backup_day(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_day", value)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> pulumi.Input[str]:
        return pulumi.get(self, "backup_time")

    @backup_time.setter
    def backup_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_time", value)


@pulumi.input_type
class GetClassicLinkInstancesFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GetCvmInstancesFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GetNetDetectStatesFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


