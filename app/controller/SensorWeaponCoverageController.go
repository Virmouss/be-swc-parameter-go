package controller

import (
	"be-swc-parameter/app/generated/swc"
	"be-swc-parameter/app/router"
	"context"

	"google.golang.org/grpc"
)

type SensorWeaponCoverageController struct {
	swcService router.SwcService
	swc.UnimplementedSWCParametersServer
}

func NewGrpcSWCService(grpc *grpc.Server, swcService router.SwcService) {

	gRPCHandler := &SensorWeaponCoverageController{
		swcService: swcService,
	}

	swc.RegisterSWCParametersServer(grpc, gRPCHandler)
}

func (h *SensorWeaponCoverageController) AddSWCParameter(ctx context.Context, request *swc.AddSWCParameterReq) (*swc.AddSWCParameterRes, error) {

	swcData := &swc.AddSWCParameterReq{
		Type:        request.Type,
		Group:       request.Group,
		Item:        request.Item,
		Environment: request.Environment,
		Value:       request.Value,
		DefaultSwc:  request.DefaultSwc,
		Unit:        request.Unit,
	}

	service_response, err := h.swcService.AddSWCParameter(ctx, swcData)
	if err != nil {
		return service_response, err
	}

	return service_response, nil

}

func (h *SensorWeaponCoverageController) UpdateSWCParameter(ctx context.Context, request *swc.UpdateSWCParameterReq) (*swc.UpdateSWCParameterRes, error) {

	swcData := &swc.UpdateSWCParameterReq{
		Id:          request.Id,
		Type:        request.Type,
		Group:       request.Group,
		Item:        request.Item,
		Environment: request.Environment,
		Value:       request.Value,
		Unit:        request.Unit,
	}

	service_response, err := h.swcService.UpdateSWCParameter(ctx, swcData)
	if err != nil {
		return service_response, err
	}
	return service_response, nil

}

func (h *SensorWeaponCoverageController) DeleteSWCParameter(ctx context.Context, request *swc.DeleteSWCParameterReq) (*swc.DeleteSWCParameterRes, error) {

	swcData := &swc.DeleteSWCParameterReq{
		Id: request.Id,
	}

	service_response, err := h.swcService.DeleteSWCParameter(ctx, swcData)

	if err != nil {
		return service_response, err
	}
	return service_response, nil

}

func (h *SensorWeaponCoverageController) DeleteAllSWCParameter(ctx context.Context, request *swc.EmptyRequest) (*swc.DeleteAllSWCParameterRes, error) {

	swcData := &swc.EmptyRequest{}

	service_response, err := h.swcService.DeleteAllSWCParameter(ctx, swcData)
	if err != nil {
		return service_response, err
	}
	return service_response, nil

}

func (h *SensorWeaponCoverageController) GetAllSWCParameter(ctx context.Context, request *swc.GetAllSWCParameterReq) (*swc.GetAllSWCParameterRes, error) {

	swcData := &swc.GetAllSWCParameterReq{
		PageNumber: request.PageNumber,
		PageSize:   request.PageSize,
		Group:      request.Group,
		Field:      request.Field,
		Ascending:  request.Ascending,
	}

	service_response, err := h.swcService.GetAllSWCParameter(ctx, swcData)
	if err != nil {
		return nil, err
	}
	return service_response, nil

}
