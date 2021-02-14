package cvr

type Virksomhed struct {
	Aarsbeskaeftigelse   []VirksomhedAarsbeskaeftigelse `json:"aarsbeskaeftigelse"`
	Attributter          []VirksomhedAttribut           `json:"attributter"`
	Beliggenhedsadresser []VirksomhedAdresse            `json:"beliggenhedsadresse"`

	Bibranche1         []VirksomhedBranche `json:"bibranche1"`
	Bibranche2         []VirksomhedBranche `json:"bibranche2"`
	Bibranche3         []VirksomhedBranche `json:"bibranche3"`
	Binavne            []VirksomhedNavn    `json:"binavne"`
	BrancheAnsvarskode *int                `json:"brancheAnsvarskode"`

	CVRNummer int `json:"cvrNummer"`

	DataAdgang       int                          `json:"dataAdgang"`
	DeltagerRelation []VirksomhedDeltagerRelation `json:"deltagerRelation"`
	ElektroniskPost  []VirksomhedKontaktoplysning `json:"elektroniskPost"`
	Enhedsnummer     int                          `json:"enhedsNummer"`
	Enhedstype       string                       `json:"enhedstype"`

	FejlBeskrivelse         *string                            `json:"fejlBeskrivelse"`
	FejlRegistreret         bool                               `json:"fejlRegistreret"`
	FejlVedIndlaesning      bool                               `json:"fejlVedIndlaesning"`
	Fusioner                []VirksomhedFusion                 `json:"fusioner"`
	Hjemmeside              []VirksomhedKontaktoplysning       `json:"hjemmeside"`
	Hovedbranche            []VirksomhedBranche                `json:"hovedbranche"`
	Kvartalsbeskaeftigelse  []VirksomhedKvartalsbeskaeftigelse `json:"kvartalsbeskaeftigelse"`
	Livsforloeb             []VirksomhedLivsforloeb            `json:"livsforloeb"`
	Maanedsbeskaeftigelse   []VirksomhedMaanedsbeskaeftigelse  `json:"maanedsbeskaeftigelse"`
	NaermesteFremtidigeDato *string                            `json:"naermesteFremtidigeDato"`
	Navne                   []VirksomhedNavn                   `json:"navne"`
	ObligatoriskEmail       []VirksomhedObligatoriskEmail      `json:"obligatoriskEmail"`
	PEnheder                []interface{}                      `json:"penheder"`
	Postadresse             []VirksomhedAdresse                `json:"postadresse"`
	RegNummer               []VirksomhedRegNummer              `json:"regNummer"`
	Reklamebeskyttet        bool                               `json:"reklamebeskyttet"`
	SamtID                  int                                `json:"samtId"`

	SekundaertTelefaxNummer []VirksomhedKontaktoplysning `json:"sekundaertTelefaxNummer"`
	SekundaertTelefonNummer []VirksomhedKontaktoplysning `json:"sekundaertTelefonNummer"`
	SidstIndlaest
	SidstOpdateret
	Spaltninger        []VirksomhedSpaltning        `json:"spaltninger"`
	Status             []VirksomhedStatus           `json:"status"`
	TelefaxNummer      []VirksomhedKontaktoplysning `json:"telefaxNummer"`
	TelefonNummer      []VirksomhedKontaktoplysning `json:"telefonNummer"`
	VirkningsAktoer    string                       `json:"virkningsAktoer"`
	VirksomhedMetadata VirksomhedMetadata           `json:"virksomhedMetadata"`
	Virksomhedsstatus  []Virksomhedsstatus          `json:"virksomhedsstatus"`
}

type VirksomhedAttribut struct {
	Sekvensnummer int                `json:"sekvensnr"`
	Type          string             `json:"type"`
	Vaerdier      []VirksomhedVaerdi `json:"vaerdier"`
	Vaerditype    string             `json:"string"`
}

type VirksomhedVaerdi struct {
	Vaerdi  string `json:"vaerdi"`
	Periode `json:"periode"`
	SidstOpdateret
}

