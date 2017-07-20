#include <stdio.h>
#include <time.h>
#include "config.h"
#include "loragw_hal.h"
#include "loragw_gps.h"

int main() {
	// Configuration
    int nb_cycles = 1000;
    int nb_uplink_cycle = 10;
    int nb_downlink_cycle = 2;

    // Loop variables
    int i;

    // Setting board
    struct lgw_conf_board_s board_conf;
    board_conf.clksrc = 0;
    board_conf.lorawan_public = 1;
    lgw_board_setconf(board_conf);

    // Setting TX Gain
    struct lgw_tx_gain_lut_s gain_lut;
    gain_lut.size = 16;
    for (i = 0; i < TX_GAIN_LUT_SIZE_MAX; i++) {
        gain_lut.lut[i].dig_gain = 0;
        gain_lut.lut[i].dac_gain = 3;
    }
    for (i = 0; i < 3; i++) {
        gain_lut.lut[i].pa_gain = 0;
    }
    for (i = 3; i < 7; i++) {
        gain_lut.lut[i].pa_gain = 1;
    }
    gain_lut.lut[7].pa_gain = 2;
    gain_lut.lut[8].pa_gain = 1;
    gain_lut.lut[9].pa_gain = 2;
    gain_lut.lut[10].pa_gain = 2;
    for (i = 11; i < 16; i++) {
        gain_lut.lut[i].pa_gain = 3;
    }
    int mix_gains[16] = {8, 10, 12, 8, 10, 12, 13, 9, 15, 10, 11, 9, 10, 11, 12, 14};
    for (i = 0; i < 16; i++) {
        gain_lut.lut[i].mix_gain = mix_gains[i];
    }
    int rf_powers[16] = {-6,-3,0,3,6,10,11,12,13,14,16,20,23,25,26,27};
    for (i = 0; i < 16; i++) {
        gain_lut.lut[i].rf_power = rf_powers[i];
    }

    // Setting RF channels
    struct lgw_conf_rxrf_s radio0;
    radio0.enable = 1;
    radio0.freq_hz = 867500000;
    radio0.rssi_offset = -166;
    radio0.tx_enable = 1;
    radio0.type = LGW_RADIO_TYPE_SX1257;
    struct lgw_conf_rxrf_s radio1;
    radio1.enable = 1;
    radio1.freq_hz = 868500000;
    radio1.rssi_offset = -166;
    radio1.tx_enable = 0;
    radio1.type = LGW_RADIO_TYPE_SX1257;
    lgw_rxrf_setconf(0, radio0);
    lgw_rxrf_setconf(1, radio1);

    // Setting SF channels
    struct lgw_conf_rxif_s channels[8];
    int if_values[8] = {-400000, -200000, 0, -400000, -200000, 0, 200000, 400000};

    for (i = 0; i < 8; i++) {
        channels[i].enable = 1;
        if (i <= 2) {
            channels[i].rf_chain = 1;
        } else {
            channels[i].rf_chain = 0;
        }
        channels[i].freq_hz = if_values[i];
        channels[i].bandwidth = BW_UNDEFINED;
        channels[i].datarate = DR_UNDEFINED;
        lgw_rxif_setconf(i, channels[i]);
    }

    // Setting LoRa channel
    struct lgw_conf_rxif_s lora_channel;
    lora_channel.enable = 1;
    lora_channel.rf_chain = 1;
    lora_channel.freq_hz = -200000;
    lora_channel.bandwidth = BW_250KHZ;
    lora_channel.datarate = DR_UNDEFINED;
    lgw_rxif_setconf(8, lora_channel);

    // Setting FSK channel
    struct lgw_conf_rxif_s fsk_channel;
    fsk_channel.enable = 1;
    fsk_channel.rf_chain = 1;
    fsk_channel.freq_hz = 300000;
    fsk_channel.bandwidth = BW_125KHZ;
    fsk_channel.datarate = 50000;
    lgw_rxif_setconf(9, fsk_channel);

    // Start
    if (lgw_start() != LGW_HAL_SUCCESS) {
        printf("Concentrator start unsuccessful\n");
        return 1;
    }

    // Loop
    struct lgw_pkt_rx_s packets[8];
    for (i = 0; i < 10000; i++) {
        lgw_receive(8, packets);
    }

    // Stop
    lgw_stop();

    return 0;
}
