package aliddns

import log "github.com/sirupsen/logrus"

// DnsCheck 与返回的解析内容对比，确认是否需要更新、新增操作
func (qs *QueryStruct) DnsCheck() {
	qs.createClient()
	qs.queryDomain()
	if len(qs.Body.DomainRecords.Record) == 0 {
		log.Infof("类型:%v\n本地解析结果:%v\n远端不存在此解析\n现在开始添加!...", qs.ValueType, qs.Value)
		qs.addDomain()
		return
	}
	for _, record := range qs.Body.DomainRecords.Record {
		if qs.Value == record.Value {
			log.Infof("类型:%v\n本地解析结果:%v\n远端解析结果:%v\n无需更新", qs.ValueType, qs.Value, record.Value)
			return
		} else {
			log.Infof("类型:%v\n本地解析结果:%v\n远端解析结果:%v\n现在更新", qs.ValueType, qs.Value, record.Value)
			qs.updateDomain()
			return
		}
	}
}
