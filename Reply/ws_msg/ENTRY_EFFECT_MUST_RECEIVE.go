package part

type ENTRY_EFFECT_MUST_RECEIVE struct {
	Cmd string `json:"cmd"`
	Data struct {
		Id int `json:"id"`
		Uid int `json:"uid"`
		Target_id int `json:"target_id"`
		Mock_effect int `json:"mock_effect"`
		Face string `json:"face"`
        Privilege_type int `json:"privilege_type"`
        Copy_writing string `json:"copy_writing"`
        Copy_color string `json:"copy_color"`
        Highlight_color string `json:"highlight_color"`
        Priority int `json:"priority"`
        Basemap_url string `json:"basemap_url"`
        Show_avatar int `json:"show_avatar"`
        Effective_time int `json:"effective_time"`
        Web_basemap_url string `json:"web_basemap_url"`
        Web_effective_time int `json:"web_effective_time"`
        Web_effect_close int `json:"web_effect_close"`
        Web_close_time int `json:"web_close_time"`
        Business int `json:"business"`
        Copy_writing_v2 string `json:"copy_writing_v2"`
        Icon_list []int `json:"icon_list"`
        Max_delay_time int `json:"max_delay_time"`
	} `json:"data"`
}
/*
{
    "cmd": "ENTRY_EFFECT_MUST_RECEIVE",
    "data": {
        "id": 136,
        "uid": 29183321,
        "target_id": 612524985,
        "mock_effect": 0,
        "face": "https://i2.hdslb.com/bfs/face/8c6d1ce3f96dc86d0fc7876a2824910c92ae6802.jpg",
        "privilege_type": 0,
        "copy_writing": "欢迎 \u003c%qydysky...%\u003e 进入直播间",
        "copy_color": "#000000",
        "highlight_color": "#FFF100",
        "priority": 1,
        "basemap_url": "https://i0.hdslb.com/bfs/live/mlive/586f12135b6002c522329904cf623d3f13c12d2c.png",
        "show_avatar": 1,
        "effective_time": 2,
        "web_basemap_url": "https://i0.hdslb.com/bfs/live/mlive/586f12135b6002c522329904cf623d3f13c12d2c.png",
        "web_effective_time": 2,
        "web_effect_close": 0,
        "web_close_time": 900,
        "business": 3,
        "copy_writing_v2": "欢迎 \u003c^icon^\u003e \u003c%qydysk…%\u003e 进入直播间",
        "icon_list": [
            2
        ],
        "max_delay_time": 7
    }
}
*/