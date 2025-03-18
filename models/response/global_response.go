package response

type MessageSucces struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Error struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Code    uint16 `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Data struct {
	Data interface{} `json:"data"`
}
