# Baa 中间件

baa 支持通过 中间件 机制，注入请求过程，实现类似插件的功能

## 常规路由

func (b *Baa) Use(m ...Middleware)
