# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RocketmqRoleArgs', 'RocketmqRole']

@pulumi.input_type
class RocketmqRoleArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 perm_read: pulumi.Input[bool],
                 perm_write: pulumi.Input[bool],
                 remark: pulumi.Input[str],
                 role: pulumi.Input[str]):
        """
        The set of arguments for constructing a RocketmqRole resource.
        :param pulumi.Input[str] instance_id: ID of instance.
        :param pulumi.Input[bool] perm_read: Whether to enable consumption permission.
        :param pulumi.Input[bool] perm_write: Whether to enable production permission.
        :param pulumi.Input[str] remark: remark.
        :param pulumi.Input[str] role: Name of role.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "perm_read", perm_read)
        pulumi.set(__self__, "perm_write", perm_write)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="permRead")
    def perm_read(self) -> pulumi.Input[bool]:
        """
        Whether to enable consumption permission.
        """
        return pulumi.get(self, "perm_read")

    @perm_read.setter
    def perm_read(self, value: pulumi.Input[bool]):
        pulumi.set(self, "perm_read", value)

    @property
    @pulumi.getter(name="permWrite")
    def perm_write(self) -> pulumi.Input[bool]:
        """
        Whether to enable production permission.
        """
        return pulumi.get(self, "perm_write")

    @perm_write.setter
    def perm_write(self, value: pulumi.Input[bool]):
        pulumi.set(self, "perm_write", value)

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Input[str]:
        """
        remark.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: pulumi.Input[str]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        Name of role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class _RocketmqRoleState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 created_time: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 modified_time: Optional[pulumi.Input[int]] = None,
                 perm_read: Optional[pulumi.Input[bool]] = None,
                 perm_write: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RocketmqRole resources.
        :param pulumi.Input[str] access_key: Access key.
        :param pulumi.Input[int] created_time: Created time.
        :param pulumi.Input[str] instance_id: ID of instance.
        :param pulumi.Input[int] modified_time: Modified time.
        :param pulumi.Input[bool] perm_read: Whether to enable consumption permission.
        :param pulumi.Input[bool] perm_write: Whether to enable production permission.
        :param pulumi.Input[str] remark: remark.
        :param pulumi.Input[str] role: Name of role.
        :param pulumi.Input[str] secret_key: Secret key.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if created_time is not None:
            pulumi.set(__self__, "created_time", created_time)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if modified_time is not None:
            pulumi.set(__self__, "modified_time", modified_time)
        if perm_read is not None:
            pulumi.set(__self__, "perm_read", perm_read)
        if perm_write is not None:
            pulumi.set(__self__, "perm_write", perm_write)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Access key.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[pulumi.Input[int]]:
        """
        Created time.
        """
        return pulumi.get(self, "created_time")

    @created_time.setter
    def created_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> Optional[pulumi.Input[int]]:
        """
        Modified time.
        """
        return pulumi.get(self, "modified_time")

    @modified_time.setter
    def modified_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "modified_time", value)

    @property
    @pulumi.getter(name="permRead")
    def perm_read(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable consumption permission.
        """
        return pulumi.get(self, "perm_read")

    @perm_read.setter
    def perm_read(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "perm_read", value)

    @property
    @pulumi.getter(name="permWrite")
    def perm_write(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable production permission.
        """
        return pulumi.get(self, "perm_write")

    @perm_write.setter
    def perm_write(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "perm_write", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        remark.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Name of role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret key.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)


class RocketmqRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 perm_read: Optional[pulumi.Input[bool]] = None,
                 perm_write: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a RocketmqRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: ID of instance.
        :param pulumi.Input[bool] perm_read: Whether to enable consumption permission.
        :param pulumi.Input[bool] perm_write: Whether to enable production permission.
        :param pulumi.Input[str] remark: remark.
        :param pulumi.Input[str] role: Name of role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RocketmqRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RocketmqRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 perm_read: Optional[pulumi.Input[bool]] = None,
                 perm_write: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RocketmqRoleArgs.__new__(RocketmqRoleArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if perm_read is None and not opts.urn:
                raise TypeError("Missing required property 'perm_read'")
            __props__.__dict__["perm_read"] = perm_read
            if perm_write is None and not opts.urn:
                raise TypeError("Missing required property 'perm_write'")
            __props__.__dict__["perm_write"] = perm_write
            if remark is None and not opts.urn:
                raise TypeError("Missing required property 'remark'")
            __props__.__dict__["remark"] = remark
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["access_key"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["modified_time"] = None
            __props__.__dict__["secret_key"] = None
        super(RocketmqRole, __self__).__init__(
            'tencentcloud:Trocket/rocketmqRole:RocketmqRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            created_time: Optional[pulumi.Input[int]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            modified_time: Optional[pulumi.Input[int]] = None,
            perm_read: Optional[pulumi.Input[bool]] = None,
            perm_write: Optional[pulumi.Input[bool]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None) -> 'RocketmqRole':
        """
        Get an existing RocketmqRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access key.
        :param pulumi.Input[int] created_time: Created time.
        :param pulumi.Input[str] instance_id: ID of instance.
        :param pulumi.Input[int] modified_time: Modified time.
        :param pulumi.Input[bool] perm_read: Whether to enable consumption permission.
        :param pulumi.Input[bool] perm_write: Whether to enable production permission.
        :param pulumi.Input[str] remark: remark.
        :param pulumi.Input[str] role: Name of role.
        :param pulumi.Input[str] secret_key: Secret key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqRoleState.__new__(_RocketmqRoleState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["created_time"] = created_time
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["modified_time"] = modified_time
        __props__.__dict__["perm_read"] = perm_read
        __props__.__dict__["perm_write"] = perm_write
        __props__.__dict__["remark"] = remark
        __props__.__dict__["role"] = role
        __props__.__dict__["secret_key"] = secret_key
        return RocketmqRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        Access key.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[int]:
        """
        Created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="modifiedTime")
    def modified_time(self) -> pulumi.Output[int]:
        """
        Modified time.
        """
        return pulumi.get(self, "modified_time")

    @property
    @pulumi.getter(name="permRead")
    def perm_read(self) -> pulumi.Output[bool]:
        """
        Whether to enable consumption permission.
        """
        return pulumi.get(self, "perm_read")

    @property
    @pulumi.getter(name="permWrite")
    def perm_write(self) -> pulumi.Output[bool]:
        """
        Whether to enable production permission.
        """
        return pulumi.get(self, "perm_write")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[str]:
        """
        remark.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Name of role.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Secret key.
        """
        return pulumi.get(self, "secret_key")
