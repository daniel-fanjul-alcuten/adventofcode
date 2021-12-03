package p

var S1 = []int{
	0b_00100,
	0b_11110,
	0b_10110,
	0b_10111,
	0b_10101,
	0b_01111,
	0b_00111,
	0b_11100,
	0b_10000,
	0b_11001,
	0b_00010,
	0b_01010,
}

var X1 = []int{
	0b_011111111100,
	0b_100001011111,
	0b_011010010010,
	0b_100110110110,
	0b_001000110001,
	0b_110010001010,
	0b_110000010010,
	0b_010110100110,
	0b_000101000110,
	0b_100101010010,
	0b_101000100011,
	0b_001100011000,
	0b_111100100011,
	0b_000000110111,
	0b_001000000010,
	0b_000000001110,
	0b_000000001001,
	0b_010111010000,
	0b_100010100010,
	0b_010110100000,
	0b_111111101001,
	0b_001011100001,
	0b_111011110110,
	0b_011010110011,
	0b_101110011011,
	0b_010100111101,
	0b_010100000011,
	0b_101001011000,
	0b_000011011100,
	0b_101101010000,
	0b_100110110011,
	0b_101001111001,
	0b_011111001100,
	0b_010110000100,
	0b_100000111011,
	0b_111110011110,
	0b_110011000011,
	0b_101101101110,
	0b_111111100011,
	0b_111010110000,
	0b_101010110100,
	0b_100111010110,
	0b_101010100000,
	0b_011101000111,
	0b_111111010111,
	0b_110100001010,
	0b_011000111000,
	0b_010010011111,
	0b_110111001111,
	0b_000001110100,
	0b_010100010011,
	0b_100101100010,
	0b_100111101000,
	0b_010110100011,
	0b_101000000101,
	0b_000111100011,
	0b_110110000100,
	0b_001100101001,
	0b_110001110101,
	0b_111110101001,
	0b_100001100001,
	0b_101001000000,
	0b_101011011011,
	0b_001011111010,
	0b_101011101011,
	0b_010000101001,
	0b_001011010000,
	0b_010000010110,
	0b_111010101000,
	0b_000110010100,
	0b_000011110111,
	0b_110100001100,
	0b_110001100100,
	0b_110000001101,
	0b_101101111000,
	0b_110010111010,
	0b_011001011000,
	0b_100011010100,
	0b_010010101010,
	0b_011011111010,
	0b_100000000100,
	0b_000001001000,
	0b_011010010001,
	0b_000011110101,
	0b_110111011111,
	0b_101110101011,
	0b_010010000000,
	0b_000100101110,
	0b_000001001011,
	0b_110010101101,
	0b_100011100001,
	0b_010010111001,
	0b_000001011011,
	0b_001110100101,
	0b_001010000100,
	0b_110001010000,
	0b_110010010110,
	0b_000001111000,
	0b_101101001101,
	0b_111011010010,
	0b_000001100110,
	0b_111000101001,
	0b_011000001010,
	0b_011001010100,
	0b_001011000000,
	0b_101101111111,
	0b_010101010001,
	0b_101000010001,
	0b_111010100011,
	0b_000110000010,
	0b_000110110011,
	0b_001011000110,
	0b_100000010100,
	0b_100010001001,
	0b_101011100010,
	0b_101000100010,
	0b_101000000111,
	0b_111101000100,
	0b_100010111011,
	0b_010110000101,
	0b_111000000011,
	0b_000100111000,
	0b_110101110010,
	0b_101100010111,
	0b_101011011010,
	0b_111111110100,
	0b_001110010000,
	0b_011001111111,
	0b_001100011110,
	0b_000100101000,
	0b_101111010100,
	0b_110010111000,
	0b_010100000010,
	0b_100101100101,
	0b_101101010101,
	0b_000011010100,
	0b_000110010000,
	0b_010101111110,
	0b_011011001110,
	0b_011011110101,
	0b_010011111000,
	0b_111111000001,
	0b_000010001100,
	0b_110100000101,
	0b_001000111010,
	0b_100001000100,
	0b_100011001010,
	0b_000000001011,
	0b_011110000110,
	0b_010101000100,
	0b_100110110000,
	0b_000111010010,
	0b_010000100110,
	0b_101010101011,
	0b_001001010111,
	0b_111001101001,
	0b_100101000010,
	0b_110101101001,
	0b_011001110100,
	0b_100000001111,
	0b_000110000111,
	0b_101100100001,
	0b_111010011110,
	0b_101001110110,
	0b_100100001010,
	0b_100101101010,
	0b_000000001010,
	0b_111001001010,
	0b_011110100011,
	0b_001000000110,
	0b_000010010000,
	0b_000010001001,
	0b_001000001111,
	0b_101011100000,
	0b_011001011010,
	0b_101001101100,
	0b_100010101010,
	0b_011010111100,
	0b_001111000001,
	0b_000101010100,
	0b_011111011010,
	0b_100100001110,
	0b_111010111100,
	0b_010111010010,
	0b_001000001101,
	0b_101110110011,
	0b_010000100010,
	0b_111111111100,
	0b_001110001000,
	0b_001110000001,
	0b_011111111111,
	0b_000101100001,
	0b_011000100000,
	0b_011000011011,
	0b_011101001001,
	0b_101110100010,
	0b_111101110001,
	0b_111110110100,
	0b_110010000101,
	0b_111000000000,
	0b_110111010110,
	0b_111110011001,
	0b_111000001111,
	0b_101000110101,
	0b_100111011101,
	0b_011101001000,
	0b_100011100011,
	0b_101110111100,
	0b_110111101110,
	0b_011101110101,
	0b_011011011101,
	0b_010000001011,
	0b_000101001001,
	0b_000111001111,
	0b_011010101001,
	0b_101111010101,
	0b_000101111100,
	0b_111010110001,
	0b_011110100101,
	0b_011111000100,
	0b_010100100101,
	0b_111111110000,
	0b_011110101100,
	0b_110110100000,
	0b_011111111110,
	0b_101010101001,
	0b_110100110111,
	0b_011110110110,
	0b_011111001110,
	0b_010101001010,
	0b_111011001000,
	0b_101100100111,
	0b_110110110011,
	0b_101001111111,
	0b_010100010000,
	0b_000101101011,
	0b_000110110000,
	0b_000111000011,
	0b_111001010111,
	0b_011011100000,
	0b_001110111011,
	0b_111011111001,
	0b_101111010110,
	0b_111110100101,
	0b_011101001010,
	0b_110101010000,
	0b_011011000000,
	0b_010001010100,
	0b_010010101100,
	0b_111010001111,
	0b_100100110111,
	0b_111100110111,
	0b_000100000110,
	0b_111000100000,
	0b_100011001100,
	0b_101001110101,
	0b_001011010100,
	0b_110101100000,
	0b_101101011110,
	0b_000100011011,
	0b_111011000011,
	0b_110010000111,
	0b_111111110110,
	0b_010001001101,
	0b_110011010100,
	0b_111010110101,
	0b_100111111010,
	0b_010110111010,
	0b_111110000100,
	0b_001111101100,
	0b_110111001010,
	0b_101111011001,
	0b_010000111000,
	0b_111111111000,
	0b_100100111101,
	0b_101000001010,
	0b_011001000111,
	0b_000000011110,
	0b_011001011110,
	0b_111010000110,
	0b_001111111010,
	0b_100101100001,
	0b_011101011101,
	0b_101010110101,
	0b_101001101110,
	0b_010111111001,
	0b_001000111101,
	0b_001011001010,
	0b_101100111110,
	0b_001111010000,
	0b_011101101011,
	0b_101111111100,
	0b_000000110011,
	0b_101110110110,
	0b_000110011110,
	0b_101110111000,
	0b_100010001010,
	0b_001000101000,
	0b_001010110000,
	0b_010001101111,
	0b_101000011001,
	0b_001010001111,
	0b_000111000111,
	0b_001011101010,
	0b_111001001011,
	0b_101100110101,
	0b_000110111100,
	0b_001001100001,
	0b_101011011110,
	0b_110000001110,
	0b_110110111011,
	0b_001001000111,
	0b_101000100110,
	0b_000011000001,
	0b_011000001001,
	0b_001100110111,
	0b_000110010011,
	0b_111101111111,
	0b_110000010101,
	0b_000101111001,
	0b_100110011000,
	0b_001001011101,
	0b_101101110001,
	0b_000000100010,
	0b_000001010101,
	0b_110010110011,
	0b_010110111000,
	0b_000111001100,
	0b_101110100011,
	0b_001111101111,
	0b_000111010001,
	0b_110111000001,
	0b_011111101110,
	0b_110101101010,
	0b_001111001000,
	0b_101101001100,
	0b_001001101101,
	0b_100101101011,
	0b_000010001010,
	0b_011000110011,
	0b_011101111110,
	0b_000010000000,
	0b_101100110001,
	0b_000100110110,
	0b_100000111110,
	0b_100110111100,
	0b_000111110101,
	0b_100011000111,
	0b_100111000101,
	0b_100011001110,
	0b_011001110011,
	0b_010111000100,
	0b_101000001100,
	0b_011010010011,
	0b_110111011010,
	0b_110010100010,
	0b_011101100010,
	0b_111000001000,
	0b_001111010100,
	0b_110110110101,
	0b_011110000111,
	0b_011010000010,
	0b_100011010001,
	0b_000000111110,
	0b_011010011010,
	0b_000111000100,
	0b_000000001101,
	0b_101001101010,
	0b_011011100001,
	0b_110100001111,
	0b_000110110100,
	0b_011100010101,
	0b_101111011000,
	0b_001000011011,
	0b_101010001111,
	0b_100100000100,
	0b_101001001100,
	0b_110110111010,
	0b_110110100110,
	0b_101101100010,
	0b_111111101010,
	0b_000000011000,
	0b_111111110101,
	0b_111110101011,
	0b_111011111110,
	0b_101010001110,
	0b_101010001011,
	0b_110101001110,
	0b_000001010111,
	0b_011101001100,
	0b_000011010101,
	0b_110000110011,
	0b_011110110100,
	0b_011010000000,
	0b_011100011000,
	0b_100010101111,
	0b_000000111000,
	0b_011000010000,
	0b_010000010101,
	0b_011000010101,
	0b_010011111110,
	0b_111101000111,
	0b_110001000111,
	0b_001110101010,
	0b_010110111111,
	0b_001010000001,
	0b_011110101010,
	0b_101100001101,
	0b_011110011101,
	0b_001001100011,
	0b_100100001111,
	0b_100000100010,
	0b_001110101111,
	0b_111100111111,
	0b_100001001101,
	0b_011101001101,
	0b_000101111111,
	0b_100101011111,
	0b_010110001111,
	0b_011101111101,
	0b_110111011110,
	0b_111110111111,
	0b_110110111110,
	0b_001101101101,
	0b_001111100000,
	0b_101010100101,
	0b_101011011001,
	0b_110100011100,
	0b_011110010010,
	0b_000010011000,
	0b_001011100010,
	0b_000000100000,
	0b_101111011011,
	0b_110000001111,
	0b_010011101111,
	0b_011001010010,
	0b_001011110001,
	0b_110011011101,
	0b_101101110000,
	0b_101110010011,
	0b_000011110010,
	0b_011000010001,
	0b_011110001110,
	0b_101010110110,
	0b_011100001000,
	0b_000000111111,
	0b_111001010010,
	0b_011010101101,
	0b_001111011101,
	0b_000010001110,
	0b_000010000010,
	0b_011101101010,
	0b_010110100001,
	0b_000101000100,
	0b_110110101001,
	0b_011001000110,
	0b_100111001111,
	0b_100000101010,
	0b_010111001111,
	0b_011111011111,
	0b_010010111110,
	0b_110010110100,
	0b_010100110110,
	0b_011010101110,
	0b_111000000101,
	0b_001100001111,
	0b_010100111010,
	0b_001111111001,
	0b_110011111010,
	0b_000101010110,
	0b_110111111010,
	0b_000110110110,
	0b_001011000010,
	0b_011000111110,
	0b_010000110111,
	0b_000101010010,
	0b_011001001001,
	0b_100110100000,
	0b_110110011101,
	0b_000100110000,
	0b_000011011001,
	0b_111010001110,
	0b_100011100100,
	0b_011101010111,
	0b_110110010001,
	0b_100000001000,
	0b_100001010010,
	0b_001110101110,
	0b_100000000000,
	0b_100001111111,
	0b_001001011010,
	0b_111110000111,
	0b_111100101011,
	0b_011111110101,
	0b_011101010000,
	0b_000110100010,
	0b_110111001100,
	0b_111011010101,
	0b_010001011011,
	0b_001010101100,
	0b_011001010011,
	0b_000111010100,
	0b_100110101110,
	0b_001111001100,
	0b_011011010100,
	0b_000110111001,
	0b_110010001110,
	0b_000011000101,
	0b_001010100110,
	0b_001000110011,
	0b_010111110001,
	0b_100110001101,
	0b_000010011110,
	0b_111101101111,
	0b_100000001011,
	0b_101111000001,
	0b_110000001011,
	0b_111101010100,
	0b_011101011001,
	0b_010110001000,
	0b_001011111110,
	0b_111100110001,
	0b_011110100000,
	0b_100011101100,
	0b_011010100100,
	0b_110011110001,
	0b_111011101010,
	0b_101110101000,
	0b_011111111000,
	0b_010011101010,
	0b_110000010001,
	0b_111000101100,
	0b_101001111011,
	0b_001001000011,
	0b_100110010101,
	0b_000000011101,
	0b_001011001111,
	0b_000011100110,
	0b_000011110001,
	0b_101101100011,
	0b_110010010100,
	0b_110101100011,
	0b_000001111011,
	0b_101010111100,
	0b_010100000111,
	0b_001011110110,
	0b_111000010010,
	0b_111111001100,
	0b_111111111101,
	0b_100110111111,
	0b_101010001100,
	0b_001010011011,
	0b_110011101101,
	0b_011010011101,
	0b_110110110000,
	0b_000001011001,
	0b_111001100000,
	0b_101011101111,
	0b_111100111110,
	0b_010101000111,
	0b_110110101101,
	0b_011000110010,
	0b_011110110011,
	0b_110011100010,
	0b_111011010110,
	0b_100101010100,
	0b_100110011110,
	0b_101000111011,
	0b_100001101100,
	0b_011011010000,
	0b_011111110001,
	0b_100000011111,
	0b_101101010111,
	0b_111110000110,
	0b_000101101111,
	0b_001100110011,
	0b_011010001011,
	0b_101011001101,
	0b_010001000000,
	0b_011000101110,
	0b_111000000110,
	0b_110101010100,
	0b_110001110010,
	0b_001011101101,
	0b_111010101010,
	0b_101101010010,
	0b_010011111111,
	0b_000111011000,
	0b_100110010111,
	0b_111011111111,
	0b_101111010111,
	0b_101110011001,
	0b_011111100000,
	0b_111110001100,
	0b_101001001111,
	0b_011001010111,
	0b_010011011110,
	0b_101111011100,
	0b_011011110100,
	0b_111101111000,
	0b_001101101010,
	0b_100010110110,
	0b_110110111101,
	0b_011000110100,
	0b_011000100101,
	0b_111010100000,
	0b_010000100000,
	0b_110011111110,
	0b_000101000101,
	0b_111111011001,
	0b_110101100110,
	0b_100000110110,
	0b_010000011001,
	0b_000111111110,
	0b_011110011001,
	0b_111001101011,
	0b_100001111000,
	0b_100100101001,
	0b_111010100001,
	0b_110100110110,
	0b_011001100000,
	0b_111000100001,
	0b_000011100000,
	0b_111011000111,
	0b_101000111100,
	0b_001100011010,
	0b_111010101110,
	0b_101101001011,
	0b_111011001111,
	0b_000111001101,
	0b_101001001000,
	0b_110101100111,
	0b_111110010101,
	0b_011011000010,
	0b_001110101100,
	0b_011110000010,
	0b_010001100101,
	0b_110011001111,
	0b_011101011010,
	0b_101100111011,
	0b_101110010101,
	0b_101001000001,
	0b_111000110101,
	0b_100000111100,
	0b_110000100111,
	0b_001010100010,
	0b_011111101000,
	0b_001101001111,
	0b_010010010110,
	0b_000011101011,
	0b_111101010000,
	0b_001001000000,
	0b_001000111001,
	0b_110011010111,
	0b_011001100001,
	0b_110001100000,
	0b_111000011110,
	0b_010100101100,
	0b_100000000010,
	0b_101111110111,
	0b_000101011100,
	0b_111111011101,
	0b_011011000111,
	0b_000000111010,
	0b_001001110001,
	0b_101101100000,
	0b_100100011111,
	0b_110100110000,
	0b_001011011010,
	0b_011000110001,
	0b_011101100101,
	0b_011001101100,
	0b_011011011001,
	0b_010001111111,
	0b_111001011111,
	0b_000100001010,
	0b_110101010010,
	0b_101011001000,
	0b_001011011101,
	0b_000111110001,
	0b_001111000000,
	0b_100111000010,
	0b_101111110110,
	0b_010101111100,
	0b_101000111010,
	0b_011000111001,
	0b_111000001011,
	0b_011100011011,
	0b_101111011110,
	0b_111001011110,
	0b_100010111101,
	0b_010100100000,
	0b_001000001001,
	0b_110011010001,
	0b_011110001101,
	0b_001100110001,
	0b_110000011100,
	0b_111110111100,
	0b_001110001011,
	0b_100010000001,
	0b_100010000011,
	0b_011111100111,
	0b_000000010100,
	0b_111100011111,
	0b_011111100011,
	0b_100101011100,
	0b_011111000110,
	0b_110110010101,
	0b_000100000000,
	0b_010101101110,
	0b_100010110100,
	0b_011000000110,
	0b_001000100011,
	0b_001010101001,
	0b_111000010101,
	0b_001000001110,
	0b_011101001011,
	0b_110010011101,
	0b_001010000011,
	0b_011011011011,
	0b_100110100110,
	0b_101111001011,
	0b_110101100010,
	0b_111111000111,
	0b_011001010000,
	0b_001000110111,
	0b_101011100100,
	0b_110100101110,
	0b_111101110100,
	0b_111100101110,
	0b_110000101101,
	0b_101011001111,
	0b_100101001001,
	0b_000011101111,
	0b_010100111110,
	0b_011101101000,
	0b_010001010011,
	0b_110101110000,
	0b_010011101011,
	0b_111100101111,
	0b_010100001011,
	0b_001011100100,
	0b_010001000001,
	0b_111111111110,
	0b_100101000111,
	0b_010110111011,
	0b_001111101110,
	0b_100110001110,
	0b_010110101010,
	0b_001111001111,
	0b_001010011110,
	0b_010111101010,
	0b_001101110011,
	0b_101100000001,
	0b_001101100101,
	0b_111111011100,
	0b_010100101011,
	0b_101000011000,
	0b_100010100100,
	0b_101101010011,
	0b_100000111001,
	0b_110110100111,
	0b_011000111011,
	0b_101000011110,
	0b_000011100101,
	0b_001000110010,
	0b_001101000110,
	0b_011110111011,
	0b_000011111011,
	0b_011110011011,
	0b_100100101111,
	0b_011111000111,
	0b_000010101111,
	0b_101101001110,
	0b_101000001000,
	0b_111010001101,
	0b_000100110111,
	0b_001100000110,
	0b_101011111000,
	0b_111010111101,
	0b_010001011101,
	0b_001101010101,
	0b_111000000010,
	0b_100000111101,
	0b_011000101111,
	0b_010100011111,
	0b_010000101111,
	0b_011111011011,
	0b_000101100011,
	0b_010101000000,
	0b_100100111100,
	0b_110010111001,
	0b_011001100010,
	0b_011101111010,
	0b_000111110100,
	0b_011111101111,
	0b_100011011010,
	0b_100000110101,
	0b_110110110111,
	0b_010011101000,
	0b_001000000111,
	0b_101100001011,
	0b_100101001101,
	0b_101010101100,
	0b_110010101001,
	0b_001011000111,
	0b_101010010100,
	0b_101011100101,
	0b_010110000111,
	0b_010111110010,
	0b_000001001110,
	0b_000100101111,
	0b_011111001101,
	0b_101111111001,
	0b_000111111001,
	0b_111000001010,
	0b_011110000101,
	0b_101100110010,
	0b_010010001000,
	0b_111110001101,
	0b_110110011111,
	0b_101001010111,
	0b_111110101101,
	0b_010011111011,
	0b_110000110101,
	0b_011011110000,
	0b_101111111110,
	0b_000110001010,
	0b_011001001000,
	0b_010001111001,
	0b_110001111101,
	0b_100000000001,
	0b_011100011001,
	0b_001001101111,
	0b_101001011011,
	0b_111101110111,
	0b_101011000111,
	0b_010100001100,
	0b_010100101111,
	0b_010100100111,
	0b_101010111001,
	0b_010000011010,
	0b_100010010010,
	0b_100001011000,
	0b_110101101111,
	0b_000000111011,
	0b_011011101000,
	0b_101000000010,
	0b_011100011110,
	0b_100001001110,
	0b_000001010110,
	0b_110001010111,
	0b_011110100110,
	0b_111100011100,
	0b_001010001110,
	0b_001001001000,
	0b_000001100000,
	0b_001101011001,
	0b_010001001010,
	0b_000011001110,
	0b_001110001101,
	0b_011110100100,
	0b_101001010110,
	0b_111010010110,
	0b_001000111100,
	0b_101001010010,
	0b_010011111001,
	0b_111001001101,
	0b_111001011001,
	0b_011110010101,
	0b_100010110011,
	0b_111101000000,
	0b_110101010111,
	0b_101110010000,
	0b_011001110001,
	0b_011000101011,
	0b_100001100000,
	0b_000011010001,
	0b_001000000001,
	0b_010011111101,
	0b_111010111111,
	0b_000010110101,
	0b_101010011001,
	0b_101000110110,
	0b_010110001010,
	0b_101011010011,
	0b_101010000001,
	0b_101101111011,
	0b_101011101101,
	0b_111000111110,
	0b_000010000111,
	0b_110011010011,
	0b_101011110100,
	0b_110001011111,
	0b_001110011010,
	0b_010101110001,
	0b_011100001011,
	0b_110101011001,
	0b_100001101000,
	0b_101101000011,
	0b_110010100111,
	0b_111100010100,
	0b_010101011110,
	0b_001000110000,
	0b_110011111001,
	0b_100101101111,
	0b_001101001101,
	0b_100001011100,
	0b_101111000100,
	0b_011110001011,
	0b_110001001001,
	0b_011000111100,
	0b_110111100000,
	0b_011100000010,
	0b_000001011010,
	0b_001111110101,
	0b_110111011001,
	0b_011111111101,
	0b_110000010000,
	0b_110010110101,
	0b_110010001001,
	0b_010001000011,
	0b_110001000010,
	0b_001100010100,
	0b_001111000100,
	0b_110110100010,
	0b_110101101011,
	0b_001110001110,
	0b_000011011110,
	0b_111000100100,
	0b_101010011010,
	0b_011110001111,
	0b_011001011011,
	0b_001100101111,
	0b_101001001011,
	0b_100100001101,
	0b_010001101101,
	0b_010111111110,
	0b_110000000100,
	0b_111110111010,
	0b_011010011000,
	0b_100001010110,
	0b_100011000000,
	0b_110100010011,
	0b_010101111001,
	0b_100001011010,
	0b_100000111010,
	0b_111101110011,
	0b_100100111011,
	0b_011000011110,
	0b_010000110000,
	0b_100010111100,
	0b_101101110100,
	0b_001011110111,
	0b_001001000110,
	0b_001110110001,
	0b_001100001100,
	0b_100110100001,
	0b_010011001110,
	0b_010110001110,
	0b_001100111011,
	0b_101001111010,
	0b_010111001100,
	0b_011100110000,
	0b_111010000100,
	0b_101100100110,
	0b_100110101011,
	0b_111000010000,
	0b_001100111001,
	0b_001100100101,
	0b_000100000011,
	0b_111011010001,
	0b_100000000011,
	0b_010110000001,
	0b_000111110011,
	0b_001010111001,
	0b_111100011000,
	0b_111100001111,
	0b_011001000101,
	0b_001000110101,
	0b_100001001111,
	0b_010000100001,
	0b_010111100011,
	0b_110001110001,
	0b_000101111110,
	0b_101000000100,
	0b_100110000011,
	0b_000111100101,
	0b_101011110101,
	0b_101010101111,
	0b_100111000001,
	0b_001101001011,
	0b_111000101011,
	0b_111000100010,
	0b_100010001011,
	0b_111000111101,
	0b_110100000110,
	0b_100100101011,
	0b_010011101110,
	0b_110001100010,
}