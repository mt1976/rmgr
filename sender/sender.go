package sender

import (
	"encoding/csv"
	"fmt"
	"os"

	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
)

var config = C.Configuration

func Run() error {
	file := "./data/" + config.SimulatorSourceFile
	f, err := os.Open(file)
	if err != nil {
		panic(E.ErrCannotOpenFile + err.Error())
		//return err
	}
	defer f.Close()

	//
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	//	fmt.Printf("records: %v\n", records)
	fmt.Printf("len(records): %v\n", len(records))
	// o := NewScanner(strings.NewReader(records))
	// for o.Scan() {
	// 	println(o.Text("Month"), o.Text("Day"))
	// }
	for recNo, rec := range records {
		fmt.Printf("rec %v: %v %v\n", recNo, rec, len(rec))

	}

	return nil

}
