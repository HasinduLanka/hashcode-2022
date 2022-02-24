package main

import (
	"os"
	"strconv"

	"github.com/HasinduLanka/console"
)

type TestCase struct {
	ContribCount int
	ProjCount    int

	Contribs []*Contrib
	Projects []*Project

	ContribsByName map[string]*Contrib
	ProjectsByName map[string]*Project
}

type Contrib struct {
	Name      string
	Skills    map[string]int
	Available int
}

type Project struct {
	Name       string
	Days       int
	Score      int
	BestBefore int

	Roles []Skill

	Start int
	End   int
}

type Skill struct {
	Name  string
	Level int
}

type Solution struct {
	Assignments []*Assignment
	Score       int
	TC          *TestCase `json:"-"`
}

type Assignment struct {
	Proj     *Project
	Contribs []*Contrib
}

func main() {
	os.Mkdir("out", 0777)

	console.GlobalWriter = console.NewWriterToStandardOutput()
	console.GlobalReader = console.NewReaderFromStandardInput()

	console.Print("Hashcode 2022")

	FC := console.NewWriterToFile("out/out.txt")
	FC.Print("Hashcode 2022 File!")

	tc := ParseTestCaseFromFile("cases/5.txt")
	// console.Log(tc)

	sol1 := tc.ParseSolutionFromFile("out/5.txt")
	// console.Log(sol1)
	console.Log(sol1.Eval())

}

func ParseTestCase(S string) *TestCase {
	tc := &TestCase{}

	fc := console.NewReaderFromString(S)

	// A := strings.Split(S, "\n")

	l1 := fc.ReadArrayClean(" ")
	C, _ := strconv.ParseInt(l1[0], 10, 32)
	tc.ContribCount = int(C)

	P, _ := strconv.ParseInt(l1[1], 10, 32)
	tc.ProjCount = int(P)

	tc.Contribs = make([]*Contrib, tc.ContribCount)
	tc.ContribsByName = make(map[string]*Contrib, tc.ContribCount)

	for i := 0; i < tc.ContribCount; i++ {
		l := fc.ReadArrayClean(" ")

		cntr := &Contrib{
			Name:      l[0],
			Available: 0,
		}

		skillCount, _ := strconv.ParseInt(l[1], 10, 32)
		cntr.Skills = make(map[string]int, skillCount)

		for j := 0; j < int(skillCount); j++ {
			skl := fc.ReadArrayClean(" ")

			sklLvl, _ := strconv.ParseInt(skl[1], 10, 32)

			cntr.Skills[skl[0]] = int(sklLvl)
		}

		tc.Contribs[i] = cntr
		tc.ContribsByName[cntr.Name] = cntr

	}

	tc.Projects = make([]*Project, tc.ProjCount)
	tc.ProjectsByName = make(map[string]*Project, tc.ProjCount)

	for i := 0; i < tc.ProjCount; i++ {
		l := fc.ReadArrayClean(" ")

		days, _ := strconv.ParseInt(l[1], 10, 32)
		score, _ := strconv.ParseInt(l[2], 10, 32)
		bestBefore, _ := strconv.ParseInt(l[3], 10, 32)
		roleCount, _ := strconv.ParseInt(l[4], 10, 32)

		proj := &Project{
			Name:       l[0],
			Days:       int(days),
			Score:      int(score),
			BestBefore: int(bestBefore),
		}

		proj.Roles = make([]Skill, roleCount)

		for j := 0; j < int(roleCount); j++ {
			skl := fc.ReadArrayClean(" ")

			sklLvl, _ := strconv.ParseInt(skl[1], 10, 32)

			proj.Roles[j] = Skill{
				Name:  skl[0],
				Level: int(sklLvl),
			}
		}

		tc.Projects[i] = proj
		tc.ProjectsByName[proj.Name] = proj
	}

	return tc
}

func ParseTestCaseFromFile(filename string) *TestCase {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return ParseTestCase(string(file))
}

