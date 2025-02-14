# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 db_instance_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 type: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 lock_status: Optional[pulumi.Input[bool]] = None,
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] db_instance_id: Instance ID in the format of postgres-4wdeb0zv.
        :param pulumi.Input[str] password: Password, which can contain 8-32 letters, digits, and symbols
               (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        :param pulumi.Input[str] type: The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        :param pulumi.Input[str] user_name: Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
               can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        :param pulumi.Input[bool] lock_status: whether lock account. true: locked; false: unlock.
        :param pulumi.Input[str] remark: Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        """
        pulumi.set(__self__, "db_instance_id", db_instance_id)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "user_name", user_name)
        if lock_status is not None:
            pulumi.set(__self__, "lock_status", lock_status)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID in the format of postgres-4wdeb0zv.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password, which can contain 8-32 letters, digits, and symbols
        (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
        can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="lockStatus")
    def lock_status(self) -> Optional[pulumi.Input[bool]]:
        """
        whether lock account. true: locked; false: unlock.
        """
        return pulumi.get(self, "lock_status")

    @lock_status.setter
    def lock_status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lock_status", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 lock_status: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[str] db_instance_id: Instance ID in the format of postgres-4wdeb0zv.
        :param pulumi.Input[bool] lock_status: whether lock account. true: locked; false: unlock.
        :param pulumi.Input[str] password: Password, which can contain 8-32 letters, digits, and symbols
               (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        :param pulumi.Input[str] remark: Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        :param pulumi.Input[str] type: The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        :param pulumi.Input[str] user_name: Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
               can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        """
        if db_instance_id is not None:
            pulumi.set(__self__, "db_instance_id", db_instance_id)
        if lock_status is not None:
            pulumi.set(__self__, "lock_status", lock_status)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID in the format of postgres-4wdeb0zv.
        """
        return pulumi.get(self, "db_instance_id")

    @db_instance_id.setter
    def db_instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_instance_id", value)

    @property
    @pulumi.getter(name="lockStatus")
    def lock_status(self) -> Optional[pulumi.Input[bool]]:
        """
        whether lock account. true: locked; false: unlock.
        """
        return pulumi.get(self, "lock_status")

    @lock_status.setter
    def lock_status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lock_status", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password, which can contain 8-32 letters, digits, and symbols
        (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
        can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 lock_status: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Account resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: Instance ID in the format of postgres-4wdeb0zv.
        :param pulumi.Input[bool] lock_status: whether lock account. true: locked; false: unlock.
        :param pulumi.Input[str] password: Password, which can contain 8-32 letters, digits, and symbols
               (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        :param pulumi.Input[str] remark: Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        :param pulumi.Input[str] type: The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        :param pulumi.Input[str] user_name: Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
               can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Account resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_id: Optional[pulumi.Input[str]] = None,
                 lock_status: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            if db_instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'db_instance_id'")
            __props__.__dict__["db_instance_id"] = db_instance_id
            __props__.__dict__["lock_status"] = lock_status
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["remark"] = remark
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Account, __self__).__init__(
            'tencentcloud:Postgresql/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            db_instance_id: Optional[pulumi.Input[str]] = None,
            lock_status: Optional[pulumi.Input[bool]] = None,
            password: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] db_instance_id: Instance ID in the format of postgres-4wdeb0zv.
        :param pulumi.Input[bool] lock_status: whether lock account. true: locked; false: unlock.
        :param pulumi.Input[str] password: Password, which can contain 8-32 letters, digits, and symbols
               (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        :param pulumi.Input[str] remark: Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        :param pulumi.Input[str] type: The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        :param pulumi.Input[str] user_name: Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
               can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["db_instance_id"] = db_instance_id
        __props__.__dict__["lock_status"] = lock_status
        __props__.__dict__["password"] = password
        __props__.__dict__["remark"] = remark
        __props__.__dict__["type"] = type
        __props__.__dict__["user_name"] = user_name
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dbInstanceId")
    def db_instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID in the format of postgres-4wdeb0zv.
        """
        return pulumi.get(self, "db_instance_id")

    @property
    @pulumi.getter(name="lockStatus")
    def lock_status(self) -> pulumi.Output[bool]:
        """
        whether lock account. true: locked; false: unlock.
        """
        return pulumi.get(self, "lock_status")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password, which can contain 8-32 letters, digits, and symbols
        (()`~!@#$%^&amp;amp;amp;*-+=_|{}[]:;&amp;amp;#39;&amp;amp;lt;&amp;amp;gt;,.?/); can&amp;amp;#39;t start with slash /.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remarks correspond to user `UserName`, which can contain 0-60 letters, digits, symbols (-_), and Chinese characters.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of user. Valid values: 1. normal: regular user; 2. tencentDBSuper: user with the pg_tencentdb_superuser role.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Instance username, which can contain 1-16 letters, digits, and underscore (_); can&amp;amp;#39;t be postgres;
        can&amp;amp;#39;t start with numbers, pg_, and tencentdb_.
        """
        return pulumi.get(self, "user_name")

