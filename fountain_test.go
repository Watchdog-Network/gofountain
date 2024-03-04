package gofountain

// import (
// 	"encoding/json"
// 	"fmt"

// 	"math/rand"
// 	"reflect"
// 	"testing"
// )

// func TestAvailableKsize(t *testing.T) {
// 	availableKsize := make([]int, 0)
// 	data := make(map[string]interface{}) // 문자열을 키로하고 모든 자료형을 저장할 수 있는 맵 생성

// 	data["name"] = "kim"
// 	data["age"] = 25

// 	message, _ := json.Marshal(data)

// 	for kSize := 4; kSize < 100; kSize++ {
// 		c := NewRaptorCodec(kSize, 4)
// 		ids := make([]int64, kSize+50)
// 		random := rand.New(rand.NewSource(8923489))
// 		for i := range ids {
// 			ids[i] = int64(random.Intn(60000))
// 		}

// 		messageCopy := make([]byte, len(message))

// 		copy(messageCopy, message)

// 		codeBlocks := EncodeLTBlocks(messageCopy, ids, c)

// 		t.Log("DECODE--------")
// 		decoder := newRaptorDecoder(c.(*raptorCodec), len(message))
// 		for i := 0; i < kSize+50; i++ {
// 			decoder.AddBlocks([]LTBlock{codeBlocks[i]})
// 		}
// 		if decoder.matrix.determined() {
// 			out := decoder.Decode()

// 			if !reflect.DeepEqual(message, out) {
// 				fmt.Printf("Not Equal: kSize %v\n", kSize)
// 			} else {
// 				availableKsize = append(availableKsize, kSize)
// 			}
// 		} else {
// 			fmt.Printf("Not determined: kSize %v \n", kSize)
// 		}
// 	}

// 	fmt.Printf("AvailableKsize: %v\n", availableKsize)
// }

// func Test4c3Combination(t *testing.T) {
// 	kSize := 17
// 	encodedSymbolSize := 7

// 	data := make(map[string]interface{}) // 문자열을 키로하고 모든 자료형을 저장할 수 있는 맵 생성

// 	data["name"] = "kim"
// 	data["age"] = 25

// 	message1, _ := json.Marshal(data)

// 	data1 := make(map[string]interface{}) // 문자열을 키로하고 모든 자료형을 저장할 수 있는 맵 생성

// 	data["name"] = "kim"
// 	data["age"] = 290

// 	message2, _ := json.Marshal(data1)

// 	producer1 := RaptorEncode(message1, kSize, encodedSymbolSize, 8923483)
// 	producer2 := RaptorEncode(message1, kSize, encodedSymbolSize, 8923484)
// 	producer3 := RaptorEncode(message1, kSize, encodedSymbolSize, 9234855)
// 	producer4 := RaptorEncode(message2, kSize, encodedSymbolSize, 8923486) // malicious

// 	serializedEncodedBlocks1 := make([][]byte, 3)
// 	serializedEncodedBlocks1[0] = producer2
// 	serializedEncodedBlocks1[1] = producer3
// 	serializedEncodedBlocks1[2] = producer4

// 	serializedEncodedBlocks2 := make([][]byte, 3)
// 	serializedEncodedBlocks2[0] = producer1
// 	serializedEncodedBlocks2[1] = producer3
// 	serializedEncodedBlocks2[2] = producer4

// 	serializedEncodedBlocks3 := make([][]byte, 3)
// 	serializedEncodedBlocks3[0] = producer1
// 	serializedEncodedBlocks3[1] = producer2
// 	serializedEncodedBlocks3[2] = producer4

// 	serializedEncodedBlocks4 := make([][]byte, 3) // Not malicious
// 	serializedEncodedBlocks4[0] = producer1
// 	serializedEncodedBlocks4[1] = producer2
// 	serializedEncodedBlocks4[2] = producer3

// 	out1, err := RaptorDecode(serializedEncodedBlocks1, kSize, len(message1))
// 	out2, err := RaptorDecode(serializedEncodedBlocks2, kSize, len(message1))
// 	out3, err := RaptorDecode(serializedEncodedBlocks3, kSize, len(message1))
// 	out4, err := RaptorDecode(serializedEncodedBlocks4, kSize, len(message1))

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	var jsonData []structures.EncodedBlockData
// 	err = json.Unmarshal(out1, &jsonData)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("out1: %v\n", string(out1))

// 	err = json.Unmarshal(out2, &jsonData)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("out2: %v\n", string(out2))

// 	err = json.Unmarshal(out3, &jsonData)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("out3: %v\n", string(out3))

// 	err = json.Unmarshal(out4, &jsonData)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("out4: %v\n", string(out4))

// }
