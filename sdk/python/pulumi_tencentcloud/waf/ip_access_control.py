# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['IpAccessControlArgs', 'IpAccessControl']

@pulumi.input_type
class IpAccessControlArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 edition: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 items: pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]]):
        """
        The set of arguments for constructing a IpAccessControl resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] instance_id: Waf instance Id.
        :param pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]] items: Ip parameter list.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "edition", edition)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Input[str]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: pulumi.Input[str]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Waf instance Id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def items(self) -> pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]]:
        """
        Ip parameter list.
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]]):
        pulumi.set(self, "items", value)


@pulumi.input_type
class _IpAccessControlState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]]] = None):
        """
        Input properties used for looking up and filtering IpAccessControl resources.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] instance_id: Waf instance Id.
        :param pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]] items: Ip parameter list.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if items is not None:
            pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Waf instance Id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]]]:
        """
        Ip parameter list.
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpAccessControlItemArgs']]]]):
        pulumi.set(self, "items", value)


class IpAccessControl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAccessControlItemArgs']]]]] = None,
                 __props__=None):
        """
        Create a IpAccessControl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] instance_id: Waf instance Id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAccessControlItemArgs']]]] items: Ip parameter list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpAccessControlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IpAccessControl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IpAccessControlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpAccessControlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAccessControlItemArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpAccessControlArgs.__new__(IpAccessControlArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if edition is None and not opts.urn:
                raise TypeError("Missing required property 'edition'")
            __props__.__dict__["edition"] = edition
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if items is None and not opts.urn:
                raise TypeError("Missing required property 'items'")
            __props__.__dict__["items"] = items
        super(IpAccessControl, __self__).__init__(
            'tencentcloud:Waf/ipAccessControl:IpAccessControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            items: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAccessControlItemArgs']]]]] = None) -> 'IpAccessControl':
        """
        Get an existing IpAccessControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] instance_id: Waf instance Id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpAccessControlItemArgs']]]] items: Ip parameter list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpAccessControlState.__new__(_IpAccessControlState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["edition"] = edition
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["items"] = items
        return IpAccessControl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Waf instance Id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def items(self) -> pulumi.Output[Sequence['outputs.IpAccessControlItem']]:
        """
        Ip parameter list.
        """
        return pulumi.get(self, "items")
