// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AddressTemplateArgs, AddressTemplateState } from "./addressTemplate";
export type AddressTemplate = import("./addressTemplate").AddressTemplate;
export const AddressTemplate: typeof import("./addressTemplate").AddressTemplate = null as any;
utilities.lazyLoad(exports, ["AddressTemplate"], () => require("./addressTemplate"));

export { BlockIgnoreArgs, BlockIgnoreState } from "./blockIgnore";
export type BlockIgnore = import("./blockIgnore").BlockIgnore;
export const BlockIgnore: typeof import("./blockIgnore").BlockIgnore = null as any;
utilities.lazyLoad(exports, ["BlockIgnore"], () => require("./blockIgnore"));

export { EdgeFirewallSwitchArgs, EdgeFirewallSwitchState } from "./edgeFirewallSwitch";
export type EdgeFirewallSwitch = import("./edgeFirewallSwitch").EdgeFirewallSwitch;
export const EdgeFirewallSwitch: typeof import("./edgeFirewallSwitch").EdgeFirewallSwitch = null as any;
utilities.lazyLoad(exports, ["EdgeFirewallSwitch"], () => require("./edgeFirewallSwitch"));

export { EdgePolicyArgs, EdgePolicyState } from "./edgePolicy";
export type EdgePolicy = import("./edgePolicy").EdgePolicy;
export const EdgePolicy: typeof import("./edgePolicy").EdgePolicy = null as any;
utilities.lazyLoad(exports, ["EdgePolicy"], () => require("./edgePolicy"));

export { GetEdgeFwSwitchesArgs, GetEdgeFwSwitchesResult, GetEdgeFwSwitchesOutputArgs } from "./getEdgeFwSwitches";
export const getEdgeFwSwitches: typeof import("./getEdgeFwSwitches").getEdgeFwSwitches = null as any;
export const getEdgeFwSwitchesOutput: typeof import("./getEdgeFwSwitches").getEdgeFwSwitchesOutput = null as any;
utilities.lazyLoad(exports, ["getEdgeFwSwitches","getEdgeFwSwitchesOutput"], () => require("./getEdgeFwSwitches"));

export { GetNatFwSwitchesArgs, GetNatFwSwitchesResult, GetNatFwSwitchesOutputArgs } from "./getNatFwSwitches";
export const getNatFwSwitches: typeof import("./getNatFwSwitches").getNatFwSwitches = null as any;
export const getNatFwSwitchesOutput: typeof import("./getNatFwSwitches").getNatFwSwitchesOutput = null as any;
utilities.lazyLoad(exports, ["getNatFwSwitches","getNatFwSwitchesOutput"], () => require("./getNatFwSwitches"));

export { GetVpcFwSwitchesArgs, GetVpcFwSwitchesResult, GetVpcFwSwitchesOutputArgs } from "./getVpcFwSwitches";
export const getVpcFwSwitches: typeof import("./getVpcFwSwitches").getVpcFwSwitches = null as any;
export const getVpcFwSwitchesOutput: typeof import("./getVpcFwSwitches").getVpcFwSwitchesOutput = null as any;
utilities.lazyLoad(exports, ["getVpcFwSwitches","getVpcFwSwitchesOutput"], () => require("./getVpcFwSwitches"));

export { NatFirewallSwitchArgs, NatFirewallSwitchState } from "./natFirewallSwitch";
export type NatFirewallSwitch = import("./natFirewallSwitch").NatFirewallSwitch;
export const NatFirewallSwitch: typeof import("./natFirewallSwitch").NatFirewallSwitch = null as any;
utilities.lazyLoad(exports, ["NatFirewallSwitch"], () => require("./natFirewallSwitch"));

export { NatInstanceArgs, NatInstanceState } from "./natInstance";
export type NatInstance = import("./natInstance").NatInstance;
export const NatInstance: typeof import("./natInstance").NatInstance = null as any;
utilities.lazyLoad(exports, ["NatInstance"], () => require("./natInstance"));

export { NatPolicyArgs, NatPolicyState } from "./natPolicy";
export type NatPolicy = import("./natPolicy").NatPolicy;
export const NatPolicy: typeof import("./natPolicy").NatPolicy = null as any;
utilities.lazyLoad(exports, ["NatPolicy"], () => require("./natPolicy"));

export { SyncAssetArgs, SyncAssetState } from "./syncAsset";
export type SyncAsset = import("./syncAsset").SyncAsset;
export const SyncAsset: typeof import("./syncAsset").SyncAsset = null as any;
utilities.lazyLoad(exports, ["SyncAsset"], () => require("./syncAsset"));

export { SyncRouteArgs, SyncRouteState } from "./syncRoute";
export type SyncRoute = import("./syncRoute").SyncRoute;
export const SyncRoute: typeof import("./syncRoute").SyncRoute = null as any;
utilities.lazyLoad(exports, ["SyncRoute"], () => require("./syncRoute"));

export { VpcFirewallSwitchArgs, VpcFirewallSwitchState } from "./vpcFirewallSwitch";
export type VpcFirewallSwitch = import("./vpcFirewallSwitch").VpcFirewallSwitch;
export const VpcFirewallSwitch: typeof import("./vpcFirewallSwitch").VpcFirewallSwitch = null as any;
utilities.lazyLoad(exports, ["VpcFirewallSwitch"], () => require("./vpcFirewallSwitch"));

export { VpcInstanceArgs, VpcInstanceState } from "./vpcInstance";
export type VpcInstance = import("./vpcInstance").VpcInstance;
export const VpcInstance: typeof import("./vpcInstance").VpcInstance = null as any;
utilities.lazyLoad(exports, ["VpcInstance"], () => require("./vpcInstance"));

export { VpcPolicyArgs, VpcPolicyState } from "./vpcPolicy";
export type VpcPolicy = import("./vpcPolicy").VpcPolicy;
export const VpcPolicy: typeof import("./vpcPolicy").VpcPolicy = null as any;
utilities.lazyLoad(exports, ["VpcPolicy"], () => require("./vpcPolicy"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cfw/addressTemplate:AddressTemplate":
                return new AddressTemplate(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/blockIgnore:BlockIgnore":
                return new BlockIgnore(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/edgeFirewallSwitch:EdgeFirewallSwitch":
                return new EdgeFirewallSwitch(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/edgePolicy:EdgePolicy":
                return new EdgePolicy(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/natFirewallSwitch:NatFirewallSwitch":
                return new NatFirewallSwitch(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/natInstance:NatInstance":
                return new NatInstance(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/natPolicy:NatPolicy":
                return new NatPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/syncAsset:SyncAsset":
                return new SyncAsset(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/syncRoute:SyncRoute":
                return new SyncRoute(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/vpcFirewallSwitch:VpcFirewallSwitch":
                return new VpcFirewallSwitch(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/vpcInstance:VpcInstance":
                return new VpcInstance(name, <any>undefined, { urn })
            case "tencentcloud:Cfw/vpcPolicy:VpcPolicy":
                return new VpcPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/addressTemplate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/blockIgnore", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/edgeFirewallSwitch", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/edgePolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/natFirewallSwitch", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/natInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/natPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/syncAsset", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/syncRoute", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/vpcFirewallSwitch", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/vpcInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cfw/vpcPolicy", _module)