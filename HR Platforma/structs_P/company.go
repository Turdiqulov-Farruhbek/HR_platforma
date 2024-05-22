package structs_P

type Company struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Workers  int    `json:"workers"`
	
}


type CompanyCreate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Workers  int    `json:"workers"`

}

type CompanyUpdate struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Workers  int    `json:"workers"`
}

type CompanyDeleted struct {
    ID string
}


type Companies struct {
    Companies []Company
    Count      int
}