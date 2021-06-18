package common

import (
	"testing"

	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	"github.com/stretchr/testify/assert"

	st "github.com/srjnm/demoinfocs-golang/pkg/demoinfocs/sendtables"
	stfake "github.com/srjnm/demoinfocs-golang/pkg/demoinfocs/sendtables/fake"
)

func TestInferno_UniqueID(t *testing.T) {
	entity := new(stfake.Entity)
	assert.NotEqual(t, NewInferno(nil, entity).UniqueID(), NewInferno(nil, entity).UniqueID(), "UniqueIDs of different infernos should be different")
}

func TestFires_Active(t *testing.T) {
	inf := Fires{
		s: []Fire{
			{
				IsBurning: false,
				Vector:    r3.Vector{X: 1, Y: 2, Z: 3},
			},
		},
	}

	assert.Empty(t, inf.Active().s, "Inferno should have no active fires")

	activeFires := []Fire{
		{
			IsBurning: true,
			Vector:    r3.Vector{X: 4, Y: 5, Z: 6},
		},
		{
			IsBurning: true,
			Vector:    r3.Vector{X: 7, Y: 8, Z: 9},
		},
	}
	inf.s = append(inf.s, activeFires...)

	assert.Equal(t, activeFires, inf.Active().s, "Active inferno should contain active fires")
}

func TestInferno_ConvexHull2D(t *testing.T) {
	// Construct a Inferno that looks roughly like this.
	// D should be inside the 2D Convex Hull but a corner of the 3D Convex Hull
	//
	//         C
	//       /   \
	//     /   D   \
	//   /           \
	// A - - - - - - - B
	//
	inf := Fires{
		s: []Fire{
			{
				Vector: r3.Vector{X: 1, Y: 2, Z: 3},
			},
			{
				Vector: r3.Vector{X: 4, Y: 7, Z: 6},
			},
			{
				Vector: r3.Vector{X: 7, Y: 2, Z: 9},
			},
			{
				Vector: r3.Vector{X: 4, Y: 4, Z: 12}, // This fire is inside the 2D hull
			},
		},
	}

	expectedHull := []r2.Point{
		{X: 1, Y: 2},
		{X: 4, Y: 7},
		{X: 7, Y: 2},
	}

	assert.ElementsMatch(t, expectedHull, inf.ConvexHull2D(), "ConvexHull2D should be as expected")

	// 3D-hull should be different
	assert.NotEqual(t, len(expectedHull), len(inf.ConvexHull3D().Vertices), "3D hull should contain the vertex 'D'")
}

// Just check that all fires are passed to quickhull.ConvexHull()
func TestInferno_ConvexHull3D(t *testing.T) {
	inf := Fires{
		s: []Fire{
			{
				Vector: r3.Vector{X: 1, Y: 2, Z: 3},
			},
			{
				Vector: r3.Vector{X: 4, Y: 7, Z: 6},
			},
			{
				Vector: r3.Vector{X: 7, Y: 2, Z: 9},
			},
			{
				Vector: r3.Vector{X: 4, Y: 4, Z: 12},
			},
		},
	}

	expectedHull := []r3.Vector{
		{X: 1, Y: 2, Z: 3},
		{X: 4, Y: 7, Z: 6},
		{X: 7, Y: 2, Z: 9},
		{X: 4, Y: 4, Z: 12},
	}

	assert.ElementsMatch(t, expectedHull, inf.ConvexHull3D().Vertices, "ConvexHull3D should contain all fire locations")
}

func TestInferno_Thrower(t *testing.T) {
	entity := entityWithProperty("m_hOwnerEntity", st.PropertyValue{IntVal: 1})

	player := new(Player)
	provider := demoInfoProviderMock{
		playersByHandle: map[int]*Player{1: player},
	}

	assert.Equal(t, player, NewInferno(provider, entity).Thrower())
}
