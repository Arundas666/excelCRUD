package models

type Record struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Company   string `json:"company"`
    Address   string `json:"address"`
    City      string `json:"city"`
    County    string `json:"county"`
    Postal    string `json:"postal"`
    Phone     string `json:"phone"`
    Email     string `json:"email"`
    Web       string `json:"web"`
}