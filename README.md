# 
用golang实现aliddns，同时对certbot进行txt信息更新提交
## 使用方法
### 申请证书
注意，aliddns通过读取config.yaml来决定更新内容
```shell
certbot certonly  -d *.example.com --manual --preferred-challenges dns --dry-run  --manual-auth-hook "aliddns -rt cert -f /xxxx/xxx.yaml"
```