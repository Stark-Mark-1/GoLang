package main
import(
	"fmt"
	"time"
)
func main(){
	a := 20
	switch a{
		case 10:
			fmt.Println("Value of a is 10")
		case 20:
			fmt.Println("Value of a is 20")
		case 30:
			fmt.Println("Value of a is 30")
	}
	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
		}
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}
	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("Its a bool")
		case int:
			fmt.Println("Its an int")
		default:
			fmt.Printf("Dont know the type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}