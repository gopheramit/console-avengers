package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Avenger struct {
	Name       string
	PersonName string
	Assigned   int
	Completed  int
	Abilities  string
}

type Mission struct {
	Team           []string
	MissionName    string
	MissionStatus  string
	MissionDetails string
}

func updateCount(avenger []Avenger, name string) []Avenger {
	for i, _ := range avenger {
		if avenger[i].Name == name {
			avenger[i].Assigned = avenger[i].Assigned + 1
		}
	}
	return avenger

}

func checkCount(avenger []Avenger, name string) (int, int) {

	for i, _ := range avenger {
		if avenger[i].Name == name {
			//fmt.Println(avenger[i].Name)
			//fmt.Println(name)
			return avenger[i].Assigned, 1
		}
	}
	return 999, 999
}
func checkMission(mission []Mission) {
	for i, _ := range mission {
		fmt.Println(mission[i])
	}
}
func getStatus(avenger []Avenger, name string) string {
	var a int
	for i, _ := range avenger {
		if avenger[i].Name == name {
			a = avenger[i].Assigned
		}
		if a == 0 {
			return "Available"
		} else {
			return "On Mission"
		}

	}
	return ""
}

func getMission(mission []Mission, name string) []string {
	task := []string{}
	for i, _ := range mission {
		for j, _ := range mission[i].Team {
			if mission[i].Team[j] == name {
				task = append(task, mission[i].MissionName)
			}
		}

	}
	return task
}

func listAvengers(avenger []Avenger, mission []Mission) {
	fmt.Println("Avenger Name |Status |Assigned Mission ")
	var status1 string

	for i, _ := range avenger {
		status1 = getStatus(avenger, avenger[i].Name)
		task := getMission(mission, avenger[i].Name)
		if len(task) == 0 {
			fmt.Println(avenger[i].Name, status1)
		} else {
			fmt.Println(avenger[i].Name, status1, task)
		}
	}
}

func assignMission(mission []Mission, team []string, name, details string) []Mission {
	newMission := &Mission{
		Team:           team,
		MissionName:    name,
		MissionStatus:  "Assigned",
		MissionDetails: details,
	}
	mission = append(mission, *newMission)
	return mission

}

func getAvengersDetails(avenger []Avenger, name string) (Avenger, int) {
	//fmt.Println(name)
	for i, _ := range avenger {
		//fmt.Println(avenger[i].Name)
		//fmt.Println(name)
		//fmt.Println(avenger[i].Name == name)
		if avenger[i].Name == name {

			//fmt.Println(avenger[i].Name)
			//fmt.Println(name)
			return avenger[i], 1

		}
	}
	return avenger[0], 0

}

func getMissionDetails(mission []Mission, name string) (Mission, int) {

	for _, mi := range mission {
		if mi.MissionName == name {
			return mi, 1
		}
	}
	return mission[0], 0

}

func updateMissionStatus(mission []Mission, name, status string) []Mission {

	for i, _ := range mission {
		if mission[i].MissionName == name {
			mission[i].MissionStatus = status

		}
	}
	return mission

}
func consolePrint() {
	fmt.Println("=======------S.H.I.E.L.D ------=========")
	fmt.Println("Welcome to Captain Fury")
	fmt.Println("1.Check the missions")
	fmt.Println("2.Assign mission to Avengers")
	fmt.Println("3.Check mission’s details")
	fmt.Println("4.Check Avenger’s details")
	fmt.Println("5.Update Mission’s status")
	fmt.Println("6.List Avengers")

}

func strInput() string {

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	return name

}

func printOutput(mi Mission) {
	fmt.Println("Mission Details:", mi.MissionDetails)
	fmt.Println("Mission Status:", mi.MissionStatus)
	fmt.Println("Avengers:t", mi.Team)
}

