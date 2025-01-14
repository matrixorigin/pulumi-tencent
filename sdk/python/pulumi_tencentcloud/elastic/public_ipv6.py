# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PublicIpv6Args', 'PublicIpv6']

@pulumi.input_type
class PublicIpv6Args:
    def __init__(__self__, *,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 address_name: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_service_provider: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a PublicIpv6 resource.
        :param pulumi.Input[str] address_ip: External network IP address.
        :param pulumi.Input[str] address_name: EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        :param pulumi.Input[str] address_type: Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
               the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
               that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        :param pulumi.Input[str] egress: Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
               CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
               contact the product for clarification Default value: CENTER_EGRESS_1.
        :param pulumi.Input[str] internet_charge_type: Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
               TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        :param pulumi.Input[int] internet_max_bandwidth_out: Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
               BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        :param pulumi.Input[str] internet_service_provider: Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
               values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
               some regions.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags.
        """
        if address_ip is not None:
            pulumi.set(__self__, "address_ip", address_ip)
        if address_name is not None:
            pulumi.set(__self__, "address_name", address_name)
        if address_type is not None:
            pulumi.set(__self__, "address_type", address_type)
        if bandwidth_package_id is not None:
            pulumi.set(__self__, "bandwidth_package_id", bandwidth_package_id)
        if egress is not None:
            pulumi.set(__self__, "egress", egress)
        if internet_charge_type is not None:
            pulumi.set(__self__, "internet_charge_type", internet_charge_type)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if internet_service_provider is not None:
            pulumi.set(__self__, "internet_service_provider", internet_service_provider)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> Optional[pulumi.Input[str]]:
        """
        External network IP address.
        """
        return pulumi.get(self, "address_ip")

    @address_ip.setter
    def address_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_ip", value)

    @property
    @pulumi.getter(name="addressName")
    def address_name(self) -> Optional[pulumi.Input[str]]:
        """
        EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        """
        return pulumi.get(self, "address_name")

    @address_name.setter
    def address_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_name", value)

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
        the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        """
        return pulumi.get(self, "address_type")

    @address_type.setter
    def address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_type", value)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
        that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @bandwidth_package_id.setter
    def bandwidth_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_id", value)

    @property
    @pulumi.getter
    def egress(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
        CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
        contact the product for clarification Default value: CENTER_EGRESS_1.
        """
        return pulumi.get(self, "egress")

    @egress.setter
    def egress(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "egress", value)

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
        TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        """
        return pulumi.get(self, "internet_charge_type")

    @internet_charge_type.setter
    def internet_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_charge_type", value)

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[pulumi.Input[int]]:
        """
        Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
        BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @internet_max_bandwidth_out.setter
    def internet_max_bandwidth_out(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth_out", value)

    @property
    @pulumi.getter(name="internetServiceProvider")
    def internet_service_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
        values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
        some regions.
        """
        return pulumi.get(self, "internet_service_provider")

    @internet_service_provider.setter
    def internet_service_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_service_provider", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _PublicIpv6State:
    def __init__(__self__, *,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 address_name: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_service_provider: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering PublicIpv6 resources.
        :param pulumi.Input[str] address_ip: External network IP address.
        :param pulumi.Input[str] address_name: EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        :param pulumi.Input[str] address_type: Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
               the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
               that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        :param pulumi.Input[str] egress: Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
               CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
               contact the product for clarification Default value: CENTER_EGRESS_1.
        :param pulumi.Input[str] internet_charge_type: Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
               TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        :param pulumi.Input[int] internet_max_bandwidth_out: Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
               BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        :param pulumi.Input[str] internet_service_provider: Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
               values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
               some regions.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags.
        """
        if address_ip is not None:
            pulumi.set(__self__, "address_ip", address_ip)
        if address_name is not None:
            pulumi.set(__self__, "address_name", address_name)
        if address_type is not None:
            pulumi.set(__self__, "address_type", address_type)
        if bandwidth_package_id is not None:
            pulumi.set(__self__, "bandwidth_package_id", bandwidth_package_id)
        if egress is not None:
            pulumi.set(__self__, "egress", egress)
        if internet_charge_type is not None:
            pulumi.set(__self__, "internet_charge_type", internet_charge_type)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if internet_service_provider is not None:
            pulumi.set(__self__, "internet_service_provider", internet_service_provider)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> Optional[pulumi.Input[str]]:
        """
        External network IP address.
        """
        return pulumi.get(self, "address_ip")

    @address_ip.setter
    def address_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_ip", value)

    @property
    @pulumi.getter(name="addressName")
    def address_name(self) -> Optional[pulumi.Input[str]]:
        """
        EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        """
        return pulumi.get(self, "address_name")

    @address_name.setter
    def address_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_name", value)

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
        the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        """
        return pulumi.get(self, "address_type")

    @address_type.setter
    def address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_type", value)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
        that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @bandwidth_package_id.setter
    def bandwidth_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_id", value)

    @property
    @pulumi.getter
    def egress(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
        CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
        contact the product for clarification Default value: CENTER_EGRESS_1.
        """
        return pulumi.get(self, "egress")

    @egress.setter
    def egress(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "egress", value)

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
        TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        """
        return pulumi.get(self, "internet_charge_type")

    @internet_charge_type.setter
    def internet_charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_charge_type", value)

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[pulumi.Input[int]]:
        """
        Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
        BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @internet_max_bandwidth_out.setter
    def internet_max_bandwidth_out(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth_out", value)

    @property
    @pulumi.getter(name="internetServiceProvider")
    def internet_service_provider(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
        values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
        some regions.
        """
        return pulumi.get(self, "internet_service_provider")

    @internet_service_provider.setter
    def internet_service_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_service_provider", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class PublicIpv6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 address_name: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_service_provider: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Create a PublicIpv6 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_ip: External network IP address.
        :param pulumi.Input[str] address_name: EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        :param pulumi.Input[str] address_type: Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
               the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
               that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        :param pulumi.Input[str] egress: Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
               CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
               contact the product for clarification Default value: CENTER_EGRESS_1.
        :param pulumi.Input[str] internet_charge_type: Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
               TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        :param pulumi.Input[int] internet_max_bandwidth_out: Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
               BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        :param pulumi.Input[str] internet_service_provider: Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
               values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
               some regions.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PublicIpv6Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PublicIpv6 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PublicIpv6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PublicIpv6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_ip: Optional[pulumi.Input[str]] = None,
                 address_name: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[str]] = None,
                 internet_charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_service_provider: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PublicIpv6Args.__new__(PublicIpv6Args)

            __props__.__dict__["address_ip"] = address_ip
            __props__.__dict__["address_name"] = address_name
            __props__.__dict__["address_type"] = address_type
            __props__.__dict__["bandwidth_package_id"] = bandwidth_package_id
            __props__.__dict__["egress"] = egress
            __props__.__dict__["internet_charge_type"] = internet_charge_type
            __props__.__dict__["internet_max_bandwidth_out"] = internet_max_bandwidth_out
            __props__.__dict__["internet_service_provider"] = internet_service_provider
            __props__.__dict__["tags"] = tags
        super(PublicIpv6, __self__).__init__(
            'tencentcloud:Elastic/publicIpv6:PublicIpv6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_ip: Optional[pulumi.Input[str]] = None,
            address_name: Optional[pulumi.Input[str]] = None,
            address_type: Optional[pulumi.Input[str]] = None,
            bandwidth_package_id: Optional[pulumi.Input[str]] = None,
            egress: Optional[pulumi.Input[str]] = None,
            internet_charge_type: Optional[pulumi.Input[str]] = None,
            internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
            internet_service_provider: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'PublicIpv6':
        """
        Get an existing PublicIpv6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_ip: External network IP address.
        :param pulumi.Input[str] address_name: EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        :param pulumi.Input[str] address_type: Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
               the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
               that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        :param pulumi.Input[str] egress: Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
               CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
               contact the product for clarification Default value: CENTER_EGRESS_1.
        :param pulumi.Input[str] internet_charge_type: Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
               TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        :param pulumi.Input[int] internet_max_bandwidth_out: Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
               BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        :param pulumi.Input[str] internet_service_provider: Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
               values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
               some regions.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PublicIpv6State.__new__(_PublicIpv6State)

        __props__.__dict__["address_ip"] = address_ip
        __props__.__dict__["address_name"] = address_name
        __props__.__dict__["address_type"] = address_type
        __props__.__dict__["bandwidth_package_id"] = bandwidth_package_id
        __props__.__dict__["egress"] = egress
        __props__.__dict__["internet_charge_type"] = internet_charge_type
        __props__.__dict__["internet_max_bandwidth_out"] = internet_max_bandwidth_out
        __props__.__dict__["internet_service_provider"] = internet_service_provider
        __props__.__dict__["tags"] = tags
        return PublicIpv6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> pulumi.Output[str]:
        """
        External network IP address.
        """
        return pulumi.get(self, "address_ip")

    @property
    @pulumi.getter(name="addressName")
    def address_name(self) -> pulumi.Output[str]:
        """
        EIP name, used to customize the personalized name of the EIP when applying for EIP. Default value: unnamed.
        """
        return pulumi.get(self, "address_name")

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> pulumi.Output[str]:
        """
        Elastic IPv6 type, optional values: - EIPv6: Ordinary IPv6 - HighQualityEIPv6: Premium IPv6 Note: You need to contact
        the product to open a premium IPv6 white list, and only some regions support premium IPv6 Default value: EIPv6.
        """
        return pulumi.get(self, "address_type")

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> pulumi.Output[str]:
        """
        Bandwidth packet unique ID parameter. If this parameter is set and the InternetChargeType is BANDWIDTH_PACKAGE, it means
        that the EIP created is added to the BGP bandwidth packet and the bandwidth packet is charged.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @property
    @pulumi.getter
    def egress(self) -> pulumi.Output[str]:
        """
        Elastic IPv6 network exit, optional values: - CENTER_EGRESS_1: Center Exit 1 - CENTER_EGRESS_2: Center Exit 2 -
        CENTER_EGRESS_3: Center Exit 3 Note: Network exports corresponding to different operators or resource types need to
        contact the product for clarification Default value: CENTER_EGRESS_1.
        """
        return pulumi.get(self, "egress")

    @property
    @pulumi.getter(name="internetChargeType")
    def internet_charge_type(self) -> pulumi.Output[str]:
        """
        Elastic IPv6 charging method, optional values: - BANDWIDTH_PACKAGE: Payment for Shared Bandwidth Package -
        TRAFFIC_POSTPAID_BY_HOUR: Traffic is paid by the hour Default value: TRAFFIC_POSTPAID_BY_HOUR.
        """
        return pulumi.get(self, "internet_charge_type")

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> pulumi.Output[int]:
        """
        Elastic IPv6 bandwidth limit in Mbps. The range of selectable values depends on the EIP billing method: -
        BANDWIDTH_PACKAGE: 1 Mbps to 2000 Mbps - TRAFFIC_POSTPAID_BY_HOUR: 1 Mbps to 100 Mbps Default value: 1 Mbps.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @property
    @pulumi.getter(name="internetServiceProvider")
    def internet_service_provider(self) -> pulumi.Output[str]:
        """
        Elastic IPv6 line type, default value: BGP. For users who have activated a static single-line IP whitelist, selectable
        values: - CMCC: China Mobile - CTCC: China Telecom - CUCC: China Unicom Note: Static single-wire IP is only supported in
        some regions.
        """
        return pulumi.get(self, "internet_service_provider")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

