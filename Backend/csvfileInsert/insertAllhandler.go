package csvfileinsert

import (
	"busproject/model"
	"log"
	"net/http"
)

func CreateAllHandler(w http.ResponseWriter, r *http.Request) {

	var schedules []model.Schedule

	InsertAll("csvs/Bus_Route_Shedule - Sheet1.csv", nil, nil, &schedules, nil, nil, nil)
	log.Println(schedules)
	log.Println("code is broken first complete the code")
	// for _, schedule := range schedules {
	// err := database.InsertSchedule(c.DB, schedule)
	// if err != nil {
	// log.Println(err.Error())
	// return
	// }
	// }
	log.Println("For lopp chal gaya")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfull All"))

}
