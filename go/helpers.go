package main

import "errors"

// #cgo CFLAGS: -I${SRCDIR}/../lora_gateway/libloragw/inc
// #cgo LDFLAGS: -lm ${SRCDIR}/../lora_gateway/libloragw/libloragw.a
// #include "config.h"
// #include "loragw_hal.h"
// #include "loragw_gps.h"
// void setType(struct lgw_conf_rxrf_s *rxrfConf, enum lgw_radio_type_e val) {
// 	rxrfConf->type = val;
// }
import "C"

// initRadio initiates a radio configuration in the C.struct_lgw_conf_rxrf_s format, given
// the configuration for that radio.
func initRadio(radio RadioConf) (C.struct_lgw_conf_rxrf_s, error) {
	var cRadio = C.struct_lgw_conf_rxrf_s{
		enable:      C.bool(radio.Enabled),
		freq_hz:     C.uint32_t(radio.Freq),
		rssi_offset: C.float(radio.RssiOffset),
		tx_enable:   C.bool(radio.TxEnabled),
	}

	// Checking the radio is of a pre-defined type
	switch radio.RadioType {
	case "SX1257":
		C.setType(&cRadio, C.LGW_RADIO_TYPE_SX1257)
	case "SX1255":
		C.setType(&cRadio, C.LGW_RADIO_TYPE_SX1255)
	default:
		return cRadio, errors.New("Invalid radio type (should be SX1255 or SX1257)")
	}
	return cRadio, nil
}
