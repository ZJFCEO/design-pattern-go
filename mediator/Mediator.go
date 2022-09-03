package main

type MediatorX struct {
	Ccpu  *CPU
	Cgpu  *GPU
	Cdisk *Disk
	Cmem  *MEM
}

var mediator *MediatorX = nil

func GetMediatorInstace() *MediatorX {
	if mediator == nil {
		mediator = &MediatorX{nil, nil, nil, nil}
		mediator.Ccpu = &CPU{}
		mediator.Cgpu = &GPU{}
		mediator.Cmem = &MEM{}
		mediator.Cdisk = &Disk{}
	}
	return mediator
}
func (m *MediatorX) Changed(i interface{}) {
	switch inst := i.(type) {
	case *CPU:
		m.Ccpu.Process(inst.data)
	case *GPU:
		m.Cgpu.Display(inst.data)
	case *Disk:
		m.Cdisk.Store(inst.data)
	case *MEM:
		m.Cmem.Dump(inst.data)
	default:

	}
}
