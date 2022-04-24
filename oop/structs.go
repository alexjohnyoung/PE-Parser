package structs

type PE_RICH_FORMAT64 struct {
	DosHeader  IMAGE_DOS_HEADER
	DosStub    [64]byte
	RichHeader IMAGE_RICH_HEADER
	NTHeader   NTHeader64
}

type PE_FORMAT64 struct {
	DosHeader IMAGE_DOS_HEADER
	DosStub   [64]byte
	NTHeader  NTHeader64
}

type PE_FORMAT32 struct {
	DosHeader IMAGE_DOS_HEADER
	DosStub   [64]byte
	NTHeader  NTHeader32
}

type PE_RICH_FORMAT32 struct {
	DosHeader  IMAGE_DOS_HEADER
	DosStub    [64]byte
	RichHeader IMAGE_RICH_HEADER
	NTHeader   NTHeader32
}

type IMAGE_DOS_HEADER struct {
	E_magic    [2]byte
	E_cblp     uint16
	E_cp       uint16
	E_crlc     uint16
	E_cparhdr  uint16
	E_minalloc uint16
	E_maxalloc uint16
	E_ss       uint16
	E_sp       uint16
	E_csum     uint16
	E_ip       uint16
	E_cs       uint16
	E_lfarlc   uint16
	E_ovno     uint16
	E_res      uint64
	E_oemid    uint16
	E_oeminfo  uint16
	E_res2     [20]byte
	E_lfanew   int32
}

type PRODITEM struct {
	Buildid uint16
	Prodid  uint16
	Count   uint32
}

type IMAGE_RICH_HEADER struct {
	DansID             [4]byte
	ChecksumedPadding1 [4]byte
	ChecksumedPadding2 PRODITEM
	CompID1            PRODITEM
	CompID2            PRODITEM
	CompID3            PRODITEM
	CompID4            PRODITEM
	CompID5            PRODITEM
	CompID6            PRODITEM
	CompID7            PRODITEM
	CompID8            PRODITEM
	CompID9            PRODITEM
	CompID10           PRODITEM
	RichIdentifier     [4]byte
	Checksum           uint32
	ChecksumedPadding3 PRODITEM
	ChecksumedPadding4 PRODITEM
}

type NTHeader32 struct {
	Signature      [4]byte
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER32
}

type NTHeader64 struct {
	Signature      [4]byte
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}

type IMAGE_FILE_HEADER struct {
	ImageMachine         uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress uint32
	Size           uint32
}

type IMAGE_OPTIONAL_HEADER32 struct {
	Magic                       uint16
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       uint16
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_SECTION_HEADER struct { // 40
	Name                 [8]byte
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLinenumbers uint32
	NumberOfRelocations  uint16
	NumberOfLinenumbers  uint16
	Characteristics      uint32
}

type IMAGE_IMPORT_DESCRIPTOR struct {
	OriginalFirstThunk uint32
	TimeDateStamp      int32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
}

type IMAGE_THUNK_DATA32 struct {
	HINT_NAME_TABLE uint32
}

type IMAGE_THUNK_DATA64 struct {
	HINT_NAME_TABLE   uint32
	ORDINAL_NAME_FLAG uint32
}

type IMAGE_IMPORT_BY_NAME struct {
	Hint [2]byte
	Name [1]byte
}

type IMAGE_BASE_RELOCATION struct {
	VirtualAddress uint32
	SizeOfBlock    uint32
}

type IMAGE_BASE_RELOCATION_ENTRY struct {
	PageVA     uint32
	PageSize   uint32
	PageEntry  uint32
	PageOffset uint32
	PageValue  uint16
	PageType   [1]byte
}
