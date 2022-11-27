package main

import alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"

type QueryStruct struct {
	client       *alidns20150109.Client
	Headers      Headers `json:"headers"`
	StatusCode   int     `json:"statusCode"`
	Body         Body    `json:"body"`
	accessKeyId  string  `json:"AccessKeyId"`
	accessSecret string  `json:"AccessSecret"`
	mainDomain   string  `json:"MainDomain"`
	subDomain    string  `json:"SubDomain"`
	Action       string  `json:"Action"`
	Signature    string  `json:"Signature"`
	Version      string  `json:"Version"`
	Ip           string
	Logconfig    LogConfig
}
type Headers struct {
	PageNumber               int    `json:"PageNumber"`
	AccessControlAllowOrigin string `json:"access-control-allow-origin"`
	Connection               string `json:"connection"`
	ContentLength            string `json:"content-length"`
	ContentType              string `json:"content-type"`
	Date                     string `json:"date"`
	XAcsRequestID            string `json:"x-acs-request-id"`
	XAcsTraceID              string `json:"x-acs-trace-id"`
}
type Record struct {
	DomainName string `json:"DomainName"`
	Line       string `json:"Line"`
	Locked     bool   `json:"Locked"`
	RR         string `json:"RR"`
	RecordID   string `json:"RecordId"`
	Status     string `json:"Status"`
	TTL        int    `json:"TTL"`
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	Weight     int    `json:"Weight"`
}
type DomainRecords struct {
	Record []Record `json:"Record"`
}
type Body struct {
	DomainRecords DomainRecords `json:"DomainRecords"`
	PageNumber    int           `json:"PageNumber"`
	PageSize      int           `json:"PageSize"`
	RequestID     string        `json:"RequestId"`
	TotalCount    int           `json:"TotalCount"`
}

type LogConfig struct {
	Ip    string `yaml:"ip"`
	Type  string `yaml:"type"`
	LogIs bool   `yaml:"logis"`
	Tag   string `yaml:"tag"`
}
