package part

type PK_BATTLE_PRE_NEW struct {
	Cmd       string `json:"cmd"`
	PkStatus  int    `json:"pk_status"`
	PkID      int    `json:"pk_id"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType  int         `json:"battle_type"`
		MatchType   int         `json:"match_type"`
		Uname       string      `json:"uname"`
		Face        string      `json:"face"`
		UID         int         `json:"uid"`
		RoomID      int         `json:"room_id"`
		SeasonID    int         `json:"season_id"`
		PreTimer    int         `json:"pre_timer"`
		PkVotesName string      `json:"pk_votes_name"`
		EndWinTask  interface{} `json:"end_win_task"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}