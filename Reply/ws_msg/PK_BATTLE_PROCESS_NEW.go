package part

type PK_BATTLE_PROCESS_NEW struct {
	Cmd       string `json:"cmd"`
	PkID      int    `json:"pk_id"`
	PkStatus  int    `json:"pk_status"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType int `json:"battle_type"`
		InitInfo   struct {
			RoomID    int    `json:"room_id"`
			Votes     int    `json:"votes"`
			BestUname string `json:"best_uname"`
		} `json:"init_info"`
		MatchInfo struct {
			RoomID    int    `json:"room_id"`
			Votes     int    `json:"votes"`
			BestUname string `json:"best_uname"`
		} `json:"match_info"`
	} `json:"data"`
}