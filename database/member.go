package saas_database

import (
	"github.com/loeffel-io/go-saas/model"
)

func (database *Database) GetMembers(users []*saas_model.User, teamId uint, userId uint) ([]*saas_model.User, error) {
	database.connection.LogMode(true)

	return users, database.GetConnection().
		Model(&saas_model.Team{
			Model:  saas_model.Model{Id: teamId},
			UserId: &userId,
		}).
		Related(&users, "Users").
		Error
}

func (database *Database) RemoveMember(user *saas_model.User, teamId uint, userId uint) error {
	return database.GetConnection().
		Model(&saas_model.User{
			Model: saas_model.Model{Id: user.GetId()},
		}).
		Association("Teams").
		Delete(&saas_model.Team{
			Model:  saas_model.Model{Id: teamId},
			UserId: &userId,
		}).
		Error
}
