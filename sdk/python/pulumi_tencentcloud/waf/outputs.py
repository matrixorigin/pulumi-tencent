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
    'AntiInfoLeakStrategy',
    'ClbDomainLoadBalancerSet',
    'CustomRuleStrategy',
    'CustomWhiteRuleStrategy',
    'IpAccessControlItem',
    'SaasDomainPort',
    'GetAttackLogHistogramDataResult',
    'GetAttackLogListDataResult',
    'GetCiphersCipherResult',
    'GetDomainsDomainResult',
    'GetDomainsDomainLoadBalancerSetResult',
    'GetDomainsDomainPortResult',
    'GetFindDomainsListResult',
    'GetInstanceQpsLimitQpsDataResult',
    'GetPeakPointsPointResult',
    'GetTlsVersionsTlResult',
    'GetUserClbRegionsRichDataResult',
    'GetUserDomainsUsersInfoResult',
    'GetWafInfosHostListResult',
    'GetWafInfosHostListLoadBalancerResult',
    'GetWafInfosParamResult',
]

@pulumi.output_type
class AntiInfoLeakStrategy(dict):
    def __init__(__self__, *,
                 content: str,
                 field: str):
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def content(self) -> str:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def field(self) -> str:
        return pulumi.get(self, "field")


@pulumi.output_type
class ClbDomainLoadBalancerSet(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "listenerId":
            suggest = "listener_id"
        elif key == "listenerName":
            suggest = "listener_name"
        elif key == "loadBalancerId":
            suggest = "load_balancer_id"
        elif key == "loadBalancerName":
            suggest = "load_balancer_name"
        elif key == "loadBalancerType":
            suggest = "load_balancer_type"
        elif key == "numericalVpcId":
            suggest = "numerical_vpc_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClbDomainLoadBalancerSet. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClbDomainLoadBalancerSet.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClbDomainLoadBalancerSet.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 listener_id: str,
                 listener_name: str,
                 load_balancer_id: str,
                 load_balancer_name: str,
                 protocol: str,
                 region: str,
                 vip: str,
                 vport: int,
                 zone: str,
                 load_balancer_type: Optional[str] = None,
                 numerical_vpc_id: Optional[int] = None):
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "listener_name", listener_name)
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "vip", vip)
        pulumi.set(__self__, "vport", vport)
        pulumi.set(__self__, "zone", zone)
        if load_balancer_type is not None:
            pulumi.set(__self__, "load_balancer_type", load_balancer_type)
        if numerical_vpc_id is not None:
            pulumi.set(__self__, "numerical_vpc_id", numerical_vpc_id)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="listenerName")
    def listener_name(self) -> str:
        return pulumi.get(self, "listener_name")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> str:
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def vip(self) -> str:
        return pulumi.get(self, "vip")

    @property
    @pulumi.getter
    def vport(self) -> int:
        return pulumi.get(self, "vport")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> Optional[str]:
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter(name="numericalVpcId")
    def numerical_vpc_id(self) -> Optional[int]:
        return pulumi.get(self, "numerical_vpc_id")


@pulumi.output_type
class CustomRuleStrategy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "compareFunc":
            suggest = "compare_func"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomRuleStrategy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomRuleStrategy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomRuleStrategy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arg: str,
                 compare_func: str,
                 content: str,
                 field: str):
        pulumi.set(__self__, "arg", arg)
        pulumi.set(__self__, "compare_func", compare_func)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def arg(self) -> str:
        return pulumi.get(self, "arg")

    @property
    @pulumi.getter(name="compareFunc")
    def compare_func(self) -> str:
        return pulumi.get(self, "compare_func")

    @property
    @pulumi.getter
    def content(self) -> str:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def field(self) -> str:
        return pulumi.get(self, "field")


