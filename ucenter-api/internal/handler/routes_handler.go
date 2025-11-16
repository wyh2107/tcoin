// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serviceCtx *svc.ServiceContext) {
	register := NewRegisterHandler(serviceCtx)
	registerRouter := r.Group()
	registerRouter.POST("/uc/register/phone", register.Register)
}
