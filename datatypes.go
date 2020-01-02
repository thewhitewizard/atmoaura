package atmoaura


type CurrentIndex struct {
	City string `json:"commune"`
	Insee string `json:"code_insee"`
	Index Index `json:"indices"`
}

type Index struct {
	Color string `json:"couleur_html"`
	Value string `json:"valeur"`
	Date string `json:"date"`
	State string `json:"type_valeur"`	
	Qualifier string `json:"qualificatif"`
}

type CurrentVigilance struct {
	City string `json:"commune"`
	Insee string `json:"code_insee"`
	Vigilance Vigilance `json:"vigilances"`
}

type Vigilance struct {
	StartAt string `json:"date_debut"`
	EndAt string `json:"date_fin"`
	Name string `json:"nom_procedure"`
	Area string `json:"zone"`	
	Level string `json:"niveau"`
	Polluant string `json:"polluant"`
	Threshold string `json:"seuil"`
	Comment string `json:"commentaire"`
	UpdateAt string `json:"date_modification"`
}
 