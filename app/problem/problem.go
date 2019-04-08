package problem

/*Problem struct is the service representation
for the problem+json response object.*/
type Problem struct {
  Status int    `json:"status"`
  Title  string `json:"title"`
  Detail string `json:"detail"`
  Type   string `json:"type"`
}
