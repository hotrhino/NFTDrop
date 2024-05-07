package NFTDrop

type MerkleProofResp struct {
	Code int `json:"code"`
	Data struct {
		Sale   int      `json:"sale"`
		Limit  int      `json:"limit"`
		Bought int      `json:"bought"`
		Paths  []string `json:"paths"`
	} `json:"data"`
}
