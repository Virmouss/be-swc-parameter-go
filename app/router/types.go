package router

import (
	"context"

	"be-swc-parameter/app/generated/swc"
)

type SwcService interface {
	AddSWCParameter(context.Context, *swc.AddSWCParameterReq) (*swc.AddSWCParameterRes, error)
	UpdateSWCParameter(context.Context, *swc.UpdateSWCParameterReq) (*swc.UpdateSWCParameterRes, error)
	DeleteSWCParameter(context.Context, *swc.DeleteSWCParameterReq) (*swc.DeleteSWCParameterRes, error)
	DeleteAllSWCParameter(context.Context, *swc.EmptyRequest) (*swc.DeleteAllSWCParameterRes, error)
	GetAllSWCParameter(context.Context, *swc.GetAllSWCParameterReq) (*swc.GetAllSWCParameterRes, error)
}
