package main

type celsius float64

func (c celsius) farenheight() farenheight {
	return farenheight((c * 9 / 5) + 32)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

type farenheight float64

func (f farenheight) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

func (f farenheight) kelvin() kelvin {
	return f.celsius().kelvin()
}

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) farenheight() farenheight {
	return k.celsius().farenheight()
}
