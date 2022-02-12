package medidas //Creacion del paquete medidas

//Definimos tres constantes para realizar conversiones
const KmPorMIlla = 1.60934
const MetrosPorPie = 0.3048
const CentimetroPorPulgada = 2.54

// CReamos tres funciones pra la conversion de unidades del sistema inglés a sistema decimal 
func MIllasEnKm(s float64) float64{
	return s * KmPorMIlla
}

func PiesEnMetros(s float64) float64{
	return s * MetrosPorPie
}

func PulgadasENCentimetros(s float64) float64{
	return s * CentimetroPorPulgada
}