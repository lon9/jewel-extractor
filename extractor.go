package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
)

// Extractor extract information of decorations
type Extractor struct {
	key4096   []uint32
	key72     []uint32
	jewelList JewelList
}

// NewExtractor is constructor
func NewExtractor() (res *Extractor, err error) {
	res = &Extractor{}
	key4096Buf := bytes.NewBuffer(key4096Bytes)
	res.key4096 = make([]uint32, len(key4096Bytes)/4)
	for i := 0; i < len(key4096Bytes); i += 4 {
		if err = binary.Read(key4096Buf, binary.LittleEndian, &res.key4096[i/4]); err != nil {
			return
		}
	}
	key72Buf := bytes.NewBuffer(key72Bytes)
	res.key72 = make([]uint32, len(key72Bytes))
	for i := 0; i < len(key72Bytes); i += 4 {
		if err = binary.Read(key72Buf, binary.LittleEndian, &res.key72[i/4]); err != nil {
			return
		}
	}
	if res.jewelList, err = NewJewelList(); err != nil {
		return
	}
	return
}

// Extract extracts decorations
func (e *Extractor) Extract(fname string, slot int, lang string) (err error) {

	if slot < 1 || slot > 3 {
		return errors.New("slot is must be 1-3")
	}

	if lang != "en" && lang != "ja" && lang != "zh" && lang != "ko" {
		return errors.New("supported languages are en, ja, zh, ko")
	}

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}
	decryptedByte, err := e.decrypt(b)
	if err != nil {
		return
	}
	decrypted, err := ioutil.ReadAll(decryptedByte)
	if err != nil {
		return
	}
	r := bytes.NewReader(decrypted)
	// Go to section3
	if err = e.gotoSection3PastSignature(r); err != nil {
		return
	}
	r.Seek(4+8+4, io.SeekCurrent)

	// Extract decoration for each sections
	for i := 0; i < 3; i++ {
		verbose := false
		if i == slot-1 {
			verbose = true
		}
		decorations, err := e.readSaveSlot(r, verbose)
		if err != nil {
			return err
		}
		if verbose {
			if err = e.export(decorations, lang); err != nil {
				return err
			}
		}
	}
	return
}

func (e *Extractor) export(decorations map[uint32]uint32, lang string) (err error) {
	output := make(map[string]uint32)
	for _, jewel := range e.jewelList {
		if v, ok := decorations[jewel.ItemID]; ok {
			output[jewel.Locales[lang]] = uint32(math.Min(float64(jewel.Max), float64(v)))
		} else if jewel.EquippedItemID != 0xFFFFFFFF {
			output[jewel.Locales[lang]] = 0
		}
	}
	enc := json.NewEncoder(os.Stdout)
	fmt.Printf("\n====   Jewels   ====\n\n")
	return enc.Encode(output)
}

func (e *Extractor) decrypt(b []byte) (io.Reader, error) {
	buf := bytes.NewReader(b)

	size := buf.Len()

	input := make([]uint32, size/4)
	for i := 0; i < len(input); i++ {
		if err := binary.Read(buf, binary.LittleEndian, &input[i]); err != nil {
			return nil, err
		}
	}
	output := make([]uint64, size/8)
	for i := 0; i < len(output); i++ {
		output[i] = e.decryptBlock(input[i*2], input[i*2+1])
	}

	outBuf := new(bytes.Buffer)
	if err := binary.Write(outBuf, binary.LittleEndian, output); err != nil {
		return nil, err
	}
	return outBuf, nil
}

func (e *Extractor) decryptBlock(data, nextData uint32) uint64 {
	key72Offset := uint32(17)
	var A uint32
	var B uint32

	A = data ^ e.key72[key72Offset]
	key72Offset--
	B = e.basicAlg(key72Offset, A, nextData)
	key72Offset--

	for i := 0; i < 7; i++ {
		A = e.basicAlg(key72Offset, B, A)
		key72Offset--
		B = e.basicAlg(key72Offset, A, B)
		key72Offset--
	}
	A = e.basicAlg(key72Offset, B, A)
	key72Offset--
	B ^= e.key72[key72Offset]
	return uint64(A)<<32 | uint64(B)
}

func (e *Extractor) basicAlg(arr720Offset, prevStep1, prevStep2 uint32) uint32 {
	A := e.b2(prevStep1) + 0x100
	result := e.key4096[A]
	B := e.b3(prevStep1)
	result += e.key4096[B]
	B = e.b1(prevStep1) + 0x200
	result ^= e.key4096[B]
	A = e.b(prevStep1) + 0x300
	result += e.key4096[A]
	result ^= e.key72[arr720Offset]
	result ^= prevStep2
	return result
}

func (*Extractor) b(data uint32) uint32 {
	return data & 0xFF
}

func (*Extractor) b1(data uint32) uint32 {
	return (data >> 8) & 0xFF
}

