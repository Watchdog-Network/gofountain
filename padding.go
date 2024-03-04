package gofountain

import (
	"encoding/json"
	"os"
	"strconv"
)

// A HLFEncodedBlock represents Hyperledger fabric block that store transaction
// information in the ledger.
type HLFEncodedBlock struct {
	BlockHeader   BlockHeader
	BlockData     BlockData
	BlockMetadata BlockMetadata
	// How many padding bytes this block has at the end.
	Padding []int
}

// The BlockHeader consist of block number, copy of the previous block hash and
// the current block hash.
type BlockHeader struct {
	Number       uint64
	PreviousHash []byte
	DataHash     []byte
}

// The BlockData fields of a block are the essentials segment that contents the
// transaction details sorted in the byte array.
type BlockData struct {
	Data [][]byte
}

// The BlockMetadata fields contain the created time of the block, certificate
// details and signature of the block writer.
type BlockMetadata struct {
	Metadata [][]byte
}

// EqualizeParsedBlockLentghs adds padding to parsed blocks to make them devided by number of source symbols.
// Returns a slice of serialized parsedblock with padding, numSourceSymbols, symbolAlignmentSize, numEncodedSourceSymbols.
func EqualizeParsedBlockLengths(parsedBlock HLFEncodedBlock) ([]byte, int, int, int) {
	// numSourceSymbols means how many source symbols the input message will be divided into
	// and the minimum number of source symbols required for decoding.
	numSourceSymbols, _ := strconv.Atoi(os.Getenv("NUM_SOURCE_SYMBOLS"))

	// symbolAlignmentSize, the size of ach symbol in the source message in bytes.
	symbolAlignmentSize, _ := strconv.Atoi(os.Getenv("SYMBOL_ALIGNMENT_SIZE"))

	// numEncodedSourceSymbols means how many encoded source symbols will be created using source symbols
	// and the maximum number of source symbols required for decoding.
	numEncodedSourceSymbols, _ := strconv.Atoi(os.Getenv("NUM_ENCODED_SOURCE_SYMBOLS"))

	marshalledBlock, _ := json.Marshal(parsedBlock)

	var padding []int
	if len(marshalledBlock)%numSourceSymbols != 0 {
		if ((numSourceSymbols-(len(marshalledBlock)%numSourceSymbols))/2)%2 == 0 {
			padding = make([]int, ((numSourceSymbols-(len(marshalledBlock)%numSourceSymbols))/2)+2)
		} else {
			padding = make([]int, ((numSourceSymbols-(len(marshalledBlock)%numSourceSymbols))/2)+1)
		}
	}

	parsedBlock.Padding = padding
	marshalledBlock, _ = json.Marshal(parsedBlock)

	return marshalledBlock, numSourceSymbols, symbolAlignmentSize, numEncodedSourceSymbols
}
