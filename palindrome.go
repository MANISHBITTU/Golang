package main  
import "fmt"  
func main() { 
   var n int
   fmt.Println("Enter any integer")  
   fmt.Scanln(&n)
   old_n := n
   var a int=0;
   for n>0 {
		 a=a*10+n%10
		 n=n/10
   }
   if(old_n==a){
	fmt.Println("",old_n,"is Palindrome")  
   }else{
	fmt.Println(old_n,"is not Palindrome") 
   }
}  