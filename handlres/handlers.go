package handlers

import (
	"context"

	myfile "github.com/Mukunth-arya/grcp/proto"
	"github.com/hashicorp/go-hclog"
)

type datas struct {
	l hclog.Logger
}

func LoggerValue(l hclog.Logger) *datas {
	return &datas{l}
}

var datas1 = []int32{}
var data int

func (l *datas) GetValue(ctx context.Context, file *myfile.ValueRequest) (*myfile.ValueReponse, error) {
	l.l.Info("Here it do the cgpa calculator", file.GetSub1())

	var Calc_total = (file.Sub1 + file.Sub2 + file.Sub3 + file.Sub4)

	datas1 = append(datas1, file.Sub1)
	datas1 = append(datas1, file.Sub2)
	datas1 = append(datas1, file.Sub3)
	datas1 = append(datas1, file.Sub4)
	for index, _ := range datas1 {
		data = index + 1

	}
	var Avg_cgpa = Calc_total / int32(data)
	var mul = int32(data) * Avg_cgpa
	var _ = mul - Calc_total // remainder for sample purpose
	return &myfile.ValueReponse{Finalvalue: int32(Avg_cgpa)}, nil
}
