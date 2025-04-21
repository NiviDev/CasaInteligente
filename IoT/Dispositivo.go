package IoT

import "errors"

type Dispositivo struct {
	Nombre string
	Estado bool
}

func (d *Dispositivo) Encender() error {
	if (*d).Estado {
		return errors.New("el dispositivo ya esta encendido")
	}
	(*d).Estado = true
	return nil
}

func (d *Dispositivo) Apagar() error {
	if !(*d).Estado {
		return errors.New("el dispositivo ya esta apagado")
	}
	(*d).Estado = false
	return nil
}

func (d *Dispositivo) EstadoActual() string {
	if (*d).Estado {
		return "Encendido"
	}
	return "Apagado"
}
