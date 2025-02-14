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

__all__ = ['RocketmqVipInstanceArgs', 'RocketmqVipInstance']

@pulumi.input_type
class RocketmqVipInstanceArgs:
    def __init__(__self__, *,
                 node_count: pulumi.Input[int],
                 spec: pulumi.Input[str],
                 storage_size: pulumi.Input[int],
                 time_span: pulumi.Input[int],
                 vpc_info: pulumi.Input['RocketmqVipInstanceVpcInfoArgs'],
                 zone_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RocketmqVipInstance resource.
        :param pulumi.Input[int] node_count: Number of nodes, minimum 2, maximum 20.
        :param pulumi.Input[str] spec: Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
               `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        :param pulumi.Input[int] storage_size: Single node storage space, in GB, minimum 200GB.
        :param pulumi.Input[int] time_span: Purchase period, in months.
        :param pulumi.Input['RocketmqVipInstanceVpcInfoArgs'] vpc_info: VPC information.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
               website of Tencent Cloud.
        :param pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]] ip_rules: Public IP access control rules.
        :param pulumi.Input[str] name: Instance name.
        """
        pulumi.set(__self__, "node_count", node_count)
        pulumi.set(__self__, "spec", spec)
        pulumi.set(__self__, "storage_size", storage_size)
        pulumi.set(__self__, "time_span", time_span)
        pulumi.set(__self__, "vpc_info", vpc_info)
        pulumi.set(__self__, "zone_ids", zone_ids)
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Input[int]:
        """
        Number of nodes, minimum 2, maximum 20.
        """
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Input[str]:
        """
        Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
        `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: pulumi.Input[str]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> pulumi.Input[int]:
        """
        Single node storage space, in GB, minimum 200GB.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> pulumi.Input[int]:
        """
        Purchase period, in months.
        """
        return pulumi.get(self, "time_span")

    @time_span.setter
    def time_span(self, value: pulumi.Input[int]):
        pulumi.set(self, "time_span", value)

    @property
    @pulumi.getter(name="vpcInfo")
    def vpc_info(self) -> pulumi.Input['RocketmqVipInstanceVpcInfoArgs']:
        """
        VPC information.
        """
        return pulumi.get(self, "vpc_info")

    @vpc_info.setter
    def vpc_info(self, value: pulumi.Input['RocketmqVipInstanceVpcInfoArgs']):
        pulumi.set(self, "vpc_info", value)

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
        website of Tencent Cloud.
        """
        return pulumi.get(self, "zone_ids")

    @zone_ids.setter
    def zone_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "zone_ids", value)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]]]:
        """
        Public IP access control rules.
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RocketmqVipInstanceState:
    def __init__(__self__, *,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 vpc_info: Optional[pulumi.Input['RocketmqVipInstanceVpcInfoArgs']] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RocketmqVipInstance resources.
        :param pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]] ip_rules: Public IP access control rules.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] node_count: Number of nodes, minimum 2, maximum 20.
        :param pulumi.Input[str] spec: Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
               `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        :param pulumi.Input[int] storage_size: Single node storage space, in GB, minimum 200GB.
        :param pulumi.Input[int] time_span: Purchase period, in months.
        :param pulumi.Input['RocketmqVipInstanceVpcInfoArgs'] vpc_info: VPC information.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
               website of Tencent Cloud.
        """
        if ip_rules is not None:
            pulumi.set(__self__, "ip_rules", ip_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_count is not None:
            pulumi.set(__self__, "node_count", node_count)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if storage_size is not None:
            pulumi.set(__self__, "storage_size", storage_size)
        if time_span is not None:
            pulumi.set(__self__, "time_span", time_span)
        if vpc_info is not None:
            pulumi.set(__self__, "vpc_info", vpc_info)
        if zone_ids is not None:
            pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]]]:
        """
        Public IP access control rules.
        """
        return pulumi.get(self, "ip_rules")

    @ip_rules.setter
    def ip_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqVipInstanceIpRuleArgs']]]]):
        pulumi.set(self, "ip_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of nodes, minimum 2, maximum 20.
        """
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[str]]:
        """
        Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
        `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> Optional[pulumi.Input[int]]:
        """
        Single node storage space, in GB, minimum 200GB.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase period, in months.
        """
        return pulumi.get(self, "time_span")

    @time_span.setter
    def time_span(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_span", value)

    @property
    @pulumi.getter(name="vpcInfo")
    def vpc_info(self) -> Optional[pulumi.Input['RocketmqVipInstanceVpcInfoArgs']]:
        """
        VPC information.
        """
        return pulumi.get(self, "vpc_info")

    @vpc_info.setter
    def vpc_info(self, value: Optional[pulumi.Input['RocketmqVipInstanceVpcInfoArgs']]):
        pulumi.set(self, "vpc_info", value)

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
        website of Tencent Cloud.
        """
        return pulumi.get(self, "zone_ids")

    @zone_ids.setter
    def zone_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "zone_ids", value)


class RocketmqVipInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqVipInstanceIpRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 vpc_info: Optional[pulumi.Input[pulumi.InputType['RocketmqVipInstanceVpcInfoArgs']]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a RocketmqVipInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqVipInstanceIpRuleArgs']]]] ip_rules: Public IP access control rules.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] node_count: Number of nodes, minimum 2, maximum 20.
        :param pulumi.Input[str] spec: Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
               `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        :param pulumi.Input[int] storage_size: Single node storage space, in GB, minimum 200GB.
        :param pulumi.Input[int] time_span: Purchase period, in months.
        :param pulumi.Input[pulumi.InputType['RocketmqVipInstanceVpcInfoArgs']] vpc_info: VPC information.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
               website of Tencent Cloud.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqVipInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RocketmqVipInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RocketmqVipInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqVipInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqVipInstanceIpRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[int]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 time_span: Optional[pulumi.Input[int]] = None,
                 vpc_info: Optional[pulumi.Input[pulumi.InputType['RocketmqVipInstanceVpcInfoArgs']]] = None,
                 zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RocketmqVipInstanceArgs.__new__(RocketmqVipInstanceArgs)

            __props__.__dict__["ip_rules"] = ip_rules
            __props__.__dict__["name"] = name
            if node_count is None and not opts.urn:
                raise TypeError("Missing required property 'node_count'")
            __props__.__dict__["node_count"] = node_count
            if spec is None and not opts.urn:
                raise TypeError("Missing required property 'spec'")
            __props__.__dict__["spec"] = spec
            if storage_size is None and not opts.urn:
                raise TypeError("Missing required property 'storage_size'")
            __props__.__dict__["storage_size"] = storage_size
            if time_span is None and not opts.urn:
                raise TypeError("Missing required property 'time_span'")
            __props__.__dict__["time_span"] = time_span
            if vpc_info is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_info'")
            __props__.__dict__["vpc_info"] = vpc_info
            if zone_ids is None and not opts.urn:
                raise TypeError("Missing required property 'zone_ids'")
            __props__.__dict__["zone_ids"] = zone_ids
        super(RocketmqVipInstance, __self__).__init__(
            'tencentcloud:Tdmq/rocketmqVipInstance:RocketmqVipInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqVipInstanceIpRuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_count: Optional[pulumi.Input[int]] = None,
            spec: Optional[pulumi.Input[str]] = None,
            storage_size: Optional[pulumi.Input[int]] = None,
            time_span: Optional[pulumi.Input[int]] = None,
            vpc_info: Optional[pulumi.Input[pulumi.InputType['RocketmqVipInstanceVpcInfoArgs']]] = None,
            zone_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'RocketmqVipInstance':
        """
        Get an existing RocketmqVipInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqVipInstanceIpRuleArgs']]]] ip_rules: Public IP access control rules.
        :param pulumi.Input[str] name: Instance name.
        :param pulumi.Input[int] node_count: Number of nodes, minimum 2, maximum 20.
        :param pulumi.Input[str] spec: Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
               `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        :param pulumi.Input[int] storage_size: Single node storage space, in GB, minimum 200GB.
        :param pulumi.Input[int] time_span: Purchase period, in months.
        :param pulumi.Input[pulumi.InputType['RocketmqVipInstanceVpcInfoArgs']] vpc_info: VPC information.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_ids: The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
               website of Tencent Cloud.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqVipInstanceState.__new__(_RocketmqVipInstanceState)

        __props__.__dict__["ip_rules"] = ip_rules
        __props__.__dict__["name"] = name
        __props__.__dict__["node_count"] = node_count
        __props__.__dict__["spec"] = spec
        __props__.__dict__["storage_size"] = storage_size
        __props__.__dict__["time_span"] = time_span
        __props__.__dict__["vpc_info"] = vpc_info
        __props__.__dict__["zone_ids"] = zone_ids
        return RocketmqVipInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipRules")
    def ip_rules(self) -> pulumi.Output[Optional[Sequence['outputs.RocketmqVipInstanceIpRule']]]:
        """
        Public IP access control rules.
        """
        return pulumi.get(self, "ip_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[int]:
        """
        Number of nodes, minimum 2, maximum 20.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[str]:
        """
        Instance specification: Universal type, rocket-vip-basic-0, Basic type: `rocket-vip-basic-1`, Standard type:
        `rocket-vip-basic-2`, Advanced Type I: `rocket-vip-basic-3`, Advanced Type II: `rocket-vip-basic-4`.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> pulumi.Output[int]:
        """
        Single node storage space, in GB, minimum 200GB.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> pulumi.Output[int]:
        """
        Purchase period, in months.
        """
        return pulumi.get(self, "time_span")

    @property
    @pulumi.getter(name="vpcInfo")
    def vpc_info(self) -> pulumi.Output['outputs.RocketmqVipInstanceVpcInfo']:
        """
        VPC information.
        """
        return pulumi.get(self, "vpc_info")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The Zone ID list for node deployment, such as Guangzhou Zone 1, is 100001. For details, please refer to the official
        website of Tencent Cloud.
        """
        return pulumi.get(self, "zone_ids")

