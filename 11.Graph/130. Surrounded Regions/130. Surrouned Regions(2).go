/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2024-03-05 12:19:19  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2024-03-05 12:19:19  */

package main


func solve(board [][]byte)  {
  n := len(board)
  m := len(board[0])
  
  var dfs func(i,j int)

  dfs = func(i,j int) {
    if i<0 || j<0 || i>=n || j>= m || board[i][j]=='X' || board[i][j]=='*' {
      return 
    }
    board[i][j] = '*'
    dfs(i-1,j)
    dfs(i+1,j)
    dfs(i,j-1)
    dfs(i,j+1)
  }

  for i:=0 ; i<n ;i++ {
    for j:=0 ;j<m ;j++ {
      if (i==0 || j==0 || i==n-1 || j==m-1) && board[i][j] == 'O' {
        dfs(i,j)
      }
    }
  }

  for i:=0 ; i<n;i++ {
    for j:=0; j<m;j++ {
      if board[i][j]== '*'{
        board[i][j] = 'O'
      }else if board[i][j] == 'O' {
        board[i][j] = 'X'
      }
    }
  }

}
