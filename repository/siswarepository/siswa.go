package siswarepository

import (
	"project-sia/models/request"
	"project-sia/models/response"
)

type SiswaRepository interface {
	GetSiswaByEmail(rq request.GetByData, data *response.DataSiswaWLog) *response.DataSiswaWLog
}
