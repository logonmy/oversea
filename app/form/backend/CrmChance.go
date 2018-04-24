package backend

type CrmChanceForm struct {
	PageSize int
	Page     int
	CustId   int
}

func New() *CrmChanceForm {
	return &CrmChanceForm{
		Page:     1,
		PageSize: 10,
	}
}
