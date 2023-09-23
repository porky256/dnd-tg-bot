package postgres

import (
	"context"
	"database/sql"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"github.com/porky256/dnd-tg-bot/internal/models"
	"time"
)

// PGCharacterProvider implements GlobalDatabaseProvider
type PGCharacterProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGCharacterProvider creates a new postgres userProvider entity
func NewPGCharacterProvider(db *sql.DB, timeout time.Duration) *PGCharacterProvider {
	return &PGCharacterProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}

func (db *PGCharacterProvider) InsertCharacter(char *models.Character) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()
	stmt := `INSERT INTO  characters (owner, name, char_level, current_hp, max_hp,armor,gold_coins,silver_coins,copper_coins) 
			 VALUES ($1, $2, $3, $4, $5,$6,$7,$8,$9) RETURNING id`

	var newID int
	err := db.DB.QueryRowContext(ctx, stmt,
		char.OwnerTgID,
		char.Name,
		char.Level,
		char.CurrentHP,
		char.MaxHP,
		char.Armor,
		char.GoldCoins,
		char.SilverCoins,
		char.CopperCoins,
	).Scan(&newID)

	if err != nil {
		return 0, ParsePqError(err)
	}
	return newID, nil
}

func (db *PGCharacterProvider) InsertCharactersAbilities(char *models.Character) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()
	stmt := `INSERT INTO  abilities (character_owner, strength, dexterity, 
                        constitution, intelligence,wisdom,charisma) 
			 VALUES ($1, $2, $3, $4, $5,$6,$7) RETURNING id`

	var newID int
	err := db.DB.QueryRowContext(ctx, stmt,
		char.ID,
		char.Abilities.Strength,
		char.Abilities.Dexterity,
		char.Abilities.Constitution,
		char.Abilities.Intelligence,
		char.Abilities.Wisdom,
		char.Abilities.Charisma,
	).Scan(&newID)

	if err != nil {
		return 0, ParsePqError(err)
	}
	return newID, nil
}

func (db *PGCharacterProvider) InsertCharactersSkillInsights(char *models.Character) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()
	stmt := `INSERT INTO  skills_insights (character_owner, acrobatics, animalHandling, 
                        arcana, athletics,deception,history,insight,intimidation,
					    investigation,medicine,nature,perception,performance,persuasion,
                    	religion,sleight_of_hand,stealth,survival) 
			 VALUES ($1, $2, $3, $4, $5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19) RETURNING id`

	var newID int
	err := db.DB.QueryRowContext(ctx, stmt,
		char.ID,
		char.SkillInsights.Acrobatics,
		char.SkillInsights.AnimalHandling,
		char.SkillInsights.Arcana,
		char.SkillInsights.Athletics,
		char.SkillInsights.Deception,
		char.SkillInsights.History,
		char.SkillInsights.Insight,
		char.SkillInsights.Intimidation,
		char.SkillInsights.Investigation,
		char.SkillInsights.Medicine,
		char.SkillInsights.Nature,
		char.SkillInsights.Perception,
		char.SkillInsights.Performance,
		char.SkillInsights.Persuasion,
		char.SkillInsights.Religion,
		char.SkillInsights.SleightOfHand,
		char.SkillInsights.Stealth,
		char.SkillInsights.Survival,
	).Scan(&newID)

	if err != nil {
		return 0, ParsePqError(err)
	}
	return newID, nil
}

