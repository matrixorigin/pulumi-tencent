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
    'GetParameterTemplatesResult',
    'AwaitableGetParameterTemplatesResult',
    'get_parameter_templates',
    'get_parameter_templates_output',
]

@pulumi.output_type
class GetParameterTemplatesResult:
    """
    A collection of values returned by getParameterTemplates.
    """
    def __init__(__self__, filters=None, id=None, lists=None, order_by=None, order_by_type=None, result_output_file=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if order_by_type and not isinstance(order_by_type, str):
            raise TypeError("Expected argument 'order_by_type' to be a str")
        pulumi.set(__self__, "order_by_type", order_by_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetParameterTemplatesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetParameterTemplatesListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter(name="orderByType")
    def order_by_type(self) -> Optional[str]:
        return pulumi.get(self, "order_by_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetParameterTemplatesResult(GetParameterTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetParameterTemplatesResult(
            filters=self.filters,
            id=self.id,
            lists=self.lists,
            order_by=self.order_by,
            order_by_type=self.order_by_type,
            result_output_file=self.result_output_file)


def get_parameter_templates(filters: Optional[Sequence[pulumi.InputType['GetParameterTemplatesFilterArgs']]] = None,
                            order_by: Optional[str] = None,
                            order_by_type: Optional[str] = None,
                            result_output_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetParameterTemplatesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['orderBy'] = order_by
    __args__['orderByType'] = order_by_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Postgresql/getParameterTemplates:getParameterTemplates', __args__, opts=opts, typ=GetParameterTemplatesResult).value

    return AwaitableGetParameterTemplatesResult(
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        lists=pulumi.get(__ret__, 'lists'),
        order_by=pulumi.get(__ret__, 'order_by'),
        order_by_type=pulumi.get(__ret__, 'order_by_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_parameter_templates)
def get_parameter_templates_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetParameterTemplatesFilterArgs']]]]] = None,
                                   order_by: Optional[pulumi.Input[Optional[str]]] = None,
                                   order_by_type: Optional[pulumi.Input[Optional[str]]] = None,
                                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetParameterTemplatesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...