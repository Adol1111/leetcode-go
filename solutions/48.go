package solutions

func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// 1 2 3
// 4 5 6
// 7 8 9

// i=0,j=0
// 0(i),0(j) -> 2(n-j-1),0                7, 2, 1
// 2(n-j-1),0 -> 2(n-i-1),2(n-j-1)        4, 5, 6
// 2(n-i-1),2(n-j-1) -> 0(j),2(n-i-1)     9, 8, 3
// 0(j),2(n-i-1) -> 0(i),0(j)

// i=0,j=1
// 0(i),1(j) -> 1(n-j-1),0(i)             7, 4, 1
// 1(n-j-1),0(i) -> 2(n-i-1),1(n-j-1)     8, 5, 2
// 2(n-i-1),1(3-j-1) -> 1(j),2(n-i-1)     9, 6, 3
// 1(j),2(n-i-1) -> 0,1

// 1  2  3  4
// 5  6  7  8
// 9  10 11 12
// 13 14 15 16

// i=0,j=0
// 0(i),0(j) -> 3(n-j-1),0               13 2  3  1
// 3(n-j-1),0 -> 3(n-i-1),3(n-j-1)       5  6  7  8
// 3(n-i-1),3(n-j-1) -> 0(j),3(n-i-1)    9  10 11 12
// 0(j),3(n-i-1) -> 0(i),0(j)            16 14 15 4

// i=0,j=1
// 0(i),1(j) -> 2(n-j-1),0(i)            13 9  3  1
// 2(n-j-1),0(i) -> 3(n-i-1),2(n-j-1)    5  6  7  2
// 3(n-i-1),2(n-j-1) -> 1(j),3(n-i-1)    15 10 11 12
// 1(j),3(n-i-1) -> 0,1                  16 14 8  4

// i=1,j=0
// 1(i),0(j) -> 3(n-j-1),1(i)            13 9  5  1
// 3(n-j-1),1(i) -> 2(n-i-1),3(n-j-1)    14 6  7  2
// 2(n-i-1),3(n-j-1) -> 0(j),2(n-i-1)    15 10 11 3
// 0(j),2(n-i-1) -> 1,0                  16 12 8  4

// i=1,j=1
// 1(i),1(j) -> 2(n-j-1),1(i)            13 9  5  1
// 2(n-j-1),1(i) -> 2(n-i-1),2(n-j-1)    14 10 6  2
// 2(n-i-1),2(n-j-1) -> 1(j),2(n-i-1)    15 11 7  3
// 1(j),2(n-i-1) -> 1,1                  16 12 8  4
