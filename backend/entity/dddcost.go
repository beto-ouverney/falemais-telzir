package entity

type DDDCost struct {
	ID          int     `json:"id"`
	Origin      int     `json:"origin"`
	Destination int     `json:"destination"`
	Cost        float64 `json:"cost,omitempty"`
}
