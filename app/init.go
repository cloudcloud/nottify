package app

import "github.com/revel/revel"

func init() {
    revel.Filters = []revel.Filter{
        revel.PanicFilter,
        revel.RouterFilter,
        revel.FilterConfiguringFilter,
        revel.ParamsFilter,
        revel.SessionFilter,
        revel.FlashFilter,
        revel.ValidationFilter,
        revel.I18nFilter,
        HeaderFilter,
        revel.InterceptorFilter,
        revel.CompressFilter,
        revel.ActionInvoker,
    }
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
    c.Response.Out.Header().Add("X-Frame-Options", "ORIGIN")
    c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
    c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

    fc[0](c, fc[1:])
}

