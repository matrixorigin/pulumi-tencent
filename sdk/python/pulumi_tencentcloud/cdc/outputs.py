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
    'GetDedicatedClusterHostsHostInfoSetResult',
    'GetDedicatedClusterInstanceTypesDedicatedClusterInstanceTypeSetResult',
    'GetDedicatedClusterOrdersDedicatedClusterOrderSetResult',
    'GetDedicatedClusterOrdersDedicatedClusterOrderSetDedicatedClusterOrderItemResult',
]

@pulumi.output_type
class GetDedicatedClusterHostsHostInfoSetResult(dict):
    def __init__(__self__, *,
                 cpu_available: int,
                 cpu_total: int,
                 expire_time: str,
                 host_id: str,
                 host_ip: str,
                 host_status: str,
                 host_type: str,
                 mem_available: int,
                 mem_total: int,
                 run_time: str,
                 service_type: str):
        pulumi.set(__self__, "cpu_available", cpu_available)
        pulumi.set(__self__, "cpu_total", cpu_total)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "host_id", host_id)
        pulumi.set(__self__, "host_ip", host_ip)
        pulumi.set(__self__, "host_status", host_status)
        pulumi.set(__self__, "host_type", host_type)
        pulumi.set(__self__, "mem_available", mem_available)
        pulumi.set(__self__, "mem_total", mem_total)
        pulumi.set(__self__, "run_time", run_time)
        pulumi.set(__self__, "service_type", service_type)

    @property
    @pulumi.getter(name="cpuAvailable")
    def cpu_available(self) -> int:
        return pulumi.get(self, "cpu_available")

    @property
    @pulumi.getter(name="cpuTotal")
    def cpu_total(self) -> int:
        return pulumi.get(self, "cpu_total")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> str:
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter(name="hostIp")
    def host_ip(self) -> str:
        return pulumi.get(self, "host_ip")

    @property
    @pulumi.getter(name="hostStatus")
    def host_status(self) -> str:
        return pulumi.get(self, "host_status")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> str:
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter(name="memAvailable")
    def mem_available(self) -> int:
        return pulumi.get(self, "mem_available")

    @property
    @pulumi.getter(name="memTotal")
    def mem_total(self) -> int:
        return pulumi.get(self, "mem_total")

    @property
    @pulumi.getter(name="runTime")
    def run_time(self) -> str:
        return pulumi.get(self, "run_time")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        return pulumi.get(self, "service_type")


@pulumi.output_type
class GetDedicatedClusterInstanceTypesDedicatedClusterInstanceTypeSetResult(dict):
    def __init__(__self__, *,
                 cpu: int,
                 cpu_type: str,
                 fpga: int,
                 gpu: int,
                 instance_bandwidth: float,
                 instance_family: str,
                 instance_pps: int,
                 instance_type: str,
                 memory: int,
                 network_card: int,
                 remark: str,
                 status: str,
                 storage_block_amount: int,
                 type_name: str,
                 zone: str):
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "cpu_type", cpu_type)
        pulumi.set(__self__, "fpga", fpga)
        pulumi.set(__self__, "gpu", gpu)
        pulumi.set(__self__, "instance_bandwidth", instance_bandwidth)
        pulumi.set(__self__, "instance_family", instance_family)
        pulumi.set(__self__, "instance_pps", instance_pps)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "memory", memory)
        pulumi.set(__self__, "network_card", network_card)
        pulumi.set(__self__, "remark", remark)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage_block_amount", storage_block_amount)
        pulumi.set(__self__, "type_name", type_name)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def cpu(self) -> int:
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="cpuType")
    def cpu_type(self) -> str:
        return pulumi.get(self, "cpu_type")

    @property
    @pulumi.getter
    def fpga(self) -> int:
        return pulumi.get(self, "fpga")

    @property
    @pulumi.getter
    def gpu(self) -> int:
        return pulumi.get(self, "gpu")

    @property
    @pulumi.getter(name="instanceBandwidth")
    def instance_bandwidth(self) -> float:
        return pulumi.get(self, "instance_bandwidth")

    @property
    @pulumi.getter(name="instanceFamily")
    def instance_family(self) -> str:
        return pulumi.get(self, "instance_family")

    @property
    @pulumi.getter(name="instancePps")
    def instance_pps(self) -> int:
        return pulumi.get(self, "instance_pps")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def memory(self) -> int:
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="networkCard")
    def network_card(self) -> int:
        return pulumi.get(self, "network_card")

    @property
    @pulumi.getter
    def remark(self) -> str:
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageBlockAmount")
    def storage_block_amount(self) -> int:
        return pulumi.get(self, "storage_block_amount")

    @property
    @pulumi.getter(name="typeName")
    def type_name(self) -> str:
        return pulumi.get(self, "type_name")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


