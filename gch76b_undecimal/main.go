package main

import (
	"fmt"

	"github.com/bitlux/caches/util"
)

func main() {
	for T := range 11 {
		for E := range 11 {
			if !util.IsUnique(T, E) {
				continue
			}
			for N := range 11 {
				if !util.IsUnique(T, E, N) {
					continue
				}

				sum := T + E + E
				if sum%11 != N {
					continue
				}
				for I := range 11 {
					if !util.IsUnique(T, E, N, I) {
						continue
					}
					for M := range 11 {
						if !util.IsUnique(T, E, N, I, M) {
							continue
						}
						for O := range 11 {
							if !util.IsUnique(T, E, N, I, M, O) {
								continue
							}

							if (I+N+M+sum/11)%11 != O {
								continue
							}
							for S := range 11 {
								if !util.IsUnique(T, E, N, I, M, O, S) {
									continue
								}
								for U := range 11 {
									if !util.IsUnique(T, E, N, I, M, O, S, U) {
										continue
									}
									for L := range 11 {
										if !util.IsUnique(T, E, N, I, M, O, S, U, L) {
											continue
										}
										for P := range 11 {
											if !util.IsUnique(T, E, N, I, M, O, S, U, L, P) {
												continue
											}
											for H := range 11 {
												if !util.IsUnique(T, E, N, I, M, O, S, U, L, P, H) {
													continue
												}

												sunlit := util.FromDigitsBase([]int{S, U, N, L, I, T}, 11)
												lupine := util.FromDigitsBase([]int{L, U, P, I, N, E}, 11)
												home := util.FromDigitsBase([]int{H, O, M, E}, 11)
												nilupon := util.FromDigitsBase([]int{N, I, L, U, P, O, N}, 11)
												if sunlit+lupine+home == nilupon {
													fmt.Printf("I.UNL = %d.%d%d%d H.TPS = %d.%d%d%d\n", I, U, N, L, H, T, P, S)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
