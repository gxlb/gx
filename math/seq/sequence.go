package seq

import (
	"math"
)

type SeqOption uint32

const (
	MEM                SeqOption = 1 //only in memory
	STORE              SeqOption = 2 //with storage
	_storage_type_bits           = 4
	_get_storage                 = (1 << _storage_type_bits) - 1

	NO_REC             SeqOption = 1 << _storage_type_bits //never recycle any id
	REC_LOOP           SeqOption = 2 << _storage_type_bits //gen id by loop
	REC_FST            SeqOption = 3 << _storage_type_bits //use recycled-id first
	REC_LST            SeqOption = 4 << _storage_type_bits //use unused-id first
	_recycle_type_bits           = 4
	_get_cecycle                 = (1<<_recycle_type_bits - 1) << _storage_type_bits

	NOUSE_FST          = 1 << (_storage_type_bits + _recycle_type_bits)
	PANIC_ERR          = 2 << (_storage_type_bits + _recycle_type_bits) //panic with error
	_usage_option_bits = 4
	_get_usage_option  = (1<<_usage_option_bits - 1) << (_storage_type_bits + _recycle_type_bits)

	DFT_OP_ID_INC    = MEM | NO_REC | NOUSE_FST
	DFT_OP_ID_NO_ERR = MEM | NO_REC | NOUSE_FST | PANIC_ERR
	DFT_OP_ID_LOOP   = MEM | REC_LOOP | NOUSE_FST
	DFT_OP_ID_RECF   = MEM | REC_FST | NOUSE_FST
	DFT_OP_ID_RECL   = MEM | REC_LST | NOUSE_FST
	DFT_OP_ID_GBL    = STORE | NO_REC | REC_FST | NOUSE_FST

	MAX_SEQ64 = math.MaxUint64
	MAX_SEQ32 = math.MaxUint32
)
