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

__all__ = [
    'GetResetInstanceBlueprintResult',
    'AwaitableGetResetInstanceBlueprintResult',
    'get_reset_instance_blueprint',
    'get_reset_instance_blueprint_output',
]

@pulumi.output_type
class GetResetInstanceBlueprintResult:
    """
    A collection of values returned by getResetInstanceBlueprint.
    """
    def __init__(__self__, filters=None, id=None, instance_id=None, limit=None, offset=None, reset_instance_blueprint_sets=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if offset and not isinstance(offset, int):
            raise TypeError("Expected argument 'offset' to be a int")
        pulumi.set(__self__, "offset", offset)
        if reset_instance_blueprint_sets and not isinstance(reset_instance_blueprint_sets, list):
            raise TypeError("Expected argument 'reset_instance_blueprint_sets' to be a list")
        pulumi.set(__self__, "reset_instance_blueprint_sets", reset_instance_blueprint_sets)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetResetInstanceBlueprintFilterResult']]:
        return pulumi.get(self, "filters")

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
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def offset(self) -> Optional[int]:
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter(name="resetInstanceBlueprintSets")
    def reset_instance_blueprint_sets(self) -> Sequence['outputs.GetResetInstanceBlueprintResetInstanceBlueprintSetResult']:
        return pulumi.get(self, "reset_instance_blueprint_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetResetInstanceBlueprintResult(GetResetInstanceBlueprintResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResetInstanceBlueprintResult(
            filters=self.filters,
            id=self.id,
            instance_id=self.instance_id,
            limit=self.limit,
            offset=self.offset,
            reset_instance_blueprint_sets=self.reset_instance_blueprint_sets,
            result_output_file=self.result_output_file)


def get_reset_instance_blueprint(filters: Optional[Sequence[pulumi.InputType['GetResetInstanceBlueprintFilterArgs']]] = None,
                                 instance_id: Optional[str] = None,
                                 limit: Optional[int] = None,
                                 offset: Optional[int] = None,
                                 result_output_file: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResetInstanceBlueprintResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['instanceId'] = instance_id
    __args__['limit'] = limit
    __args__['offset'] = offset
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Lighthouse/getResetInstanceBlueprint:getResetInstanceBlueprint', __args__, opts=opts, typ=GetResetInstanceBlueprintResult).value

    return AwaitableGetResetInstanceBlueprintResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        limit=pulumi.get(__ret__, 'limit'),
        offset=pulumi.get(__ret__, 'offset'),
        reset_instance_blueprint_sets=pulumi.get(__ret__, 'reset_instance_blueprint_sets'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_reset_instance_blueprint)
def get_reset_instance_blueprint_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetResetInstanceBlueprintFilterArgs']]]]] = None,
                                        instance_id: Optional[pulumi.Input[str]] = None,
                                        limit: Optional[pulumi.Input[Optional[int]]] = None,
                                        offset: Optional[pulumi.Input[Optional[int]]] = None,
                                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResetInstanceBlueprintResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...