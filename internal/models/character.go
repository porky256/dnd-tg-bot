package models

import "time"

type Character struct {
	ID            int
	Name          string
	OwnerTgID     string
	Level         int
	CurrentHP     int
	MaxHP         int
	Armor         int
	GoldCoins     int
	SilverCoins   int
	CopperCoins   int
	Abilities     Abilities
	SkillInsights Skills
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Abilities struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

type Skills struct {
	Acrobatics     int
	AnimalHandling int
	Arcana         int
	Athletics      int
	Deception      int
	History        int
	Insight        int
	Intimidation   int
	Investigation  int
	Medicine       int
	Nature         int
	Perception     int
	Performance    int
	Persuasion     int
	Religion       int
	SleightOfHand  int
	Stealth        int
	Survival       int
}

func calcModificatorFromSkill(value int) int {
	return (value - 10) / 2
}

func (c Character) MasteryBonus() int {
	return (c.Level-1)/4 + 2
}

func (c Character) Skills() Skills {
	return Skills{
		Acrobatics:     c.Abilities.Dexterity + c.SkillInsights.Acrobatics*c.MasteryBonus(),
		AnimalHandling: c.Abilities.Wisdom + c.SkillInsights.AnimalHandling*c.MasteryBonus(),
		Arcana:         c.Abilities.Intelligence + c.SkillInsights.Arcana*c.MasteryBonus(),
		Athletics:      c.Abilities.Strength + c.SkillInsights.Athletics*c.MasteryBonus(),
		Deception:      c.Abilities.Charisma + c.SkillInsights.Deception*c.MasteryBonus(),
		History:        c.Abilities.Intelligence + c.SkillInsights.History*c.MasteryBonus(),
		Insight:        c.Abilities.Dexterity + c.SkillInsights.Insight*c.MasteryBonus(),
		Intimidation:   c.Abilities.Charisma + c.SkillInsights.Intimidation*c.MasteryBonus(),
		Investigation:  c.Abilities.Intelligence + c.SkillInsights.Investigation*c.MasteryBonus(),
		Medicine:       c.Abilities.Wisdom + c.SkillInsights.Medicine*c.MasteryBonus(),
		Nature:         c.Abilities.Intelligence + c.SkillInsights.Nature*c.MasteryBonus(),
		Perception:     c.Abilities.Wisdom + c.SkillInsights.Perception*c.MasteryBonus(),
		Performance:    c.Abilities.Charisma + c.SkillInsights.Performance*c.MasteryBonus(),
		Persuasion:     c.Abilities.Charisma + c.SkillInsights.Persuasion*c.MasteryBonus(),
		Religion:       c.Abilities.Intelligence + c.SkillInsights.Religion*c.MasteryBonus(),
		SleightOfHand:  c.Abilities.Dexterity + c.SkillInsights.SleightOfHand*c.MasteryBonus(),
		Stealth:        c.Abilities.Dexterity + c.SkillInsights.Stealth*c.MasteryBonus(),
		Survival:       c.Abilities.Wisdom + c.SkillInsights.Survival*c.MasteryBonus(),
	}
}
