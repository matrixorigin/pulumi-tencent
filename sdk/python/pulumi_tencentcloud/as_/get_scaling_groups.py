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
    'GetScalingGroupsResult',
    'AwaitableGetScalingGroupsResult',
    'get_scaling_groups',
    'get_scaling_groups_output',
]

@pulumi.output_type
class GetScalingGroupsResult:
    """
    A collection of values returned by getScalingGroups.
    """
    def __init__(__self__, configuration_id=None, id=None, result_output_file=None, scaling_group_id=None, scaling_group_lists=None, scaling_group_name=None, tags=None):
        if configuration_id and not isinstance(configuration_id, str):
            raise TypeError("Expected argument 'configuration_id' to be a str")
        pulumi.set(__self__, "configuration_id", configuration_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        if scaling_group_lists and not isinstance(scaling_group_lists, list):
            raise TypeError("Expected argument 'scaling_group_lists' to be a list")
        pulumi.set(__self__, "scaling_group_lists", scaling_group_lists)
        if scaling_group_name and not isinstance(scaling_group_name, str):
            raise TypeError("Expected argument 'scaling_group_name' to be a str")
        pulumi.set(__self__, "scaling_group_name", scaling_group_name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="configurationId")
    def configuration_id(self) -> Optional[str]:
        return pulumi.get(self, "configuration_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> Optional[str]:
        return pulumi.get(self, "scaling_group_id")

    @property
    @pulumi.getter(name="scalingGroupLists")
    def scaling_group_lists(self) -> Sequence['outputs.GetScalingGroupsScalingGroupListResult']:
        return pulumi.get(self, "scaling_group_lists")

    @property
    @pulumi.getter(name="scalingGroupName")
    def scaling_group_name(self) -> Optional[str]:
        return pulumi.get(self, "scaling_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetScalingGroupsResult(GetScalingGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingGroupsResult(
            configuration_id=self.configuration_id,
            id=self.id,
            result_output_file=self.result_output_file,
            scaling_group_id=self.scaling_group_id,
            scaling_group_lists=self.scaling_group_lists,
            scaling_group_name=self.scaling_group_name,
            tags=self.tags)


def get_scaling_groups(configuration_id: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       scaling_group_id: Optional[str] = None,
                       scaling_group_name: Optional[str] = None,
                       tags: Optional[Mapping[str, Any]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['configurationId'] = configuration_id
    __args__['resultOutputFile'] = result_output_file
    __args__['scalingGroupId'] = scaling_group_id
    __args__['scalingGroupName'] = scaling_group_name
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:As/getScalingGroups:getScalingGroups', __args__, opts=opts, typ=GetScalingGroupsResult).value

    return AwaitableGetScalingGroupsResult(
        configuration_id=pulumi.get(__ret__, 'configuration_id'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        scaling_group_id=pulumi.get(__ret__, 'scaling_group_id'),
        scaling_group_lists=pulumi.get(__ret__, 'scaling_group_lists'),
        scaling_group_name=pulumi.get(__ret__, 'scaling_group_name'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_scaling_groups)
def get_scaling_groups_output(configuration_id: Optional[pulumi.Input[Optional[str]]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              scaling_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                              scaling_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                              tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScalingGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...