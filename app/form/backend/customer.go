package backend

type CustomerForm struct {
	PageSize     int
	Page         int
	Id           int
	Name         string
	AssignTo     string
	AssignStatus string
	Source       []int
	FollowStatus string
	Director   string
}
