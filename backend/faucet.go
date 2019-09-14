package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/tendermint/tmlibs/bech32"
	"github.com/tomasen/realip"
)

type Data struct {
	FAUCET_CHAIN                string `json:"FAUCET_CHAIN"`
	FAUCET_RECAPTCHA_SECRET_KEY string `json:"FAUCET_RECAPTCHA_SECRET_KEY"`
	FAUCET_PUBLIC_URL           string `json:"FAUCET_PUBLIC_URL"`
	FAUCET_AMOUNT_FAUCET        string `json:"FAUCET_AMOUNT_FAUCET"`
	FAUCET_AMOUNT_STEAK         string `json:"FAUCET_AMOUNT_STEAK"`
	FAUCET_KEY                  string `json:"FAUCET_KEY"`
	FAUCET_PASS                 string `json:"FAUCET_PASS"`
	FAUCET_NODE                 string `json:"FAUCET_NODE"`
}

var chain string
var recaptchaSecretKey string
var amountFaucet string
var amountSteak string
var key string
var pass string
var node string
var publicUrl string

type claim_struct struct {
	Address  string
	Response string
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(key, "=", value)
		return value
	} else {
		log.Fatal("Error loading environment variable: ", key)
		return ""
	}
}

func main() {

	jsonFile, err := os.Open("env.local.json")
	// err := godotenv.Load(".env.local", ".env")
	if err != nil {
		log.Fatal("Error loading env.local.json file")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Data
	json.Unmarshal(byteValue, &data)

	chain = data.FAUCET_CHAIN
	recaptchaSecretKey = data.FAUCET_RECAPTCHA_SECRET_KEY
	amountFaucet = data.FAUCET_AMOUNT_STEAK
	amountSteak = data.FAUCET_AMOUNT_STEAK
	key = data.FAUCET_KEY
	pass = data.FAUCET_PASS
	node = data.FAUCET_NODE
	publicUrl = data.FAUCET_PUBLIC_URL

	r := mux.NewRouter()
	recaptcha.Init(recaptchaSecretKey)

	r.HandleFunc("/claim", getCoinsHandler)
	fmt.Println("faucet server started at 0.0.0.0:8080")

	log.Fatal(http.ListenAndServe(publicUrl, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}

func executeCmd(command string, writes ...string) {
	cmd, wc, _ := goExecute(command)

	for _, write := range writes {
		wc.Write([]byte(write + "\n"))
	}
	cmd.Wait()
}

func goExecute(command string) (cmd *exec.Cmd, pipeIn io.WriteCloser, pipeOut io.ReadCloser) {
	cmd = getCmd(command)
	pipeIn, _ = cmd.StdinPipe()
	pipeOut, _ = cmd.StdoutPipe()
	go cmd.Start()
	time.Sleep(time.Second)
	return cmd, pipeIn, pipeOut
}

func getCmd(command string) *exec.Cmd {
	// split command into command and args
	split := strings.Split(command, " ")

	var cmd *exec.Cmd
	if len(split) == 1 {
		cmd = exec.Command(split[0])
	} else {
		cmd = exec.Command(split[0], split[1:]...)
	}

	return cmd
}

func getCoinsHandler(w http.ResponseWriter, request *http.Request) {
	var claim claim_struct

	// decode JSON response from front end
	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&claim)
	if decoderErr != nil {
		panic(decoderErr)
	}

	// make sure address is bech32
	readableAddress, decodedAddress, decodeErr := bech32.DecodeAndConvert(claim.Address)
	if decodeErr != nil {
		panic(decodeErr)
	}
	// re-encode the address in bech32
	encodedAddress, encodeErr := bech32.ConvertAndEncode(readableAddress, decodedAddress)
	if encodeErr != nil {
		panic(encodeErr)
	}

	// make sure captcha is valid
	clientIP := realip.FromRequest(request)
	captchaResponse := claim.Response
	captchaPassed, captchaErr := recaptcha.Confirm(clientIP, captchaResponse)
	if captchaErr != nil {
		panic(captchaErr)
	}

	// send the coins!
	if captchaPassed {

		fmt.Println(encodedAddress)

		sendFaucet := fmt.Sprintf("colorcli tx send " + encodedAddress + " " + amountFaucet + " --from=" + key + " --chain-id=" + chain + " --fees=2color --home /home/ubuntu/goApps/src/github.com/RNSSolution/color-sdk/build/node1/colorcli")
		fmt.Println(sendFaucet)
		fmt.Println(time.Now().UTC().Format(time.RFC3339), encodedAddress, "[1]")
		executeCmd(sendFaucet, "y", pass)
	}
	return
}