@pulumi.output_type
class CustomWhiteRuleStrategy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "compareFunc":
            suggest = "compare_func"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CustomWhiteRuleStrategy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CustomWhiteRuleStrategy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CustomWhiteRuleStrategy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 arg: str,
                 compare_func: str,
                 content: str,
                 field: str):
        pulumi.set(__self__, "arg", arg)
        pulumi.set(__self__, "compare_func", compare_func)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def arg(self) -> str:
        return pulumi.get(self, "arg")

    @property
    @pulumi.getter(name="compareFunc")
    def compare_func(self) -> str:
        return pulumi.get(self, "compare_func")

    @property
    @pulumi.getter
    def content(self) -> str:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def field(self) -> str:
        return pulumi.get(self, "field")


@pulumi.output_type
class IpAccessControlItem(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "validTs":
            suggest = "valid_ts"
        elif key == "validStatus":
            suggest = "valid_status"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IpAccessControlItem. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IpAccessControlItem.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IpAccessControlItem.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: int,
                 ip: str,
                 note: str,
                 valid_ts: int,
                 id: Optional[str] = None,
                 source: Optional[str] = None,
                 valid_status: Optional[int] = None):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "note", note)
        pulumi.set(__self__, "valid_ts", valid_ts)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if valid_status is not None:
            pulumi.set(__self__, "valid_status", valid_status)

    @property
    @pulumi.getter
    def action(self) -> int:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def ip(self) -> str:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def note(self) -> str:
        return pulumi.get(self, "note")

    @property
    @pulumi.getter(name="validTs")
    def valid_ts(self) -> int:
        return pulumi.get(self, "valid_ts")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="validStatus")
    def valid_status(self) -> Optional[int]:
        return pulumi.get(self, "valid_status")


@pulumi.output_type
class SaasDomainPort(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "upstreamPort":
            suggest = "upstream_port"
        elif key == "upstreamProtocol":
            suggest = "upstream_protocol"
        elif key == "nginxServerId":
            suggest = "nginx_server_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SaasDomainPort. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SaasDomainPort.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SaasDomainPort.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 port: str,
                 protocol: str,
                 upstream_port: str,
                 upstream_protocol: str,
                 nginx_server_id: Optional[str] = None):
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "upstream_port", upstream_port)
        pulumi.set(__self__, "upstream_protocol", upstream_protocol)
        if nginx_server_id is not None:
            pulumi.set(__self__, "nginx_server_id", nginx_server_id)

    @property
    @pulumi.getter
    def port(self) -> str:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="upstreamPort")
    def upstream_port(self) -> str:
        return pulumi.get(self, "upstream_port")

    @property
    @pulumi.getter(name="upstreamProtocol")
    def upstream_protocol(self) -> str:
        return pulumi.get(self, "upstream_protocol")

    @property
    @pulumi.getter(name="nginxServerId")
    def nginx_server_id(self) -> Optional[str]:
        return pulumi.get(self, "nginx_server_id")


@pulumi.output_type
class GetAttackLogHistogramDataResult(dict):
    def __init__(__self__, *,
                 count: int,
                 time_stamp: int):
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "time_stamp", time_stamp)

    @property
    @pulumi.getter
    def count(self) -> int:
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="timeStamp")
    def time_stamp(self) -> int:
        return pulumi.get(self, "time_stamp")


@pulumi.output_type
class GetAttackLogListDataResult(dict):
    def __init__(__self__, *,
                 content: str,
                 file_name: str,
                 source: str,
                 time_stamp: str):
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "file_name", file_name)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "time_stamp", time_stamp)

    @property
    @pulumi.getter
    def content(self) -> str:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> str:
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter
    def source(self) -> str:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="timeStamp")
    def time_stamp(self) -> str:
        return pulumi.get(self, "time_stamp")


@pulumi.output_type
class GetCiphersCipherResult(dict):
    def __init__(__self__, *,
                 cipher_id: int,
                 cipher_name: str,
                 version_id: int):
        pulumi.set(__self__, "cipher_id", cipher_id)
        pulumi.set(__self__, "cipher_name", cipher_name)
        pulumi.set(__self__, "version_id", version_id)

    @property
    @pulumi.getter(name="cipherId")
    def cipher_id(self) -> int:
        return pulumi.get(self, "cipher_id")

    @property
    @pulumi.getter(name="cipherName")
    def cipher_name(self) -> str:
        return pulumi.get(self, "cipher_name")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> int:
        return pulumi.get(self, "version_id")


