# matrix-search-research
Research project aimed at improving average search time in a row and column wise sorted matrix

A 2D matrix of this type has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.

At present, the best average time was reached using the following algorithm:

``` go
func search004(table [][]int, target int) (int,int) {
    i,j := 0,len(table[0])-1
    // binary-search to find the best starting position
    if table[0][j] < target {
        i1 := len(table)
        for i < i1 {
            m := (i+i1)>>1
            if table[m][j] < target {
                i = m+1
            } else {
                i1 = m
            }
        }
    } else {
        j0 := 0
        for j0 < j {
            m := (j0+j)>>1
            if table[0][m] < target {
                j0 = m+1
            } else {
                j = m
            }
        }
    }
    // main search
    for i < len(table) && j >= 0 {
        if table[i][j] > target {
            j--
        } else if table[i][j] < target {
            i++
        } else {
            return i,j
        }
    }
    return -1,-1
}
```

It uses binary-search to find the best starting position for the main search.
For most of the test inputs this simple optimization works better than the
other methods present in this research.


