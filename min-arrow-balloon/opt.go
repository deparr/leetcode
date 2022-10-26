import "sort"

func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func (i, j int) bool {return points[i][1] < points[j][1]})

    res, arrow := 0, 0

    for _, p := range points {
        if res == 0 || p[0] > arrow{
            res++
            arrow = p[1]
        }
    }

    return res
}