@pulumi.output_type
class GetDomainsDomainResult(dict):
    def __init__(__self__, *,
                 alb_type: str,
                 api_status: int,
                 app_id: int,
                 bot_status: int,
                 cc_lists: Sequence[str],
                 cdc_clusters: str,
                 cls_status: int,
                 cname: str,
                 create_time: str,
                 domain: str,
                 domain_id: str,
                 edition: str,
                 engine: int,
                 flow_mode: int,
                 instance_id: str,
                 instance_name: str,
                 ipv6_status: int,
                 level: int,
                 load_balancer_sets: Sequence['outputs.GetDomainsDomainLoadBalancerSetResult'],
                 mode: int,
                 ports: Sequence['outputs.GetDomainsDomainPortResult'],
                 post_ckafka_status: int,
                 post_cls_status: int,
                 region: str,
                 rs_lists: Sequence[str],
                 sg_detail: str,
                 sg_state: int,
                 state: int,
                 status: int):
        pulumi.set(__self__, "alb_type", alb_type)
        pulumi.set(__self__, "api_status", api_status)
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "bot_status", bot_status)
        pulumi.set(__self__, "cc_lists", cc_lists)
        pulumi.set(__self__, "cdc_clusters", cdc_clusters)
        pulumi.set(__self__, "cls_status", cls_status)
        pulumi.set(__self__, "cname", cname)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "domain_id", domain_id)
        pulumi.set(__self__, "edition", edition)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "flow_mode", flow_mode)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "ipv6_status", ipv6_status)
        pulumi.set(__self__, "level", level)
        pulumi.set(__self__, "load_balancer_sets", load_balancer_sets)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "ports", ports)
        pulumi.set(__self__, "post_ckafka_status", post_ckafka_status)
        pulumi.set(__self__, "post_cls_status", post_cls_status)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "rs_lists", rs_lists)
        pulumi.set(__self__, "sg_detail", sg_detail)
        pulumi.set(__self__, "sg_state", sg_state)
        pulumi.set(__self__, "state", state)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="albType")
    def alb_type(self) -> str:
        return pulumi.get(self, "alb_type")

    @property
    @pulumi.getter(name="apiStatus")
    def api_status(self) -> int:
        return pulumi.get(self, "api_status")

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> int:
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="botStatus")
    def bot_status(self) -> int:
        return pulumi.get(self, "bot_status")

    @property
    @pulumi.getter(name="ccLists")
    def cc_lists(self) -> Sequence[str]:
        return pulumi.get(self, "cc_lists")

    @property
    @pulumi.getter(name="cdcClusters")
    def cdc_clusters(self) -> str:
        return pulumi.get(self, "cdc_clusters")

    @property
    @pulumi.getter(name="clsStatus")
    def cls_status(self) -> int:
        return pulumi.get(self, "cls_status")

    @property
    @pulumi.getter
    def cname(self) -> str:
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def edition(self) -> str:
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter
    def engine(self) -> int:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="flowMode")
    def flow_mode(self) -> int:
        return pulumi.get(self, "flow_mode")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="ipv6Status")
    def ipv6_status(self) -> int:
        return pulumi.get(self, "ipv6_status")

    @property
    @pulumi.getter
    def level(self) -> int:
        return pulumi.get(self, "level")

    @property
    @pulumi.getter(name="loadBalancerSets")
    def load_balancer_sets(self) -> Sequence['outputs.GetDomainsDomainLoadBalancerSetResult']:
        return pulumi.get(self, "load_balancer_sets")

    @property
    @pulumi.getter
    def mode(self) -> int:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def ports(self) -> Sequence['outputs.GetDomainsDomainPortResult']:
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="postCkafkaStatus")
    def post_ckafka_status(self) -> int:
        return pulumi.get(self, "post_ckafka_status")

    @property
    @pulumi.getter(name="postClsStatus")
    def post_cls_status(self) -> int:
        return pulumi.get(self, "post_cls_status")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="rsLists")
    def rs_lists(self) -> Sequence[str]:
        return pulumi.get(self, "rs_lists")

    @property
    @pulumi.getter(name="sgDetail")
    def sg_detail(self) -> str:
        return pulumi.get(self, "sg_detail")

    @property
    @pulumi.getter(name="sgState")
    def sg_state(self) -> int:
        return pulumi.get(self, "sg_state")

    @property
    @pulumi.getter
    def state(self) -> int:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> int:
        return pulumi.get(self, "status")


