package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	structs "pestructs/pe/oop"
	"strconv"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetImageMachine(ImageMachine uint16) (strImageMachine string) {

	switch ImageMachine {
	case 0:
		strImageMachine = "Unknown"
	case 467:
		strImageMachine = "Matsushita AM33"
	case 34404:
		strImageMachine = "x64"
	case 448:
		strImageMachine = "ARM little endian"
	case 43620:
		strImageMachine = "ARM64 little endian"
	case 452:
		strImageMachine = "ARM Thumb-2 little endian"
	case 3772:
		strImageMachine = "EFI byte code"
	case 332:
		strImageMachine = "Intel 386 or later processors and compatible processors"
	case 512:
		strImageMachine = "Intel Itanium processor family"
	case 25138:
		strImageMachine = "LoongArch 32-bit processor family"
	case 25188:
		strImageMachine = "LoongArch 64-bit processor family"
	case 36929:
		strImageMachine = "Mitsubishi M32R little endian"
	case 614:
		strImageMachine = "MIPS16"
	case 870:
		strImageMachine = "MIPS with FPU"
	case 1126:
		strImageMachine = "MIPS16 with FPU"
	case 496:
		strImageMachine = "Power PC little endian"
	case 497:
		strImageMachine = "Power PC with floating point support"
	case 358:
		strImageMachine = "MIPS little endian"
	case 20530:
		strImageMachine = "RISC-V 32-bit address space"
	case 20580:
		strImageMachine = "RISC-V 64-bit address space"
	case 20776:
		strImageMachine = "RISC-V 128-bit address space"
	case 418:
		strImageMachine = "Hitachi SH3"
	case 419:
		strImageMachine = "Hitachi SH3 DSP"
	case 422:
		strImageMachine = "Hitachi SH4"
	case 424:
		strImageMachine = "Hitachi SH5"
	case 450:
		strImageMachine = "Thumb"
	case 361:
		strImageMachine = "MIPS little-endian WCE v2"
	}

	return strImageMachine
}

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

