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

// func (h *SensorWeaponCoverageController) GetSWCDatabyID(ctx context.Context, request *swc.GetSWCParameterByIdrequest) (*swc.GetSWCParameterByIdRes, error) {
// 	swcId := &swc.GetSWCParameterByIdrequest{Id: int32(request.Id)}

// 	swcData, err := h.swcService.GetSwcDatabyId(ctx, swcId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	swcResponses := &swc.GetSWCParameterByIdRes{
// 		Id:          swcData.ID,
// 		Type:        swcData.Type,
// 		GroupSWC:    swcData.GroupSWC,
// 		Item:        swcData.Item,
// 		Environment: swcData.Environment,
// 		Value:       swcData.Value,
// 		Default:     swcData.Value,
// 		Unit:        swcData.Unit,
// 		UpdatedAt:   swcData.UpdatedAt,
// 		CreatedAt:   swcData.CreatedAt,
// 	}

// 	return swcResponses, nil

// }

// func (h *SensorWeaponCoverageController) GetSWCList(ctx context.Context, request *swc.GetSWCListrequest) (*swc.GetSWCListRes, error) {
// 	swcListrequestuest := swc.GetSWCListrequest{
// 		PageNumber:  request.PageNumber,
// 		PageSize:    request.PageSize,
// 		Environment: request.Environment,
// 		Group:       request.Group,
// 	}

// 	swcList, err := h.swcService.GetSwcList(ctx, &swcListrequestuest)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return swcList, err

// }
