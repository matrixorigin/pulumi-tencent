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

__all__ = ['ResourceArgs', 'Resource']

@pulumi.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 resource_loc: pulumi.Input['ResourceResourceLocArgs'],
                 resource_type: pulumi.Input[int],
                 folder_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_config_remark: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Resource resource.
        :param pulumi.Input['ResourceResourceLocArgs'] resource_loc: Resource location.
        :param pulumi.Input[int] resource_type: Resource type, only support JAR now, value is 1.
        :param pulumi.Input[str] folder_id: Folder id.
        :param pulumi.Input[str] name: Resource name.
        :param pulumi.Input[str] remark: Resource description.
        :param pulumi.Input[str] resource_config_remark: Resource version description.
        :param pulumi.Input[str] work_space_id: Workspace serialId.
        """
        pulumi.set(__self__, "resource_loc", resource_loc)
        pulumi.set(__self__, "resource_type", resource_type)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if resource_config_remark is not None:
            pulumi.set(__self__, "resource_config_remark", resource_config_remark)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="resourceLoc")
    def resource_loc(self) -> pulumi.Input['ResourceResourceLocArgs']:
        """
        Resource location.
        """
        return pulumi.get(self, "resource_loc")

    @resource_loc.setter
    def resource_loc(self, value: pulumi.Input['ResourceResourceLocArgs']):
        pulumi.set(self, "resource_loc", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[int]:
        """
        Resource type, only support JAR now, value is 1.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Folder id.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Resource description.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="resourceConfigRemark")
    def resource_config_remark(self) -> Optional[pulumi.Input[str]]:
        """
        Resource version description.
        """
        return pulumi.get(self, "resource_config_remark")

    @resource_config_remark.setter
    def resource_config_remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_config_remark", value)

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        Workspace serialId.
        """
        return pulumi.get(self, "work_space_id")

    @work_space_id.setter
    def work_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_space_id", value)


@pulumi.input_type
class _ResourceState:
    def __init__(__self__, *,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_config_remark: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_loc: Optional[pulumi.Input['ResourceResourceLocArgs']] = None,
                 resource_type: Optional[pulumi.Input[int]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Resource resources.
        :param pulumi.Input[str] folder_id: Folder id.
        :param pulumi.Input[str] name: Resource name.
        :param pulumi.Input[str] remark: Resource description.
        :param pulumi.Input[str] resource_config_remark: Resource version description.
        :param pulumi.Input[str] resource_id: Resource ID.
        :param pulumi.Input['ResourceResourceLocArgs'] resource_loc: Resource location.
        :param pulumi.Input[int] resource_type: Resource type, only support JAR now, value is 1.
        :param pulumi.Input[int] version: Resource Version.
        :param pulumi.Input[str] work_space_id: Workspace serialId.
        """
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if resource_config_remark is not None:
            pulumi.set(__self__, "resource_config_remark", resource_config_remark)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_loc is not None:
            pulumi.set(__self__, "resource_loc", resource_loc)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        Folder id.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Resource description.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="resourceConfigRemark")
    def resource_config_remark(self) -> Optional[pulumi.Input[str]]:
        """
        Resource version description.
        """
        return pulumi.get(self, "resource_config_remark")

    @resource_config_remark.setter
    def resource_config_remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_config_remark", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource ID.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceLoc")
    def resource_loc(self) -> Optional[pulumi.Input['ResourceResourceLocArgs']]:
        """
        Resource location.
        """
        return pulumi.get(self, "resource_loc")

    @resource_loc.setter
    def resource_loc(self, value: Optional[pulumi.Input['ResourceResourceLocArgs']]):
        pulumi.set(self, "resource_loc", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[int]]:
        """
        Resource type, only support JAR now, value is 1.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Resource Version.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        Workspace serialId.
        """
        return pulumi.get(self, "work_space_id")

    @work_space_id.setter
    def work_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_space_id", value)


class Resource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_config_remark: Optional[pulumi.Input[str]] = None,
                 resource_loc: Optional[pulumi.Input[pulumi.InputType['ResourceResourceLocArgs']]] = None,
                 resource_type: Optional[pulumi.Input[int]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_id: Folder id.
        :param pulumi.Input[str] name: Resource name.
        :param pulumi.Input[str] remark: Resource description.
        :param pulumi.Input[str] resource_config_remark: Resource version description.
        :param pulumi.Input[pulumi.InputType['ResourceResourceLocArgs']] resource_loc: Resource location.
        :param pulumi.Input[int] resource_type: Resource type, only support JAR now, value is 1.
        :param pulumi.Input[str] work_space_id: Workspace serialId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Resource resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 resource_config_remark: Optional[pulumi.Input[str]] = None,
                 resource_loc: Optional[pulumi.Input[pulumi.InputType['ResourceResourceLocArgs']]] = None,
                 resource_type: Optional[pulumi.Input[int]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceArgs.__new__(ResourceArgs)

            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["name"] = name
            __props__.__dict__["remark"] = remark
            __props__.__dict__["resource_config_remark"] = resource_config_remark
            if resource_loc is None and not opts.urn:
                raise TypeError("Missing required property 'resource_loc'")
            __props__.__dict__["resource_loc"] = resource_loc
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["work_space_id"] = work_space_id
            __props__.__dict__["resource_id"] = None
            __props__.__dict__["version"] = None
        super(Resource, __self__).__init__(
            'tencentcloud:Oceanus/resource:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            resource_config_remark: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_loc: Optional[pulumi.Input[pulumi.InputType['ResourceResourceLocArgs']]] = None,
            resource_type: Optional[pulumi.Input[int]] = None,
            version: Optional[pulumi.Input[int]] = None,
            work_space_id: Optional[pulumi.Input[str]] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder_id: Folder id.
        :param pulumi.Input[str] name: Resource name.
        :param pulumi.Input[str] remark: Resource description.
        :param pulumi.Input[str] resource_config_remark: Resource version description.
        :param pulumi.Input[str] resource_id: Resource ID.
        :param pulumi.Input[pulumi.InputType['ResourceResourceLocArgs']] resource_loc: Resource location.
        :param pulumi.Input[int] resource_type: Resource type, only support JAR now, value is 1.
        :param pulumi.Input[int] version: Resource Version.
        :param pulumi.Input[str] work_space_id: Workspace serialId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceState.__new__(_ResourceState)

        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["name"] = name
        __props__.__dict__["remark"] = remark
        __props__.__dict__["resource_config_remark"] = resource_config_remark
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_loc"] = resource_loc
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["version"] = version
        __props__.__dict__["work_space_id"] = work_space_id
        return Resource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[Optional[str]]:
        """
        Folder id.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Resource description.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="resourceConfigRemark")
    def resource_config_remark(self) -> pulumi.Output[Optional[str]]:
        """
        Resource version description.
        """
        return pulumi.get(self, "resource_config_remark")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        Resource ID.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceLoc")
    def resource_loc(self) -> pulumi.Output['outputs.ResourceResourceLoc']:
        """
        Resource location.
        """
        return pulumi.get(self, "resource_loc")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[int]:
        """
        Resource type, only support JAR now, value is 1.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        Resource Version.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> pulumi.Output[Optional[str]]:
        """
        Workspace serialId.
        """
        return pulumi.get(self, "work_space_id")
