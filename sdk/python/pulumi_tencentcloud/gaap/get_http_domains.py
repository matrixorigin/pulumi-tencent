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
    'GetHttpDomainsResult',
    'AwaitableGetHttpDomainsResult',
    'get_http_domains',
    'get_http_domains_output',
]

@pulumi.output_type
class GetHttpDomainsResult:
    """
    A collection of values returned by getHttpDomains.
    """
    def __init__(__self__, domain=None, domains=None, id=None, listener_id=None, result_output_file=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listener_id and not isinstance(listener_id, str):
            raise TypeError("Expected argument 'listener_id' to be a str")
        pulumi.set(__self__, "listener_id", listener_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetHttpDomainsDomainResult']:
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> str:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetHttpDomainsResult(GetHttpDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHttpDomainsResult(
            domain=self.domain,
            domains=self.domains,
            id=self.id,
            listener_id=self.listener_id,
            result_output_file=self.result_output_file)


def get_http_domains(domain: Optional[str] = None,
                     listener_id: Optional[str] = None,
                     result_output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHttpDomainsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['listenerId'] = listener_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Gaap/getHttpDomains:getHttpDomains', __args__, opts=opts, typ=GetHttpDomainsResult).value

    return AwaitableGetHttpDomainsResult(
        domain=pulumi.get(__ret__, 'domain'),
        domains=pulumi.get(__ret__, 'domains'),
        id=pulumi.get(__ret__, 'id'),
        listener_id=pulumi.get(__ret__, 'listener_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_http_domains)
def get_http_domains_output(domain: Optional[pulumi.Input[str]] = None,
                            listener_id: Optional[pulumi.Input[str]] = None,
                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHttpDomainsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...