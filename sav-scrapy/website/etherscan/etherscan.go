package etherscan

func ScrapyEth() {
	_, err := EtherscanRequest()
	if err != nil {
		return
	}
}
