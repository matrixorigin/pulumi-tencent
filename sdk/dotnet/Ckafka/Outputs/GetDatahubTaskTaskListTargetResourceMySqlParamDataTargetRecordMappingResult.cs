// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetDatahubTaskTaskListTargetResourceMySqlParamDataTargetRecordMappingResult
    {
        public readonly bool AllowNull;
        public readonly bool AutoIncrement;
        public readonly string ColumnName;
        public readonly string ColumnSize;
        public readonly string DecimalDigits;
        public readonly string DefaultValue;
        public readonly string ExtraInfo;
        public readonly string JsonKey;
        public readonly string Type;

        [OutputConstructor]
        private GetDatahubTaskTaskListTargetResourceMySqlParamDataTargetRecordMappingResult(
            bool allowNull,

            bool autoIncrement,

            string columnName,

            string columnSize,

            string decimalDigits,

            string defaultValue,

            string extraInfo,

            string jsonKey,

            string type)
        {
            AllowNull = allowNull;
            AutoIncrement = autoIncrement;
            ColumnName = columnName;
            ColumnSize = columnSize;
            DecimalDigits = decimalDigits;
            DefaultValue = defaultValue;
            ExtraInfo = extraInfo;
            JsonKey = jsonKey;
            Type = type;
        }
    }
}