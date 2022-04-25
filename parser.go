package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	structs "pe/oop"
	util "pe/util"
	"strconv"
	"strings"
)

func DisplayDosHeader(DosHeader *structs.IMAGE_DOS_HEADER) {

	fmt.Println("============= DOS Header =============")
	fmt.Printf("\te_magic: %v\n", string(DosHeader.E_magic[:]))
	fmt.Printf("\te_cblp: %v\n", DosHeader.E_cblp)
	fmt.Printf("\te_cp: %v\n", DosHeader.E_cp)
	fmt.Printf("\te_crlc: %v\n", DosHeader.E_crlc)
	fmt.Printf("\te_cparhdr: %v\n", DosHeader.E_cparhdr)
	fmt.Printf("\te_minalloc: %v\n", DosHeader.E_minalloc)
	fmt.Printf("\te_maxalloc: %v\n", DosHeader.E_maxalloc)
	fmt.Printf("\te_ss: %v\n", DosHeader.E_ss)
	fmt.Printf("\te_sp: %v\n", DosHeader.E_sp)
	fmt.Printf("\te_csum: %v\n", DosHeader.E_csum)
	fmt.Printf("\te_ip: %v\n", DosHeader.E_ip)
	fmt.Printf("\te_cs: %v\n", DosHeader.E_cs)
	fmt.Printf("\te_lfarlc: %v\n", DosHeader.E_lfarlc)
	fmt.Printf("\te_ovno: %v\n", DosHeader.E_ovno)
	fmt.Printf("\te_res: %v\n", DosHeader.E_res)
	fmt.Printf("\te_oemid: %v\n", DosHeader.E_oemid)
	fmt.Printf("\te_oeminfo: %v\n", DosHeader.E_oeminfo)
	fmt.Printf("\te_res2: %v\n", DosHeader.E_res2)
	fmt.Printf("\te_lfanew: %v\n", DosHeader.E_lfanew)

}

