package main

import (
	"fmt"
	"html/template"
	"monegoo/controllers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	hour := t.Hour()
	minutes := t.Minute()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d", year, month, day, hour, minutes)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles(
		"./testdata/raw.html",
		"./testdata/version.html")

	router.GET("/", func(c *gin.Context) {

		var currency = controllers.GetCurrencies()

		c.HTML(http.StatusOK, "raw.html", gin.H{
			"now": time.Now(),

			"aed": currency[0].AED_Currency,
			"afn": currency[0].AFN_Currency,
			"all": currency[0].ALL_Currency,
			"amd": currency[0].AMD_Currency,
			"ang": currency[0].ANG_Currency,
			"aoa": currency[0].AOA_Currency,
			"ars": currency[0].ARS_Currency,
			"aud": currency[0].AUD_Currency,
			"awg": currency[0].AWG_Currency,
			"azn": currency[0].AZN_Currency,
			"bam": currency[0].BAM_Currency,
			"bbd": currency[0].BBD_Currency,
			"bdt": currency[0].BDT_Currency,
			"bgn": currency[0].BGN_Currency,
			"bhd": currency[0].BHD_Currency,
			"bif": currency[0].BIF_Currency,
			"bmd": currency[0].BMD_Currency,
			"bnd": currency[0].BND_Currency,
			"bob": currency[0].BOB_Currency,
			"brl": currency[0].BRL_Currency,
			"bsd": currency[0].BSD_Currency,
			"btc": currency[0].BTC_Currency,
			"btn": currency[0].BTN_Currency,
			"bwp": currency[0].BWP_Currency,
			"byn": currency[0].BYN_Currency,
			"bzd": currency[0].BZD_Currency,
			"cad": currency[0].CAD_Currency,
			"cdf": currency[0].CDF_Currency,
			"chf": currency[0].CHF_Currency,
			"clf": currency[0].CLF_Currency,
			"clp": currency[0].CLP_Currency,
			"cnh": currency[0].CNH_Currency,
			"cny": currency[0].CNY_Currency,
			"cop": currency[0].COP_Currency,
			"crc": currency[0].CRC_Currency,
			"cuc": currency[0].CUC_Currency,
			"cup": currency[0].CUP_Currency,
			"cve": currency[0].CVE_Currency,
			"czk": currency[0].CZK_Currency,
			"djf": currency[0].DJF_Currency,
			"dkk": currency[0].DKK_Currency,
			"dop": currency[0].DOP_Currency,
			"dzd": currency[0].DZD_Currency,
			"egp": currency[0].EGP_Currency,
			"ern": currency[0].ERN_Currency,
			"etb": currency[0].ETB_Currency,
			"eur": currency[0].EUR_Currency,
			"fjd": currency[0].FJD_Currency,
			"fkp": currency[0].FKP_Currency,
			"gbp": currency[0].GBP_Currency,
			"gel": currency[0].GEL_Currency,
			"ggp": currency[0].GGP_Currency,
			"ghs": currency[0].GHS_Currency,
			"gip": currency[0].GIP_Currency,
			"gmd": currency[0].GMD_Currency,
			"gnf": currency[0].GNF_Currency,
			"gtq": currency[0].GTQ_Currency,
			"gyd": currency[0].GYD_Currency,
			"hkd": currency[0].HKD_Currency,
			"hnl": currency[0].HNL_Currency,
			"hrk": currency[0].HRK_Currency,
			"htg": currency[0].HTG_Currency,
			"huf": currency[0].HUF_Currency,
			"idr": currency[0].IDR_Currency,
			"ils": currency[0].ILS_Currency,
			"imp": currency[0].IMP_Currency,
			"inr": currency[0].INR_Currency,
			"iqd": currency[0].IQD_Currency,
			"irr": currency[0].IRR_Currency,
			"isk": currency[0].ISK_Currency,
			"jep": currency[0].JEP_Currency,
			"jmd": currency[0].JMD_Currency,
			"jod": currency[0].JOD_Currency,
			"jpy": currency[0].JPY_Currency,
			"kes": currency[0].KES_Currency,
			"kgs": currency[0].KGS_Currency,
			"khr": currency[0].KHR_Currency,
			"kmf": currency[0].KMF_Currency,
			"kpw": currency[0].KPW_Currency,
			"krw": currency[0].KRW_Currency,
			"kwd": currency[0].KWD_Currency,
			"kyd": currency[0].KYD_Currency,
			"kzt": currency[0].KZT_Currency,
			"lak": currency[0].LAK_Currency,
			"lbp": currency[0].LBP_Currency,
			"lkr": currency[0].LKR_Currency,
			"ldr": currency[0].LRD_Currency,
			"lsl": currency[0].LSL_Currency,
			"lyd": currency[0].LYD_Currency,
			"mad": currency[0].MAD_Currency,
			"mdl": currency[0].MDL_Currency,
			"mga": currency[0].MGA_Currency,
			"mkd": currency[0].MKD_Currency,
			"mmk": currency[0].MMK_Currency,
			"mnt": currency[0].MNT_Currency,
			"mop": currency[0].MOP_Currency,
			"mru": currency[0].MRU_Currency,
			"mur": currency[0].MUR_Currency,
			"mvr": currency[0].MVR_Currency,
			"mwk": currency[0].MWK_Currency,
			"mxn": currency[0].MXN_Currency,
			"myr": currency[0].MYR_Currency,
			"mzn": currency[0].MZN_Currency,
			"nad": currency[0].NAD_Currency,
			"ngn": currency[0].NGN_Currency,
			"nio": currency[0].NIO_Currency,
			"nok": currency[0].NOK_Currency,
			"npr": currency[0].NPR_Currency,
			"nzd": currency[0].NZD_Currency,
			"omr": currency[0].OMR_Currency,
			"pab": currency[0].PAB_Currency,
			"pen": currency[0].PEN_Currency,
			"pgk": currency[0].PGK_Currency,
			"php": currency[0].PHP_Currency,
			"pkr": currency[0].PKR_Currency,
			"pln": currency[0].PLN_Currency,
			"pyg": currency[0].PYG_Currency,
			"qar": currency[0].QAR_Currency,
			"ron": currency[0].RON_Currency,
			"rsd": currency[0].RSD_Currency,
			"rub": currency[0].RUB_Currency,
			"rwf": currency[0].RWF_Currency,
			"sar": currency[0].SAR_Currency,
			"sbd": currency[0].SBD_Currency,
			"scr": currency[0].SCR_Currency,
			"sdg": currency[0].SDG_Currency,
			"sek": currency[0].SEK_Currency,
			"sgd": currency[0].SGD_Currency,
			"shp": currency[0].SHP_Currency,
			"sll": currency[0].SLL_Currency,
			"sos": currency[0].SOS_Currency,
			"srd": currency[0].SRD_Currency,
			"ssp": currency[0].SSP_Currency,
			"std": currency[0].STD_Currency,
			"stn": currency[0].STN_Currency,
			"svc": currency[0].SVC_Currency,
			"syp": currency[0].SYP_Currency,
			"szl": currency[0].SZL_Currency,
			"thb": currency[0].THB_Currency,
			"tjs": currency[0].TJS_Currency,
			"tmt": currency[0].TMT_Currency,
			"tnd": currency[0].TND_Currency,
			"top": currency[0].TOP_Currency,
			"try": currency[0].TRY_Currency,
			"ttd": currency[0].TTD_Currency,
			"twd": currency[0].TWD_Currency,
			"tzs": currency[0].TZS_Currency,
			"uah": currency[0].UAH_Currency,
			"ugx": currency[0].UGX_Currency,
			"usd": currency[0].USD_Currency,
			"uyu": currency[0].UYU_Currency,
			"uzs": currency[0].UZS_Currency,
			"ves": currency[0].VES_Currency,
			"vnd": currency[0].VND_Currency,
			"vyv": currency[0].VUV_Currency,
			"wst": currency[0].WST_Currency,
			"xaf": currency[0].XAF_Currency,
			"xag": currency[0].XAG_Currency,
			"xau": currency[0].XAU_Currency,
			"xcd": currency[0].XCD_Currency,
			"xdr": currency[0].XDR_Currency,
			"xof": currency[0].XOF_Currency,
			"xpd": currency[0].XPD_Currency,
			"xpf": currency[0].XPF_Currency,
			"xpt": currency[0].XPT_Currency,
			"yer": currency[0].YER_Currency,
			"zar": currency[0].ZAR_Currency,
			"zmw": currency[0].ZMW_Currency,
			"zwl": currency[0].ZWL_Currency,
		})
	})

	router.GET("/version", func(c *gin.Context) {
		c.HTML(http.StatusOK, "version.html", gin.H{
			"version": "1.0.4",
		})
	})

	router.Run(":8080")
}
