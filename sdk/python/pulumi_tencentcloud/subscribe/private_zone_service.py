# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateZoneServiceArgs', 'PrivateZoneService']

@pulumi.input_type
class PrivateZoneServiceArgs:
    def __init__(__self__):
        """
        The set of arguments for constructing a PrivateZoneService resource.
        """
        pass


@pulumi.input_type
class _PrivateZoneServiceState:
    def __init__(__self__, *,
                 service_status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateZoneService resources.
        :param pulumi.Input[str] service_status: Private domain resolution service activation status.
        """
        if service_status is not None:
            pulumi.set(__self__, "service_status", service_status)

    @property
    @pulumi.getter(name="serviceStatus")
    def service_status(self) -> Optional[pulumi.Input[str]]:
        """
        Private domain resolution service activation status.
        """
        return pulumi.get(self, "service_status")

    @service_status.setter
    def service_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_status", value)


class PrivateZoneService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None):
        """
        Create a PrivateZoneService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PrivateZoneServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PrivateZoneService resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PrivateZoneServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateZoneServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateZoneServiceArgs.__new__(PrivateZoneServiceArgs)

            __props__.__dict__["service_status"] = None
        super(PrivateZoneService, __self__).__init__(
            'tencentcloud:Subscribe/privateZoneService:PrivateZoneService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            service_status: Optional[pulumi.Input[str]] = None) -> 'PrivateZoneService':
        """
        Get an existing PrivateZoneService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] service_status: Private domain resolution service activation status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateZoneServiceState.__new__(_PrivateZoneServiceState)

        __props__.__dict__["service_status"] = service_status
        return PrivateZoneService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="serviceStatus")
    def service_status(self) -> pulumi.Output[str]:
        """
        Private domain resolution service activation status.
        """
        return pulumi.get(self, "service_status")

