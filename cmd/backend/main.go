package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	t2 := t1.Add(time.Minute)
	fmt.Printf("__%s\n", convertTimeRangeToWhereClause2(&t1, &t2))
	fmt.Printf("__%s\n", convertTimeRangeToWhereClause2(&t1, nil))
	fmt.Printf("__%s\n", convertTimeRangeToWhereClause2(nil, &t2))
	fmt.Printf("__%s\n", convertTimeRangeToWhereClause2(nil, nil))
}

func convertTimeRangeToWhereClause2(from, to *time.Time) string {
	var conditions []string

	if from != nil {
		conditions = append(conditions, fmt.Sprintf("episode_start >= '%s'", from.Format("2006-01-02 15:04:05")))
	}
	if to != nil {
		conditions = append(conditions, fmt.Sprintf("episode_start <= '%s'", to.Format("2006-01-02 15:04:05")))
	}

	return strings.Join(conditions, " AND ")
}
