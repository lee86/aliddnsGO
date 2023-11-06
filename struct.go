package main

import alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"

type Account struct {
	AccessKeyId  string `json:"AccessKeyId"`
	AccessSecret string `json:"AccessSecret"`
	MainDomain   string `json:"MainDomain"`
	SubDomain    string `json:"SubDomain"`
}
type QueryStruct struct {
	client       *alidns20150109.Client
	Headers      Headers `json:"headers"`
	StatusCode   int     `json:"statusCode"`
	Body         Body    `json:"body"`
	Action       string  `json:"Action"`
	Signature    string  `json:"Signature"`
	Version      string  `json:"Version"`
	AccessKeyId  string  `json:"AccessKeyId"`
	AccessSecret string  `json:"AccessSecret"`
	MainDomain   string  `json:"MainDomain"`
	SubDomain    string  `json:"SubDomain"`
	Ipv4         string
	Ipv6         string
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
type Body struct {
	DomainRecords DomainRecords `json:"DomainRecords"`
	PageNumber    int           `json:"PageNumber"`
	PageSize      int           `json:"PageSize"`
	RequestID     string        `json:"RequestId"`
	TotalCount    int           `json:"TotalCount"`
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
