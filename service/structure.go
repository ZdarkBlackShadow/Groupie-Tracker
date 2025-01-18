package service

//Structure for the artifacts

type ArtifactDetails struct {
	Name                 string `json:"name"`
	MaxRarity            int    `json:"max_rarity"`
	OnePieceBonus        string `json:"1-piece_bonus"`
	TwoPiecesBonus       string `json:"2-piece_bonus"`
	FourPiecesBonus      string `json:"4-piece_bonus"`
	Id                   string `json:"id"`
	ImageURL             string
	AllUrlImageAvailable []string
}

//Structur for the boss
type BossStruct struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Drops       []Drop     `json:"drops"`
	Artifacts   []Artifact `json:"artifacts"`
	Id          string     `json:"id"`
	ImageURL    string
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
	Id                 string            `json:"id"`
	ImageUrl           string
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
	Id                  string        `json:"id"`
	ImageUrl            string
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
	Id        string     `json:"id"`
	ImageUrl  string
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
	ImageUrl              string
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

//Structure for weapons
type Weapon struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Rarity             int    `json:"rarity"`
	BaseAttack         int    `json:"baseAttack"`
	SubStat            string `json:"subStat"`
	PassiveName        string `json:"passiveName"`
	PassiveDescription string `json:"passiveDesc"`
	Location           string `json:"location"`
	AscensionMaterial  string `json:"ascensionMaterial"`
	ImageUrl           string
	Id                 string
}

//Structure for materials :
//	-Boss material
type BossMaterial struct {
	Name       string   `json:"name"`
	Source     string   `json:"source"`
	Characters []string `json:"characters"`
}

//	-character-ascension
//a revoir
type CharactersAscensionMaterial struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
	Rarity  int      `json:"rarity"`
}

type CharactersAscensionElement struct {
	Sliver   Material `json:"sliver"`
	Fragment Material `json:"fragment"`
	Chunk    Material `json:"chunk"`
	Gemstone Material `json:"gemstone"`
}

type CharactersAscensionDataStructure struct {
	Anemo   Element `json:"anemo"`
	Cryo    Element `json:"cryo"`
	Electro Element `json:"electro"`
	Geo     Element `json:"geo"`
	Hydro   Element `json:"hydro"`
	Pyro    Element `json:"pyro"`
	Dendro  Element `json:"dendro"`
}

