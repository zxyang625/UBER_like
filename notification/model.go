package no

import "notification/pkg/grpc/pb"

type NoticeTripRequest struct {
	*pb.NoticeTripRequest `json:"_pb_notice_bill_request,omitempty" protobuf:"bytes,1,opt,name=notice_trip_request,json=tripMsg,proto3"`
}

type NoticeTripReply struct {
	*pb.NoticeTripReply `json:"_pb_notice_trip_reply,omitempty" protobuf:"bytes,1,opt,name=notice_trip_reply,json=tripMsg,proto3"`
}

//type NoticeBillRequest struct {
//	*pb.NoticeBillRequest `json:"_pb_notice_bill_request,omitempty"`
//}
//
//type NoticeBillReply struct {
//	*pb.NoticeBillReply `json:"_pb_notice_bill_reply,omitempty"`
//}