func (*Extractor) b2(data uint32) uint32 {
	return (data >> 16) & 0xFF
}

func (*Extractor) b3(data uint32) uint32 {
	return (data >> 24) & 0xFF
}

func (e *Extractor) gotoSection3PastSignature(r *bytes.Reader) error {
	r.Seek(64+8*3, io.SeekStart)
	var section3Offset int64
	if err := binary.Read(r, binary.LittleEndian, &section3Offset); err != nil {
		return err
	}
	r.Seek(section3Offset, io.SeekStart)
	var sig uint32
	if err := binary.Read(r, binary.LittleEndian, &sig); err != nil {
		return err
	}
	if sig != section3Signature {
		return fmt.Errorf("section 3 signature must be %x (actualy: %x)", section3Signature, sig)
	}
	return nil
}

func (e *Extractor) readUntilPlaytimeIncluded(r *bytes.Reader, verbose bool) (err error) {
	if verbose {

		fmt.Printf("\n=====   Hunter Informations   ====\n\n")

		hunterNameBytes := make([]byte, 64)
		if _, err = r.Read(hunterNameBytes); err != nil {
			return
		}
		hunterName := bytes.TrimSpace(hunterNameBytes)
		fmt.Println("Hunter Name:", string(hunterName))
		var hunterRank uint32
		if err = binary.Read(r, binary.LittleEndian, &hunterRank); err != nil {
			return
		}
		fmt.Println("Hunter Rank:", hunterRank)
		var zeni uint32
		if err = binary.Read(r, binary.LittleEndian, &zeni); err != nil {
			return
		}
		fmt.Println("Zeni:", zeni)
		var researchPoint uint32
		if err = binary.Read(r, binary.LittleEndian, &researchPoint); err != nil {
			return
		}
		fmt.Println("Research Point:", researchPoint)
		var hunterXP uint32
		if err = binary.Read(r, binary.LittleEndian, &hunterXP); err != nil {
			return
		}
		fmt.Println("Hunter XP:", hunterXP)
		var playtime uint32
		if err = binary.Read(r, binary.LittleEndian, &playtime); err != nil {
			return
		}
		fmt.Println("Play Time:", playtime)
	} else {
		r.Seek(64+4+4+4+4+4, io.SeekCurrent)
	}
	return
}

func (e *Extractor) readSaveSlot(r *bytes.Reader, verbose bool) (decorations map[uint32]uint32, err error) {

	// Read player info
	if err = e.readUntilPlaytimeIncluded(r, verbose); err != nil {
		return
	}
	r.Seek(4+
		hunterAppearanceStructureSize+
		palicoAppearanceStructureSize+
		guildCardStructureSize+
		guildCardStructureSize*100+
		0x019e36+
		itemLoadoutsStructureSize+
		8+
		itemPouchStructureSize, io.SeekCurrent)
	r.Seek(8*200+8*200+8*800, io.SeekCurrent)
	decorations = make(map[uint32]uint32)
	for i := 0; i < 200; i++ {
		var itemID uint32
		if err = binary.Read(r, binary.LittleEndian, &itemID); err != nil {
			return
		}
		var itemQuantity uint32
		if err = binary.Read(r, binary.LittleEndian, &itemQuantity); err != nil {
			return
		}
		if itemID > 0 {
			decorations[itemID] = itemQuantity
		}
	}
	for i := 0; i < 1000; i++ {
		e.readEquipmentSlot(r, decorations)
	}
	r.Seek(0x2449C+
		0x2a*250+
		0x0FB9+
		equipLoadoutsStructureSize+
		0x6521+
		dlcTypeSize*256+
		0x2A5D, io.SeekCurrent)
	return
}

func (e *Extractor) readEquipmentSlot(r *bytes.Reader, decorations map[uint32]uint32) (err error) {
	r.Seek(4, io.SeekCurrent)
	var equipmentType uint32
	if err = binary.Read(r, binary.LittleEndian, &equipmentType); err != nil {
		return
	}
	r.Seek(4+4+4+4, io.SeekCurrent)
	if equipmentType == 0 || equipmentType == 1 {
		for i := 0; i < 3; i++ {
			var deco uint32
			if err = binary.Read(r, binary.LittleEndian, &deco); err != nil {
				return
			}
			if deco == math.MaxUint32 {
				continue
			}
			jewel, err := e.jewelList.FindJewelByEquippedItemID(deco)
			if err != nil {
				return err
			}
			if _, ok := decorations[jewel.ItemID]; ok {
				decorations[jewel.ItemID]++
			} else {
				decorations[jewel.ItemID] = 1
			}
		}
	} else {
		r.Seek(4+4+4, io.SeekCurrent)
	}
	r.Seek(4+4+4+4+4+4+8, io.SeekCurrent)
	return
}
