package main

import "fmt"

type Matrix struct {
	rows int
	col int
	values [][]int
}

func (t Matrix) getrows() int{
	return t.rows
}

func (t Matrix) getcolumns() int{
	return t.col
}

func (t *Matrix) setelement(i,j,val int) string{
	if i >= t.rows || j >= t.col {
		return "Failed"
	}
	t.values[i][j] = val
	return "Success"
}

func (t Matrix) add(m [][]int) [][]int{
	for i:=0; i<t.rows; i++{
		for j:=0; j<t.col; j++{
			t.values[i][j] += m[i][j]
		}
	}
	return	t.values
}


func main(){
	var a Matrix
	a.rows = 2
	a.col = 2
	a.values = [][]int{{1,2},{3,9}}
	fmt.Println(a.setelement(1,2,55))
	fmt.Println(a)
	fmt.Println(a.add([][]int{{1,2},{2,1}}))


}

