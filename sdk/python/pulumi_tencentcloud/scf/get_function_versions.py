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
    'GetFunctionVersionsResult',
    'AwaitableGetFunctionVersionsResult',
    'get_function_versions',
    'get_function_versions_output',
]

@pulumi.output_type
class GetFunctionVersionsResult:
    """
    A collection of values returned by getFunctionVersions.
    """
    def __init__(__self__, function_name=None, id=None, namespace=None, order=None, order_by=None, result_output_file=None, versions=None):
        if function_name and not isinstance(function_name, str):
            raise TypeError("Expected argument 'function_name' to be a str")
        pulumi.set(__self__, "function_name", function_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if order and not isinstance(order, str):
            raise TypeError("Expected argument 'order' to be a str")
        pulumi.set(__self__, "order", order)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> str:
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.GetFunctionVersionsVersionResult']:
        return pulumi.get(self, "versions")


class AwaitableGetFunctionVersionsResult(GetFunctionVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionVersionsResult(
            function_name=self.function_name,
            id=self.id,
            namespace=self.namespace,
            order=self.order,
            order_by=self.order_by,
            result_output_file=self.result_output_file,
            versions=self.versions)


def get_function_versions(function_name: Optional[str] = None,
                          namespace: Optional[str] = None,
                          order: Optional[str] = None,
                          order_by: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionVersionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['functionName'] = function_name
    __args__['namespace'] = namespace
    __args__['order'] = order
    __args__['orderBy'] = order_by
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Scf/getFunctionVersions:getFunctionVersions', __args__, opts=opts, typ=GetFunctionVersionsResult).value

    return AwaitableGetFunctionVersionsResult(
        function_name=pulumi.get(__ret__, 'function_name'),
        id=pulumi.get(__ret__, 'id'),
        namespace=pulumi.get(__ret__, 'namespace'),
        order=pulumi.get(__ret__, 'order'),
        order_by=pulumi.get(__ret__, 'order_by'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        versions=pulumi.get(__ret__, 'versions'))


@_utilities.lift_output_func(get_function_versions)
def get_function_versions_output(function_name: Optional[pulumi.Input[str]] = None,
                                 namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                 order: Optional[pulumi.Input[Optional[str]]] = None,
                                 order_by: Optional[pulumi.Input[Optional[str]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionVersionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...