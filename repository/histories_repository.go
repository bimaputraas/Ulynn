package repository

import (
	"errors"
	"time"
	"videogames_rent_api/model"
)

// create histories
func (r Histories) UpdateCreate(videoGameId int,durationMonth int,user model.Users) (*model.ResBodyHistoryVideoGame,error){
	tx := r.DB.Begin()

	// get video game
	var videoGame model.VideoGames
	result := tx.First(&videoGame,videoGameId)
	if result.Error != nil{
		return nil,result.Error
	}

	// check video game availability
	if !videoGame.Availability{
		return nil,errors.New("video game is not ready to rent")
	}

	// update video game avalability
	videoGame.Availability = false
	result = tx.Save(&videoGame)
	if result.Error != nil{
		tx.Rollback()
		return nil,result.Error
	}

	// get total rental cost
	// set default duration month to 1
	if durationMonth < 1 {
		durationMonth = 1
	}
	totalrentalcost := videoGame.RentalCost*float64(durationMonth)

	// check user deposit amount
	if user.DepositAmount < totalrentalcost{
		tx.Rollback()
		return nil,errors.New("insufficient funds")
	}

	// update user deposit amount
	user.DepositAmount -= totalrentalcost
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
		DueDate: time.Now().AddDate(0,durationMonth,0),
		Status: "In-Progress",
		TotalRentalCost: totalrentalcost,
	}
	result = tx.Create(&history)
	if result.Error != nil{
		tx.Rollback()
		return nil,result.Error
	}

	// commit transaction
	tx.Commit()

	// resBody
	var output model.ResBodyHistoryVideoGame
	result = r.DB.Raw(`SELECT video_games.title,histories.start_date,histories.due_date,histories.status,histories.total_rental_cost
	FROM histories
	JOIN video_games ON histories.video_game_id = video_games.id
	WHERE histories.id = ? AND histories.user_id = ?;`,history.ID,history.UserID).Scan(&output)
	if result.Error != nil{
		return nil,result.Error
	}

	return &output,nil
}

func (r Histories) Update(historyId,userId int) (*model.Histories,error){
	tx := r.DB.Begin()

	// find history
	var history model.Histories
	result := tx.Where("user_id = ?",userId).First(&history,historyId)
	if result.Error != nil {
		return nil,result.Error
	}

	if history.Status == "Done" || history.Status == "done" || history.Status == "DONE" {
		return nil, errors.New("failed update, video game is already available")
	}

	// update history
	history.Status = "Done"
	history.DueDate = time.Now()

	result = tx.Save(&history)
	if result.Error != nil {
		tx.Rollback()
		return nil,result.Error
	}

	// find video game
	var videoGame model.VideoGames
	result = tx.First(&videoGame,history.VideoGameID)
	if result.Error != nil {
		tx.Rollback()
		return nil,result.Error
	}

	// update video game
	videoGame.Availability = true
	result = tx.Save(&videoGame)
	if result.Error != nil {
		tx.Rollback()
		return nil,result.Error
	}

	// updated history
	var updatedHistory model.Histories
	result = tx.Where("user_id = ?",userId).Preload("VideoGame").First(&updatedHistory,historyId)
	if result.Error != nil {
		tx.Rollback()
		return nil,result.Error
	}

	// commit tx
	tx.Commit()
	return &updatedHistory,nil
}

// find all
func (r Histories) FindAll(userId int) (*[]model.Histories,error){
	var histories []model.Histories
	result := r.DB.Where("user_id = ?",userId).Preload("VideoGame").Find(&histories)
	if result.Error != nil {
		return nil,result.Error
	}

	return &histories,nil
}

// find all with status
func (r Histories) FindAllWithStatus(userId int,ok bool) (*[]model.Histories,error){
	status := "Done"
	if !ok {
		status = "In-Progress"
	}
	var histories []model.Histories
	result := r.DB.Where("status = ? AND user_id = ?",status,userId).Preload("VideoGame").Find(&histories)
	if result.Error != nil {
		return nil,result.Error
	}

	return &histories,nil
}

// find by id
func (r Histories) FindById(historyId,userId int) (*model.Histories,error){
	var history model.Histories
	result := r.DB.Where("user_id = ?",userId).Preload("VideoGame").First(&history,historyId)
	if result.Error != nil {
		return nil,result.Error
	}

	return &history,nil
}

func (r Histories) Delete(histories model.Histories) (*model.Histories,error){
	result := r.DB.Delete(&histories)
	if result.Error != nil {
		return nil,result.Error
	}

	return &histories,nil
}
