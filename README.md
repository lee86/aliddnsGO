#   
[![Security Status](https://www.murphysec.com/platform3/v31/badge/1726604684981915648.svg)](https://www.murphysec.com/console/report/1721468007020584960/1726604684981915648)
用golang实现aliddns，同时对certbot进行txt信息更新提交

## 更新ipv6

```shell
aliddns -rt ipv6 -f /xxxx/xxxx.yaml # 默认配置文件 /etc/ddns/config.yaml
```

```shell
aliddns # 如果采用默认配置，且目标为更新ddns，则只需要此条命令
```

## 证书相关

### 申请证书

注意，aliddns通过读取config.yaml来决定更新内容

```shell
certbot certonly  -d *.example.com --manual --preferred-challenges dns --dry-run  --manual-auth-hook "aliddns -rt cert -f /xxxx/xxx.yaml"
```

### 更新证书(同时重启gitlab nginx服务)

```shell
certbot renew  --manual --preferred-challenges dns --manual-auth-hook "aliddns -rt cert" --deploy-hook "gitlab-ctl restart nginx"
```