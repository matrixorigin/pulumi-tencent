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
    'GetProductNamespaceResult',
    'AwaitableGetProductNamespaceResult',
    'get_product_namespace',
    'get_product_namespace_output',
]

@pulumi.output_type
class GetProductNamespaceResult:
    """
    A collection of values returned by getProductNamespace.
    """
    def __init__(__self__, id=None, lists=None, name=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetProductNamespaceListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetProductNamespaceResult(GetProductNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProductNamespaceResult(
            id=self.id,
            lists=self.lists,
            name=self.name,
            result_output_file=self.result_output_file)


def get_product_namespace(name: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProductNamespaceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Monitor/getProductNamespace:getProductNamespace', __args__, opts=opts, typ=GetProductNamespaceResult).value

    return AwaitableGetProductNamespaceResult(
        id=pulumi.get(__ret__, 'id'),
        lists=pulumi.get(__ret__, 'lists'),
        name=pulumi.get(__ret__, 'name'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_product_namespace)
def get_product_namespace_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProductNamespaceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...