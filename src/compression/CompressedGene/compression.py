class CompressedGene:
    def __init__(self, gene: str) -> None:
        self.__compress(gene)

    def __compress(self, gene: str) -> None:
        self.bit_string: int = 1 # begin label
        for nucleotide in gene.upper():
            self.bit_string <<= 2 # shift to the left by 2 bits
            if nucleotide == "A": # change 2 last bits by 00
                self.bit_string |= 0b00
            elif nucleotide == "C": # change 2 last bits by 01
                self.bit_string |= 0b01
            elif nucleotide == "G": # change 2 last bits by 10
                self.bit_string |= 0b10
            elif nucleotide == "T": # change 2 last bits by 11
                self.bit_string |= 0b11
            else:
                raise ValueError("Invalid Nucleotide:{}".format(nucleotide))

    def decompress(self) -> str:
        gene: str = ""
        for i in range(0, self.bit_string.bit_length() - 1, 2):
            bits: int = self.bit_string >> i & 0b11
            # get only 2 significant bits
            if bits == 0b00: # A
                gene += "A"
            elif bits == 0b01: # C
                gene += "C"
            elif bits == 0b10: # G
                gene += "G"
            elif bits == 0b11: # T
                gene += "T"
            else:
                raise ValueError("Invalid bits:{}".format(bits))
        return gene[::-1]

    def __str__(self) -> str:
        return self.decompress()


if __name__ == '__main__':
    from sys import getsizeof

    original: str = "TAGGGGATTAACCCGTTAATATATATATATGGATTAACCCGTTAATATATATATAGGATTAACCCGTTAATATATATATA" * 100
    print("original is {} bytes".format(getsizeof(original)))

    compressed: CompressedGene = CompressedGene(original)
    print("compressed is {} byte".format(getsizeof(compressed.bit_string)))
    print(compressed)
    print("original and decompressed are the same: {}".format(original == compressed.decompress()))

