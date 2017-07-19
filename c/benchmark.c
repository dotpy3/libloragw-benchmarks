#include "config.h"
#include "loragw_hal.h"
#include "loragw_gps.h"

int main(int argc, char * argv) {
	// Configuration
    int nb_cycles = 1000;
    int nb_uplink_cycle = 10;
    int nb_downlink_cycle = 2;

    // Setting board
    struct lgw_conf_board_s board_conf;
    board_conf.clksrc = 8;
    board_conf.lorawan_public = true;

    // Setting TX Gain

    // Setting RF channels

    // Setting SF channels

    // Setting LoRa channel

    // Setting FSK channel

    // Start

    // Loop

    // Stop

    return 0;
}
