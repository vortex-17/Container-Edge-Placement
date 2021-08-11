package m_helper

import (
	"fmt"
)

type Edge struct {
	Id         int
	Resources  int
	Power      int
	Containers int
	Active     bool
}

type Container struct {
	Id        int
	Resources int
	Power     int
	Active    bool
}

var Edge_num int
var Edge_list []Edge

var Container_num int
var Container_list []Container

func Initialise_edge(resources, power int) Edge {
	edge := Edge{
		Id:         Edge_num,
		Resources:  resources,
		Power:      power,
		Containers: 0,
		Active:     false,
	}

	Edge_num++
	return edge
}

func Initialise_container(resources, power int) Container {
	container := Container{
		Id:        Container_num,
		Resources: resources,
		Power:     power,
		Active:    false,
	}

	Container_num++
	return container
}

func Check_resource_constraints(e Edge, c Container) bool {
	if e.Resources > c.Resources {
		if e.Power > c.Power {
			return true
		}
	}
	return false
}

func Sort_data(Edge_list []Edge) []Edge {
	for i := 0; i < len(Edge_list); i++ {
		min := i
		for j := i; j < len(Edge_list); j++ {
			if Edge_list[j].Power <= Edge_list[min].Power {
				min = j
			}
		}

		Edge_list[i], Edge_list[min] = Edge_list[min], Edge_list[i]
	}

	return Edge_list

}

func Total_power_loss(Edge_list []Edge) int {
	tpl := 0
	for e := range Edge_list {
		if Edge_list[e].Active {
			tpl += Edge_list[e].Power
		}
	}

	return tpl
}

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

// #########################################################################
// Bin Packing Algorithms

// Online Algorithms
// #########################################################################

func Bestfit(Edge_list []Edge, c Container) []Edge {
	// fmt.Println(c)
	active, inactive := Active_inactive_list(Edge_list)
	if len(active) < 1 {
		// No edge nodes are active now. We will now active an edge node
		fmt.Println("No edge nodes are active now. We will now active an edge node")
		inactive[0].Active = true
		// active = append(active, inactive...) //merging active and inactive list
		Edge_list = Bestfit(append(active, inactive...), c)
		// fmt.Println("Return list :", Edge_list)
		return Edge_list
	} else {
		//There are few active nodes

		active := Sort_data(active)
		// fmt.Println("Äctive :", active)
		for e := range active {
			if Check_resource_constraints(active[e], c) == true {
				active[e].Active = true
				active[e].Power -= c.Power
				active[e].Resources -= c.Resources
				active[e].Containers++
				c.Active = true
				break

			}
		}

		if c.Active == false {
			// Could not place the container into any of the active edge node
			fmt.Println("Could not place the container into any of the active edge node")
			// Need to activate another node
			if len(inactive) > 0 {
				inactive[0].Active = true
				// active = append(active, inactive...) //merging active and inactive list
				Edge_list = Bestfit(append(active, inactive...), c)
				return Edge_list
			} else {
				fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

			}
		}

	}

	// active = append(active, inactive...)
	// fmt.Println("Result list: ", append(active, inactive...))
	return append(active, inactive...)

}

func Firstfit(Edge_list []Edge, c Container) []Edge {
	// fmt.Println(c)
	active, inactive := Active_inactive_list(Edge_list)
	if len(active) < 1 {
		// No edge nodes are active now. We will now active an edge node
		fmt.Println("No edge nodes are active now. We will now active an edge node")
		inactive[0].Active = true
		// active = append(active, inactive...) //merging active and inactive list
		Edge_list = Firstfit(append(active, inactive...), c)
		// fmt.Println("Return list :", Edge_list)
		return Edge_list
	} else {
		//There are few active nodes
		// fmt.Println("Äctive :", active)
		for e := range active {
			if Check_resource_constraints(active[e], c) == true {
				active[e].Active = true
				active[e].Power -= c.Power
				active[e].Resources -= c.Resources
				active[e].Containers++
				c.Active = true
				break

			}
		}

		if c.Active == false {
			// Could not place the container into any of the active edge node
			fmt.Println("Could not place the container into any of the active edge node")
			// Need to activate another node
			if len(inactive) > 0 {
				inactive[0].Active = true
				// active = append(active, inactive...) //merging active and inactive list
				Edge_list = Firstfit(append(active, inactive...), c)
				return Edge_list
			} else {
				fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

			}
		}

	}

	// active = append(active, inactive...)
	fmt.Println("Result list: ", append(active, inactive...))
	return append(active, inactive...)

}

func Worstfit(Edge_list []Edge, c Container) []Edge {
	// fmt.Println(c)
	active, inactive := Active_inactive_list(Edge_list)
	if len(active) < 1 {
		// No edge nodes are active now. We will now active an edge node
		fmt.Println("No edge nodes are active now. We will now active an edge node")
		inactive[0].Active = true
		// active = append(active, inactive...) //merging active and inactive list
		Edge_list = Worstfit(append(active, inactive...), c)
		// fmt.Println("Return list :", Edge_list)
		return Edge_list
	} else {
		//There are few active nodes
		active := Sort_data(active)

		//reversing the active list
		for i, j := 0, len(active)-1; i < j; i, j = i+1, j-1 {
			active[i], active[j] = active[j], active[i]
		}

		// fmt.Println("Äctive :", active)
		for e := range active {
			if Check_resource_constraints(active[e], c) == true {
				active[e].Active = true
				active[e].Power -= c.Power
				active[e].Resources -= c.Resources
				active[e].Containers++
				c.Active = true
				break

			}
		}

		if c.Active == false {
			// Could not place the container into any of the active edge node
			fmt.Println("Could not place the container into any of the active edge node")
			// Need to activate another node
			if len(inactive) > 0 {
				inactive[0].Active = true
				// active = append(active, inactive...) //merging active and inactive list
				Edge_list = Worstfit(append(active, inactive...), c)
				return Edge_list
			} else {
				fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

			}
		}

	}

	// active = append(active, inactive...)
	fmt.Println("Result list: ", append(active, inactive...))
	return append(active, inactive...)

}

func Nextfit(Edge_list []Edge, c Container) []Edge {

	// When processing next item, check if it fits in the same bin as the last item.
	// Use a new bin only if it does not.

	// fmt.Println(c)
	active, inactive := Active_inactive_list(Edge_list)
	if len(active) < 1 {
		// No edge nodes are active now. We will now active an edge node
		fmt.Println("No edge nodes are active now. We will now active an edge node")
		inactive[0].Active = true
		// active = append(active, inactive...) //merging active and inactive list
		Edge_list = Nextfit(append(active, inactive...), c)
		// fmt.Println("Return list :", Edge_list)
		return Edge_list
	} else {
		//There are few active nodes
		// fmt.Println("Äctive :", active)
		for e := range active {
			if Check_resource_constraints(active[len(active)-1], c) == true {
				active[e].Active = true
				active[e].Power -= c.Power
				active[e].Resources -= c.Resources
				active[e].Containers++
				c.Active = true
				break

			}
		}

		if c.Active == false {
			// Could not place the container into any of the active edge node
			fmt.Println("Could not place the container into any of the active edge node")
			// Need to activate another node
			if len(inactive) > 0 {
				inactive[0].Active = true
				// active = append(active, inactive...) //merging active and inactive list
				Edge_list = Nextfit(append(active, inactive...), c)
				return Edge_list
			} else {
				fmt.Println("Sorry, no edge nodes can accomodate the container ", c.Id)

			}
		}

	}

	// active = append(active, inactive...)
	fmt.Println("Result list: ", append(active, inactive...))
	return append(active, inactive...)

}
