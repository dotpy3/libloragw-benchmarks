package main

type GainTableConf struct {
	PaGain      uint8  `json:"pa_gain"`
	MixGain     uint8  `json:"mix_gain"`
	RfPower     int8   `json:"rf_power"`
	DigGain     uint8  `json:"dig_gain"`
	Description string `json:"desc,omitempty"`
	DacGain     *uint8 `json:"dac_gain,omitempty"`
}

type RadioConf struct {
	Enabled     bool    `json:"enable"`
	RadioType   string  `json:"type"`
	Freq        int     `json:"freq"`
	RssiOffset  float32 `json:"rssi_offset"`
	TxEnabled   bool    `json:"tx_enable"`
	TxNotchFreq *int    `json:"tx_notch_freq,omitempty"`
	TxMinFreq   *int    `json:"tx_freq_min,omitempty"`
	TxMaxFreq   *int    `json:"tx_freq_max,omitempty"`
}

type ChannelConf struct {
	Enabled      bool    `json:"enable"`
	Description  string  `json:"desc,omitempty"`
	Radio        uint8   `json:"radio"`
	IfValue      int32   `json:"if"`
	Bandwidth    *uint32 `json:"bandwidth,omitempty"`
	Datarate     *uint32 `json:"datarate,omitempty"`
	SpreadFactor *uint8  `json:"spread_factor,omitempty"`
}
