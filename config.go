package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func (qs *QueryStruct) Confget() {
	sysconf := "/etc/aliddns/config.yaml"
	yamlFile, err := ioutil.ReadFile(sysconf)
	if err != nil {
		log.Fatal("yaml Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, qs.Logconfig)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatal("Unmarshal: %v", err)
	}
	if *isTest {
		//log.Println(string(yamlFile))
		log.Info("conf: +-+-+-+-+->>>>> ", qs.Logconfig)
	}
}
