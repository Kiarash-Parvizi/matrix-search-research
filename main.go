package main

import(
    "fmt"
    "strconv"
    "time"
    "math/rand"
)

func search000(table [][]int, target int) (int,int) {
    i,j := 0,len(table[0])-1
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

func search004(table [][]int, target int) (int,int) {
    i,j := 0,len(table[0])-1
    //
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
    //
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

// Not guaranteed to find
// It must be used in combination with another search method
func search003(table [][]int, target, d int) (int,int) {
    l,r := 0, len(table)
    for l < r {
        m := (l+r)>>1
        if table[m][len(table[m])-1] < target {
            l = m+1
        } else {
            r = m
        }
    }
    if r == len(table) {
        return -1,-1
    }
    //
    for i := r; i < len(table); i+=d {
        l, r := 0, len(table[i])
        for l < r {
            m := (l+r)>>1
            if table[i][m] < target {
                l = m+1
            } else {
                r = m
            }
        }
        if r != len(table[i]) && table[i][r] == target {
            return i,r
        }
    }
    return -1,-1
}

// Not guaranteed to find
// It must be used in combination with another search method
func search002(table [][]int, si,sj,n,m int, target int) (int,int) {
    if target > table[si+n-1][sj+m-1] {
        return -1,-1
    }
    if n > m {
        l, r := 0, n/m
        for l < r {
            mid := (l+r)>>1
            if table[si + mid*m+m-1][sj + m-1] < target {
                l = mid+1
            } else {
                r = mid
            }
        }
        if l == n/m {
            return search002(table, si+m*l, sj, n%m, m, target)
        }
        return search002(table, si+m*l, sj, m, m, target)
    } else if m > n {
        l, r := 0, m/n
        for l < r {
            mid := (l+r)>>1
            if table[si + n-1][sj + mid*n+n-1] < target {
                l = mid+1
            } else {
                r = mid
            }
        }
        if l == m/n {
            return search002(table, si, sj+n*l, n, m%n, target)
        }
        return search002(table, si, sj+n*l, n, n, target)
    } else {
        l, r := 0, n
        for l < r {
            mid := (l+r)>>1
            if table[si+mid][sj+mid] < target {
                l = mid+1
            } else {
                r = mid
            }
        }
        if table[si+l][sj+l] == target {
            return si+l, sj+l
        }
        if l == 0 {
            return -1,-1
        }
        x, y := search002(table, si+l, sj, n-l, l, target)
        if x != -1 {
            return x,y
        }
        x, y = search002(table, si, sj+l, l, n-l, target)
        return x,y
    }
}

func search001(table [][]int, si,sj,n,m int, target int) (int,int) {
    if n == 1 {
        C++
        if table[si][sj] == target {
            return si,sj
        }
        return -1,-1
    }
    iMid, jMid := si+n/2-1, sj+m/2-1
    // switch
    if table[iMid][jMid] == target {
        return iMid,jMid
    } else if table[iMid][jMid] < target {
        r0,r1 := search001(table, si+n/2,sj,n/2,m/2, target)
        if r0 != -1 {
            return r0,r1
        }
        r0,r1 = search001(table, si,sj+m/2,n/2,m/2, target)
        if r0 != -1 {
            return r0,r1
        }
        r0,r1 = search001(table, si+n/2,sj+m/2,n/2,m/2, target)
        if r0 != -1 {
            return r0,r1
        }
    } else {
        r0,r1 := search001(table, si+n/2,sj,n/2,m/2, target)
        if r0 != -1 {
            return r0,r1
        }
        r0,r1 = search001(table, si,sj+m/2,n/2,m/2, target)
        if r0 != -1 {
            return r0,r1
        }
        r0,r1 = search001(table, si,sj,n/2,m/2, target)
        if r0 != -1 {
            return r0,r1
        }
    }
    return -1,-1
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func mkTable(_size int) [][]int {
    table := make([][]int, _size)
    for i := 0; i < _size; i++ {
        table[i] = make([]int, _size)
    }
    r := 0
    for i := 0; i < _size; i++ {
        r += (i+7)%11 + 4
        table[i][0] = r
    }
    r = table[0][0]
    for i := 1; i < _size; i++ {
        r += (i+7)%11 + 4
        table[0][i] = r
    }
    for i := 1; i < _size; i++ {
        for j := 1; j < _size; j++ {
            r := (j+7)%59 + 7
            mx0 := max(table[i-1][j],table[i][j-1])
            table[i][j] = max(mx0,(table[i-1][j]+table[i][j-1])/2) + r
        }
    }
    return table
}

func mkTable001(_size int) [][]int {
    table := make([][]int, _size)
    for i := 0; i < _size; i++ {
        table[i] = make([]int, _size)
    }
    table[0][0] = 0
    for i := 1; i < _size; i++ {
        table[i][0] = table[i-1][0] + rand.Intn(100)
    }
    for i := 1; i < _size; i++ {
        table[0][i] = table[0][i-1] + rand.Intn(100)
    }
    for i := 1; i < _size; i++ {
        for j := 1; j < _size; j++ {
            mx0 := max(table[i-1][j],table[i][j-1])
            table[i][j] = mx0 + rand.Intn(100)
        }
    }
    return table
}

func mkTable000(_size int) [][]int {
    table := make([][]int, _size)
    for i := 0; i < _size; i++ {
        table[i] = make([]int, _size)
    }
    for i, r := 0, 0; i < _size; i++ {
        for j := 0; j < _size; j++ {
            table[i][j] = r
            r += 2
        }
        r -= _size*3/2
    }
    return table
}

func zeroPad(a, l int) string {
    s := strconv.Itoa(a)
    if len(s) < l {
        t := []byte{}
        for i := 0; i < l-len(s); i++ {
            t = append(t, ' ')
        }
        return string(t) + s
    }
    return s
}

func printTable(table [][]int) {
    for i := 0; i < len(table); i++ {
        for j := 0; j < len(table[0]); j++ {
            fmt.Print(zeroPad(table[i][j], 7))
        }
        fmt.Println("")
    }
}
func printTable_withD(table [][]int, n, m int) {
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Print(zeroPad(table[i][j], 7))
        }
        fmt.Println("")
    }
}

var(
    C int
)

func main() {
    table := mkTable001(3000)
    //printTable(table)
    // print a sample of the table
    {
        fmt.Println("sample of the table:")
        printTable_withD(table, 10,10)
    }
    // search
    mn, mx := table[0][0], table[len(table)-1][len(table[0])-1]
    fmt.Println("min,max: ", mn, mx)
    // average time
    fmt.Println("search000")
    {
        start := time.Now()
        for i := mn; i <= mx; i++ {
            search000(table, i)
        }
        elapsed := time.Since(start)
        fmt.Printf("Time taken: %s\n", elapsed)
    }
    fmt.Println("search004")
    {
        start := time.Now()
        for i := mn; i <= mx; i++ {
            search004(table, i)
        }
        elapsed := time.Since(start)
        fmt.Printf("Time taken: %s\n", elapsed)
    }
    fmt.Println("search002")
    {
        start := time.Now()
        for i := mn; i <= mx; i++ {
            x,_ := search002(table, 0,0, len(table), len(table[0]), i)
            if x == -1 {
                x,_ = search004(table, i)
            }
        }
        elapsed := time.Since(start)
        fmt.Printf("Time taken: %s\n", elapsed)
    }
    fmt.Println("search003")
    {
        start := time.Now()
        for i := mn; i <= mx; i++ {
            x,_ := search003(table, i,2)
            if x == -1 {
                x,_ = search004(table, i)
            }
        }
        elapsed := time.Since(start)
        fmt.Printf("Time taken: %s\n", elapsed)
    }
    // correctness test
    isEmpt := map[int]struct{}{}
    fmt.Println("\n--------------\n\ncorrectness test:")
    fmt.Println("search000")
    {
        notFoundCnt := 0
        for i := mn; i <= mx; i++ {
            x,y := search000(table, i)
            if x == -1 {
                notFoundCnt++
                isEmpt[i] = struct{}{}
            } else {
                if table[x][y] != i {
                    fmt.Println("ERR")
                }
            }
        }
        fmt.Println("notFoundCnt: ", notFoundCnt)
    }
    fmt.Println("search004")
    {
        notFoundCnt := 0
        for i := mn; i <= mx; i++ {
            x,y := search004(table, i)
            if x == -1 {
                notFoundCnt++
            } else {
                if notFoundCnt == 100 {
                    fmt.Println(i, x,y, table[x][y:y+2])
                }
                if table[x][y] != i {
                    fmt.Println("ERR")
                }
            }
        }
        fmt.Println("notFoundCnt: ", notFoundCnt)
    }
    fmt.Println("search002")
    {
        notFoundCnt, notFoundByFirst, foundByFirst := 0, 0, 0
        for i := mn; i <= mx; i++ {
            x,y := search002(table, 0,0, len(table), len(table[0]), i)
            if x == -1 {
                x,y = search004(table, i)
                if x != -1 {
                    notFoundByFirst++
                }
            } else {
                foundByFirst++
            }
            if x == -1 {
                notFoundCnt++
                //_, found := isEmpt[i]
                //if found {
                //    fmt.Println("ok")
                //} else {
                //    fmt.Println("ERR at i =", i)
                //}
            }
            if x != -1 && table[x][y] != i {
                fmt.Println("ERR")
            }
        }
        fmt.Println("notFoundCnt: ", notFoundCnt)
        fmt.Println("notFoundByFirst: ", notFoundByFirst)
        fmt.Println("foundByFirst: ", foundByFirst)
    }
    fmt.Println("search003")
    {
        notFoundCnt, notFoundByFirst, foundByFirst := 0, 0, 0
        for i := mn; i <= mx; i++ {
            x,y := search003(table, i, 2)
            if x == -1 {
                x,y = search004(table, i)
                if x != -1 {
                    notFoundByFirst++
                }
            } else {
                foundByFirst++
            }
            if x == -1 {
                notFoundCnt++
            }
            if x != -1 && table[x][y] != i {
                fmt.Println("ERR")
            }
        }
        fmt.Println("notFoundCnt: ", notFoundCnt)
        fmt.Println("notFoundByFirst: ", notFoundByFirst)
        fmt.Println("foundByFirst: ", foundByFirst)
    }
    // test for each value
    //func () {
    //    for i := 301; i < 350; i++ {
    //        r000_0,r000_1 := 0,0
    //        r001_0,r001_1 := 0,0
    //        //if r000 {
    //        //    fmt.Println("ERR <-----------------")
    //        //}
    //        {
    //            start := time.Now()
    //            // calc
    //            r000_0,r000_1 = search000(table, i)
    //            //
    //            elapsed := time.Since(start)
    //            fmt.Printf("Time taken: %s\n", elapsed)
    //        }
    //        {
    //            start := time.Now()
    //            // calc
    //            C = 0
    //            r001_0,r001_1 = search002(table, 0,0,len(table),len(table[0]), i)
    //            //
    //            elapsed := time.Since(start)
    //            fmt.Printf("Time taken: %s\n", elapsed)
    //            fmt.Println("C: ", C)
    //        }
    //        fmt.Print(i, ":")
    //        fmt.Println(r000_0, r000_1)
    //        fmt.Print(i, ":")
    //        fmt.Println(r001_0, r001_1)
    //        fmt.Println("-----------")
    //    }
    //}()
}