@pulumi.output_type
class GetDomainsDomainLoadBalancerSetResult(dict):
    def __init__(__self__, *,
                 listener_id: str,
                 listener_name: str,
                 load_balancer_id: str,
                 load_balancer_name: str,
                 load_balancer_type: str,
                 numerical_vpc_id: int,
                 protocol: str,
                 region: str,
                 vip: str,
                 vport: int,
                 zone: str):
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "listener_name", listener_name)
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        pulumi.set(__self__, "load_balancer_type", load_balancer_type)
        pulumi.set(__self__, "numerical_vpc_id", numerical_vpc_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "vip", vip)
        pulumi.set(__self__, "vport", vport)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="listenerName")
    def listener_name(self) -> str:
        return pulumi.get(self, "listener_name")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> str:
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> str:
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter(name="numericalVpcId")
    def numerical_vpc_id(self) -> int:
        return pulumi.get(self, "numerical_vpc_id")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def vip(self) -> str:
        return pulumi.get(self, "vip")

    @property
    @pulumi.getter
    def vport(self) -> int:
        return pulumi.get(self, "vport")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


@pulumi.output_type
class GetDomainsDomainPortResult(dict):
    def __init__(__self__, *,
                 nginx_server_id: int,
                 port: str,
                 protocol: str,
                 upstream_port: str,
                 upstream_protocol: str):
        pulumi.set(__self__, "nginx_server_id", nginx_server_id)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "upstream_port", upstream_port)
        pulumi.set(__self__, "upstream_protocol", upstream_protocol)

    @property
    @pulumi.getter(name="nginxServerId")
    def nginx_server_id(self) -> int:
        return pulumi.get(self, "nginx_server_id")

    @property
    @pulumi.getter
    def port(self) -> str:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="upstreamPort")
    def upstream_port(self) -> str:
        return pulumi.get(self, "upstream_port")

    @property
    @pulumi.getter(name="upstreamProtocol")
    def upstream_protocol(self) -> str:
        return pulumi.get(self, "upstream_protocol")


@pulumi.output_type
class GetFindDomainsListResult(dict):
    def __init__(__self__, *,
                 appid: int,
                 domain: str,
                 domain_id: str,
                 edition: str,
                 find_time: str,
                 instance_id: str,
                 ips: Sequence[str],
                 is_waf_domain: int):
        pulumi.set(__self__, "appid", appid)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "domain_id", domain_id)
        pulumi.set(__self__, "edition", edition)
        pulumi.set(__self__, "find_time", find_time)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "ips", ips)
        pulumi.set(__self__, "is_waf_domain", is_waf_domain)

    @property
    @pulumi.getter
    def appid(self) -> int:
        return pulumi.get(self, "appid")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def edition(self) -> str:
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="findTime")
    def find_time(self) -> str:
        return pulumi.get(self, "find_time")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="isWafDomain")
    def is_waf_domain(self) -> int:
        return pulumi.get(self, "is_waf_domain")


