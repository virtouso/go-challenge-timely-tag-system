package request

type SubscribeUserRequest struct {
	UserId    string `json:"user_id"`
	SegmentId string `json:"segment_id"`
}
