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
    'GetEdgeFwSwitchesResult',
    'AwaitableGetEdgeFwSwitchesResult',
    'get_edge_fw_switches',
    'get_edge_fw_switches_output',
]

@pulumi.output_type
class GetEdgeFwSwitchesResult:
    """
    A collection of values returned by getEdgeFwSwitches.
    """
    def __init__(__self__, datas=None, id=None, result_output_file=None):
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def datas(self) -> Sequence['outputs.GetEdgeFwSwitchesDataResult']:
        return pulumi.get(self, "datas")

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


class AwaitableGetEdgeFwSwitchesResult(GetEdgeFwSwitchesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEdgeFwSwitchesResult(
            datas=self.datas,
            id=self.id,
            result_output_file=self.result_output_file)


def get_edge_fw_switches(result_output_file: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEdgeFwSwitchesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cfw/getEdgeFwSwitches:getEdgeFwSwitches', __args__, opts=opts, typ=GetEdgeFwSwitchesResult).value

    return AwaitableGetEdgeFwSwitchesResult(
        datas=pulumi.get(__ret__, 'datas'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_edge_fw_switches)
def get_edge_fw_switches_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEdgeFwSwitchesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...