$fa=12.000000;
$fs=2.000000;
$fn=0;
difference(){
union(){
union(){
translate([0.000000, 7.825000, 2.000000]) {
cube([17.150000, 1.500000, 4.000000], center=true);
}
translate([7.825000, 0.000000, 2.000000]) {
cube([1.500000, 17.150000, 4.000000], center=true);
}
}
mirror([1.000000, 0.000000, 0.000000]) {
union(){
translate([0.000000, 7.825000, 2.000000]) {
cube([17.150000, 1.500000, 4.000000], center=true);
}
translate([7.825000, 0.000000, 2.000000]) {
cube([1.500000, 17.150000, 4.000000], center=true);
}
}
}
mirror([0.000000, 1.000000, 0.000000]) {
union(){
translate([0.000000, 7.825000, 2.000000]) {
cube([17.150000, 1.500000, 4.000000], center=true);
}
translate([7.825000, 0.000000, 2.000000]) {
cube([1.500000, 17.150000, 4.000000], center=true);
}
}
}
}
union(){
translate([5.660000, 0.000000, 1.250000]) {
cube([5.000000, 5.000000, 2.500000], center=true);
}
mirror([1.000000, 0.000000, 0.000000]) {
translate([5.660000, 0.000000, 1.250000]) {
cube([5.000000, 5.000000, 2.500000], center=true);
}
}
mirror([0.000000, 1.000000, 0.000000]) {
translate([5.660000, 0.000000, 1.250000]) {
cube([5.000000, 5.000000, 2.500000], center=true);
}
}
}
}