@pulumi.output_type
class GetInstanceQpsLimitQpsDataResult(dict):
    def __init__(__self__, *,
                 elastic_billing_default: int,
                 elastic_billing_max: int,
                 elastic_billing_min: int,
                 qps_extend_intl_max: int,
                 qps_extend_max: int):
        pulumi.set(__self__, "elastic_billing_default", elastic_billing_default)
        pulumi.set(__self__, "elastic_billing_max", elastic_billing_max)
        pulumi.set(__self__, "elastic_billing_min", elastic_billing_min)
        pulumi.set(__self__, "qps_extend_intl_max", qps_extend_intl_max)
        pulumi.set(__self__, "qps_extend_max", qps_extend_max)

    @property
    @pulumi.getter(name="elasticBillingDefault")
    def elastic_billing_default(self) -> int:
        return pulumi.get(self, "elastic_billing_default")

    @property
    @pulumi.getter(name="elasticBillingMax")
    def elastic_billing_max(self) -> int:
        return pulumi.get(self, "elastic_billing_max")

    @property
    @pulumi.getter(name="elasticBillingMin")
    def elastic_billing_min(self) -> int:
        return pulumi.get(self, "elastic_billing_min")

    @property
    @pulumi.getter(name="qpsExtendIntlMax")
    def qps_extend_intl_max(self) -> int:
        return pulumi.get(self, "qps_extend_intl_max")

    @property
    @pulumi.getter(name="qpsExtendMax")
    def qps_extend_max(self) -> int:
        return pulumi.get(self, "qps_extend_max")


@pulumi.output_type
class GetPeakPointsPointResult(dict):
    def __init__(__self__, *,
                 access: int,
                 attack: int,
                 bot_access: int,
                 cc: int,
                 down: int,
                 status_client_error: int,
                 status_ok: int,
                 status_redirect: int,
                 status_server_error: int,
                 time: int,
                 up: int,
                 upstream_client_error: int,
                 upstream_redirect: int,
                 upstream_server_error: int):
        pulumi.set(__self__, "access", access)
        pulumi.set(__self__, "attack", attack)
        pulumi.set(__self__, "bot_access", bot_access)
        pulumi.set(__self__, "cc", cc)
        pulumi.set(__self__, "down", down)
        pulumi.set(__self__, "status_client_error", status_client_error)
        pulumi.set(__self__, "status_ok", status_ok)
        pulumi.set(__self__, "status_redirect", status_redirect)
        pulumi.set(__self__, "status_server_error", status_server_error)
        pulumi.set(__self__, "time", time)
        pulumi.set(__self__, "up", up)
        pulumi.set(__self__, "upstream_client_error", upstream_client_error)
        pulumi.set(__self__, "upstream_redirect", upstream_redirect)
        pulumi.set(__self__, "upstream_server_error", upstream_server_error)

    @property
    @pulumi.getter
    def access(self) -> int:
        return pulumi.get(self, "access")

    @property
    @pulumi.getter
    def attack(self) -> int:
        return pulumi.get(self, "attack")

    @property
    @pulumi.getter(name="botAccess")
    def bot_access(self) -> int:
        return pulumi.get(self, "bot_access")

    @property
    @pulumi.getter
    def cc(self) -> int:
        return pulumi.get(self, "cc")

    @property
    @pulumi.getter
    def down(self) -> int:
        return pulumi.get(self, "down")

    @property
    @pulumi.getter(name="statusClientError")
    def status_client_error(self) -> int:
        return pulumi.get(self, "status_client_error")

    @property
    @pulumi.getter(name="statusOk")
    def status_ok(self) -> int:
        return pulumi.get(self, "status_ok")

    @property
    @pulumi.getter(name="statusRedirect")
    def status_redirect(self) -> int:
        return pulumi.get(self, "status_redirect")

    @property
    @pulumi.getter(name="statusServerError")
    def status_server_error(self) -> int:
        return pulumi.get(self, "status_server_error")

    @property
    @pulumi.getter
    def time(self) -> int:
        return pulumi.get(self, "time")

    @property
    @pulumi.getter
    def up(self) -> int:
        return pulumi.get(self, "up")

    @property
    @pulumi.getter(name="upstreamClientError")
    def upstream_client_error(self) -> int:
        return pulumi.get(self, "upstream_client_error")

    @property
    @pulumi.getter(name="upstreamRedirect")
    def upstream_redirect(self) -> int:
        return pulumi.get(self, "upstream_redirect")

    @property
    @pulumi.getter(name="upstreamServerError")
    def upstream_server_error(self) -> int:
        return pulumi.get(self, "upstream_server_error")


