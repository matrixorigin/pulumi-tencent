// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AlarmArgs, AlarmState } from "./alarm";
export type Alarm = import("./alarm").Alarm;
export const Alarm: typeof import("./alarm").Alarm = null as any;
utilities.lazyLoad(exports, ["Alarm"], () => require("./alarm"));

export { AlarmNoticeArgs, AlarmNoticeState } from "./alarmNotice";
export type AlarmNotice = import("./alarmNotice").AlarmNotice;
export const AlarmNotice: typeof import("./alarmNotice").AlarmNotice = null as any;
utilities.lazyLoad(exports, ["AlarmNotice"], () => require("./alarmNotice"));

export { CkafkaConsumerArgs, CkafkaConsumerState } from "./ckafkaConsumer";
export type CkafkaConsumer = import("./ckafkaConsumer").CkafkaConsumer;
export const CkafkaConsumer: typeof import("./ckafkaConsumer").CkafkaConsumer = null as any;
utilities.lazyLoad(exports, ["CkafkaConsumer"], () => require("./ckafkaConsumer"));

export { ConfigArgs, ConfigState } from "./config";
export type Config = import("./config").Config;
export const Config: typeof import("./config").Config = null as any;
utilities.lazyLoad(exports, ["Config"], () => require("./config"));

export { ConfigAttachmentArgs, ConfigAttachmentState } from "./configAttachment";
export type ConfigAttachment = import("./configAttachment").ConfigAttachment;
export const ConfigAttachment: typeof import("./configAttachment").ConfigAttachment = null as any;
utilities.lazyLoad(exports, ["ConfigAttachment"], () => require("./configAttachment"));

export { ConfigExtraArgs, ConfigExtraState } from "./configExtra";
export type ConfigExtra = import("./configExtra").ConfigExtra;
export const ConfigExtra: typeof import("./configExtra").ConfigExtra = null as any;
utilities.lazyLoad(exports, ["ConfigExtra"], () => require("./configExtra"));

export { CosRechargeArgs, CosRechargeState } from "./cosRecharge";
export type CosRecharge = import("./cosRecharge").CosRecharge;
export const CosRecharge: typeof import("./cosRecharge").CosRecharge = null as any;
utilities.lazyLoad(exports, ["CosRecharge"], () => require("./cosRecharge"));

export { CosShipperArgs, CosShipperState } from "./cosShipper";
export type CosShipper = import("./cosShipper").CosShipper;
export const CosShipper: typeof import("./cosShipper").CosShipper = null as any;
utilities.lazyLoad(exports, ["CosShipper"], () => require("./cosShipper"));

export { DataTransformArgs, DataTransformState } from "./dataTransform";
export type DataTransform = import("./dataTransform").DataTransform;
export const DataTransform: typeof import("./dataTransform").DataTransform = null as any;
utilities.lazyLoad(exports, ["DataTransform"], () => require("./dataTransform"));

export { ExportArgs, ExportState } from "./export";
export type Export = import("./export").Export;
export const Export: typeof import("./export").Export = null as any;
utilities.lazyLoad(exports, ["Export"], () => require("./export"));

export { GetMachineGroupConfigsArgs, GetMachineGroupConfigsResult, GetMachineGroupConfigsOutputArgs } from "./getMachineGroupConfigs";
export const getMachineGroupConfigs: typeof import("./getMachineGroupConfigs").getMachineGroupConfigs = null as any;
export const getMachineGroupConfigsOutput: typeof import("./getMachineGroupConfigs").getMachineGroupConfigsOutput = null as any;
utilities.lazyLoad(exports, ["getMachineGroupConfigs","getMachineGroupConfigsOutput"], () => require("./getMachineGroupConfigs"));

export { GetMachinesArgs, GetMachinesResult, GetMachinesOutputArgs } from "./getMachines";
export const getMachines: typeof import("./getMachines").getMachines = null as any;
export const getMachinesOutput: typeof import("./getMachines").getMachinesOutput = null as any;
utilities.lazyLoad(exports, ["getMachines","getMachinesOutput"], () => require("./getMachines"));

