package entity

// Plan is the entity for the plan
type Plan struct {
	Name    string  `json:"name"`
	With    float64 `json:"with"`
	Without float64 `json:"without"`
}

// PlanComparation is the entity for the plan comparation
type PlanComparation struct {
	Origin      int    `json:"origin"`
	Destination int    `json:"destination"`
	Minutes     int    `json:"minutes"`
	Plans       []Plan `json:"plan"`
}
