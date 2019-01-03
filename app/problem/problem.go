package problem


type Problem struct {
  Status int    `json: "status"`
  Title  string `json: "title"`
  Detail string `json: "detail"`
  Type   string `json: "type"`
}
