package cbp

//this file contains class based placement policy code

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type Edge struct {
	Id         int
	R_max      int
	R_ul       int
	R_cc       int
	R_total    int
	Containers []Container
	Active     bool
}

type Container struct {
	Container_class int //Class A = 0 || Class B = 1
	Id              int
	Active          bool
	Max_r           int //Max resource requirement if autoscaling
	Min_r           int //Min resource requirement if autoscaling
}

var Edge_num int
var Edge_list []Edge

var Container_num int
var Container_list []Container
var Excess_container_list []Container

func Initialise_edge(resources, R_ul int) Edge {
	edge := Edge{
		Id:         Edge_num,
		R_total:    resources,
		R_ul:       R_ul,
		R_max:      0,
		R_cc:       0,
		Containers: make([]Container, 0),
		Active:     false,
	}

	Edge_num++
	return edge
}

func Initialise_container(min_r, max_r, Container_class int) Container {

	container := Container{
		Id:              Container_num,
		Container_class: Container_class,
		Min_r:           min_r,
		Max_r:           max_r,
		Active:          false,
	}

	Container_num++
	return container
}

func Check_resource_constraints(e Edge, c Container, scaled bool) bool {
	if scaled == true {
		if e.R_total-e.R_cc >= c.Max_r {
			return true
		}
		return false
	} else {
		if e.R_total-e.R_cc >= c.Min_r {
			return true
		}
		return false
	}

}

func Check_resource_constraints_Scaled(e Edge, c Container) bool {
	if e.R_total-e.R_cc >= c.Max_r {
		return true
	}
	return false
}

func Sort_data(Edge_list []Edge, c Container) []Edge {
	// We are sorting the data to find the argmin |(resource left at peak level) - container_max/min|
	var val int
	if c.Container_class == 0 {
		val = c.Max_r
	} else {
		val = c.Min_r
	}

	for i := 0; i < len(Edge_list); i++ {
		min := i
		for j := i; j < len(Edge_list); j++ {
			if math.Abs((float64(Edge_list[j].R_total - Edge_list[j].R_max + val))) <= math.Abs((float64(Edge_list[min].R_total - Edge_list[min].R_max + val))) {
				min = j
			}
		}

		Edge_list[i], Edge_list[min] = Edge_list[min], Edge_list[i]
	}

	return Edge_list

}

func Sort_data_BF(Edge_list []Edge) []Edge {
	for i := 0; i < len(Edge_list); i++ {
		min := i
		for j := i; j < len(Edge_list); j++ {
			if (Edge_list[j].R_total - Edge_list[j].R_cc) <= (Edge_list[min].R_total - Edge_list[min].R_cc) {
				min = j
			}
		}

		Edge_list[i], Edge_list[min] = Edge_list[min], Edge_list[i]
	}

	return Edge_list

}

func Total_resource_loss(Edge_list []Edge) int {
	tpl := 0
	for e := range Edge_list {
		if len(Edge_list[e].Containers) > 0 {
			tpl += (Edge_list[e].R_total - Edge_list[e].R_cc)
		}
	}

	return tpl
}

//this function returns the list of active and inactive edge nodes from the list of edge nodes.
func Active_inactive_list(Edge_list []Edge) ([]Edge, []Edge) {
	//This function returns the list of active and inactive lists
	var Active_list []Edge
	var Inactive_list []Edge
	for e := range Edge_list {
		if Edge_list[e].Active == true {
			Active_list = append(Active_list, Edge_list[e])
		} else {
			Inactive_list = append(Inactive_list, Edge_list[e])
		}
	}

	return Active_list, Inactive_list
}

