package service

//Structure for the artifacts

type ArtifactDetails struct {
	Name            string `json:"name"`
	MaxRarity       int    `json:"max_rarity"`
	OnePieceBonus   string `json:"1-piece_bonus"`
	TwoPiecesBonus  string `json:"2-piece_bonus"`
	FourPiecesBonus string `json:"4-piece_bonus"`
}

//Structur for the boss
type BossStruct struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Drops       []Drop     `json:"drops"`
	Artifacts   []Artifact `json:"artifacts"`
}

type Drop struct {
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
	Source string `json:"source"`
}

type Artifact struct {
	Name      string `json:"name"`
	MaxRarity int    `json:"max_rarity"`
}

//Structur for the characters
type Characters struct {
	Name               string            `json:"name"`
	Title              string            `json:"title"`
	Vision             string            `json:"vision"`
	Weapon             string            `json:"weapon"`
	Gender             string            `json:"gender"`
	Nation             string            `json:"nation"`
	Affiliation        string            `json:"affiliation"`
	Rarity             int               `json:"rarity"`
	Release            string            `json:"release"`
	Constellation      string            `json:"constellation"`
	Birthday           string            `json:"birthday"`
	Description        string            `json:"description"`
	SkillTalents       []SkillTalent     `json:"skillTalents"`
	PassiveTalents     []PassiveTalent   `json:"passiveTalents"`
	Constellations     []Constellation   `json:"constellations"`
	VisionKey          string            `json:"vision_key"`
	WeaponType         string            `json:"weapon_type"`
	AscensionMaterials AscensionMaterial `json:"ascension_materials"`
}

type SkillTalent struct {
	Name             string                   `json:"name"`
	Unlock           string                   `json:"unlock"`
	Description      string                   `json:"description"`
	Upgrades         []Upgrade                `json:"upgrades"`
	Type             string                   `json:"type"`
	AttributeScaling []AttributeScalingStruct `json:"attribute-scaling"`
}

type Upgrade struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AttributeScalingStruct struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PassiveTalent struct {
	Name        string `json:"name"`
	Unlock      string `json:"unlock"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type Constellation struct {
	Name        string `json:"name"`
	Unlock      string `json:"unlock"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type AscensionMaterial struct {
	Level20 []Material `json:"level_20"`
	Level40 []Material `json:"level_40"`
	Level50 []Material `json:"level_50"`
	Level60 []Material `json:"level_60"`
	Level70 []Material `json:"level_70"`
	Level80 []Material `json:"level_80"`
}

type Material struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

//consumables

//to do

//Structur for Domains
type Domain struct {
	Name                string        `json:"name"`
	Type                string        `json:"type"`
	Description         string        `json:"description"`
	Location            string        `json:"location"`
	Nation              string        `json:"nation"`
	Requirements        []Requirement `json:"requirements"`
	RecommendedElements []string      `json:"recommendedElements"`
	Rewards             []DayReward   `json:"rewards"`
}

type Requirement struct {
	Level            int      `json:"level"`
	AdventureRank    int      `json:"adventureRank"`
	RecommendedLevel int      `json:"recommendedLevel"`
	LeyLineDisorder  []string `json:"leyLineDisorder"`
}

type DayReward struct {
	Day     string         `json:"day"`
	Details []RewardDetail `json:"details"`
}

type RewardDetail struct {
	Level                   int    `json:"level"`
	AdventureExperience     int    `json:"adventureExperience"`
	CompanionshipExperience int    `json:"companionshipExperience"`
	Mora                    int    `json:"mora"`
	Drops                   []Item `json:"drops"`
}

type Item struct {
	Name    string `json:"name"`
	DropMin int    `json:"drop_min"`
	DropMax int    `json:"drop_max"`
}

//Structur for element
type Element struct {
	Name      string     `json:"name"`
	Key       string     `json:"key"`
	Reactions []Reaction `json:"reactions"`
}

type Reaction struct {
	Name        string   `json:"name"`
	Elements    []string `json:"elements"`
	Description string   `json:"description"`
}

//Structur for enemies
type Enemies struct {
	Id                    string                       `json:"id"`
	Name                  string                       `json:"name"`
	Description           string                       `json:"description"`
	Region                string                       `json:"region"`
	Type                  string                       `json:"type"`
	Family                string                       `json:"family"`
	Faction               string                       `json:"faction"`
	Elements              []string                     `json:"elements"`
	Drops                 []EnemieDrop                 `json:"drops"`
	Artifacts             []EnemieArtifact             `json:"artifacts"`
	Descriptions          []EnemieDescription          `json:"descriptions"`
	ElementalDescriptions []EnemieElementalDescription `json:"elemental-description"`
	MoraGained            int                          `json:"mora-gained"`
}

type EnemieDrop struct {
	Name         string `json:"name"`
	Rarity       int    `json:"rarity"`
	MinimumLevel int    `json:"minimum-level"`
}

type EnemieArtifact struct {
	Name   string `json:"name"`
	Set    string `json:"set"`
	Rarity string `json:"rarity"`
}

type EnemieDescription struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EnemieElementalDescription struct {
	Element     string `json:"element"`
	Description string `json:"description"`
}
