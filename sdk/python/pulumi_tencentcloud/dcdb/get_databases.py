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
    'GetDatabasesResult',
    'AwaitableGetDatabasesResult',
    'get_databases',
    'get_databases_output',
]

@pulumi.output_type
class GetDatabasesResult:
    """
    A collection of values returned by getDatabases.
    """
    def __init__(__self__, id=None, instance_id=None, lists=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetDatabasesListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDatabasesResult(GetDatabasesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasesResult(
            id=self.id,
            instance_id=self.instance_id,
            lists=self.lists,
            result_output_file=self.result_output_file)


def get_databases(instance_id: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dcdb/getDatabases:getDatabases', __args__, opts=opts, typ=GetDatabasesResult).value

    return AwaitableGetDatabasesResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        lists=pulumi.get(__ret__, 'lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_databases)
def get_databases_output(instance_id: Optional[pulumi.Input[str]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabasesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...