func (tc *TestCase) ParseSolution(S string) *Solution {
	fc := console.NewReaderFromString(S)

	nOfAssignments, _ := strconv.ParseInt(fc.ReadLine(), 10, 32)

	sol := &Solution{TC: tc}
	sol.Assignments = make([]*Assignment, nOfAssignments)

	for i := 0; i < int(nOfAssignments); i++ {

		name := fc.ReadLine()

		assignment := &Assignment{
			Proj: tc.ProjectsByName[name],
		}

		contribNames := fc.ReadArrayClean(" ")
		assignment.Contribs = make([]*Contrib, len(contribNames))

		for i, name := range contribNames {
			assignment.Contribs[i] = tc.ContribsByName[name]
		}

		sol.Assignments[i] = assignment
	}

	return sol
}

func (tc *TestCase) ParseSolutionFromFile(filename string) *Solution {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return tc.ParseSolution(string(file))
}

func (sol *Solution) Eval() int {

	FinalScore := 0

	for _, assignment := range sol.Assignments {

		if len(assignment.Proj.Roles) != len(assignment.Contribs) {
			console.Log(assignment)
			panic("Assignment has different number of roles and contribs")
		}

		for i, role := range assignment.Proj.Roles {
			contr := assignment.Contribs[i]

			skillLvl, skillFound := contr.Skills[role.Name]
			if !skillFound {
				skillLvl = 0
			}

			if skillLvl == role.Level {
				contr.Skills[role.Name]++

			} else if skillLvl == role.Level-1 {
				for _, mentor := range assignment.Contribs {
					mentSkl, mentSklFound := mentor.Skills[role.Name]
					if mentSklFound {
						if mentSkl >= role.Level {
							contr.Skills[role.Name]++
						}
					}
				}

			} else if skillLvl < role.Level {
				console.Log(sol)
				panic("Not all roles were assigned to a contributor in project " + assignment.Proj.Name)

			}
		}

		// for _, roleSkil := range assignment.Proj.Roles {
		// 	for _, contr := range assignment.Contribs {
		// 		role := roleSkil.Name
		// 		roleLvl := roleSkil.Level

		// 		sklLvl, sklFound := contr.Skills[role]
		// 		if !sklFound {
		// 			if roleLvl == 1 {
		// 				sklLvl = 0
		// 				contr.Skills[role] = sklLvl
		// 			}
		// 		}

		// 		if sklLvl == roleLvl {
		// 			contr.Skills[role]++
		// 			delete(roles, role)
		// 			console.Print("Assigned " + role + " to " + contr.Name)
		// 			break

		// 		} else if sklLvl == roleLvl-1 {
		// 			for _, mentor := range assignment.Contribs {
		// 				mentSkl, mentSklFound := mentor.Skills[role]
		// 				if mentSklFound {
		// 					if mentSkl >= roleLvl {
		// 						contr.Skills[role]++
		// 					}
		// 				}
		// 			}
		// 			delete(roles, role)
		// 			console.Print("Assigned " + role + " to " + contr.Name + " MENTORED")
		// 			break

		// 		}

		// 	}
		// }

		// if len(roles) > 0 {
		// 	console.Log(sol)
		// 	console.Log(roles)
		// 	panic("Not all roles were assigned to a contributor in project " + assignment.Proj.Name)
		// }

		start := 0

		for _, contr := range assignment.Contribs {
			if contr.Available < start {
				start = contr.Available
			}
		}

		end := start + assignment.Proj.Days

		for _, contr := range assignment.Contribs {
			contr.Available = end

		}

		assignment.Proj.Start = start
		assignment.Proj.End = end

		// assignment.Proj.Score

		score := 0

		if assignment.Proj.End > assignment.Proj.BestBefore {
			score = assignment.Proj.Score - (assignment.Proj.End - assignment.Proj.BestBefore)
			if score < 0 {
				score = 0
			}

		} else {
			score = assignment.Proj.Score

		}

		FinalScore += score

	}

	return FinalScore
}
