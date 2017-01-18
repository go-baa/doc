# Baa 依赖注入

依赖注入(dependency injection)简称 DI，是 baa 实现的核心，baa 所有组件基于DI组装起来的。

## 常规路由

func (b *Baa) SetDIer(v DIer)
func (b *Baa) SetDI(name string, h interface{})
func (b *Baa) GetDI(name string) interface{}
func (c *Context) DI(name string) interface{}
