package test

import "app"

func InitContext() *app.GlobalContext {
    ctx := app.GlobalContext{}
    ctx.Initialize()
    return &ctx
}

func InitDummyContext() *app.GlobalContext {
    ctx := app.GlobalContext{}
    ctx.InitializeDummy()
    return &ctx
}
