# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StartBackupFullMigrationArgs', 'StartBackupFullMigration']

@pulumi.input_type
class StartBackupFullMigrationArgs:
    def __init__(__self__, *,
                 backup_migration_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a StartBackupFullMigration resource.
        :param pulumi.Input[str] backup_migration_id: Backup import task ID, returned by the CreateBackupMigration interface.
        :param pulumi.Input[str] instance_id: ID of imported target instance.
        """
        pulumi.set(__self__, "backup_migration_id", backup_migration_id)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="backupMigrationId")
    def backup_migration_id(self) -> pulumi.Input[str]:
        """
        Backup import task ID, returned by the CreateBackupMigration interface.
        """
        return pulumi.get(self, "backup_migration_id")

    @backup_migration_id.setter
    def backup_migration_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_migration_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of imported target instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _StartBackupFullMigrationState:
    def __init__(__self__, *,
                 backup_migration_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StartBackupFullMigration resources.
        :param pulumi.Input[str] backup_migration_id: Backup import task ID, returned by the CreateBackupMigration interface.
        :param pulumi.Input[str] instance_id: ID of imported target instance.
        """
        if backup_migration_id is not None:
            pulumi.set(__self__, "backup_migration_id", backup_migration_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="backupMigrationId")
    def backup_migration_id(self) -> Optional[pulumi.Input[str]]:
        """
        Backup import task ID, returned by the CreateBackupMigration interface.
        """
        return pulumi.get(self, "backup_migration_id")

    @backup_migration_id.setter
    def backup_migration_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_migration_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of imported target instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class StartBackupFullMigration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_migration_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a StartBackupFullMigration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_migration_id: Backup import task ID, returned by the CreateBackupMigration interface.
        :param pulumi.Input[str] instance_id: ID of imported target instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StartBackupFullMigrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a StartBackupFullMigration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param StartBackupFullMigrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StartBackupFullMigrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_migration_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StartBackupFullMigrationArgs.__new__(StartBackupFullMigrationArgs)

            if backup_migration_id is None and not opts.urn:
                raise TypeError("Missing required property 'backup_migration_id'")
            __props__.__dict__["backup_migration_id"] = backup_migration_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(StartBackupFullMigration, __self__).__init__(
            'tencentcloud:Sqlserver/startBackupFullMigration:StartBackupFullMigration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_migration_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'StartBackupFullMigration':
        """
        Get an existing StartBackupFullMigration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_migration_id: Backup import task ID, returned by the CreateBackupMigration interface.
        :param pulumi.Input[str] instance_id: ID of imported target instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StartBackupFullMigrationState.__new__(_StartBackupFullMigrationState)

        __props__.__dict__["backup_migration_id"] = backup_migration_id
        __props__.__dict__["instance_id"] = instance_id
        return StartBackupFullMigration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupMigrationId")
    def backup_migration_id(self) -> pulumi.Output[str]:
        """
        Backup import task ID, returned by the CreateBackupMigration interface.
        """
        return pulumi.get(self, "backup_migration_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of imported target instance.
        """
        return pulumi.get(self, "instance_id")
