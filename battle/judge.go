package battle

import (
	"github.com/google/uuid"
	"github.com/jackmrzhou/gc-ai-backend/conf"
	"github.com/jackmrzhou/gc-ai-backend/models"
	"github.com/jackmrzhou/gc-ai-backend/utils"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

func Judge(ATKSrc, DEFSrc *models.SourceCode, game *models.Game, battle *models.Battle) {
	//todo:check pointers
	atkUUID, err := uuid.NewRandom()
	if err != nil{
		// todo:handle
	}
	defUUID, err := uuid.NewRandom()
	if err != nil{
		// todo
	}
	resUUID, err := uuid.NewRandom()
	if err != nil{
		// todo
	}

	atkFilePath := filepath.Join(conf.DataDir, atkUUID.String())
	err = ioutil.WriteFile(atkFilePath, []byte(ATKSrc.Content), 0666)
	if err != nil{
		log.Println(err)
	}
	defFilePath := filepath.Join(conf.DataDir, defUUID.String())
	err = ioutil.WriteFile(defFilePath, []byte(DEFSrc.Content), 0666)
	if err != nil{
		log.Println(err)
	}
	resFilePath := filepath.Join(conf.DataDir, resUUID.String())
	err = ioutil.WriteFile(resFilePath, []byte(""), 0666)
	if err != nil{
		log.Println(err)
		return
	}

	cmd := exec.Command(conf.JudgerPath, atkFilePath, atkUUID.String(), defFilePath, defUUID.String(), resFilePath)

	if err:=models.UpdateBattleStatusByID(battle.ID, models.Judeging); err != nil{
		log.Println(err)
	}
	err = cmd.Run()
	if err != nil{
		log.Println(err)
	}else {
		b, err := ioutil.ReadFile(resFilePath)
		if err != nil{
			log.Println("Read result file failed!")
			// todo:handle unexpect error
		}else{
			UpdateBattleInfo(battle, string(b))

			// todo:error handle
		}
	}
}

const (
	attacker = 1
	defender = 2
)

type Result struct {
	Winner uint8
}

func ParseResult(res string) (*Result, error) {
	ret := Result{}

	ch := res[len(res) - 1]
	if ch != 'A' && ch != 'B' && ch != 'T'{
		ret.Winner = 0
	}else if ch == 'A'{
		ret.Winner = attacker
	}else{
		ret.Winner = defender
	}
	return &ret, nil
}

func UpdateBattleInfo(bt *models.Battle, result string) error {
	bt.Status = models.Finished
	bt.Detail = string(result)
	_res, err := ParseResult(result)
	if _res.Winner == attacker{
		bt.WinnerID = bt.AttackerID
	}else if _res.Winner == defender{
		bt.WinnerID = bt.DefenderID
	}else{
		bt.WinnerID = 0
	}

	UpdateScore(bt)
	// todo:handle update error

	return err
}

func UpdateScore(bt *models.Battle) {

	if bt.WinnerID == 0{
		// tie
		bt.RewardScore = 0
		bt.PenaltyScore = 0
		return
	}

	bt.RewardScore = 5
	bt.PenaltyScore = 5

	// update winner's score
	if winnerRank, err := models.QueryRankByUserAndGameID(bt.WinnerID, bt.GameID); err != nil{
		winnerRank, err = models.CreateRankByID(bt.WinnerID, bt.GameID, bt.RewardScore)
		utils.LogIfNotNil(err)
	}else{
		err = models.UpdateRankByID(winnerRank.ID, int(bt.RewardScore))
		utils.LogIfNotNil(err)
	}

	var loserID uint
	if bt.WinnerID == bt.AttackerID{
		loserID = bt.DefenderID
	}else{
		loserID = bt.AttackerID
	}

	//update loser's score
	if loserRank, err := models.QueryRankByUserAndGameID(loserID, bt.GameID); err != nil{
		loserRank, err = models.CreateRankByID(loserID, bt.GameID, 0)
		utils.LogIfNotNil(err)
	}else{
		err = models.UpdateRankByID(loserRank.ID, -int(bt.PenaltyScore))
		utils.LogIfNotNil(err)
	}
}
