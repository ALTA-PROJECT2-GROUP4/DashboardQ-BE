package handler

import "dashboardq-be/features/class"

type CreateReq struct {
	Name       string `json:"name" form:"name"`
	StartClass string `json:"start_class" form:"start_class"`
	EndClass   string `json:"end_class" form:"end_class"`
	// IdUser     uint   `json:"id_user" form:"id_user"`
}

func CrtToCore(data interface{}) *class.Core {
	res := class.Core{}

	switch data.(type) {
	case CreateReq:
		cln := data.(CreateReq)
		res.Name = cln.Name
		res.StartClass = cln.StartClass
		res.EndClass = cln.EndClass
		// res.IdUser = cln.IdUser
	default:
		return nil
	}

	return &res
}
