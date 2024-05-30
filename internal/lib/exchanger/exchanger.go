package exchanger

import (
	"errors"
)

func moneySlice2MoneyList(bnotes []int, bnotesCnt []int, res *[][]int) {
	sl := []int{}
	for i, cnt := range bnotesCnt {
		if cnt > 0 {
			for j := 0; j < cnt; j++ {
				sl = append(sl, bnotes[i])
			}
		}
	}

	*res = append(*res, sl)
	// fmt.Println(" -- ", summaOf(sl), sl)
}

func genCombinations(amount int, bnotes []int, bIndex int, startInd int, sl []int, res *[][]int) {
	slCopy := make([]int, len(sl))
	copy(slCopy, sl)

	if bIndex == len(bnotes)-1 {
		slCopy[bIndex] = amount / bnotes[bIndex]
		moneySlice2MoneyList(bnotes, slCopy, res)
		return
	}

	for j := startInd; ; j++ {
		slCopy = make([]int, len(sl))
		copy(slCopy, sl)
		if j*bnotes[bIndex] <= amount {
			slCopy[bIndex] = j

			if bnotes[bIndex]*j < amount {
				genCombinations(amount-bnotes[bIndex]*j, bnotes, bIndex+1, 0, slCopy, res)
			} else {
				moneySlice2MoneyList(bnotes, slCopy, res)
			}
		} else {
			break
		}
	}
}

func Exchange(amount int, bnotes []int) ([][]int, error) {
	if len(bnotes) < 1 {
		return [][]int{}, errors.New("failed banknotes list")
	}

	if amount < 0 || amount%bnotes[len(bnotes)-1] > 0 {
		return [][]int{}, errors.New("failed amount")
	}

	if amount == 0 {
		return [][]int{{0}}, nil
	}

	res := [][]int{}

	for i := len(bnotes) - 1; i >= 0; i-- {
		if bnotes[i] > amount {
			break
		}

		sl := make([]int, len(bnotes))
		genCombinations(amount, bnotes, i, 1, sl, &res)
	}

	return res, nil
}
