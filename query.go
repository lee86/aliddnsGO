package main

import (
	"encoding/json"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// queryDomain 查询解析记录
func (qs *QueryStruct) queryDomain() {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(qs.mainDomain),
		RRKeyWord:  tea.String(qs.subDomain),
	}
	runtime := &util.RuntimeOptions{}
	if r := tea.Recover(recover()); r != nil {
		log.Fatal(r)
	}
	//json 解析返回参数，获取rrid
	resp, _err := qs.client.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
	checkError(_err)
	jsString := tea.StringValue(util.ToJSONString(resp))
	if *isTest {
		log.Info(jsString)
	}
	err := json.Unmarshal([]byte(jsString), &qs)
	checkError(err)
	//更新luci全局参数
	/*	cmd := exec.Command("/bin/sh", "-c", "uci", "set", "aliddns.base.record_id="+qs.DomainRecords.Record[0].RecordID)
		err = cmd.Start()
		checkError(err)
		commit := exec.Command("/bin/sh", "-c", "uci", "commit", "aliddns")
		err = commit.Start()
		checkError(err)*/
}

// addDomain 添加解析
func (qs *QueryStruct) addDomain() {
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: tea.String(qs.mainDomain),
		RR:         tea.String(qs.subDomain),
		Type:       tea.String("A"),
		Value:      tea.String(qs.Ip),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := qs.client.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
	checkError(_err)
	if *isTest {
		log.Info(tea.StringValue(util.ToJSONString(resp)))
	}
}

// updateDomain 更新解析
func (qs *QueryStruct) updateDomain() {
	UpdateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(qs.Body.DomainRecords.Record[0].RecordID),
		RR:       tea.String(qs.subDomain),
		Type:     tea.String("A"),
		Value:    tea.String(qs.Ip),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := qs.client.UpdateDomainRecordWithOptions(UpdateDomainRecordRequest, runtime)
	checkError(_err)
	if *isTest {
		log.Info(tea.StringValue(util.ToJSONString(resp)))
	}
}

func (qs *QueryStruct) CreateClient() {
	config := &openapi.Config{
		AccessKeyId:     &qs.accessKeyId,
		AccessKeySecret: &qs.accessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	qs.client, _ = alidns20150109.NewClient(config)
}
