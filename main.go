package main

import (
	"fmt"
	"mitacs/m_helper"
)

func Place_containers(Edge_list []m_helper.Edge, Container_list []m_helper.Container) {
	for c := range Container_list {
		Edge_list = m_helper.Nextfit(Edge_list, Container_list[c])
	}

	fmt.Println(Edge_list)
	fmt.Println(Container_list)
	fmt.Println("TPL :", m_helper.Total_power_loss(Edge_list))
}

func main() {
	m_helper.Edge_num = 0
	fmt.Println("Container placement in edge computing using Best Fit Algorithm")
	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(200, 200))
	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(80, 150))
	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(120, 100))
	m_helper.Edge_list = append(m_helper.Edge_list, m_helper.Initialise_edge(100, 100))

	//Initialising containers
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 40))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 30))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(20, 30))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(50, 30))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))
	m_helper.Container_list = append(m_helper.Container_list, m_helper.Initialise_container(10, 20))

	// Bestfit(Edge_list, Container_list)
	// Firstfit(Edge_list, Container_list)
	// fmt.Println(Edge_list)
	// fmt.Println("TPL :", Total_power_loss(Edge_list))

	Place_containers(m_helper.Edge_list, m_helper.Container_list)
}