func DetermineProdId(Prodid uint16) string {
	switch Prodid {
	case 0x0000:
		return "Unknown"
	case 0x0001:
		return "Import0"
	case 0x0002:
		return "Linker510"
	case 0x0003:
		return "Cvtomf510"
	case 0x0004:
		return "Linker600"
	case 0x0005:
		return "Cvtomf600"
	case 0x0006:
		return "Cvtres500"
	case 0x0007:
		return "Utc11_Basic"
	case 0x0008:
		return "Utc11_C"
	case 0x0009:
		return "Utc12_Basic"
	case 0x000a:
		return "Utc12_C"
	case 0x000b:
		return "Utc12_CPP"
	case 0x000c:
		return "AliasObj60"
	case 0x000d:
		return "VisualBasic60"
	case 0x000e:
		return "Masm613"
	case 0x000f:
		return "Masm710"
	case 0x0010:
		return "Linker511"
	case 0x0011:
		return "Cvtomf511"
	case 0x0012:
		return "Masm614"
	case 0x0013:
		return "Linker512"
	case 0x0014:
		return "Cvtomf512"
	case 0x0015:
		return "Utc12_C_Std"
	case 0x0016:
		return "Utc12_CPP_Std"
	case 0x0017:
		return "Utc12_C_Book"
	case 0x0018:
		return "Utc12_CPP_Book"
	case 0x0019:
		return "Implib700"
	case 0x001a:
		return "Cvtomf700"
	case 0x001b:
		return "Utc13_Basic"
	case 0x001c:
		return "Utc13_C"
	case 0x001d:
		return "Utc13_CPP"
	case 0x001e:
		return "Linker610"
	case 0x001f:
		return "Cvtomf610"
	case 0x0020:
		return "Linker601"
	case 0x0021:
		return "Cvtomf601"
	case 0x0022:
		return "Utc12_1_Basic"
	case 0x0023:
		return "Utc12_1_C"
	case 0x0024:
		return "Utc12_1_CPP"
	case 0x0025:
		return "Linker620"
	case 0x0026:
		return "Cvtomf620"
	case 0x0027:
		return "AliasObj70"
	case 0x0028:
		return "Linker621"
	case 0x0029:
		return "Cvtomf621"
	case 0x002a:
		return "Masm615"
	case 0x002b:
		return "Utc13_LTCG_C"
	case 0x002c:
		return "Utc13_LTCG_CPP"
	case 0x002d:
		return "Masm620"
	case 0x002e:
		return "ILAsm100"
	case 0x002f:
		return "Utc12_2_Basic"
	case 0x0030:
		return "Utc12_2_C"
	case 0x0031:
		return "Utc12_2_CPP"
	case 0x0032:
		return "Utc12_2_C_Std"
	case 0x0033:
		return "Utc12_2_CPP_Std"
	case 0x0034:
		return "Utc12_2_C_Book"
	case 0x0035:
		return "Utc12_2_CPP_Book"
	case 0x0036:
		return "Implib622"
	case 0x0037:
		return "Cvtomf622"
	case 0x0038:
		return "Cvtres501"
	case 0x0039:
		return "Utc13_C_Std"
	case 0x003a:
		return "Utc13_CPP_Std"
	case 0x003b:
		return "Cvtpgd1300"
	case 0x003c:
		return "Linker622"
	case 0x003d:
		return "Linker700"
	case 0x003e:
		return "Export622"
	case 0x003f:
		return "Export700"
	case 0x0040:
		return "Masm700"
	case 0x0041:
		return "Utc13_POGO_I_C"
	case 0x0042:
		return "Utc13_POGO_I_CPP"
	case 0x0043:
		return "Utc13_POGO_O_C"
	case 0x0044:
		return "Utc13_POGO_O_CPP"
	case 0x0045:
		return "Cvtres700"
	case 0x0046:
		return "Cvtres710p"
	case 0x0047:
		return "Linker710p"
	case 0x0048:
		return "Cvtomf710p"
	case 0x0049:
		return "Export710p"
	case 0x004a:
		return "Implib710p"
	case 0x004b:
		return "Masm710p"
	case 0x004c:
		return "Utc1310p_C"
	case 0x004d:
		return "Utc1310p_CPP"
	case 0x004e:
		return "Utc1310p_C_Std"
	case 0x004f:
		return "Utc1310p_CPP_Std"
	case 0x0050:
		return "Utc1310p_LTCG_C"
	case 0x0051:
		return "Utc1310p_LTCG_CPP"
	case 0x0052:
		return "Utc1310p_POGO_I_C"
	case 0x0053:
		return "Utc1310p_POGO_I_CPP"
	case 0x0054:
		return "Utc1310p_POGO_O_C"
	case 0x0055:
		return "Utc1310p_POGO_O_CPP"
	case 0x0056:
		return "Linker624"
	case 0x0057:
		return "Cvtomf624"
	case 0x0058:
		return "Export624"
	case 0x0059:
		return "Implib624"
	case 0x005a:
		return "Linker710"
	case 0x005b:
		return "Cvtomf710"
	case 0x005c:
		return "Export710"
	case 0x005d:
		return "Implib710"
	case 0x005e:
		return "Cvtres710"
	case 0x005f:
		return "Utc1310_C"
	case 0x0060:
		return "Utc1310_CPP"
	case 0x0061:
		return "Utc1310_C_Std"
	case 0x0062:
		return "Utc1310_CPP_Std"
	case 0x0063:
		return "Utc1310_LTCG_C"
	case 0x0064:
		return "Utc1310_LTCG_CPP"
	case 0x0065:
		return "Utc1310_POGO_I_C"
	case 0x0066:
		return "Utc1310_POGO_I_CPP"
	case 0x0067:
		return "Utc1310_POGO_O_C"
	case 0x0068:
		return "Utc1310_POGO_O_CPP"
	case 0x0069:
		return "AliasObj710"
	case 0x006a:
		return "AliasObj710p"
	case 0x006b:
		return "Cvtpgd1310"
	case 0x006c:
		return "Cvtpgd1310p"
	case 0x006d:
		return "Utc1400_C"
	case 0x006e:
		return "Utc1400_CPP"
	case 0x006f:
		return "Utc1400_C_Std"
	case 0x0070:
		return "Utc1400_CPP_Std"
	case 0x0071:
		return "Utc1400_LTCG_C"
	case 0x0072:
		return "Utc1400_LTCG_CPP"
	case 0x0073:
		return "Utc1400_POGO_I_C"
	case 0x0074:
		return "Utc1400_POGO_I_CPP"
	case 0x0075:
		return "Utc1400_POGO_O_C"
	case 0x0076:
		return "Utc1400_POGO_O_CPP"
	case 0x0077:
		return "Cvtpgd1400"
	case 0x0078:
		return "Linker800"
	case 0x0079:
		return "Cvtomf800"
	case 0x007a:
		return "Export800"
	case 0x007b:
		return "Implib800"
	case 0x007c:
		return "Cvtres800"
	case 0x007d:
		return "Masm800"
	case 0x007e:
		return "AliasObj800"
	case 0x007f:
		return "PhoenixPrerelease"
	case 0x0080:
		return "Utc1400_CVTCIL_C"
	case 0x0081:
		return "Utc1400_CVTCIL_CPP"
	case 0x0082:
		return "Utc1400_LTCG_MSIL"
	case 0x0083:
		return "Utc1500_C"
	case 0x0084:
		return "Utc1500_CPP"
	case 0x0085:
		return "Utc1500_C_Std"
	case 0x0086:
		return "Utc1500_CPP_Std"
	case 0x0087:
		return "Utc1500_CVTCIL_C"
	case 0x0088:
		return "Utc1500_CVTCIL_CPP"
	case 0x0089:
		return "Utc1500_LTCG_C"
	case 0x008a:
		return "Utc1500_LTCG_CPP"
	case 0x008b:
		return "Utc1500_LTCG_MSIL"
	case 0x008c:
		return "Utc1500_POGO_I_C"
	case 0x008d:
		return "Utc1500_POGO_I_CPP"
	case 0x008e:
		return "Utc1500_POGO_O_C"
	case 0x008f:
		return "Utc1500_POGO_O_CPP"
	case 0x0090:
		return "Cvtpgd1500"
	case 0x0091:
		return "Linker900"
	case 0x0092:
		return "Export900"
	case 0x0093:
		return "Implib900"
	case 0x0094:
		return "Cvtres900"
	case 0x0095:
		return "Masm900"
	case 0x0096:
		return "AliasObj900"
	case 0x0097:
		return "Resource"
	case 0x0098:
		return "AliasObj1000"
	case 0x0099:
		return "Cvtpgd1600"
	case 0x009a:
		return "Cvtres1000"
	case 0x009b:
		return "Export1000"
	case 0x009c:
		return "Implib1000"
	case 0x009d:
		return "Linker1000"
	case 0x009e:
		return "Masm1000"
	case 0x009f:
		return "Phx1600_C"
	case 0x00a0:
		return "Phx1600_CPP"
	case 0x00a1:
		return "Phx1600_CVTCIL_C"
	case 0x00a2:
		return "Phx1600_CVTCIL_CPP"
	case 0x00a3:
		return "Phx1600_LTCG_C"
	case 0x00a4:
		return "Phx1600_LTCG_CPP"
	case 0x00a5:
		return "Phx1600_LTCG_MSIL"
	case 0x00a6:
		return "Phx1600_POGO_I_C"
	case 0x00a7:
		return "Phx1600_POGO_I_CPP"
	case 0x00a8:
		return "Phx1600_POGO_O_C"
	case 0x00a9:
		return "Phx1600_POGO_O_CPP"
	case 0x00aa:
		return "Utc1600_C"
	case 0x00ab:
		return "Utc1600_CPP"
	case 0x00ac:
		return "Utc1600_CVTCIL_C"
	case 0x00ad:
		return "Utc1600_CVTCIL_CPP"
	case 0x00ae:
		return "Utc1600_LTCG_C"
	case 0x00af:
		return "Utc1600_LTCG_CPP"
	case 0x00b0:
		return "Utc1600_LTCG_MSIL"
	case 0x00b1:
		return "Utc1600_POGO_I_C"
	case 0x00b2:
		return "Utc1600_POGO_I_CPP"
	case 0x00b3:
		return "Utc1600_POGO_O_C"
	case 0x00b4:
		return "Utc1600_POGO_O_CPP"
	case 0x00b5:
		return "AliasObj1010"
	case 0x00b6:
		return "Cvtpgd1610"
	case 0x00b7:
		return "Cvtres1010"
	case 0x00b8:
		return "Export1010"
	case 0x00b9:
		return "Implib1010"
	case 0x00ba:
		return "Linker1010"
	case 0x00bb:
		return "Masm1010"
	case 0x00bc:
		return "Utc1610_C"
	case 0x00bd:
		return "Utc1610_CPP"
	case 0x00be:
		return "Utc1610_CVTCIL_C"
	case 0x00bf:
		return "Utc1610_CVTCIL_CPP"
	case 0x00c0:
		return "Utc1610_LTCG_C"
	case 0x00c1:
		return "Utc1610_LTCG_CPP"
	case 0x00c2:
		return "Utc1610_LTCG_MSIL"
	case 0x00c3:
		return "Utc1610_POGO_I_C"
	case 0x00c4:
		return "Utc1610_POGO_I_CPP"
	case 0x00c5:
		return "Utc1610_POGO_O_C"
	case 0x00c6:
		return "Utc1610_POGO_O_CPP"
	case 0x00c7:
		return "AliasObj1100"
	case 0x00c8:
		return "Cvtpgd1700"
	case 0x00c9:
		return "Cvtres1100"
	case 0x00ca:
		return "Export1100"
	case 0x00cb:
		return "Implib1100"
	case 0x00cc:
		return "Linker1100"
	case 0x00cd:
		return "Masm1100"
	case 0x00ce:
		return "Utc1700_C"
	case 0x00cf:
		return "Utc1700_CPP"
	case 0x00d0:
		return "Utc1700_CVTCIL_C"
	case 0x00d1:
		return "Utc1700_CVTCIL_CPP"
	case 0x00d2:
		return "Utc1700_LTCG_C"
	case 0x00d3:
		return "Utc1700_LTCG_CPP"
	case 0x00d4:
		return "Utc1700_LTCG_MSIL"
	case 0x00d5:
		return "Utc1700_POGO_I_C"
	case 0x00d6:
		return "Utc1700_POGO_I_CPP"
	case 0x00d7:
		return "Utc1700_POGO_O_C"
	case 0x00d8:
		return "Utc1700_POGO_O_CPP"
	case 0x00d9:
		return "AliasObj1200"
	case 0x00da:
		return "Cvtpgd1800"
	case 0x00db:
		return "Cvtres1200"
	case 0x00dc:
		return "Export1200"
	case 0x00dd:
		return "Implib1200"
	case 0x00de:
		return "Linker1200"
	case 0x00df:
		return "Masm1200"
	case 0x00e0:
		return "Utc1800_C"
	case 0x00e1:
		return "Utc1800_CPP"
	case 0x00e2:
		return "Utc1800_CVTCIL_C"
	case 0x00e3:
		return "Utc1800_CVTCIL_CPP"
	case 0x00e4:
		return "Utc1800_LTCG_C"
	case 0x00e5:
		return "Utc1800_LTCG_CPP"
	case 0x00e6:
		return "Utc1800_LTCG_MSIL"
	case 0x00e7:
		return "Utc1800_POGO_I_C"
	case 0x00e8:
		return "Utc1800_POGO_I_CPP"
	case 0x00e9:
		return "Utc1800_POGO_O_C"
	case 0x00ea:
		return "Utc1800_POGO_O_CPP"
	case 0x00eb:
		return "AliasObj1210"
	case 0x00ec:
		return "Cvtpgd1810"
	case 0x00ed:
		return "Cvtres1210"
	case 0x00ee:
		return "Export1210"
	case 0x00ef:
		return "Implib1210"
	case 0x00f0:
		return "Linker1210"
	case 0x00f1:
		return "Masm1210"
	case 0x00f2:
		return "Utc1810_C"
	case 0x00f3:
		return "Utc1810_CPP"
	case 0x00f4:
		return "Utc1810_CVTCIL_C"
	case 0x00f5:
		return "Utc1810_CVTCIL_CPP"
	case 0x00f6:
		return "Utc1810_LTCG_C"
	case 0x00f7:
		return "Utc1810_LTCG_CPP"
	case 0x00f8:
		return "Utc1810_LTCG_MSIL"
	case 0x00f9:
		return "Utc1810_POGO_I_C"
	case 0x00fa:
		return "Utc1810_POGO_I_CPP"
	case 0x00fb:
		return "Utc1810_POGO_O_C"
	case 0x00fc:
		return "Utc1810_POGO_O_CPP"
	case 0x00fd:
		return "AliasObj1400"
	case 0x00fe:
		return "Cvtpgd1900"
	case 0x00ff:
		return "Cvtres1400"
	case 0x0100:
		return "Export1400"
	case 0x0101:
		return "Implib1400"
	case 0x0102:
		return "Linker1400"
	case 0x0103:
		return "Masm1400"
	case 0x0104:
		return "Utc1900_C"
	case 0x0105:
		return "Utc1900_CPP"
	case 0x0106:
		return "Utc1900_CVTCIL_C"
	case 0x0107:
		return "Utc1900_CVTCIL_CPP"
	case 0x0108:
		return "Utc1900_LTCG_C"
	case 0x0109:
		return "Utc1900_LTCG_CPP"
	case 0x010a:
		return "Utc1900_LTCG_MSIL"
	case 0x010b:
		return ": 'Utc1900_POGO_I_C"
	case 0x010c:
		return "Utc1900_POGO_I_CPP"
	case 0x010d:
		return "Utc1900_POGO_O_C"
	case 0x010e:
		return "Utc1900_POGO_O_CPP"
	}

	return "Unknown"
}

