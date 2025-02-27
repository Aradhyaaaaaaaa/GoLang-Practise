package main
type Employee struct{
	name string
	age integer
	isRemote bool
}
type Address struct{
	street string
	city string
}
func madsin(){
	adress:= Address{
		street: "123 Main street ",
		city: "xyz",

	}
	employee1 :=Employee{
		name : "Alice",
		age : 3,
		isRemote : true,
		Adress: adress,
	}
	fmt.prinln("Address:", employee.street, employee.city)
	fmt.println("Employee Name:", employee1.name)
	fmt.println("Employee Age:", employee1.age)

	job:= struct // define anonymous struct 
	{
		title string
		salary int
	}{
		title: "softeware Engineer", //initializing struct*
	}
	fmt.println("job:" job.title) //acess the feild by dot notation these anonymous feild are helpfull to encapsulate the data 
	empoyeeptr:= &employee1//not copying the value but passing by refrence
	employeeptr.age = 31
	}
	//pointers are need to mutate feilds

	// nested struct
	