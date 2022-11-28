# 根据阿里官方接口查询，添加，更新DNS解析记录

该项目为调用阿里官方接口查询，添加，更新DNS解析记录，由于原aliddns插件无法添加*泛解析，遂使用golang参考官方接口开发

## 运行条件

- 仓库中的aliddns架构为AMD64，平台linux，x86_64软路由可直接替换运行，其余架构请自行编译

- openwrt 安装aliddns，ddns插件

- 在阿里云帐户中获取您的 [凭证](https://usercenter.console.aliyun.com/#/manage/ak)并通过它替换aliddns插件中的
  ACCESS_KEY_ID 以及 ACCESS_KEY_SECRET;

- 替换 /usr/sbin/aliddns,授权运行

## 使用的 API

- AddDomainRecord
  调用AddDomainRecord根据传入参数添加解析记录。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/Alidns/2015-01-09/AddDomainRecord)
- UpdateDomainRecord
  调用UpdateDomainRecord根据传入参数更新解析记录。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/Alidns/2015-01-09/UpdateDomainRecord)
- DescribeDomainRecords
  调用DescribeDomainRecords根据传入参数查询解析记录。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/Alidns/2015-01-09/DescribeDomainRecords)
