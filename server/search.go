package server

import (
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
)

type SearchResultStruct struct {
	Name       string
	Image      string
	Link       string
	Percentage float64
}

type DataSearchStruct struct {
	Data    []SearchResultStruct
	IsLogin bool
}

func Search(w http.ResponseWriter, r *http.Request) {
	UserSearch := r.URL.Query().Get(("query"))
	Data := DataSearchStruct{
		Data:    SearchFunction(UserSearch),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "ResultOfTheSearch", Data)
	if err != nil {
		log.Fatal(err)
	}
}

func SearchFunction(UserSearch string) []SearchResultStruct {
	var PercentageAcceptable float64 = 60.0
	var temp SearchResultStruct
	var ListOfPercentage []float64
	var ResultOfSerch []SearchResultStruct = []SearchResultStruct{}
	//artifact
	ListOfPercentage = []float64{}
	var wordTwoPieceBonus string
	var wordFourPieceBonus string
	var wordOnePieceBonus string
	for _, artifact := range API_Data.ALLArtifacts {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  artifact.Name,
			Image: artifact.ImageURL,
			Link:  "/artifacts/details?id=" + artifact.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, artifact.Name))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, strconv.Itoa(artifact.MaxRarity)))
		if artifact.OnePieceBonus == "" {
			wordTwoPieceBonus = ""
			for _, char := range artifact.TwoPiecesBonus {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordTwoPieceBonus)))
					wordTwoPieceBonus = ""
				} else {
					wordTwoPieceBonus += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordTwoPieceBonus)))
			wordFourPieceBonus = ""
			for _, char := range artifact.FourPiecesBonus {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordFourPieceBonus)))
					wordFourPieceBonus = ""
				} else {
					wordFourPieceBonus += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordFourPieceBonus)))
		} else {
			wordOnePieceBonus = ""
			for _, char := range artifact.OnePieceBonus {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordOnePieceBonus)))
					wordOnePieceBonus = ""
				} else {
					wordOnePieceBonus += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordOnePieceBonus)))
		}
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	//boss
	ListOfPercentage = []float64{}
	var wordBossDescription string
	for _, boss := range API_Data.AllBoss {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  boss.Name,
			Image: boss.ImageURL,
			Link:  "/boss/details?id=" + boss.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, boss.Name))
		wordBossDescription = ""
		for _, char := range boss.Description {
			if char == ' ' {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordBossDescription)))
				wordBossDescription = ""
			} else {
				wordBossDescription += string(char)
			}
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordBossDescription)))
		for _, artifact := range boss.Artifacts {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, artifact.Name))
		}
		for _, drop := range boss.Drops {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, drop.Name))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, drop.Source))
		}
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	//characters
	var wordcharacterConstellation string
	for _, character := range API_Data.AllCharacters {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  character.Name,
			Image: character.ImageUrl,
			Link:  "/characters/details?id=" + character.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Name))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Affiliation))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Birthday))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Constellation))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Gender))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Nation))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Release))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Title))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Vision))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.VisionKey))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.Weapon))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, character.WeaponType))
		for _, element := range character.AscensionMaterials.Level20 {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		}
		for _, element := range character.AscensionMaterials.Level40 {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		}
		for _, element := range character.AscensionMaterials.Level50 {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		}
		for _, element := range character.AscensionMaterials.Level60 {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		}
		for _, element := range character.AscensionMaterials.Level70 {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		}
		for _, element := range character.AscensionMaterials.Level80 {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		}
		for _, element := range character.Constellations {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Unlock))
			wordcharacterConstellation = ""
			for _, char := range element.Description {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordcharacterConstellation)))
					wordcharacterConstellation = ""
				} else {
					wordcharacterConstellation += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordcharacterConstellation)))
		}
		for _, element := range character.PassiveTalents {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Unlock))
			wordcharacterConstellation = ""
			for _, char := range element.Description {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordcharacterConstellation)))
					wordcharacterConstellation = ""
				} else {
					wordcharacterConstellation += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordcharacterConstellation)))
		}
		for _, element := range character.SkillTalents {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Unlock))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Type))
			wordcharacterConstellation = ""
			for _, char := range element.Description {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordcharacterConstellation)))
					wordcharacterConstellation = ""
				} else {
					wordcharacterConstellation += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordcharacterConstellation)))
			for _, attribute := range element.AttributeScaling {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, attribute.Name))
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, attribute.Value))
			}
			for _, upgrade := range element.Upgrades {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, upgrade.Name))
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, upgrade.Value))
			}
		}
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	//domain
	ListOfPercentage = []float64{}
	for _, domain := range API_Data.AllDomain {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  domain.Name,
			Image: domain.ImageUrl,
			Link:  "/domains/details?id=" + domain.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, domain.Name))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, domain.Id))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, domain.Location))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, domain.Nation))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, domain.Type))
		wordelement := ""
		for _, char := range domain.Description {
			if char == ' ' {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordelement)))
				wordelement = ""
			} else {
				wordelement += string(char)
			}
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordelement)))
		for _, recomendedelement := range domain.RecommendedElements {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, recomendedelement))
		}
		for _, requirement := range domain.Requirements {
			for _, sub := range requirement.LeyLineDisorder {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, sub))
			}
		}
		for _, reward := range domain.Rewards {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, reward.Day))
			for _, detail := range reward.Details {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, strconv.Itoa(detail.AdventureExperience)))
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, strconv.Itoa(detail.CompanionshipExperience)))
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, strconv.Itoa(detail.Mora)))
				for _, drop := range detail.Drops {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, drop.Name))
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, strconv.Itoa(drop.DropMax)))
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, strconv.Itoa(drop.DropMin)))
				}
			}
		}
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	//element
	ListOfPercentage = []float64{}
	var wordelement string
	for _, element := range API_Data.AllElement {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  element.Name,
			Image: element.ImageUrl,
			Link:  "/elements/details?id=" + element.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Name))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Id))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, element.Key))
		for _, reaction := range element.Reactions {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, reaction.Name))
			wordelement = ""
			for _, char := range reaction.Description {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordelement)))
					wordelement = ""
				} else {
					wordelement += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordelement)))
			for _, el := range reaction.Elements {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, el))
			}
		}
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	//Enemies
	ListOfPercentage = []float64{}
	var wordEnemie string
	for _, enemie := range API_Data.AllEnnemies {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  enemie.Name,
			Image: enemie.ImageUrl,
			Link:  "/enemies/details?id=" + enemie.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, enemie.Name))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, enemie.Type))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, enemie.Region))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, enemie.Faction))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, enemie.Family))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, enemie.Id))
		for _, artifact := range enemie.Artifacts {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, artifact.Name))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, artifact.Rarity))
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, artifact.Set))
		}
		for _, description := range enemie.Descriptions {
			wordEnemie = ""
			for _, char := range description.Description {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordEnemie)))
					wordEnemie = ""
				} else {
					wordEnemie += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordEnemie)))
		}
		wordEnemie = ""
		for _, char := range enemie.Description {
			if char == ' ' {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordEnemie)))
				wordEnemie = ""

			} else {
				wordEnemie += string(char)
			}
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordEnemie)))
		for _, drop := range enemie.Drops {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, drop.Name))
		}
		for _, elementaldescription := range enemie.ElementalDescriptions {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, elementaldescription.Element))
			wordEnemie = ""
			for _, char := range elementaldescription.Description {
				if char == ' ' {
					ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordEnemie)))
					wordEnemie = ""
				} else {
					wordEnemie += string(char)
				}
			}
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordEnemie)))
		}
		for _, el := range enemie.Elements {
			ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, el))
		}
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	//weapons
	ListOfPercentage = []float64{}
	var wordweapons string
	for _, weapon := range API_Data.AllWeapons {
		ListOfPercentage = []float64{}
		temp = SearchResultStruct{
			Name:  weapon.Name,
			Image: weapon.ImageUrl,
			Link:  "/weapons/details?id=" + weapon.Id,
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, weapon.Name))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, weapon.Id))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, weapon.AscensionMaterial))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, weapon.Location))
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, weapon.SubStat))
		wordweapons = ""
		for _, char := range weapon.PassiveDescription {
			if char == ' ' {
				ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordweapons)))
				wordweapons = ""
			} else {
				wordweapons += string(char)
			}
		}
		ListOfPercentage = append(ListOfPercentage, SimilarityPercentage(UserSearch, string(wordweapons)))
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			ResultOfSerch = append(ResultOfSerch, temp)
		}
	}
	sort.Slice(ResultOfSerch, func(i, j int) bool {
		return ResultOfSerch[i].Percentage > ResultOfSerch[j].Percentage
	})
	return ResultOfSerch
}

// Levenshtein
// Fonction pour calculer la distance de Levenshtein entre deux mots
func LevenshteinDistance(a, b string) int {
	la, lb := len(a), len(b)
	matrix := make([][]int, la+1)

	for i := range matrix {
		matrix[i] = make([]int, lb+1)
	}

	// Initialiser la première ligne et la première colonne
	for i := 0; i <= la; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		matrix[0][j] = j
	}

	// Calculer la distance
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			matrix[i][j] = MinLevenshtein(
				matrix[i-1][j]+1,      // Suppression
				matrix[i][j-1]+1,      // Insertion
				matrix[i-1][j-1]+cost, // Substitution
			)
		}
	}

	return matrix[la][lb]
}

// Fonction utilitaire pour trouver le minimum de trois entiers
func MinLevenshtein(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

// Fonction pour calculer le pourcentage de ressemblance
func SimilarityPercentage(a, b string) float64 {
	distance := LevenshteinDistance(a, b)
	maxLen := math.Max(float64(len(a)), float64(len(b)))

	// Si les deux mots sont vides, ils sont considérés comme 100% similaires
	if maxLen == 0 {
		return 100.0
	}

	// Calculer le pourcentage de ressemblance
	return (1 - float64(distance)/maxLen) * 100
}
