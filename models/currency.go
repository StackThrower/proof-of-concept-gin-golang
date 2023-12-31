package models

type Currency struct {
	Id           int64   `db:"id"`
	Timestamp    int64   `db:"timestamp"`
	AED_Currency float64 `db: "AED_CURRENCY"`
	AFN_Currency float64 `db: "AFN_CURRENCY"`
	ALL_Currency float64 `db: "ALL_CURRENCY"`
	AMD_Currency float64 `db: "AMD_CURRENCY"`
	ANG_Currency float64 `db: "ANG_CURRENCY"`
	AOA_Currency float64 `db: "AOA_CURRENCY"`
	ARS_Currency float64 `db: "ARS_CURRENCY"`
	AUD_Currency float64 `db: "AUD_CURRENCY"`
	AWG_Currency float64 `db: "AWG_CURRENCY"`
	AZN_Currency float64 `db: "AZN_CURRENCY"`
	BAM_Currency float64 `db: "BAM_CURRENCY"`
	BBD_Currency float64 `db: "BBD_CURRENCY"`
	BDT_Currency float64 `db: "BDT_CURRENCY"`
	BGN_Currency float64 `db: "BGN_CURRENCY"`
	BHD_Currency float64 `db: "BHD_CURRENCY"`
	BIF_Currency float64 `db: "BIF_CURRENCY"`
	BMD_Currency float64 `db: "BMD_CURRENCY"`
	BND_Currency float64 `db: "BND_CURRENCY"`
	BOB_Currency float64 `db: "BOB_CURRENCY"`
	BRL_Currency float64 `db: "BRL_CURRENCY"`
	BSD_Currency float64 `db: "BSD_CURRENCY"`
	BTC_Currency float64 `db: "BTC_CURRENCY"`
	BTN_Currency float64 `db: "BTN_CURRENCY"`
	BWP_Currency float64 `db: "BWP_CURRENCY"`
	BYN_Currency float64 `db: "BYN_CURRENCY"`
	BZD_Currency float64 `db: "BZD_CURRENCY"`
	CAD_Currency float64 `db: "CAD_CURRENCY"`
	CDF_Currency float64 `db: "CDF_CURRENCY"`
	CHF_Currency float64 `db: "CHF_CURRENCY"`
	CLF_Currency float64 `db: "CLF_CURRENCY"`
	CLP_Currency float64 `db: "CLP_CURRENCY"`
	CNH_Currency float64 `db: "CNH_CURRENCY"`
	CNY_Currency float64 `db: "CNY_CURRENCY"`
	COP_Currency float64 `db: "COP_CURRENCY"`
	CRC_Currency float64 `db: "CRC_CURRENCY"`
	CUC_Currency float64 `db: "CUC_CURRENCY"`
	CUP_Currency float64 `db: "CUP_CURRENCY"`
	CVE_Currency float64 `db: "CVE_CURRENCY"`
	CZK_Currency float64 `db: "CZK_CURRENCY"`
	DJF_Currency float64 `db: "DJF_CURRENCY"`
	DKK_Currency float64 `db: "DKK_CURRENCY"`
	DOP_Currency float64 `db: "DOP_CURRENCY"`
	DZD_Currency float64 `db: "DZD_CURRENCY"`
	EGP_Currency float64 `db: "EGP_CURRENCY"`
	ERN_Currency float64 `db: "ERN_CURRENCY"`
	ETB_Currency float64 `db: "ETB_CURRENCY"`
	EUR_Currency float64 `db: "EUR_CURRENCY"`
	FJD_Currency float64 `db: "FJD_CURRENCY"`
	FKP_Currency float64 `db: "FKP_CURRENCY"`
	GBP_Currency float64 `db: "GBP_CURRENCY"`
	GEL_Currency float64 `db: "GEL_CURRENCY"`
	GGP_Currency float64 `db: "GGP_CURRENCY"`
	GHS_Currency float64 `db: "GHS_CURRENCY"`
	GIP_Currency float64 `db: "GIP_CURRENCY"`
	GMD_Currency float64 `db: "GMD_CURRENCY"`
	GNF_Currency float64 `db: "GNF_CURRENCY"`
	GTQ_Currency float64 `db: "GTQ_CURRENCY"`
	GYD_Currency float64 `db: "GYD_CURRENCY"`
	HKD_Currency float64 `db: "HKD_CURRENCY"`
	HNL_Currency float64 `db: "HNL_CURRENCY"`
	HRK_Currency float64 `db: "HRK_CURRENCY"`
	HTG_Currency float64 `db: "HTG_CURRENCY"`
	HUF_Currency float64 `db: "HUF_CURRENCY"`
	IDR_Currency float64 `db: "IDR_CURRENCY"`
	ILS_Currency float64 `db: "ILS_CURRENCY"`
	IMP_Currency float64 `db: "IMP_CURRENCY"`
	INR_Currency float64 `db: "INR_CURRENCY"`
	IQD_Currency float64 `db: "IQD_CURRENCY"`
	IRR_Currency float64 `db: "IRR_CURRENCY"`
	ISK_Currency float64 `db: "ISK_CURRENCY"`
	JEP_Currency float64 `db: "JEP_CURRENCY"`
	JMD_Currency float64 `db: "JMD_CURRENCY"`
	JOD_Currency float64 `db: "JOD_CURRENCY"`
	JPY_Currency float64 `db: "JPY_CURRENCY"`
	KES_Currency float64 `db: "KES_CURRENCY"`
	KGS_Currency float64 `db: "KGS_CURRENCY"`
	KHR_Currency float64 `db: "KHR_CURRENCY"`
	KMF_Currency float64 `db: "KMF_CURRENCY"`
	KPW_Currency float64 `db: "KPW_CURRENCY"`
	KRW_Currency float64 `db: "KRW_CURRENCY"`
	KWD_Currency float64 `db: "KWD_CURRENCY"`
	KYD_Currency float64 `db: "KYD_CURRENCY"`
	KZT_Currency float64 `db: "KZT_CURRENCY"`
	LAK_Currency float64 `db: "LAK_CURRENCY"`
	LBP_Currency float64 `db: "LBP_CURRENCY"`
	LKR_Currency float64 `db: "LKR_CURRENCY"`
	LRD_Currency float64 `db: "LRD_CURRENCY"`
	LSL_Currency float64 `db: "LSL_CURRENCY"`
	LYD_Currency float64 `db: "LYD_CURRENCY"`
	MAD_Currency float64 `db: "MAD_CURRENCY"`
	MDL_Currency float64 `db: "MDL_CURRENCY"`
	MGA_Currency float64 `db: "MGA_CURRENCY"`
	MKD_Currency float64 `db: "MKD_CURRENCY"`
	MMK_Currency float64 `db: "MMK_CURRENCY"`
	MNT_Currency float64 `db: "MNT_CURRENCY"`
	MOP_Currency float64 `db: "MOP_CURRENCY"`
	MRU_Currency float64 `db: "MRU_CURRENCY"`
	MUR_Currency float64 `db: "MUR_CURRENCY"`
	MVR_Currency float64 `db: "MVR_CURRENCY"`
	MWK_Currency float64 `db: "MWK_CURRENCY"`
	MXN_Currency float64 `db: "MXN_CURRENCY"`
	MYR_Currency float64 `db: "MYR_CURRENCY"`
	MZN_Currency float64 `db: "MZN_CURRENCY"`
	NAD_Currency float64 `db: "NAD_CURRENCY"`
	NGN_Currency float64 `db: "NGN_CURRENCY"`
	NIO_Currency float64 `db: "NIO_CURRENCY"`
	NOK_Currency float64 `db: "NOK_CURRENCY"`
	NPR_Currency float64 `db: "NPR_CURRENCY"`
	NZD_Currency float64 `db: "NZD_CURRENCY"`
	OMR_Currency float64 `db: "OMR_CURRENCY"`
	PAB_Currency float64 `db: "PAB_CURRENCY"`
	PEN_Currency float64 `db: "PEN_CURRENCY"`
	PGK_Currency float64 `db: "PGK_CURRENCY"`
	PHP_Currency float64 `db: "PHP_CURRENCY"`
	PKR_Currency float64 `db: "PKR_CURRENCY"`
	PLN_Currency float64 `db: "PLN_CURRENCY"`
	PYG_Currency float64 `db: "PYG_CURRENCY"`
	QAR_Currency float64 `db: "QAR_CURRENCY"`
	RON_Currency float64 `db: "RON_CURRENCY"`
	RSD_Currency float64 `db: "RSD_CURRENCY"`
	RUB_Currency float64 `db: "RUB_CURRENCY"`
	RWF_Currency float64 `db: "RWF_CURRENCY"`
	SAR_Currency float64 `db: "SAR_CURRENCY"`
	SBD_Currency float64 `db: "SBD_CURRENCY"`
	SCR_Currency float64 `db: "SCR_CURRENCY"`
	SDG_Currency float64 `db: "SDG_CURRENCY"`
	SEK_Currency float64 `db: "SEK_CURRENCY"`
	SGD_Currency float64 `db: "SGD_CURRENCY"`
	SHP_Currency float64 `db: "SHP_CURRENCY"`
	SLL_Currency float64 `db: "SLL_CURRENCY"`
	SOS_Currency float64 `db: "SOS_CURRENCY"`
	SRD_Currency float64 `db: "SRD_CURRENCY"`
	SSP_Currency float64 `db: "SSP_CURRENCY"`
	STD_Currency float64 `db: "STD_CURRENCY"`
	STN_Currency float64 `db: "STN_CURRENCY"`
	SVC_Currency float64 `db: "SVC_CURRENCY"`
	SYP_Currency float64 `db: "SYP_CURRENCY"`
	SZL_Currency float64 `db: "SZL_CURRENCY"`
	THB_Currency float64 `db: "THB_CURRENCY"`
	TJS_Currency float64 `db: "TJS_CURRENCY"`
	TMT_Currency float64 `db: "TMT_CURRENCY"`
	TND_Currency float64 `db: "TND_CURRENCY"`
	TOP_Currency float64 `db: "TOP_CURRENCY"`
	TRY_Currency float64 `db: "TRY_CURRENCY"`
	TTD_Currency float64 `db: "TTD_CURRENCY"`
	TWD_Currency float64 `db: "TWD_CURRENCY"`
	TZS_Currency float64 `db: "TZS_CURRENCY"`
	UAH_Currency float64 `db: "UAH_CURRENCY"`
	UGX_Currency float64 `db: "UGX_CURRENCY"`
	USD_Currency float64 `db: "USD_CURRENCY"`
	UYU_Currency float64 `db: "UYU_CURRENCY"`
	UZS_Currency float64 `db: "UZS_CURRENCY"`
	VES_Currency float64 `db: "VES_CURRENCY"`
	VND_Currency float64 `db: "VND_CURRENCY"`
	VUV_Currency float64 `db: "VUV_CURRENCY"`
	WST_Currency float64 `db: "WST_CURRENCY"`
	XAF_Currency float64 `db: "XAF_CURRENCY"`
	XAG_Currency float64 `db: "XAG_CURRENCY"`
	XAU_Currency float64 `db: "XAU_CURRENCY"`
	XCD_Currency float64 `db: "XCD_CURRENCY"`
	XDR_Currency float64 `db: "XDR_CURRENCY"`
	XOF_Currency float64 `db: "XOF_CURRENCY"`
	XPD_Currency float64 `db: "XPD_CURRENCY"`
	XPF_Currency float64 `db: "XPF_CURRENCY"`
	XPT_Currency float64 `db: "XPT_CURRENCY"`
	YER_Currency float64 `db: "YER_CURRENCY"`
	ZAR_Currency float64 `db: "ZAR_CURRENCY"`
	ZMW_Currency float64 `db: "ZMW_CURRENCY"`
	ZWL_Currency float64 `db: "ZWL_CURRENCY"`
}