type VirksomhedAdresse struct {
	AdresseID      *string           `json:"adresseId"`
	BogstavFra     *string           `json:"bogstavFra"`
	BogstavTil     *string           `json:"bogstavTil"`
	Bynavn         *string           `json:"bynavn"`
	COnavn         *string           `json:"conavn"`
	Etage          *string           `json:"etage"`
	Fritekst       *string           `json:"fritekst"`
	HusnummerFra   *int              `json:"husnummerFra"`
	HusnummerTil   *int              `json:"husnummerTil"`
	Kommune        VirksomhedKommune `json:"kommune"`
	Landekode      *string           `json:"landekode"`
	Periode        `json:"periode"`
	Postboks       *string `json:"postboks"`
	Postdistrikt   *string `json:"postdistrikt"`
	Postnummer     int     `json:"postnummer"`
	Sidedoer       *string `json:"sidedoer"`
	SidstOpdateret string  `json:"sidstOpdateret"`
	SidstValideret *string `json:"sidstValideret"`
	Vejkode        *int    `json:"vejkode"`
	Vejnavn        *string `json:"vejnavn"`
}

type VirksomhedKommune struct {
	Kommunekode int    `json:"kommuneKode"`
	KommuneNavn string `json:"kommuneNavn"`
	Periode     `json:"periode"`
	SidstOpdateret
}

type VirksomhedBranche struct {
	Branchekode  string `json:"branchekode"`
	Branchetekst string `json:"branchetekst"`
	Periode      `json:"periode"`
	SidstOpdateret
}

type VirksomhedPenhed struct {
	PNummer int `json:"pNummer"`
	Periode `json:"periode"`
	SidstOpdateret
}

type VirksomhedKontaktoplysning struct {
	Hemmelig         bool   `json:"hemmelig"`
	Kontaktoplysning string `json:"kontaktoplysning"`
	Periode          `json:"periode"`
	SidstOpdateret
}

type VirksomhedAarsbeskaeftigelse struct {
	Aar                             int      `json:"aar"`
	AntalAarsvaerk                  *float32 `json:"antalAarsvaerk"`
	AntalAnsatte                    *int     `json:"antalAnsatte"`
	AntalInklusivEjere              *int     `json:"antalInklusivEjere"`
	IntervalKodeAntalAarsvaerk      string   `json:"intervalKodeAntalAarsvaerk"`
	IntervalKodeAntalAnsatte        string   `json:"intervalKodeAntalAnsatte"`
	IntervalKodeAntalInklusiveEjere string   `json:"intervalKodeAntalInklusivEjere"`
	SidstOpdateret
}

type VirksomhedKvartalsbeskaeftigelse struct {
	Aar                        int     `json:"aar"`
	AntalAarsvaerk             *int    `json:"antalAarsvaerk"`
	AntalAnsatte               *int    `json:"antalAnsatte"`
	IntervalKodeAntalAarsvaerk *string `json:"intervalKodeAntalAarsvaerk"`
	IntervalKodeAntalAnsatte   *string `json:"intervalKodeAntalAnsatte"`
	Kvartal                    int     `json:"kvartal"`
	SidstOpdateret             string  `json:"sidstOpdateret"`
}

type VirksomhedLivsforloeb struct {
	Periode `json:"periode"`
	SidstOpdateret
}

type VirksomhedMaanedsbeskaeftigelse struct {
	Aar                        int     `json:"aar"`
	AntalAarsvaerk             *int    `json:"antalAarsvaerk"`
	AntalAnsatte               *int    `json:"antalAnsatte"`
	IntervalKodeAntalAarsvaerk *string `json:"intervalKodeAntalAarsvaerk"`
	IntervalKodeAntalAnsatte   *string `json:"intervalKodeAntalAnsatte"`
	Maaned                     int     `json:"maaned"`
	SidstOpdateret             string  `json:"sidstOpdateret"`
}

type VirksomhedNavn struct {
	Navn    string `json:"navn"`
	Periode `json:"periode"`
	SidstOpdateret
}

func (vn VirksomhedNavn) GetPeriode() Periode {
	return vn.Periode
}

type VirksomhedObligatoriskEmail struct {
	Hemmelig         bool   `json:"hemmelig"`
	Kontaktoplysning string `json:"kontaktoplysning"`
	Periode          `json:"periode"`
	SidstOpdateret
}

type VirksomhedRegNummer struct {
	RegNummer string `json:"regNummer"`
	Periode   `json:"periode"`
	SidstOpdateret
}

type VirksomhedStatus struct {
	Kreditoplysningskode  int     `json:"kreditoplysningskode"`
	Kreditoplysningstekst *string `json:"kreditoplysningstekst"`
	Statuskode            int     `json:"statuskode"`
	Statustekst           *string `json:"statustekst"`
	Periode               `json:"periode"`
	SidstOpdateret
}

type Virksomhedsstatus struct {
	Status  string `json:"status"`
	Periode `json:"periode"`
	SidstOpdateret
}

