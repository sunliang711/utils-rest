package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"nft-studio-backend/types"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// https://docs.etherscan.io/tutorials/verifying-contracts-programmatically
func VerifyEtherscan(network uint, apiKey string, module string, sourceCode string, contractAddress string, codeFormat string, contractName string, compilerVersion string,
	optimizationUsed uint, runs uint, constructorArguments string) error {

	var URL string

	data := url.Values{}
	data.Set("apikey", apiKey)
	data.Set("module", module)
	data.Set("action", "verifysourcecode")
	data.Set("sourceCode", sourceCode)
	data.Set("contractaddress", contractAddress)
	data.Set("codeformat", codeFormat)
	data.Set("contractname", contractName)
	data.Set("compilerversion", compilerVersion)
	data.Set("optimizationUsed", fmt.Sprintf("%v", optimizationUsed))
	if optimizationUsed == 1 {
		data.Set("runs", fmt.Sprintf("%v", runs))
	}
	if constructorArguments != "" {
		data.Set("constructorArguments", constructorArguments)
	}

	encodedData := data.Encode()

	switch network {
	case types.MainNetwork:
		URL = viper.GetString("contract.etherscan_main_url")
	case types.RopstenNetwork:
		URL = viper.GetString("contract.etherscan_ropsten_url")
	default:
		return errors.New("invalid network")
	}

	req, err := http.NewRequest("POST", URL, strings.NewReader(encodedData))
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("verify etherscan error: %v", err)
		return err
	}
	defer resp.Body.Close()
	io.Copy(os.Stderr, resp.Body)

	return nil
}