func Class_constrained_best_fit(Edge_list []Edge, c Container, scaled bool) []Edge {

	active, inactive := Active_inactive_list(Edge_list)
	if len(active) < 1 {
		// No edge nodes are active now. We will now active an edge node

		// fmt.Println("No edge nodes are active now. We will now active an edge node")
		inactive[0].Active = true

		// active = append(active, inactive...) //merging active and inactive list

		Edge_list = Class_constrained_best_fit(append(active, inactive...), c, scaled)
		// fmt.Println("Return list :", Edge_list)
		return Edge_list

	} else {

		//There are few active nodes
		//We sort the edge nodes because we want to find the node which has the tightest space.

		active := Sort_data(active, c)
		// fmt.Println("??ctive :", active)

		for e := range active {
			if Check_resource_constraints(active[e], c, scaled) == true {
				if c.Container_class == 0 { //class A
					active[e].R_max += c.Max_r
					if scaled == true {
						active[e].R_cc += c.Max_r
					} else {
						active[e].R_cc += c.Min_r
					}
				} else {
					active[e].R_max += c.Min_r
					active[e].R_cc += c.Min_r
				}
				active[e].Active = true
				// active[e].R_ul -= c.Min_r
				// active[e].R_cc += c.Min_r
				active[e].Containers = append(active[e].Containers, c)
				c.Active = true
				break

			}
		}

		if c.Active == false {

			// Could not place the container into any of the active edge node
			// fmt.Println("Could not place the container into any of the active edge node")

			// Need to activate another node
			if len(inactive) > 0 {
				inactive[0].Active = true
				// active = append(active, inactive...) //merging active and inactive list
				Edge_list = Class_constrained_best_fit(append(active, inactive...), c, scaled)
				return Edge_list
			} else {
				// fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

			}
		}

	}

	// active = append(active, inactive...)
	// fmt.Println("Result list: ", append(active, inactive...))
	return append(active, inactive...)

}

func Best_fit(Edge_list []Edge, c Container, scaled bool) []Edge {

	active, inactive := Active_inactive_list(Edge_list)
	if len(active) < 1 {
		// No edge nodes are active now. We will now active an edge node

		// fmt.Println("No edge nodes are active now. We will now active an edge node")
		inactive[0].Active = true

		// active = append(active, inactive...) //merging active and inactive list

		Edge_list = Best_fit(append(active, inactive...), c, scaled)
		// fmt.Println("Return list :", Edge_list)
		return Edge_list

	} else {

		//There are few active nodes
		//We sort the edge nodes because we want to find the node which has the tightest space.

		active := Sort_data_BF(active)
		// fmt.Println("??ctive :", active)

		for e := range active {
			if Check_resource_constraints(active[e], c, scaled) == true {
				if c.Container_class == 0 { //class A
					active[e].R_max += c.Max_r
					if scaled == true {
						active[e].R_cc += c.Max_r
					} else {
						active[e].R_cc += c.Min_r
					}
				} else {
					active[e].R_max += c.Min_r
					active[e].R_cc += c.Min_r
				}
				active[e].Active = true
				// active[e].R_ul -= c.Min_r

				active[e].Containers = append(active[e].Containers, c)
				c.Active = true
				break

			}
		}

		if c.Active == false {

			// Could not place the container into any of the active edge node
			// fmt.Println("Could not place the container into any of the active edge node")

			// Need to activate another node
			if len(inactive) > 0 {
				inactive[0].Active = true
				// active = append(active, inactive...) //merging active and inactive list
				Edge_list = Best_fit(append(active, inactive...), c, scaled)
				return Edge_list
			} else {
				// fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

			}
		}

	}

	// active = append(active, inactive...)
	// fmt.Println("Result list: ", append(active, inactive...))
	return append(active, inactive...)

}

// func Class_constrained_best_fit2(Edge_list []Edge, c Container) []Edge {

// 	active, inactive := Active_inactive_list(Edge_list)
// 	if len(active) < 1 {
// 		// No edge nodes are active now. We will now active an edge node

// 		fmt.Println("No edge nodes are active now. We will now active an edge node")
// 		inactive[0].Active = true

// 		// active = append(active, inactive...) //merging active and inactive list

// 		Edge_list = Class_constrained_best_fit2(append(active, inactive...), c)
// 		// fmt.Println("Return list :", Edge_list)
// 		return Edge_list

