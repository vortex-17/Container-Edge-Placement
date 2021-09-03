package main

import (
	"fmt"
	// "mitacs/m_helper"
	"mitacs/cbp"
)

// func Place_containers(Edge_list []m_helper.Edge, Container_list []m_helper.Container) {
// 	for c := range Container_list {
// 		Edge_list = m_helper.Nextfit(Edge_list, Container_list[c])
// 	}

// 	fmt.Println(Edge_list)
// 	fmt.Println(Container_list)
// 	fmt.Println("TPL :", m_helper.Total_power_loss(Edge_list))
// }

// func main() {
// 	m_helper.Edge_num = 0
// 	fmt.Println("Container placement in edge computing using Bin packing Algorithms")
// 	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(200, 200))
// 	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(80, 150))
// 	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(120, 100))
// 	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(100, 100))

// 	//Initialising containers
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 40))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 30))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(20, 30))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(50, 30))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
// 	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))

// 	Place_containers(m_helper.Edge_list, m_helper.Container_list)
// }

func place_container(Edge_list []cbp.Edge, Container_list []cbp.Container) {
	fmt.Println("placing containers...")
	for c := range Container_list {
		// fmt.Println("placing container : ", Container_list[c])
		// Edge_list = cbp.Class_constrained_best_fit(Edge_list, Container_list[c])
		Edge_list = cbp.Best_fit(Edge_list, Container_list[c])
		// fmt.Printf("Edge nodes after placement: %+v", Edge_list)
		// fmt.Println()
	}

	// fmt.Println()
	// fmt.Printf("%+v", Edge_list)
	cbp.Print_edges(Edge_list)
	fmt.Println()
	fmt.Println("Total resource loss: ", cbp.Total_resource_loss(Edge_list))
	// fmt.Println("Activating scaling function...")
	// ecl := make([]cbp.Container, 0)
	// Edge_list, ecl = cbp.Random_scaling_event(0.2, Edge_list)
	// fmt.Printf("Edge list after scale up event: %+v", Edge_list)
	// fmt.Println("\n\nContainers kicked out: ", ecl)

}

func main() {
	cbp.Edge_num = 0
	cbp.Edge_list = cbp.StartEdges("edge.csv")
	// fmt.Println(cbp.Edge_list)

	cbp.Container_num = 0
	cbp.Container_list = cbp.StartContainers("container.csv")
	// fmt.Println(cbp.Container_list)

	place_container(cbp.Edge_list, cbp.Container_list)

}
