# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'RuleDataArgs',
]

@pulumi.input_type
class RuleDataArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 dest_content: pulumi.Input[str],
                 dest_type: pulumi.Input[str],
                 rule_action: pulumi.Input[str],
                 source_content: pulumi.Input[str],
                 source_type: pulumi.Input[str],
                 order_index: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 service_template_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "dest_content", dest_content)
        pulumi.set(__self__, "dest_type", dest_type)
        pulumi.set(__self__, "rule_action", rule_action)
        pulumi.set(__self__, "source_content", source_content)
        pulumi.set(__self__, "source_type", source_type)
        if order_index is not None:
            pulumi.set(__self__, "order_index", order_index)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if service_template_id is not None:
            pulumi.set(__self__, "service_template_id", service_template_id)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destContent")
    def dest_content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dest_content")

    @dest_content.setter
    def dest_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "dest_content", value)

    @property
    @pulumi.getter(name="destType")
    def dest_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dest_type")

    @dest_type.setter
    def dest_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "dest_type", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="sourceContent")
    def source_content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_content")

    @source_content.setter
    def source_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_content", value)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source_type")

    @source_type.setter
    def source_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_type", value)

    @property
    @pulumi.getter(name="orderIndex")
    def order_index(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "order_index")

    @order_index.setter
    def order_index(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "order_index", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="serviceTemplateId")
    def service_template_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_template_id")

    @service_template_id.setter
    def service_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_template_id", value)

