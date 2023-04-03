package score

import (
	"fmt"
	"math"

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
	fmt.Println("collectionList Len", len(collectionList))
	for _, item := range collectionList {
		fmt.Println("collectionList item", item)

		// calc BlueChip
		bluceChip, err := calcBlueChip(item.TotalGiantWhaleHolder, item.TotalHolder)
		if err != nil {
			fmt.Println("calcBlueChip error", err.Error())
			continue
		}
		fmt.Println("bluceChip", bluceChip)

		// calc Fluidity

		// calc Reliability

		// calc CommunityActive

		// calc Heat

		// calc PotentialIncome
	}

	return nil
}

func calcBlueChip(totalGiantWhaleHolder uint64, totalHolder uint64) (string, error) {
	blueChip := (totalGiantWhaleHolder * 100 / totalHolder)
	// fmt.Println(totalGiantWhaleHolder, totalHolder, blueChip)
	blueChip = uint64(math.Ceil(float64(blueChip)))
	return fmt.Sprintf("%v", blueChip), nil
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
