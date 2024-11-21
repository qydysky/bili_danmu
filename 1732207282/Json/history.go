package part

type History struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Cursor struct {
			Max      int    `json:"max"`
			ViewAt   int    `json:"view_at"`
			Business string `json:"business"`
			Ps       int    `json:"ps"`
		} `json:"cursor"`
		Tab []struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"tab"`
		List []struct {
			Title     string      `json:"title"`
			LongTitle string      `json:"long_title"`
			Cover     string      `json:"cover"`
			Covers    interface{} `json:"covers"`
			URI       string      `json:"uri"`
			History   struct {
				Oid      int    `json:"oid"`
				Epid     int    `json:"epid"`
				Bvid     string `json:"bvid"`
				Page     int    `json:"page"`
				Cid      int    `json:"cid"`
				Part     string `json:"part"`
				Business string `json:"business"`
				Dt       int    `json:"dt"`
			} `json:"history"`
			Videos     int    `json:"videos"`
			AuthorName string `json:"author_name"`
			AuthorFace string `json:"author_face"`
			AuthorMid  int    `json:"author_mid"`
			ViewAt     int    `json:"view_at"`
			Progress   int    `json:"progress"`
			Badge      string `json:"badge"`
			ShowTitle  string `json:"show_title"`
			Duration   int    `json:"duration"`
			Current    string `json:"current"`
			Total      int    `json:"total"`
			NewDesc    string `json:"new_desc"`
			IsFinish   int    `json:"is_finish"`
			IsFav      int    `json:"is_fav"`
			Kid        int    `json:"kid"`
			TagName    string `json:"tag_name"`
			LiveStatus int    `json:"live_status"`
		} `json:"list"`
	} `json:"data"`
}
