# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CngwNetworkArgs', 'CngwNetwork']

@pulumi.input_type
class CngwNetworkArgs:
    def __init__(__self__, *,
                 gateway_id: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 internet_address_version: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_pay_mode: Optional[pulumi.Input[str]] = None,
                 master_zone_id: Optional[pulumi.Input[str]] = None,
                 multi_zone_flag: Optional[pulumi.Input[bool]] = None,
                 sla_type: Optional[pulumi.Input[str]] = None,
                 slave_zone_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CngwNetwork resource.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] group_id: gateway group ID.
        :param pulumi.Input[str] description: description of clb.
        :param pulumi.Input[str] internet_address_version: internet type. Reference value:`IPV4` (default value), `IPV6`.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network bandwidth.
        :param pulumi.Input[str] internet_pay_mode: trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        :param pulumi.Input[str] master_zone_id: primary availability zone.
        :param pulumi.Input[bool] multi_zone_flag: Whether load balancing has multiple availability zones.
        :param pulumi.Input[str] sla_type: specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
               when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
               `clb.c4.large`, `clb.c4.xlarge`.
        :param pulumi.Input[str] slave_zone_id: alternate availability zone.
        """
        pulumi.set(__self__, "gateway_id", gateway_id)
        pulumi.set(__self__, "group_id", group_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if internet_address_version is not None:
            pulumi.set(__self__, "internet_address_version", internet_address_version)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if internet_pay_mode is not None:
            pulumi.set(__self__, "internet_pay_mode", internet_pay_mode)
        if master_zone_id is not None:
            pulumi.set(__self__, "master_zone_id", master_zone_id)
        if multi_zone_flag is not None:
            pulumi.set(__self__, "multi_zone_flag", multi_zone_flag)
        if sla_type is not None:
            pulumi.set(__self__, "sla_type", sla_type)
        if slave_zone_id is not None:
            pulumi.set(__self__, "slave_zone_id", slave_zone_id)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Input[str]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        gateway group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        description of clb.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="internetAddressVersion")
    def internet_address_version(self) -> Optional[pulumi.Input[str]]:
        """
        internet type. Reference value:`IPV4` (default value), `IPV6`.
        """
        return pulumi.get(self, "internet_address_version")

    @internet_address_version.setter
    def internet_address_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_address_version", value)

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[pulumi.Input[int]]:
        """
        public network bandwidth.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @internet_max_bandwidth_out.setter
    def internet_max_bandwidth_out(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth_out", value)

    @property
    @pulumi.getter(name="internetPayMode")
    def internet_pay_mode(self) -> Optional[pulumi.Input[str]]:
        """
        trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        """
        return pulumi.get(self, "internet_pay_mode")

    @internet_pay_mode.setter
    def internet_pay_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_pay_mode", value)

    @property
    @pulumi.getter(name="masterZoneId")
    def master_zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        primary availability zone.
        """
        return pulumi.get(self, "master_zone_id")

    @master_zone_id.setter
    def master_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_zone_id", value)

    @property
    @pulumi.getter(name="multiZoneFlag")
    def multi_zone_flag(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether load balancing has multiple availability zones.
        """
        return pulumi.get(self, "multi_zone_flag")

    @multi_zone_flag.setter
    def multi_zone_flag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_zone_flag", value)

    @property
    @pulumi.getter(name="slaType")
    def sla_type(self) -> Optional[pulumi.Input[str]]:
        """
        specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
        when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
        `clb.c4.large`, `clb.c4.xlarge`.
        """
        return pulumi.get(self, "sla_type")

    @sla_type.setter
    def sla_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sla_type", value)

    @property
    @pulumi.getter(name="slaveZoneId")
    def slave_zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        alternate availability zone.
        """
        return pulumi.get(self, "slave_zone_id")

    @slave_zone_id.setter
    def slave_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slave_zone_id", value)


@pulumi.input_type
class _CngwNetworkState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 internet_address_version: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_pay_mode: Optional[pulumi.Input[str]] = None,
                 master_zone_id: Optional[pulumi.Input[str]] = None,
                 multi_zone_flag: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 sla_type: Optional[pulumi.Input[str]] = None,
                 slave_zone_id: Optional[pulumi.Input[str]] = None,
                 vip: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CngwNetwork resources.
        :param pulumi.Input[str] description: description of clb.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] group_id: gateway group ID.
        :param pulumi.Input[str] internet_address_version: internet type. Reference value:`IPV4` (default value), `IPV6`.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network bandwidth.
        :param pulumi.Input[str] internet_pay_mode: trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        :param pulumi.Input[str] master_zone_id: primary availability zone.
        :param pulumi.Input[bool] multi_zone_flag: Whether load balancing has multiple availability zones.
        :param pulumi.Input[str] network_id: network id.
        :param pulumi.Input[str] sla_type: specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
               when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
               `clb.c4.large`, `clb.c4.xlarge`.
        :param pulumi.Input[str] slave_zone_id: alternate availability zone.
        :param pulumi.Input[str] vip: clb vip.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if gateway_id is not None:
            pulumi.set(__self__, "gateway_id", gateway_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if internet_address_version is not None:
            pulumi.set(__self__, "internet_address_version", internet_address_version)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if internet_pay_mode is not None:
            pulumi.set(__self__, "internet_pay_mode", internet_pay_mode)
        if master_zone_id is not None:
            pulumi.set(__self__, "master_zone_id", master_zone_id)
        if multi_zone_flag is not None:
            pulumi.set(__self__, "multi_zone_flag", multi_zone_flag)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if sla_type is not None:
            pulumi.set(__self__, "sla_type", sla_type)
        if slave_zone_id is not None:
            pulumi.set(__self__, "slave_zone_id", slave_zone_id)
        if vip is not None:
            pulumi.set(__self__, "vip", vip)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        description of clb.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        gateway group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="internetAddressVersion")
    def internet_address_version(self) -> Optional[pulumi.Input[str]]:
        """
        internet type. Reference value:`IPV4` (default value), `IPV6`.
        """
        return pulumi.get(self, "internet_address_version")

    @internet_address_version.setter
    def internet_address_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_address_version", value)

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[pulumi.Input[int]]:
        """
        public network bandwidth.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @internet_max_bandwidth_out.setter
    def internet_max_bandwidth_out(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth_out", value)

    @property
    @pulumi.getter(name="internetPayMode")
    def internet_pay_mode(self) -> Optional[pulumi.Input[str]]:
        """
        trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        """
        return pulumi.get(self, "internet_pay_mode")

    @internet_pay_mode.setter
    def internet_pay_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internet_pay_mode", value)

    @property
    @pulumi.getter(name="masterZoneId")
    def master_zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        primary availability zone.
        """
        return pulumi.get(self, "master_zone_id")

    @master_zone_id.setter
    def master_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_zone_id", value)

    @property
    @pulumi.getter(name="multiZoneFlag")
    def multi_zone_flag(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether load balancing has multiple availability zones.
        """
        return pulumi.get(self, "multi_zone_flag")

    @multi_zone_flag.setter
    def multi_zone_flag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_zone_flag", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        network id.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="slaType")
    def sla_type(self) -> Optional[pulumi.Input[str]]:
        """
        specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
        when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
        `clb.c4.large`, `clb.c4.xlarge`.
        """
        return pulumi.get(self, "sla_type")

    @sla_type.setter
    def sla_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sla_type", value)

    @property
    @pulumi.getter(name="slaveZoneId")
    def slave_zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        alternate availability zone.
        """
        return pulumi.get(self, "slave_zone_id")

    @slave_zone_id.setter
    def slave_zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slave_zone_id", value)

    @property
    @pulumi.getter
    def vip(self) -> Optional[pulumi.Input[str]]:
        """
        clb vip.
        """
        return pulumi.get(self, "vip")

    @vip.setter
    def vip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vip", value)


class CngwNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 internet_address_version: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_pay_mode: Optional[pulumi.Input[str]] = None,
                 master_zone_id: Optional[pulumi.Input[str]] = None,
                 multi_zone_flag: Optional[pulumi.Input[bool]] = None,
                 sla_type: Optional[pulumi.Input[str]] = None,
                 slave_zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CngwNetwork resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description of clb.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] group_id: gateway group ID.
        :param pulumi.Input[str] internet_address_version: internet type. Reference value:`IPV4` (default value), `IPV6`.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network bandwidth.
        :param pulumi.Input[str] internet_pay_mode: trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        :param pulumi.Input[str] master_zone_id: primary availability zone.
        :param pulumi.Input[bool] multi_zone_flag: Whether load balancing has multiple availability zones.
        :param pulumi.Input[str] sla_type: specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
               when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
               `clb.c4.large`, `clb.c4.xlarge`.
        :param pulumi.Input[str] slave_zone_id: alternate availability zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CngwNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CngwNetwork resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CngwNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CngwNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 internet_address_version: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 internet_pay_mode: Optional[pulumi.Input[str]] = None,
                 master_zone_id: Optional[pulumi.Input[str]] = None,
                 multi_zone_flag: Optional[pulumi.Input[bool]] = None,
                 sla_type: Optional[pulumi.Input[str]] = None,
                 slave_zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CngwNetworkArgs.__new__(CngwNetworkArgs)

            __props__.__dict__["description"] = description
            if gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_id'")
            __props__.__dict__["gateway_id"] = gateway_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["internet_address_version"] = internet_address_version
            __props__.__dict__["internet_max_bandwidth_out"] = internet_max_bandwidth_out
            __props__.__dict__["internet_pay_mode"] = internet_pay_mode
            __props__.__dict__["master_zone_id"] = master_zone_id
            __props__.__dict__["multi_zone_flag"] = multi_zone_flag
            __props__.__dict__["sla_type"] = sla_type
            __props__.__dict__["slave_zone_id"] = slave_zone_id
            __props__.__dict__["network_id"] = None
            __props__.__dict__["vip"] = None
        super(CngwNetwork, __self__).__init__(
            'tencentcloud:Tse/cngwNetwork:CngwNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            internet_address_version: Optional[pulumi.Input[str]] = None,
            internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
            internet_pay_mode: Optional[pulumi.Input[str]] = None,
            master_zone_id: Optional[pulumi.Input[str]] = None,
            multi_zone_flag: Optional[pulumi.Input[bool]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            sla_type: Optional[pulumi.Input[str]] = None,
            slave_zone_id: Optional[pulumi.Input[str]] = None,
            vip: Optional[pulumi.Input[str]] = None) -> 'CngwNetwork':
        """
        Get an existing CngwNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description of clb.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] group_id: gateway group ID.
        :param pulumi.Input[str] internet_address_version: internet type. Reference value:`IPV4` (default value), `IPV6`.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network bandwidth.
        :param pulumi.Input[str] internet_pay_mode: trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        :param pulumi.Input[str] master_zone_id: primary availability zone.
        :param pulumi.Input[bool] multi_zone_flag: Whether load balancing has multiple availability zones.
        :param pulumi.Input[str] network_id: network id.
        :param pulumi.Input[str] sla_type: specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
               when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
               `clb.c4.large`, `clb.c4.xlarge`.
        :param pulumi.Input[str] slave_zone_id: alternate availability zone.
        :param pulumi.Input[str] vip: clb vip.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CngwNetworkState.__new__(_CngwNetworkState)

        __props__.__dict__["description"] = description
        __props__.__dict__["gateway_id"] = gateway_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["internet_address_version"] = internet_address_version
        __props__.__dict__["internet_max_bandwidth_out"] = internet_max_bandwidth_out
        __props__.__dict__["internet_pay_mode"] = internet_pay_mode
        __props__.__dict__["master_zone_id"] = master_zone_id
        __props__.__dict__["multi_zone_flag"] = multi_zone_flag
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["sla_type"] = sla_type
        __props__.__dict__["slave_zone_id"] = slave_zone_id
        __props__.__dict__["vip"] = vip
        return CngwNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        description of clb.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        gateway group ID.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="internetAddressVersion")
    def internet_address_version(self) -> pulumi.Output[Optional[str]]:
        """
        internet type. Reference value:`IPV4` (default value), `IPV6`.
        """
        return pulumi.get(self, "internet_address_version")

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> pulumi.Output[Optional[int]]:
        """
        public network bandwidth.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @property
    @pulumi.getter(name="internetPayMode")
    def internet_pay_mode(self) -> pulumi.Output[Optional[str]]:
        """
        trade type of internet. Reference value:`BANDWIDTH` (default value), `TRAFFIC`.
        """
        return pulumi.get(self, "internet_pay_mode")

    @property
    @pulumi.getter(name="masterZoneId")
    def master_zone_id(self) -> pulumi.Output[Optional[str]]:
        """
        primary availability zone.
        """
        return pulumi.get(self, "master_zone_id")

    @property
    @pulumi.getter(name="multiZoneFlag")
    def multi_zone_flag(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether load balancing has multiple availability zones.
        """
        return pulumi.get(self, "multi_zone_flag")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        network id.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="slaType")
    def sla_type(self) -> pulumi.Output[Optional[str]]:
        """
        specification type of clb. Default `shared` type when this parameter is empty, Note: input `shared` is not supported
        when creating. Reference value:`clb.c2.medium`, `clb.c3.small`, `clb.c3.medium`, `clb.c4.small`, `clb.c4.medium`,
        `clb.c4.large`, `clb.c4.xlarge`.
        """
        return pulumi.get(self, "sla_type")

    @property
    @pulumi.getter(name="slaveZoneId")
    def slave_zone_id(self) -> pulumi.Output[Optional[str]]:
        """
        alternate availability zone.
        """
        return pulumi.get(self, "slave_zone_id")

    @property
    @pulumi.getter
    def vip(self) -> pulumi.Output[str]:
        """
        clb vip.
        """
        return pulumi.get(self, "vip")
