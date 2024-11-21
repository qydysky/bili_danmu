package part

type StreamType struct {
	ProtocolName string `json:"protocol_name"`
	Format       []struct {
		FormatName string `json:"format_name"`
		Codec      []struct {
			CodecName string `json:"codec_name"`
			CurrentQn int    `json:"current_qn"`
			AcceptQn  []int  `json:"accept_qn"`
			BaseURL   string `json:"base_url"`
			URLInfo   []struct {
				Host      string `json:"host"`
				Extra     string `json:"extra"`
				StreamTTL int    `json:"stream_ttl"`
			} `json:"url_info"`
			HdrQn     interface{} `json:"hdr_qn"`
			DolbyType int         `json:"dolby_type"`
			AttrName  string      `json:"attr_name"`
		} `json:"codec"`
	} `json:"format"`
}
