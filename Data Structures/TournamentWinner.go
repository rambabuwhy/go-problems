/*

  There's an algorithms tournament taking place in which teams of programmers
  compete against each other to solve algorithmic problems as fast as possible.
  Teams compete in a round robin, where each team faces off against all other
  teams. Only two teams compete against each other at a time, and for each
  competition, one team is designated the home team, while the other team is the
  away team. In each competition there's always one winner and one loser; there
  are no ties. A team receives 3 points if it wins and 0 points if it loses. The
  winner of the tournament is the team that receives the most amount of points.
  array contains information about the winner of each corresponding competition in the
  "competitions" [home team, awayteam]
   array. Specifically, results[i]  denotes the winner of  competitions[i]
   0 means away team wins
   1 means home team wins
  
  */



package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	//step 1:  fill map
	lmap := map[string]int{}
	for i := 0; i < len(results); i++ {

		homeTeam, awayTeam := competitions[i][0], competitions[i][1]
		result := results[i]
		if result == 0 {
			lmap[awayTeam] += 3

		} else {
			lmap[homeTeam] += 3

		}
	}

	//step 2: find max in map
	max := 0
	team := ""
	for k, v := range lmap {
		if max < v {
			max = v
			team = k
		}
	}
	return team
}
