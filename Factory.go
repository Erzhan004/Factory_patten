package main
import "fmt"

type Cars interface {
	SetName(name string)
	SetPrice(price int)
	SetPower(power int)
	SetCategory(category string)
	GetName() string
	GetPrice() int
	GetPower() int
	GetCategory() string
}

type Car struct{
	name string
	price int
	power int
	category string
}

func(c *Car) SetName(name string){
   c.name = name
}
func(c *Car) SetPrice(price int){
	c.price = price
 }
 func(c *Car) SetPower(power int){
	c.power = power
 }
 func(c *Car) SetCategory(category string){
	c.category = category
 }
 func(c *Car) GetName() string{
	return c.name
 }
 func(c *Car) GetPrice() int{
	return c.price
 } 
 func(c *Car) GetPower() int{
	return c.power
 } 
 func(c *Car) GetCategory() string{
	return c.category
 }

 type SedanFactory struct{}

 func (sf SedanFactory) CreateCar() Cars {
	return &Car{category: "Седан"}
}

type SUVFactory struct{}

func (sf SUVFactory) CreateCar() Cars {
	return &Car{category: "Внедорожник"}
}

func printDetails(cs Cars) {
    fmt.Printf("Name: %s", cs.GetName())
    fmt.Println()
    fmt.Printf("Power: %d", cs.GetPower())
    fmt.Println()
	fmt.Printf("Category: %v", cs.GetCategory())
    fmt.Println()
	fmt.Printf("Price: %d", cs.GetPrice())
    fmt.Println()
	fmt.Println()
	
}

func main(){
	sedanFactory := SedanFactory{}
	suvFactory := SUVFactory{}

	// Создаем седан и внедорожник с помощью фабрик
	sedan := sedanFactory.CreateCar()
	suv := suvFactory.CreateCar()

	// Устанавливаем характеристики машин
	sedan.SetName("Toyota Camry")
	sedan.SetPrice(30000)
	sedan.SetPower(200)
	suv.SetName("Jeep ")
	suv.SetPrice(40000)
	suv.SetPower(250)
	printDetails(sedan)
	printDetails(suv)


}


