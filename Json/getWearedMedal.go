package part

type GetWearedMedal struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type GetWearedMedal_Roominfo struct {
	Title            string `json:"title"`
	RoomID           int    `json:"room_id"`
	UID              int    `json:"uid"`
	Online           int    `json:"online"`
	LiveTime         int    `json:"live_time"`
	LiveStatus       int    `json:"live_status"`
	ShortID          int    `json:"short_id"`
	Area             int    `json:"area"`
	AreaName         string `json:"area_name"`
	AreaV2ID         int    `json:"area_v2_id"`
	AreaV2Name       string `json:"area_v2_name"`
	AreaV2ParentName string `json:"area_v2_parent_name"`
	AreaV2ParentID   int    `json:"area_v2_parent_id"`
	Uname            string `json:"uname"`
	Face             string `json:"face"`
	TagName          string `json:"tag_name"`
	Tags             string `json:"tags"`
	CoverFromUser    string `json:"cover_from_user"`
	Keyframe         string `json:"keyframe"`
	LockTill         string `json:"lock_till"`
	HiddenTill       string `json:"hidden_till"`
	BroadcastType    int    `json:"broadcast_type"`
}
type GetWearedMedal_Data struct {
	GuardType        int                     `json:"guard_type"`
	Intimacy         int                     `json:"intimacy"`
	IsReceive        int                     `json:"is_receive"`
	LastWearTime     int                     `json:"last_wear_time"`
	Level            int                     `json:"level"`
	LplStatus        int                     `json:"lpl_status"`
	MasterAvailable  int                     `json:"master_available"`
	MasterStatus     int                     `json:"master_status"`
	MedalID          int                     `json:"medal_id"`
	MedalName        string                  `json:"medal_name"`
	ReceiveChannel   int                     `json:"receive_channel"`
	ReceiveTime      string                  `json:"receive_time"`
	Score            int                     `json:"score"`
	Source           int                     `json:"source"`
	Status           int                     `json:"status"`
	TargetID         int                     `json:"target_id"`
	TodayIntimacy    int                     `json:"today_intimacy"`
	UID              int                     `json:"uid"`
	TargetName       string                  `json:"target_name"`
	TargetFace       string                  `json:"target_face"`
	LiveStreamStatus int                     `json:"live_stream_status"`
	IconCode         int                     `json:"icon_code"`
	IconText         string                  `json:"icon_text"`
	Rank             string                  `json:"rank"`
	MedalColor       int                     `json:"medal_color"`
	MedalColorStart  int                     `json:"medal_color_start"`
	MedalColorEnd    int                     `json:"medal_color_end"`
	GuardLevel       int                     `json:"guard_level"`
	MedalColorBorder int                     `json:"medal_color_border"`
	IsLighted        int                     `json:"is_lighted"`
	TodayFeed        int                     `json:"today_feed"`
	DayLimit         int                     `json:"day_limit"`
	NextIntimacy     int                     `json:"next_intimacy"`
	CanDelete        bool                    `json:"can_delete"`
	IsUnion          int                     `json:"is_union"`
	Roominfo         GetWearedMedal_Roominfo `json:"roominfo"`
}
