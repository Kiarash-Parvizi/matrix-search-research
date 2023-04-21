package main

import(
    "fmt"
    "strconv"
    "time"
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

func mkTable000(_size int) [][]int {
    table := make([][]int, _size)
    for i := 0; i < _size; i++ {
        table[i] = make([]int, _size)
    }
    for i, r := 0, 0; i < _size; i++ {
        for j := 0; j < _size; j++ {
            table[i][j] = r
            r += 3
        }
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
            fmt.Print(zeroPad(table[i][j], 5))
        }
        fmt.Println("")
    }
}

var(
    C int
)

func main() {
    table := mkTable(12000)
    //printTable(table)
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
    fmt.Println("search002")
    {
        start := time.Now()
        for i := mn; i <= mx; i++ {
            x,_ := search002(table, 0,0, len(table), len(table[0]), i)
            if x == -1 {
                x,_ = search000(table, i)
            }
        }
        elapsed := time.Since(start)
        fmt.Printf("Time taken: %s\n", elapsed)
    }
    // correctness test
    fmt.Println("correctness test:")
    fmt.Println("search000")
    isEmpt := map[int]struct{}{}
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
    fmt.Println("search002")
    {
        notFoundCnt, notFoundByFirst, foundByFirst := 0, 0, 0
        for i := mn; i <= mx; i++ {
            x,y := search002(table, 0,0, len(table), len(table[0]), i)
            if x == -1 {
                x,y = search000(table, i)
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

