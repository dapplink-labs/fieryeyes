package score

import (
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/models"
)

func (as *FeScoreService) CalcScores() error {
	log.Info("begin to calc scores")
	// get all collection infos from db
	collectionModel := models.Collection{}
	collectionList, err := collectionModel.GetCollectionList(1, 100000, 1, as.Cfg.Database.Db)
	if err != nil {
		log.Error("GetCollectionList error", err.Error())
		return err
	}
	fmt.Println("collectionList", collectionList)

	// calc BlueChip

	// calc Fluidity

	// calc Reliability

	// calc CommunityActive

	// calc Heat

	// calc PotentialIncome

	return nil
}

func calcBlueChip() (string, error) {
	return "", nil
}

func calcFluidity() (string, error) {
	return "", nil
}

func calcReliability() (string, error) {
	return "", nil
}

func calcCommunityActive() (string, error) {
	return "", nil
}

func calcPotentialIncome() (string, error) {
	return "", nil
}
