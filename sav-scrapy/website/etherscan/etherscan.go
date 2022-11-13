package etherscan

func (client *EtherscanClient) ScrapyEth() {
	_, err := client.EtherscanRequest()
	if err != nil {
		return
	}
}
