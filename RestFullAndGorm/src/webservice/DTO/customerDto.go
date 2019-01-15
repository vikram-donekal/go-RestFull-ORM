package DTO

import "../../../bin/gorm"

type Customer struct {
	gorm.Model
	FirstName string `json:"FirstName" ;sql: "type:VARCHAR(20);not null;unique"`
	LastName string `json:"LastName" ;sql: "type:VARCHAR(20);not null;unique`

}


