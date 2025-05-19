package main
import "fmt"

func main(){
	// age:= 19

	// if age>=18{
	// 	fmt.Println("person is an adult")
	// } else if age>=12{
	// 	fmt.Println("Person is Teenager")
	// } else{
	// 	fmt.Println("Person is kid")
	// }


	var role = "admin"
	var hasPermission = true
	var gender = "Male"

	if role == "admin" && hasPermission{
		fmt.Println("Can Access")
	} else {
		fmt.Println("Cannot Access the data")
	}


	if(role == "admin"){
		if(gender=="Female"){
			fmt.Println("Aurat admin h")
		}else{
			fmt.Println("Mard admin h")
		}
	}
}	