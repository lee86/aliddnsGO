package aliddns

import alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"

type QueryStruct struct {
	client       *alidns20150109.Client
	Headers      Headers `json:"headers,omitempty"`
	StatusCode   int     `json:"statusCode,omitempty"`
	Body         Body    `json:"body,omitempty"`
	Action       string  `json:"Action,omitempty"`
	Signature    string  `json:"Signature,omitempty"`
	Version      string  `json:"Version,omitempty"`
	AccessKeyId  string  `json:"AccessKeyId,omitempty"`
	AccessSecret string  `json:"AccessSecret,omitempty"`
	MainDomain   string  `json:"MainDomain,omitempty"`
	SubDomain    string  `json:"SubDomain,omitempty"`
	Value        string
	ValueType    string
	IsTest       bool
}
type Headers struct {
	PageNumber               int    `json:"PageNumber,omitempty"`
	AccessControlAllowOrigin string `json:"access-control-allow-origin,omitempty"`
	Connection               string `json:"connection,omitempty"`
	ContentLength            string `json:"content-length,omitempty"`
	ContentType              string `json:"content-type,omitempty"`
	Date                     string `json:"date,omitempty"`
	XAcsRequestID            string `json:"x-acs-request-id,omitempty"`
	XAcsTraceID              string `json:"x-acs-trace-id,omitempty"`
}
type Body struct {
	DomainRecords DomainRecords `json:"DomainRecords,omitempty"`
	PageNumber    int           `json:"PageNumber,omitempty"`
	PageSize      int           `json:"PageSize,omitempty"`
	RequestID     string        `json:"RequestId,omitempty"`
	TotalCount    int           `json:"TotalCount,omitempty"`
}
type Record struct {
	DomainName string `json:"DomainName,omitempty"`
	Line       string `json:"Line,omitempty"`
	Locked     bool   `json:"Locked,omitempty"`
	RR         string `json:"RR,omitempty"`
	RecordID   string `json:"RecordId,omitempty"`
	Status     string `json:"Status,omitempty"`
	TTL        int    `json:"TTL,omitempty"`
	Type       string `json:"Type,omitempty"`
	Value      string `json:"Value,omitempty"`
	Weight     int    `json:"Weight,omitempty"`
}
type DomainRecords struct {
	Record []Record `json:"Record,omitempty"`
}
