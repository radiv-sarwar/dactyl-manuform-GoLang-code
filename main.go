package main

import (
	"math"

	// . "github.com/ljanyst/ghostscad/lib/shapes"

	. "github.com/go-gl/mathgl/mgl64"
	// . "github.com/ljanyst/ghostscad/lib/degmath"
	. "github.com/ljanyst/ghostscad/primitive"
	"github.com/ljanyst/ghostscad/sys"
)

func main() {
	//	sys.Initialize()
	//	sys.RenderOne(NewSphere(10))
	// a := 4.0
	// b := 2.0
	// c := a + b
	// fmt.Println(c)

	nrows := 5
	ncols := 7

	alpha := (math.Pi / 12)
	Beta := (math.Pi / 36)
	// centerrow := (nrows - 3)
	// centercol := 4.0
	// tenting_angle := (math.Pi/12)
	// pinky_15u := true
	// first_15u_row := 3.0
	// extra_row := true
	// inner_column := true
	// mini_thumb := false

	// column_size := :standard ???????????
	// if inner_column == true {
	// 	var column_offset [3]float64
	// 	if ncols <= 1 {
	// 		column_offset = [3]float64{0.0, -2.0, 0.0}
	// 	} else if ncols == 3 || ncols == 2 {
	// 		column_offset = [3]float64{0.0, 2.82, -4.5}
	// 	} else if ncols >= 4 {
	// 		column_offset = [3]float64{0.0, -12.0, 5.64}
	// 	} else {
	// 		column_offset = [3]float64{0, 0, 0}
	// 	}
	// }
	// thumb_offsets := [3]float64{6, -3, 7}
	// keyboard_z_offset := 11.0

	extra_width := 2.5
	extra_height := 1.0
	// wall_z_offset := -8.0
	// wall_xy_offset := 5.0
	// wall-thickness := 2.0

	// fixed_angles := [7]float64{deg2rad(10), deg2rad(10), 0, 0, 0, deg2rad(-15), deg2rad(-15)}
	// fixed_x := [7]float64{-41.5, -22.5, 0, 20.3, 41.4, 65.5, 89.6}
	// fixed_z := [7]float64{12.1, 8.3, 0, 5, 10.7, 14.5, 17.5}
	// fixed_tenting := deg2rad(0)

	// GENERAL VARIABLES

	innercol_offset := 1

	// create_side_nubs := false

	// SWITCH HOLE //////////////////////////
	keyswitch_height := 14.15
	keyswitch_width := 14.15

	plate_thickness := 4.0
	// side_nub_thickness := 4.0
	retention_tab_thickness := 1.5
	retention_tab_hole_thickness := plate_thickness - retention_tab_thickness
	mount_width := keyswitch_width + 3.2
	mount_height := keyswitch_height + 2.7

	// for i := 0; i < 8; i++ {
	// 	println(fixed_angles[i])
	// }

	// 1) learn how to input shape in variable.
	// shape1 := NewCube(Vec3{5, 5, 5})

	// 2) how would you then make unions, differences and hulls?
	// shape2 := NewSphere(4.0)
	// shape3 := NewDifference(shape1, shape2)
	// slice1 := Vec3{40, 20, 30}

	// 3) try translating and rotating the shape variables.
	// shape4 := NewRotation(Vec3{50, 50, 50}, shape3)
	// sys.RenderOne(shape4)
	// for i := 0; i < 3; i++ {
	// 	println(slice1[i])
	// }
	// 4) make a unit switch hole

	top_wall_shape := NewCube(Vec3{keyswitch_width + 3, 1.5, plate_thickness})
	top_wall := NewTranslation(Vec3{0, (1.5 / 2) + (keyswitch_height / 2), plate_thickness / 2}, top_wall_shape)

	left_wall_shape := NewCube(Vec3{1.5, keyswitch_height + 3, plate_thickness})
	left_wall := NewTranslation(Vec3{(1.5 / 2) + (keyswitch_width / 2), 0, (plate_thickness / 2)}, left_wall_shape)
	plate_half := NewUnion(top_wall, left_wall)
	plate_other_half1 := NewMirror(Vec3{1, 0, 0}, plate_half)
	plate_other_half2 := NewMirror(Vec3{0, 1, 0}, plate_half)
	plate_full := NewUnion(plate_half, plate_other_half1, plate_other_half2)

	top_nub_shape := NewCube(Vec3{5, 5, retention_tab_hole_thickness})
	top_nub := NewTranslation(Vec3{(keyswitch_width / 2.5), 0, (retention_tab_hole_thickness / 2)}, top_nub_shape)
	top_nub_mirror1 := NewMirror(Vec3{1, 0, 0}, top_nub)
	top_nub_mirror2 := NewMirror(Vec3{0, 1, 0}, top_nub)
	top_nub_final := NewUnion(top_nub, top_nub_mirror1, top_nub_mirror2)

	switch_plate := NewDifference(plate_full, top_nub_final)

	var columns []int
	var rows []int

	for i := innercol_offset; i < int(ncols); i++ {
		columns = append(columns, i)
	}
	for i := 0; i < int(nrows); i++ {
		rows = append(rows, i)
	}

	// inner_column := 0
	var inner_rows []int

	for i := 0; i < (nrows - 2); i++ {
		inner_rows = append(inner_rows, i)
	}

	// for value := range columns {
	// 	println(columns[value])
	// }
	// for value := range rows {
	// 	println(rows[value])
	// }
	// for value := range inner_rows {
	// 	println(inner_rows[value])
	// }

	sa_profile_key_height := 12.7
	cap_top_height := plate_thickness + sa_profile_key_height

	row_radius := (((mount_height + extra_height) / 2) / math.Sin(alpha/2)) + cap_top_height
	column_radius := (((mount_width + extra_width) / 2) / math.Sin(Beta/2)) + cap_top_height
	column_x_delta := (column_radius * math.Sin(Beta/2)) - 1

	println(row_radius)
	println(column_radius)
	println(column_x_delta)

	sys.RenderOne(switch_plate)

}

func deg2rad(a float64) float64 {
	rad := ((a / 180) / math.Pi)
	return rad
}

// function declaration structure. For more than one returns, use brackets
// func funcName(parameter type, parameter type) return-type {
// 	...
// 	return ...
// }
