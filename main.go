package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
)

var log = logrus.New()
var qs = new(QueryStruct)
var isTest = flag.Bool("test", false, "是否开启测试，true -> 开启;false -> 关闭，默认 false")

func init() {
	/*qs.Confget()
	if !qs.Logconfig.LogIs {

	}
	if qs.Logconfig.LogIs {
		hook, err := lSyslog.NewSyslogHook(qs.Logconfig.Type, qs.Logconfig.Ip, syslog.LOG_INFO, qs.Logconfig.Tag)
		if err == nil {
			log.Hooks.Add(hook)
		}
	}*/
	hook, err := lSyslog.NewSyslogHook(qs.Logconfig.Type, qs.Logconfig.Ip, syslog.LOG_INFO, qs.Logconfig.Tag)
	if err == nil {
		log.Hooks.Add(hook)
	}
}

func main() {
	flag.Parse()
	if flag.Parsed() {
		localIp, err := GetOutBoundIP()
		checkError(err)
		qs.mainDomain = getKeys(main_dm)
		qs.subDomain = getKeys(sub_dm)
		qs.accessKeyId = getKeys(ak_id)
		qs.accessSecret = getKeys(ak_sec)
		qs.Ip = localIp
		qs.CreateClient()
		checkError(err)
		qs.queryDomain()

		if len(qs.Body.DomainRecords.Record) == 0 {
			log.Infof("本地ip:%v\t远端不存在此解析\t现在更新!...", localIp)
			qs.addDomain()
			return
		}
		if localIp == qs.Body.DomainRecords.Record[0].Value {
			log.Infof("本地ip:%v\t远端ip:%v\t无需更新", localIp, qs.Body.DomainRecords.Record[0].Value)
			return
		} else {
			log.Infof("本地ip:%v\t远端ip:%v\t现在更新", localIp, qs.Body.DomainRecords.Record[0].Value)
			qs.addDomain()
			return
		}
	}
}

func _main() {
	jss := "{\"headers\":{\"access-control-allow-origin\":\"*\",\"connection\":\"keep-alive\",\"content-length\":\"298\",\"content-type\":\"application/json;charset=utf-8\",\"date\":\"Fri, 25 Nov 2022 07:21:41 GMT\",\"x-acs-request-id\":\"1ED3ED60-D135-5931-9C6A-D262A35EF126\",\"x-acs-trace-id\":\"c3f2893ba00e1537be8de1485022af14\"},\"statusCode\":200,\"body\":{\"DomainRecords\":{\"Record\":[{\"DomainName\":\"ie8.pub\",\"Line\":\"default\",\"Locked\":false,\"RR\":\"*\",\"RecordId\":\"798006993888129024\",\"Status\":\"ENABLE\",\"TTL\":600,\"Type\":\"A\",\"Value\":\"14.145.166.234\",\"Weight\":1}]},\"PageNumber\":1,\"PageSize\":20,\"RequestId\":\"1ED3ED60-D135-5931-9C6A-D262A35EF126\",\"TotalCount\":1}}"
	var data QueryStruct
	if err := json.Unmarshal([]byte(jss), &data); err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}
}