@pulumi.output_type
class GetDedicatedClusterOrdersDedicatedClusterOrderSetResult(dict):
    def __init__(__self__, *,
                 action: str,
                 cpu: int,
                 create_time: str,
                 dedicated_cluster_id: str,
                 dedicated_cluster_order_id: str,
                 dedicated_cluster_order_items: Sequence['outputs.GetDedicatedClusterOrdersDedicatedClusterOrderSetDedicatedClusterOrderItemResult'],
                 dedicated_cluster_type_id: str,
                 gpu: int,
                 mem: int,
                 order_status: str,
                 order_type: str,
                 pay_status: int,
                 pay_type: str,
                 power_draw: float,
                 supported_instance_families: Sequence[str],
                 supported_storage_types: Sequence[str],
                 supported_uplink_speeds: Sequence[int],
                 time_span: int,
                 time_unit: str,
                 weight: int):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "dedicated_cluster_id", dedicated_cluster_id)
        pulumi.set(__self__, "dedicated_cluster_order_id", dedicated_cluster_order_id)
        pulumi.set(__self__, "dedicated_cluster_order_items", dedicated_cluster_order_items)
        pulumi.set(__self__, "dedicated_cluster_type_id", dedicated_cluster_type_id)
        pulumi.set(__self__, "gpu", gpu)
        pulumi.set(__self__, "mem", mem)
        pulumi.set(__self__, "order_status", order_status)
        pulumi.set(__self__, "order_type", order_type)
        pulumi.set(__self__, "pay_status", pay_status)
        pulumi.set(__self__, "pay_type", pay_type)
        pulumi.set(__self__, "power_draw", power_draw)
        pulumi.set(__self__, "supported_instance_families", supported_instance_families)
        pulumi.set(__self__, "supported_storage_types", supported_storage_types)
        pulumi.set(__self__, "supported_uplink_speeds", supported_uplink_speeds)
        pulumi.set(__self__, "time_span", time_span)
        pulumi.set(__self__, "time_unit", time_unit)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def cpu(self) -> int:
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dedicatedClusterId")
    def dedicated_cluster_id(self) -> str:
        return pulumi.get(self, "dedicated_cluster_id")

    @property
    @pulumi.getter(name="dedicatedClusterOrderId")
    def dedicated_cluster_order_id(self) -> str:
        return pulumi.get(self, "dedicated_cluster_order_id")

    @property
    @pulumi.getter(name="dedicatedClusterOrderItems")
    def dedicated_cluster_order_items(self) -> Sequence['outputs.GetDedicatedClusterOrdersDedicatedClusterOrderSetDedicatedClusterOrderItemResult']:
        return pulumi.get(self, "dedicated_cluster_order_items")

    @property
    @pulumi.getter(name="dedicatedClusterTypeId")
    def dedicated_cluster_type_id(self) -> str:
        return pulumi.get(self, "dedicated_cluster_type_id")

    @property
    @pulumi.getter
    def gpu(self) -> int:
        return pulumi.get(self, "gpu")

    @property
    @pulumi.getter
    def mem(self) -> int:
        return pulumi.get(self, "mem")

    @property
    @pulumi.getter(name="orderStatus")
    def order_status(self) -> str:
        return pulumi.get(self, "order_status")

    @property
    @pulumi.getter(name="orderType")
    def order_type(self) -> str:
        return pulumi.get(self, "order_type")

    @property
    @pulumi.getter(name="payStatus")
    def pay_status(self) -> int:
        return pulumi.get(self, "pay_status")

    @property
    @pulumi.getter(name="payType")
    def pay_type(self) -> str:
        return pulumi.get(self, "pay_type")

    @property
    @pulumi.getter(name="powerDraw")
    def power_draw(self) -> float:
        return pulumi.get(self, "power_draw")

    @property
    @pulumi.getter(name="supportedInstanceFamilies")
    def supported_instance_families(self) -> Sequence[str]:
        return pulumi.get(self, "supported_instance_families")

    @property
    @pulumi.getter(name="supportedStorageTypes")
    def supported_storage_types(self) -> Sequence[str]:
        return pulumi.get(self, "supported_storage_types")

    @property
    @pulumi.getter(name="supportedUplinkSpeeds")
    def supported_uplink_speeds(self) -> Sequence[int]:
        return pulumi.get(self, "supported_uplink_speeds")

    @property
    @pulumi.getter(name="timeSpan")
    def time_span(self) -> int:
        return pulumi.get(self, "time_span")

    @property
    @pulumi.getter(name="timeUnit")
    def time_unit(self) -> str:
        return pulumi.get(self, "time_unit")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")


