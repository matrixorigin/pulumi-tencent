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

__all__ = ['DetachUserPolicyOperationArgs', 'DetachUserPolicyOperation']

@pulumi.input_type
class DetachUserPolicyOperationArgs:
    def __init__(__self__, *,
                 user_id: pulumi.Input[str],
                 policy_sets: Optional[pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]]] = None):
        """
        The set of arguments for constructing a DetachUserPolicyOperation resource.
        :param pulumi.Input[str] user_id: User id, the same as the sub-user uin.
        :param pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]] policy_sets: Authentication policy collection.
        """
        pulumi.set(__self__, "user_id", user_id)
        if policy_sets is not None:
            pulumi.set(__self__, "policy_sets", policy_sets)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        User id, the same as the sub-user uin.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="policySets")
    def policy_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]]]:
        """
        Authentication policy collection.
        """
        return pulumi.get(self, "policy_sets")

    @policy_sets.setter
    def policy_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]]]):
        pulumi.set(self, "policy_sets", value)


@pulumi.input_type
class _DetachUserPolicyOperationState:
    def __init__(__self__, *,
                 policy_sets: Optional[pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DetachUserPolicyOperation resources.
        :param pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]] policy_sets: Authentication policy collection.
        :param pulumi.Input[str] user_id: User id, the same as the sub-user uin.
        """
        if policy_sets is not None:
            pulumi.set(__self__, "policy_sets", policy_sets)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="policySets")
    def policy_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]]]:
        """
        Authentication policy collection.
        """
        return pulumi.get(self, "policy_sets")

    @policy_sets.setter
    def policy_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DetachUserPolicyOperationPolicySetArgs']]]]):
        pulumi.set(self, "policy_sets", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        User id, the same as the sub-user uin.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class DetachUserPolicyOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetachUserPolicyOperationPolicySetArgs']]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DetachUserPolicyOperation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetachUserPolicyOperationPolicySetArgs']]]] policy_sets: Authentication policy collection.
        :param pulumi.Input[str] user_id: User id, the same as the sub-user uin.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DetachUserPolicyOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DetachUserPolicyOperation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DetachUserPolicyOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DetachUserPolicyOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetachUserPolicyOperationPolicySetArgs']]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DetachUserPolicyOperationArgs.__new__(DetachUserPolicyOperationArgs)

            __props__.__dict__["policy_sets"] = policy_sets
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(DetachUserPolicyOperation, __self__).__init__(
            'tencentcloud:Dlc/detachUserPolicyOperation:DetachUserPolicyOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetachUserPolicyOperationPolicySetArgs']]]]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'DetachUserPolicyOperation':
        """
        Get an existing DetachUserPolicyOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetachUserPolicyOperationPolicySetArgs']]]] policy_sets: Authentication policy collection.
        :param pulumi.Input[str] user_id: User id, the same as the sub-user uin.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DetachUserPolicyOperationState.__new__(_DetachUserPolicyOperationState)

        __props__.__dict__["policy_sets"] = policy_sets
        __props__.__dict__["user_id"] = user_id
        return DetachUserPolicyOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policySets")
    def policy_sets(self) -> pulumi.Output[Optional[Sequence['outputs.DetachUserPolicyOperationPolicySet']]]:
        """
        Authentication policy collection.
        """
        return pulumi.get(self, "policy_sets")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        User id, the same as the sub-user uin.
        """
        return pulumi.get(self, "user_id")
