package reactorreboot

func expandOverlapFaces(overlap []bool, a, smaller *Cuboid) []*Cuboid {
	out := []*Cuboid{}
	if overlap[0] {
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: smaller.z0,
			z1: smaller.z1,
		})
	}
	if overlap[1] {
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: smaller.z0,
			z1: smaller.z1,
		})
	}
	if overlap[2] {
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: smaller.z0,
			z1: smaller.z1,
		})
	}
	if overlap[3] {
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: smaller.z0,
			z1: smaller.z1,
		})
	}
	if overlap[4] {
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[5] {
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[0] && overlap[2] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: smaller.z1,
			z1: smaller.z1,
		})
	}
	if overlap[0] && overlap[3] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: smaller.z1,
			z1: smaller.z1,
		})
	}
	if overlap[1] && overlap[2] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: smaller.z1,
			z1: smaller.z1,
		})
	}
	if overlap[1] && overlap[3] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1,
			x1: a.x1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: smaller.z1,
			z1: smaller.z1,
		})
	}
	if overlap[0] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[0] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[1] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[1] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: smaller.y0,
			y1: smaller.y1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[2] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[2] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[3] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[3] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x0,
			x1: smaller.x1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[0] && overlap[2] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[0] && overlap[2] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[0] && overlap[3] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[0] && overlap[3] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: a.x0,
			x1: smaller.x0 - 1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[1] && overlap[2] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[1] && overlap[2] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: a.y0,
			y1: smaller.y0 - 1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	if overlap[1] && overlap[3] && overlap[4] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: a.z0,
			z1: smaller.z0 - 1,
		})
	}
	if overlap[1] && overlap[3] && overlap[5] {
		// x0 y0
		out = append(out, &Cuboid{
			on: true,
			x0: smaller.x1 + 1,
			x1: a.x1,
			y0: smaller.y1 + 1,
			y1: a.y1,
			z0: smaller.z1 + 1,
			z1: a.z1,
		})
	}
	return out
}
