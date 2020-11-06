package main

import (  
   "fmt"  
   "time"  
)

func main(){    
   go f1()  
   go f2()  
   time.Sleep(time.Second)
}
  
func f1(){  
   for i:=0; i<10; i++ {  
      fmt.Println("f1 - ", i)  
      time.Sleep(100)  
   }  
}
  
func f2(){  
   for i:=0; i<10; i++ {  
      fmt.Println("f2 - ", i)  
      time.Sleep(100)
   }  
}