func printAvenger(avg Avenger) {
	fmt.Println("Person Name:", avg.PersonName)
	fmt.Println("Abilities:", avg.Abilities)
	fmt.Println("Mission Assigned :", avg.Assigned)
	fmt.Println("Mission Completed :", avg.Completed)
}
func main() {

	var avenger = []Avenger{
		Avenger{
			Name:       "Iron Man",
			PersonName: "Tony Stark",
			Assigned:   0,
			Completed:  0,
			Abilities:  "Genius,Rich,Tech armored suit",
		},
		Avenger{
			Name:       "Captain America",
			PersonName: "Steve Rogers",
			Assigned:   0,
			Completed:  0,
			Abilities:  "Leadrship,ethics,super soldier",
		},
		Avenger{
			Name:       "Hulk",
			PersonName: "Bruce Banner",
			Assigned:   0,
			Completed:  0,
			Abilities:  "smart,strongest avenger",
		},
		Avenger{
			Name:       "Thor",
			PersonName: "Thor",
			Assigned:   0,
			Completed:  0,
			Abilities:  "God of Thunder",
		},
		Avenger{
			Name:       "Black Widow",
			PersonName: "Natasha Rommanof",
			Assigned:   0,
			Completed:  0,
			Abilities:  "smart,spy",
		},
		Avenger{
			Name:       "Hawkeye",
			PersonName: "Clint Barton",
			Assigned:   0,
			Completed:  0,
			Abilities:  "Guy who can aim very well and see evrything",
		},
	}

	var mission = []Mission{}

	for {
		consolePrint()
		var a int
		fmt.Println("Enter the option:")
		_, err := fmt.Scanf("%d", &a)
		if err != nil {
			fmt.Println("Invalid Input from usser")
		}

		switch a {
		case 1:
			if len(mission) == 0 {
				fmt.Println("No Mission has been assigned to an Avenger.")
			}
			checkMission(mission)

		case 2:
			var name string
			flag1 := true
			var details string
			fmt.Println("Enter Avengers:")
			name = strInput()
			team := strings.Split(name, ",")

			for i, _ := range team {
				cnt, _ := checkCount(avenger, team[i])
				//if flag == 1 {
				if cnt >= 2 {
					fmt.Printf("Sorry, %s as already been working on two missions", team[i])
					flag1 = false
				} else {
					avenger = updateCount(avenger, team[i])
					//fmt.Println("Enter Mission Name:")
					//name = strInput()
					//fmt.Println("Enter Mission Deatails")
					//details = strInput()
					//mission = assignMission(mission, team, name, details)
					//fmt.Println("Mission has been assigned.Email notification sent and/or SMS notification sent.")
				}
				//}
			}
			if flag1 == true {

				fmt.Println("Enter Mission Name:")
				name = strInput()
				fmt.Println("Enter Mission Deatails")
				details = strInput()
				mission = assignMission(mission, team, name, details)
				fmt.Println("Mission has been assigned.")
				fmt.Println("notification sent and/or SMS notification sent.")

			}

		case 3:
			var name string
			fmt.Println("Enter Mission Name:")
			name = strInput()
			mi, i := getMissionDetails(mission, name)
			if i != 1 {
				fmt.Println("Not Found")
			} else {
				printOutput(mi)
			}

		case 4:
			var name string
			fmt.Println("Enter Avenger Name:")
			name = strInput()
			//fmt.Println(name)
			avg, i := getAvengersDetails(avenger, name)
			//fmt.Println(avg, i)
			if i != 1 {
				fmt.Println("Avenger is not in the list")
			} else {
				printAvenger(avg)
			}
		case 5:
			var name string
			var status string
			fmt.Println("Enter Mission Name:")
			name = strInput()
			fmt.Println("Enter Mission Status")
			status = strInput()
			mission = updateMissionStatus(mission, name, status)
			team, _ := getMissionDetails(mission, name)

			fmt.Println("Email and Sms are sent to", team.Team)

		case 6:
			listAvengers(avenger, mission)

		}

	}

}
