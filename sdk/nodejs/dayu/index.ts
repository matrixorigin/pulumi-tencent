// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CcHttpPolicyArgs, CcHttpPolicyState } from "./ccHttpPolicy";
export type CcHttpPolicy = import("./ccHttpPolicy").CcHttpPolicy;
export const CcHttpPolicy: typeof import("./ccHttpPolicy").CcHttpPolicy = null as any;
utilities.lazyLoad(exports, ["CcHttpPolicy"], () => require("./ccHttpPolicy"));

export { CcHttpsPolicyArgs, CcHttpsPolicyState } from "./ccHttpsPolicy";
export type CcHttpsPolicy = import("./ccHttpsPolicy").CcHttpsPolicy;
export const CcHttpsPolicy: typeof import("./ccHttpsPolicy").CcHttpsPolicy = null as any;
utilities.lazyLoad(exports, ["CcHttpsPolicy"], () => require("./ccHttpsPolicy"));

export { CcPolicyV2Args, CcPolicyV2State } from "./ccPolicyV2";
export type CcPolicyV2 = import("./ccPolicyV2").CcPolicyV2;
export const CcPolicyV2: typeof import("./ccPolicyV2").CcPolicyV2 = null as any;
utilities.lazyLoad(exports, ["CcPolicyV2"], () => require("./ccPolicyV2"));

export { DdosIpAttachmentV2Args, DdosIpAttachmentV2State } from "./ddosIpAttachmentV2";
export type DdosIpAttachmentV2 = import("./ddosIpAttachmentV2").DdosIpAttachmentV2;
export const DdosIpAttachmentV2: typeof import("./ddosIpAttachmentV2").DdosIpAttachmentV2 = null as any;
utilities.lazyLoad(exports, ["DdosIpAttachmentV2"], () => require("./ddosIpAttachmentV2"));

export { DdosPolicyArgs, DdosPolicyState } from "./ddosPolicy";
export type DdosPolicy = import("./ddosPolicy").DdosPolicy;
export const DdosPolicy: typeof import("./ddosPolicy").DdosPolicy = null as any;
utilities.lazyLoad(exports, ["DdosPolicy"], () => require("./ddosPolicy"));

export { DdosPolicyAttachmentArgs, DdosPolicyAttachmentState } from "./ddosPolicyAttachment";
export type DdosPolicyAttachment = import("./ddosPolicyAttachment").DdosPolicyAttachment;
export const DdosPolicyAttachment: typeof import("./ddosPolicyAttachment").DdosPolicyAttachment = null as any;
utilities.lazyLoad(exports, ["DdosPolicyAttachment"], () => require("./ddosPolicyAttachment"));

export { DdosPolicyCaseArgs, DdosPolicyCaseState } from "./ddosPolicyCase";
export type DdosPolicyCase = import("./ddosPolicyCase").DdosPolicyCase;
export const DdosPolicyCase: typeof import("./ddosPolicyCase").DdosPolicyCase = null as any;
utilities.lazyLoad(exports, ["DdosPolicyCase"], () => require("./ddosPolicyCase"));

export { DdosPolicyV2Args, DdosPolicyV2State } from "./ddosPolicyV2";
export type DdosPolicyV2 = import("./ddosPolicyV2").DdosPolicyV2;
export const DdosPolicyV2: typeof import("./ddosPolicyV2").DdosPolicyV2 = null as any;
utilities.lazyLoad(exports, ["DdosPolicyV2"], () => require("./ddosPolicyV2"));

export { EipArgs, EipState } from "./eip";
export type Eip = import("./eip").Eip;
export const Eip: typeof import("./eip").Eip = null as any;
utilities.lazyLoad(exports, ["Eip"], () => require("./eip"));

export { GetCcHttpPoliciesArgs, GetCcHttpPoliciesResult, GetCcHttpPoliciesOutputArgs } from "./getCcHttpPolicies";
export const getCcHttpPolicies: typeof import("./getCcHttpPolicies").getCcHttpPolicies = null as any;
export const getCcHttpPoliciesOutput: typeof import("./getCcHttpPolicies").getCcHttpPoliciesOutput = null as any;
utilities.lazyLoad(exports, ["getCcHttpPolicies","getCcHttpPoliciesOutput"], () => require("./getCcHttpPolicies"));

export { GetCcHttpsPoliciesArgs, GetCcHttpsPoliciesResult, GetCcHttpsPoliciesOutputArgs } from "./getCcHttpsPolicies";
export const getCcHttpsPolicies: typeof import("./getCcHttpsPolicies").getCcHttpsPolicies = null as any;
export const getCcHttpsPoliciesOutput: typeof import("./getCcHttpsPolicies").getCcHttpsPoliciesOutput = null as any;
utilities.lazyLoad(exports, ["getCcHttpsPolicies","getCcHttpsPoliciesOutput"], () => require("./getCcHttpsPolicies"));

export { GetDdosPoliciesArgs, GetDdosPoliciesResult, GetDdosPoliciesOutputArgs } from "./getDdosPolicies";
export const getDdosPolicies: typeof import("./getDdosPolicies").getDdosPolicies = null as any;
export const getDdosPoliciesOutput: typeof import("./getDdosPolicies").getDdosPoliciesOutput = null as any;
utilities.lazyLoad(exports, ["getDdosPolicies","getDdosPoliciesOutput"], () => require("./getDdosPolicies"));

