package algogram

import (
	heap "tdas/cola_prioridad"
)

type Usuario struct {
	nombre string
	feed *heap.ColaPrioridad[]
}

func PublicarPost()
func VerProximoPost()
func LikearPost()