func (db *PGCharacterProvider) UpdateCharacter(char *models.Character) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE characters SET name = $1, char_level = $2, current_hp = $3, max_hp = $4, 
                    armor = $5, gold_coins = $6, silver_coins = $7, copper_coins = $8 WHERE id=$9`
	res, err := db.DB.ExecContext(ctx, stmt,
		char.Name,
		char.Level,
		char.CurrentHP,
		char.MaxHP,
		char.Armor,
		char.GoldCoins,
		char.SilverCoins,
		char.CopperCoins,
		char.ID,
	)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

func (db *PGCharacterProvider) UpdateCharactersAbilities(char *models.Character) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE abilities SET strength=$1, dexterity=$2, constitution=$3, 
                    intelligence=$4 ,wisdom=$5 ,charisma=$6 WHERE character_owner=$7`
	res, err := db.DB.ExecContext(ctx, stmt,
		char.Abilities.Strength,
		char.Abilities.Dexterity,
		char.Abilities.Constitution,
		char.Abilities.Intelligence,
		char.Abilities.Wisdom,
		char.Abilities.Charisma,
		char.ID,
	)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

func (db *PGCharacterProvider) UpdateCharactersSkillInsights(char *models.Character) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE skills_insights SET acrobatics=$1, animalHandling=$2, 
                    arcana=$3, athletics=$4, deception=$5, history=$6,
                    insight=$7, intimidation=$8, investigation=$9,
                    medicine=$10, nature=$11, perception=$12, performance=$13,
                    persuasion=$14, religion=$15, sleight_of_hand=$16,
                    stealth=$17, survival=$18 WHERE character_owner=$19 `
	res, err := db.DB.ExecContext(ctx, stmt,
		char.SkillInsights.Acrobatics,
		char.SkillInsights.AnimalHandling,
		char.SkillInsights.Arcana,
		char.SkillInsights.Athletics,
		char.SkillInsights.Deception,
		char.SkillInsights.History,
		char.SkillInsights.Insight,
		char.SkillInsights.Intimidation,
		char.SkillInsights.Investigation,
		char.SkillInsights.Medicine,
		char.SkillInsights.Nature,
		char.SkillInsights.Perception,
		char.SkillInsights.Performance,
		char.SkillInsights.Persuasion,
		char.SkillInsights.Religion,
		char.SkillInsights.SleightOfHand,
		char.SkillInsights.Stealth,
		char.SkillInsights.Survival,
		char.ID,
	)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

func (db *PGCharacterProvider) GetFullCharacterByID(charID int) (*models.Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT characters.id, owner, name, char_level, current_hp, max_hp,armor,
                    gold_coins,silver_coins,copper_coins, characters.created_at, characters.updated_at,
                    strength, dexterity, constitution, intelligence,wisdom,charisma,
                    acrobatics, animalHandling, arcana, athletics,deception,history,
                    insight,intimidation, investigation,medicine,nature,perception,
                    performance,persuasion, religion,sleight_of_hand,stealth,survival
                    FROM characters 
                        LEFT JOIN abilities ON characters.id = abilities.character_owner
                        LEFT JOIN skills_insights ON characters.id = skills_insights.character_owner
					WHERE characters.id = $1`
	char := new(models.Character)
	row := db.DB.QueryRowContext(ctx, stmt, charID)
	err := row.Scan(
		&char.ID,
		&char.OwnerTgID,
		&char.Name,
		&char.Level,
		&char.CurrentHP,
		&char.MaxHP,
		&char.Armor,
		&char.GoldCoins,
		&char.SilverCoins,
		&char.CopperCoins,
		&char.CreatedAt,
		&char.UpdatedAt,
		&char.Abilities.Strength,
		&char.Abilities.Dexterity,
		&char.Abilities.Constitution,
		&char.Abilities.Intelligence,
		&char.Abilities.Wisdom,
		&char.Abilities.Charisma,
		&char.SkillInsights.Acrobatics,
		&char.SkillInsights.AnimalHandling,
		&char.SkillInsights.Arcana,
		&char.SkillInsights.Athletics,
		&char.SkillInsights.Deception,
		&char.SkillInsights.History,
		&char.SkillInsights.Insight,
		&char.SkillInsights.Intimidation,
		&char.SkillInsights.Investigation,
		&char.SkillInsights.Medicine,
		&char.SkillInsights.Nature,
		&char.SkillInsights.Perception,
		&char.SkillInsights.Performance,
		&char.SkillInsights.Persuasion,
		&char.SkillInsights.Religion,
		&char.SkillInsights.SleightOfHand,
		&char.SkillInsights.Stealth,
		&char.SkillInsights.Survival,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return char, nil
}

