package main

import (
        "fmt"
)

func main() {
        var cases int
        var rental, team int

        fmt.Scanf("%d", &cases)

        for cases > 0 {
                fmt.Scanf("%d%d", &rental, &team)

                price := make([]int, rental)
                for i := 0; i < rental; i++ {
                        fmt.Scanf("%d", &price[i])
                }

                // 가장 저렴한 대여료가 되는 페스티벌 시작-끝 날짜 확인
                // startDay, endDay, min := minPrice(rental, team, price)
                // fmt.Printf("(%d - %d): %.11f\n", startDay, endDay, min)
                _, _, min := minPrice(rental, team, price)
                fmt.Printf("%.11f\n", min)

                cases--
        }
}

func minPrice(rental, team int, price []int) (startDay, endDay int, min float32) {
        min = 3.0

        for rental >= team {
                offset := 0
                end := team

                for rental >= end {
                        tmp := float32(sum(price[offset:end])) / float32(team)
                        if tmp < min {
                                startDay = offset + 1
                                endDay = end
                                min = tmp
                        }
                        offset++
                        end++
                }

                team++
        }
        return
}

func sum(price []int) (sum int) {
        for _, v := range price {
                sum += v
        }
        return sum
}