@pulumi.output_type
class GetTlsVersionsTlResult(dict):
    def __init__(__self__, *,
                 version_id: int,
                 version_name: str):
        pulumi.set(__self__, "version_id", version_id)
        pulumi.set(__self__, "version_name", version_name)

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> int:
        return pulumi.get(self, "version_id")

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> str:
        return pulumi.get(self, "version_name")


@pulumi.output_type
class GetUserClbRegionsRichDataResult(dict):
    def __init__(__self__, *,
                 code: str,
                 id: str,
                 text: str,
                 value: str):
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "text", text)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def code(self) -> str:
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def text(self) -> str:
        return pulumi.get(self, "text")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


@pulumi.output_type
class GetUserDomainsUsersInfoResult(dict):
    def __init__(__self__, *,
                 appid: int,
                 cls: int,
                 domain: str,
                 domain_id: str,
                 edition: str,
                 instance_id: str,
                 instance_name: str,
                 level: str,
                 write_config: str):
        pulumi.set(__self__, "appid", appid)
        pulumi.set(__self__, "cls", cls)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "domain_id", domain_id)
        pulumi.set(__self__, "edition", edition)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "level", level)
        pulumi.set(__self__, "write_config", write_config)

    @property
    @pulumi.getter
    def appid(self) -> int:
        return pulumi.get(self, "appid")

    @property
    @pulumi.getter
    def cls(self) -> int:
        return pulumi.get(self, "cls")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def edition(self) -> str:
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter
    def level(self) -> str:
        return pulumi.get(self, "level")

    @property
    @pulumi.getter(name="writeConfig")
    def write_config(self) -> str:
        return pulumi.get(self, "write_config")


@pulumi.output_type
class GetWafInfosHostListResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 domain_id: str,
                 flow_mode: int,
                 load_balancers: Sequence['outputs.GetWafInfosHostListLoadBalancerResult'],
                 status: int):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "domain_id", domain_id)
        pulumi.set(__self__, "flow_mode", flow_mode)
        pulumi.set(__self__, "load_balancers", load_balancers)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="flowMode")
    def flow_mode(self) -> int:
        return pulumi.get(self, "flow_mode")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Sequence['outputs.GetWafInfosHostListLoadBalancerResult']:
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter
    def status(self) -> int:
        return pulumi.get(self, "status")


@pulumi.output_type
class GetWafInfosHostListLoadBalancerResult(dict):
    def __init__(__self__, *,
                 listener_id: str,
                 listener_name: str,
                 load_balancer_id: str,
                 load_balancer_name: str,
                 load_balancer_type: str,
                 numerical_vpc_id: int,
                 protocol: str,
                 region: str,
                 vip: str,
                 vport: int,
                 zone: str):
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "listener_name", listener_name)
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        pulumi.set(__self__, "load_balancer_type", load_balancer_type)
        pulumi.set(__self__, "numerical_vpc_id", numerical_vpc_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "vip", vip)
        pulumi.set(__self__, "vport", vport)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="listenerName")
    def listener_name(self) -> str:
        return pulumi.get(self, "listener_name")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> str:
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> str:
        return pulumi.get(self, "load_balancer_type")

    @property
    @pulumi.getter(name="numericalVpcId")
    def numerical_vpc_id(self) -> int:
        return pulumi.get(self, "numerical_vpc_id")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def vip(self) -> str:
        return pulumi.get(self, "vip")

    @property
    @pulumi.getter
    def vport(self) -> int:
        return pulumi.get(self, "vport")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


@pulumi.output_type
class GetWafInfosParamResult(dict):
    def __init__(__self__, *,
                 load_balancer_id: str,
                 domain_id: Optional[str] = None,
                 listener_id: Optional[str] = None):
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[str]:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        return pulumi.get(self, "listener_id")

