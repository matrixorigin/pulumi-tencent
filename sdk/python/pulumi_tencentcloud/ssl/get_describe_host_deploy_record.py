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
    'GetDescribeHostDeployRecordResult',
    'AwaitableGetDescribeHostDeployRecordResult',
    'get_describe_host_deploy_record',
    'get_describe_host_deploy_record_output',
]

@pulumi.output_type
class GetDescribeHostDeployRecordResult:
    """
    A collection of values returned by getDescribeHostDeployRecord.
    """
    def __init__(__self__, certificate_id=None, deploy_record_lists=None, id=None, resource_type=None, result_output_file=None):
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if deploy_record_lists and not isinstance(deploy_record_lists, list):
            raise TypeError("Expected argument 'deploy_record_lists' to be a list")
        pulumi.set(__self__, "deploy_record_lists", deploy_record_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="deployRecordLists")
    def deploy_record_lists(self) -> Sequence['outputs.GetDescribeHostDeployRecordDeployRecordListResult']:
        return pulumi.get(self, "deploy_record_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDescribeHostDeployRecordResult(GetDescribeHostDeployRecordResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeHostDeployRecordResult(
            certificate_id=self.certificate_id,
            deploy_record_lists=self.deploy_record_lists,
            id=self.id,
            resource_type=self.resource_type,
            result_output_file=self.result_output_file)


def get_describe_host_deploy_record(certificate_id: Optional[str] = None,
                                    resource_type: Optional[str] = None,
                                    result_output_file: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeHostDeployRecordResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['certificateId'] = certificate_id
    __args__['resourceType'] = resource_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ssl/getDescribeHostDeployRecord:getDescribeHostDeployRecord', __args__, opts=opts, typ=GetDescribeHostDeployRecordResult).value

    return AwaitableGetDescribeHostDeployRecordResult(
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        deploy_record_lists=pulumi.get(__ret__, 'deploy_record_lists'),
        id=pulumi.get(__ret__, 'id'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_describe_host_deploy_record)
def get_describe_host_deploy_record_output(certificate_id: Optional[pulumi.Input[str]] = None,
                                           resource_type: Optional[pulumi.Input[Optional[str]]] = None,
                                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeHostDeployRecordResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...