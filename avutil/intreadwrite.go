package avutil

//#include <libavutil/intreadwrite.h>
/*
uint32_t avRL32(uint8_t *i, int o) {
	return AV_RL32(i+o);
}
*/
import "C"

func AvRL32(i *uint8, offset uint) uint32 {
	return uint32(C.avRL32((*C.uint8_t)(i), C.int(offset)))
}
