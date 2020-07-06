func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0]==1 {
        return 0
    }
    obstacleGrid[0][0] = 1
    for j:=1;j<len(obstacleGrid[0]);j++ {
        if obstacleGrid[0][j] == 1 {
            obstacleGrid[0][j] = 0
        } else if obstacleGrid[0][j-1] == 1 {
             obstacleGrid[0][j] = 1
        }
    }
    for i:=1;i<len(obstacleGrid);i++ {
        for j:=0;j<len(obstacleGrid[0]);j++ {
            if obstacleGrid[i][j] == 1 {
                obstacleGrid[i][j] = 0
            } else {
                if j ==0 {
                    obstacleGrid[i][j] =  obstacleGrid[i-1][j]
                } else {
                    obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
                }
            }
        }
    }
    return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
