// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

import (
	"context"
	"reflect"

	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRecords(ctx *pulumi.Context, args *LookupRecordsArgs, opts ...pulumi.InvokeOption) (*LookupRecordsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRecordsResult
	err := ctx.Invoke("tencentcloud:Dnspod/getRecords:getRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecords.
type LookupRecordsArgs struct {
	Domain           *string `pulumi:"domain"`
	DomainId         *string `pulumi:"domainId"`
	GroupId          *string `pulumi:"groupId"`
	Keyword          *string `pulumi:"keyword"`
	Limit            *int    `pulumi:"limit"`
	Offset           *int    `pulumi:"offset"`
	RecordLine       *string `pulumi:"recordLine"`
	RecordLineId     *string `pulumi:"recordLineId"`
	RecordType       *string `pulumi:"recordType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SortField        *string `pulumi:"sortField"`
	SortType         *string `pulumi:"sortType"`
	Subdomain        *string `pulumi:"subdomain"`
}

// A collection of values returned by getRecords.
type LookupRecordsResult struct {
	Domain   *string `pulumi:"domain"`
	DomainId *string `pulumi:"domainId"`
	GroupId  *string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                      `pulumi:"id"`
	Keyword          *string                     `pulumi:"keyword"`
	Limit            *int                        `pulumi:"limit"`
	Offset           *int                        `pulumi:"offset"`
	RecordCountInfos []GetRecordsRecordCountInfo `pulumi:"recordCountInfos"`
	RecordLine       *string                     `pulumi:"recordLine"`
	RecordLineId     *string                     `pulumi:"recordLineId"`
	RecordType       *string                     `pulumi:"recordType"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
	Results          []GetRecordsResult          `pulumi:"results"`
	SortField        *string                     `pulumi:"sortField"`
	SortType         *string                     `pulumi:"sortType"`
	Subdomain        *string                     `pulumi:"subdomain"`
}

func LookupRecordsOutput(ctx *pulumi.Context, args LookupRecordsOutputArgs, opts ...pulumi.InvokeOption) LookupRecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRecordsResult, error) {
			args := v.(LookupRecordsArgs)
			r, err := LookupRecords(ctx, &args, opts...)
			var s LookupRecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRecordsResultOutput)
}

// A collection of arguments for invoking getRecords.
type LookupRecordsOutputArgs struct {
	Domain           pulumi.StringPtrInput `pulumi:"domain"`
	DomainId         pulumi.StringPtrInput `pulumi:"domainId"`
	GroupId          pulumi.StringPtrInput `pulumi:"groupId"`
	Keyword          pulumi.StringPtrInput `pulumi:"keyword"`
	Limit            pulumi.IntPtrInput    `pulumi:"limit"`
	Offset           pulumi.IntPtrInput    `pulumi:"offset"`
	RecordLine       pulumi.StringPtrInput `pulumi:"recordLine"`
	RecordLineId     pulumi.StringPtrInput `pulumi:"recordLineId"`
	RecordType       pulumi.StringPtrInput `pulumi:"recordType"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	SortField        pulumi.StringPtrInput `pulumi:"sortField"`
	SortType         pulumi.StringPtrInput `pulumi:"sortType"`
	Subdomain        pulumi.StringPtrInput `pulumi:"subdomain"`
}

func (LookupRecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordsArgs)(nil)).Elem()
}

// A collection of values returned by getRecords.
type LookupRecordsResultOutput struct{ *pulumi.OutputState }

func (LookupRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordsResult)(nil)).Elem()
}

func (o LookupRecordsResultOutput) ToLookupRecordsResultOutput() LookupRecordsResultOutput {
	return o
}

func (o LookupRecordsResultOutput) ToLookupRecordsResultOutputWithContext(ctx context.Context) LookupRecordsResultOutput {
	return o
}

func (o LookupRecordsResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRecordsResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o LookupRecordsResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

func (o LookupRecordsResultOutput) RecordCountInfos() GetRecordsRecordCountInfoArrayOutput {
	return o.ApplyT(func(v LookupRecordsResult) []GetRecordsRecordCountInfo { return v.RecordCountInfos }).(GetRecordsRecordCountInfoArrayOutput)
}

func (o LookupRecordsResultOutput) RecordLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.RecordLine }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) RecordLineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.RecordLineId }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) RecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.RecordType }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) Results() GetRecordsResultArrayOutput {
	return o.ApplyT(func(v LookupRecordsResult) []GetRecordsResult { return v.Results }).(GetRecordsResultArrayOutput)
}

func (o LookupRecordsResultOutput) SortField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.SortField }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) SortType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.SortType }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.Subdomain }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordsResultOutput{})
}