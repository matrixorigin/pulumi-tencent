# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConfigInstanceSecurityGroupsArgs', 'ConfigInstanceSecurityGroups']

@pulumi.input_type
class ConfigInstanceSecurityGroupsArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 security_group_id_sets: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a ConfigInstanceSecurityGroups resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_id_sets: A list of security group IDs to modify, an array of one or more security group IDs.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "security_group_id_sets", security_group_id_sets)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="securityGroupIdSets")
    def security_group_id_sets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of security group IDs to modify, an array of one or more security group IDs.
        """
        return pulumi.get(self, "security_group_id_sets")

    @security_group_id_sets.setter
    def security_group_id_sets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_id_sets", value)


@pulumi.input_type
class _ConfigInstanceSecurityGroupsState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 security_group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ConfigInstanceSecurityGroups resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_id_sets: A list of security group IDs to modify, an array of one or more security group IDs.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if security_group_id_sets is not None:
            pulumi.set(__self__, "security_group_id_sets", security_group_id_sets)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="securityGroupIdSets")
    def security_group_id_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of security group IDs to modify, an array of one or more security group IDs.
        """
        return pulumi.get(self, "security_group_id_sets")

    @security_group_id_sets.setter
    def security_group_id_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_id_sets", value)


class ConfigInstanceSecurityGroups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 security_group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a ConfigInstanceSecurityGroups resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_id_sets: A list of security group IDs to modify, an array of one or more security group IDs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigInstanceSecurityGroupsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ConfigInstanceSecurityGroups resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ConfigInstanceSecurityGroupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigInstanceSecurityGroupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 security_group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigInstanceSecurityGroupsArgs.__new__(ConfigInstanceSecurityGroupsArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if security_group_id_sets is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id_sets'")
            __props__.__dict__["security_group_id_sets"] = security_group_id_sets
        super(ConfigInstanceSecurityGroups, __self__).__init__(
            'tencentcloud:Sqlserver/configInstanceSecurityGroups:ConfigInstanceSecurityGroups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            security_group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ConfigInstanceSecurityGroups':
        """
        Get an existing ConfigInstanceSecurityGroups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_id_sets: A list of security group IDs to modify, an array of one or more security group IDs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigInstanceSecurityGroupsState.__new__(_ConfigInstanceSecurityGroupsState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["security_group_id_sets"] = security_group_id_sets
        return ConfigInstanceSecurityGroups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="securityGroupIdSets")
    def security_group_id_sets(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of security group IDs to modify, an array of one or more security group IDs.
        """
        return pulumi.get(self, "security_group_id_sets")
