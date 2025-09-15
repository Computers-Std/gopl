// Exercise 4.10: Modify issues to report the results in age
// categories, say less than a month old, less than a year old, and
// more than a year old.
package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"ukiran/gopl/ch4/github"
)

type issueGroup struct {
	Heading    string
	IssuesList []*github.Issue
}

var monthOld, lessYearOld, moreYearOld issueGroup

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	categorizeByAge(result.Items)

	groups := []issueGroup{monthOld, lessYearOld, moreYearOld}
	for _, group := range groups {
		fmt.Printf("%v: %d issues\n", group.Heading, len(group.IssuesList))
		for _, item := range group.IssuesList {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func categorizeByAge(issues []*github.Issue) {
	for _, issue := range issues {
		switch calcAge(issue.CreatedAt) {
		case "<=1m":
			monthOld.Heading = "Less than a month old"
			monthOld.IssuesList = append(monthOld.IssuesList, issue)
		case "<=1y":
			lessYearOld.Heading = "\nLess than a year old"
			lessYearOld.IssuesList = append(lessYearOld.IssuesList, issue)
		case ">1y":
			moreYearOld.Heading = "\nMore than a year old"
			moreYearOld.IssuesList = append(moreYearOld.IssuesList, issue)
		}
	}

}

func calcAge(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	years := int(diff.Hours() / (24 * 365))
	months := int(diff.Hours() / (24 * 30))

	switch {
	case months <= 1:
		return "<=1m"
	case years <= 1:
		return "<=1y"
	default:
		return ">1y"
	}
}
