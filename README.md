# goweb-cors
Cross-origin_resource_sharing util for golang

What is [CORS](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing) ?

CORS请求保护不当可导致敏感信息泄漏，因此应当严格设置Access-Control-Allow-Origin使用同源策略进行保护。

```
c := cors.New(cors.Options{
	AllowedOrigins:   []string{"http://qq.com", "https://qq.com"},
	AllowCredentials: true,
	Debug:            false,
})

// 引入中间件
handler = c.Handler(handler)
```