func DecryptAndDisplayRichHeader(RichHeader *structs.IMAGE_RICH_HEADER) {
	fmt.Println("============= Rich Header =============")

	fmt.Printf("\tDanS ID: %v\n", string(RichHeader.DansID[:]))

	fmt.Printf("\tComp ID 1\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID1.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID1.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID1.Count)

	fmt.Printf("\tComp ID 2\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID2.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID2.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID2.Count)

	fmt.Printf("\tComp ID 3\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID3.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID3.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID3.Count)

	fmt.Printf("\tComp ID 4\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID4.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID4.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID4.Count)

	fmt.Printf("\tComp ID 5\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID5.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID5.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID5.Count)

	fmt.Printf("\tComp ID 6\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID6.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID6.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID6.Count)

	fmt.Printf("\tComp ID 7\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID7.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID7.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID7.Count)

	fmt.Printf("\tComp ID 8\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID8.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID8.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID8.Count)

	fmt.Printf("\tComp ID 9\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID9.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID9.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID9.Count)

	fmt.Printf("\tComp ID 10\n")
	fmt.Printf("\t\tProduct Id: %v\n", DetermineProdId(RichHeader.CompID10.Prodid))
	fmt.Printf("\t\tBuild Id: %v\n", RichHeader.CompID10.Buildid)
	fmt.Printf("\t\tCount: %v\n", RichHeader.CompID10.Count)

	fmt.Printf("\tRich Identifier: %v\n", string(RichHeader.RichIdentifier[:]))
	fmt.Printf("\tChecksum: %v\n", RichHeader.Checksum)
}

func DisplayFileHeader(FileHeader *structs.IMAGE_FILE_HEADER) {

	ImageMachine := GetImageMachine(FileHeader.ImageMachine)

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

func GetDataDirectoryName32(offset int) string {

	var directoryName string

	switch offset {
	case 96:
		directoryName = "export table"
	case 104:
		directoryName = "import table"
	case 112:
		directoryName = "resource table"
	case 120:
		directoryName = "exception table"
	case 128:
		directoryName = "certificate table"
	case 136:
		directoryName = "base relocation table"
	case 144:
		directoryName = "debugging information"
	case 152:
		directoryName = "architecture-specific data"
	case 160:
		directoryName = "global pointer register"
	case 168:
		directoryName = "thread local storage table"
	case 176:
		directoryName = "load configuration table"
	case 184:
		directoryName = "bound import table"
	case 192:
		directoryName = "import address table"
	case 200:
		directoryName = "delay import descriptor"
	case 208:
		directoryName = "CLR header"
	case 216:
		directoryName = "Reserved"
	}

	return directoryName
}

func GetDataDirectoryName64(offset int) string {

	var directoryName string

	switch offset {
	case 112:
		directoryName = "export table"
	case 120:
		directoryName = "import table"
	case 128:
		directoryName = "resource table"
	case 136:
		directoryName = "exception table"
	case 144:
		directoryName = "certificate table"
	case 152:
		directoryName = "base relocation table"
	case 160:
		directoryName = "debugging information"
	case 168:
		directoryName = "architecture-specific data"
	case 176:
		directoryName = "global pointer register"
	case 184:
		directoryName = "thread local storage table"
	case 192:
		directoryName = "load configuration"
	case 200:
		directoryName = "bound import table"
	case 208:
		directoryName = "import address table"
	case 216:
		directoryName = "delay import descriptor"
	case 224:
		directoryName = "CLR header"
	case 232:
		directoryName = "Reserved"
	}

	return directoryName
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
			directoryName = GetDataDirectoryName64(nOffset)
		} else {
			directoryName = GetDataDirectoryName32(nOffset)
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
	CheckError(err)

	return &sectionHeaderArray
}

func ParseImportDirectory(data *[]byte, importVA uint32, importSectionAddr uint32, importSize uint32, importSectionPtr uint32, importSectionName string) (*[]structs.IMAGE_IMPORT_DESCRIPTOR, uint32) {

	importTableRVA := importVA - importSectionAddr
	importTableLocation := importSectionPtr + importTableRVA
	fmt.Println("============= Import Directory =============")
	fmt.Printf("\tLocated in %v (VA: %v)\n", importSectionName, importSectionAddr)
	fmt.Printf("\tPointerToRawData: %v\n\tFile Offset: %v\n\n", importSectionPtr, importTableRVA)

	var numImports uint32 = 0
	var importEndStruct uint32 = 0

	numBytes := 0
	numZeroes := 0

	for i := importTableLocation; i < importTableLocation+200; i++ {
		var byteIndex [1]byte = [1]byte{(*data)[i]}

		numBytes += 1

		if byteIndex[0] == 0 {
			numZeroes += 1

			if numZeroes == 20 {
				importEndStruct = i
				numBytes = 0
				break
			}
		} else {
			numZeroes = 0
		}

		if numBytes == 20 {
			numImports += 1
			numBytes = 0
			numZeroes = 0
		}

	}

	importArray := make([]structs.IMAGE_IMPORT_DESCRIPTOR, numImports)

	fmt.Printf("\tNum. imports: %v\n", numImports)

	importBuffer := bytes.Buffer{}
	importBuffer.Write((*data)[importTableLocation:importEndStruct])

	fmt.Printf("\tImport Table VA: %v\n", importVA)
	fmt.Printf("\tImport Table Size: %v\n\n", importSize)

	err := binary.Read(&importBuffer, binary.LittleEndian, &importArray)
	CheckError(err)

	return &importArray, numImports
}

func GetNumCharacters(data *[]byte, startPos uint32) int {

	var numChars int

	for byteIndex := startPos; byteIndex < byteIndex+64; byteIndex++ {
		if (*data)[byteIndex] == 0 {
			break
		} else {
			numChars += 1
		}
	}

	return numChars
}

func DisplayImportDirectory(bitRange string, data *[]byte, importArray *[]structs.IMAGE_IMPORT_DESCRIPTOR, importSectionAddr uint32, importSectionPtr, numImports uint32) {

	for nImport := uint32(0); nImport < numImports; nImport++ {

		importObj := (*importArray)[nImport]

		var nameRVA uint32 = importObj.Name - importSectionAddr
		var nameLocation uint32 = nameRVA + importSectionPtr
		var numChars int = GetNumCharacters(data, nameLocation)

		fmt.Printf("\tImport %v\n", nImport+1)
		fmt.Printf("\t_____________________________\n")
		fmt.Printf("\t| OriginalFirstThunk: %v\n", importObj.OriginalFirstThunk)
		fmt.Printf("\t| FirstThunk: %v\n", importObj.FirstThunk)
		fmt.Printf("\t| TimeDateStamp: %v (%v)\n", importObj.TimeDateStamp, GetBoundStatus(importObj.TimeDateStamp))
		fmt.Printf("\t| ForwarderChain: %v\n", importObj.ForwarderChain)
		fmt.Printf("\t| Name: %v\n", string((*data)[nameLocation:nameLocation+uint32(numChars)]))

		importTableRVA := importObj.OriginalFirstThunk - importSectionAddr
		importTableOffset := importSectionPtr + importTableRVA

		numZeroes := 0
		numBytes := 0

		var funcIndex uint32
		var numFuncs uint32 = 0

		if bitRange == "64bit" {

			// get number of funcs
			for funcIndex = importTableOffset; funcIndex < funcIndex+364; funcIndex++ {
				var byteIndex [1]byte = [1]byte{(*data)[funcIndex]}
				if byteIndex[0] == 0 {
					numZeroes += 1
					numBytes += 1
				} else {
					numZeroes = 0
					numBytes += 1
				}

				if numZeroes >= 8 {
					numBytes = 0

					break
				}
				if numBytes >= 8 {
					numBytes = 0
					numFuncs++
				}
			}

			importTableArray := make([]structs.IMAGE_THUNK_DATA64, numFuncs)

			bILTBuffer := bytes.Buffer{}
			bILTBuffer.Write((*data)[importTableOffset : importTableOffset+(8*numFuncs)])

			err := binary.Read(&bILTBuffer, binary.LittleEndian, &importTableArray)
			CheckError(err)

			fmt.Printf("\t\tImage Lookup Table (%v)\n", numFuncs)
			for funcIndex = 0; funcIndex < numFuncs; funcIndex++ {
				ordinalFlagByte := importTableArray[funcIndex].ORDINAL_NAME_FLAG

				fmt.Printf("\t\t\tFunction Import %v\n", funcIndex+1)
				if ordinalFlagByte == 0 { // rva

					if importTableArray[funcIndex].HINT_NAME_TABLE > 0 {
						hintTableRVA := importTableArray[funcIndex].HINT_NAME_TABLE + 2
						hintTableOffset := hintTableRVA - importSectionAddr
						hintTableLocation := importSectionPtr + hintTableOffset

						numChars = GetNumCharacters(data, hintTableLocation)
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
		} else { // 32 bit
			// get number of funcs
			for funcIndex = importTableOffset; funcIndex < funcIndex+364; funcIndex++ {
				var byteIndex [1]byte = [1]byte{(*data)[funcIndex]}
				if byteIndex[0] == 0 {
					numZeroes += 1
					numBytes += 1
				} else {
					numZeroes = 0
					numBytes += 1
				}

				if numZeroes >= 4 {
					numBytes = 0

					break
				}
				if numBytes >= 4 {
					numBytes = 0
					numFuncs++
				}
			}

			importTableArray := make([]structs.IMAGE_THUNK_DATA32, numFuncs)

			bILTBuffer := bytes.Buffer{}
			bILTBuffer.Write((*data)[importTableOffset : importTableOffset+(8*numFuncs)])

			err := binary.Read(&bILTBuffer, binary.LittleEndian, &importTableArray)
			CheckError(err)

			fmt.Printf("\t\tImage Lookup Table (%v)\n", numFuncs)
			for funcIndex = 0; funcIndex < numFuncs; funcIndex++ {
				ordinalFlagByte := importTableArray[funcIndex].HINT_NAME_TABLE << 31
				fmt.Printf("\t\t\tFunction Import %v\n", funcIndex+1)
				if ordinalFlagByte == 0 { // rva

					if importTableArray[funcIndex].HINT_NAME_TABLE > 0 {
						hintTableRVA := importTableArray[funcIndex].HINT_NAME_TABLE + 2
						hintTableOffset := hintTableRVA - importSectionAddr
						hintTableLocation := importSectionPtr + hintTableOffset

						numChars = GetNumCharacters(data, hintTableLocation)
						funcName := string((*data)[hintTableLocation : hintTableLocation+uint32(numChars)])

						fmt.Printf("\t\t\t\tName: %v\n", funcName)
						fmt.Printf("\t\t\t\t\t| HINT_NAME_TABLE\n")
						fmt.Printf("\t\t\t\t\t| RVA: %v\n", hintTableRVA)
						fmt.Printf("\t\t\t\t\t| Physical Offset: %v\n", hintTableOffset)
					} else {
						fmt.Printf("\t\t\tInvalid entry (HINT_NAME_TABLE == 0)\n")
					}
				} else { // ordinal
					var ordinalVal uint16 = uint16(importTableArray[funcIndex].HINT_NAME_TABLE)
					fmt.Printf("\t\t\tOrdinal: %v\n", ordinalVal)
				}
				fmt.Printf("\t\t\t___________________________________________________\n")
			}
		}
	}
}

func GetBoundStatus(TimeDateStamp int32) string {
	if TimeDateStamp == -1 {
		return "Bound"
	}

	return "Not Bound"
}

func DetermineBitRange(magicBytes [2]byte) string {

	var bitRange string = ""

	magicValue := binary.LittleEndian.Uint16(magicBytes[:])

	switch magicValue {
	case 263:
		bitRange = "ROM"
	case 267:
		bitRange = "32bit"
	case 523:
		bitRange = "64bit"
	}

	return bitRange
}

func HasRichHeader(data *[]byte) bool {

	richIdentifier := [4]byte{(*data)[224], (*data)[225], (*data)[226], (*data)[227]}

	return string(richIdentifier[:]) == "Rich"
}

func DecryptRichHeader(data *[]byte) {

	xorBytes := [4]byte{(*data)[228], (*data)[229], (*data)[230], (*data)[231]}
	xorKey := binary.LittleEndian.Uint32(xorBytes[:])

	for i := 128; i < 224; i++ {
		richBytes := [4]byte{(*data)[i], (*data)[i+1], (*data)[i+2], (*data)[i+3]}
		richIndex := binary.LittleEndian.Uint32(richBytes[:])
		richIndex = richIndex ^ xorKey

		var encodedIndex [4]byte
		binary.LittleEndian.PutUint32(encodedIndex[:], richIndex)

		(*data)[i] = encodedIndex[0]
		(*data)[i+1] = encodedIndex[1]
		(*data)[i+2] = encodedIndex[2]
		(*data)[i+3] = encodedIndex[3]

		i += 3
	}
}

func DisplayRelocations(relocPages map[uint32]bool, relocArray map[uint32]string) {

	fmt.Println("============= Base Relocations =============")

	numEntry := 0

	entryArray := make(map[int]uint32)

	for page, _ := range relocPages {
		fmt.Printf("Relocation struct %v (Page RVA: %v)\n", numEntry, page)
		entryArray[numEntry] = page
		numEntry++
	}

	fmt.Printf("Enter relocation struct to view by entry: ")

	var userEntry int

	fmt.Scan(&userEntry)

	if userEntry > numEntry || userEntry < 0 {
		fmt.Printf("Not a valid entry!\n")
		return
	}

	userStruct := entryArray[userEntry]
	relocStr := relocArray[userStruct]

	fmt.Printf("%v\n", relocStr)

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
			CheckError(err)

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
			CheckError(err)

			DosHeader = &ObjPE.DosHeader
			FileHeader = &ObjPE.NTHeader.FileHeader
			OptionalHeader64 = &ObjPE.NTHeader.OptionalHeader
			Signature = string(ObjPE.NTHeader.Signature[:])
			numSections = ObjPE.NTHeader.FileHeader.NumberOfSections
		}
	} else if bitRange == "32bit" { // 32 bit or ROM

		if hasRichHdr {
			ObjPE := structs.PE_RICH_FORMAT32{}

			err := binary.Read(&dataBuffer, binary.LittleEndian, &ObjPE)
			CheckError(err)

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
			CheckError(err)

			Signature = string(ObjPE.NTHeader.Signature[:])
			numSections = ObjPE.NTHeader.FileHeader.NumberOfSections
		}
	} else {
		fmt.Printf("Cannot read this type of file!")
		return
	}

	DisplayDosHeader(DosHeader)

	if hasRichHdr {
		DecryptAndDisplayRichHeader(RichHeader)
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
				DisplayRelocations(relocPages, relocArray)
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
		CheckError(err)

		var hasRichHdr bool = HasRichHeader(&data)
		var bitRange string

		if hasRichHdr {
			bitRange = DetermineBitRange([2]byte{data[272], data[273]})
			DecryptRichHeader(&data)
		} else {
			bitRange = DetermineBitRange([2]byte{data[152], data[153]})
		}

		DisplayPE(&data, bitRange, hasRichHdr)

	}
}
