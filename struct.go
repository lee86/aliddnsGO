package main

type Account struct {
	AccessKeyId  string `json:"AccessKeyId"`
	AccessSecret string `json:"AccessSecret"`
	MainDomain   string `json:"MainDomain"`
	SubDomain    string `json:"SubDomain"`
}
