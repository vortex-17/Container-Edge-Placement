package main

import (
	"fmt"
	// "mitacs/m_helper"
	"mitacs/cbp"
)

func place_container(Edge_list []cbp.Edge, Container_list []cbp.Container, val int) {
	fmt.Println("placing containers...")
	Edge_list = cbp.Sort_Edge(Edge_list)
	for c := range Container_list {
		// fmt.Println("placing container : ", Container_list[c])
		if val == 0 {
			Edge_list = cbp.Class_constrained_best_fit(Edge_list, Container_list[c], false)
		} else {
			Edge_list = cbp.Best_fit(Edge_list, Container_list[c], false)
		}

		// fmt.Printf("Edge nodes after placement: %+v", Edge_list)
		// fmt.Println()
	}

	Edge_list = cbp.Sort_Edge(Edge_list)
	fmt.Println("Before Scaling")
	fmt.Println("Total resource loss: ", cbp.Total_resource_loss(Edge_list))
	used, unused := cbp.No_nodes_used(Edge_list)
	fmt.Println("Nodes used and unused : ", used, unused)

	//Activating scaling function
	ecl := make([]cbp.Container, 0)
	Edge_list = cbp.Sort_Edge(Edge_list)
	Edge_list, ecl = cbp.Random_scaling_event(0.95, Edge_list)

	// fmt.Println("\n\nContainers kicked out: ", len(ecl))
	container_kicked, a, na, tr := cbp.Containers_kicked(ecl)
	fmt.Printf("\nContainer kicked :%d\nContainers with autoscale: %d\nContainers without autoscale: %d\nTotal resource migration: %d\n", container_kicked, a, na, tr)

	Edge_list = cbp.Sort_Edge(Edge_list)

	for c := range ecl {

		if val == 0 {
			Edge_list = cbp.Class_constrained_best_fit(Edge_list, ecl[c], true)
		} else {
			Edge_list = cbp.Best_fit(Edge_list, ecl[c], true)
		}
	}

	// cbp.Print_edges(Edge_list)
	fmt.Println()
	fmt.Println("After Scaling")
	fmt.Println("Total resource loss: ", cbp.Total_resource_loss(Edge_list))
	used, unused = cbp.No_nodes_used(Edge_list)
	fmt.Println("Nodes used and unused : ", used, unused)

}

func main() {
	cbp.Edge_num = 0
	cbp.Edge_list = cbp.StartEdges("Data/edge2.csv")
	// fmt.Println(cbp.Edge_list)

	cbp.Container_num = 0
	var autoscale, nonautoscale int
	cbp.Container_list, autoscale, nonautoscale = cbp.StartContainers("Data/container2.csv")
	// fmt.Println(cbp.Container_list)

	fmt.Println("No of Edge Nodes: ", len(cbp.Edge_list))
	fmt.Println("Number of Autoscale containers: ", autoscale)
	fmt.Println("Number of Non-autoscale containers: ", nonautoscale)
	fmt.Println("******************************************")
	fmt.Println("Class Contrained best fit algorithm")
	place_container(cbp.Edge_list, cbp.Container_list, 0)
	fmt.Println("*******************************************")
	fmt.Println("best fit algorithm")
	place_container(cbp.Edge_list, cbp.Container_list, 1)
	fmt.Println("*******************************************")

}
