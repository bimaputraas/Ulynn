package repository

import (
	"errors"
	"time"
	"videogames_rent_api/model"
)

// create histories
func (r Rent) UpdateCreate(videoGameId int,user model.Users) (*model.ResBodyHistoryVideoGame,error){
	tx := r.DB.Begin()

	// check video game availability
	var videoGame model.VideoGames
	result := tx.First(&videoGame,videoGameId)
	if result.Error != nil{
		tx.Rollback()
		return nil,result.Error
	}
	if !videoGame.Availability{
		tx.Rollback()
		return nil,errors.New("video game is not available")
	}

	// check user deposit amount
	if user.DepositAmount < videoGame.RentalCost{
		tx.Rollback()
		return nil,errors.New("insufficient funds")
	}

	// update video game avalability
	videoGame.Availability = false
	result = tx.Save(&videoGame)
	if result.Error != nil{
		tx.Rollback()
		return nil,result.Error
	}

	// update user deposit amount
	user.DepositAmount -= videoGame.RentalCost
	result = tx.Save(&user)
	if result.Error != nil{
		tx.Rollback()
		return nil,result.Error
	}

	// create history
	history := model.Histories{
		UserID: user.ID,
		VideoGameID: videoGame.ID,
		StartDate: time.Now(),
		DueDate: time.Now().AddDate(0,0,3),
		Status: "In-Progress",
	}
	result = tx.Save(&history)
	if result.Error != nil{
		tx.Rollback()
		return nil,result.Error
	}

	// commit transaction
	tx.Commit()

	// resBody
	var output model.ResBodyHistoryVideoGame
	result = r.DB.Raw(`SELECT video_games.title,histories.start_date,histories.due_date,histories.status
	FROM histories
	JOIN video_games ON histories.video_game_id = video_games.id
	WHERE histories.id = ? AND histories.user_id = ?;`,history.ID,history.UserID).Scan(&output)
	if result.Error != nil{
		return nil,result.Error
	}

	return &output,nil
}

func (r Rent) Update(historyId,userId int,reqStatus string) (*model.Histories,error){
	var history model.Histories
	result := r.DB.Where("user_id = ?",userId).First(&history,historyId)
	if result.Error != nil {
		return nil,result.Error
	}

	if history.Status == "done" || history.Status == "DONE"{
		return nil, errors.New("failed update, video game is already available")
	}

	// update history
	history.Status = reqStatus
	history.DueDate = time.Now()

	result = r.DB.Save(&history)
	if result.Error != nil {
		return nil,result.Error
	}

	// find video game
	var videoGame model.VideoGames
	result = r.DB.First(&videoGame,history.VideoGameID)
	if result.Error != nil {
		return nil,result.Error
	}

	// update video game
	videoGame.Availability = true
	result = r.DB.Save(&videoGame)
	if result.Error != nil {
		return nil,result.Error
	}

	return &history,nil
}