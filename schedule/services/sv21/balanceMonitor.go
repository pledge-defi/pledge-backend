package sv21

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"math/big"
	"pledge-backend/config"
	"pledge-backend/log"
	"pledge-backend/utils"
)

type BalanceMonitor struct {
}

func NewBalanceMonitor() *BalanceMonitor {
	return &BalanceMonitor{}
}

// Monitor Sending email when balance is insufficient
func (s *BalanceMonitor) Monitor() {

	//check on bsc test-net
	tokenPoolBalance, err := s.GetBalance(config.Config.TestNet.NetUrl, config.Config.TestNet.PledgePoolToken)
	thresholdPoolToken, ok := new(big.Int).SetString(config.Config.Threshold.PledgePoolTokenThresholdBnb, 10)
	if ok && (err == nil) && (tokenPoolBalance.Cmp(thresholdPoolToken) <= 0) {
		emailBody, err := s.EmailBody(config.Config.TestNet.PledgePoolToken, "TBNB", tokenPoolBalance.String(), thresholdPoolToken.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}

	//check on bsc main-net
	tokenPoolBalance, err = s.GetBalance(config.Config.MainNet.NetUrl, config.Config.MainNet.PledgePoolToken)
	thresholdPoolToken, ok = new(big.Int).SetString(config.Config.Threshold.PledgePoolTokenThresholdBnb, 10)
	if ok && (err == nil) && (tokenPoolBalance.Cmp(thresholdPoolToken) <= 0) {
		emailBody, err := s.EmailBody(config.Config.MainNet.PledgePoolToken, "BNB", tokenPoolBalance.String(), thresholdPoolToken.String())
		if err != nil {
			log.Logger.Error(err.Error())
		} else {
			err = utils.SendEmail(emailBody, 2)
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
	}
}

// GetBalance get balance of ERC20 token
func (s *BalanceMonitor) GetBalance(netUrl, token string) (*big.Int, error) {

	ethereumClient, err := ethclient.Dial(netUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}
	defer ethereumClient.Close()

	balance, err := ethereumClient.BalanceAt(context.Background(), common.HexToAddress(token), nil)
	if err != nil {
		log.Logger.Error(err.Error())
		return big.NewInt(0), err
	}

	return balance, err
}

// EmailBody email body
func (s *BalanceMonitor) EmailBody(token, currency, balance, threshold string) ([]byte, error) {
	e18, err := decimal.NewFromString("1000000000000000000")
	if err != nil {
		return nil, err
	}

	balanceDeci, err := decimal.NewFromString(balance)
	if err != nil {
		return nil, err
	}
	balanceStr := balanceDeci.Div(e18).String()

	thresholdDeci, err := decimal.NewFromString(threshold)
	if err != nil {
		return nil, err
	}

	thresholdStr := thresholdDeci.Div(e18).String()
	log.Logger.Sugar().Info("balance not enough ", token, " ", currency, " ", balanceStr, " ", thresholdStr)
	body := fmt.Sprintf(`<p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;The balance of <strong><span style="color: rgb(255, 0, 0);"> %s </span></strong> is <strong>%s %s </strong>. Please recharge it in time. The current minimum balance limit is %s %s 
</p>`, token, balanceStr, currency, thresholdStr, currency)
	return []byte(body), nil
}
