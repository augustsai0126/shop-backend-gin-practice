package handler

import "shop-backend-gin-practice/internal/service"

type HandlerGroup struct {
	User        *UserHandler
	Category    *CategoryHandler
	Product     *ProductHandler
	Cart        *CartHandler
	UserService service.UserService // for middleware
}
