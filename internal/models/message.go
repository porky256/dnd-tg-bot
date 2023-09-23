package models

import (
	"fmt"
	"strconv"
)

var layout = "Name: <strong>%s</strong>\n" +
	"Lvl:%s\tHP:%s/%s\tArmor:%s\n" +
	"<pre>STR:%s(%s) DEX:%s(%s) CON:%s(%s)\n" +
	"INT:%s(%s) WIS:%s(%s) CHA:%s(%s)</pre>\n" +
	"Money: %sg %ss %sc"

func FormMessage(character *Character) string {
	return fmt.Sprintf(layout, character.Name,
		strconv.Itoa(character.Level), strconv.Itoa(character.CurrentHP),
		strconv.Itoa(character.MaxHP), strconv.Itoa(character.Armor),
		fillToSize(strconv.Itoa(calcModificatorFromSkill(character.Abilities.Strength)), 2),
		fillToSize(strconv.Itoa(character.Abilities.Strength), 2),
		fillToSize(strconv.Itoa(calcModificatorFromSkill(character.Abilities.Dexterity)), 2),
		fillToSize(strconv.Itoa(character.Abilities.Dexterity), 2),
		fillToSize(strconv.Itoa(calcModificatorFromSkill(character.Abilities.Constitution)), 2),
		fillToSize(strconv.Itoa(character.Abilities.Constitution), 2),
		fillToSize(strconv.Itoa(calcModificatorFromSkill(character.Abilities.Intelligence)), 2),
		fillToSize(strconv.Itoa(character.Abilities.Intelligence), 2),
		fillToSize(strconv.Itoa(calcModificatorFromSkill(character.Abilities.Wisdom)), 2),
		fillToSize(strconv.Itoa(character.Abilities.Wisdom), 2),
		fillToSize(strconv.Itoa(calcModificatorFromSkill(character.Abilities.Charisma)), 2),
		fillToSize(strconv.Itoa(character.Abilities.Charisma), 2),
		strconv.Itoa(character.GoldCoins), strconv.Itoa(character.SilverCoins), strconv.Itoa(character.CopperCoins))
}

func fillToSize(x string, length int) string {
	left := ""
	right := ""
	for i := 0; i < (length-len(x))/2; i++ {
		left += " "
		right += " "
	}
	if (length-len(x))%2 == 1 {
		left += " "
	}
	return left + x + right
}