@pulumi.output_type
class GetDedicatedClusterOrdersDedicatedClusterOrderSetDedicatedClusterOrderItemResult(dict):
    def __init__(__self__, *,
                 compute_format: str,
                 count: int,
                 create_time: str,
                 dedicated_cluster_type_id: str,
                 description: str,
                 name: str,
                 power_draw: float,
                 sub_order_id: str,
                 sub_order_pay_status: int,
                 sub_order_status: str,
                 supported_instance_families: Sequence[str],
                 supported_storage_types: Sequence[str],
                 supported_uplink_speeds: Sequence[int],
                 total_cpu: int,
                 total_gpu: int,
                 total_mem: int,
                 type_family: str,
                 type_name: str,
                 weight: int):
        pulumi.set(__self__, "compute_format", compute_format)
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "dedicated_cluster_type_id", dedicated_cluster_type_id)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "power_draw", power_draw)
        pulumi.set(__self__, "sub_order_id", sub_order_id)
        pulumi.set(__self__, "sub_order_pay_status", sub_order_pay_status)
        pulumi.set(__self__, "sub_order_status", sub_order_status)
        pulumi.set(__self__, "supported_instance_families", supported_instance_families)
        pulumi.set(__self__, "supported_storage_types", supported_storage_types)
        pulumi.set(__self__, "supported_uplink_speeds", supported_uplink_speeds)
        pulumi.set(__self__, "total_cpu", total_cpu)
        pulumi.set(__self__, "total_gpu", total_gpu)
        pulumi.set(__self__, "total_mem", total_mem)
        pulumi.set(__self__, "type_family", type_family)
        pulumi.set(__self__, "type_name", type_name)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="computeFormat")
    def compute_format(self) -> str:
        return pulumi.get(self, "compute_format")

    @property
    @pulumi.getter
    def count(self) -> int:
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dedicatedClusterTypeId")
    def dedicated_cluster_type_id(self) -> str:
        return pulumi.get(self, "dedicated_cluster_type_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="powerDraw")
    def power_draw(self) -> float:
        return pulumi.get(self, "power_draw")

    @property
    @pulumi.getter(name="subOrderId")
    def sub_order_id(self) -> str:
        return pulumi.get(self, "sub_order_id")

    @property
    @pulumi.getter(name="subOrderPayStatus")
    def sub_order_pay_status(self) -> int:
        return pulumi.get(self, "sub_order_pay_status")

    @property
    @pulumi.getter(name="subOrderStatus")
    def sub_order_status(self) -> str:
        return pulumi.get(self, "sub_order_status")

    @property
    @pulumi.getter(name="supportedInstanceFamilies")
    def supported_instance_families(self) -> Sequence[str]:
        return pulumi.get(self, "supported_instance_families")

    @property
    @pulumi.getter(name="supportedStorageTypes")
    def supported_storage_types(self) -> Sequence[str]:
        return pulumi.get(self, "supported_storage_types")

    @property
    @pulumi.getter(name="supportedUplinkSpeeds")
    def supported_uplink_speeds(self) -> Sequence[int]:
        return pulumi.get(self, "supported_uplink_speeds")

    @property
    @pulumi.getter(name="totalCpu")
    def total_cpu(self) -> int:
        return pulumi.get(self, "total_cpu")

    @property
    @pulumi.getter(name="totalGpu")
    def total_gpu(self) -> int:
        return pulumi.get(self, "total_gpu")

    @property
    @pulumi.getter(name="totalMem")
    def total_mem(self) -> int:
        return pulumi.get(self, "total_mem")

    @property
    @pulumi.getter(name="typeFamily")
    def type_family(self) -> str:
        return pulumi.get(self, "type_family")

    @property
    @pulumi.getter(name="typeName")
    def type_name(self) -> str:
        return pulumi.get(self, "type_name")

    @property
    @pulumi.getter
    def weight(self) -> int:
        return pulumi.get(self, "weight")


