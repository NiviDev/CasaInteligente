package IoT

type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}
