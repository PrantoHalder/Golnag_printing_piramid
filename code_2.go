package main

import "fmt"

func forloop (){
	fmt.Println("home work - 1")
	var i int
	var j int
	var n int = 5
	for i=1;i<=n;i++ {
		for j=1;j<=2*n-1;j++{
	        if (j>=n-(i-1) && j<= n+(i-1)){
				fmt.Printf("*")
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
		
	
}