// 	} else {

// 		//There are few active nodes
// 		//We sort the edge nodes because we want to find the node which has the tightest space.

// 		active := Sort_data(active, c)
// 		// fmt.Println("??ctive :", active)
// 		for e := range active {
// 			if Check_resource_constraints_Scaled(active[e], c) == true {
// 				// fmt.Printf("Edge | container: %+v | \n%+v\n ", active[e], c)
// 				if c.Container_class == 0 { //class A
// 					active[e].R_max += c.Max_r
// 					active[e].R_cc += c.Max_r
// 				} else {
// 					active[e].R_max += c.Min_r
// 					active[e].R_cc += c.Min_r
// 				}
// 				active[e].Active = true
// 				// active[e].R_ul -= c.Min_r

// 				active[e].Containers = append(active[e].Containers, c)
// 				c.Active = true
// 				break

// 			}
// 		}

// 		if c.Active == false {

// 			// Could not place the container into any of the active edge node
// 			fmt.Println("Could not place the container into any of the active edge node")

// 			// Need to activate another node
// 			if len(inactive) > 0 {
// 				inactive[0].Active = true
// 				// active = append(active, inactive...) //merging active and inactive list
// 				Edge_list = Class_constrained_best_fit2(append(active, inactive...), c)
// 				return Edge_list
// 			} else {
// 				fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

// 			}
// 		}

// 	}

// 	// active = append(active, inactive...)
// 	// fmt.Println("Result list: ", append(active, inactive...))
// 	return append(active, inactive...)

// }

// func Best_fit2(Edge_list []Edge, c Container) []Edge {

// 	active, inactive := Active_inactive_list(Edge_list)
// 	if len(active) < 1 {
// 		// No edge nodes are active now. We will now active an edge node

// 		fmt.Println("No edge nodes are active now. We will now active an edge node")
// 		inactive[0].Active = true

// 		// active = append(active, inactive...) //merging active and inactive list

// 		Edge_list = Best_fit2(append(active, inactive...), c)
// 		// fmt.Println("Return list :", Edge_list)
// 		return Edge_list

// 	} else {

// 		//There are few active nodes
// 		//We sort the edge nodes because we want to find the node which has the tightest space.

// 		active := Sort_data_BB(active)
// 		// fmt.Println("??ctive :", active)

// 		for e := range active {
// 			if Check_resource_constraints_Scaled(active[e], c) == true {
// 				if c.Container_class == 0 { //class A
// 					active[e].R_max += c.Max_r
// 					active[e].R_cc += c.Max_r
// 				} else {
// 					active[e].R_max += c.Min_r
// 					active[e].R_cc += c.Min_r
// 				}
// 				active[e].Active = true
// 				// active[e].R_ul -= c.Min_r
// 				active[e].Containers = append(active[e].Containers, c)
// 				c.Active = true
// 				break

// 			}
// 		}

// 		if c.Active == false {

// 			// Could not place the container into any of the active edge node
// 			fmt.Println("Could not place the container into any of the active edge node")

// 			// Need to activate another node
// 			if len(inactive) > 0 {
// 				inactive[0].Active = true
// 				// active = append(active, inactive...) //merging active and inactive list
// 				Edge_list = Best_fit2(append(active, inactive...), c)
// 				return Edge_list
// 			} else {
// 				fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

// 			}
// 		}

// 	}

// 	// active = append(active, inactive...)
// 	// fmt.Println("Result list: ", append(active, inactive...))
// 	return append(active, inactive...)

// }

func Check_scaling(e Edge, c Container) bool {
	if e.R_cc+(c.Max_r-c.Min_r) <= e.R_total {
		return true
	}

	return false
}

