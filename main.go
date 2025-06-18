package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	zip_streamer "github.com/scosman/zipstreamer/zip_streamer"
)

func main() {
	fmt.Println(zip_streamer.EncryptIt([]byte(`{
  "suggestedFilename": "W409101448A_2025_06_18_09_04_27.zip",
  "files": [
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EED8FE7EA7AAFEDD16F/TI01392TRU_0522-00.pdf",
      "zipPath": "Technical_information/RU/TI01392TRU_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9B9B5F06DBE88985/BA01915TEL_0118.pdf",
      "zipPath": "Operating_instruction/EL/BA01915TEL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEDBBB0339BB08E9292/BA01915TZH_0118-00.pdf",
      "zipPath": "Operating_instruction/ZH/BA01915TZH_0118-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED997E92C8D8CDCD5BE/TI01442TDE_0118.pdf",
      "zipPath": "Technical_information/DE/TI01442TDE_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B88723E5485F2BAC/KA01414TJA_0219.pdf",
      "zipPath": "Short_operating_instruction/JA/KA01414TJA_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA4EA81BC1A6F0ED9/TI01442TPL_0118.pdf",
      "zipPath": "Technical_information/PL/TI01442TPL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA99EEB9FEE1C18EC1/BA01915TCS_0118.pdf",
      "zipPath": "Operating_instruction/CS/BA01915TCS_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DRAWXL/005056A500261EDF82F2DB6A1BF7BB49/1016280294_W407CF144D9_PZ31.pdf",
      "zipPath": "Inspection_certificate_acc._EN_10204-3.1/EN/1016280294_W407CF144D9_PZ31.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEAA18EA76922C9FE79/KA01414TSK_0219.pdf",
      "zipPath": "Short_operating_instruction/SK/KA01414TSK_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9ACE08F6FAF521F5/BA01915TSL_0118.pdf",
      "zipPath": "Operating_instruction/SL/BA01915TSL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EED8FE7E8732AEED114/TI01392TPT_0522-00.pdf",
      "zipPath": "Technical_information/PT/TI01392TPT_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDD8ECACB5F39065419/TI01392TES_0522-00.pdf",
      "zipPath": "Technical_information/ES/TI01392TES_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEA9CFE33E109675B28/BA01915TTR_0118.pdf",
      "zipPath": "Operating_instruction/TR/BA01915TTR_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE99198B0E4DD7/BA01915TFR_0118.pdf",
      "zipPath": "Operating_instruction/FR/BA01915TFR_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE99B763762DD9/BA01915TJA_0118.pdf",
      "zipPath": "Operating_instruction/JA/BA01915TJA_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA97F131854E40FEF2/BA01915TPL_0118.pdf",
      "zipPath": "Operating_instruction/PL/BA01915TPL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EED8FE7EB72BCB0F16F/TI01392TJA_0522-00.pdf",
      "zipPath": "Technical_information/JA/TI01392TJA_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B9A46F11B6402C97/KA01414TPT_0219.pdf",
      "zipPath": "Short_operating_instruction/PT/KA01414TPT_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA29208DB2F7712D3/KA01414TDA_0219.pdf",
      "zipPath": "Short_operating_instruction/DA/KA01414TDA_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA292B1D26A04D66D/KA01414TZH_0219.pdf",
      "zipPath": "Short_operating_instruction/ZH/KA01414TZH_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEA9A8A0B153973C00A/BA01915TBG_0118.pdf",
      "zipPath": "Operating_instruction/BG/BA01915TBG_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE939CB4E3A711F/BA01915TSV_0118.pdf",
      "zipPath": "Operating_instruction/SV/BA01915TSV_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDC86F699ABF0F9DAC8/BA01915TRU_0118.pdf",
      "zipPath": "Operating_instruction/RU/BA01915TRU_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEAA18EA67719471E77/KA01414TEL_0219.pdf",
      "zipPath": "Short_operating_instruction/EL/KA01414TEL_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9ABBD3FD6E9117DB4/BA01915TIT_0118.pdf",
      "zipPath": "Operating_instruction/IT/BA01915TIT_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEBA8EA3B34EC275930/BA01854TRU_0420.pdf",
      "zipPath": "Operating_instruction/RU/BA01854TRU_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE84A4DFAEECCFB/BA01915TNO_0118.pdf",
      "zipPath": "Operating_instruction/NO/BA01915TNO_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE4899E53AAC0C/TI01442TPT_0118.pdf",
      "zipPath": "Technical_information/PT/TI01442TPT_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE9209E51EB30BB/BA01915THU_0118.pdf",
      "zipPath": "Operating_instruction/HU/BA01915THU_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA2929EB9CB2F1600/KA01414TSV_0219.pdf",
      "zipPath": "Short_operating_instruction/SV/KA01414TSV_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA2924BBD8B075472/KA01414TNO_0219.pdf",
      "zipPath": "Short_operating_instruction/NO/KA01414TNO_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE49FA1B53EC0E/TI01442TES_0118.pdf",
      "zipPath": "Technical_information/ES/TI01442TES_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A1BFA6744BBA5CFB/KA01414TDE_0219.pdf",
      "zipPath": "Short_operating_instruction/DE/KA01414TDE_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B2AFD71708DE363F/TI01442TRU_0118.pdf",
      "zipPath": "Technical_information/RU/TI01442TRU_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDD8BE78DFCBEC74E43/TI01392TEN_0522-00.pdf",
      "zipPath": "Technical_information/EN/TI01392TEN_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA2927D756AE9753C/KA01414THU_0219.pdf",
      "zipPath": "Short_operating_instruction/HU/KA01414THU_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA99EEBB38AAD00EC3/BA01915THR_0118.pdf",
      "zipPath": "Operating_instruction/HR/BA01915THR_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE985653974DD7/BA01915TPT_0118.pdf",
      "zipPath": "Operating_instruction/PT/BA01915TPT_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EED8FE7E971AD55316E/TI01392TFR_0522-00.pdf",
      "zipPath": "Technical_information/FR/TI01392TFR_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDD8ECACA703B089414/TI01392TIT_0522-00.pdf",
      "zipPath": "Technical_information/IT/TI01392TIT_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEAB5AF948EBAB5BB64/KA01414TTR_0219.pdf",
      "zipPath": "Short_operating_instruction/TR/KA01414TTR_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA292B0934B47B66C/KA01414THR_0219.pdf",
      "zipPath": "Short_operating_instruction/HR/KA01414THR_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE9383CBDC1111E/BA01915TTH_0118.pdf",
      "zipPath": "Operating_instruction/TH/BA01915TTH_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE4964BABF2C0E/TI01442TFR_0118.pdf",
      "zipPath": "Technical_information/FR/TI01442TFR_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EE9B0CC7C64BCC5D201/KA01414TFR_0219.pdf",
      "zipPath": "Short_operating_instruction/FR/KA01414TFR_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA4A4590484F1CAB9/KA01414TID_0219.pdf",
      "zipPath": "Short_operating_instruction/ID/KA01414TID_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A8AE4AB73422CC11/TI01442TJA_0118.pdf",
      "zipPath": "Technical_information/JA/TI01442TJA_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA29222647BF6134A/KA01414TET_0219.pdf",
      "zipPath": "Short_operating_instruction/ET/KA01414TET_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B20FAE3936D6E9/BA01854TDE_0420.pdf",
      "zipPath": "Operating_instruction/DE/BA01854TDE_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEAA18E92F960889E0B/KA01414TNL_0219.pdf",
      "zipPath": "Short_operating_instruction/NL/KA01414TNL_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED996C0FAE3E556D899/BA01915TDE_0118.pdf",
      "zipPath": "Operating_instruction/DE/BA01915TDE_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B20CB718C6B6D5/BA01854TEN_0420.pdf",
      "zipPath": "Operating_instruction/EN/BA01854TEN_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE98BD0C0755300/BA01915TDA_0118.pdf",
      "zipPath": "Operating_instruction/DA/BA01915TDA_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED996C0FA5BCD7BD898/BA01915TEN_0118.pdf",
      "zipPath": "Operating_instruction/EN/BA01915TEN_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDC86F69A9B2D207ACE/KA01414TRU_0219.pdf",
      "zipPath": "Short_operating_instruction/RU/KA01414TRU_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE907123EA81057/BA01915TID_0118.pdf",
      "zipPath": "Operating_instruction/ID/BA01915TID_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEAA18EC40DAC49BEE0/KA01414TBG_0219.pdf",
      "zipPath": "Short_operating_instruction/BG/KA01414TBG_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9A1BFA5CF966DBCF8/KA01414TEN_0219.pdf",
      "zipPath": "Short_operating_instruction/EN/KA01414TEN_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE8B28B4B6BCEBA/BA01915TLV_0118.pdf",
      "zipPath": "Operating_instruction/LV/BA01915TLV_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE91E9EAE0C50B6/BA01915TKO_0118.pdf",
      "zipPath": "Operating_instruction/KO/BA01915TKO_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA4A457B1C619EAB7/KA01414TRO_0219.pdf",
      "zipPath": "Short_operating_instruction/RO/KA01414TRO_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA292239532F7934E/KA01414TLT_0219.pdf",
      "zipPath": "Short_operating_instruction/LT/KA01414TLT_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE94C7ED21651D0/BA01915TNL_0118.pdf",
      "zipPath": "Operating_instruction/NL/BA01915TNL_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B2AE8746A7CA0EEF/BA01915T23ES_0118.pdf",
      "zipPath": "Operating_instruction/ES/BA01915T23ES_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EE9B0CC78438BE191DF/KA01414TIT_0219.pdf",
      "zipPath": "Short_operating_instruction/IT/KA01414TIT_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEA9FC91994B6CF581B/KA01414TTH_0219.pdf",
      "zipPath": "Short_operating_instruction/TH/KA01414TTH_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA2927C3B18F3B53B/KA01414TSL_0219.pdf",
      "zipPath": "Short_operating_instruction/SL/KA01414TSL_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEA9CFE354101697B2A/BA01915TFI_0118.pdf",
      "zipPath": "Operating_instruction/FI/BA01915TFI_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B2078F065796B1/BA01854TIT_0420.pdf",
      "zipPath": "Operating_instruction/IT/BA01854TIT_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEA9A9BADBF0A1D3E88/BA01915TSK_0118.pdf",
      "zipPath": "Operating_instruction/SK/BA01915TSK_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B20615F6919666/BA01854TJA_0420.pdf",
      "zipPath": "Operating_instruction/JA/BA01854TJA_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE97C3C4B9FB2A2/BA01915TET_0118.pdf",
      "zipPath": "Operating_instruction/ET/BA01915TET_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DRAWXL/005056A500261EDF82F2DB9B338B7B49/1016280294_W407CF144D9_PMZ.pdf",
      "zipPath": "Material_certificate/EN/1016280294_W407CF144D9_PMZ.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA4A4686C2B84CB1F/KA01414TPL_0219.pdf",
      "zipPath": "Short_operating_instruction/PL/KA01414TPL_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDD8BE79009FADB4E46/TI01392TDE_0522-00.pdf",
      "zipPath": "Technical_information/DE/TI01392TDE_0522-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B8872327586F8BA9/KA01414TES_0219.pdf",
      "zipPath": "Short_operating_instruction/ES/KA01414TES_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA1C3806A0C2BD7E6/KA01414TKO_0219.pdf",
      "zipPath": "Short_operating_instruction/KO/KA01414TKO_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B20B7D53C656CC/BA01854TES_0420.pdf",
      "zipPath": "Operating_instruction/ES/BA01854TES_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED997E92DB88877F5C1/TI01442TEN_0118.pdf",
      "zipPath": "Technical_information/EN/TI01442TEN_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B208D582CA96BC/BA01854TPT_0420.pdf",
      "zipPath": "Operating_instruction/PT/BA01854TPT_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE97DACABB7B2A4/BA01915TLT_0118.pdf",
      "zipPath": "Operating_instruction/LT/BA01915TLT_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDBA3B20A27874356C5/BA01854TFR_0420.pdf",
      "zipPath": "Operating_instruction/FR/BA01854TFR_0420.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDD91983581DA2503D8/BA01854TPL_0420-00.pdf",
      "zipPath": "Operating_instruction/PL/BA01854TPL_0420-00.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA29207A3A5A2F2D1/KA01414TCS_0219.pdf",
      "zipPath": "Short_operating_instruction/CS/KA01414TCS_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDA9CE94DBE6FC5D1D3/BA01915TRO_0118.pdf",
      "zipPath": "Operating_instruction/RO/BA01915TRO_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261ED9B2AFD7E6E4125642/TI01442TIT_0118.pdf",
      "zipPath": "Technical_information/IT/TI01442TIT_0118.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDAA2929FEC3CB0F600/KA01414TLV_0219.pdf",
      "zipPath": "Short_operating_instruction/LV/KA01414TLV_0219.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEF8DF1F6B6BD834650/1942%20TM1x1.pdf",
      "zipPath": "Certificate/RU/1942_TM1x1.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEF8DF2165318154650/1943%20TM1x1.pdf",
      "zipPath": "Certificate/RU/1943_TM1x1.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEEA4E81BA6B0789D50/02-2.0171%20%20iTHERM%20ModuLine%20%20%2012.07.2027.pdf",
      "zipPath": "Certificate/RU/02-2.0171__iTHERM_ModuLine___12.07.2027.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEEA4D0992DD9D29D1B/02-2.0170%2012.07.2027.pdf",
      "zipPath": "Certificate/RU/02-2.0170_12.07.2027.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDEBCE7DF2B184D3555/EC_00136_04.24.pdf",
      "zipPath": "Manufact._declaration/EN/EC_00136_04.24.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDEBCE7DF2B184D3555/EC_00136_04.24.pdf",
      "zipPath": "Manufact._declaration/DE/EC_00136_04.24.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEBB2A33EB9FCF2D8AB/EAEC%20N%20RU%20D-DE.BE02.B.12645_20.pdf",
      "zipPath": "Certificate/RU/EAEC_N_RU_D-DE.BE02.B.12645_20.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EEE8BA4C0B74A7B62FB/HE_00886_07_17.pdf",
      "zipPath": "Manufact._declaration/EN/HE_00886_07_17.pdf"
    },
    {
      "url": "https://bdih-download.endress.com/files/DLA/005056A500261EDEBE9BF032A81A56B8/TM1x1%20Suppliers%20Dec%20of%20Conformity_240412.pdf",
      "zipPath": "Manufact._declaration/EN/TM1x1_Suppliers_Dec_of_Conformity_240412.pdf"
    }
  ]
}`), os.Getenv("ZIP_STREAMER_KEY_PHRASE")))
	zipServer := zip_streamer.NewServer()
	zipServer.Compression = (os.Getenv("ZS_COMPRESSION") == "DEFLATE")
	zipServer.ListfileUrlPrefix = os.Getenv("ZS_LISTFILE_URL_PREFIX")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4008"
	}

	httpServer := &http.Server{
		Addr:        ":" + port,
		Handler:     zipServer,
		ReadTimeout: 10 * time.Second,
	}

	shutdownChannel := make(chan os.Signal, 10)
	go func() {
		log.Printf("Server starting on port %s", port)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Server Error: %s", err)
		}
		shutdownChannel <- syscall.SIGUSR1
	}()

	// Listen for os signal for graceful shutdown
	signal.Notify(shutdownChannel, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Wait for shutdown signal, then shut down
	shutdownSignal := <-shutdownChannel
	log.Printf("Received signal (%s), shutting down...", shutdownSignal.String())
	httpServer.Shutdown(context.Background())

	// Exit was not expected, return non 0 exit code
	if shutdownSignal == syscall.SIGUSR1 {
		os.Exit(1)
	}
}
