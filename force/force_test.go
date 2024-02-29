package force

import (
	"testing"

	"github.com/arthurlee945/suhag/vec"
)

func TestAcceleration(t *testing.T) {
	v2Cases := []vec.Vec2{
		*vec.NewVec2(5, 2),
		*vec.NewVec2(24, 5),
	}
	v3Cases := []vec.Vec3{
		*vec.NewVec3(5, 2, 5),
		*vec.NewVec3(24, 5, 1),
	}

	for _, v2 := range v2Cases {
		acc := Acceleration(v2, 10)
		if acc[0] != v2[0]/10 || acc[1] != v2[1]/10 {
			v2.Div(10)
			t.Errorf("Expected acceleration to match vector / mass; expected=%v; got=%v", v2, acc)
		}
	}

	for _, v3 := range v3Cases {
		acc := Acceleration(v3, 8)
		if acc[0] != v3[0]/8 || acc[1] != v3[1]/8 {
			v3.Div(8)
			t.Errorf("Expected acceleration to match vector / mass; expected=%v; got=%v", v3, acc)
		}
	}
}

func TestForce(t *testing.T) {
	v2Cases := []vec.Vec2{
		*vec.NewVec2(5, 2),
		*vec.NewVec2(24, 5),
	}
	v3Cases := []vec.Vec3{
		*vec.NewVec3(5, 2, 5),
		*vec.NewVec3(24, 5, 1),
	}

	for _, v2 := range v2Cases {
		acc := Force(v2, 10)
		if acc[0] != v2[0]*10 || acc[1] != v2[1]*10 {
			v2.Div(10)
			t.Errorf("Expected Force to match vector * mass; expected=%v; got=%v", v2, acc)
		}
	}

	for _, v3 := range v3Cases {
		acc := Force(v3, 8)
		if acc[0] != v3[0]*8 || acc[1] != v3[1]*8 {
			v3.Div(8)
			t.Errorf("Expected Force to match vector * mass; expected=%v; got=%v", v3, acc)
		}
	}
}

func TestFriction(t *testing.T) {
	v2Cases := []*vec.Vec2{
		vec.NewVec2(5, 2),
		vec.NewVec2(24, 5),
	}
	v3Cases := []*vec.Vec3{
		vec.NewVec3(5, 2, 5),
		vec.NewVec3(24, 5, 1),
	}

	coef := 8.0
	normal := 2.0
	for _, v2 := range v2Cases {
		frc := Friction(v2, coef, normal)
		v2.Mult(-1)
		v2.Normalize()
		v2.Mult(coef * normal)
		if frc[0] != v2[0] || frc[1] != v2[1] {
			t.Errorf("Expected Friction to be -1 * coef * Normal * velocity vector (v2). expected=%v, got=%v", frc, v2)
		}
	}

	for _, v3 := range v3Cases {
		frc := Friction(v3, coef, normal)
		v3.Mult(-1)
		v3.Normalize()
		v3.Mult(coef * normal)
		if frc[0] != v3[0] || frc[1] != v3[1] || frc[2] != v3[2] {
			t.Errorf("Expected Friction to be -1 * coef * Normal * velocity vector (v2). expected=%v, got=%v", frc, v3)
		}
	}
}

func TestDrag(t *testing.T) {
	v2Cases := []*vec.Vec2{
		vec.NewVec2(5, 2),
		vec.NewVec2(24, 5),
	}
	v3Cases := []*vec.Vec3{
		vec.NewVec3(5, 2, 5),
		vec.NewVec3(24, 5, 1),
	}

	coef := 8.0
	crossSectional := 1.5
	density := 2.0
	for _, v2 := range v2Cases {
		drg := Drag(v2, density, crossSectional, coef)
		speed := v2.Mag()
		v2.Mult(-0.5)
		v2.Normalize()
		v2.Mult(density * crossSectional * coef * speed * speed)
		if drg[0] != v2[0] || drg[1] != v2[1] {
			t.Errorf("Expected Friction to be -0.5 * density * A * c * speed * speed. expected=%v, got=%v", v2, drg)
		}
	}

	for _, v3 := range v3Cases {
		drg := Drag(v3, density, crossSectional, coef)
		speed := v3.Mag()
		v3.Mult(-0.5)
		v3.Normalize()
		v3.Mult(density * crossSectional * coef * speed * speed)
		if drg[0] != v3[0] || drg[1] != v3[1] || drg[2] != v3[2] {
			t.Errorf("Expected Friction to be -0.5 * density * A * c * speed * speed. expected=%v, got=%v", v3, drg)
		}
	}
}

