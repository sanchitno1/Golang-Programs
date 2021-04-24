package main
import "fmt"
func main(){
	fmt.Printf("Welcome to Golang Calculator\n")
	calc()
}
func calc(){
	var operator string;
	var op1,op2 int;
	fmt.Printf("Enter Two Numbers:   ")
	fmt.Scanf("%d %d",&op1,&op2)
	fmt.Printf("Enter the operator(+,-,*,/)...For remainder press R: ")
	fmt.Scan(&operator)
	switch operator{
		case "+":
			fmt.Printf("Sum is:  %d\n",op1+op2)
			break
		case "-":
			fmt.Printf("Difference is:  %d\n",op1-op2)
			break
		case "*":
			fmt.Printf("Product is:  %d\n",op1*op2)
			break
		case "/":
			fmt.Printf("Quotient is:  %d\n",op1/op2)
			break
		case "R":
			fmt.Printf("Remainder is:  %d\n",op1%op2)
			break
		default:
			fmt.Printf("Please Enter Correct Operator!!!")
		}
}
		
		
