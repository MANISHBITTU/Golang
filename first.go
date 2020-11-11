package main  
import "fmt"  
func main() { 
   var n int
   fmt.Println("Enter any integer")  
   fmt.Scanln(&n)
   var sum int=0
   for a := 1; a <=n; a++ {  
      sum += a 
   }  
   fmt.Println("total sum upto n is ",sum)
}  