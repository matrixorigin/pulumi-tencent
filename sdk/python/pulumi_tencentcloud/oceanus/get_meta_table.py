# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetMetaTableResult',
    'AwaitableGetMetaTableResult',
    'get_meta_table',
    'get_meta_table_output',
]

@pulumi.output_type
class GetMetaTableResult:
    """
    A collection of values returned by getMetaTable.
    """
    def __init__(__self__, catalog=None, create_time=None, database=None, ddl=None, id=None, result_output_file=None, serial_id=None, table=None, work_space_id=None):
        if catalog and not isinstance(catalog, str):
            raise TypeError("Expected argument 'catalog' to be a str")
        pulumi.set(__self__, "catalog", catalog)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if ddl and not isinstance(ddl, str):
            raise TypeError("Expected argument 'ddl' to be a str")
        pulumi.set(__self__, "ddl", ddl)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if serial_id and not isinstance(serial_id, str):
            raise TypeError("Expected argument 'serial_id' to be a str")
        pulumi.set(__self__, "serial_id", serial_id)
        if table and not isinstance(table, str):
            raise TypeError("Expected argument 'table' to be a str")
        pulumi.set(__self__, "table", table)
        if work_space_id and not isinstance(work_space_id, str):
            raise TypeError("Expected argument 'work_space_id' to be a str")
        pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter
    def catalog(self) -> str:
        return pulumi.get(self, "catalog")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def database(self) -> str:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def ddl(self) -> str:
        return pulumi.get(self, "ddl")

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
    @pulumi.getter(name="serialId")
    def serial_id(self) -> str:
        return pulumi.get(self, "serial_id")

    @property
    @pulumi.getter
    def table(self) -> str:
        return pulumi.get(self, "table")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> str:
        return pulumi.get(self, "work_space_id")


class AwaitableGetMetaTableResult(GetMetaTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetaTableResult(
            catalog=self.catalog,
            create_time=self.create_time,
            database=self.database,
            ddl=self.ddl,
            id=self.id,
            result_output_file=self.result_output_file,
            serial_id=self.serial_id,
            table=self.table,
            work_space_id=self.work_space_id)


def get_meta_table(catalog: Optional[str] = None,
                   database: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   table: Optional[str] = None,
                   work_space_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMetaTableResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['catalog'] = catalog
    __args__['database'] = database
    __args__['resultOutputFile'] = result_output_file
    __args__['table'] = table
    __args__['workSpaceId'] = work_space_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Oceanus/getMetaTable:getMetaTable', __args__, opts=opts, typ=GetMetaTableResult).value

    return AwaitableGetMetaTableResult(
        catalog=pulumi.get(__ret__, 'catalog'),
        create_time=pulumi.get(__ret__, 'create_time'),
        database=pulumi.get(__ret__, 'database'),
        ddl=pulumi.get(__ret__, 'ddl'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        serial_id=pulumi.get(__ret__, 'serial_id'),
        table=pulumi.get(__ret__, 'table'),
        work_space_id=pulumi.get(__ret__, 'work_space_id'))


@_utilities.lift_output_func(get_meta_table)
def get_meta_table_output(catalog: Optional[pulumi.Input[str]] = None,
                          database: Optional[pulumi.Input[str]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          table: Optional[pulumi.Input[str]] = None,
                          work_space_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMetaTableResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
