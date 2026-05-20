package main

import (
	"fmt"
	"time"
)

type ParkingSpot struct {
	floor int
	spot_id int
	occupied bool
	vehicle string
	entry_time time.Time
}

type ParkingSystem struct {
	parking_lot [5][100]*ParkingSpot
	floor_stats [5]int
}

func NewParkingSystem() *ParkingSystem {
	p := &ParkingSystem{
		floor_stats: [5]int{100,100,100,100,100},
	}
	for i,_ := range p.parking_lot {
		for j,_ := range p.parking_lot[i] {
			p.parking_lot[i][j] = &ParkingSpot{
				floor: i,
				spot_id: j,
			}
		}
	}

	return p
}

func(p *ParkingSystem) Entry(floor int, vehicle string) error {
	// just find the the nearest spot on the floor to get itself filled
	if floor < 0 || floor > 4 {
		return fmt.Errorf("Not valid floor: %v", floor)
	}

	if p.floor_stats[floor] == 0 {
		return fmt.Errorf("No free slot available")
	}

	// now find the nearest empty place
	for i, v := range p.parking_lot[floor] {
		if !v.occupied {
			v := p.parking_lot[floor][i]
			v.occupied = true
			v.vehicle = vehicle
			v.entry_time = time.Now()
			p.floor_stats[floor]--;

			fmt.Printf("Vehicle %v parked at floor %v, spot: %v", vehicle, floor, i)
			break 
		}
	}

	return nil
}

func(p *ParkingSystem) Exit(vehicle string) (error) {
	floor, spot_id := p.search(vehicle)
	// handle errors to check vehicle is present or not and then update the stats and parkingspot
	if spot_id < 0 {
		return fmt.Errorf("no vehicle: %v found", vehicle)
	}

	// now just return the statement with the duration
	// Vehicle MH-12-AB-1234 exited. Duration: 2h 15m. Fee: ₹60, 20 rupees/hr -> convert to minutes: 1 rupee / 3 min
	duration := time.Since(p.parking_lot[floor][spot_id].entry_time)
	
	hour := int(duration.Hours())
	min := int(duration.Minutes()) % 60
	p.floor_stats[floor]++;

	cost := duration.Minutes() * 1/3.0
	fmt.Printf("Vehicle %v exited. Duration: %vh %vm. Fee: ₹%v\n", vehicle, hour, min, cost)

	return nil
}

func (p *ParkingSystem) search(vehicle string) (int,int) {
	parking := p.parking_lot

	for i:=0;i<len(parking);i++ {
		for idx, val := range parking[i] {
			if val.vehicle == vehicle {
				return i, idx
			}
		}
	}

	return -1,-1
}

func (p *ParkingSystem) Status() {
	// print the filled/total spots per floor
	// Floor 1: 73/100 available
	for i,v := range p.floor_stats {
		if i < len(p.floor_stats) - 1 {
			fmt.Printf("Floor %v: %v/100 available | ", i+1, v)
		} else {
			fmt.Printf("Floor %v: %v/100 available\n", i+1, v)
		}
	}
}

func callParkingManagamentSystem() {
	p := NewParkingSystem()

	// 1. enter
	fmt.Println("--- Step 1: parking vehicles ---")
	parkings := []struct {
		plate string
		floor int
	}{
		{"MH-12-AB-1234", 2},
		{"DL-08-CD-5678", 0},
		{"KA-01-EF-9012", 4},
	}
	for _, entry := range parkings {
		if err := p.Entry(entry.floor, entry.plate); err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println()
	}

	// 2. status
	fmt.Println("--- Step 2: status after parking ---")
	p.Status()
	fmt.Println()

	// 3. exit one vehicle
	time.Sleep(2 * time.Second)

	fmt.Println("--- Step 3: vehicle exit ---")
	if err := p.Exit("MH-12-AB-1234"); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println()

	// 5. Search
	fmt.Println("--- Step 4: search by plate ---")
	floor, spot := p.search("DL-08-CD-5678")
	if spot >= 0 {
		fmt.Printf("Found DL-08-CD-5678 at Floor %d, Spot %d\n", floor+1, spot)
	} else {
		fmt.Println("Vehicle not found")
	}
	fmt.Println()

	// 6. Final status.
	fmt.Println("--- Step 6: final status ---")
	p.Status()
}
