package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeMap struct {
    heightmap [][]*Coordinate
    maxCol, maxRow int
}

type Coordinate struct {
    height int
    visible bool
    score int
}

func (tm *TreeMap) genMapFromStr(input string) {
    for _, currentLine := range strings.Split(input, "\n") {
        if len(currentLine) > 0 {
            tm.heightmap = append(tm.heightmap, make([]*Coordinate, 0))
            for _, val := range strings.Split(currentLine, "") {
                currentRow := &(tm.heightmap[len(tm.heightmap) - 1])
                intVal, _ := strconv.Atoi(val)
                newCoordinate := Coordinate {
                    height: intVal,
                    visible: false,
                    score: 1,
                }
                *currentRow = append(*currentRow, &newCoordinate)
            }
        }
    }
}

func (tm *TreeMap) getRow(rowNum int) []*Coordinate {
    if rowNum >= tm.maxRow {
        fmt.Fprintln(os.Stderr, "Tried to read TreeMap out of bounds")
        os.Exit(3)
    }
    row := []*Coordinate{}
    for colIdx := 0; colIdx < tm.maxCol; colIdx++ {
        row = append(row, tm.heightmap[rowNum][colIdx])
    }
    return row
}

func (tm *TreeMap) getCol(colNum int) []*Coordinate {
    if colNum >= tm.maxCol {
        fmt.Fprintln(os.Stderr, "Tried to read TreeMap out of bounds")
        os.Exit(3)
    }
    col := []*Coordinate{}
    for _, row := range tm.heightmap {
        col = append(col, row[colNum])
    }
    return col
}

func isMarkVisible(treeIdx int, line []*Coordinate) bool {
    currentCoord := line[treeIdx]
    // if we already know its visible then why bother counting
    visibleFromLeft := true
    visibleFromRight := true
    lScore := 0
    rScore := 0
    // check to left
    for i := treeIdx - 1; i >= 0; i-- {
        lScore++;
        if line[i].height >= currentCoord.height {
            visibleFromLeft = false
            break
        }
    }
    for i := treeIdx + 1; i < len(line); i++ {
        rScore++;
        if line[i].height >= currentCoord.height {
            visibleFromRight = false
            break
        }
    }
    currentCoord.visible = (visibleFromLeft || visibleFromRight || currentCoord.visible)
    currentCoord.score = currentCoord.score * lScore * rScore
    return currentCoord.visible 
}

func (tm TreeMap) countVisible() int {
    for currentRow := 0; currentRow < tm.maxRow; currentRow++ {
        for currentCol := 0; currentCol < tm.maxCol; currentCol++ {
            isMarkVisible(currentCol, tm.getRow(currentRow))
            isMarkVisible(currentRow, tm.getCol(currentCol))
        }
    }
    totalVisible := 0
    for currentRow := 0; currentRow < tm.maxRow; currentRow++ {
        for currentCol := 0; currentCol < tm.maxCol; currentCol++ {
            if tm.heightmap[currentRow][currentCol].visible {
                totalVisible += 1;
            }
        }
    }
    return totalVisible
}

func (tm TreeMap) calculateMaxScore() int {
    maxScore := 0;
    for currentRow := 0; currentRow < tm.maxRow; currentRow++ {
        for currentCol := 0; currentCol < tm.maxCol; currentCol++ {
            if tm.heightmap[currentRow][currentCol].score > maxScore {
                maxScore = tm.heightmap[currentRow][currentCol].score
            }
        }
    }
    return maxScore
}

func (tm TreeMap) printHeightMap() {
    for _, currentLine := range tm.heightmap {
        // fmt.Printf("%2d: ", idx)
        for _, val := range currentLine {
            fmt.Print(val.height)
        }
        fmt.Printf("\n")
    }
    return
}

func main() {
    input, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, "Could not open input")
        os.Exit(2)
    }
    myMap := &TreeMap {
        maxCol: 99, 
        maxRow: 99,
    }
    myMap.genMapFromStr(string(input))
    fmt.Printf("# of visible: %d\n", myMap.countVisible())
    fmt.Printf("score: %d\n", myMap.calculateMaxScore())
}
