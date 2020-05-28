package main

import (
	"fmt"

	"github.com/baking-bad/bcdhub/internal/database"
	"github.com/baking-bad/bcdhub/internal/elastic"
	"github.com/schollz/progressbar/v3"
)

func createTasks(dbConn, esConn string, userID uint, offset, size int) error {
	es := elastic.WaitNew([]string{esConn})

	db, err := database.New(dbConn)
	if err != nil {
		return err
	}
	defer db.Close()

	allTasks, err := es.GetDiffTasks()
	if err != nil {
		return err
	}

	fmt.Printf("Total %d pairs, picking %d:%d\n", len(allTasks), offset, offset+size)

	tasks := allTasks[offset : offset+size]

	bar := progressbar.NewOptions(len(tasks), progressbar.OptionSetPredictTime(false))
	for _, diff := range tasks {
		bar.Add(1)
		a := database.Assessments{
			Address1:   diff.Address1,
			Network1:   diff.Network1,
			Address2:   diff.Address2,
			Network2:   diff.Network2,
			UserID:     userID,
			Assessment: database.AssessmentUndefined,
		}
		if err := db.CreateAssessment(&a); err != nil {
			fmt.Print("\033[2K\r")
			return err
		}
	}
	fmt.Print("\033[2K\r")

	return nil
}