//It will select containers with autoscaling feature and scale them for time t
func Random_scaling_event(percent float64, Edge_list []Edge) ([]Edge, []Container) {
	var Excess_container_list []Container
	scaling_edge := int(percent * float64(len(Edge_list)))
	// fmt.Println(scaling_edge)

	//Going through selected edge nodes to look for scaling containers
	for i := 0; i < scaling_edge; i++ {

		container_list := sort_containers(Edge_list[i].Containers)
		index := 0

		//going through containers in the ith edge node
		for _ = range container_list {

			container := Edge_list[i].Containers[index]

			//Will scale only the containers with class 0
			if container.Container_class == 0 {
				if Check_scaling(Edge_list[i], container) == true {
					Edge_list[i].R_cc += (container.Max_r - container.Min_r)
					index++
				} else {
					// fmt.Println("Edge node at max capacity! Cannot scale the container, Removing the container")

					//removing containers which cannot be scaled. They will be placed somewhere else
					Edge_list[i].R_max -= container.Max_r
					Edge_list[i].R_cc -= container.Min_r
					Excess_container_list = append(Excess_container_list, container)

					Edge_list[i].Containers = append(Edge_list[i].Containers[:index], Edge_list[i].Containers[index+1:]...)
				}
			} else {
				index++
			}
		}
		// Edge_node = Edge_list[i]
		// fmt.Printf("\nAfter Scaling for edge node %+v", Edge_node)
		// fmt.Printf("\n")
	}

	return Edge_list, Excess_container_list

}

//It will read from file
func parse(filename string) *csv.Reader {

	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvfile)
	return r

}

func StartEdges(filename string) []Edge {
	r := parse(filename)
	var E_list []Edge
	for {
		data, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(data)
		r_total, err1 := strconv.Atoi(data[0])
		r_ul, err2 := strconv.Atoi(data[1])

		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}

		//Initialise edge
		E_list = append(E_list, Initialise_edge(r_total, r_ul))
	}

	return E_list
}

func StartContainers(filename string) ([]Container, int, int) {
	r := parse(filename)
	var autoscale, nonautoscale int
	var C_list []Container
	for {
		data, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(data)
		class, err0 := strconv.Atoi(data[0])
		min_r, err1 := strconv.Atoi(data[1])
		max_r, err2 := strconv.Atoi(data[2])

		if err1 != nil || err2 != nil || err0 != nil {
			log.Fatal(err1, err2)
		}

		if class == 0 {
			autoscale++
		} else {
			nonautoscale++
		}

		//Initialise containr
		C_list = append(C_list, Initialise_container(min_r, max_r, class))
	}

	return C_list, autoscale, nonautoscale
}

func Print_edges(Edge_list []Edge) {
	for i := range Edge_list {
		fmt.Printf("\n%+v\n", Edge_list[i])
	}
}

func Sort_Edge(Edge_list []Edge) []Edge {
	for i := 0; i < len(Edge_list); i++ {
		min := i
		for j := i; j < len(Edge_list); j++ {
			if Edge_list[j].Id <= Edge_list[min].Id {
				min = j
			}
		}

		Edge_list[i], Edge_list[min] = Edge_list[min], Edge_list[i]
	}

	return Edge_list
}

func No_nodes_used(Edge_list []Edge) (int, int) {
	used, unused := 0, 0
	for e := range Edge_list {
		if len(Edge_list[e].Containers) > 0 {
			used++
		} else {
			unused++
		}
	}

	return used, unused
}

func Containers_kicked(Container_list []Container) (int, int, int, int) {
	autoscale, nonautoscale := 0, 0
	tr := 0
	for c := range Container_list {
		if Container_list[c].Container_class == 0 {
			autoscale++
			tr += Container_list[c].Max_r
		} else {
			nonautoscale++
			tr += Container_list[c].Min_r
		}
	}

	return len(Container_list), autoscale, nonautoscale, tr
}

func sort_containers(Container_list []Container) []Container {

	for i := 0; i < len(Container_list); i++ {
		max := i
		for j := i; j < len(Container_list); j++ {
			if Container_list[j].Max_r >= Container_list[max].Max_r {
				max = j
			}
		}

		Container_list[i], Container_list[max] = Container_list[max], Container_list[i]
	}

	return Container_list

}
