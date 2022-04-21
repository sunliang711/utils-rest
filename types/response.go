package types

// Resp stands for object to response to client
// 2020/02/19 22:25:52
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type WithdrawResp struct {
	Code       int    `json:"code" structs:"code"`
	Msg        string `json:"msg" structs:"msg,omitempty"`
	AppId      string `json:"appid" structs:"appid,omitempty"`
	Nonce      string `json:"nonce" structs:"nonce,omitempty"`
	OutTradeNo string `json:"out_trade_no" structs:"out_trade_no,omitempty"`
	Attach     string `json:"attach" structs:"attach,omitempty"`
	Sign       string `json:"sign" structs:"-"`
	SignType   string `json:"sign_type" structs:"sign_type,omitempty"`
	ResultCode int    `json:"result_code" structs:"result_code,omitempty"`
	ResultMsg  string `json:"result_msg" structs:"result_msg,omitempty"`
	TxId       string `json:"tx_id" structs:"tx_id,omitempty"`
}

type WithdrawStatusResp struct {
	Code           int    `json:"code" structs:"code"`
	Msg            string `json:"msg" structs:"msg,omitempty"`
	AppId          string `json:"appid" structs:"appid,omitempty"`
	Nonce          string `json:"nonce" structs:"nonce,omitempty"`
	OutTradeNo     string `json:"out_trade_no" structs:"out_trade_no,omitempty"`
	Attach         string `json:"attach" structs:"attach,omitempty"`
	Sign           string `json:"sign" structs:"-"`
	SignType       string `json:"sign_type" structs:"sign_type,omitempty"`
	TxId           string `json:"tx_id" structs:"tx_id,omitempty"`
	TxStatus       string `json:"tx_status" structs:"tx_status,omitempty"`
	TxHash         string `json:"tx_hash" structs:"tx_hash,omitempty"`
	BlockHash      string `json:"block_hash" structs:"block_hash,omitempty"`
	BlockNumber    uint64 `json:"block_number" structs:"block_number,omitempty"`
	BlockTimestamp uint64 `json:"block_timestamp" structs:"block_timestamp,omitempty"`
	Fee            string `json:"fee" structs:"fee,omitempty"`
}
