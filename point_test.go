package goldilocks

import (
	"testing"

	"github.com/olabini/goldilocks/internal/field"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type pointSuite struct{}

var _ = Suite(&pointSuite{})

var testPointData1 = [][]uint64{
	{0x00e9f7e08f3ea0b6, 0x0079bb839e6cd664, 0x006db64036695b64, 0x00aa5d3fc7616622, 0x00857f0dc411550e, 0x002393d82efde209, 0x00f82748376fb2a0, 0x009efff9bceaeed3},
	{0x00ad6c3b508f7174, 0x001ed696f08213b1, 0x00f0686909fd2b0c, 0x00c5fe825432aa3d, 0x00d13ee025530425, 0x00558d8e4c4eada9, 0x00c9586080d08c17, 0x00a1326063a2bc4e},
	{0x004116747b4743de, 0x001a73752a3141b3, 0x004bcb1f9e57a702, 0x005170749ef95fd6, 0x00d0e7d7c372ecf4, 0x008284a2df2f3e3a, 0x00a363abed9f3b3a, 0x0028e9a599c15212},
	{0x0046bc190d1e1d1a, 0x0041c270b896cfbe, 0x0095971a20ce9265, 0x005807ec99bab3ca, 0x00412c12bda3e742, 0x0062318f598340b2, 0x007f750ca3b450d2, 0x001a3641722217fd},
}

var testPointData2 = [][]uint64{
	{0x00e9f7e08f3ea0b6, 0x0079bb839e6cd664, 0x006db64036695b64, 0x00aa5d3fc7616622, 0x00857f0dc411550e, 0x002393d82efde209, 0x00f82748376fb2a0, 0x009efff9bceaeed3},
	{0x00ad6c3b508f7174, 0x001ed696f08213b1, 0x00f0686909fd2b0c, 0x00c5fe825432aa3d, 0x00d13ee025530425, 0x00558d8e4c4eada9, 0x00c9586080d08c17, 0x00a1326063a2bc4e},
	{0x004116747b4743de, 0x001a73752a3141b3, 0x004bcb1f9e57a702, 0x005170749ef95fd6, 0x00d0e7d7c372ecf4, 0x008284a2df2f3e3a, 0x00a363abed9f3b3a, 0x0028e9a599c15212},
	{0x0046bc190d1e1d1a, 0x0041c270b896cfbe, 0x0095971a20ce9265, 0x005807ec99bab3ca, 0x00412c12bda3e742, 0x0062318f598340b2, 0x007f750ca3b450d2, 0x001a3641722217fd},
}

// testPointData3 = testPointData1 + testPointData2
var testPointData3 = [][]uint64{
	{0x001d859ce5314305, 0x00a80de5b42916ec, 0x007970f754771f30, 0x003ab63a14107f61, 0x0081f9bf23619f32, 0x0074300516f41504, 0x00bf80eb72e26901, 0x00c1a8a097ee58ca},
	{0x00e6637863317a54, 0x00318396a8682dcf, 0x00bc163928c29568, 0x00c561d09046898e, 0x009d8f55d40f377e, 0x0045874755f9a70a, 0x00392c8ffdb7373c, 0x00baae66240659f9},
	{0x00c42304039d747b, 0x001d9510c627343a, 0x00cd08f59aa09c0b, 0x000b80df2ed5d2f5, 0x00ef5bd55ac82343, 0x0072ac6706ee3768, 0x0078f9f95aee4aea, 0x002b63a4d93ffdf9},
	{0x00f71ad866015d1c, 0x00870711c1584289, 0x00ee4d7bdbfa883d, 0x00548824379834f1, 0x00ee911feaf8d241, 0x00fdb66827586951, 0x008f51e1a765503c, 0x0015f85a4e6a135d},
}

// point_add
// q = {
//     x = {0x00a3f760229c8330, 0x000abff01241b1c6, 0x0050ad4746887f47, 0x006f3c16c2e39094, 0x006a4d7d18c89659, 0x0031f3f205422228, 0x00f2aa87950a5d97, 0x00477c8d49cd90c3}
//     y = {0x00d9bc3116f547ca, 0x00ddd6e150ea845c, 0x00610da27a494d6b, 0x00fe126e00eacd75, 0x007f85b42c7ecc8a, 0x0088486a3659d71c, 0x00c6d87303fe9e83, 0x00564d0370012096}
//     z = {0x0075645a9b9be26d, 0x00ef5b0e12160f34, 0x001e7663b8429daf, 0x008e4ec330e5a9c3, 0x00b4ab6f6e6791f7, 0x00b85fc51ee87719, 0x0034f7a53dece326, 0x009d6092e3567b77}
//     t = {0x000e58af83d5c8a4, 0x00bb6c3f4b6ce5de, 0x001bb017c9b27e5c, 0x002f7b1abc0ac1c7, 0x00a7b509fe1560b3, 0x0027a399020bfaa7, 0x00ddc9c305a6616d, 0x000110cb607d19af}
// } // q
// r = {
//     x = {0x0096140cfd15f6d2, 0x00c46b8574f9eef4, 0x004b5e747ecdf050, 0x00dd788b7edd7bcd, 0x00704d53eab37b0e, 0x00e0aa7602f23570, 0x007db7aded4a082b, 0x00748ee73db512c0}
//     y = {0x0054db083996a145, 0x00255a9abbc3b56e, 0x00509f5f76c153f3, 0x0079d0117810f5c7, 0x004261fe7c4622d1, 0x00b0e55119948424, 0x00a786ba482703ac, 0x0069c739cc61f3d0}
//     z = {0x00fd4e10cda7083a, 0x00007cf35531cb68, 0x00d89be786066020, 0x0016141806a94f98, 0x0082e0a397950982, 0x00edf4ae6a27c078, 0x00f9f0763326ed8a, 0x009805739cb19823}
//     t = {0x00c32155683cf5a5, 0x007667016d9a5b56, 0x002bac50664d0374, 0x000163ee2a153dc1, 0x004ac19e94acf439, 0x0085ed1b37c73a23, 0x00773c44b9213219, 0x00fab8c5a0b05ce5}
// } // r
// p = {
//     x = {0x00feaf36075a1d55, 0x005afa5178a89fa8, 0x00f11051892e397f, 0x0070085af8e6fb1a, 0x00ef6dfcb431b814, 0x00374f9d0f9d52d6, 0x00a782e6e30d7ec6, 0x00b3f2eb5b87b8ed}
//     y = {0x0019e8336bc0cc7c, 0x000e189a56d66a41, 0x0014bebc6f9f9a8e, 0x0088b9610a2fe87b, 0x001e1a090b72d899, 0x00d3c5b9d5cfaba2, 0x00613f3b9d78bf9e, 0x0081ae9b71829559}
//     z = {0x006b7ea1b5330778, 0x0034f861de60fd63, 0x0096ab372e057eb0, 0x00768311f724448c, 0x00467b904a0b46aa, 0x00f7f11d2d98c8f1, 0x0052b39817d3b6d5, 0x00b55df8ea4cef70}
//     t = {0x00b9411e6e68a43c, 0x001cfd807523bd1c, 0x00396140591f5d8a, 0x002a0b195f15023e, 0x0037d4c817d15bb0, 0x008a9fdcd3400654, 0x0012da0cb125f9d5, 0x00544d82c39409eb}
// } // p
// point_add
// q = {
//     x = {0x00e9f7e08f3ea0b6, 0x0079bb839e6cd664, 0x006db64036695b64, 0x00aa5d3fc7616622, 0x00857f0dc411550e, 0x002393d82efde209, 0x00f82748376fb2a0, 0x009efff9bceaeed3}
//     y = {0x00ad6c3b508f7174, 0x001ed696f08213b1, 0x00f0686909fd2b0c, 0x00c5fe825432aa3d, 0x00d13ee025530425, 0x00558d8e4c4eada9, 0x00c9586080d08c17, 0x00a1326063a2bc4e}
//     z = {0x004116747b4743de, 0x001a73752a3141b3, 0x004bcb1f9e57a702, 0x005170749ef95fd6, 0x00d0e7d7c372ecf4, 0x008284a2df2f3e3a, 0x00a363abed9f3b3a, 0x0028e9a599c15212}
//     t = {0x0046bc190d1e1d1a, 0x0041c270b896cfbe, 0x0095971a20ce9265, 0x005807ec99bab3ca, 0x00412c12bda3e742, 0x0062318f598340b2, 0x007f750ca3b450d2, 0x001a3641722217fd}
// } // q
// r = {
//     x = {0x00feaf36075a1d55, 0x005afa5178a89fa8, 0x00f11051892e397f, 0x0070085af8e6fb1a, 0x00ef6dfcb431b814, 0x00374f9d0f9d52d6, 0x00a782e6e30d7ec6, 0x00b3f2eb5b87b8ed}
//     y = {0x0019e8336bc0cc7c, 0x000e189a56d66a41, 0x0014bebc6f9f9a8e, 0x0088b9610a2fe87b, 0x001e1a090b72d899, 0x00d3c5b9d5cfaba2, 0x00613f3b9d78bf9e, 0x0081ae9b71829559}
//     z = {0x006b7ea1b5330778, 0x0034f861de60fd63, 0x0096ab372e057eb0, 0x00768311f724448c, 0x00467b904a0b46aa, 0x00f7f11d2d98c8f1, 0x0052b39817d3b6d5, 0x00b55df8ea4cef70}
//     t = {0x00b9411e6e68a43c, 0x001cfd807523bd1c, 0x00396140591f5d8a, 0x002a0b195f15023e, 0x0037d4c817d15bb0, 0x008a9fdcd3400654, 0x0012da0cb125f9d5, 0x00544d82c39409eb}
// } // r
// p = {
//     x = {0x00eb394f0e209270, 0x000ad75267e5ed29, 0x0042f454def4485a, 0x005fc37369e5c98b, 0x002a83a6d8b176a0, 0x00d5124b8791d0a2, 0x00c51ccf64c5619e, 0x00f2749da0c23723}
//     y = {0x008fe5183f333198, 0x0085db45b707103e, 0x00d04cce5038f249, 0x00ff210653620be0, 0x001b92cb290e69fb, 0x00dbab5961be18c5, 0x00280e78df15e19c, 0x000991e1f07835c8}
//     z = {0x0088e9c811aea09f, 0x00af3a7149e85b39, 0x00eadd9992d708ac, 0x0055c3fa6d4a1979, 0x00cb95b31198df6a, 0x0092c37dc27b01d1, 0x00c3e0a1902ea0ee, 0x007689ca3e80b8dc}
//     t = {0x00bacbfac091f422, 0x0098f217bfde91ac, 0x00f9a64317481b28, 0x00321c0d022adf1a, 0x0041ca448f827cc9, 0x006cff0d6a8f670b, 0x00aedf4441e0b4e4, 0x00ccc1b1e2ace46b}
// } // p

func createElementFrom(d []uint64) *field.Element {
	return field.CreateElementFrom(d)
}

func createPointFrom(d [][]uint64) *point {
	return &point{
		x: createElementFrom(d[0]),
		y: createElementFrom(d[1]),
		z: createElementFrom(d[2]),
		t: createElementFrom(d[3]),
	}
}

func (s *pointSuite) Test_pointAdd(c *C) {
	q := createPointFrom(testPointData1)
	r := createPointFrom(testPointData2)
	p := createPointFrom(testPointData3)
	defer q.destroy()
	defer r.destroy()
	defer p.destroy()

	res := newPoint()
	defer res.destroy()

	pointAdd(res, q, r)

	c.Assert(res, DeepEquals, p)
}
