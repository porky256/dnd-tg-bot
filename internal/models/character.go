package models

import "time"

type Character struct {
	ID                int
	Name              string
	OwnerTgID         string
	Level             int
	CurrentHP         int
	MaxHP             int
	Armor             int
	GoldCoins         int
	SilverCoins       int
	CopperCoins       int
	Abilities         Abilities
	SkillModificators Skills
	CreatedAt         time.Time
	UpdatedAt         time.Time
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
		Acrobatics:     calcModificatorFromSkill(c.Abilities.Dexterity) + c.SkillModificators.Acrobatics,
		AnimalHandling: calcModificatorFromSkill(c.Abilities.Wisdom) + c.SkillModificators.AnimalHandling,
		Arcana:         calcModificatorFromSkill(c.Abilities.Intelligence) + c.SkillModificators.Arcana,
		Athletics:      calcModificatorFromSkill(c.Abilities.Strength) + c.SkillModificators.Athletics,
		Deception:      calcModificatorFromSkill(c.Abilities.Charisma) + c.SkillModificators.Deception,
		History:        calcModificatorFromSkill(c.Abilities.Intelligence) + c.SkillModificators.History,
		Insight:        calcModificatorFromSkill(c.Abilities.Dexterity) + c.SkillModificators.Insight,
		Intimidation:   calcModificatorFromSkill(c.Abilities.Charisma) + c.SkillModificators.Intimidation,
		Investigation:  calcModificatorFromSkill(c.Abilities.Intelligence) + c.SkillModificators.Investigation,
		Medicine:       calcModificatorFromSkill(c.Abilities.Wisdom) + c.SkillModificators.Medicine,
		Nature:         calcModificatorFromSkill(c.Abilities.Intelligence) + c.SkillModificators.Nature,
		Perception:     calcModificatorFromSkill(c.Abilities.Wisdom) + c.SkillModificators.Perception,
		Performance:    calcModificatorFromSkill(c.Abilities.Charisma) + c.SkillModificators.Performance,
		Persuasion:     calcModificatorFromSkill(c.Abilities.Charisma) + c.SkillModificators.Persuasion,
		Religion:       calcModificatorFromSkill(c.Abilities.Intelligence) + c.SkillModificators.Religion,
		SleightOfHand:  calcModificatorFromSkill(c.Abilities.Dexterity) + c.SkillModificators.SleightOfHand,
		Stealth:        calcModificatorFromSkill(c.Abilities.Dexterity) + c.SkillModificators.Stealth,
		Survival:       calcModificatorFromSkill(c.Abilities.Wisdom) + c.SkillModificators.Survival,
	}
}