export { GetShipperTasksArgs, GetShipperTasksResult, GetShipperTasksOutputArgs } from "./getShipperTasks";
export const getShipperTasks: typeof import("./getShipperTasks").getShipperTasks = null as any;
export const getShipperTasksOutput: typeof import("./getShipperTasks").getShipperTasksOutput = null as any;
utilities.lazyLoad(exports, ["getShipperTasks","getShipperTasksOutput"], () => require("./getShipperTasks"));

export { IndexArgs, IndexState } from "./index_";
export type Index = import("./index_").Index;
export const Index: typeof import("./index_").Index = null as any;
utilities.lazyLoad(exports, ["Index"], () => require("./index_"));

export { KafkaRechargeArgs, KafkaRechargeState } from "./kafkaRecharge";
export type KafkaRecharge = import("./kafkaRecharge").KafkaRecharge;
export const KafkaRecharge: typeof import("./kafkaRecharge").KafkaRecharge = null as any;
utilities.lazyLoad(exports, ["KafkaRecharge"], () => require("./kafkaRecharge"));

export { LogsetArgs, LogsetState } from "./logset";
export type Logset = import("./logset").Logset;
export const Logset: typeof import("./logset").Logset = null as any;
utilities.lazyLoad(exports, ["Logset"], () => require("./logset"));

export { MachineGroupArgs, MachineGroupState } from "./machineGroup";
export type MachineGroup = import("./machineGroup").MachineGroup;
export const MachineGroup: typeof import("./machineGroup").MachineGroup = null as any;
utilities.lazyLoad(exports, ["MachineGroup"], () => require("./machineGroup"));

export { ScheduledSqlArgs, ScheduledSqlState } from "./scheduledSql";
export type ScheduledSql = import("./scheduledSql").ScheduledSql;
export const ScheduledSql: typeof import("./scheduledSql").ScheduledSql = null as any;
utilities.lazyLoad(exports, ["ScheduledSql"], () => require("./scheduledSql"));

export { TopicArgs, TopicState } from "./topic";
export type Topic = import("./topic").Topic;
export const Topic: typeof import("./topic").Topic = null as any;
utilities.lazyLoad(exports, ["Topic"], () => require("./topic"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cls/alarm:Alarm":
                return new Alarm(name, <any>undefined, { urn })
            case "tencentcloud:Cls/alarmNotice:AlarmNotice":
                return new AlarmNotice(name, <any>undefined, { urn })
            case "tencentcloud:Cls/ckafkaConsumer:CkafkaConsumer":
                return new CkafkaConsumer(name, <any>undefined, { urn })
            case "tencentcloud:Cls/config:Config":
                return new Config(name, <any>undefined, { urn })
            case "tencentcloud:Cls/configAttachment:ConfigAttachment":
                return new ConfigAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cls/configExtra:ConfigExtra":
                return new ConfigExtra(name, <any>undefined, { urn })
            case "tencentcloud:Cls/cosRecharge:CosRecharge":
                return new CosRecharge(name, <any>undefined, { urn })
            case "tencentcloud:Cls/cosShipper:CosShipper":
                return new CosShipper(name, <any>undefined, { urn })
            case "tencentcloud:Cls/dataTransform:DataTransform":
                return new DataTransform(name, <any>undefined, { urn })
            case "tencentcloud:Cls/export:Export":
                return new Export(name, <any>undefined, { urn })
            case "tencentcloud:Cls/index:Index":
                return new Index(name, <any>undefined, { urn })
            case "tencentcloud:Cls/kafkaRecharge:KafkaRecharge":
                return new KafkaRecharge(name, <any>undefined, { urn })
            case "tencentcloud:Cls/logset:Logset":
                return new Logset(name, <any>undefined, { urn })
            case "tencentcloud:Cls/machineGroup:MachineGroup":
                return new MachineGroup(name, <any>undefined, { urn })
            case "tencentcloud:Cls/scheduledSql:ScheduledSql":
                return new ScheduledSql(name, <any>undefined, { urn })
            case "tencentcloud:Cls/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/alarm", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/alarmNotice", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/ckafkaConsumer", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/config", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/configAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/configExtra", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/cosRecharge", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/cosShipper", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/dataTransform", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/export", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/index", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/kafkaRecharge", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/logset", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/machineGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/scheduledSql", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/topic", _module)