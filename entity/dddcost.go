package entity

type DDDCost struct {
	Origin      int     `json:"origin"`
	Destination int     `json:"destination"`
	Cost        float64 `json:"cost"`
}
