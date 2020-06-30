http 反向代理 灰度转发
----

## 部署

部署文件 请看 `example/proxy-grey.yaml`  
需要有 `nginx` namespace

## 配置文件
配置文件是采用 `ConfigMap` 挂载的形式

需要新建一个 `grey-conf` 的 `ConfigMap`  
它由一个`key` 为 `config.yaml`  
`value` 就是配置文件的内容

请参考 默认配置文件 `example/config.yaml`

## 配置
### header
语法 
- `header("X-C-Version", "4.7.0")`
`header`字段`X-C-Version` 是否 `等于` `4.7.0`
- `header("X-C-Version", "start#4.7.0")`
`header`字段`X-C-Version` 是否 `startsWith` `4.7.0`
- `header("X-C-Version", "end#4.7.0")`
`header`字段`X-C-Version` 是否 `endsWith` `4.7.0`
- `header("X-C-Version", "contains#4.7.0")`
`header`字段`X-C-Version` 是否 `contains` `4.7.0`

### cookie
语法 `cookie("nihao", "234")`

### host
语法 `host("suffix", "-sit")`  
`suffix`的意思是，对于一个域名 `a-b-c.hello.com` 来说，第一个点号`.` 之前的字符串 `a-b-c` 是否以  `-sit` 结尾  
暂时只支持`suffix` `prefix`

## 待办
1. 规则匹配优化， cookie 取出来缓存一下  
2. 规则检测的时候，多个规则，一个合法，其它不合法，目前不会报错  
