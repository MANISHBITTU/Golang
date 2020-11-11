package main  
import "fmt"  
func main() { 
   var n int
   fmt.Println("Enter any integer")  
   fmt.Scanln(&n)
   old_n := n
   var size int=0
   for n>0 {
		 size++
		 n=n/10
   }
   var a,p int=0,1
   n=old_n
   var temp int=0
   for n>0 {
	  a=n%10
	  p=1
       for i:=1; i<=size; i++ {
           p=p*a
       }
      temp=temp+p
      n=n/10
  }
  fmt.Println(temp)
  if(old_n==temp){
	fmt.Println(old_n,"is Armstrong number")  
   }else{
	fmt.Println(old_n,"is not Armstrong number") 
   }   
}