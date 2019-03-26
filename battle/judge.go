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
	"strings"
)

func Judge(ATKSrc, DEFSrc *models.SourceCode, game *models.Game, battle *models.Battle) {
	//todo:check pointers
	log.Println("Start judging...")
	atkUUID, err := uuid.NewRandom()
	if err != nil{
		log.Println(err)
		// todo:handle
	}
	atkUUIDStr := strings.Replace(atkUUID.String(), "-", "_", -1)

	defUUID, err := uuid.NewRandom()
	if err != nil{
		log.Println(err)
		// todo
	}
	defUUIDStr := strings.Replace(defUUID.String(), "-", "_", -1)

	resUUID, err := uuid.NewRandom()
	if err != nil{
		log.Println(err)
		// todo
	}
	resUUIDStr := strings.Replace(resUUID.String(), "-", "_", -1)

	atkFilePath := filepath.Join(conf.DataDir, atkUUIDStr)+".cpp"
	err = ioutil.WriteFile(atkFilePath, []byte(ATKSrc.Content), 0666)
	if err != nil{
		log.Println(err)
	}
	defFilePath := filepath.Join(conf.DataDir, defUUIDStr)+".cpp"
	err = ioutil.WriteFile(defFilePath, []byte(DEFSrc.Content), 0666)
	if err != nil{
		log.Println(err)
	}
	resFilePath := filepath.Join(conf.DataDir, resUUIDStr)
	err = ioutil.WriteFile(resFilePath, []byte(""), 0666)
	if err != nil{
		log.Println(err)
		return
	}

	cmd := exec.Command(conf.JudgerPath, resFilePath, atkFilePath, atkUUIDStr, defFilePath, defUUIDStr)

	if err:=models.UpdateBattleStatusByID(battle.ID, models.Judeging); err != nil{
		log.Println(err)
	}
	err = cmd.Run()
	if err != nil{
		log.Println(err)
	}else {
		b, err := ioutil.ReadFile(resFilePath)
		if err != nil{
			log.Println(err)
			log.Println("Read result file failed!")
			// todo:handle unexpect error
		}else{
			log.Println("End judging...")
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

	models.UpdateBattle(bt)
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
