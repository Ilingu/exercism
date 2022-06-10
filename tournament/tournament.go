package tournament

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

type Teams = map[string]TeamInfo
type TeamInfo struct {
	rank   int
	points int
	MP     int
	W      int
	D      int
	L      int
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func SortTeams(UnorderTeams Teams) Teams {
	PointToTeams := make(map[int][]string)
	pointsSlice := make([]int, 0, len(PointToTeams))
	for TeamName, teamI := range UnorderTeams {
		if PointToTeams[teamI.points] == nil {
			pointsSlice = append(pointsSlice, teamI.points)
		}
		PointToTeams[teamI.points] = append(PointToTeams[teamI.points], TeamName)
	}
	sort.Ints(pointsSlice)
	reverse(pointsSlice)

	OrderedTeams := make(Teams)
	TeamRank := 1
	for _, OrderedPoints := range pointsSlice {
		whoseTeamPoints := PointToTeams[OrderedPoints]
		sort.Strings(whoseTeamPoints)
		for _, TeamName := range whoseTeamPoints {
			CurrentTeam := UnorderTeams[TeamName]
			OrderedTeams[TeamName] = TeamInfo{points: OrderedPoints, rank: TeamRank, MP: CurrentTeam.MP, W: CurrentTeam.W, D: CurrentTeam.D, L: CurrentTeam.L}
			TeamRank++
		}
	}
	return OrderedTeams
}

func ConvertToOutputTable(TeamsPoints Teams, writer io.Writer) {
	templateString := [5]string{"Team                           | MP |  W |  D |  L |  P\n"}
	for TeamName, Team := range TeamsPoints {
		espace := ""
		for i := 0; i < 31-len(TeamName); i++ {
			espace += " "
		}
		templateString[Team.rank] = fmt.Sprintf("%s%s|  %d |  %d |  %d |  %d |  %d\n", TeamName, espace, Team.MP, Team.W, Team.D, Team.L, Team.points)
	}

	TableString := strings.Join(templateString[:], "")
	writer.Write([]byte(TableString))
	// io.WriteString(writer, templateString)
}

func Tally(reader io.Reader, writer io.Writer) error {
	buffer, _ := io.ReadAll(reader)
	input := string(buffer)

	TeamsPoints := make(Teams)
	matches := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	if len(matches) == 1 {
		return errors.New("error")
	}
	for _, match := range matches {
		matchInfo := strings.Split(match, ";")
		if len(matchInfo) < 3 {
			continue
		}
		TeamA, TeamB, Outcome := matchInfo[0], matchInfo[1], matchInfo[2]

		TeamAInfo, TeamBInfo := TeamsPoints[TeamA], TeamsPoints[TeamB]
		NewTeamAInfo := TeamInfo{rank: TeamAInfo.rank, points: TeamAInfo.points, MP: TeamAInfo.MP, W: TeamAInfo.W, D: TeamAInfo.D, L: TeamAInfo.L}
		NewTeamBInfo := TeamInfo{rank: TeamBInfo.rank, points: TeamBInfo.points, MP: TeamBInfo.MP, W: TeamBInfo.W, D: TeamBInfo.D, L: TeamBInfo.L}
		if Outcome == "win" { // Team A wins
			NewTeamAInfo.points += 3
			NewTeamAInfo.W++
			NewTeamBInfo.L++
		}
		if Outcome == "draw" {
			NewTeamAInfo.points++
			NewTeamBInfo.points++
			NewTeamAInfo.D++
			NewTeamBInfo.D++
		}
		if Outcome == "loss" {
			NewTeamBInfo.points += 3
			NewTeamBInfo.W++
			NewTeamAInfo.L++
		}

		NewTeamAInfo.MP++
		NewTeamBInfo.MP++

		TeamsPoints[TeamA] = NewTeamAInfo
		TeamsPoints[TeamB] = NewTeamBInfo
	}
	TeamsPoints = SortTeams(TeamsPoints) // Sort with MP/W/D/l
	DebugString(TeamsPoints)
	ConvertToOutputTable(TeamsPoints, writer)

	return nil
}

func DebugString(teams Teams) {
	for TeamName, TeamI := range teams {
		fmt.Printf("%s --> %s \n", TeamName, TeamI)
	}
	fmt.Println("")
}

func (TI TeamInfo) String() string {
	return fmt.Sprintf("MP: %d // Win: %d // Draw: %d // Loss: %d // Point: %d // Rank: %d", TI.MP, TI.W, TI.D, TI.L, TI.points, TI.rank)
}