type VirksomhedFusion struct {
	EnhedsNummerOrganisation int                  `json:"enhedsNummerOrganisation"`
	OrganisationsNavn        []VirksomhedNavn     `json:"organisationsNavn"`
	Indgaaende               []VirksomhedAttribut `json:"indgaaende"`
	Udgaaende                []VirksomhedAttribut `json:"udgaaende"`
}

type VirksomhedSpaltning struct {
	EnhedsNummerOrganisation int                  `json:"enhedsNummerOrganisation"`
	OrganisationsNavn        []VirksomhedNavn     `json:"organisationsNavn"`
	Indgaaende               []VirksomhedAttribut `json:"indgaaende"`
	Udgaaende                []VirksomhedAttribut `json:"udgaaende"`
}

type VirksomhedMetadata struct {
	NyesteNavn                   VirksomhedNavn                    `json:"nyesteNavn"`
	NyesteBinavne                []string                          `json:"nyesteBinavne"`
	NyesteVirksomhedsform        VirksomhedVirksomhedsform         `json:"nyesteVirksomhedsform"`
	NyesteBeliggenhedsadresse    VirksomhedAdresse                 `json:"nyesteBeliggenhedsadresse"`
	NyesteHovedbranche           VirksomhedBranche                 `json:"nyesteHovedbranche"`
	NyesteBibranche1             VirksomhedBranche                 `json:"nyesteBibranche1"`
	NyesteBibranche2             VirksomhedBranche                 `json:"nyesteBibranche2"`
	NyesteBibranche3             VirksomhedBranche                 `json:"nyesteBibranche3"`
	NyesteStatus                 VirksomhedStatus                  `json:"nyesteStatus"`
	NyesteKontaktoplysninger     []string                          `json:"nyesteKontaktoplysninger"`
	AntalPenheder                int                               `json:"antalPenheder"`
	NyesteAarsbeskaeftigelse     *VirksomhedAarsbeskaeftigelse     `json:"nyesteAarsbeskaeftigelse"`
	NyesteKvartalsbeskaeftigelse *VirksomhedKvartalsbeskaeftigelse `json:"nyesteKvartalsbeskaeftigelse"`
	NyesteMaanedsbeskaeftigelse  *VirksomhedMaanedsbeskaeftigelse  `json:"nyesteMaanedsbeskaeftigelse"`
	SammensatStatus              string                            `json:"sammensatStatus"`
	StiftelsesDato               *GyldigDato                       `json:"stiftelsesDato"`
	VirkningsDato                *GyldigDato                       `json:"virkningsDato"`
}

type VirksomhedVirksomhedsform struct {
	AnsvarligDataleverandoer string `json:"ansvarligDataleverandoer"`
	KortBeskrivelse          string `json:"kortBeskrivelse"`
	LangBeskrivelse          string `json:"langBeskrivelse"`
	Virksomhedsformkode      int    `json:"virksomhedsformkode"`
	Periode                  `json:"periode"`
	SidstOpdateret
}

type VirksomhedDeltagerRelation struct {
	Deltager       `json:"deltager"`
	Organisationer []VirksomhedOrganisation `json:"organisationer"`
	Kontorsteder   []VirksomhedKontorsted   `json:"kontorsteder"`
}

type Deltager struct {
	Navne        []VirksomhedNavn `json:"navne"`
	Enhedsnummer int              `json:"enhedsNummer"`
	Enhedstype   string           `json:"enhedstype"`
	SidstIndlaest
	SidstOpdateret
	Organisationstype *string `json:"organisationstype"`
	AdresseHemmelig   bool    `json:"adresseHemmelig"`
}

type VirksomhedKontorsted struct{}

type VirksomhedOrganisation struct {
	Enhedsnummer int              `json:"enhedsNummerOrganisation"`
	Hovedtype    string           `json:"hovedtype"`
	Navne        []VirksomhedNavn `json:"organisationsNavn"`
}

type VirksomhedMinimal struct {
	CVRNummer        int    `json:"cvrNummer"`
	Reklamebeskyttet bool   `json:"reklamebeskyttet"`
	DataAdgang       int    `json:"dataAdgang"`
	Enhedsnummer     int    `json:"enhedsNummer"`
	Enhedstype       string `json:"enhedstype"`
	SidstIndlaest
	SidstOpdateret
	Navne       []VirksomhedNavn        `json:"navne"`
	Livsforloeb []VirksomhedLivsforloeb `json:"livsforloeb"`
}
