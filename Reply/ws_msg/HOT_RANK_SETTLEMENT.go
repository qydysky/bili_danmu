package part

type HOT_RANK_SETTLEMENT struct{
    Cmd string `json:"cmd"`
    Data struct{
        Rank int `json:"rank"`
        Uname string `json:"uname"`
        Face string `json:"face"`
        Timestamp int `json:"timestamp"`
        Icon string `json:"icon"`
        Area_name string `json:"area_name"`
        Url string `json:"url"`
        Cache_key string `json:"cache_key"`
        Dm_msg string `json:"dm_msg"`
    } `json:"data"`
}
