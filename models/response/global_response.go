package response

type MessageSucces struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Meta struct {
	Code    uint16 `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Data struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Meta Meta `json:"meta"`
}