//	-Structure for character-experience
type CharacterExperienceItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"experience"`
	Rarity     int    `json:"rarity"`
}

type CharacterExperienceDataStruct struct {
	Items []CharacterExperienceItem `json:"items"`
	Id    string                    `json:"id"`
}

//	-Structure for common-ascension
type CommonAscensionDetails struct {
	Characters []string              `json:"characters"`
	Weapons    []string              `json:"weapons"`
	Items      []CommonAscensionItem `json:"items"`
	Sources    []string              `json:"sources"`
}

type CommonAscensionItem struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type CommonAscensionDataStruct struct {
	Slime                    CommonAscensionDetails `json:"slime"`
	HilichurlMasks           CommonAscensionDetails `json:"hilichurl-masks"`
	HilichurlArrowheads      CommonAscensionDetails `json:"hilichurl-arrowheads"`
	SamachurlScrolls         CommonAscensionDetails `json:"samachurl-scrolls"`
	TreasureHoarderInsignias CommonAscensionDetails `json:"treasure-hoarder-insignias"`
	FatuiInsignias           CommonAscensionDetails `json:"fatui-insignias"`
	WhopperflowerNectar      CommonAscensionDetails `json:"whopperflower-nectar"`
	HilichurlHorns           CommonAscensionDetails `json:"hilichurl-horns"`
	LeyLine                  CommonAscensionDetails `json:"ley-line"`
	Nobushi                  CommonAscensionDetails `json:"nobushi"`
	BoneShards               CommonAscensionDetails `json:"bone-shards"`
	MistGrass                CommonAscensionDetails `json:"mist-grass"`
	FatuiKnives              CommonAscensionDetails `json:"fatui-knives"`
	Specters                 CommonAscensionDetails `json:"specters"`
	ChaosParts               CommonAscensionDetails `json:"chaos-parts"`
	EremitesRed              CommonAscensionDetails `json:"eremites-red"`
	FungalsSpores            CommonAscensionDetails `json:"fungals-spores"`
	ClockworkMeka            CommonAscensionDetails `json:"clockwork-meka"`
	FontemerAberrant         CommonAscensionDetails `json:"fontemer-aberrant"`
	PrimalConstructPrisms    CommonAscensionDetails `json:"primal-construct-prisms"`
	BreacherPrimus           CommonAscensionDetails `json:"breacher-primus"`
	BlackSerpents            CommonAscensionDetails `json:"black-serpents"`
	HilichurlRogue           CommonAscensionDetails `json:"hilichurl-rogue"`
	ConsecratedBeast         CommonAscensionDetails `json:"consecrated-beast"`
	FatuiOperative           CommonAscensionDetails `json:"fatui-operative"`
	MirrorMaidenPrisms       CommonAscensionDetails `json:"mirror-maiden-prisms"`
	RuinSentinelsChaos       CommonAscensionDetails `json:"ruin-sentinels-chaos"`
	RiftwolfConcealds        CommonAscensionDetails `json:"riftwolf-concealds"`
	TaintedHydroPhantasm     CommonAscensionDetails `json:"tainted-hydro-phantasm"`
	StateShiftedFungus       CommonAscensionDetails `json:"state-shifted-fungus"`
	RuinDrakesChaos          CommonAscensionDetails `json:"ruin-drakes-chaos"`
	Hilt                     CommonAscensionDetails `json:"hilt"`
	Whistle                  CommonAscensionDetails `json:"whistle"`
	Fang                     CommonAscensionDetails `json:"fang"`
}

//	Structur for cooking-ingredients
type Ingredient struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Rarity      *int     `json:"rarity,omitempty"`
	Sources     []string `json:"sources"`
}

type CookingIngredients struct {
	Almond         Ingredient `json:"almond"`
	Bacon          Ingredient `json:"bacon"`
	BambooShoot    Ingredient `json:"bamboo-shoot"`
	Berry          Ingredient `json:"berry"`
	BirdEgg        Ingredient `json:"bird-egg"`
	Butter         Ingredient `json:"butter"`
	Cabbage        Ingredient `json:"cabbage"`
	CallaLily      Ingredient `json:"calla-lily"`
	Carrot         Ingredient `json:"carrot"`
	Cheese         Ingredient `json:"cheese"`
	ChilledMeat    Ingredient `json:"chilled-meat"`
	Crab           Ingredient `json:"crab"`
	CrabRoe        Ingredient `json:"crab-roe"`
	Cream          Ingredient `json:"cream"`
	Fish           Ingredient `json:"fish"`
	Flour          Ingredient `json:"flour"`
	Fowl           Ingredient `json:"fowl"`
	Ham            Ingredient `json:"ham"`
	Jam            Ingredient `json:"jam"`
	JueyunChili    Ingredient `json:"jueyun-chili"`
	LotusHead      Ingredient `json:"lotus-head"`
	Matsutake      Ingredient `json:"matsutake"`
	Milk           Ingredient `json:"milk"`
	Mint           Ingredient `json:"mint"`
	Mushroom       Ingredient `json:"mushroom"`
	Onion          Ingredient `json:"onion"`
	Pepper         Ingredient `json:"pepper"`
	Pinecone       Ingredient `json:"pinecone"`
	Potato         Ingredient `json:"potato"`
	Qingxin        Ingredient `json:"qingxin"`
	Radish         Ingredient `json:"radish"`
	RawMeat        Ingredient `json:"raw-meat"`
	Rice           Ingredient `json:"rice"`
	Salt           Ingredient `json:"salt"`
	Sausage        Ingredient `json:"sausage"`
	ShrimpMeat     Ingredient `json:"shrimp-meat"`
	SmallLampGrass Ingredient `json:"small-lamp-grass"`
	SmokedFowl     Ingredient `json:"smoked-fowl"`
	Snapdragon     Ingredient `json:"snapdragon"`
	Sugar          Ingredient `json:"sugar"`
	SweetFlower    Ingredient `json:"sweet-flower"`
	Tofu           Ingredient `json:"tofu"`
	Tomato         Ingredient `json:"tomato"`
	Violetgrass    Ingredient `json:"violetgrass"`
	Wheat          Ingredient `json:"wheat"`
	ID             string     `json:"id"`
}

//	-Structur for local-specialties
type Specialty struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type RegionSpecialties struct {
	Mondstadt []Specialty `json:"mondstadt"`
	Liyue     []Specialty `json:"liyue"`
	Inazuma   []Specialty `json:"inazuma"`
	Sumeru    []Specialty `json:"sumeru"`
	Fontaine  []Specialty `json:"fontaine"`
	Natlan    []Specialty `json:"natlan"`
	ID        string      `json:"id"`
}

//	-Structure for talent-book
type TalentCategory struct {
	Characters   []string     `json:"characters"`
	Availability []string     `json:"availability"`
	Source       string       `json:"source"`
	Items        []TalentItem `json:"items"`
}

type TalentItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type TalentBook struct {
	ID         string         `json:"id"`
	Freedom    TalentCategory `json:"freedom"`
	Resistance TalentCategory `json:"resistance"`
	Ballad     TalentCategory `json:"ballad"`
	Prosperity TalentCategory `json:"prosperity"`
	Diligence  TalentCategory `json:"diligence"`
	Gold       TalentCategory `json:"gold"`
	Transience TalentCategory `json:"transience"`
	Elegance   TalentCategory `json:"elegance"`
	Light      TalentCategory `json:"light"`
	Admonition TalentCategory `json:"admonition"`
	Ingenuity  TalentCategory `json:"ingenuity"`
	Praxis     TalentCategory `json:"praxis"`
	Equity     TalentCategory `json:"equity"`
	Justice    TalentCategory `json:"justice"`
	Order      TalentCategory `json:"order"`
	Conflict   TalentCategory `json:"conflict"`
	Contention TalentCategory `json:"contention"`
	Kindling   TalentCategory `json:"kindling"`
}

//	-Structur for talent-boss
type Talent struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type TalentBoss map[string]Talent

//	-Structur for weapon-ascension
type WeaponAscensionSet struct {
	Weapons      []string              `json:"weapons"`
	Availability []string              `json:"availability"`
	Source       string                `json:"source"`
	Items        []WeaponAscensionItem `json:"items"`
}

type WeaponAscensionItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

type WeaponAscensionData map[string]WeaponAscensionSet

//	-Structur for weapon-experience
type WeaponExperienceItem struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Experience int      `json:"experience"`
	Rarity     int      `json:"rarity"`
	Source     []string `json:"source"`
}

type WeaponExperienceData struct {
	Items []WeaponExperienceItem `json:"items"`
	Id    string                 `json:"id"`
}

//Structure for the food
type FoodStruct struct {
	Name        string   `json:"name"`
	Rarity      int      `json:"rarity"`
	Type        string   `json:"type"`
	Effect      string   `json:"effect"`
	HasRecipe   bool     `json:"hasRecipe"`
	Description string   `json:"description"`
	Recipe      []Recipe `json:"recipe"`
}

type Recipe struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

type Dishes map[string]FoodStruct

//Structure for the potions :
type CraftingItem struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

type Potion struct {
	Name     string         `json:"name"`
	Effect   string         `json:"effect"`
	Rarity   int            `json:"rarity"`
	Crafting []CraftingItem `json:"crafting"`
}

type PotionsData struct {
	ID      string            `json:"id"`
	Potions map[string]Potion `json:"-"`
}

type Register struct {
	Email        string
	Password     string
	Collecttions Collecttions
}

//Structure for the collections
type Collecttions struct {
	Collections []CollectionsStruct
}

type CollectionsStruct struct {
	Name               string
	Image              string
	Type               string
	LinkToTheRessource string
	DateAdded          string
}

type AllData struct {
	ALLArtifacts  []ArtifactDetails
	AllBoss       []BossStruct
	AllCharacters []Characters
	AllDomain     []Domain
	AllElement    []Element
	AllEnnemies   []Enemies
	AllWeapons    []Weapon
}
