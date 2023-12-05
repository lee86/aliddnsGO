package aliddns

import (
	"encoding/json"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	log "github.com/sirupsen/logrus"
)

// CreateClient 创建client
func (qs *QueryStruct) createClient() {
	config := &openapi.Config{
		AccessKeyId:     &qs.AccessKeyId,
		AccessKeySecret: &qs.AccessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	qs.client, _ = alidns20150109.NewClient(config)
}

// queryDomain 查询解析记录
func (qs *QueryStruct) queryDomain() {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(qs.MainDomain),
		RRKeyWord:  tea.String(qs.SubDomain),
		Type:       tea.String(qs.ValueType),
	}
	runtime := &util.RuntimeOptions{}
	if r := tea.Recover(recover()); r != nil {
		log.Fatal(r)
	}
	//json 解析返回参数，获取rrid
	resp, _err := qs.client.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
	checkError(_err)
	jsString := tea.StringValue(util.ToJSONString(resp))
	if qs.IsTest {
		log.Info(jsString)
	}
	err := json.Unmarshal([]byte(jsString), &qs)
	checkError(err)
}

// addDomain 添加解析
func (qs *QueryStruct) addDomain() {
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: tea.String(qs.MainDomain),
		RR:         tea.String(qs.SubDomain),
		Type:       tea.String(qs.ValueType),
		Value:      tea.String(qs.Value),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := qs.client.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
	checkError(_err)
	if qs.IsTest {
		log.Info(tea.StringValue(util.ToJSONString(resp)))
	}
}

// updateDomain 更新解析
func (qs *QueryStruct) updateDomain() {
	UpdateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(qs.Body.DomainRecords.Record[0].RecordID),
		RR:       tea.String(qs.SubDomain),
		Type:     tea.String(qs.ValueType),
		Value:    tea.String(qs.Value),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := qs.client.UpdateDomainRecordWithOptions(UpdateDomainRecordRequest, runtime)
	checkError(_err)
	if qs.IsTest {
		log.Info(tea.StringValue(util.ToJSONString(resp)))
	}
}
