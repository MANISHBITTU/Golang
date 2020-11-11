package main  
import "fmt"  
func main() { 
   var n int=7
   fmt.Println("Enter any integer")  
   fmt.Scanln(&n)
   fmt.Println("Factorial of ",n," is ",fact(n))  
}  
func fact(n int) int{
	if(n==0||n==1){
         return 1;
	} else{
		return n*fact(n-1)
	}
}