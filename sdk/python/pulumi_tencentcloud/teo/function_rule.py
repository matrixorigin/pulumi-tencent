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

__all__ = ['FunctionRuleArgs', 'FunctionRule']

@pulumi.input_type
class FunctionRuleArgs:
    def __init__(__self__, *,
                 function_id: pulumi.Input[str],
                 function_rule_conditions: pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]],
                 zone_id: pulumi.Input[str],
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FunctionRule resource.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]] function_rule_conditions: The list of rule conditions, where the conditions are connected by an "OR" relationship.
        :param pulumi.Input[str] zone_id: ID of the site.
        :param pulumi.Input[str] remark: Rule description, maximum support of 60 characters.
        """
        pulumi.set(__self__, "function_id", function_id)
        pulumi.set(__self__, "function_rule_conditions", function_rule_conditions)
        pulumi.set(__self__, "zone_id", zone_id)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Input[str]:
        """
        ID of the Function.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="functionRuleConditions")
    def function_rule_conditions(self) -> pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]]:
        """
        The list of rule conditions, where the conditions are connected by an "OR" relationship.
        """
        return pulumi.get(self, "function_rule_conditions")

    @function_rule_conditions.setter
    def function_rule_conditions(self, value: pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]]):
        pulumi.set(self, "function_rule_conditions", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Rule description, maximum support of 60 characters.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _FunctionRuleState:
    def __init__(__self__, *,
                 function_id: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 function_rule_conditions: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FunctionRule resources.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[str] function_name: The name of the function.
        :param pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]] function_rule_conditions: The list of rule conditions, where the conditions are connected by an "OR" relationship.
        :param pulumi.Input[int] priority: The priority of the function trigger rule. A higher numerical value indicates a higher priority.
        :param pulumi.Input[str] remark: Rule description, maximum support of 60 characters.
        :param pulumi.Input[str] rule_id: ID of the Function Rule.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        if function_id is not None:
            pulumi.set(__self__, "function_id", function_id)
        if function_name is not None:
            pulumi.set(__self__, "function_name", function_name)
        if function_rule_conditions is not None:
            pulumi.set(__self__, "function_rule_conditions", function_rule_conditions)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Function.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the function.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter(name="functionRuleConditions")
    def function_rule_conditions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]]]:
        """
        The list of rule conditions, where the conditions are connected by an "OR" relationship.
        """
        return pulumi.get(self, "function_rule_conditions")

    @function_rule_conditions.setter
    def function_rule_conditions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionRuleFunctionRuleConditionArgs']]]]):
        pulumi.set(self, "function_rule_conditions", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the function trigger rule. A higher numerical value indicates a higher priority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Rule description, maximum support of 60 characters.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Function Rule.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class FunctionRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 function_rule_conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionRuleFunctionRuleConditionArgs']]]]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FunctionRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionRuleFunctionRuleConditionArgs']]]] function_rule_conditions: The list of rule conditions, where the conditions are connected by an "OR" relationship.
        :param pulumi.Input[str] remark: Rule description, maximum support of 60 characters.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FunctionRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FunctionRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 function_rule_conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionRuleFunctionRuleConditionArgs']]]]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionRuleArgs.__new__(FunctionRuleArgs)

            if function_id is None and not opts.urn:
                raise TypeError("Missing required property 'function_id'")
            __props__.__dict__["function_id"] = function_id
            if function_rule_conditions is None and not opts.urn:
                raise TypeError("Missing required property 'function_rule_conditions'")
            __props__.__dict__["function_rule_conditions"] = function_rule_conditions
            __props__.__dict__["remark"] = remark
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["function_name"] = None
            __props__.__dict__["priority"] = None
            __props__.__dict__["rule_id"] = None
        super(FunctionRule, __self__).__init__(
            'tencentcloud:Teo/functionRule:FunctionRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            function_id: Optional[pulumi.Input[str]] = None,
            function_name: Optional[pulumi.Input[str]] = None,
            function_rule_conditions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionRuleFunctionRuleConditionArgs']]]]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            rule_id: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'FunctionRule':
        """
        Get an existing FunctionRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] function_id: ID of the Function.
        :param pulumi.Input[str] function_name: The name of the function.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FunctionRuleFunctionRuleConditionArgs']]]] function_rule_conditions: The list of rule conditions, where the conditions are connected by an "OR" relationship.
        :param pulumi.Input[int] priority: The priority of the function trigger rule. A higher numerical value indicates a higher priority.
        :param pulumi.Input[str] remark: Rule description, maximum support of 60 characters.
        :param pulumi.Input[str] rule_id: ID of the Function Rule.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionRuleState.__new__(_FunctionRuleState)

        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["function_name"] = function_name
        __props__.__dict__["function_rule_conditions"] = function_rule_conditions
        __props__.__dict__["priority"] = priority
        __props__.__dict__["remark"] = remark
        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["zone_id"] = zone_id
        return FunctionRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[str]:
        """
        ID of the Function.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        """
        The name of the function.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="functionRuleConditions")
    def function_rule_conditions(self) -> pulumi.Output[Sequence['outputs.FunctionRuleFunctionRuleCondition']]:
        """
        The list of rule conditions, where the conditions are connected by an "OR" relationship.
        """
        return pulumi.get(self, "function_rule_conditions")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        The priority of the function trigger rule. A higher numerical value indicates a higher priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Rule description, maximum support of 60 characters.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[str]:
        """
        ID of the Function Rule.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

