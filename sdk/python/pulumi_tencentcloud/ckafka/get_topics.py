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
    'GetTopicsResult',
    'AwaitableGetTopicsResult',
    'get_topics',
    'get_topics_output',
]

@pulumi.output_type
class GetTopicsResult:
    """
    A collection of values returned by getTopics.
    """
    def __init__(__self__, id=None, instance_id=None, instance_lists=None, result_output_file=None, topic_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.GetTopicsInstanceListResult']:
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[str]:
        return pulumi.get(self, "topic_name")


class AwaitableGetTopicsResult(GetTopicsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTopicsResult(
            id=self.id,
            instance_id=self.instance_id,
            instance_lists=self.instance_lists,
            result_output_file=self.result_output_file,
            topic_name=self.topic_name)


def get_topics(instance_id: Optional[str] = None,
               result_output_file: Optional[str] = None,
               topic_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTopicsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['topicName'] = topic_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ckafka/getTopics:getTopics', __args__, opts=opts, typ=GetTopicsResult).value

    return AwaitableGetTopicsResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        instance_lists=pulumi.get(__ret__, 'instance_lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        topic_name=pulumi.get(__ret__, 'topic_name'))


@_utilities.lift_output_func(get_topics)
def get_topics_output(instance_id: Optional[pulumi.Input[str]] = None,
                      result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      topic_name: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTopicsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...