package api

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/svkorch/mmine/internal/lib/exchanger"

	"github.com/sirupsen/logrus"
)

type Req struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type Resp struct {
	Exchanges [][]int `json:"exchanges"`
}

func Exchange(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("RemoteAddr: [%s]", r.RemoteAddr)

	inData := Req{}
	err := json.NewDecoder(r.Body).Decode(&inData)
	if err != nil {
		logrus.Error("json body parsing error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	logrus.Infof("[1] amount: %d, banknotes: %v", inData.Amount, inData.Banknotes)
	slices.Reverse(inData.Banknotes)
	if !slices.IsSorted(inData.Banknotes) {
		logrus.Warn("need to be sorted")
		slices.Sort(inData.Banknotes)
	}
	slices.Reverse(inData.Banknotes)
	logrus.Infof("[2] amount: %d, banknotes: %v", inData.Amount, inData.Banknotes)

	logrus.Info("exchange of ", inData.Amount)

	sl, err := exchanger.Exchange(inData.Amount, inData.Banknotes)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(sl)

	outData := Resp{sl}
	// outJSON, err := json.MarshalIndent(outData, "", "\t")
	outJSON, err := json.Marshal(outData)
	if err != nil {
		logrus.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(outJSON)

	logrus.Info(string(outJSON))
}