func DisplayRichHeader(RichHeader *structs.IMAGE_RICH_HEADER) {
	fmt.Println("============= Rich Header =============")

	fmt.Printf("\tDanS ID: %v\n", string(RichHeader.DansID[:]))

	fmt.Printf("\tComp ID 1\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID1.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID1.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID1.Count)

	fmt.Printf("\tComp ID 2\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID2.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID2.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID2.Count)

	fmt.Printf("\tComp ID 3\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID3.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID3.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID3.Count)

	fmt.Printf("\tComp ID 4\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID4.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID4.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID4.Count)

	fmt.Printf("\tComp ID 5\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID5.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID5.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID5.Count)

	fmt.Printf("\tComp ID 6\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID6.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID6.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID6.Count)

	fmt.Printf("\tComp ID 7\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID7.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID7.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID7.Count)

	fmt.Printf("\tComp ID 8\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID8.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID8.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID8.Count)

	fmt.Printf("\tComp ID 9\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID9.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID9.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID9.Count)

	fmt.Printf("\tComp ID 10\n")
	fmt.Printf("\t\tProduct Id: %v\n", util.DetermineProdId(RichHeader.CompID10.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID10.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID10.Count)

	fmt.Printf("\tRich Identifier: %v\n", string(RichHeader.RichIdentifier[:]))
	fmt.Printf("\tChecksum: %v\n", RichHeader.Checksum)
}

func DisplayFileHeader(FileHeader *structs.IMAGE_FILE_HEADER) {

	ImageMachine := util.GetImageMachine(FileHeader.ImageMachine)

	fmt.Println("============= File Header =============")
	fmt.Printf("\tMachine: %v\n", ImageMachine)
	fmt.Printf("\tNumberOfSections: %v\n", FileHeader.NumberOfSections)
	fmt.Printf("\tTimeDateStamp: %v\n", FileHeader.TimeDateStamp)
	fmt.Printf("\tPointerToSymbolTable: %v\n", FileHeader.PointerToSymbolTable)
	fmt.Printf("\tNumberOfSymbols: %v\n", FileHeader.NumberOfSymbols)
	fmt.Printf("\tSizeOfOptionalHeader: %v\n", FileHeader.SizeOfOptionalHeader)
	fmt.Printf("\tCharacteristics: %v\n", FileHeader.Characteristics)
}

func DisplayOptionalHeader32(OptionalHeader *structs.IMAGE_OPTIONAL_HEADER32, bitRange string) {

	fmt.Println("============= Optional Header =============")
	fmt.Printf("\tMagic: %v (%v)\n", OptionalHeader.Magic, bitRange)
	fmt.Printf("\tMajorLinkerVersion: %v\n", OptionalHeader.MajorLinkerVersion)
	fmt.Printf("\tMinorLinkerVersion: %v\n", OptionalHeader.MinorLinkerVersion)
	fmt.Printf("\tSizeOfCode: %v bytes\n", OptionalHeader.SizeOfCode)
	fmt.Printf("\tSizeOfInitializedData: %v bytes\n", OptionalHeader.SizeOfInitializedData)
	fmt.Printf("\tAddressOfEntryPoint: %v\n", OptionalHeader.AddressOfEntryPoint)
	fmt.Printf("\tBaseOfCode: %v\n", OptionalHeader.BaseOfCode)
	fmt.Printf("\tBaseOfData: %v\n", OptionalHeader.BaseOfData)
	fmt.Printf("\tImageBase: %v\n", OptionalHeader.ImageBase)
	fmt.Printf("\tSectionAlignment: %v bytes\n", OptionalHeader.SectionAlignment)
	fmt.Printf("\tFileAlignment: %v bytes\n", OptionalHeader.FileAlignment)
	fmt.Printf("\tMajorOperatingSystemVersion: %v\n", OptionalHeader.MajorOperatingSystemVersion)
	fmt.Printf("\tMinorOperatingSystemVersion: %v\n", OptionalHeader.MinorOperatingSystemVersion)
	fmt.Printf("\tMajorImageVersion: %v\n", OptionalHeader.MajorImageVersion)
	fmt.Printf("\tMinorImageVersion: %v\n", OptionalHeader.MinorImageVersion)
	fmt.Printf("\tMajorSubsystemVersion: %v\n", OptionalHeader.MajorSubsystemVersion)
	fmt.Printf("\tMinorSubsystemVersion: %v\n", OptionalHeader.MinorSubsystemVersion)
	fmt.Printf("\tWin32Version: 0\n")
	fmt.Printf("\tSizeOfImage: %v bytes\n", OptionalHeader.SizeOfImage)
	fmt.Printf("\tSizeOfHeaders: %v bytes\n", OptionalHeader.SizeOfHeaders)
	fmt.Printf("\tCheckSum: %v\n", OptionalHeader.CheckSum)
	fmt.Printf("\tSubsystem: %v\n", OptionalHeader.Subsystem)
	fmt.Printf("\tDllCharacteristics: %v\n", OptionalHeader.DllCharacteristics)
	fmt.Printf("\tSizeOfStackReserve: %v bytes\n", OptionalHeader.SizeOfStackReserve)
	fmt.Printf("\tSizeOfStackCommit: %v bytes\n", OptionalHeader.SizeOfStackCommit)
	fmt.Printf("\tSizeOfHeapReserve: %v bytes\n", OptionalHeader.SizeOfHeapReserve)
	fmt.Printf("\tSizeOfHeapCommit: %v bytes\n", OptionalHeader.SizeOfHeapCommit)
	fmt.Printf("\tLoaderFlags: %v\n", OptionalHeader.LoaderFlags)
	fmt.Printf("\tNumberOfRvaAndSizes: %v\n", OptionalHeader.NumberOfRvaAndSizes)
}

func DisplayOptionalHeader64(OptionalHeader *structs.IMAGE_OPTIONAL_HEADER64, bitRange string) {

	fmt.Println("============= Optional Header =============")
	fmt.Printf("\tMagic: %v (%v)\n", OptionalHeader.Magic, bitRange)
	fmt.Printf("\tMajorLinkerVersion: %v\n", OptionalHeader.MajorLinkerVersion)
	fmt.Printf("\tMinorLinkerVersion: %v\n", OptionalHeader.MinorLinkerVersion)
	fmt.Printf("\tSizeOfCode: %v bytes\n", OptionalHeader.SizeOfCode)
	fmt.Printf("\tSizeOfInitializedData: %v bytes\n", OptionalHeader.SizeOfInitializedData)
	fmt.Printf("\tAddressOfEntryPoint: %v\n", OptionalHeader.AddressOfEntryPoint)
	fmt.Printf("\tBaseOfCode: %v\n", OptionalHeader.BaseOfCode)
	fmt.Printf("\tImageBase: %v\n", OptionalHeader.ImageBase)
	fmt.Printf("\tSectionAlignment: %v bytes\n", OptionalHeader.SectionAlignment)
	fmt.Printf("\tFileAlignment: %v bytes\n", OptionalHeader.FileAlignment)
	fmt.Printf("\tMajorOperatingSystemVersion: %v\n", OptionalHeader.MajorOperatingSystemVersion)
	fmt.Printf("\tMinorOperatingSystemVersion: %v\n", OptionalHeader.MinorOperatingSystemVersion)
	fmt.Printf("\tMajorImageVersion: %v\n", OptionalHeader.MajorImageVersion)
	fmt.Printf("\tMinorImageVersion: %v\n", OptionalHeader.MinorImageVersion)
	fmt.Printf("\tMajorSubsystemVersion: %v\n", OptionalHeader.MajorSubsystemVersion)
	fmt.Printf("\tMinorSubsystemVersion: %v\n", OptionalHeader.MinorSubsystemVersion)
	fmt.Printf("\tWin32Version: 0\n")
	fmt.Printf("\tSizeOfImage: %v bytes\n", OptionalHeader.SizeOfImage)
	fmt.Printf("\tSizeOfHeaders: %v bytes\n", OptionalHeader.SizeOfHeaders)
	fmt.Printf("\tCheckSum: %v\n", OptionalHeader.CheckSum)
	fmt.Printf("\tSubsystem: %v\n", OptionalHeader.Subsystem)
	fmt.Printf("\tDllCharacteristics: %v\n", OptionalHeader.DllCharacteristics)
	fmt.Printf("\tSizeOfStackReserve: %v bytes\n", OptionalHeader.SizeOfStackReserve)
	fmt.Printf("\tSizeOfStackCommit: %v bytes\n", OptionalHeader.SizeOfStackCommit)
	fmt.Printf("\tSizeOfHeapReserve: %v bytes\n", OptionalHeader.SizeOfHeapReserve)
	fmt.Printf("\tSizeOfHeapCommit: %v bytes\n", OptionalHeader.SizeOfHeapCommit)
	fmt.Printf("\tLoaderFlags: %v\n", OptionalHeader.LoaderFlags)
	fmt.Printf("\tNumberOfRvaAndSizes: %v\n", OptionalHeader.NumberOfRvaAndSizes)
}

func DisplayDataDirectories(DataDirectory *[16]structs.IMAGE_DATA_DIRECTORY, bitRange string) (uint32, uint32) {

	fmt.Println("============= Data Directories =============")

	var importVA uint32
	var importSize uint32

	var baseOffset int = 96

	if bitRange == "64bit" {
		baseOffset = 112
	}

	for i := 0; i < 16; i++ {
		nOffset := baseOffset + (i * 8)

		var directoryName string
		if bitRange == "64bit" {
			directoryName = util.GetDataDirectoryName64(nOffset)
		} else {
			directoryName = util.GetDataDirectoryName32(nOffset)
		}

		fmt.Printf("\tData Directory (%v)\n", directoryName)
		fmt.Printf("\t\timportVA: %v\n", DataDirectory[i].VirtualAddress)
		fmt.Printf("\t\timportSize: %v\n", DataDirectory[i].Size)

		if directoryName == "import table" {
			importVA = DataDirectory[i].VirtualAddress
			importSize = DataDirectory[i].Size
		}
	}

	return importVA, importSize
}

func DisplaySectionTable(SectionTable *[]structs.IMAGE_SECTION_HEADER, numSections uint16, importVirtualAddr uint32) (uint32, uint32, string, uint32, uint32) {

	var sectionNumber uint16
	var highestAddress uint32
	var importPtrToRawData uint32
	var importSectionName string

	var relocPtrToRawData uint32
	var relocSize uint32

	fmt.Println("============= Section Tables =============")
	for sectionNumber = 0; sectionNumber < numSections; sectionNumber++ {

		sectionIndex := (*SectionTable)[sectionNumber]

		sectionName := string(sectionIndex.Name[:])
		sectionVirtualSize := sectionIndex.VirtualSize
		sectionVA := sectionIndex.VirtualAddress
		sectionSizeOfRawData := sectionIndex.SizeOfRawData
		sectionPointerToRawData := sectionIndex.PointerToRawData
		sectionPointerToRelocations := sectionIndex.PointerToRelocations
		sectionPointerToLinenumbers := sectionIndex.PointerToLinenumbers
		sectionNumberOfRelocations := sectionIndex.NumberOfRelocations
		sectionNumberOfLinenumbers := sectionIndex.NumberOfLinenumbers
		sectionCharacteristics := sectionIndex.Characteristics

		if sectionVA >= highestAddress && sectionVA <= importVirtualAddr && sectionVA+sectionVirtualSize >= importVirtualAddr {
			highestAddress = sectionVA
			importPtrToRawData = sectionPointerToRawData
			importSectionName = sectionName
		}

		if strings.Index(sectionName, "reloc") > -1 {
			relocPtrToRawData = sectionPointerToRawData
			relocSize = sectionSizeOfRawData
		}

		fmt.Printf("\t%v table\n", sectionName)
		fmt.Printf("\t\tVirtualSize: %v\n", sectionVirtualSize)
		fmt.Printf("\t\tVirtualAddress: %v\n", sectionVA)
		fmt.Printf("\t\tSizeOfRawData: %v\n", sectionSizeOfRawData)
		fmt.Printf("\t\tPointerToRawData: %v\n", sectionPointerToRawData)
		fmt.Printf("\t\tPointerToRelocations: %v\n", sectionPointerToRelocations)
		fmt.Printf("\t\tPointerToLinenumbers: %v\n", sectionPointerToLinenumbers)
		fmt.Printf("\t\tNumberOfRelocations: %v\n", sectionNumberOfRelocations)
		fmt.Printf("\t\tNumberOfLinenumbers: %v\n", sectionNumberOfLinenumbers)
		fmt.Printf("\t\tCharacteristics: %v\n", sectionCharacteristics)
	}

	return highestAddress, importPtrToRawData, importSectionName, relocPtrToRawData, relocSize
}

func ParseSectionTable(data *[]byte, sectionTableOffset uint16, numSections uint16) *[]structs.IMAGE_SECTION_HEADER {

	sectionBuffer := bytes.Buffer{}
	sectionBuffer.Write((*data)[sectionTableOffset : sectionTableOffset+(numSections*40)])

	sectionHeaderArray := make([]structs.IMAGE_SECTION_HEADER, numSections)

	err := binary.Read(&sectionBuffer, binary.LittleEndian, &sectionHeaderArray)
	util.CheckError(err)

	return &sectionHeaderArray
}

func ParseImportDirectory(data *[]byte, importVA uint32, importSectionAddr uint32, importSize uint32, importSectionPtr uint32, importSectionName string) (*[]structs.IMAGE_IMPORT_DESCRIPTOR, int) {

	importTableRVA := importVA - importSectionAddr
	importTableLocation := importSectionPtr + importTableRVA
	fmt.Println("============= Import Directory =============")
	fmt.Printf("\tLocated in %v (VA: %v)\n", importSectionName, importSectionAddr)
	fmt.Printf("\tPointerToRawData: %v\n\tFile Offset: %v\n\n", importSectionPtr, importTableRVA)

	numImports, importEndStruct := util.GetNumImports(data, importTableLocation)

	importArray := make([]structs.IMAGE_IMPORT_DESCRIPTOR, numImports)

	fmt.Printf("\tNum. imports: %v\n", numImports)

	importBuffer := bytes.Buffer{}
	importBuffer.Write((*data)[importTableLocation:importEndStruct])

	fmt.Printf("\tImport Table VA: %v\n", importVA)
	fmt.Printf("\tImport Table Size: %v\n\n", importSize)

	err := binary.Read(&importBuffer, binary.LittleEndian, &importArray)
	util.CheckError(err)

	return &importArray, numImports
}

func ParseImportLookupTable(data *[]byte, bitRange string, importSectionAddr uint32, importSectionPtr uint32, importTableOffset uint32, importTableRVA uint32) {

	var funcIndex uint32

	if bitRange == "64bit" {

		numFuncs := util.GetNumImportFuncs(data, bitRange, importTableOffset)

		importTableArray := make([]structs.IMAGE_THUNK_DATA64, numFuncs)

		bILTBuffer := bytes.Buffer{}
		bILTBuffer.Write((*data)[importTableOffset : importTableOffset+(8*uint32(numFuncs))])

		err := binary.Read(&bILTBuffer, binary.LittleEndian, &importTableArray)
		util.CheckError(err)

		fmt.Printf("\t\tImport Lookup Table (%v)\n", numFuncs)

		for funcIndex = 0; funcIndex < uint32(numFuncs); funcIndex++ {

			ordinalFlagByte := importTableArray[funcIndex].ORDINAL_NAME_FLAG

			fmt.Printf("\t\t\tFunction Import %v\n", funcIndex+1)

			if ordinalFlagByte == 0 { // HINT_NAME_TABLE Value
				if importTableArray[funcIndex].HINT_NAME_TABLE > 0 {
					hintTableRVA := importTableArray[funcIndex].HINT_NAME_TABLE + 2
					hintTableOffset := hintTableRVA - importSectionAddr
					hintTableLocation := importSectionPtr + hintTableOffset

					numChars := util.GetNumCharacters(data, hintTableLocation)
					funcName := string((*data)[hintTableLocation : hintTableLocation+uint32(numChars)])

					fmt.Printf("\t\t\t\tName: %v\n", funcName)
					fmt.Printf("\t\t\t\t\t| HINT_NAME_TABLE\n")
					fmt.Printf("\t\t\t\t\t| RVA: %v\n", hintTableRVA)
					fmt.Printf("\t\t\t\t\t| Physical Offset: %v\n", hintTableOffset)

				} else {
					fmt.Printf("\t\t\tInvalid entry (HINT_NAME_TABLE == 0)\n")
				}

			} else { // Ordinal Value
				var ordinalVal uint16 = uint16(importTableArray[funcIndex].HINT_NAME_TABLE)
				fmt.Printf("\t\t\tOrdinal: %v\n", ordinalVal)
			}

			fmt.Printf("\t\t\t___________________________________________________\n")
		}

	} else { // 32 bit

		numFuncs := util.GetNumImportFuncs(data, bitRange, importTableOffset)

		importTableArray := make([]structs.IMAGE_THUNK_DATA32, numFuncs)

		bILTBuffer := bytes.Buffer{}
		bILTBuffer.Write((*data)[importTableOffset : importTableOffset+(4*uint32(numFuncs))])

		err := binary.Read(&bILTBuffer, binary.LittleEndian, &importTableArray)
		util.CheckError(err)

		fmt.Printf("\t\tImport Lookup Table (%v)\n", numFuncs)

		for funcIndex = 0; funcIndex < uint32(numFuncs); funcIndex++ {

			ordinalFlagByte := importTableArray[funcIndex].HINT_NAME_TABLE << 31
			fmt.Printf("\t\t\tFunction Import %v\n", funcIndex+1)

			if ordinalFlagByte == 0 {

				if importTableArray[funcIndex].HINT_NAME_TABLE > 0 {
					hintTableRVA := importTableArray[funcIndex].HINT_NAME_TABLE + 2
					hintTableOffset := hintTableRVA - importSectionAddr
					hintTableLocation := importSectionPtr + hintTableOffset

					numChars := util.GetNumCharacters(data, hintTableLocation)
					funcName := string((*data)[hintTableLocation : hintTableLocation+uint32(numChars)])

					fmt.Printf("\t\t\t\tName: %v\n", funcName)
					fmt.Printf("\t\t\t\t\t| HINT_NAME_TABLE\n")
					fmt.Printf("\t\t\t\t\t| RVA: %v\n", hintTableRVA)
					fmt.Printf("\t\t\t\t\t| Physical Offset: %v\n", hintTableOffset)

				} else {
					fmt.Printf("\t\t\tInvalid entry (HINT_NAME_TABLE == 0)\n")
				}

			} else {
				var ordinalVal uint16 = uint16(importTableArray[funcIndex].HINT_NAME_TABLE)
				fmt.Printf("\t\t\tOrdinal: %v\n", ordinalVal)
			}

			fmt.Printf("\t\t\t___________________________________________________\n")
		}
	}
}

func DisplayImportDirectory(bitRange string, data *[]byte, importArray *[]structs.IMAGE_IMPORT_DESCRIPTOR, importSectionAddr uint32, importSectionPtr uint32, numImports int) {

	for nImport := uint32(0); nImport < uint32(numImports); nImport++ {

		importObj := (*importArray)[nImport]

		nameRVA := importObj.Name - importSectionAddr
		nameLocation := nameRVA + importSectionPtr
		numChars := util.GetNumCharacters(data, nameLocation)

		fmt.Printf("\tImport %v\n", nImport+1)
		fmt.Printf("\t_____________________________\n")
		fmt.Printf("\t| OriginalFirstThunk: %v\n", importObj.OriginalFirstThunk)
		fmt.Printf("\t| FirstThunk: %v\n", importObj.FirstThunk)
		fmt.Printf("\t| TimeDateStamp: %v (%v)\n", importObj.TimeDateStamp, util.GetBoundStatus(importObj.TimeDateStamp))
		fmt.Printf("\t| ForwarderChain: %v\n", importObj.ForwarderChain)
		fmt.Printf("\t| Name: %v\n", string((*data)[nameLocation:nameLocation+uint32(numChars)]))

		importTableRVA := importObj.OriginalFirstThunk - importSectionAddr
		importTableOffset := importSectionPtr + importTableRVA

		ParseImportLookupTable(data, bitRange, importSectionAddr, importSectionPtr, importTableOffset, importTableRVA)
	}
}

func ParseRelocations(data *[]byte, relocPtr uint32) (map[uint32]bool, map[uint32]string) {

	relocStr := ""
	numRelocs := 0

	relocArray := make(map[uint32]string)
	relocPages := make(map[uint32]bool)

	for dataIndex := relocPtr; dataIndex < relocPtr+4096; dataIndex++ {

		byteArray := [8]byte{(*data)[dataIndex], (*data)[dataIndex+1], (*data)[dataIndex+2], (*data)[dataIndex+3], (*data)[dataIndex+4], (*data)[dataIndex+5], (*data)[dataIndex+6], (*data)[dataIndex+7]}
		byteVal := binary.LittleEndian.Uint64(byteArray[:])

		if byteVal == 0 {
			break
		}

		bytesVA := [4]byte{(*data)[dataIndex], (*data)[dataIndex+1], (*data)[dataIndex+2], (*data)[dataIndex+3]}
		VirtualAddress := binary.LittleEndian.Uint32(bytesVA[:])

		bSizeOfBlock := [4]byte{(*data)[dataIndex+4], (*data)[dataIndex+5], (*data)[dataIndex+6], (*data)[dataIndex+7]}
		SizeOfBlock := binary.LittleEndian.Uint32(bSizeOfBlock[:])

		var numEntry uint32 = 0

		relocStr += "\tPage RVA: " + strconv.FormatUint(uint64(VirtualAddress), 10) + "\n"
		relocStr += "\t-- SizeOfBlock: " + strconv.FormatUint(uint64(SizeOfBlock), 10) + "\n"
		relocStr += "\t-- Num. entries: " + strconv.FormatUint(uint64((SizeOfBlock-8)/2), 10) + "\n\n"

		relocPages[VirtualAddress] = true

		for i := relocPtr + 8; i < relocPtr+SizeOfBlock; i++ {

			offsetType := [2]byte{(*data)[i], (*data)[i+1]}
			offsetVal := binary.LittleEndian.Uint16(offsetType[:])

			typeByte := [1]byte{(*data)[i+1]}
			typeByte[0] = typeByte[0] << 4
			typeByte[0] = typeByte[0] >> 4
			typeVal := uint8(typeByte[0])

			relocStr += "\t-- Entry " + strconv.FormatUint(uint64(numEntry+1), 10) + "\n"
			relocStr += "\tOffset: " + strconv.FormatUint(uint64(i-relocPtr), 10) + "\n"
			relocStr += "\tValue: " + strconv.FormatUint(uint64(offsetVal), 10) + "\n"
			relocStr += "\tType: " + strconv.FormatUint(uint64(typeVal), 10) + "\n\n"

			numEntry++
			i++
		}

		relocArray[VirtualAddress] = relocStr
		relocStr = ""

		dataIndex = relocPtr + SizeOfBlock - 1
		relocPtr = dataIndex + 1
		numRelocs++
	}

	return relocPages, relocArray
}

func DisplayRelocations(relocPages *map[uint32]bool, relocArray *map[uint32]string) {

	fmt.Println("============= Base Relocations =============")

	numEntry := 0

	entryArray := make(map[int]uint32)

	for page, _ := range *relocPages { // loop relocPages
		fmt.Printf("Relocation struct %v (Page RVA: %v)\n", numEntry, page)
		entryArray[numEntry] = page
		numEntry++
	}

	fmt.Printf("Enter relocation struct to view by entry: ")

	var userEntry int

	fmt.Scan(&userEntry)

	if userEntry > numEntry || userEntry < 0 { // if invalid entry, quit
		fmt.Printf("Not a valid entry!\n")
		return
	}

	userStruct := entryArray[userEntry]
	relocStr := (*relocArray)[userStruct]

	fmt.Printf("%v\n", relocStr)
}

func DisplayPE(data *[]byte, bitRange string, hasRichHdr bool) {

	var DosHeader *structs.IMAGE_DOS_HEADER
	var RichHeader *structs.IMAGE_RICH_HEADER
	var FileHeader *structs.IMAGE_FILE_HEADER
	var OptionalHeader64 *structs.IMAGE_OPTIONAL_HEADER64
	var OptionalHeader32 *structs.IMAGE_OPTIONAL_HEADER32

	var sectionTableOffset uint16 = 392
	var Signature string
	var numSections uint16

	dataBuffer := bytes.Buffer{}
	dataBuffer.Write(*data)

	if bitRange == "64bit" {
		if hasRichHdr { // 64-bit program has a Rich header
			ObjPE := structs.PE_RICH_FORMAT64{}

			err := binary.Read(&dataBuffer, binary.LittleEndian, &ObjPE)
			util.CheckError(err)

			DosHeader = &ObjPE.DosHeader
			RichHeader = &ObjPE.RichHeader
			FileHeader = &ObjPE.NTHeader.FileHeader
			OptionalHeader64 = &ObjPE.NTHeader.OptionalHeader
			Signature = string(ObjPE.NTHeader.Signature[:])
			sectionTableOffset = 0x200
			numSections = ObjPE.NTHeader.FileHeader.NumberOfSections

		} else { // 64-bit program does not have a Rich header
			ObjPE := structs.PE_FORMAT64{}

			err := binary.Read(&dataBuffer, binary.LittleEndian, &ObjPE)
			util.CheckError(err)

			DosHeader = &ObjPE.DosHeader
			FileHeader = &ObjPE.NTHeader.FileHeader
			OptionalHeader64 = &ObjPE.NTHeader.OptionalHeader
			Signature = string(ObjPE.NTHeader.Signature[:])
			numSections = ObjPE.NTHeader.FileHeader.NumberOfSections
		}
	} else if bitRange == "32bit" {

		if hasRichHdr {
			ObjPE := structs.PE_RICH_FORMAT32{}

			err := binary.Read(&dataBuffer, binary.LittleEndian, &ObjPE)
			util.CheckError(err)

			DosHeader = &ObjPE.DosHeader
			RichHeader = &ObjPE.RichHeader
			FileHeader = &ObjPE.NTHeader.FileHeader
			OptionalHeader32 = &ObjPE.NTHeader.OptionalHeader
			Signature = string(ObjPE.NTHeader.Signature[:])
			sectionTableOffset = 0x1F0
			numSections = ObjPE.NTHeader.FileHeader.NumberOfSections
		} else {
			ObjPE := structs.PE_FORMAT32{}
			DosHeader = &ObjPE.DosHeader
			FileHeader = &ObjPE.NTHeader.FileHeader
			OptionalHeader32 = &ObjPE.NTHeader.OptionalHeader
			Signature = string(ObjPE.NTHeader.Signature[:])

			err := binary.Read(&dataBuffer, binary.LittleEndian, &ObjPE)
			util.CheckError(err)

			Signature = string(ObjPE.NTHeader.Signature[:])
			numSections = ObjPE.NTHeader.FileHeader.NumberOfSections
		}
	} else {
		fmt.Printf("Cannot read this type of file!")
		return
	}

	DisplayDosHeader(DosHeader)

	if hasRichHdr {
		DisplayRichHeader(RichHeader)
	}

	fmt.Println("============= NT Header =============")
	fmt.Printf("\tSignature: %v\n", Signature)

	DisplayFileHeader(FileHeader)

	var importVA uint32
	var importSize uint32

	if bitRange != "64bit" {
		DisplayOptionalHeader32(OptionalHeader32, bitRange)
		importVA, importSize = DisplayDataDirectories(&OptionalHeader32.DataDirectory, bitRange)
	} else {
		DisplayOptionalHeader64(OptionalHeader64, bitRange)
		importVA, importSize = DisplayDataDirectories(&OptionalHeader64.DataDirectory, bitRange)
	}

	sectionHeaderArray := ParseSectionTable(data, sectionTableOffset, numSections)
	importSectionAddr, importSectionPtr, importSectionName, relocPtr, relocSize := DisplaySectionTable(sectionHeaderArray, numSections, importVA)

	if importSize > 0 {
		importArray, numImports := ParseImportDirectory(data, importVA, importSectionAddr, importSize, importSectionPtr, importSectionName)
		DisplayImportDirectory(bitRange, data, importArray, importSectionAddr, importSectionPtr, numImports)
	}

	if relocSize > 0 {
		var wishToParseRelocations string

		fmt.Printf("Do you wish to parse relocations? (y/n): ")
		fmt.Scan(&wishToParseRelocations)

		var relocPages map[uint32]bool
		var relocArray map[uint32]string

		if wishToParseRelocations == "y" {
			fmt.Printf("Parsing base relocations..\n")
			relocPages, relocArray = ParseRelocations(data, relocPtr)

			var displayRelocAgain string = "y"

			for displayRelocAgain == "y" {
				DisplayRelocations(&relocPages, &relocArray)
				fmt.Printf("Do you wish to display base relocations again? (y/n): ")
				fmt.Scan(&displayRelocAgain)
			}
		}
	}
}

func main() {

	var pe string

	for {
		fmt.Println("== PE Parser ==")
		fmt.Print("Enter file name of PE: ")

		fmt.Scan(&pe)

		data, err := os.ReadFile(pe)
		util.CheckError(err)

		var hasRichHdr bool = util.HasRichHeader(&data)
		var bitRange string

		if hasRichHdr { // if pe has a rich header
			bitRange = util.DetermineBitRange([2]byte{data[272], data[273]})
			util.DecryptRichHeader(&data)
		} else {
			bitRange = util.DetermineBitRange([2]byte{data[152], data[153]})
		}

		DisplayPE(&data, bitRange, hasRichHdr)

	}
}
