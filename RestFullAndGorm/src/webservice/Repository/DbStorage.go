package Repository

import (
	"../DTO"
	"../Helper"
)



func GetAllCustomer() []DTO.Customer {

	CustomerListResponse:=[]DTO.Customer{}
	Helper.GetDB().Debug().Find(&CustomerListResponse)
	return CustomerListResponse
}

func GetById(id int32) DTO.Customer {

	CustomerReponse:=DTO.Customer{}
	Helper.GetDB().Debug().First(&CustomerReponse,id)
	return CustomerReponse
}

func SaveCustomer (input DTO.Customer) DTO.Customer{

	checkCustomer:=GetById( int32(input.ID))
	if(checkCustomer.LastName =="" && checkCustomer.FirstName =="" ){
		Helper.GetDB().Debug().Save(&input)

	}else {
		checkCustomer.FirstName=input.FirstName
		checkCustomer.LastName=input.LastName
		Helper.GetDB().Debug().Save(&checkCustomer)
	}

	return input
}

func CreateMessage(FirstNameInput string, LastNameInput string) DTO.Customer {


	return DTO.Customer{

		FirstName:FirstNameInput,
		LastName:LastNameInput,
	}
}


func DeleteCustomer(  id int32) DTO.Customer {

	checkCustomer:=GetById(id)
	if(checkCustomer.LastName !="" && checkCustomer.FirstName !="" ){
		Helper.GetDB().Debug().Delete(&checkCustomer)

	}else {
		return  DTO.Customer{}
	}
	return checkCustomer

}