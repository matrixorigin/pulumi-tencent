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

__all__ = [
    'GetGatewaysResult',
    'AwaitableGetGatewaysResult',
    'get_gateways',
    'get_gateways_output',
]

@pulumi.output_type
class GetGatewaysResult:
    """
    A collection of values returned by getGateways.
    """
    def __init__(__self__, gateway_lists=None, id=None, name=None, public_ip_address=None, result_output_file=None, tags=None, vpc_id=None, zone=None):
        if gateway_lists and not isinstance(gateway_lists, list):
            raise TypeError("Expected argument 'gateway_lists' to be a list")
        pulumi.set(__self__, "gateway_lists", gateway_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if public_ip_address and not isinstance(public_ip_address, str):
            raise TypeError("Expected argument 'public_ip_address' to be a str")
        pulumi.set(__self__, "public_ip_address", public_ip_address)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="gatewayLists")
    def gateway_lists(self) -> Sequence['outputs.GetGatewaysGatewayListResult']:
        return pulumi.get(self, "gateway_lists")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> Optional[str]:
        return pulumi.get(self, "public_ip_address")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetGatewaysResult(GetGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewaysResult(
            gateway_lists=self.gateway_lists,
            id=self.id,
            name=self.name,
            public_ip_address=self.public_ip_address,
            result_output_file=self.result_output_file,
            tags=self.tags,
            vpc_id=self.vpc_id,
            zone=self.zone)


def get_gateways(id: Optional[str] = None,
                 name: Optional[str] = None,
                 public_ip_address: Optional[str] = None,
                 result_output_file: Optional[str] = None,
                 tags: Optional[Mapping[str, Any]] = None,
                 vpc_id: Optional[str] = None,
                 zone: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewaysResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['publicIpAddress'] = public_ip_address
    __args__['resultOutputFile'] = result_output_file
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpn/getGateways:getGateways', __args__, opts=opts, typ=GetGatewaysResult).value

    return AwaitableGetGatewaysResult(
        gateway_lists=pulumi.get(__ret__, 'gateway_lists'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        public_ip_address=pulumi.get(__ret__, 'public_ip_address'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        tags=pulumi.get(__ret__, 'tags'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_gateways)
def get_gateways_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        public_ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                        vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                        zone: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewaysResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...