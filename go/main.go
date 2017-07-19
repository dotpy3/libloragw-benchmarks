package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
)

// #cgo CFLAGS: -I${SRCDIR}/../lora_gateway/libloragw/inc
// #cgo LDFLAGS: -lm ${SRCDIR}/../lora_gateway/libloragw/libloragw.a -lmpsse -lrt
// #include "config.h"
// #include "loragw_hal.h"
// #include "loragw_gps.h"
import "C"

const NbMaxPackets = 8

func main() {
	// Configuration
	nbCycles := 1000
	nbUplinkCycle := 10
	nbDownlinkCycle := 2

	args := os.Args
	var err error
	if len(args) >= 2 {
		nbCycles, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Number of cycles:", nbCycles)
	if len(args) >= 3 {
		nbUplinkCycle, err = strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Number of uplinks per cycles:", nbUplinkCycle)
	if len(args) >= 4 {
		nbUplinkCycle, err = strconv.Atoi(args[3])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Number of downlinks per cycle:", nbDownlinkCycle)

	// Setting board
	var boardConf = C.struct_lgw_conf_board_s{
		clksrc:         C.uint8_t(0),
		lorawan_public: C.bool(true),
	}
	fmt.Println("Board configuration:")
	pp.Println(boardConf)
	C.lgw_board_setconf(boardConf)

	// Setting TX Gain
	var gainLut = C.struct_lgw_tx_gain_lut_s{
		lut: [C.TX_GAIN_LUT_SIZE_MAX]C.struct_lgw_tx_gain_s{},
	}
	txLuts := []GainTableConf{
		{Description: "TX gain table, index 0", PaGain: 0, MixGain: 8, RfPower: -6, DigGain: 0},
		{Description: "TX gain table, index 1", PaGain: 0, MixGain: 10, RfPower: -3, DigGain: 0},
		{Description: "TX gain table, index 2", PaGain: 0, MixGain: 12, RfPower: 0, DigGain: 0},
		{Description: "TX gain table, index 3", PaGain: 1, MixGain: 8, RfPower: 3, DigGain: 0},
		{Description: "TX gain table, index 4", PaGain: 1, MixGain: 10, RfPower: 6, DigGain: 0},
		{Description: "TX gain table, index 5", PaGain: 1, MixGain: 12, RfPower: 10, DigGain: 0},
		{Description: "TX gain table, index 6", PaGain: 1, MixGain: 13, RfPower: 11, DigGain: 0},
		{Description: "TX gain table, index 7", PaGain: 2, MixGain: 9, RfPower: 12, DigGain: 0},
		{Description: "TX gain table, index 8", PaGain: 1, MixGain: 15, RfPower: 13, DigGain: 0},
		{Description: "TX gain table, index 9", PaGain: 2, MixGain: 10, RfPower: 14, DigGain: 0},
		{Description: "TX gain table, index 10", PaGain: 2, MixGain: 11, RfPower: 16, DigGain: 0},
		{Description: "TX gain table, index 11", PaGain: 3, MixGain: 9, RfPower: 20, DigGain: 0},
		{Description: "TX gain table, index 12", PaGain: 3, MixGain: 10, RfPower: 23, DigGain: 0},
		{Description: "TX gain table, index 13", PaGain: 3, MixGain: 11, RfPower: 25, DigGain: 0},
		{Description: "TX gain table, index 14", PaGain: 3, MixGain: 12, RfPower: 26, DigGain: 0},
		{Description: "TX gain table, index 15", PaGain: 3, MixGain: 14, RfPower: 27, DigGain: 0},
	}
	for i, txConf := range txLuts {
		txLut := C.struct_lgw_tx_gain_s{}
		if txConf.DacGain != nil {
			txLut.dac_gain = C.uint8_t(*txConf.DacGain)
		} else {
			txLut.dac_gain = 3
		}
		txLut.dig_gain = C.uint8_t(txConf.DigGain)
		txLut.mix_gain = C.uint8_t(txConf.MixGain)
		txLut.rf_power = C.int8_t(txConf.RfPower)
		txLut.pa_gain = C.uint8_t(txConf.PaGain)
		gainLut.lut[i] = txLut
	}
	gainLut.size = C.uint8_t(len(txLuts))

	fmt.Println("TX gain tables:")
	pp.Println(gainLut)
	C.lgw_txgain_setconf(&gainLut)

	int1 := 863000000
	int2 := 870000000

	// Setting RF channels
	radios := []RadioConf{
		{Enabled: true, RadioType: "SX1257", Freq: 867500000, RssiOffset: -166, TxEnabled: true, TxMinFreq: &int1, TxMaxFreq: &int2},
		{Enabled: true, RadioType: "SX1257", Freq: 868500000, RssiOffset: -166, TxEnabled: false},
	}
	for nbRadio, radio := range radios {
		cRadio, err := initRadio(radio)
		if err != nil {
			continue
		}
		fmt.Println("Radio", index, ":")
		pp.Println(cRadio)
		C.lgw_rxrf_setconf(C.uint8_t(nbRadio), cRadio)
	}

	// Setting SF channels
	channels := []ChannelConf{
		{Description: "Lora MAC, 125kHz, all SF, 868.1 MHz", Enabled: true, Radio: 1, IfValue: -400000},
		{Description: "Lora MAC, 125kHz, all SF, 868.3 MHz", Enabled: true, Radio: 1, IfValue: -200000},
		{Description: "Lora MAC, 125kHz, all SF, 868.5 MHz", Enabled: true, Radio: 1, IfValue: 0},
		{Description: "Lora MAC, 125kHz, all SF, 867.1 MHz", Enabled: true, Radio: 0, IfValue: -400000},
		{Description: "Lora MAC, 125kHz, all SF, 867.3 MHz", Enabled: true, Radio: 0, IfValue: -200000},
		{Description: "Lora MAC, 125kHz, all SF, 867.5 MHz", Enabled: true, Radio: 0, IfValue: 0},
		{Description: "Lora MAC, 125kHz, all SF, 867.7 MHz", Enabled: true, Radio: 0, IfValue: 200000},
		{Description: "Lora MAC, 125kHz, all SF, 867.9 MHz", Enabled: true, Radio: 0, IfValue: 400000},
	}

	for index, channel := range channels {
		if !channel.Enabled {
			continue
		}
		var cChannel = C.struct_lgw_conf_rxif_s{
			enable:   C.bool(channel.Enabled),
			rf_chain: C.uint8_t(channel.Radio),
			freq_hz:  C.int32_t(channel.IfValue),
		}
		fmt.Println("Channel", index, ":")
		pp.Println(cChannel)
		C.lgw_rxif_setconf(C.uint8_t(index), cChannel)
	}

	// Setting LoRa channel
	LoRaChannel := C.struct_lgw_conf_rxif_s{
		enable:    true,
		rf_chain:  1,
		freq_hz:   -200000,
		bandwidth: C.BW_250KHZ,
		datarate:  C.DR_UNDEFINED,
	}
	fmt.Println("LoRa channel:")
	pp.Println(LoRaChannel)
	C.lgw_rxif_setconf(C.uint8_t(8), LoRaChannel)

	// Setting FSK channel
	fsk := C.struct_lgw_conf_rxif_s{
		enable:    true,
		rf_chain:  1,
		freq_hz:   30000,
		bandwidth: C.BW_125KHZ,
		datarate:  50000,
	}
	fmt.Println("FSK channel:")
	pp.Println(fsk)
	C.lgw_rxif_setconf(C.uint8_t(9), fsk)

	// Start
	fmt.Println("Starting concentrator:")
	if C.lgw_start() != C.LGW_HAL_SUCCESS {
		fmt.Println("Concentrator start unsuccessful")
		return
	}

	for i := 0; i < nbCycles; i++ {
		for j := 0; j < nbUplinkCycle; j++ {
			var packets [NbMaxPackets]C.struct_lgw_pkt_rx_s
			C.lgw_receive(NbMaxPackets, &packets[0])
		}
		// No downlinks for the moment
	}

	// Stop
	C.lgw_stop()
}
