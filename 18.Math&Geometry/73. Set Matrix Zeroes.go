/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2022-12-10 16:45:37  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2022-12-10 16:45:37  */

package main

// import "fmt"



func main(){
	li := [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(li)



}


func setZeroes(matrix [][]int)  {
    row1:=1

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if matrix[i][j]==0 && i==0{
				row1=0
				// matrix[0][j]=0
			}else if matrix[i][j]==0{
				matrix[i][0]=0
				matrix[0][j]=0
			}
		}
	}
	// fmt.Println(matrix)
	
	for i:=1;i<len(matrix);i++{
		for j:=1;j<len(matrix[i]);j++{
			
			 if matrix[i][0]==0 || matrix[0][j]==0{
				matrix[i][j]=0
			}
		}
	}

	if matrix[0][0]==0{
		for i:=0;i<len(matrix);i++{
			matrix[i][0]=0
		}
	}
	if row1==0{
		for i:=0;i<len(matrix[0]);i++{
			matrix[0][i]=0
		}
	}
	// fmt.Println(matrix)

}