func TestAttraction(t *testing.T) {
	v2Cases := []*vec.Vec2{
		vec.NewVec2(5, 2),
		vec.NewVec2(24, 5),
	}
	v3Cases := []*vec.Vec3{
		vec.NewVec3(5, 2, 5),
		vec.NewVec3(24, 5, 1),
	}

	G, m1, m2 := 6.2, 23.0, 5.0

	attractionv2 := Attraction(G, m1, m2, *v2Cases[0], *v2Cases[1])
	attrValv2 := vec.Sub(*v2Cases[0], *v2Cases[1])
	distv2 := attrValv2.Mag()
	attrMag := -1 * (G * m1 * m2) / (distv2 * distv2)
	attrValv2.Normalize()
	attrValv2.Mult(attrMag)
	if attractionv2[0] != attrValv2[0] || attractionv2[1] != attrValv2[1] {
		t.Errorf("F = ((G * mass1 * mass2) / r ^ 2) * r^. expected=%v, got=%v", attrValv2, attractionv2)
	}

	attractionv3 := Attraction(G, m1, m2, *v3Cases[0], *v3Cases[1])
	attrValv3 := vec.Sub(*v3Cases[0], *v3Cases[1])
	distv3 := attrValv3.Mag()
	attrMagv3 := -1 * (G * m1 * m2) / (distv3 * distv3)
	attrValv3.Normalize()
	attrValv3.Mult(attrMagv3)
	if attractionv3[0] != attrValv3[0] || attractionv3[1] != attrValv3[1] {
		t.Errorf("F = ((G * mass1 * mass2) / r ^ 2) * r^. expected=%v, got=%v", attrValv3, attractionv3)
	}
}

func TestRepulsion(t *testing.T) {
	v2Cases := []*vec.Vec2{
		vec.NewVec2(5, 2),
		vec.NewVec2(24, 5),
	}
	v3Cases := []*vec.Vec3{
		vec.NewVec3(5, 2, 5),
		vec.NewVec3(24, 5, 1),
	}

	G, m1, m2 := 6.2, 23.0, 5.0

	repulsionv2 := Repulsion(G, m1, m2, *v2Cases[0], *v2Cases[1])
	repulValv2 := vec.Sub(*v2Cases[0], *v2Cases[1])
	distv2 := repulValv2.Mag()
	repulMag := (G * m1 * m2) / (distv2 * distv2)
	repulValv2.Normalize()
	repulValv2.Mult(repulMag)
	if repulsionv2[0] != repulValv2[0] || repulsionv2[1] != repulValv2[1] {
		t.Errorf("F = ((G * mass1 * mass2) / r ^ 2) * r^. expected=%v, got=%v", repulValv2, repulsionv2)
	}

	repulsionv3 := Repulsion(G, m1, m2, *v3Cases[0], *v3Cases[1])
	repulValv3 := vec.Sub(*v3Cases[0], *v3Cases[1])
	distv3 := repulValv3.Mag()
	repulMagv3 := (G * m1 * m2) / (distv3 * distv3)
	repulValv3.Normalize()
	repulValv3.Mult(repulMagv3)
	if repulsionv3[0] != repulValv3[0] || repulsionv3[1] != repulValv3[1] {
		t.Errorf("F = ((G * mass1 * mass2) / r ^ 2) * r^. expected=%v, got=%v", repulValv3, repulsionv3)
	}
}
