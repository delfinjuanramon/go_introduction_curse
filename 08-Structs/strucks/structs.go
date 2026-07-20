package structs

type Product struct {
	ID    uint
	Name  string
	Type  string
	Price float32
}

type NewProduct struct {
	ID    uint
	Name  string
	Type  Tipo
	Price float32
	Tags  []string
}

type Tipo struct {
	ID          uint
	Code        string
	Description string
}

// se pueden crear metodos para la estructura

// aqui solo mustra un resultado.  Es como un GET de Java
func (p Product) PrecioTotal() float64{   
	return float64(p.count) * float64(p.Price)
}

 // con el * se cambian los valores de la estructura.  Es como un SET de Java 
// solo cambian SET y GET con el *. 
// El asterisco que indica un puntero que almacena la dirección de memoria
// si no se le pone * el valor no se cambia en el Struct original, 
// solo en la copia que se pasa como referencia.

func (p *Product) CambiaNombre(nuevoNombre string) {  
	p.Name= nuevoNombre	
}