func (db *PGCharacterProvider) GetCharacterByID(charID int) (*models.Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, owner, name, char_level, current_hp, max_hp,armor,
                    gold_coins,silver_coins,copper_coins, created_at, updated_at
                    FROM characters WHERE id = $1`
	char := new(models.Character)
	row := db.DB.QueryRowContext(ctx, stmt, charID)
	err := row.Scan(
		&char.ID,
		&char.OwnerTgID,
		&char.Name,
		&char.Level,
		&char.CurrentHP,
		&char.MaxHP,
		&char.Armor,
		&char.GoldCoins,
		&char.SilverCoins,
		&char.CopperCoins,
		&char.CreatedAt,
		&char.UpdatedAt,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return char, nil
}

func (db *PGCharacterProvider) GetCharactersByOwnerTgID(ownerTgID string) ([]models.Character, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, owner, name, char_level, current_hp, max_hp,armor,
                    gold_coins,silver_coins,copper_coins, created_at, updated_at
                    FROM characters WHERE owner = $1`
	rows, err := db.DB.QueryContext(ctx, stmt, ownerTgID)
	if err != nil {
		return nil, ParsePqError(err)
	}
	var chars []models.Character
	for rows.Next() {
		var char models.Character
		err = rows.Scan(
			&char.ID,
			&char.OwnerTgID,
			&char.Name,
			&char.Level,
			&char.CurrentHP,
			&char.MaxHP,
			&char.Armor,
			&char.GoldCoins,
			&char.SilverCoins,
			&char.CopperCoins,
			&char.CreatedAt,
			&char.UpdatedAt,
		)
		if err != nil {
			return nil, ParsePqError(err)
		}
		chars = append(chars, char)
	}

	return chars, nil
}

func (db *PGCharacterProvider) GetCharactersAbilitiesByCharacterID(charID int) (*models.Abilities, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT strength, dexterity, constitution, intelligence,wisdom,charisma
                    FROM abilities WHERE character_owner = $1`
	abilities := new(models.Abilities)
	row := db.DB.QueryRowContext(ctx, stmt, charID)
	err := row.Scan(
		&abilities.Strength,
		&abilities.Dexterity,
		&abilities.Constitution,
		&abilities.Intelligence,
		&abilities.Wisdom,
		&abilities.Charisma,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return abilities, nil
}

func (db *PGCharacterProvider) GetCharactersSkillsInsightsByCharacterID(charID int) (*models.Skills, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT acrobatics, animalHandling, arcana, athletics,deception,history,
                    insight,intimidation, investigation,medicine,nature,perception,
                    performance,persuasion, religion,sleight_of_hand,stealth,survival
                    FROM skills_insights WHERE character_owner = $1`
	skills := new(models.Skills)
	row := db.DB.QueryRowContext(ctx, stmt, charID)
	err := row.Scan(
		&skills.Acrobatics,
		&skills.AnimalHandling,
		&skills.Arcana,
		&skills.Athletics,
		&skills.Deception,
		&skills.History,
		&skills.Insight,
		&skills.Intimidation,
		&skills.Investigation,
		&skills.Medicine,
		&skills.Nature,
		&skills.Perception,
		&skills.Performance,
		&skills.Persuasion,
		&skills.Religion,
		&skills.SleightOfHand,
		&skills.Stealth,
		&skills.Survival,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return skills, nil
}

func (db *PGCharacterProvider) DeleteCharacterByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM characters WHERE id=$1", id)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}

func (db *PGCharacterProvider) DeleteAbilityByCharacterID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM abilities WHERE character_owner=$1", id)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}

func (db *PGCharacterProvider) DeleteSkillInsightsByCharacterID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM skills_insights WHERE character_owner=$1", id)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