export { GetDdosPolicyAttachmentsArgs, GetDdosPolicyAttachmentsResult, GetDdosPolicyAttachmentsOutputArgs } from "./getDdosPolicyAttachments";
export const getDdosPolicyAttachments: typeof import("./getDdosPolicyAttachments").getDdosPolicyAttachments = null as any;
export const getDdosPolicyAttachmentsOutput: typeof import("./getDdosPolicyAttachments").getDdosPolicyAttachmentsOutput = null as any;
utilities.lazyLoad(exports, ["getDdosPolicyAttachments","getDdosPolicyAttachmentsOutput"], () => require("./getDdosPolicyAttachments"));

export { GetDdosPolicyCasesArgs, GetDdosPolicyCasesResult, GetDdosPolicyCasesOutputArgs } from "./getDdosPolicyCases";
export const getDdosPolicyCases: typeof import("./getDdosPolicyCases").getDdosPolicyCases = null as any;
export const getDdosPolicyCasesOutput: typeof import("./getDdosPolicyCases").getDdosPolicyCasesOutput = null as any;
utilities.lazyLoad(exports, ["getDdosPolicyCases","getDdosPolicyCasesOutput"], () => require("./getDdosPolicyCases"));

export { GetEipArgs, GetEipResult, GetEipOutputArgs } from "./getEip";
export const getEip: typeof import("./getEip").getEip = null as any;
export const getEipOutput: typeof import("./getEip").getEipOutput = null as any;
utilities.lazyLoad(exports, ["getEip","getEipOutput"], () => require("./getEip"));

export { GetL4RulesArgs, GetL4RulesResult, GetL4RulesOutputArgs } from "./getL4Rules";
export const getL4Rules: typeof import("./getL4Rules").getL4Rules = null as any;
export const getL4RulesOutput: typeof import("./getL4Rules").getL4RulesOutput = null as any;
utilities.lazyLoad(exports, ["getL4Rules","getL4RulesOutput"], () => require("./getL4Rules"));

export { GetL4RulesV2Args, GetL4RulesV2Result, GetL4RulesV2OutputArgs } from "./getL4RulesV2";
export const getL4RulesV2: typeof import("./getL4RulesV2").getL4RulesV2 = null as any;
export const getL4RulesV2Output: typeof import("./getL4RulesV2").getL4RulesV2Output = null as any;
utilities.lazyLoad(exports, ["getL4RulesV2","getL4RulesV2Output"], () => require("./getL4RulesV2"));

export { GetL7RulesArgs, GetL7RulesResult, GetL7RulesOutputArgs } from "./getL7Rules";
export const getL7Rules: typeof import("./getL7Rules").getL7Rules = null as any;
export const getL7RulesOutput: typeof import("./getL7Rules").getL7RulesOutput = null as any;
utilities.lazyLoad(exports, ["getL7Rules","getL7RulesOutput"], () => require("./getL7Rules"));

export { GetL7RulesV2Args, GetL7RulesV2Result, GetL7RulesV2OutputArgs } from "./getL7RulesV2";
export const getL7RulesV2: typeof import("./getL7RulesV2").getL7RulesV2 = null as any;
export const getL7RulesV2Output: typeof import("./getL7RulesV2").getL7RulesV2Output = null as any;
utilities.lazyLoad(exports, ["getL7RulesV2","getL7RulesV2Output"], () => require("./getL7RulesV2"));

export { L4RuleArgs, L4RuleState } from "./l4rule";
export type L4Rule = import("./l4rule").L4Rule;
export const L4Rule: typeof import("./l4rule").L4Rule = null as any;
utilities.lazyLoad(exports, ["L4Rule"], () => require("./l4rule"));

export { L4RuleV2Args, L4RuleV2State } from "./l4ruleV2";
export type L4RuleV2 = import("./l4ruleV2").L4RuleV2;
export const L4RuleV2: typeof import("./l4ruleV2").L4RuleV2 = null as any;
utilities.lazyLoad(exports, ["L4RuleV2"], () => require("./l4ruleV2"));

export { L7RuleArgs, L7RuleState } from "./l7rule";
export type L7Rule = import("./l7rule").L7Rule;
export const L7Rule: typeof import("./l7rule").L7Rule = null as any;
utilities.lazyLoad(exports, ["L7Rule"], () => require("./l7rule"));

export { L7RuleV2Args, L7RuleV2State } from "./l7ruleV2";
export type L7RuleV2 = import("./l7ruleV2").L7RuleV2;
export const L7RuleV2: typeof import("./l7ruleV2").L7RuleV2 = null as any;
utilities.lazyLoad(exports, ["L7RuleV2"], () => require("./l7ruleV2"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Dayu/ccHttpPolicy:CcHttpPolicy":
                return new CcHttpPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ccHttpsPolicy:CcHttpsPolicy":
                return new CcHttpsPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ccPolicyV2:CcPolicyV2":
                return new CcPolicyV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosIpAttachmentV2:DdosIpAttachmentV2":
                return new DdosIpAttachmentV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicy:DdosPolicy":
                return new DdosPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment":
                return new DdosPolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicyCase:DdosPolicyCase":
                return new DdosPolicyCase(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicyV2:DdosPolicyV2":
                return new DdosPolicyV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/eip:Eip":
                return new Eip(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l4Rule:L4Rule":
                return new L4Rule(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l4RuleV2:L4RuleV2":
                return new L4RuleV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l7Rule:L7Rule":
                return new L7Rule(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l7RuleV2:L7RuleV2":
                return new L7RuleV2(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ccHttpPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ccHttpsPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ccPolicyV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosIpAttachmentV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicyCase", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicyV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/eip", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l4Rule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l4RuleV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l7Rule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l7RuleV2", _module)