// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmConditionsTemplateTemplateGroupListConditionResult
    {
        public readonly int AlarmNotifyPeriod;
        public readonly int AlarmNotifyType;
        public readonly string CalcType;
        public readonly string CalcValue;
        public readonly string ContinueTime;
        public readonly int IsAdvanced;
        public readonly int IsOpen;
        public readonly string MetricDisplayName;
        public readonly int MetricId;
        public readonly int Period;
        public readonly string ProductId;
        public readonly int RuleId;
        public readonly string Unit;

        [OutputConstructor]
        private GetAlarmConditionsTemplateTemplateGroupListConditionResult(
            int alarmNotifyPeriod,

            int alarmNotifyType,

            string calcType,

            string calcValue,

            string continueTime,

            int isAdvanced,

            int isOpen,

            string metricDisplayName,

            int metricId,

            int period,

            string productId,

            int ruleId,

            string unit)
        {
            AlarmNotifyPeriod = alarmNotifyPeriod;
            AlarmNotifyType = alarmNotifyType;
            CalcType = calcType;
            CalcValue = calcValue;
            ContinueTime = continueTime;
            IsAdvanced = isAdvanced;
            IsOpen = isOpen;
            MetricDisplayName = metricDisplayName;
            MetricId = metricId;
            Period = period;
            ProductId = productId;
            RuleId = ruleId;
            Unit = unit;
        }
    }
}
