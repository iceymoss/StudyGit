package trip

import (
	"context"
	"fmt"
	proto "grpc_gateway/server/trip/api"
)

type TripService struct{}

func (t *TripService) GetTrip(ctc context.Context, req *proto.GetTripRequest) (*proto.GetTripResponse, error) {
	fmt.Println(req.Id)

	return &proto.GetTripResponse{
		TripId: 2022,
		UserId: 520,
		Start:  "北京",
		End:    "上海",
		Path: []*proto.PathData{
			{
				Path: "济南",
			}, {
				Path: "郑州",
			}, {
				Path: "南京",
			},
		},
		Distans: 1213,
		Fee:     1400,
		Time:    10200,
	}, nil
}

func (t *TripService) CreateTrip(ctx context.Context, req *proto.CreateTripRequest) (*proto.CreateTripRespnse, error) {
	fmt.Println(req)
	return &proto.CreateTripRespnse{
		TripId:  00001,
		UserId:  req.UserId,
		Start:   req.Start,
		End:     req.End,
		Distans: 1213,
		Fee:     1400,
		Time:    10200,
	}, nil

}
