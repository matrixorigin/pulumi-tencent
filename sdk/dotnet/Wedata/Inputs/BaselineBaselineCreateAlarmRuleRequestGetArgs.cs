// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class BaselineBaselineCreateAlarmRuleRequestGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("alarmLevel")]
        public Input<int>? AlarmLevel { get; set; }

        [Input("alarmRecipientIds")]
        private InputList<string>? _alarmRecipientIds;
        public InputList<string> AlarmRecipientIds
        {
            get => _alarmRecipientIds ?? (_alarmRecipientIds = new InputList<string>());
            set => _alarmRecipientIds = value;
        }

        [Input("alarmRecipientType")]
        public Input<int>? AlarmRecipientType { get; set; }

        [Input("alarmRecipients")]
        private InputList<string>? _alarmRecipients;
        public InputList<string> AlarmRecipients
        {
            get => _alarmRecipients ?? (_alarmRecipients = new InputList<string>());
            set => _alarmRecipients = value;
        }

        [Input("alarmTypes")]
        private InputList<string>? _alarmTypes;
        public InputList<string> AlarmTypes
        {
            get => _alarmTypes ?? (_alarmTypes = new InputList<string>());
            set => _alarmTypes = value;
        }

        [Input("alarmWays")]
        private InputList<string>? _alarmWays;
        public InputList<string> AlarmWays
        {
            get => _alarmWays ?? (_alarmWays = new InputList<string>());
            set => _alarmWays = value;
        }

        [Input("creator")]
        public Input<string>? Creator { get; set; }

        [Input("creatorId")]
        public Input<string>? CreatorId { get; set; }

        [Input("extInfo")]
        public Input<string>? ExtInfo { get; set; }

        [Input("monitorObjectIds")]
        private InputList<string>? _monitorObjectIds;
        public InputList<string> MonitorObjectIds
        {
            get => _monitorObjectIds ?? (_monitorObjectIds = new InputList<string>());
            set => _monitorObjectIds = value;
        }

        [Input("monitorType")]
        public Input<int>? MonitorType { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        public BaselineBaselineCreateAlarmRuleRequestGetArgs()
        {
        }
        public static new BaselineBaselineCreateAlarmRuleRequestGetArgs Empty => new BaselineBaselineCreateAlarmRuleRequestGetArgs();
    }
}