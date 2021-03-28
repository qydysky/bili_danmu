package part

type GetGuardNum struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Info struct {
			Num                     int `json:"num"`
			Page                    int `json:"page"`
			Now                     int `json:"now"`
			AchievementLevel        int `json:"achievement_level"`
			AnchorGuardAchieveLevel int `json:"anchor_guard_achieve_level"`
		} `json:"info"`
		List []struct {
			UID           int    `json:"uid"`
			Ruid          int    `json:"ruid"`
			Rank          int    `json:"rank"`
			Username      string `json:"username"`
			Face          string `json:"face"`
			IsAlive       int    `json:"is_alive"`
			GuardLevel    int    `json:"guard_level"`
			GuardSubLevel int    `json:"guard_sub_level"`
			MedalInfo     struct {
				MedalName        string `json:"medal_name"`
				MedalLevel       int    `json:"medal_level"`
				MedalColorStart  int    `json:"medal_color_start"`
				MedalColorEnd    int    `json:"medal_color_end"`
				MedalColorBorder int    `json:"medal_color_border"`
			} `json:"medal_info"`
		} `json:"list"`
		Top3 []struct {
			UID           int    `json:"uid"`
			Ruid          int    `json:"ruid"`
			Rank          int    `json:"rank"`
			Username      string `json:"username"`
			Face          string `json:"face"`
			IsAlive       int    `json:"is_alive"`
			GuardLevel    int    `json:"guard_level"`
			GuardSubLevel int    `json:"guard_sub_level"`
			MedalInfo     struct {
				MedalName        string `json:"medal_name"`
				MedalLevel       int    `json:"medal_level"`
				MedalColorStart  int    `json:"medal_color_start"`
				MedalColorEnd    int    `json:"medal_color_end"`
				MedalColorBorder int    `json:"medal_color_border"`
			} `json:"medal_info"`
		} `json:"top3"`
		MyFollowInfo struct {
			GuardLevel    int    `json:"guard_level"`
			AccompanyDays int    `json:"accompany_days"`
			ExpiredTime   string `json:"expired_time"`
			AutoRenew     int    `json:"auto_renew"`
			RenewRemind   struct {
				Content string `json:"content"`
				Type    int    `json:"type"`
				Hint    string `json:"hint"`
			} `json:"renew_remind"`
			MedalInfo struct {
				MedalName        string `json:"medal_name"`
				MedalLevel       int    `json:"medal_level"`
				MedalColorStart  int    `json:"medal_color_start"`
				MedalColorEnd    int    `json:"medal_color_end"`
				MedalColorBorder int    `json:"medal_color_border"`
			} `json:"medal_info"`
			Rank int    `json:"rank"`
			Ruid int    `json:"ruid"`
			Face string `json:"face"`
		} `json:"my_follow_info"`
		GuardWarn struct {
			IsWarn      int    `json:"is_warn"`
			Warn        string `json:"warn"`
			Expired     int    `json:"expired"`
			WillExpired int    `json:"will_expired"`
			Address     string `json:"address"`
		} `json:"guard_warn"`
	} `json:"data"`
}
